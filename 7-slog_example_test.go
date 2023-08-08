package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"time"
)

// The new log/slog package provides structured logging with levels.
// Structured logging emits key-value pairs to enable fast, accurate processing of large amounts of log data.
// The package supports integration with popular log analysis tools and services.
//
// https://tip.golang.org/doc/go1.21#slog

var opts = &slog.HandlerOptions{
	ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
		// remove time
		if a.Key == slog.TimeKey && len(groups) == 0 {
			return slog.Attr{}
		}
		return a
	},
	Level: slog.LevelDebug,
}

func Example_slog_TextHandler() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))

	logger.Debug("debugging", "is", "fun")

	logger.Info("hello", "world", 42)

	logger.Warn("hello", slog.String("name", "world"))

	logger.With(
		"a", 1,
		"b", 2,
	).Error("something went wrong",
		slog.Duration("duration", 42*time.Second),
		slog.Int("c", 3),
	)

	// Output:
	// level=DEBUG msg=debugging is=fun
	// level=INFO msg=hello world=42
	// level=WARN msg=hello name=world
	// level=ERROR msg="something went wrong" a=1 b=2 duration=42s c=3
}

type Password string

func (Password) LogValue() slog.Value {
	return slog.StringValue("[SECRET]")
}

func Example_slog_secret() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))

	p := Password("super password!")

	logger.Info("demo", "password", p)

	// Output:
	// level=INFO msg=demo password=[SECRET]
}

type User struct {
	Name  string
	Passw Password
}

func (u User) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("name", u.Name),
		slog.Any("password", u.Passw),
	)
}

func Example_slog_group() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))

	u := User{
		Name:  "user",
		Passw: Password("super password!"),
	}

	logger.Info("demo", "user", u)

	// Output:
	// level=INFO msg=demo user.name=user user.password=[SECRET]
}

func Example_slog_JSONHanlder() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))

	u := User{
		Name:  "user",
		Passw: Password("super password!"),
	}

	logger.Info("demo", "user", u)

	// Output:
	// {"level":"INFO","msg":"demo","user":{"name":"user","password":"[SECRET]"}}
}

func Example_slog_override_log() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))

	// override log output
	slog.SetDefault(logger)

	log.Print("hello")

	// Output:
	// {"level":"INFO","msg":"hello"}
}

func ServiceName(name string) slog.Attr {
	return slog.String("service", name)
}

func Example_slog_custom_attr() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, opts))

	logger = logger.With(ServiceName("my-service"))

	logger.Info("hello")

	// Output:
	// {"level":"INFO","msg":"hello","service":"my-service"}
}

//--- context data propagation

type metadataKey struct{}

func WithMetadata(ctx context.Context, metadata map[string]any) context.Context {
	return context.WithValue(ctx, metadataKey{}, metadata)
}

type MetadataHanlder struct {
	h slog.Handler
}

func (m *MetadataHanlder) Enabled(ctx context.Context, level slog.Level) bool {
	return m.h.Enabled(ctx, level)
}

func (m *MetadataHanlder) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &MetadataHanlder{
		h: m.h.WithAttrs(attrs),
	}
}

func (m *MetadataHanlder) WithGroup(name string) slog.Handler {
	return &MetadataHanlder{
		h: m.h.WithGroup(name),
	}
}

func (m *MetadataHanlder) Handle(ctx context.Context, r slog.Record) error {
	if metadata, ok := ctx.Value(metadataKey{}).(map[string]any); ok {
		r = r.Clone()

		for k, v := range metadata {
			r.Add(k, v)
		}
	}

	return m.h.Handle(ctx, r)
}

func Example_slog_context() {
	ctx := context.Background()

	ctx = WithMetadata(ctx, map[string]any{
		"request_id": "123",
	})

	logger := slog.New(&MetadataHanlder{
		h: slog.NewTextHandler(os.Stdout, opts),
	})

	logger.InfoContext(ctx, "request started")

	// Output:
	// level=INFO msg="request started" request_id=123
}
