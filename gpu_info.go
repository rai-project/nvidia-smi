package nvidiasmi

import (
	"github.com/NVIDIA/gpu-monitoring-tools/bindings/go/nvml"
	"github.com/pkg/errors"
)

var devices = map[int]*nvml.Device{}

func getNVMLDevice(id int) (*nvml.Device, error) {
	device, ok := devices[id]
	if ok {
		return device, nil
	}
	device, err := nvml.NewDevice(uint(id))
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get device with id = %d", id)
	}
	devices[id] = device
	return device, nil
}

func (g GPU) NumSMs() int {
	return 0
}

func (g GPU) ComputeCapability() float64 {
	return 0
}

func (g GPU) TheoreticalFlops() float64 {
	return 0
}

func pciBandwidth(gen, width *uint) *uint {
	m := map[uint]uint{
		1: 250, // MB/s
		2: 500,
		3: 985,
		4: 1969,
	}
	if gen == nil || width == nil {
		return nil
	}
	bw := m[*gen] * *width
	return &bw
}

func (g GPU) Bandwidth() (uint, error) {
	device, err := getNVMLDevice(g.ID)
}
