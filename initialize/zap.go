package initialize

import (
	"gin-practice/config"
	"os"
	"path/filepath"
	"strings"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func dirCheck(path string) error {
	dir := path
	if !strings.HasSuffix(path, "/") && filepath.Ext(path) != "" {
		dir = filepath.Dir(path)
	}
	return os.MkdirAll(dir, 0777)
}

func getLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func encodeConfig() zapcore.EncoderConfig {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	return encoderConfig
}

func GetZap() *zap.Logger {
	z := config.SELF_CONFIG.Zap
	dirCheck(z.Path)
	hook := lumberjack.Logger{
		Filename:   z.Path,
		MaxSize:    z.MaxSize,
		MaxBackups: z.MaxBackup,
		MaxAge:     z.MaxAge,
		Compress:   z.Compress,
	}
	level := getLevel(z.Level)
	write := zapcore.AddSync(&hook)

	encoder := encodeConfig()

	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoder),
		write,
		level,
	)
	// 开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 文件及行号
	development := zap.Development()
	logger := zap.New(core, caller, development)
	return logger
}
