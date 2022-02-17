package main

import (
	"context"
	"fmt"
	"go-scaffold/controllers"
	"go-scaffold/dao/cache"
	"go-scaffold/dao/rds"
	"go-scaffold/pkg/configs"
	"go-scaffold/pkg/logger"
	"go-scaffold/pkg/snowflake"
	"go-scaffold/router"
	"go-scaffold/tasks"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"github.com/urfave/cli/v2"
)

var (
	app *cli.App
)

// 初始化函数
func init() {
	app = cli.NewApp()
	// APP 的名称
	app.Name = "go-scaffold"
	// APP 的作者
	app.Authors = []*cli.Author{
		{Name: "wangyuxiao", Email: "xiao.950901@gmail.com"},
	}
	// APP 的版权
	app.Copyright = "©2021-2021 XiaoCoder Corporation,All Rights Reserved"
	// APP 的版本
	app.Version = "0.0.1"
}

// 主函数
func main() {
	// 设置随机数的随机种子，保证随机性
	rand.Seed(time.Now().UnixNano())

	// 注册服务
	app.Commands = []*cli.Command{
		{
			Name:  "api",
			Usage: "API Service",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "conf",
					Aliases:  []string{`c`},
					Usage:    "指定配置文件",
					Value:    "",
					Required: true,
				},
			},
			Action: func(c *cli.Context) error {
				if err := api(c.String("conf")); err != nil {
					log.Printf("failed to start api service, err: %e\n", err)
					return cli.Exit(err.Error(), 1)
				}
				return nil
			},
		},
	}

	// 运行服务
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("Service failed to start, err: [%s]\n", err)
	}
}

// API Service
func api(configFile string) error {
	// 解析配置文件
	if err := configs.ParseConfigFile(configFile); err != nil {
		log.Printf("failed to parse configuration file, err: %e\n", err)
		return err
	}

	// 初始化日志对象
	if err := logger.Init("api"); err != nil {
		log.Printf("failed to initialize logger, err: %e\n", err)
		return err
	}
	// 延迟注册：将缓存区的日志追加到日志文件中
	defer logger.Sync()

	// 初始化分布式ID生成器
	if err := snowflake.Init(); err != nil {
		return err
	}

	// 初始化数据库连接
	if err := rds.Init(); err != nil {
		zap.L().Fatal("failed to initialize rds service", zap.Error(err))
		return err
	}
	// 延迟注册：关闭数据库连接
	defer rds.Close()

	// 初始化缓存连接
	if err := cache.Init(); err != nil {
		zap.L().Fatal("failed to initialize cache service", zap.Error(err))
		return err
	}
	// 延迟注册：关闭缓存连接
	defer cache.Close()

	// TODO: 初始化ES（ElasticSearch）连接

	// 定时任务
	go tasks.Init()

	// 初始化参数校验器使用的翻译器
	if err := controllers.InitTrans("zh"); err != nil {
		zap.L().Fatal("failed to initialize trans", zap.Error(err))
	}

	// 注册路由
	engine := router.Setup()

	// 运行服务（优雅关闭）
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", configs.AllConfig.Basic.Host, configs.AllConfig.Basic.Port),
		Handler: engine,
	}

	// 通过协程启动服务
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("failed to start api service", zap.Error(err))
		}
	}()

	// 等待中断信号来优雅地关闭服务，为关闭服务操作设置一个5秒的超时
	// 创建一个接收信号的通道
	quit := make(chan os.Signal, 1)

	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	// 此处不会阻塞
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 阻塞在此，当接收到上述两种信号时才会往下执行
	<-quit

	// 接收到信号，关闭服务
	zap.L().Info("the service will shut down in 5 seconds")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := server.Shutdown(ctx); err != nil {
		zap.L().Fatal("failed to close service", zap.Error(err))
	}
	zap.L().Info("the service has exited normally")
	return nil
}
