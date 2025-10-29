// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"os"

	"github.com/allegro-irp/service6-video/apis/services/sales/route/sys/checkapi"
	"github.com/allegro-irp/service6-video/foundation/web"
)

func WebAPI(shutdown chan os.Signal) *web.App {
	mux := web.NewApp(shutdown)

	checkapi.Routes(mux)

	return mux
}
