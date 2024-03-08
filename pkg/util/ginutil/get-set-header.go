package ginutil

import (
	"context"
	"errors"
	"net/http"
)

type requestHeader struct{}
type requestMethod struct{}
type requestUri struct{}

// ContextWithRequest: extract header, method, uri to context
func ContextWithRequest(ctx context.Context, request *http.Request) context.Context {
	ctx = context.WithValue(ctx, requestHeader{}, request.Header)
	ctx = context.WithValue(ctx, requestMethod{}, request.Method)
	ctx = context.WithValue(ctx, requestUri{}, request.RequestURI)
	return ctx
}

func GetRequestHeader(ctx context.Context, key string) (string, error) {
	if h, ok := ctx.Value(requestHeader{}).(http.Header); ok {
		return h.Get(key), nil
	}
	return "", errors.New("unable to get request header from context")
}

func GetRequestMethod(ctx context.Context) (string, error) {
	if m, ok := ctx.Value(requestMethod{}).(string); ok {
		return m, nil
	}
	return "", errors.New("unable to get request method from context")
}

func GetRequestUri(ctx context.Context) (string, error) {
	if uri, ok := ctx.Value(requestUri{}).(string); ok {
		return uri, nil
	}
	return "", errors.New("unable to get request request uri from context")
}
