package automation

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// system/automation/roles_handler.yaml

import (
	"context"
	atypes "github.com/cortezaproject/corteza-server/automation/types"
	"github.com/cortezaproject/corteza-server/pkg/expr"
	"github.com/cortezaproject/corteza-server/pkg/wfexec"
	"github.com/cortezaproject/corteza-server/system/types"
)

var _ wfexec.ExecResponse

type (
	rolesHandlerRegistry interface {
		AddFunctions(ff ...*atypes.Function)
		Type(ref string) expr.Type
	}
)

func (h rolesHandler) register() {
	h.reg.AddFunctions(
		h.Lookup(),
		h.Create(),
		h.Update(),
		h.Delete(),
		h.Recover(),
		h.Archive(),
		h.Unarchive(),
	)
}

type (
	rolesLookupArgs struct {
		hasQuery    bool
		Query       interface{}
		queryID     uint64
		queryHandle string
	}

	rolesLookupResults struct {
		Role *types.Role
	}
)

// Lookup function Looks-up for compose role by ID
//
// expects implementation of lookup function:
// func (h rolesHandler) lookup(ctx context.Context, args *rolesLookupArgs) (results *rolesLookupResults, err error) {
//    return
// }
func (h rolesHandler) Lookup() *atypes.Function {
	return &atypes.Function{
		Ref:  "rolesLookup",
		Kind: "function",
		Meta: &atypes.FunctionMeta{
			Short: "Looks-up for compose role by ID",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "query",
				Types: []string{"ID", "Handle"}, Required: true,
			},
		},

		Results: []*atypes.Param{

			{
				Name:  "role",
				Types: []string{"Role"},
			},
		},

		Handler: func(ctx context.Context, in *expr.Vars) (out *expr.Vars, err error) {
			var (
				args = &rolesLookupArgs{
					hasQuery: in.Has("query"),
				}
			)

			if err = in.Decode(args); err != nil {
				return
			}

			// Converting Query argument
			if args.hasQuery {
				aux := expr.Must(expr.Select(in, "query"))
				switch aux.Type() {
				case h.reg.Type("ID").Type():
					args.queryID = aux.Get().(uint64)
				case h.reg.Type("Handle").Type():
					args.queryHandle = aux.Get().(string)
				}
			}

			var results *rolesLookupResults
			if results, err = h.lookup(ctx, args); err != nil {
				return
			}

			out = &expr.Vars{}
			_ = expr.Set(out, "role", expr.Must(h.reg.Type("Role").Cast(results.Role)))

			return
		},
	}
}

type (
	rolesCreateArgs struct {
		hasRole bool
		Role    *types.Role
	}

	rolesCreateResults struct {
		Role *types.Role
	}
)

// Create function Creates new role
//
// expects implementation of create function:
// func (h rolesHandler) create(ctx context.Context, args *rolesCreateArgs) (results *rolesCreateResults, err error) {
//    return
// }
func (h rolesHandler) Create() *atypes.Function {
	return &atypes.Function{
		Ref:  "rolesCreate",
		Kind: "function",
		Meta: &atypes.FunctionMeta{
			Short: "Creates new role",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "role",
				Types: []string{"Role"}, Required: true,
			},
		},

		Results: []*atypes.Param{

			{
				Name:  "role",
				Types: []string{"Role"},
			},
		},

		Handler: func(ctx context.Context, in *expr.Vars) (out *expr.Vars, err error) {
			var (
				args = &rolesCreateArgs{
					hasRole: in.Has("role"),
				}
			)

			if err = in.Decode(args); err != nil {
				return
			}

			var results *rolesCreateResults
			if results, err = h.create(ctx, args); err != nil {
				return
			}

			out = &expr.Vars{}
			_ = expr.Set(out, "role", expr.Must(h.reg.Type("Role").Cast(results.Role)))

			return
		},
	}
}

type (
	rolesUpdateArgs struct {
		hasRole bool
		Role    *types.Role
	}

	rolesUpdateResults struct {
		Role *types.Role
	}
)

// Update function Updates exiting role
//
// expects implementation of update function:
// func (h rolesHandler) update(ctx context.Context, args *rolesUpdateArgs) (results *rolesUpdateResults, err error) {
//    return
// }
func (h rolesHandler) Update() *atypes.Function {
	return &atypes.Function{
		Ref:  "rolesUpdate",
		Kind: "function",
		Meta: &atypes.FunctionMeta{
			Short: "Updates exiting role",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "role",
				Types: []string{"Role"}, Required: true,
			},
		},

		Results: []*atypes.Param{

			{
				Name:  "role",
				Types: []string{"Role"},
			},
		},

		Handler: func(ctx context.Context, in *expr.Vars) (out *expr.Vars, err error) {
			var (
				args = &rolesUpdateArgs{
					hasRole: in.Has("role"),
				}
			)

			if err = in.Decode(args); err != nil {
				return
			}

			var results *rolesUpdateResults
			if results, err = h.update(ctx, args); err != nil {
				return
			}

			out = &expr.Vars{}
			_ = expr.Set(out, "role", expr.Must(h.reg.Type("Role").Cast(results.Role)))

			return
		},
	}
}

type (
	rolesDeleteArgs struct {
		hasQuery    bool
		Query       interface{}
		queryID     uint64
		queryHandle string

		hasRole bool
		Role    *types.Role
	}
)

