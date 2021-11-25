package provider

import (
	"fmt"
	"os"
)

// ConsoleLog 代表控制台输出
type ConsoleLog struct {
	PingLog
}

//NewConsoleLog 表示控制台输出，定义初始化实例方法 ；
func NewConsoleLog(params []interface{}) (interface{}, error) {

	level := params[0].(LogLevel)
	ctxFielder := params[1].(CtxFielder)
	formatter := params[2].(Formatter)

	log := &ConsoleLog{}
	log.SetLevel(level)
	log.SetCtxFielder(ctxFielder)
	log.SetFormatter(formatter)

	// 最重要的将内容输出到控制台
	log.SetOutput(os.Stdout)

	fmt.Println("formatter-log", fmt.Sprintf("%T", log.PingLog))
	return log.PingLog, nil
}
