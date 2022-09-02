package uda

import (
	"context"
)

type executor interface {
	add(_ context.Context, domain string, rr string, value string, options *options) (err error)

	resolve(ctx context.Context, domain string, rr string, value string, options *options) (result *Result, err error)

	get(ctx context.Context, domain string, rr string, options *options) (record *Record, err error)

	update(_ context.Context, record *Record, rr string, value string, options *options) (err error)
}
