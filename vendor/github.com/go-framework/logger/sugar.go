package logger

import "go.uber.org/zap"

// Default sugared logger.
type SugaredLogger struct {
	*zap.SugaredLogger
}

func newSugaredLogger(sugar *zap.SugaredLogger) *SugaredLogger {
	return &SugaredLogger{
		SugaredLogger: sugar,
	}
}

// New sugared logger by default config.
func NewSugaredLogger(opts ...Option) *SugaredLogger {
	logger, _ := NewDevelopmentConfig().NewLogger()
	return &SugaredLogger{
		SugaredLogger: logger.Sugar(),
	}
}

// wrap args split by space
func wrap(args []interface{}) (a []interface{}) {
	for n, arg := range args {
		if len(args) > 0 && n < len(args)-1 {
			a = append(a, arg, " ")
		} else {
			a = append(a, arg)
		}
	}
	return
}

// Print uses fmt.Sprint to construct and log a message.
func (l *SugaredLogger) Print(args ...interface{}) {
	l.SugaredLogger.Info(args...)
}

// Printf uses fmt.Sprintf to log a templated message.
func (l *SugaredLogger) Printf(template string, args ...interface{}) {
	l.SugaredLogger.Infof(template, args...)
}

// Debugln uses fmt.Sprint to construct and log a message.
func (l *SugaredLogger) Debugln(args ...interface{}) {
	l.SugaredLogger.Debug(wrap(args)...)
}

// Infoln uses fmt.Sprint to construct and log a message.
func (l *SugaredLogger) Infoln(args ...interface{}) {
	l.SugaredLogger.Info(wrap(args)...)
}

// Println uses fmt.Sprint to construct and log a message.
func (l *SugaredLogger) Println(args ...interface{}) {
	l.SugaredLogger.Info(wrap(args)...)
}

// Warnln uses fmt.Sprint to construct and log a message.
func (l *SugaredLogger) Warnln(args ...interface{}) {
	l.SugaredLogger.Warn(wrap(args)...)
}

// Errorln uses fmt.Sprint to construct and log a message.
func (l *SugaredLogger) Errorln(args ...interface{}) {
	l.SugaredLogger.Error(wrap(args)...)
}

// DPanicln uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (See DPanicLevel for details.)
func (l *SugaredLogger) DPanicln(args ...interface{}) {
	l.SugaredLogger.DPanic(wrap(args)...)
}

// Panicln uses fmt.Sprint to construct and log a message, then panics.
func (l *SugaredLogger) Panicln(args ...interface{}) {
	l.SugaredLogger.Panic(wrap(args)...)
}

// Fatalln uses fmt.Sprint to construct and log a message, then calls os.Exit.
func (l *SugaredLogger) Fatalln(args ...interface{}) {
	l.SugaredLogger.Fatal(wrap(args)...)
}

// With adds a variadic number of fields to the logging context. It accepts a
// mix of strongly-typed Field objects and loosely-typed key-value pairs. When
// processing pairs, the first element of the pair is used as the field key
// and the second as the field value.
//
// For example,
//   sugaredLogger.With(
//     "hello", "world",
//     "failure", errors.New("oh no"),
//     Stack(),
//     "count", 42,
//     "user", User{Name: "alice"},
//  )
// is the equivalent of
//   unsugared.With(
//     String("hello", "world"),
//     String("failure", "oh no"),
//     Stack(),
//     Int("count", 42),
//     Object("user", User{Name: "alice"}),
//   )
//
// Note that the keys in key-value pairs should be strings. In development,
// passing a non-string key panics. In production, the logger is more
// forgiving: a separate error is logged, but the key-value pair is skipped
// and execution continues. Passing an orphaned key triggers similar behavior:
// panics in development and errors in production.
func (l *SugaredLogger) With(args ...interface{}) *SugaredLogger {
	return &SugaredLogger{SugaredLogger: l.SugaredLogger.With(args...)}
}

// Debug uses fmt.Sprint to construct and log a message.
func Debug(args ...interface{}) {
	DefaultSugaredLogger.Debug(args...)
}

