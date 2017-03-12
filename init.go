package nvidiasmi

import (
	"github.com/rai-project/config"
	logger "github.com/rai-project/logger"
)

var (
	log = logger.New().WithField("pkg", "nvidia-smi")
)

func init() {
	config.AfterInit(func() {
		log = logger.New().WithField("pkg", "nvidia-smi")
	})
}
