package logger

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"teaching_evaluation_backend/consts"
	"time"
)

// LoggingMiddleware 打印访问日志，包括请求信息和响应状态及响应体
func LoggingMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		start := time.Now()

		// 生成 traceId
		traceID := string(c.GetHeader(consts.TraceIDHeader))
		if traceID == "" {
			timeStr := time.Now().Format("20060102150405")

			randomBytes := make([]byte, 4)
			_, err := rand.Read(randomBytes)
			if err != nil {
				hlog.Error("trace id err:", err)
			}
			randomHex := hex.EncodeToString(randomBytes)

			traceID = timeStr + "-" + randomHex
		}

		bodyBytes := c.Request.Body()
		c.Request.SetBody(bodyBytes)

		var reqBody interface{}
		if len(bodyBytes) > 0 {
			_ = json.Unmarshal(bodyBytes, &reqBody)
		} else {
			reqBody = "{}"
		}

		hlog.CtxInfof(ctx, "[%s] 请求开始: method=%s path=%s client=%s request=%+v",
			traceID, c.Request.Method(), c.Request.URI().String(), c.ClientIP(), reqBody)

		c.Header("X-Trace-ID", traceID)

		c.Next(ctx)

		duration := time.Since(start)

		respBody := string(c.Response.Body())

		hlog.CtxInfof(ctx, "[%s] 请求结束: method=%s path=%s status=%d duration=%v response=%s",
			traceID, c.Request.Method(), c.Request.URI().String(), c.Response.StatusCode(), duration, respBody)
	}
}
