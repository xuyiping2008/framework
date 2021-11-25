package provider

import (
	"fmt"
	"io"
	"strings"
)

// LogServiceProvider 服务提供者
type LogServiceProvider struct {

	//framework.ServiceProvider
	Driver string // Driver
	// 日志级别
	Level LogLevel
	// 日志输出格式方法
	Formatter Formatter
	// 日志context上下文信息获取函数
	CtxFielder CtxFielder
	// 日志输出信息
	Output io.Writer
}

// Register 注册一个服务实例
func (l *LogServiceProvider) Register() NewLog {
	if l.Driver == "" {
		l.Driver = strings.ToLower("console")
	}

	if l.CtxFielder == nil {
		l.CtxFielder = func(i interface{}) map[string]interface{} {
			return map[string]interface{}{}
		}
	}
	// 根据driver的配置项确定
	switch l.Driver {
	case "single":
		return NewSingleLog
	case "rotate":
		return NewRotateLog
	case "console":
		fmt.Println("console-log")
		return NewConsoleLog
	case "custom":
		return NewCustomLog
	default:
		return NewConsoleLog
	}
}

func (l *LogServiceProvider) Params() []interface{} {
	if l.Formatter == nil {
		l.Formatter = TextFormatter
	}

	if l.Level == UnknownLevel {
		l.Level = InfoLevel
	}

	return []interface{}{l.Level, l.CtxFielder, l.Formatter, l.Output}
}
