package loggerfactory

import (
	"corebanking/pkg/helper/config"

	"corebanking/pkg/helper/logger_factory/zap"

	"github.com/pkg/errors"
)

// ZapFactory receiver for zap factory.
type ZapFactory struct{}

// Build zap logger.
func (mf *ZapFactory) Build(lc *config.LogConfig) error {
	err := zap.RegisterLog(*lc)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}
