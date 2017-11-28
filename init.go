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

func Init() {
	defer close(initialized)
	HasGPU = false
	info, err := New()
	if err != nil {
		log.WithError(err).Info("was not able to get nvidia-smi info")
		return
	}
	Info = info
	GPUCount = len(Info.GPUS)
	HasGPU = GPUCount > 0
}

func init() {
	config.BeforeInit(func() {
		Init()
	})
	config.AfterInit(func() {
		log = logger.New().WithField("pkg", "nvidia-smi")
	})
}
