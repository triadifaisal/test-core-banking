package zap

import (
	"corebanking/pkg/helper/config"
	"corebanking/pkg/helper/logger"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// RegisterLog ...
func RegisterLog(lc config.LogConfig) error {
	zLogger, err := initLog(lc)
	if err != nil {
		return errors.Wrap(err, "RegisterLog")
	}
	//nolint: errcheck // TODO: add to backlog
	defer zLogger.Sync()
	zSugarlog := zLogger.Sugar()
	logger.SetLogger(zSugarlog)

	return nil
}

// initLog create logger.
func initLog(lc config.LogConfig) (zap.Logger, error) {
	var (
		config  zap.Config
		zLogger *zap.Logger
	)

	// customize production configuration
	config = zap.NewProductionConfig()
	config.Encoding = lc.Encoding
	config.OutputPaths = append(config.OutputPaths, lc.FilePath)
	config.ErrorOutputPaths = append(config.ErrorOutputPaths, lc.FilePath)
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder

	// customize it from configuration file
	err := customizeLogFromConfig(&config, lc)
	if err != nil {
		return zap.Logger{}, errors.Wrap(err, "cfg.Build()")
	}
	zLogger, err = config.Build(zap.AddCallerSkip(lc.CallerSkip))
	if err != nil {
		return zap.Logger{}, errors.Wrap(err, "cfg.Build()")
	}

	zLogger.Debug("logger construction succeeded")
	return *zLogger, nil
}

// customizeLogFromConfig customize log based on parameters from configuration file.
func customizeLogFromConfig(cfg *zap.Config, lc config.LogConfig) error {
	cfg.DisableCaller = !lc.EnableCaller

	// set log level
	l := zap.NewAtomicLevel().Level()
	err := l.Set(lc.Level)
	if err != nil {
		return errors.Wrap(err, "")
	}
	cfg.Level.SetLevel(l)

	return nil
}
