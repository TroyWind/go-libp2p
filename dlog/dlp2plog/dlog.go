package dlp2plog

import (
	"github.com/libp2p/go-libp2p/util"
	"go.uber.org/zap"
)

var L *zap.Logger

func init() {
	L = util.GetXDebugLog("lp2p")
}
