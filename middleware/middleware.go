package middleware

import (
	"context"
	"github.com/kubemq-hub/kubemq-targets/types"
)

type Middleware interface {
	Do(ctx context.Context, request *types.Request) (*types.Response, error)
}

type DoFunc func(ctx context.Context, request *types.Request) (*types.Response, error)

func (df DoFunc) Do(ctx context.Context, request *types.Request) (*types.Response, error) {
	return df(ctx, request)
}

type MiddlewareFunc func(Middleware) Middleware

func Log(log *LogMiddleware) MiddlewareFunc {
	return func(df Middleware) Middleware {
		return DoFunc(func(ctx context.Context, request *types.Request) (*types.Response, error) {
			result, err := df.Do(ctx, request)
			switch log.minLevel {
			case "debug":
				reqStr := ""
				if request != nil {
					reqStr = request.String()
				}
				resStr := ""
				if result != nil {
					resStr = result.String()
				}
				log.Infof("request: %s, response: %s, error:%+v", reqStr, resStr, err)
			case "info":
				if err != nil {
					log.Errorf("error processing request: %s", err.Error())
				} else {
					log.Infof("request processed with successful response")
				}
			case "error":
				if err != nil {
					log.Errorf("error processing request: %s", err.Error())
				}
			}
			return result, err
		})
	}
}

func Chain(md Middleware, list ...MiddlewareFunc) Middleware {
	chain := md
	for _, middleware := range list {
		chain = middleware(chain)
	}
	return chain
}
