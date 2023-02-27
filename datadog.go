package datadog_caddy

import (
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"

	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func init() {
	caddy.RegisterModule(Datadog{})
}

// Datadog is an example; put your own type here.
type Datadog struct {
}

func (d Datadog) Provision(ctx caddy.Context) error {
	ctx.Logger().Info("zouzou")
	tracer.Start(tracer.WithDebugMode(true))
	return nil
}

func (d Datadog) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) (err error) {
	httptrace.TraceAndServe(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err = next.ServeHTTP(w, r)
	}), w, r, nil)
	return err
}

var (
	_ caddyhttp.MiddlewareHandler = (*Datadog)(nil)
)

// CaddyModule returns the Caddy module information.
func (Datadog) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.datadog",
		New: func() caddy.Module { return new(Datadog) },
	}
}
