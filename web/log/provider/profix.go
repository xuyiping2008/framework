package provider

func Prefix(level LogLevel) string {
	prefix := ""
	switch level {
	case PanicLevel:
		prefix = "[Panic]"
	case FatalLevel:
		prefix = "[Fatal]"
	case ErrorLevel:
		prefix = "[Error]"
	case WarnLevel:
		prefix = "[Warn]"
	case InfoLevel:
		prefix = "[Info]"
	case DebugLevel:
		prefix = "[Debug]"
	case TraceLevel:
		prefix = "[Trace]"
	}
	return prefix
}
