// +build linux,arm64

package smi

import (
	"github.com/rai-project/tegra"
)

func Init() {
	defer close(initialized)
	if !tegra.IsSupported() {
		HasGPU = false
		return
	}
	HasGPU = true
	info := &NvidiaSmi{
    DriverVersion: "tegra-version",
    AttachedGpus: "tegra-gpus",
    GPUS: []GPU{
      GPU{
        ID: "tegra-gpu",
      },
    }
  }
	GPUCount = 1
}
