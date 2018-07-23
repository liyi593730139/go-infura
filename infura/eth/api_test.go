package eth

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCanGetAccountBalance(t *testing.T) {
	c := NewClient(Mainnet)
	r, err := c.GetAccountBalance("0x398137383b3d25c92898c656696e41950e47316b", Latest)
	require.Nil(t, err)
	assert.NotEmpty(t, r)
}

func TestCanGetBlockNumber(t *testing.T) {
	c := NewClient(Mainnet)
	r, err := c.GetBlockNumber()
	require.Nil(t, err)
	assert.NotEmpty(t, r)
}
