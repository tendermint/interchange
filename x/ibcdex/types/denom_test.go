package types_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tendermint/interchange/x/ibcdex/types"
	"testing"
)

func TestOriginalDenom(t *testing.T) {
	port, channel, denom := "port-0", "channel-0", "foo"
	port2, channel2, denom2 := "port-1", "channel-1", "bar"
	voucher1 := types.VoucherDenom(port, channel, denom)
	voucher2 := types.VoucherDenom(port2, channel2, denom)
	voucher3 := types.VoucherDenom(port, channel, denom2)

	// VoucherDenom generates different denoms
	require.NotEqual(t, voucher1, voucher2)
	require.NotEqual(t, voucher1, voucher3)

	// OriginalDenom gets back the original denom
	_, isOrigin := types.OriginalDenom(port2, channel2, voucher1)
	require.False(t, isOrigin)

	origin, isOrigin := types.OriginalDenom(port, channel, voucher1)
	require.True(t, isOrigin)
	require.Equal(t, denom, origin)
}