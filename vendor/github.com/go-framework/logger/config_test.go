package logger

import (
	"encoding/json"
	"testing"

	"github.com/natefinch/lumberjack"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

const rawYaml = `
# lumberjack config template
level: debug
development: false
encoding: console
encoderConfig:
  messageKey: message
writes:
  - name: lumberjack
    writer:
      filename: lumberjack.log
      maxsize: 1024
      maxage: 30
      maxbackups: 3
      localtime: true
      compress: true
    level: info
    encoding: json
    encoderConfig:
      messageKey: msg
`

func TestConfig_UnmarshalYAML(t *testing.T) {
	var cfg = NewDevelopmentConfig()

	err := yaml.Unmarshal([]byte(rawYaml), &cfg)

	assert.Nil(t, err, "UnmarshalYAML failed")
	assert.Equal(t, cfg.Level.String(), "debug")
	assert.Equal(t, cfg.Development, false)
	assert.Equal(t, cfg.Encoding, "console")
	assert.Equal(t, cfg.EncoderConfig.MessageKey, "message")
	for _, write := range cfg.Writes {
		if write.Name == "lumberjack" {
			assert.Equal(t, write.Level.String(), "info")
			assert.Equal(t, write.Encoding, "json")
			assert.Equal(t, write.EncoderConfig.MessageKey, "msg")
			obj, ok := write.Writer.(*lumberjack.Logger)
			assert.True(t, ok, "should be *lumberjack.Logger type")
			assert.Equal(t, obj.Filename, "lumberjack.log")
			assert.Equal(t, obj.MaxSize, 1024)
			assert.Equal(t, obj.MaxAge, 30)
			assert.Equal(t, obj.MaxBackups, 3)
			assert.Equal(t, obj.LocalTime, true)
			assert.Equal(t, obj.Compress, true)
		}
	}
}

const rawJson = `
{
    "level":"debug",
    "development":false,
    "encoding":"console",
    "encoderConfig":{
        "messageKey":"message"
    },
    "writes":[
        {
            "name":"lumberjack",
            "writer":{
                "filename":"lumberjack.log",
                "maxsize":1024,
                "maxage":30,
                "maxbackups":3,
                "localtime":true,
                "compress":true
            },
            "level":"info",
            "encoding":"json",
            "encoderConfig":{
                "messageKey":"msg"
            }
        }
    ]
}
`

func TestConfig_UnmarshalJSON(t *testing.T) {
	var cfg = NewDevelopmentConfig()

	err := json.Unmarshal([]byte(rawJson), &cfg)

	assert.Nil(t, err, "UnmarshalJSON failed")
	assert.Equal(t, cfg.Level.String(), "debug")
	assert.Equal(t, cfg.Development, false)
	assert.Equal(t, cfg.Encoding, "console")
	assert.Equal(t, cfg.EncoderConfig.MessageKey, "message")
	for _, write := range cfg.Writes {
		if write.Name == "lumberjack" {
			assert.Equal(t, write.Level.String(), "info")
			assert.Equal(t, write.Encoding, "json")
			assert.Equal(t, write.EncoderConfig.MessageKey, "msg")
			obj, ok := write.Writer.(*lumberjack.Logger)
			assert.True(t, ok, "should be *lumberjack.Logger type")
			assert.Equal(t, obj.Filename, "lumberjack.log")
			assert.Equal(t, obj.MaxSize, 1024)
			assert.Equal(t, obj.MaxAge, 30)
			assert.Equal(t, obj.MaxBackups, 3)
			assert.Equal(t, obj.LocalTime, true)
			assert.Equal(t, obj.Compress, true)
		}
	}
}
