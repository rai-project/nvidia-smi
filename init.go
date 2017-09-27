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
	config.AfterInit(func() {
		defer close(initialized)

		log = logger.New().WithField("pkg", "nvidia-smi")

		HasGPU = false
		info, err := New()
		if err != nil {
			log.WithError(err).Info("was not able to get nvidia-smi info")
			return
		}
		Info = info
		GPUCount = len(Info.GPUS)
		HasGPU = GPUCount > 0
	})
}
