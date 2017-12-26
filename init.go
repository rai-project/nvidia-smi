package nvidiasmi

import (
	"github.com/rai-project/config"
	logger "github.com/rai-project/logger"
)

var (
	log         = logger.New().WithField("pkg", "nvidia-smi")
	initialized = make(chan struct{}, 1)
)

func Wait() {
	<-initialized
}

func init() {
	config.BeforeInit(func() {
		Init()
	})
	config.AfterInit(func() {
		log = logger.New().WithField("pkg", "nvidia-smi")
	})
}
