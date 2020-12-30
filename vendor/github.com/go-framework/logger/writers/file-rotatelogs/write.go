package file_rotatelogs

import (
	"errors"
	"sync"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
)

type FileRotateLogs struct {
	rotateLogs *rotatelogs.RotateLogs
	once       sync.Once

	// Pattern used to generate actual log file names.
	// You should use patterns using the strftime (3) format.
	Pattern string `json:"pattern" yaml:"pattern"`
	// Filename is the file to write logs to.  Backup log files will be retained
	// in the same directory.
	Filename string `json:"filename" yaml:"filename"`

	// RotationTime is interval between file rotation.
	//By default logs are rotated every 86400 seconds.
	RotationTime int `json:"rotationtime" yaml:"rotationtime"`

	// MaxAge is the maximum number of days to retain old log files based on the
	// timestamp encoded in their filename.  Note that a day is defined as 24
	// hours and may not exactly correspond to calendar days due to daylight
	// savings, leap seconds, etc. The default is not to remove old log files
	// based on age.
	MaxAge int `json:"maxage" yaml:"maxage"`

	// MaxBackups is the maximum number of old log files to retain.  The default
	// is to retain all old log files (though MaxAge may still cause them to get
	// deleted. -1 is disabled.
	MaxBackups int `json:"maxbackups" yaml:"maxbackups"`

	// LocalTime determines if the time used for formatting the timestamps in
	// backup files is the computer's local time.  The default is to use UTC
	// time.
	LocalTime bool `json:"localtime" yaml:"localtime"`
}

func (l FileRotateLogs) NewOptions() (options []rotatelogs.Option) {
	if l.Filename != "" {
		options = append(options, rotatelogs.WithLinkName(l.Filename))
	}
	if l.RotationTime <= 0 {
		l.RotationTime = 86400
	}
	options = append(options, rotatelogs.WithRotationTime(time.Duration(l.RotationTime)*time.Second))
	if l.MaxAge <= 0 {
		l.MaxAge = 7
	}
	options = append(options, rotatelogs.WithMaxAge(time.Duration(l.MaxAge)*time.Hour*24))
	if l.MaxBackups != 0 {
		options = append(options, rotatelogs.WithRotationCount(l.MaxBackups))
	}
	if !l.LocalTime {
		options = append(options, rotatelogs.WithLocation(time.UTC))
	}

	return
}

func (l *FileRotateLogs) NewRotateLogs() (err error) {
	if l.Pattern == "" {
		return errors.New("pattern must be Required")
	}

	l.rotateLogs, err = rotatelogs.New(l.Pattern, l.NewOptions()...)
	return
}

func (l *FileRotateLogs) Write(p []byte) (n int, err error) {
	l.once.Do(func() {
		if err := l.NewRotateLogs(); err != nil {
			panic(err)
		}
	})
	return l.rotateLogs.Write(p)
}
