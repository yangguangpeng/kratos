package sentry

import (
	"errors"
	"fmt"
	"github.com/getsentry/sentry-go"
	"helloworld/pkg/global/envFlag"
	"log"
	"math/rand"
	"os"
	"time"
)

var InitSentryInstance *InitSentry

type InitSentry struct {
	ProjectName    string
	Dsn            string //启动端口
	FlushTimeOut   time.Duration
	HourSetToCheck int

	systemQuit chan int
}

func (config *InitSentry) Init(message string) {
	return
	InitSentryInstance = config

	config.systemQuit = make(chan int, 1)

	if config.Dsn == "" {
		log.Fatalf("请指定sentry的dsn")
	}
	err := sentry.Init(sentry.ClientOptions{
		Dsn: config.Dsn,
	})

	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	// Flush buffered events before the program terminates.
	if config.FlushTimeOut == 0 {
		//每二秒刷新一次
		config.FlushTimeOut = 2 * time.Second
	}

	defer sentry.Flush(config.FlushTimeOut)

	if message == "" {
		message = "可以自定义一些消息"
	}

	dir, _ := os.Getwd()

	message = fmt.Sprintf("Sentry Works! File Path = %s, Message = %s", dir, message)

	//每天每个项目发出一次报警，表示配置的sentry正在正常的工作中。

	if envFlag.Instance.IsEnvPro() || envFlag.Instance.IsUnitTestMode() {

		go func() {

			for {

				//发完第一次的消息以后，后面是每天发一次
				sentry.CaptureMessage(message)

				now := time.Now()
				next := now.Add(time.Hour * 24)
				locShanghai, _ := time.LoadLocation("Asia/Shanghai")

				rand.Seed(time.Now().Unix()) //Seed生成的随机数
				mySec := rand.Intn(30)
				myNSec := rand.Intn(100)

				next = time.Date(next.Year(), next.Month(), next.Day(), config.HourSetToCheck, 0, mySec, myNSec, locShanghai) //设置每天北京时间8点发消息
				t := time.NewTimer(next.Sub(now))
				select {
				case <-t.C:
					continue
				case <-config.systemQuit:
					log.Println("Sentry退出...")
					return
				}

			}
		}()
	}

}

func (config *InitSentry) Quit() {
	config.systemQuit <- 1
}

func (config *InitSentry) checkConfig() bool {
	if config == nil {
		log.Printf("sentry un-inited!")
		return false
	}
	return true
}

func (config *InitSentry) Exception(errorInfo error, extend string) {
	//仅在生产环境纪录错误
	if errorInfo != nil && config.checkConfig() && !envFlag.Instance.IsEnvPro() {
		sentry.CaptureException(errors.New(fmt.Sprintf("【%s-%s】%s", extend, config.ProjectName, errorInfo.Error())))
	}
}

func (config *InitSentry) Info(info interface{}, extend string) {
	//仅在生产环境纪录错误
	if !envFlag.Instance.IsEnvPro() && config.checkConfig() {
		sentry.CaptureMessage(fmt.Sprintf("【%s-%s】%+v", extend, config.ProjectName, info))
	}
}
