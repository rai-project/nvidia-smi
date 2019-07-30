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
	NVIDIARTX6000 = NvidiaGPU{
		Name:              "Quadro RTX 6000",
		NumSMs:            40,
		ComputeCapability: 7.5,
		Architecture:      "Turing",
		Interconnect:      PCIe3{},
		ClockRate:         int64(1770),
		PeekGFlops:        int64(16300),
		MemoryBandwidth:   float64(624),
	}

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

	NVIDIATeslaV100PCIe16GB = NvidiaGPU{
		Name:            "Tesla V100-PCIE-16GB",
		Architecture:    "Volta",
		Interconnect:    PCIe3{},
		ClockRate:       int64(1380),
		PeekGFlops:      int64(14000),
		MemoryBandwidth: float64(900),
	}

	NVIDIATeslaV100PCIe32GB = NvidiaGPU{
		Name:            "Tesla V100-PCIE-32GB",
		Architecture:    "Volta",
		Interconnect:    PCIe3{},
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

	NVIDIATeslaP100PCIe16GB = NvidiaGPU{
		Name:              "Tesla P100-PCIE-16GB",
		Architecture:      "Pascal",
		ComputeCapability: 6.0,
		NumSMs:            80,
		Interconnect:      PCIe3{},
		ClockRate:         int64(1380),
		PeekGFlops:        int64(9300),
		MemoryBandwidth:   float64(732),
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

	// https://www.microway.com/knowledge-center-articles/in-depth-comparison-of-nvidia-tesla-maxwell-gpu-accelerators/
	NVIDIAM60 = NvidiaGPU{
		Name:            "Tesla M60",
		Architecture:    "Maxwell",
		Interconnect:    PCIe3{},
		NumSMs:          16,
		ClockRate:       int64(1178),
		PeekGFlops:      int64(9640),
		MemoryBandwidth: float64(320),
	}

	// https://www.microway.com/knowledge-center-articles/in-depth-comparison-of-nvidia-tesla-kepler-gpu-accelerators/
	NVIDIAK80 = NvidiaGPU{
		Name:            "Tesla K80",
		Architecture:    "Kepler",
		Interconnect:    PCIe3{},
		NumSMs:          26,
		ClockRate:       int64(875),
		PeekGFlops:      int64(8730),
		MemoryBandwidth: float64(480),
	}

	NVIDIAK40 = NvidiaGPU{
		Name:            "Tesla K40",
		Architecture:    "Kepler",
		Interconnect:    PCIe3{},
		NumSMs:          15,
		ClockRate:       int64(875),
		PeekGFlops:      int64(5040),
		MemoryBandwidth: float64(288),
	}

	NVIDIAK20X = NvidiaGPU{
		Name:            "Tesla K20X",
		Architecture:    "Kepler",
		Interconnect:    PCIe2{},
		ClockRate:       int64(784),
		PeekGFlops:      int64(3935),
		MemoryBandwidth: float64(250),
	}

	NVIDIAK20 = NvidiaGPU{
		Name:            "Tesla K20",
		Architecture:    "Kepler",
		Interconnect:    PCIe2{},
		ClockRate:       int64(705),
		PeekGFlops:      int64(3520),
		MemoryBandwidth: float64(208),
	}

	NvidiaGPUs = []NvidiaGPU{
		NVIDIATeslaV100SXM2,
		NVIDIATeslaV100PCIe16GB,
		NVIDIATeslaV100PCIe32GB,
		NVIDIATeslaP4,
		NVIDIATeslaP40,
		NVIDIATeslaP100PCIe16GB,
		NVIDIATitanXP,
		NVIDIATitanX,
		NVIDIAM60,
		NVIDIAK20,
		NVIDIAK20X,
		NVIDIAK40,
		NVIDIAK80,
	}
)
