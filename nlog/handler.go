package nlog

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"
	"sync"

	"github.com/josh-silvas/nbot/shared"
)

// H : custom Handler for log messages
type H struct {
	handler slog.Handler
	record  func([]string, slog.Attr) slog.Attr
	buffer  *bytes.Buffer
	mutex   *sync.Mutex
}

// cliHandler : create a new H for CLI output
func cliHandler(opts *slog.HandlerOptions) *H {
	if opts == nil {
		opts = &slog.HandlerOptions{}
	}
	b := new(bytes.Buffer)
	return &H{
		buffer: b,
		handler: slog.NewJSONHandler(b, &slog.HandlerOptions{
			Level:       opts.Level,
			AddSource:   opts.AddSource,
			ReplaceAttr: replAttrWrap(opts.ReplaceAttr),
		}),
		record: opts.ReplaceAttr,
		mutex:  new(sync.Mutex),
	}
}

// Enabled : check if a log level is enabled
func (h *H) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

// WithAttrs : add attributes to a log message
func (h *H) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &H{handler: h.handler.WithAttrs(attrs), buffer: h.buffer, record: h.record, mutex: h.mutex}
}

// WithGroup : add a group to a log message
func (h *H) WithGroup(name string) slog.Handler {
	return &H{handler: h.handler.WithGroup(name), buffer: h.buffer, record: h.record, mutex: h.mutex}
}

// attrs : get attributes from a log message
func (h *H) attrs(ctx context.Context, r slog.Record) (map[string]any, error) {
	h.mutex.Lock()
	defer func() {
		h.buffer.Reset()
		h.mutex.Unlock()
	}()
	if err := h.handler.Handle(ctx, r); err != nil {
		return nil, NewHandleError(err, "handler.Handle error when calling Handle func")
	}

	a := make(map[string]any)
	if err := json.Unmarshal(h.buffer.Bytes(), &a); err != nil {
		return nil, NewHandleError(err, "json.Unmarshal error in the Handles result")
	}
	return a, nil
}

// Handle : H a log message
func (h *H) Handle(ctx context.Context, r slog.Record) error {
	var level string
	levelAttr := slog.Attr{
		Key:   slog.LevelKey,
		Value: slog.AnyValue(r.Level),
	}
	// If the level is a custom level, use the custom level name.
	if label, ok := LevelNames[r.Level]; ok {
		levelAttr.Value = slog.StringValue(label)
	}
	if h.record != nil {
		levelAttr = h.record([]string{}, levelAttr)
	}

	if !levelAttr.Equal(slog.Attr{}) {
		level = levelAttr.Value.String() + ":"
		switch {
		case r.Level < slog.LevelInfo:
			level = color(lightGray, level)
		case r.Level < slog.LevelWarn:
			level = color(lightBlue, level)
		case r.Level < slog.LevelError:
			level = color(lightYellow, level)
		case r.Level < LevelFatal:
			level = color(lightRed, level)
		case r.Level >= LevelFatal:
			level = color(fatalColor, level)
		default:
			level = color(lightMagenta, level)
		}
	}

	var timestamp string
	timeAttr := slog.Attr{
		Key:   slog.TimeKey,
		Value: slog.StringValue(r.Time.Format(timeFormat)),
	}
	if h.record != nil {
		timeAttr = h.record([]string{}, timeAttr)
	}
	if !timeAttr.Equal(slog.Attr{}) {
		timestamp = color(lightGray, timeAttr.Value.String())
	}

	var msg string
	msgAttr := slog.Attr{
		Key:   slog.MessageKey,
		Value: slog.StringValue(r.Message),
	}
	if h.record != nil {
		msgAttr = h.record([]string{}, msgAttr)
	}
	if !msgAttr.Equal(slog.Attr{}) {
		msg = color(white, msgAttr.Value.String())
	}

	attrs, err := h.attrs(ctx, r)
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(attrs, "", "  ")
	if err != nil {
		return NewHandleError(err, "json.MarshalIndent error when marshaling attrs")
	}

	out := strings.Builder{}
	if len(timestamp) > 0 {
		out.WriteString(timestamp)
		out.WriteString(" ")
	}
	if len(level) > 0 {
		out.WriteString(level)
		out.WriteString(" ")
	}
	if len(msg) > 0 {
		out.WriteString(msg)
		out.WriteString(" ")
	}
	if len(b) > 0 {
		out.WriteString(color(darkGray, string(b)))
	}
	fmt.Println(out.String())

	return nil
}

// replAttrWrap : wrap the slog.HandlerOptions.ReplaceAttr func
func replAttrWrap(next func([]string, slog.Attr) slog.Attr) func([]string, slog.Attr) slog.Attr {
	return func(groups []string, a slog.Attr) slog.Attr {
		// Remove time from the output for predictable test output.
		if shared.IContainsAny([]string{slog.TimeKey, slog.LevelKey, slog.MessageKey}, a.Key) {
			return slog.Attr{}
		}

		if next == nil {
			return a
		}
		return next(groups, a)
	}
}
