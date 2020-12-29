package logger

func NewLogger(cfg Config, options ...Option) (*Logger, error) {
	//if cfg.Config == nil {
	//	*cfg.Config = zap.NewDevelopmentConfig()
	//}

	if len(cfg.Writes) == 0 {
		return cfg.Config.Build(options...)
	}

	return cfg.Build(options...)
}
