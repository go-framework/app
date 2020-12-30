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
level: info
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
    level: debug
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
	log.Warn("warn")
}

const lumberjackConfig = `
# lumberjack config template
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
  - name: lumberjack
    level: info
    encoding: json
    encoderConfig:
      messageKey: msg
      levelKey: level
      nameKey: name
      callerKey: caller
      functionKey:
      stacktraceKey: stack
      lineEnding:
      levelEncoder: capital
      timeEncoder: ISO8601
      durationEncoder: string
      callerEncoder:
      nameEncoder:
      consoleSeparator:
    writer:
      filename: lumberjack.log
      maxsize: 1024
      maxage: 30
      maxbackups: 3
      localtime: true
      compress: true
`

func Test_NewLumberjackLogger(t *testing.T) {
	var cfg = NewDevelopmentConfig()

	err := yaml.Unmarshal([]byte(lumberjackConfig), &cfg)
	assert.Nil(t, err, "UnmarshalYAML failed")

	log, err := cfg.NewLogger()
	assert.Nil(t, err, "NewLumberjackLogger failed")

	log.Debug("debug")
	log.Info("info")
	log.Warn("warn")
}

const consoleAndLumberjackConfig = `
# config template
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
  - name: lumberjack
    level: info
    encoding: json
    encoderConfig:
      messageKey: msg
      timeKey: time
      nameKey: name
      callerKey: caller
      functionKey:
      stacktraceKey: stack
      lineEnding:
      levelEncoder: capital
      timeEncoder: ISO8601
      durationEncoder: string
      callerEncoder:
      nameEncoder:
      consoleSeparator:
    writer:
      filename: lumberjack.log
      maxsize: 1024
      maxage: 30
      maxbackups: 3
      localtime: true
      compress: true
  - name: console
    level: debug
    encoding: console
`

func Test_NewConsoleAndLumberjackLogger(t *testing.T) {
	var cfg = NewDevelopmentConfig()

	err := yaml.Unmarshal([]byte(consoleAndLumberjackConfig), &cfg)
	assert.Nil(t, err, "UnmarshalYAML failed")

	log, err := cfg.NewLogger()
	assert.Nil(t, err, "NewConsoleAndLumberjackLogger failed")

	log.Debug("debug")
	log.Info("info")
	log.Warn("warn")
}

const fileRotateLogsConfig = `
# file-rotatelogs config template
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
  - name: file-rotatelogs
    level: info
    encoding: json
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
    writer:
      filename: file-rotatelogs.log
      pattern: file-rotatelogs-%Y%m%d.log
      rotationtime: 86400
      maxage: 30
      maxbackups: 3
      localtime: true
`

func Test_NewFileRotateLogsConfigLogger(t *testing.T) {
	var cfg = NewDevelopmentConfig()

	err := yaml.Unmarshal([]byte(fileRotateLogsConfig), &cfg)
	assert.Nil(t, err, "UnmarshalYAML failed")

	log, err := cfg.NewLogger()
	assert.Nil(t, err, "NewFileRotateLogsConfigLogger failed")

	log.Debug("debug")
	log.Info("info")
	log.Warn("warn")
}
