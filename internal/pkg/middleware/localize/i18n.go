package localize

import (
	"context"

	"github.com/go-kratos/kratos-layout/internal/pkg/i18n"

	local "github.com/nicksnyder/go-i18n/v2/i18n"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

type localizerKey struct{}

func I18N() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				// Do something on entering
				defer func() {
					// Do something on exiting
				}()

				header := tr.RequestHeader()
				lang := header.Get("language")
				bundle, err := i18n.NewBundle()
				if err != nil {
					return nil, err
				}
				localizer := i18n.NewLocalizer(bundle, lang)
				ctx = context.WithValue(ctx, localizerKey{}, localizer)
			}
			return handler(ctx, req)
		}
	}
}

func FromContext(ctx context.Context) *local.Localizer {
	return ctx.Value(localizerKey{}).(*local.Localizer)
}

func Localize(ctx context.Context, ID string) (string, error) {
	localizer := FromContext(ctx)
	return localizer.Localize(&local.LocalizeConfig{MessageID: ID})
}
