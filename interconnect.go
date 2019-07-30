package nvidiasmi

const (
	GBps = float64(1)
	Gbps = float64(0.125)
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
	return float64(40) * Gbps
}

type HBM3 struct {
}

func (HBM3) Name() string {
	return "HBM3"
}

func (HBM3) Bandwidth() float64 {
	return float64(512) * GBps
}

type HBM2 struct {
}

func (HBM2) Name() string {
	return "HBM2"
}

func (HBM2) Bandwidth() float64 {
	return float64(256) * GBps
}

type Ethernet struct {
}

func (Ethernet) Name() string {
	return "Ethernet"
}

func (Ethernet) Bandwidth() float64 {
	return float64(10) * GBps
}

type Ethernet20 struct {
}

func (Ethernet20) Name() string {
	return "Ethernet20"
}

func (Ethernet20) Bandwidth() float64 {
	return float64(20) * GBps
}

type NVLink1 struct {
}

func (NVLink1) Name() string {
	return "NVLink1"
}

func (NVLink1) Bandwidth() float64 {
	return float64(160) * GBps
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
	return "PCIe1"
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

type SXM2 struct {
}
