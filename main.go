package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/cors"
	hertzlogrus "github.com/hertz-contrib/logger/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"os"
	"path"
	init_ "teaching_evaluation_backend/init"
	"teaching_evaluation_backend/middle"

	"time"
)

const (
	LogFileAddr = "./hlog/logs/"
)

func main() {
	ctx := context.Background()

	if err := initLog(); err != nil {
		log.Fatalf("init log error: %v", err)
	}

	if err := init_.Init(ctx); err != nil {
		hlog.CtxFatalf(ctx, "init failed: %v", err)
	}

	h := server.Default(
		server.WithHandleMethodNotAllowed(true),
	)

	h.Use(recovery.Recovery(recovery.WithRecoveryHandler(RecoveryHandler)))
	h.Use(middle.LoggingMiddleware())
	h.Use(middle.JWTAuthMiddleware())
	h.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		//超时时间设定
		MaxAge: 24 * time.Hour,
	}))

	register(h)
	h.Spin()
}

func initLog() error {
	// 创建日志目录
	if err := os.MkdirAll(LogFileAddr, 0o755); err != nil {
		return err
	}

	// 生成日志文件名（按天）
	logFileName := time.Now().Format("2006-01-02") + ".log"
	fileName := path.Join(LogFileAddr, logFileName)

	// 配置lumberjack
	lumberjackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    20,   // 单个文件最大20M
		MaxBackups: 5,    // 保留5个备份
		MaxAge:     10,   // 日志保存10天
		Compress:   true, // 压缩备份
	}

	logger := hertzlogrus.NewLogger()
	logger.SetOutput(lumberjackLogger)
	logger.SetLevel(hlog.LevelDebug)
	hlog.SetLogger(logger)

	multiWriter := io.MultiWriter(os.Stdout, lumberjackLogger)
	logger.SetOutput(multiWriter)

	logger.SetLevel(hlog.LevelDebug)
	hlog.SetLogger(logger)

	return nil
}

func RecoveryHandler(ctx context.Context, c *app.RequestContext, err interface{}, stack []byte) {
	hlog.SystemLogger().CtxErrorf(ctx, "[Recovery] err=%v\nstack=%s", err, stack)
	hlog.SystemLogger().Infof("Client: %s", c.Request.Header.UserAgent())
	c.AbortWithStatus(consts.StatusInternalServerError)
}
