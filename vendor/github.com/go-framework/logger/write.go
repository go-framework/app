package logger

import (
	"encoding/json"
	"errors"
	"fmt"
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
	Writer io.Writer `json:"-" yaml:"-"`
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

// unmarshal map[string]interface{}) data, get the name and config is exist,
// the Writer which implement structure should be tag as `json:",inline" yaml:",inline" mapstructure:",squash"` format.
func (w *Write) unmarshal(data map[string]interface{}) error {
	// get write name.
	if name, ok := data["name"]; ok {
		switch v := name.(type) {
		case string:
			w.Name = v
		default:
			w.Name = fmt.Sprintf("%v", name)
		}
	} else {
		return errors.New("write must be have name filed")
	}

	// load Writer.
	if writer, ok := GetWriter(w.Name); ok {
		w.Writer = writer
	} else {
		return errors.New("unsupported write name: " + w.Name)
	}

	// if have writer config filed then parse it.
	if writer, ok := data["writer"]; ok {
		err := mapstructure.Decode(writer, w.Writer)
		if err != nil {
			return err
		}
	}

	return nil
}

// Implement YAML Unmarshaler interface.
func (w *Write) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// // new temp as map[string]interface{}.
	// temp := make(map[string]interface{})
	//
	// // unmarshal temp as yaml.
	// if err := unmarshal(temp); err != nil {
	// 	return err
	// }
	//
	// return w.unmarshal(temp)

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
		Writer interface{} `json:"writer" yaml:"writer"`
	}

	// var _t = T{
	// 	Writer: _w.Writer,
	// }
	if err := unmarshal(_w.Writer); err != nil {
		return err
	}

	// err := mapstructure.Decode(writer, w.Writer)

	*w = Write(_w)

	return nil
}

// Implement JSON Unmarshaler interface.
func (w *Write) UnmarshalJSON(data []byte) error {
	// new temp as map[string]interface{}.
	temp := make(map[string]interface{})

	// unmarshal temp as json.
	err := json.Unmarshal(data, &temp)
	if err != nil {
		return err
	}

	return w.unmarshal(temp)
}
