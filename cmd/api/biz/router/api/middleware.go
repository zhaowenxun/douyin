// Code generated by hertz generator.

package Api

import (
	"context"
	"fmt"

	// "github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/api/biz/mw"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/requestid"
	"go.opentelemetry.io/otel/trace"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		// use recovery mw
		recovery.Recovery(recovery.WithRecoveryHandler(
			func(ctx context.Context, c *app.RequestContext, err interface{}, stack []byte) {
				hlog.SystemLogger().CtxErrorf(ctx, "[Recovery] err=%v\nstack=%s", err, stack)
				c.JSON(consts.StatusInternalServerError, utils.H{
					"code":    errno.ServiceErr.ErrCode,
					"message": fmt.Sprintf("[Recovery] err=%v\nstack=%s", err, stack),
				})
			},
		)),
		// use requestid mw
		requestid.New(
			requestid.WithGenerator(func(ctx context.Context, c *app.RequestContext) string {
				traceID := trace.SpanFromContext(ctx).SpanContext().TraceID().String()
				return traceID
			}),
		),
		// use gzip mw
		gzip.Gzip(gzip.DefaultCompression),
	}
}

func _v2Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _loginuserMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _registeruserMw() []app.HandlerFunc {
	// your code...
	return nil
}
