// Code generated by "stringer -type=PacketType"; DO NOT EDIT.

package types

import "strconv"

const _PacketType_name = "PingRPCCascadePulseGetRandomHostsBootstrapGetNonceAuthorizeDisconnectPhase1Phase2Phase3"

var _PacketType_index = [...]uint8{0, 4, 7, 14, 19, 33, 42, 50, 59, 69, 75, 81, 87}

func (i PacketType) String() string {
	i -= 1
	if i < 0 || i >= PacketType(len(_PacketType_index)-1) {
		return "PacketType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _PacketType_name[_PacketType_index[i]:_PacketType_index[i+1]]
}
