package automation

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// system/automation/users_handler.yaml

import (
	"context"
	atypes "github.com/cortezaproject/corteza-server/automation/types"
	"github.com/cortezaproject/corteza-server/pkg/expr"
	"github.com/cortezaproject/corteza-server/pkg/wfexec"
	"github.com/cortezaproject/corteza-server/system/types"
)

var _ wfexec.ExecResponse

type (
	usersHandlerRegistry interface {
		AddFunctions(ff ...*atypes.Function)
		Type(ref string) expr.Type
	}
)

func (h usersHandler) register() {
	h.reg.AddFunctions(
		h.Lookup(),
		h.Create(),
		h.Update(),
		h.Delete(),
		h.Recover(),
		h.Suspend(),
		h.Unsuspend(),
	)
}

type (
	usersLookupArgs struct {
		hasQuery    bool
		Query       interface{}
		queryID     uint64
		queryHandle string
		queryEmail  string
	}

	usersLookupResults struct {
		User *types.User
	}
)

// Lookup function Looks-up for compose user by ID
//
// expects implementation of lookup function:
// func (h usersHandler) lookup(ctx context.Context, args *usersLookupArgs) (results *usersLookupResults, err error) {
//    return
// }
func (h usersHandler) Lookup() *atypes.Function {
	return &atypes.Function{
		Ref:  "usersLookup",
		Kind: "function",
		Meta: &atypes.FunctionMeta{
			Short: "Looks-up for compose user by ID",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "query",
				Types: []string{"ID", "Handle", "String"}, Required: true,
			},
		},

		Results: []*atypes.Param{

			{
				Name:  "user",
				Types: []string{"User"},
			},
		},

		Handler: func(ctx context.Context, in *expr.Vars) (out *expr.Vars, err error) {
			var (
				args = &usersLookupArgs{
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
				case h.reg.Type("String").Type():
					args.queryEmail = aux.Get().(string)
				}
			}

			var results *usersLookupResults
			if results, err = h.lookup(ctx, args); err != nil {
				return
			}

			out = &expr.Vars{}
			_ = expr.Set(out, "user", expr.Must(h.reg.Type("User").Cast(results.User)))

			return
		},
	}
}

type (
	usersCreateArgs struct {
		hasUser bool
		User    *types.User
	}

	usersCreateResults struct {
		User *types.User
	}
)

// Create function Creates new user
//
// expects implementation of create function:
// func (h usersHandler) create(ctx context.Context, args *usersCreateArgs) (results *usersCreateResults, err error) {
//    return
// }
func (h usersHandler) Create() *atypes.Function {
	return &atypes.Function{
		Ref:  "usersCreate",
		Kind: "function",
		Meta: &atypes.FunctionMeta{
			Short: "Creates new user",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "user",
				Types: []string{"User"}, Required: true,
			},
		},

		Results: []*atypes.Param{

			{
				Name:  "user",
				Types: []string{"User"},
			},
		},

		Handler: func(ctx context.Context, in *expr.Vars) (out *expr.Vars, err error) {
			var (
				args = &usersCreateArgs{
					hasUser: in.Has("user"),
				}
			)

			if err = in.Decode(args); err != nil {
				return
			}

			var results *usersCreateResults
			if results, err = h.create(ctx, args); err != nil {
				return
			}

			out = &expr.Vars{}
			_ = expr.Set(out, "user", expr.Must(h.reg.Type("User").Cast(results.User)))

			return
		},
	}
}

type (
	usersUpdateArgs struct {
		hasUser bool
		User    *types.User
	}

	usersUpdateResults struct {
		User *types.User
	}
)

// Update function Updates exiting user
//
// expects implementation of update function:
// func (h usersHandler) update(ctx context.Context, args *usersUpdateArgs) (results *usersUpdateResults, err error) {
//    return
// }
func (h usersHandler) Update() *atypes.Function {
	return &atypes.Function{
		Ref:  "usersUpdate",
		Kind: "function",
		Meta: &atypes.FunctionMeta{
			Short: "Updates exiting user",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "user",
				Types: []string{"User"}, Required: true,
			},
		},

		Results: []*atypes.Param{

			{
				Name:  "user",
				Types: []string{"User"},
			},
		},

		Handler: func(ctx context.Context, in *expr.Vars) (out *expr.Vars, err error) {
			var (
				args = &usersUpdateArgs{
					hasUser: in.Has("user"),
				}
			)

			if err = in.Decode(args); err != nil {
				return
			}

			var results *usersUpdateResults
			if results, err = h.update(ctx, args); err != nil {
				return
			}

			out = &expr.Vars{}
			_ = expr.Set(out, "user", expr.Must(h.reg.Type("User").Cast(results.User)))

			return
		},
	}
}

