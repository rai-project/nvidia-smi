// +build nogpu

package nvidiasmi

func Init() {
	defer close(initialized)
	HasGPU = false
	Info = nil
	GPUCount = 0
}
