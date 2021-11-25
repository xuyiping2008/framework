package provider

import (
	"io"
	pkgLog "log"
	"time"
)

type PingLog struct {
	Level      LogLevel
	Formatter  Formatter
	CtxFielder CtxFielder
	Output     io.Writer
}

func (p *PingLog) IsLevelEnable(level LogLevel) bool {
	return level <= p.Level
}

func (p *PingLog) Info(msg string, fields map[string]interface{}) {
	p.logf(InfoLevel, msg, fields)
}

func (p *PingLog) logf(level LogLevel, msg string, fields map[string]interface{}) error {
	if !p.IsLevelEnable(level) {
		return nil
	}

	fs := fields
	if p.CtxFielder != nil {
		/*t := p.ctxFielder(ctx)
		if t != nil {
			for k, v := range t {
				fs[k] = v
			}
		}*/
	}

	if p.Formatter == nil {
		p.Formatter = TextFormatter
	}

	ct, err := p.Formatter(level, time.Now(), msg, fs)
	if err != nil {
		return err
	}

	if level == PanicLevel {
		pkgLog.Panicln(string(ct))
		return nil
	}

	p.Output.Write(ct)
	p.Output.Write([]byte("\r\n"))
	return nil
}

// SetOutput 设置output
func (p *PingLog) SetOutput(output io.Writer) {
	p.Output = output
}

// Panic 输出panic的日志信息
func (p *PingLog) Panic(msg string, fields map[string]interface{}) {
	p.logf(PanicLevel, msg, fields)
}

// Fatal will add fatal record which contains msg and fields
func (p *PingLog) Fatal(msg string, fields map[string]interface{}) {
	p.logf(FatalLevel, msg, fields)
}

// Error will add error record which contains msg and fields
func (p *PingLog) Error(msg string, fields map[string]interface{}) {
	p.logf(ErrorLevel, msg, fields)
}

// Warn will add warn record which contains msg and fields
func (p *PingLog) Warn(msg string, fields map[string]interface{}) {
	p.logf(WarnLevel, msg, fields)
}

// Debug will add debug record which contains msg and fields
func (p *PingLog) Debug(msg string, fields map[string]interface{}) {
	p.logf(DebugLevel, msg, fields)
}

// Trace will add trace info which contains msg and fields
func (p *PingLog) Trace(msg string, fields map[string]interface{}) {
	p.logf(TraceLevel, msg, fields)
}

// SetLevel set log level, and higher level will be recorded
func (p *PingLog) SetLevel(level LogLevel) {
	p.Level = level
}

// SetCxtFielder will get fields from context
func (p *PingLog) SetCtxFielder(handler CtxFielder) {
	p.CtxFielder = handler
}

// SetFormatter will set formatter handler will covert data to string for recording
func (p *PingLog) SetFormatter(formatter Formatter) {
	p.Formatter = formatter
}