type (
	usersDeleteArgs struct {
		hasQuery    bool
		Query       interface{}
		queryID     uint64
		queryHandle string
		queryEmail  string

		hasUser bool
		User    *types.User
	}
)

// Delete function Deletes user
//
// expects implementation of delete function:
// func (h usersHandler) delete(ctx context.Context, args *usersDeleteArgs) (err error) {
//    return
// }
func (h usersHandler) Delete() *atypes.Function {
	return &atypes.Function{
		Ref:  "usersDelete",
		Kind: "function",
		Meta: &atypes.FunctionMeta{
			Short: "Deletes user",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "query",
				Types: []string{"ID", "Handle", "String"},
			},
			{
				Name:  "user",
				Types: []string{"User"},
			},
		},

		Handler: func(ctx context.Context, in *expr.Vars) (out *expr.Vars, err error) {
			var (
				args = &usersDeleteArgs{
					hasQuery: in.Has("query"),
					hasUser:  in.Has("user"),
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
				case h.reg.Type("String").Type():
					args.queryEmail = aux.Get().(string)
				}
			}

			return out, h.delete(ctx, args)
		},
	}
}

type (
	usersRecoverArgs struct {
		hasQuery    bool
		Query       interface{}
		queryID     uint64
		queryHandle string
		queryEmail  string

		hasUser bool
		User    *types.User
	}
)

// Recover function Recovers deleted user
//
// expects implementation of recover function:
// func (h usersHandler) recover(ctx context.Context, args *usersRecoverArgs) (err error) {
//    return
// }
func (h usersHandler) Recover() *atypes.Function {
	return &atypes.Function{
		Ref:  "usersRecover",
		Kind: "function",
		Meta: &atypes.FunctionMeta{
			Short: "Recovers deleted user",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "query",
				Types: []string{"ID", "Handle", "String"},
			},
			{
				Name:  "user",
				Types: []string{"User"},
			},
		},

		Handler: func(ctx context.Context, in *expr.Vars) (out *expr.Vars, err error) {
			var (
				args = &usersRecoverArgs{
					hasQuery: in.Has("query"),
					hasUser:  in.Has("user"),
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
				case h.reg.Type("String").Type():
					args.queryEmail = aux.Get().(string)
				}
			}

			return out, h.recover(ctx, args)
		},
	}
}

type (
	usersSuspendArgs struct {
		hasQuery    bool
		Query       interface{}
		queryID     uint64
		queryHandle string
		queryEmail  string

		hasUser bool
		User    *types.User
	}
)

// Suspend function Suspends user
//
// expects implementation of suspend function:
// func (h usersHandler) suspend(ctx context.Context, args *usersSuspendArgs) (err error) {
//    return
// }
func (h usersHandler) Suspend() *atypes.Function {
	return &atypes.Function{
		Ref:  "usersSuspend",
		Kind: "function",
		Meta: &atypes.FunctionMeta{
			Short: "Suspends user",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "query",
				Types: []string{"ID", "Handle", "String"},
			},
			{
				Name:  "user",
				Types: []string{"User"},
			},
		},

		Handler: func(ctx context.Context, in *expr.Vars) (out *expr.Vars, err error) {
			var (
				args = &usersSuspendArgs{
					hasQuery: in.Has("query"),
					hasUser:  in.Has("user"),
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
				case h.reg.Type("String").Type():
					args.queryEmail = aux.Get().(string)
				}
			}

			return out, h.suspend(ctx, args)
		},
	}
}

type (
	usersUnsuspendArgs struct {
		hasQuery    bool
		Query       interface{}
		queryID     uint64
		queryHandle string
		queryEmail  string

		hasUser bool
		User    *types.User
	}
)

// Unsuspend function Unsuspends user
//
// expects implementation of unsuspend function:
// func (h usersHandler) unsuspend(ctx context.Context, args *usersUnsuspendArgs) (err error) {
//    return
// }
func (h usersHandler) Unsuspend() *atypes.Function {
	return &atypes.Function{
		Ref:  "usersUnsuspend",
		Kind: "function",
		Meta: &atypes.FunctionMeta{
			Short: "Unsuspends user",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "query",
				Types: []string{"ID", "Handle", "String"},
			},
			{
				Name:  "user",
				Types: []string{"User"},
			},
		},

		Handler: func(ctx context.Context, in *expr.Vars) (out *expr.Vars, err error) {
			var (
				args = &usersUnsuspendArgs{
					hasQuery: in.Has("query"),
					hasUser:  in.Has("user"),
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
				case h.reg.Type("String").Type():
					args.queryEmail = aux.Get().(string)
				}
			}

			return out, h.unsuspend(ctx, args)
		},
	}
}
