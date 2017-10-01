package nvidiasmi

import (
	"os/exec"

	"encoding/xml"

	"github.com/pkg/errors"
)

const (
	BinaryName = "nvidia-smi"
)

var (
	Info     *NvidiaSmi
	HasGPU   bool = false
	GPUCount      = 0
)

func New() (*NvidiaSmi, error) {
	exe, err := exec.LookPath(BinaryName)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot find %v in path", BinaryName)
	}
	cmd := exec.Command(exe, "-x", "-q", "-a")
	bts, err := cmd.CombinedOutput()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to run %v -x -q -a", exe)
	}
	res := new(NvidiaSmi)
	err = xml.Unmarshal(bts, res)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal %v -x -q -a", exe)
	}
	return res, nil
}
