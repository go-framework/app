package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func Test_NewDevelopmentConfigLogger(t *testing.T) {
	var log, err = NewDevelopmentConfig().NewLogger()
	assert.Nil(t, err, "NewDevelopmentConfigLogger failed")
	log.Debug("debug")
	log.Info("info")
}

func Test_NewProductionConfigLogger(t *testing.T) {
	var log, err = NewProductionConfig().NewLogger()
	assert.Nil(t, err, "NewProductionConfigLogger failed")
	log.Debug("debug")
	log.Info("info")
}

const consoleConfig = `
# console config template
level: debug
development: false
disableCaller: false
disableStacktrace: false
sampling:
  initial: 100
  thereafter: 100
encoding: console
encoderConfig:
  messageKey: M
  levelKey: L
  timeKey: T
  nameKey: N
  callerKey: C
  functionKey:
  stacktraceKey: S
  lineEnding:
  levelEncoder: capital
  timeEncoder: ISO8601
  durationEncoder: string
  callerEncoder:
  nameEncoder:
  consoleSeparator:
outputPaths:
errorOutputPaths:
initialFields:
writes:
  - name: console
    level: info
    encoding: console
`

func Test_NewConsoleLogger(t *testing.T) {
	var cfg = NewDevelopmentConfig()

	err := yaml.Unmarshal([]byte(consoleConfig), &cfg)
	assert.Nil(t, err, "UnmarshalYAML failed")

	log, err := cfg.NewLogger()
	assert.Nil(t, err, "NewConsoleLogger failed")

	log.Debug("debug")
	log.Info("info")
}
