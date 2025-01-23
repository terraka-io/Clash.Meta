package delay

import "github.com/metacubex/mihomo/common/atomic"

var IsDelayServer = atomic.NewBool(false)
var TunFileDescriptor = atomic.NewInt32(0)
