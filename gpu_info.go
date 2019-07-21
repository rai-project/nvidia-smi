package nvidiasmi

func (g GPU) NumSMs() int64 {
	info := map[string]int64{
		"TITAN Xp":             -1,
		"Tesla V100-SXM2-16GB": 80,
	}
	if val, ok := info[g.ProductName]; ok {
		return val
	}
	return 0
}

func (g GPU) MemoryClock() int64 {
	info := map[string]int64{
		"TITAN Xp":             5705,
		"Tesla V100-SXM2-16GB": 877,
	}
	if val, ok := info[g.ProductName]; ok {
		return val
	}
	return 0
}

func (g GPU) Frequency() int64 {
	info := map[string]int64{
		"TITAN Xp":             1582,
		"Tesla V100-SXM2-16GB": 1530,
	}
	if val, ok := info[g.ProductName]; ok {
		return val
	}
	return 0
}

func (g GPU) ComputeCapability() float64 {
	info := map[string]float64{
		"TITAN Xp":             6.1,
		"Tesla V100-SXM2-16GB": 7.0,
	}
	if val, ok := info[g.ProductName]; ok {
		return val
	}
	return 0
}

func (g GPU) TheoreticalFlops() float64 {
	return 0
}

func (g GPU) Bandwidth() uint {
	return 0
}
