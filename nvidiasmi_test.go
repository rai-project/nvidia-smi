package nvidiasmi

import (
	"encoding/xml"
	"testing"

	rice "github.com/GeertJohan/go.rice"
	"github.com/k0kubun/pp"
	"github.com/stretchr/testify/assert"
)

var (
	box         = rice.MustFindBox("_fixtures")
	example1XML = box.MustBytes("example1.xml")
	example2XML = box.MustBytes("example2.xml")
)

func TestUnmarshal(t *testing.T) {
	var info NvidiaSmi
	err := xml.Unmarshal(example1XML, &info)
	assert.NoError(t, err)
	assert.NotEmpty(t, info)
}

func TestNew(t *testing.T) {
	info, err := New()
	if err == nil {
		assert.NotNil(t, info)
		pp.Println(info)
	}
}
