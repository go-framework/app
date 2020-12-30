package logger

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/imdario/mergo"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Write struct {
	// Writer name.
	Name string `json:"name" yaml:"name"`
	// Writer config.
	Writer io.Writer `json:"writer" yaml:"-"`
	// Level is the minimum enabled logging level. Note that this is a dynamic
	// level, so calling Config.Level.SetLevel will atomically change the log
	// level of all loggers descended from this config.
	Level zap.AtomicLevel `json:"level" yaml:"level"`
	// Encoding sets the logger's encoding. Valid values are "json" and
	// "console", as well as any third-party encodings registered via
	// RegisterEncoder.
	Encoding string `json:"encoding" yaml:"encoding"`
	// Logger encoder config.
	EncoderConfig *zapcore.EncoderConfig `json:"encoderConfig" yaml:"encoderConfig"`
}

type write Write

func (w *Write) GetLevel(level zap.AtomicLevel) zap.AtomicLevel {
	if w.Level == (zap.AtomicLevel{}) {
		w.Level = level
	}

	return w.Level
}

func (w *Write) GetEncoding(encoding string) string {
	if w.Encoding == "" {
		if encoding == "" {
			w.Encoding = "console"
		} else {
			w.Encoding = encoding
		}
	}

	return w.Encoding
}

func (w *Write) GetEncoderConfig(config *zapcore.EncoderConfig) (*zapcore.EncoderConfig, error) {
	if w.EncoderConfig == nil {
		if config == nil {
			*w.EncoderConfig = zap.NewDevelopmentEncoderConfig()
		} else {
			w.EncoderConfig = config
		}
	} else {
		if config != nil {
			if err := mergo.Merge(w.EncoderConfig, config); err != nil {
				return nil, err
			}
		}
	}

	return w.EncoderConfig, nil
}

func (w Write) buildEncoder() (zapcore.Encoder, error) {
	return newEncoder(w.Encoding, *w.EncoderConfig)
}

// Implement YAML Unmarshaler interface.
func (w *Write) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var _w = write{}
	if err := unmarshal(&_w); err != nil {
		return err
	}

	var ok bool = false
	// load Writer.
	_w.Writer, ok = GetWriter(_w.Name)
	if !ok {
		return errors.New("unsupported write name: " + w.Name)
	}

	type T struct {
		// Writer config.
		Writer map[string]interface{} `json:"writer" yaml:"writer"`
	}

	var _t = T{}
	if err := unmarshal(&_t); err != nil {
		return err
	}

	if err := mapstructure.Decode(_t.Writer, _w.Writer); err != nil {
		return err
	}

	*w = Write(_w)

	return nil
}

// Implement JSON Unmarshaler interface.
func (w *Write) UnmarshalJSON(data []byte) error {
	type T struct {
		// Writer name.
		Name string `json:"name" yaml:"name"`
	}

	var _t = T{}
	if err := json.Unmarshal(data, &_t); err != nil {
		return err
	}

	// load Writer.
	writer, ok := GetWriter(_t.Name)
	if !ok {
		return errors.New("unsupported write name: " + w.Name)
	}

	var _w = write{
		Writer: writer,
	}

	if err := json.Unmarshal(data, &_w); err != nil {
		return err
	}

	*w = Write(_w)

	return nil
}
