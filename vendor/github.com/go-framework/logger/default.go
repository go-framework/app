package logger

var (
	DefaultLogger        *Logger
	DefaultSugaredLogger *SugaredLogger
)

func init() {
	DefaultLogger, _ = NewDevelopmentConfig().NewLogger()
	DefaultSugaredLogger = newSugaredLogger(DefaultLogger.Sugar())
}
