package zaplogger

import (
	"fmt"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"time"
	"github.com/mygotest/workspace/gindemo/demo1/functions"
)

var singletonZap *zap.Logger = nil
var zapOnce sync.Once

func InitLogger() *zap.Logger {
	zapOnce.Do(func() {
		workPath, err := funcs.GetCurrentDirectory()
		if nil != err {
			panic(err)
		}
		logPath := funcs.AppFilePath(workPath, "github.com/mygotest/workspace/gindemo", "logs", "udc.log", false)


		fmt.Println("log path : ", logPath)
		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   logPath,
			MaxSize:    1, // megabytes
			MaxBackups: 20,
			MaxAge:     10, // days
		})

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			w,
			zap.InfoLevel,
		)
		singletonZap = zap.New(core)
	})
	return singletonZap
}

func Panic(msg string, fields ...zapcore.Field) {
	fields = logPostfix(fields...)
	InitLogger().Panic(msg, fields...)
}
func logPostfix(fields ...zapcore.Field) []zapcore.Field {
	fields = append(fields, zap.String("curenttime", time.Now().Format(time.RFC3339)))
	return fields
}

func Info(msg string, fields ...zapcore.Field) {
	fields = logPostfix(fields...)
	InitLogger().Info(msg, fields...)
}

func Error(msg string, fields ...zapcore.Field) {
	fields = logPostfix(fields...)
	InitLogger().Error(msg, fields...)
}