// Delete function Deletes the role
//
// expects implementation of delete function:
// func (h rolesHandler) delete(ctx context.Context, args *rolesDeleteArgs) (err error) {
//    return
// }
func (h rolesHandler) Delete() *atypes.Function {
	return &atypes.Function{
		Ref:  "rolesDelete",
		Kind: "function",
		Meta: &atypes.FunctionMeta{
			Short: "Deletes the role",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "query",
				Types: []string{"ID", "Handle"},
			},
			{
				Name:  "role",
				Types: []string{"Role"},
			},
		},

		Handler: func(ctx context.Context, in *expr.Vars) (out *expr.Vars, err error) {
			var (
				args = &rolesDeleteArgs{
					hasQuery: in.Has("query"),
					hasRole:  in.Has("role"),
				}
			)

			if err = in.Decode(args); err != nil {
				return
			}

			// Converting Query argument
			if args.hasQuery {
				aux := expr.Must(expr.Select(in, "query"))
				switch aux.Type() {
				case h.reg.Type("ID").Type():
					args.queryID = aux.Get().(uint64)
				case h.reg.Type("Handle").Type():
					args.queryHandle = aux.Get().(string)
				}
			}

			return out, h.delete(ctx, args)
		},
	}
}

type (
	rolesRecoverArgs struct {
		hasQuery    bool
		Query       interface{}
		queryID     uint64
		queryHandle string

		hasRole bool
		Role    *types.Role
	}
)

// Recover function Recovers deleted role
//
// expects implementation of recover function:
// func (h rolesHandler) recover(ctx context.Context, args *rolesRecoverArgs) (err error) {
//    return
// }
func (h rolesHandler) Recover() *atypes.Function {
	return &atypes.Function{
		Ref:  "rolesRecover",
		Kind: "function",
		Meta: &atypes.FunctionMeta{
			Short: "Recovers deleted role",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "query",
				Types: []string{"ID", "Handle"},
			},
			{
				Name:  "role",
				Types: []string{"Role"},
			},
		},

		Handler: func(ctx context.Context, in *expr.Vars) (out *expr.Vars, err error) {
			var (
				args = &rolesRecoverArgs{
					hasQuery: in.Has("query"),
					hasRole:  in.Has("role"),
				}
			)

			if err = in.Decode(args); err != nil {
				return
			}

			// Converting Query argument
			if args.hasQuery {
				aux := expr.Must(expr.Select(in, "query"))
				switch aux.Type() {
				case h.reg.Type("ID").Type():
					args.queryID = aux.Get().(uint64)
				case h.reg.Type("Handle").Type():
					args.queryHandle = aux.Get().(string)
				}
			}

			return out, h.recover(ctx, args)
		},
	}
}

type (
	rolesArchiveArgs struct {
		hasQuery    bool
		Query       interface{}
		queryID     uint64
		queryHandle string

		hasRole bool
		Role    *types.Role
	}
)

// Archive function Archives the role
//
// expects implementation of archive function:
// func (h rolesHandler) archive(ctx context.Context, args *rolesArchiveArgs) (err error) {
//    return
// }
func (h rolesHandler) Archive() *atypes.Function {
	return &atypes.Function{
		Ref:  "rolesArchive",
		Kind: "function",
		Meta: &atypes.FunctionMeta{
			Short: "Archives the role",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "query",
				Types: []string{"ID", "Handle"},
			},
			{
				Name:  "role",
				Types: []string{"Role"},
			},
		},

		Handler: func(ctx context.Context, in *expr.Vars) (out *expr.Vars, err error) {
			var (
				args = &rolesArchiveArgs{
					hasQuery: in.Has("query"),
					hasRole:  in.Has("role"),
				}
			)

			if err = in.Decode(args); err != nil {
				return
			}

			// Converting Query argument
			if args.hasQuery {
				aux := expr.Must(expr.Select(in, "query"))
				switch aux.Type() {
				case h.reg.Type("ID").Type():
					args.queryID = aux.Get().(uint64)
				case h.reg.Type("Handle").Type():
					args.queryHandle = aux.Get().(string)
				}
			}

			return out, h.archive(ctx, args)
		},
	}
}

type (
	rolesUnarchiveArgs struct {
		hasQuery    bool
		Query       interface{}
		queryID     uint64
		queryHandle string

		hasRole bool
		Role    *types.Role
	}
)

// Unarchive function Unarchives the role
//
// expects implementation of unarchive function:
// func (h rolesHandler) unarchive(ctx context.Context, args *rolesUnarchiveArgs) (err error) {
//    return
// }
func (h rolesHandler) Unarchive() *atypes.Function {
	return &atypes.Function{
		Ref:  "rolesUnarchive",
		Kind: "function",
		Meta: &atypes.FunctionMeta{
			Short: "Unarchives the role",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "query",
				Types: []string{"ID", "Handle"},
			},
			{
				Name:  "role",
				Types: []string{"Role"},
			},
		},

		Handler: func(ctx context.Context, in *expr.Vars) (out *expr.Vars, err error) {
			var (
				args = &rolesUnarchiveArgs{
					hasQuery: in.Has("query"),
					hasRole:  in.Has("role"),
				}
			)

			if err = in.Decode(args); err != nil {
				return
			}

			// Converting Query argument
			if args.hasQuery {
				aux := expr.Must(expr.Select(in, "query"))
				switch aux.Type() {
				case h.reg.Type("ID").Type():
					args.queryID = aux.Get().(uint64)
				case h.reg.Type("Handle").Type():
					args.queryHandle = aux.Get().(string)
				}
			}

			return out, h.unarchive(ctx, args)
		},
	}
}
