package tasks

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

func Init() {
	// 秒字段：可选
	taskServer := cron.New(cron.WithParser(cron.NewParser(
		cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
	)))

	// 添加任务
	taskServer.AddFunc("0 10 * * *", goWork)
	taskServer.AddFunc("@every 30m", goRest)
	taskServer.AddFunc("0 20 * * *", offWork)

	// 启动服务
	taskServer.Start()
}

func goWork() {
	fmt.Println("Go work")
}

func goRest() {
	fmt.Println("It’s been 30 minutes, it’s time to rest")
}

func offWork() {
	fmt.Println("Off work")
}
