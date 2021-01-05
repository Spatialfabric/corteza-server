package rest

import (
	"github.com/go-chi/chi"

	"github.com/cortezaproject/corteza-server/automation/rest/handlers"
	"github.com/cortezaproject/corteza-server/pkg/auth"
)

func MountRoutes(r chi.Router) {
	// Protect all _private_ routes
	r.Group(func(r chi.Router) {
		r.Use(auth.MiddlewareValidOnly)

		handlers.NewWorkflow(Workflow{}.New()).MountRoutes(r)
		handlers.NewTrigger(Trigger{}.New()).MountRoutes(r)
		handlers.NewFunction(Function{}.New()).MountRoutes(r)
	})
}