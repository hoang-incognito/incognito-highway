package main

import (
	"highway/chain"
	"highway/process"
	"highway/process/topic"
	"highway/route"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger

func initLogger(level string) {
	cf := zap.NewDevelopmentConfig()
	cf.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	// Set loglevel
	switch strings.ToLower(level) {
	case "info":
		cf.Level.SetLevel(zapcore.InfoLevel)
	case "debug":
		cf.Level.SetLevel(zapcore.DebugLevel)
	}

	l, _ := cf.Build()
	logger = l.Sugar()

	// Initialize children's loggers
	chain.InitLogger(logger)
	route.InitLogger(logger)
	process.InitLogger(logger)
	topic.InitLogger(logger)
}