// Info uses fmt.Sprint to construct and log a message.
func Info(args ...interface{}) {
	DefaultSugaredLogger.Info(args...)
}

// Print uses fmt.Sprint to construct and log a message.
func Print(args ...interface{}) {
	DefaultSugaredLogger.Info(args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func Warn(args ...interface{}) {
	DefaultSugaredLogger.Warn(args...)
}

// Error uses fmt.Sprint to construct and log a message.
func Error(args ...interface{}) {
	DefaultSugaredLogger.Error(args...)
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanic(args ...interface{}) {
	DefaultSugaredLogger.DPanic(args...)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func Panic(args ...interface{}) {
	DefaultSugaredLogger.Panic(args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func Fatal(args ...interface{}) {
	DefaultSugaredLogger.Fatal(args...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func Debugf(template string, args ...interface{}) {
	DefaultSugaredLogger.Debugf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(template string, args ...interface{}) {
	DefaultSugaredLogger.Infof(template, args...)
}

// Printf uses fmt.Sprintf to log a templated message.
func Printf(template string, args ...interface{}) {
	DefaultSugaredLogger.Infof(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(template string, args ...interface{}) {
	DefaultSugaredLogger.Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...interface{}) {
	DefaultSugaredLogger.Errorf(template, args...)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanicf(template string, args ...interface{}) {
	DefaultSugaredLogger.DPanicf(template, args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func Panicf(template string, args ...interface{}) {
	DefaultSugaredLogger.Panicf(template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func Fatalf(template string, args ...interface{}) {
	DefaultSugaredLogger.Fatalf(template, args...)
}

// Debugln uses fmt.Sprint to construct and log a message.
func Debugln(args ...interface{}) {
	DefaultSugaredLogger.Debug(wrap(args)...)
}

// Infoln uses fmt.Sprint to construct and log a message.
func Infoln(args ...interface{}) {
	DefaultSugaredLogger.Info(wrap(args)...)
}

// Println uses fmt.Sprint to construct and log a message.
func Println(args ...interface{}) {
	DefaultSugaredLogger.Info(wrap(args)...)
}

// Warnln uses fmt.Sprint to construct and log a message.
func Warnln(args ...interface{}) {
	DefaultSugaredLogger.Warn(wrap(args)...)
}

// Errorln uses fmt.Sprint to construct and log a message.
func Errorln(args ...interface{}) {
	DefaultSugaredLogger.Error(wrap(args)...)
}

// DPanicln uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanicln(args ...interface{}) {
	DefaultSugaredLogger.DPanic(wrap(args)...)
}

// Panicln uses fmt.Sprint to construct and log a message, then panics.
func Panicln(args ...interface{}) {
	DefaultSugaredLogger.Panic(wrap(args)...)
}

// Fatalln uses fmt.Sprint to construct and log a message, then calls os.Exit.
func Fatalln(args ...interface{}) {
	DefaultSugaredLogger.Fatal(wrap(args)...)
}

// With adds a variadic number of fields to the logging context. It accepts a
// mix of strongly-typed Field objects and loosely-typed key-value pairs. When
// processing pairs, the first element of the pair is used as the field key
// and the second as the field value.
//
// For example,
//   sugaredLogger.With(
//     "hello", "world",
//     "failure", errors.New("oh no"),
//     Stack(),
//     "count", 42,
//     "user", User{Name: "alice"},
//  )
// is the equivalent of
//   unsugared.With(
//     String("hello", "world"),
//     String("failure", "oh no"),
//     Stack(),
//     Int("count", 42),
//     Object("user", User{Name: "alice"}),
//   )
//
// Note that the keys in key-value pairs should be strings. In development,
// passing a non-string key panics. In production, the logger is more
// forgiving: a separate error is logged, but the key-value pair is skipped
// and execution continues. Passing an orphaned key triggers similar behavior:
// panics in development and errors in production.
func With(args ...interface{}) *SugaredLogger {
	return DefaultSugaredLogger.With(args...)
}
