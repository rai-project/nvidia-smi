package nvidiasmi

import (
	"strings"

	"github.com/pkg/errors"
)

var (
	archSMMapping = map[float64]int{
		3.0: 192, // Kepler Generation (SM 3.0) GK10x class
		3.2: 192, // Kepler Generation (SM 3.2) GK10x class
		3.5: 192, // Kepler Generation (SM 3.5) GK11x class
		3.7: 192, // Kepler Generation (SM 3.7) GK21x class
		5.0: 128, // Maxwell Generation (SM 5.0) GM10x class
		5.2: 128, // Maxwell Generation (SM 5.2) GM20x class
		5.3: 128, // Maxwell Generation (SM 5.3) GM20x class
		6.0: 64,  // Pascal Generation (SM 6.0) GP100 class
		6.1: 128, // Pascal Generation (SM 6.1) GP10x class
		6.2: 128, // Pascal Generation (SM 6.2) GP10x class
		7.0: 64,  // Volta Generation (SM 7.0) GV100 class
		7.2: 64,  // Volta Generation (SM 7.2) GV11b class
		7.5: 64,  // Turing Generation (SM 7.5) TU1xx class
	}
)

func (g GPU) NumCores() (int, error) {
	sms, err := g.NumSMs()
	if err != nil {
		return 0, errors.Wrap(err, "unable to get num cores because cannot get num sms")
	}

	computeCapability, err := g.ComputeCapability()
	if err != nil {
		return 0, errors.Wrap(err, "unable to get num cores because cannot get compute capability")
	}

	cores, ok := archSMMapping[computeCapability]
	if ok {
		return cores * sms, nil
	}

	return 0, errors.Errorf("unable to find num of cores for compute capability %v", computeCapability)
}

func (g GPU) NumSMs() (int, error) {
	name := strings.ToLower(g.ProductName)
	for _, info := range NvidiaGPUs {
		if strings.ToLower(info.Name) == name {
			if info.NumSMs == 0 {
				panic("expecting a non-zero num sms")
			}
			return info.NumSMs, nil
		}
	}
	return 0, errors.Errorf("cannot find num of sms for %v", g.ProductName)
}

func (g GPU) Architecture() (string, error) {
	name := strings.ToLower(g.ProductName)
	for _, info := range NvidiaGPUs {
		if strings.ToLower(info.Name) == name {
			if info.Architecture == "" {
				panic("expecting a non-zero architecture")
			}
			return info.Architecture, nil
		}
	}
	return "", errors.Errorf("cannot find architecture name for %v", g.ProductName)
}

func (g GPU) ClockRate() (int64, error) {
	name := strings.ToLower(g.ProductName)
	for _, info := range NvidiaGPUs {
		if strings.ToLower(info.Name) == name {
			if info.ClockRate == 0 {
				panic("expecting a non-zero clock rate")
			}
			return info.ClockRate, nil
		}
	}
	return 0, errors.Errorf("cannot find clock rate for %v", g.ProductName)
}

func (g GPU) ComputeCapability() (float64, error) {
	name := strings.ToLower(g.ProductName)
	for _, info := range NvidiaGPUs {
		if strings.ToLower(info.Name) == name {
			if info.ComputeCapability == 0 {
				panic("expecting a non-zero compute capability")
			}
			return info.ComputeCapability, nil
		}
	}
	return 0, errors.Errorf("cannot find compute capability for %v", g.ProductName)
}

func (g GPU) TheoreticalGFlops() (int64, error) {
	name := strings.ToLower(g.ProductName)
	for _, info := range NvidiaGPUs {
		if strings.ToLower(info.Name) == name {
			if info.PeekGFlops == 0 {
				panic("expecting a non-zero peek gflops")
			}
			return info.PeekGFlops, nil
		}
	}
	return 0, errors.Errorf("cannot find theoretical peak gflops for %v", g.ProductName)
}

func (g GPU) MemoryBandwidth() (float64, error) {
	name := strings.ToLower(g.ProductName)
	for _, info := range NvidiaGPUs {
		if strings.ToLower(info.Name) == name {
			if info.MemoryBandwidth == 0 {
				panic("expecting a non-zero memory bandwidth")
			}
			return info.MemoryBandwidth, nil
		}
	}
	return 0, errors.Errorf("cannot find memory bandwidth for %v", g.ProductName)
}

func (g GPU) InterconnectName() (string, error) {
	name := strings.ToLower(g.ProductName)
	for _, info := range NvidiaGPUs {
		if strings.ToLower(info.Name) == name {
			if info.Interconnect.Name() == "" {
				panic("expecting a non-zero name")
			}
			return info.Interconnect.Name(), nil
		}
	}
	return "", errors.Errorf("cannot find interconnect name for %v", g.ProductName)
}

func (g GPU) InterconnectBandwidth() (float64, error) {
	name := strings.ToLower(g.ProductName)
	for _, info := range NvidiaGPUs {
		if strings.ToLower(info.Name) == name {
			if info.Interconnect.Bandwidth() == 0 {
				panic("expecting a non-zero bandwidth")
			}
			return info.Interconnect.Bandwidth(), nil
		}
	}
	return 0, errors.Errorf("cannot find interconnect bandwidth for %v", g.ProductName)
}
