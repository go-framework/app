package logger

func NewLogger(cfg Config, options ...Option) (*Logger, error) {
	return cfg.NewLogger(options...)
}
