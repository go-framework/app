package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	// Logger write list.
	Writes []*Write `json:"writes" yaml:"writes"`
	// Logger config.
	zap.Config `json:",inline" yaml:",inline"`
}

func NewDevelopmentConfig() Config {
	var cfg = Config{
		Config: zap.NewDevelopmentConfig(),
	}
	return cfg
}

func NewProductionConfig() Config {
	var cfg = Config{
		Config: zap.NewProductionConfig(),
	}
	return cfg
}

// Build constructs a logger from the Config and Options.
func (cfg Config) Build(opts ...Option) (*Logger, error) {
	if cfg.Config.Level == (zap.AtomicLevel{}) {
		cfg.Config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	sink, errSink, err := cfg.openSinks()
	if err != nil {
		return nil, err
	}

	var cores []zapcore.Core

	for i := range cfg.Writes {
		// merge
		cfg.Writes[i].GetLevel(cfg.Config.Level)
		cfg.Writes[i].GetEncoding(cfg.Config.Encoding)
		_, err := cfg.Writes[i].GetEncoderConfig(&cfg.Config.EncoderConfig)
		if err != nil {
			return nil, err
		}

		enc, err := cfg.Writes[i].buildEncoder()
		if err != nil {
			return nil, err
		}

		var writes []zapcore.WriteSyncer
		if cfg.Writes[i].Writer != nil {
			writes = append(writes, zapcore.AddSync(cfg.Writes[i].Writer))
		}
		writes = append(writes, sink)

		cores = append(cores, zapcore.NewCore(
			enc,
			zap.CombineWriteSyncers(writes...),
			cfg.Writes[i].Level,
		))
	}

	log := zap.New(
		zapcore.NewTee(cores...),
		cfg.buildOptions(errSink)...,
	)

	if len(opts) > 0 {
		log = log.WithOptions(opts...)
	}

	return log, nil
}
