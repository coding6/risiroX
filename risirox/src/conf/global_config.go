package conf

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

type GlobalConfig struct {
	//监听的ip
	Host string
	//监听的端口
	Port int
	//允许的最大链接数
	MaxConn int
	//数据包的大小
	MaxPackageSize uint32
	//工作池的工作线程数
	WorkerPoolSize uint64
	//工作池的最大线程数
	MaxWorkerPoolSize uint32
	//任务队列的最大长度
	MaxWorkerTaskQueueLen uint32
}

var GlobalConfigObj *GlobalConfig

func (g *GlobalConfig) Load() {
	file, err := ioutil.ReadFile("../conf/conf.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(file, &g)
	if err != nil {
		panic(err)
	}
}

func initLog() {
	// 设置日志格式为 JSON 格式
	log.SetFormatter(&log.JSONFormatter{})

	// 设置输出到标准输出
	log.SetOutput(os.Stdout)

	// 设置日志级别为 Debug 级别
	log.SetLevel(log.DebugLevel)
}

func init() {
	GlobalConfigObj = &GlobalConfig{
		Host:                  "0.0.0.0",
		Port:                  8999,
		MaxConn:               1000,
		MaxPackageSize:        4096,
		WorkerPoolSize:        10,
		MaxWorkerTaskQueueLen: 1024,
		MaxWorkerPoolSize:     15,
	}
	GlobalConfigObj.Load()
	initLog()
}
