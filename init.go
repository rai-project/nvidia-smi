package nvidiasmi

import (
	"github.com/Sirupsen/logrus"

	"github.com/rai-project/config"
	logger "github.com/rai-project/logger"
)

var (
	log *logrus.Entry
)

func init() {
	config.OnInit(func() {
		log = logger.New().WithField("pkg", "nvidia-smi")
	})
}
