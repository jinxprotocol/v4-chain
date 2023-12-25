package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPriceSmoothingPpm(t *testing.T) {
	require.Equal(t, PriceSmoothingPpm, uint32(300_000))
}
