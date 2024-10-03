package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"zinx/ziface"
)

/*
	对外暴露的全局变量
*/

type GlobalObj struct {
	/*
		Server
	*/
	TcpServer ziface.IServer
	Host      string
	TcpPort   int
	Name      string

	/*
		zinx框架的一些配置
	*/
	Version          string
	MaxConn          int
	MaxPackageSize   uint32
	WorkerPoolSize   uint32 // 业务工作 Worker 池的数量
	MaxWorkerTaskLen uint32 // 业务工作 Worker 对应负责的任务队列最大任务存储数量
}

func (o *GlobalObj) Reload() {
	// 从配置文件中加载一些参数
	data, err := os.ReadFile("conf/zinx.json")
	if err != nil {
		fmt.Println("os.ReadFile err:", err)
		return
	}
	// 将 json 文件解析到 GlobalObj 中
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		fmt.Println("json.Unmarshal err:")
		return
	}
}

// GlobalObject 对外暴露的的全局
var GlobalObject *GlobalObj

// init 初始化方法
func init() {
	GlobalObject = &GlobalObj{
		Name:             "ZinxServerApp",
		Version:          "V0.4",
		TcpPort:          8999,
		Host:             "0.0.0.0",
		MaxConn:          1000,
		MaxPackageSize:   4096,
		WorkerPoolSize:   10,   // 当前业务 Worker 池的数量
		MaxWorkerTaskLen: 1024, // 每个 Worker 对应的消息队列的任务的数量最大值
	}

	// 从配置文件中加载一些参数
	GlobalObject.Reload()
}
