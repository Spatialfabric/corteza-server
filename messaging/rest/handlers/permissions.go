package handlers

/*
	Hello! This file is auto-generated from `docs/src/spec.json`.

	For development:
	In order to update the generated files, edit this file under the location,
	add your struct fields, imports, API definitions and whatever you want, and:

	1. run [spec](https://github.com/titpetric/spec) in the same folder,
	2. run `./_gen.php` in this folder.

	You may edit `permissions.go`, `permissions.util.go` or `permissions_test.go` to
	implement your API calls, helper functions and tests. The file `permissions.go`
	is only generated the first time, and will not be overwritten if it exists.
*/

import (
	"context"
	"github.com/go-chi/chi"
	"net/http"

	"github.com/titpetric/factory/resputil"

	"github.com/crusttech/crust/messaging/rest/request"
)

// Internal API interface
type PermissionsAPI interface {
	Effective(context.Context, *request.PermissionsEffective) (interface{}, error)
}

// HTTP API interface
type Permissions struct {
	Effective func(http.ResponseWriter, *http.Request)
}

func NewPermissions(ph PermissionsAPI) *Permissions {
	return &Permissions{
		Effective: func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			params := request.NewPermissionsEffective()
			resputil.JSON(w, params.Fill(r), func() (interface{}, error) {
				return ph.Effective(r.Context(), params)
			})
		},
	}
}

func (ph *Permissions) MountRoutes(r chi.Router, middlewares ...func(http.Handler) http.Handler) {
	r.Group(func(r chi.Router) {
		r.Use(middlewares...)
		r.Route("/permissions", func(r chi.Router) {
			r.Get("/effective", ph.Effective)
		})
	})
}
