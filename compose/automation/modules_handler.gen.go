package automation

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// compose/automation/modules_handler.yaml

import (
	"context"
	atypes "github.com/cortezaproject/corteza-server/automation/types"
	"github.com/cortezaproject/corteza-server/compose/types"
	"github.com/cortezaproject/corteza-server/pkg/expr"
	"github.com/cortezaproject/corteza-server/pkg/wfexec"
)

var _ wfexec.ExecResponse

type (
	modulesHandlerRegistry interface {
		AddFunctions(ff ...*atypes.Function)
		Type(ref string) expr.Type
	}
)

func (h modulesHandler) register() {
	h.reg.AddFunctions(
		h.Lookup(),
	)
}

type (
	modulesLookupArgs struct {
		hasModule    bool
		Module       interface{}
		moduleID     uint64
		moduleHandle string

		hasNamespace    bool
		Namespace       interface{}
		namespaceID     uint64
		namespaceHandle string
		namespaceRes    *types.Namespace
	}

	modulesLookupResults struct {
		Module *types.Module
	}
)

// Lookup function Lookup for compose module by ID
//
// expects implementation of lookup function:
// func (h modulesHandler) lookup(ctx context.Context, args *modulesLookupArgs) (results *modulesLookupResults, err error) {
//    return
// }
func (h modulesHandler) Lookup() *atypes.Function {
	return &atypes.Function{
		Ref:  "composeModulesLookup",
		Kind: "function",
		Meta: &atypes.FunctionMeta{
			Short: "Lookup for compose module by ID",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "module",
				Types: []string{"ID", "Handle"}, Required: true,
			},
			{
				Name:  "namespace",
				Types: []string{"ID", "Handle", "ComposeNamespace"}, Required: true,
			},
		},

		Results: []*atypes.Param{

			{
				Name:  "module",
				Types: []string{"ComposeModule"},
			},
		},

		Handler: func(ctx context.Context, in *expr.Vars) (out *expr.Vars, err error) {
			var (
				args = &modulesLookupArgs{
					hasModule:    in.Has("module"),
					hasNamespace: in.Has("namespace"),
				}
			)

			if err = in.Decode(args); err != nil {
				return
			}

			// Converting Module argument
			if args.hasModule {
				aux := expr.Must(expr.Select(in, "query"))
				switch aux.Type() {
				case h.reg.Type("ID").Type():
					args.moduleID = aux.Get().(uint64)
				case h.reg.Type("Handle").Type():
					args.moduleHandle = aux.Get().(string)
				}
			}

			// Converting Namespace argument
			if args.hasNamespace {
				aux := expr.Must(expr.Select(in, "query"))
				switch aux.Type() {
				case h.reg.Type("ID").Type():
					args.namespaceID = aux.Get().(uint64)
				case h.reg.Type("Handle").Type():
					args.namespaceHandle = aux.Get().(string)
				case h.reg.Type("ComposeNamespace").Type():
					args.namespaceRes = aux.Get().(*types.Namespace)
				}
			}

			var results *modulesLookupResults
			if results, err = h.lookup(ctx, args); err != nil {
				return
			}

			out = &expr.Vars{}
			_ = expr.Set(out, "module", expr.Must(h.reg.Type("ComposeModule").Cast(results.Module)))

			return
		},
	}
}
