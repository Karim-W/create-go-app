package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitOrDie initializes the logger and panics if it fails.
func InitOrDie() *zap.Logger {
	encoderCfg := zap.NewProductionEncoderConfig()

	encoderCfg.TimeKey = "ts"

	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:       false,
		DisableCaller:     true,
		DisableStacktrace: true,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			"stdout",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
		// InitialFields: map[string]interface{}{
		// 	"pid": os.Getpid(),
		// },
	}

	logginInstance := zap.Must(config.Build())

	zap.ReplaceGlobals(logginInstance)

	return logginInstance
}
