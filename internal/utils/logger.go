package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

func SugarLog() *zap.SugaredLogger {
	writerSyncer := getLogWriter()
	encoder := getEncoder()

	// Encoder: Format log
	// WriterSyncer: Log to file (default: {"level":"info","ts":1632184403.8505602,"msg":"Test log info"})
	// Level log: Debug will include all level
	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	return logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey: "message",
		TimeKey:    "time",
		LevelKey:   "level",
		CallerKey:  "caller",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeLevel: func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString("[" + level.CapitalString() + "]")
		},
		EncodeCaller: zapcore.ShortCallerEncoder,
	})
}

func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("./logs.log")
	return zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(os.Stdout),
		zapcore.AddSync(file),
	)
}
