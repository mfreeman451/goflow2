package debug

import (
	"context"
	"runtime/debug"

	"github.com/mfreeman451/goflow2/v2/utils"
)

func PanicDecoderWrapper(wrapped utils.DecoderFunc) utils.DecoderFunc {
	return func(ctx context.Context, msg interface{}) (err error) {
		defer func() {
			if pErr := recover(); pErr != nil {
				pErrC, _ := pErr.(string)
				err = &PanicErrorMessage{Msg: msg, Inner: pErrC, Stacktrace: debug.Stack()}
			}
		}()

		err = wrapped(ctx, msg)

		return err
	}
}
