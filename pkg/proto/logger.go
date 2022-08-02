package proto

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/kyledinh/protoserver-go/pkg/proto/sys"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Logger(ctx context.Context) *zap.Logger {
	if logger, ok := ctx.Value(sys.LOG).(*zap.Logger); ok {
		return logger
	}
	return DefaultLogger()
}

func DefaultLogger() *zap.Logger {
	return zap.L()
}

func SetupLogger(ctx context.Context, serviceName string) {
	if !viper.IsSet("log.level") {
		viper.SetDefault("log.level", "info")
	}
	if !viper.IsSet("log.type") {
		viper.SetDefault("log.level", "stderr")
	}
	if !viper.IsSet("log.format") {
		viper.SetDefault("log.format", "json")
	}
	if logger, err := GetLogger(serviceName); err != nil {
		panic(err.Error())
	} else {
		zap.ReplaceGlobals(logger)
	}
}

func GetLogger(serviceName string, opts ...zap.Option) (*zap.Logger, error) {
	logCfg := zap.Config{
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "timestamp",
			LevelKey:       "level",
			NameKey:        "name",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stack",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     iso8601TimeEncoderMilli,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	switch f := strings.ToLower(viper.GetString("log.format")); f {
	case "json":
		logCfg.Encoding = "json"
	case "text":
		logCfg.Encoding = "console"
	default:
		return nil, errors.New(fmt.Sprint("Unrecognized log format ", f))
	}

	switch l := strings.ToLower(viper.GetString("log.level")); l {
	case "panic":
		logCfg.Level = zap.NewAtomicLevelAt(zap.PanicLevel)
	case "fatal":
		logCfg.Level = zap.NewAtomicLevelAt(zap.FatalLevel)
	case "error":
		logCfg.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "warn":
		logCfg.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "info":
		logCfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "debug":
		logCfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	default:
		return nil, errors.New(fmt.Sprint("Unrecognized log level ", l))
	}

	switch t := strings.ToLower(viper.GetString("log.type")); t {
	case "file":
		if viper.IsSet("log.file") {
			logCfg.OutputPaths = []string{viper.GetString("log.file")}
		} else {
			return nil, errors.New("log.file in configuration must specify a file to log to")
		}
	case "stderr":
	case "debug":
		opt := zap.WrapCore(func(z zapcore.Core) zapcore.Core {
			enc := zapcore.NewJSONEncoder(logCfg.EncoderConfig)
			buf := viper.Get("log.buffer")
			return zapcore.NewCore(enc, zapcore.AddSync(buf.(io.Writer)), logCfg.Level)
		})
		logCfg.OutputPaths = []string{}
		opts = append(opts, opt)
	default:
		return nil, errors.New(fmt.Sprint("Unrecognized log type ", t))
	}

	logger, err := logCfg.Build(opts...)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Unable to initialize logger")
		return nil, err
	}

	logger = logger.With(zap.String("servicename", serviceName))
	return logger, nil
}

func iso8601TimeEncoderMilli(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02T15:04:05.000000Z0700"))
}
