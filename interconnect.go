package nvidiasmi

const (
	Gbps = float64(1)
	GBps = float64(8)
)

type Interconnect interface {
	Name() string
	Bandwidth() float64
}

type Infiniband struct {
}

func (Infiniband) Name() string {
	return "Infiniband"
}

func (Infiniband) Bandwidth() float64 {
	return float64(70) * Gbps
}

type HBM3 struct {
}

func (HBM3) Name() string {
	return "HBM3"
}

func (HBM3) Bandwidth() float64 {
	panic("HBM3 not implemented")
	return 1.0
}

type HBM2 struct {
}

func (HBM2) Name() string {
	return "HBM2"
}

func (HBM2) Bandwidth() float64 {
	panic("HBM2 not implemented")
	return 1.0
}

type Ethernet struct {
}

func (Ethernet) Name() string {
	return "Ethernet"
}

func (Ethernet) Bandwidth() float64 {
	return float64(10) * Gbps
}

type Ethernet20 struct {
}

func (Ethernet20) Name() string {
	return "Ethernet20"
}

func (Ethernet20) Bandwidth() float64 {
	return float64(20) * Gbps
}

type NVLink1 struct {
}

func (NVLink1) Name() string {
	return "NVLink1"
}

func (NVLink1) Bandwidth() float64 {
	return float64(80) * GBps
}

type NVLink2 struct {
}

func (NVLink2) Name() string {
	return "NVLink2"
}

func (NVLink2) Bandwidth() float64 {
	return float64(300) * GBps
}

// Here we assume PCIe x16.
// PCIe 1.0: 150MB/s per lane per direction.
type PCIe1 struct {
}

func (PCIe1) Name() string {
	return "PCIe2"
}

func (PCIe1) Bandwidth() float64 {
	return 2.4 * float64(GBps)
}

// Here we assume PCIe x16.
// PCIe 1.0: 500MB/s per lane per direction.
type PCIe2 struct {
}

func (PCIe2) Name() string {
	return "PCIe2"
}

func (PCIe2) Bandwidth() float64 {
	return float64(8) * GBps
}

// Here we assume PCIe x16.
// PCIe 3.0: 1GB/s per lane per direction.
type PCIe3 struct {
}

func (PCIe3) Name() string {
	return "PCIe3"
}

func (PCIe3) Bandwidth() float64 {
	return float64(16) * GBps
}

// Here we assume PCIe x16.
// PCIe 3.0: 1GB/s per lane per direction.
type SXM2 struct {
}

func (SXM2) Name() string {
	return "SXM2"
}

func (SXM2) Bandwidth() float64 {
	panic("SXM2 bandwidth not implemented")
	return 0
}
