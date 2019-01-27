package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// MuxRouterAdapter will make sure mux.Router is compatible with RouteHandler interface
type MuxRouterAdapter struct {
	*mux.Router
}

// Handle implements RouteHandler.Handle()
func (muxRouterAdapter *MuxRouterAdapter) Handle(path string, handler http.Handler) {
	muxRouterAdapter.Router.Handle(path, handler)
}

// HandleFunc implements RouteHandler.HandleFunc()
func (muxRouterAdapter *MuxRouterAdapter) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) {
	muxRouterAdapter.Router.HandleFunc(path, f)
}
