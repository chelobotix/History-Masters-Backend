package service

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
)

// CustomEncoder to handle ANSI characters
type CustomEncoder struct {
	zapcore.Encoder
}

func (e *CustomEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	return e.Encoder.EncodeEntry(entry, fields)
}

func NewZapLogger(e *echo.Echo) (*zap.Logger, error) {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// Crear configuración personalizada
	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:      true,
		Encoding:         "console",
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		return nil, err
	}

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:     true,
		LogStatus:  true,
		LogMethod:  true,
		LogLatency: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			message := fmt.Sprintf(
				"%sMethod:%s %s | %sURI:%s %s | %sStatus:%s %d | %sLatency:%s %s | %sIP:%s %s",
				Blue, Reset, v.Method,
				Green, Reset, v.URI,
				getStatusColor(v.Status), Reset, v.Status,
				Yellow, Reset, v.Latency,
				Cyan, Reset, c.RealIP(),
			)

			logger.Info(message)
			return nil
		},
	}))

	return logger, nil
}

func getStatusColor(status int) string {
	switch {
	case status >= 500:
		return Red
	case status >= 400:
		return Yellow
	case status >= 300:
		return Cyan
	case status >= 200:
		return Green
	default:
		return Blue
	}
}

// Función auxiliar para formatear mensajes
func formatLogMessage(key, value string, color string) string {
	return fmt.Sprintf("%s%s:%s %s", color, key, Reset, value)
}

// Función auxiliar para crear un mensaje coloreado
func colorizedLog(level, msg string) string {
	switch level {
	case "ERROR":
		return fmt.Sprintf("%s%s%s", Red, msg, Reset)
	case "WARN":
		return fmt.Sprintf("%s%s%s", Yellow, msg, Reset)
	case "INFO":
		return fmt.Sprintf("%s%s%s", Green, msg, Reset)
	case "DEBUG":
		return fmt.Sprintf("%s%s%s", Blue, msg, Reset)
	default:
		return msg
	}
}
