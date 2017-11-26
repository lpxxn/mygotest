package zaplogger

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var singletonZap *zap.Logger = nil
var zapOnce sync.Once

func InitLogger() *zap.Logger {
	zapOnce.Do(func() {
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		exPath := filepath.Dir(ex)
		fmt.Println(exPath)

		logPath := path.Join(exPath, "logs", "crm.log")
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



func Panic(msg string, fields ...zapcore.Field){
	InitLogger().Panic(msg, fields...)
}
