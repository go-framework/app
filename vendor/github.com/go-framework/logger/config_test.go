package logger

import (
	"encoding/json"
	"testing"

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
	if err != nil {
		t.Fatalf("config UnmarshalYAML error: %v", err)
	}

	assert.Equal(t, cfg.Level.String(), "debug")
	assert.Equal(t, cfg.Development, false)
	assert.Equal(t, cfg.Encoding, "console")
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
	if err != nil {
		t.Fatalf("config UnmarshalJSON error: %v", err)
	}

	assert.Equal(t, cfg.Level.String(), "debug")
	assert.Equal(t, cfg.Development, false)
	assert.Equal(t, cfg.Encoding, "console")
}
