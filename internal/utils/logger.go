package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func SugarLog() *zap.SugaredLogger {
	writerSyncer := getLogWriter()
	encoder := getEncoder()

	// Encoder: Format log
	// WriterSyncer: Log to file (default: {"level":"info","ts":1632184403.8505602,"msg":"Test log info"})
	// Level log: Debug will include all level
	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)
	logger := zap.New(core)
	return logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	// default
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./logs.log")
	return zapcore.AddSync(file)
}
