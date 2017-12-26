// +build linux,arm64

package smi

import (
	"github.com/Unknwon/com"
)

func isTegra() bool {
	return com.IsFile("/etc/nv_tegra_release")
}

func Init() {
	defer close(initialized)
	if !isTegra() {
		HasGPU = false
		return
	}

}
