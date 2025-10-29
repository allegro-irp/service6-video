package checkapi

import (
	"github.com/allegro-irp/service6-video/foundation/web"
)

// Routes adds specific routes for this group.
func Routes(app *web.App) {
	app.HandlerFunc("GET /liveness", liveness)
	app.HandlerFunc("GET /readiness", readiness)
}
