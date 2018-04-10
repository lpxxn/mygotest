package interviewlogger

import (
	"go.uber.org/zap"
	"sync"
	"fmt"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"time"
	"github.com/mygotest/workspace/gindemo/demo1/functions"
	"path"
)

type InterViewLogger struct {
	*zap.Logger
}

var singleInterviewLogger *InterViewLogger = nil
var interviewOnce sync.Once


func InitInterViewLogger() *InterViewLogger {
	interviewOnce.Do(func() {
		work_path, err := funcs.GetCurrentDirectory()
		if nil != err {
			panic(err)
		}

		log_path := path.Join(work_path, "logs/log.log")
		fmt.Println("interview log path:", log_path)

		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   log_path,
			MaxSize:    20, // megabytes
			MaxBackups: 20,
			MaxAge:     10, // days
		})

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			w,
			zap.InfoLevel,
		)

		singleInterviewLogger = &InterViewLogger{Logger:zap.New(core)}

	})
	return singleInterviewLogger
}

func logWithTime(fields ...zapcore.Field) []zapcore.Field {
	fields = append(fields, zap.String("curenttime", time.Now().Format(time.RFC3339)))
	return fields
}

func LogInterView(msg string, fields ...zapcore.Field)  {
	fields = logWithTime(fields...)
	InitInterViewLogger().Info(msg, fields...)
}



