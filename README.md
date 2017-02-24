
#NVIDIA-SMI [![Build Status](https://travis-ci.org/rai-project/nvidia-smi.svg?branch=master)](https://travis-ci.org/rai-project/nvidia-smi)

## Notes

Using the output xml of `nvidia-smi -x  -q -a`.
Use github.com/rai-project/xj2s to create structure and then Use

~~~
gomodifytags -file model.go -add-tags json -w -struct NvidiaSmiLog
gomodifytags -file model.go -add-tags json -w -struct Gpu
~~~


to add the json tags.
