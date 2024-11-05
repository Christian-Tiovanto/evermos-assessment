package http

import (
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"github.com/mughieams/evermos-assessment/app/common/config"
	"github.com/mughieams/evermos-assessment/app/common/dependencies/postgresql"
	"github.com/mughieams/evermos-assessment/app/delivery/http/handler"
)

// RegisterRestHandler register REST handlers
func RegisterRestHandler(mux *runtime.ServeMux, _ config.Config, _ postgresql.PgxPoolIface) error {
	if err := mux.HandlePath(http.MethodGet, "/hello", handler.HelloHandler); err != nil {
		return err
	}
	return nil
}
