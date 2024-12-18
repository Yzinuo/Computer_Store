package comstore

import (
	g "computer_store/internal/global"
	"log/slog"
	"os"
	"time"
)

func InitLogger(conf *g.Config) {
	var level slog.Level
	switch conf.Log.Level {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelDebug
	}

	option := &slog.HandlerOptions{
		AddSource: false, //是否显示源文件的信息
		Level:     level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				if t, ok := a.Value.Any().(time.Time); ok {
					a.Value = slog.StringValue(t.Format(time.DateTime))
				}
			}
			return a
		},
	}

	var handler slog.Handler
	switch conf.Log.Format {
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, option)
	case "text":
		handler = slog.NewTextHandler(os.Stdout, option)
	default:
		handler = slog.NewTextHandler(os.Stdout, option)
	}

	logger := slog.New(handler)
	slog.SetDefault(logger) // 设定默认的slog
}
