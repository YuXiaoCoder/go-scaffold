package logger

import (
	"fmt"
	"go-scaffold/pkg/configs"
	"log"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 初始化日志对象
func Init(loggerName string) error {
	// 获取日志编码器
	encoder := getEncoder()

	// 根据配置文件设置日志级别
	level := new(zapcore.Level)
	if err := level.UnmarshalText([]byte(configs.AllConfig.Logger.Level)); err != nil {
		return err
	}

	// 设置日志核心
	var core zapcore.Core
	if *level == zap.InfoLevel {
		// 获取不同日志级别的输出流
		infoWriteSyncer := getWriteSyncer(fmt.Sprintf("%s/%s-info.log", configs.AllConfig.Logger.Directory, loggerName))
		errorWriteSyncer := getWriteSyncer(fmt.Sprintf("%s/%s-error.log", configs.AllConfig.Logger.Directory, loggerName))

		core = zapcore.NewTee(
			zapcore.NewCore(encoder, infoWriteSyncer, zap.InfoLevel),
			zapcore.NewCore(encoder, errorWriteSyncer, zap.ErrorLevel),
		)
	} else {
		// 获取指定日志级别的输出流
		writeSyncer := getWriteSyncer(fmt.Sprintf("%s/%s.log", configs.AllConfig.Logger.Directory, loggerName))

		core = zapcore.NewCore(encoder, writeSyncer, level)
	}

	// 全局替换Zap的日志对象，添加调用者信息和行数
	zap.ReplaceGlobals(zap.New(core, zap.AddCaller()))
	return nil
}

// 获取输出流
func getWriteSyncer(filename string) zapcore.WriteSyncer {
	hook, err := rotatelogs.New(
		// 替换日志文件名，采用以log为后缀
		strings.Replace(filename, ".log", "", -1)+".%Y%m%d.log",
		// 生产软链接文件
		rotatelogs.WithLinkName(filename),
		// 切割日志文件的间隔
		rotatelogs.WithRotationTime(time.Duration(configs.AllConfig.Logger.RotationTime)*time.Hour),
		// 等待清理旧日志的时间，此配置为禁用清理
		rotatelogs.WithMaxAge(-1),
		// 保留日志文件的个数
		rotatelogs.WithRotationCount(configs.AllConfig.Logger.RotationCount),
	)

	if err != nil {
		log.Printf("failed to create rotatelogs: %s", err)
		return nil
	}

	// Zap底层设置了缓存，此方法用于将缓存同步到文件中
	return zapcore.AddSync(hook)
}

// 获取日志编码器
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()

	// 设置时间格式
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	// 使用大写字母记录日志级别
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// JSON 编码器
	return zapcore.NewJSONEncoder(encoderConfig)
}

// Sync 将缓存区的日志追加到日志文件中
func Sync() {
	_ = zap.L().Sync()
}
