package main

import (
	"go.uber.org/zap"
)

func main() {
	// custom logger with zap
	//	config := zap.Config{
	//		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
	//		Development:      true,
	//		Encoding:         "json",
	//		InitialFields:    map[string]interface{}{"MyName": "wuji_logging"},
	//		OutputPaths:      []string{"stdout"},
	//		ErrorOutputPaths: []string{"stdout"},
	//	}
	//	logger, _ := config.Build()

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("info logging")
	logger.Info("format string and type", zap.String("name", "x"), zap.Int("age", 20))

	sugar := logger.Sugar()
	sugar.Infof("sugar for %s", "programing")
	sugar.Infow("sugar for kv", "name", "wujimaster", "age", "29")

	name := "wuji"
	zap.S().Info("global env: name =", name)

}
