package nvidiasmi

type NvidiaGPU struct {
	Name              string       `json:"name,omitempty"`
	NumSMs            int          `json:"num_sms,omitempty"`
	ComputeCapability float64      `json:"compute_capability,omitempty"`
	Architecture      string       `json:"architecture,omitempty"`
	Interconnect      Interconnect `json:"interconnect,omitempty"`
	ClockRate         int64        `json:"clock_rate,omitempty"`
	PeekGFlops        int64        `json:"peek_gflops,omitempty"`
	MemoryBandwidth   float64      `json:"memory_bandwidth,omitempty"`
}

var (
	NVIDIAT4 = NvidiaGPU{
		Name:              "Tesla T4",
		NumSMs:            40,
		ComputeCapability: 7.5,
		Architecture:      "Volta",
		Interconnect:      PCIe3{},
		ClockRate:         int64(1530),
		PeekGFlops:        int64(8100),
		MemoryBandwidth:   float64(300),
	}

	NVIDIATeslaV100SXM2 = NvidiaGPU{
		Name:              "Tesla V100-SXM2-16GB",
		NumSMs:            80,
		ComputeCapability: 7.0,
		Architecture:      "Volta",
		Interconnect:      NVLink2{},
		ClockRate:         int64(1530),
		PeekGFlops:        int64(15700),
		MemoryBandwidth:   float64(900),
	}

	NVIDIATeslaV100PCIe = NvidiaGPU{
		Name:            "TESLA V100 PCIE",
		Architecture:    "Volta",
		Interconnect:    PCIe2{},
		ClockRate:       int64(1380),
		PeekGFlops:      int64(14000),
		MemoryBandwidth: float64(900),
	}

	// https://devblogs.nvidia.com/new-pascal-gpus-accelerate-inference-in-the-data-center/
	NVIDIATeslaP4 = NvidiaGPU{
		Name:              "Tesla P4",
		NumSMs:            20,
		ComputeCapability: 6.1,
		Architecture:      "Pascal",
		Interconnect:      PCIe3{},
		ClockRate:         int64(1063),
		PeekGFlops:        int64(5442),
		MemoryBandwidth:   float64(192),
	}

	NVIDIATeslaP40 = NvidiaGPU{
		Name:              "Tesla P40",
		NumSMs:            30,
		ComputeCapability: 6.1,
		Architecture:      "Pascal",
		Interconnect:      PCIe3{},
		ClockRate:         int64(1531),
		PeekGFlops:        int64(11758),
		MemoryBandwidth:   float64(346),
	}

	NVIDIATeslaP100SXM2 = NvidiaGPU{
		Name:            "TESLA P100 SXM2",
		Architecture:    "Pascal",
		NumSMs:          56,
		Interconnect:    SXM2{},
		ClockRate:       int64(1481),
		PeekGFlops:      int64(10600),
		MemoryBandwidth: float64(732),
	}

	NVIDIATeslaP100PCIe = NvidiaGPU{
		Name:            "TESLA P100 PCIE",
		Architecture:    "Pascal",
		NumSMs:          56,
		Interconnect:    PCIe2{},
		ClockRate:       int64(1328),
		PeekGFlops:      int64(9300),
		MemoryBandwidth: float64(732),
	}

	NVIDIATitanXP = NvidiaGPU{
		Name:            "TITAN Xp",
		Architecture:    "Maxwell",
		Interconnect:    PCIe2{},
		ClockRate:       int64(1582),
		PeekGFlops:      int64(12100),
		MemoryBandwidth: float64(547.7),
	}

	NVIDIATitanX = NvidiaGPU{
		Name:            "TITAN X",
		Architecture:    "Maxwell",
		Interconnect:    PCIe2{},
		ClockRate:       int64(1000),
		PeekGFlops:      int64(6144),
		MemoryBandwidth: float64(336.5),
	}
	NVIDIAK20 = NvidiaGPU{
		Name:            "K20",
		Architecture:    "Kepler",
		Interconnect:    PCIe2{},
		ClockRate:       int64(1000),
		PeekGFlops:      int64(3520),
		MemoryBandwidth: float64(208),
	}
	NVIDIAK20X = NvidiaGPU{
		Name:            "K20X",
		Architecture:    "Kepler",
		Interconnect:    PCIe2{},
		ClockRate:       int64(1000),
		PeekGFlops:      int64(3935),
		MemoryBandwidth: float64(250),
	}

	NVIDIAK40 = NvidiaGPU{
		Name:            "K40",
		Architecture:    "Kepler",
		Interconnect:    PCIe2{},
		ClockRate:       int64(745),
		PeekGFlops:      int64(4290),
		MemoryBandwidth: float64(288),
	}

	NVIDIAK80 = NvidiaGPU{
		Name:            "K80",
		Architecture:    "Kepler",
		Interconnect:    PCIe2{},
		NumSMs:          13,
		ClockRate:       int64(560),
		PeekGFlops:      int64(5600),
		MemoryBandwidth: float64(480),
	}

	NvidiaGPUs = []NvidiaGPU{
		NVIDIATeslaV100SXM2,
		NVIDIATeslaV100PCIe,
		NVIDIATeslaP4,
		NVIDIATeslaP40,
		NVIDIATeslaP100SXM2,
		NVIDIATeslaP100PCIe,
		NVIDIATitanXP,
		NVIDIATitanX,
		NVIDIAK20,
		NVIDIAK20X,
		NVIDIAK40,
		NVIDIAK80,
	}
)
