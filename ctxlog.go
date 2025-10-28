package ctxlog

import (
	"context"
	"fmt"
	"log"
	"os"
)

type logger interface {
	Print(args ...any)
	Printf(f string, args ...any)
	Fatal(args ...any)
	Fatalf(f string, args ...any)
	Panic(args ...any)
	Panicf(f string, args ...any)
}

var forward logger = log.Default()

func Forward(logger logger) {
	forward = logger
}

// ctxlog extends the standard log package by adding an alternative concept of log prefix:
// Instead of tying the prefix to a Logger which would have to be handed around,
// the prefix is tied to a Context and hence can be transparently propagated.

type ctxKeyType string

const ctxKey = ctxKeyType("ctxlog")

// private: return s prefixed with concatenated prefix(es) from context (if any)
func prefixed(ctx context.Context, s string) string {
	if prefix, found := ctx.Value(ctxKey).(*string); found {
		return *prefix + " " + s
	}
	return s
}

// return new Context
func Clone(ctx context.Context) context.Context {
	if previous, found := ctx.Value(ctxKey).(*string); found {
		return Set(ctx, *previous)
	} else {
		return ctx
	}
}

// return new Context with given log prefix set
func Set(ctx context.Context, prefix string) context.Context {
	newPrefix := prefix
	return context.WithValue(ctx, ctxKey, &newPrefix)
}
func Setf(ctx context.Context, format string, args ...any) context.Context {
	return Set(ctx, fmt.Sprintf(format, args...))
}

// return new Context with given log prefix added
func Add(ctx context.Context, prefix string) context.Context {
	if previous, found := ctx.Value(ctxKey).(*string); found {
		*previous = *previous + " " + prefix
		return ctx
	} else {
		return Set(ctx, prefix)
	}
}

// like Add but formatted
func Addf(ctx context.Context, format string, args ...any) context.Context {
	return Add(ctx, fmt.Sprintf(format, args...))
}

// like log.Print() but eventually prefixed with context value (if any)
func Print(ctx context.Context, args ...any) {
	forward.Print(prefixed(ctx, fmt.Sprint(args...)))
}

// like log.Printf() but eventually prefixed with context value (if any)
func Printf(ctx context.Context, format string, args ...any) {
	forward.Print(prefixed(ctx, fmt.Sprintf(format, args...)))
}

// like log.Fatal() but eventually prefixed with context value (if any)
func Fatal(ctx context.Context, args ...any) {
	Print(ctx, args...)
	os.Exit(1)
}

// like log.Fatalf() but eventually prefixed with context value (if any)
func Fatalf(ctx context.Context, format string, args ...any) {
	Printf(ctx, format, args...)
	os.Exit(1)
}

// like log.Panic() but eventually prefixed with context value (if any)
func Panic(ctx context.Context, args ...any) {
	Print(ctx, args...)
	panic(prefixed(ctx, fmt.Sprint(args...)))
}

// like log.Panicf() but eventually prefixed with context value (if any)
func Panicf(ctx context.Context, format string, args ...any) {
	Printf(ctx, format, args...)
	panic(prefixed(ctx, fmt.Sprintf(format, args...)))
}
