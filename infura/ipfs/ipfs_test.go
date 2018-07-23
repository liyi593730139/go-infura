package ipfs

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCanGetObject(t *testing.T) {
	c := NewClient()
	objectHash := "QmZL4eT1gxnE168Pmw3KyejW6fUfMNzMgeKMgcWJUfYGRj"
	o, err := c.GetObject(objectHash)
	require.Nil(t, err)
	require.NotNil(t, o)
	require.NotEmpty(t, o.Links)
	assert.Equal(t, "Placeholder Thesis Summary.pdf", o.Links[0].Name)
}

func TestCanGetNodeInfo(t *testing.T) {
	c := NewClient()
	n, err := c.NodeInfo()
	require.Nil(t, err)
	require.NotNil(t, n)
	assert.NotEmpty(t, n.Addresses)
	assert.Equal(t, "ipfs/0.1.0", n.ProtocolVersion)
}
