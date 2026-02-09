package portsoutlog

import "context"

type Log interface {
	Error(ctx context.Context, tag string, action string, message string, metadata map[string]any)
	Warn(ctx context.Context, tag string, action string, message string, metadata map[string]any)
	Info(ctx context.Context, tag string, action string, message string, metadata map[string]any)
	Debug(ctx context.Context, tag string, action string, message string, metadata map[string]any)
}
