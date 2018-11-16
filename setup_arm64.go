// +build !nogpu
// +build linux,arm64

package nvidiasmi

import (
	"github.com/rai-project/tegra"
)

func Init() {
	defer close(initialized)
	if !tegra.IsSupported {
		HasGPU = false
		return
	}
	HasGPU = true

	soc, err := tegra.SOC()
	if err != nil {
		soc = tegra.UnknownSOC
	}

	kind, err := tegra.Kind()
	if err != nil {
		kind = tegra.UnknownKind
	}

	Info = &NvidiaSmi{
		DriverVersion: string(soc),
		AttachedGpus:  "tegra-gpus",
		GPUS: []GPU{
			GPU{
				ID: string(kind),
			},
		},
	}
	GPUCount = 1
}
