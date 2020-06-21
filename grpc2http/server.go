package grpc2http

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type Server struct {
	server   *http.Server
	mux      *runtime.ServeMux
	regFunc  RegistrationFunc
	grpcAddr string
}

type RegistrationFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error

// NewServer would create a new HTTP server to be hosted on an address (usually on broadcast IP 0.0.0.
// 0 and a port) which would map on grpc service usually running on the same machine as given by grpcAddr
// START OMIT
func NewServer(restAddr string, grpcAddr string, regFunc RegistrationFunc) *Server {
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: false, EmitDefaults: false}),
	)
	return &Server{
		&http.Server{Addr: restAddr, Handler: mux},
		mux,
		regFunc,
		grpcAddr,
	}
}

// Start would spin up the HTTP server component
func (h *Server) Start() (fErr error) {
	if rErr := h.regFunc(context.Background(), h.mux, h.grpcAddr, []grpc.DialOption{grpc.WithInsecure()}); rErr != nil {
		return errors.Wrap(rErr, "unable to spawn HTTP endpoints")
	}
	go func() {
		if err := h.server.ListenAndServe(); err != http.ErrServerClosed {
			fErr = err
			return
		}
	}()
	// END OMIT

	return nil
}

// Stop would spin down the HTTP server component
func (h *Server) Stop() error {
	return h.server.Close()
}
