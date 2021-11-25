package provider

import (
	"fmt"
	"path/filepath"

	"github.com/pkg/errors"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

// RotateLog 代表会进行切割的日志文件存储
type RotateLog struct {
	PingLog
	// 日志文件存储目录
	folder string
	// 日志文件名
	file string
}

func NewRotateLog(params []interface{}) (interface{}, error) {
	level := params[0].(LogLevel)
	ctxFielder := params[1].(CtxFielder)
	formatter := params[2].(Formatter)

	file := "ping.log"
	linkName := rotatelogs.WithLinkName(file)

	options := []rotatelogs.Option{linkName}
	rotateCount := 2
	options = append(options, rotatelogs.WithRotationCount(uint(rotateCount)))

	log := &RotateLog{}
	log.SetLevel(level)
	log.SetCtxFielder(ctxFielder)
	log.SetFormatter(formatter)
	log.folder = "./"
	log.file = file
	dateFormat := "%Y%m%d%H%M"
	w, err := rotatelogs.New(
		fmt.Sprintf("%s.%s", filepath.Join(log.folder, log.file), dateFormat), options...)
	if err != nil {
		return nil, errors.Wrap(err, "new rotatelogs error")
	}

	log.SetOutput(w)
	return log.PingLog, nil
}
