package rescue

import (
	"runtime"

	"github.com/go-kratos/kratos/v2/log"
)

// Recover is used with defer to do cleanup on panics.
// Use it like:
//  defer Recover(func() {})
func Recover(cleanups ...func()) {
	for _, cleanup := range cleanups {
		cleanup()
	}

	if p := recover(); p != nil {

		logger := log.NewHelper(log.GetLogger())
		buf := make([]byte, 64<<10) //nolint:gomnd
		n := runtime.Stack(buf, false)
		buf = buf[:n]
		logger.Errorf("%v:\n%s\n", p, buf)
	}
}
