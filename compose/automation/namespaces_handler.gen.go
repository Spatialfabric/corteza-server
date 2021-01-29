package automation

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// compose/automation/namespaces_handler.yaml

import (
	"context"
	atypes "github.com/cortezaproject/corteza-server/automation/types"
	"github.com/cortezaproject/corteza-server/compose/types"
	"github.com/cortezaproject/corteza-server/pkg/expr"
	"github.com/cortezaproject/corteza-server/pkg/wfexec"
)

var _ wfexec.ExecResponse

type (
	namespacesHandlerRegistry interface {
		AddFunctions(ff ...*atypes.Function)
		Type(ref string) expr.Type
	}
)

func (h namespacesHandler) register() {
	h.reg.AddFunctions(
		h.Lookup(),
	)
}

type (
	namespacesLookupArgs struct {
		hasNamespace    bool
		Namespace       interface{}
		namespaceID     uint64
		namespaceHandle string
	}

	namespacesLookupResults struct {
		Namespace *types.Namespace
	}
)

// Lookup function Lookup for compose namespace by ID
//
// expects implementation of lookup function:
// func (h namespacesHandler) lookup(ctx context.Context, args *namespacesLookupArgs) (results *namespacesLookupResults, err error) {
//    return
// }
func (h namespacesHandler) Lookup() *atypes.Function {
	return &atypes.Function{
		Ref:  "composeNamespacesLookup",
		Kind: "function",
		Meta: &atypes.FunctionMeta{
			Short: "Lookup for compose namespace by ID",
		},

		Parameters: []*atypes.Param{
			{
				Name:  "namespace",
				Types: []string{"ID", "Handle"}, Required: true,
			},
		},

		Results: []*atypes.Param{

			{
				Name:  "namespace",
				Types: []string{"ComposeNamespace"},
			},
		},

		Handler: func(ctx context.Context, in *expr.Vars) (out *expr.Vars, err error) {
			var (
				args = &namespacesLookupArgs{
					hasNamespace: in.Has("namespace"),
				}
			)

			if err = in.Decode(args); err != nil {
				return
			}

			// Converting Namespace argument
			if args.hasNamespace {
				aux := expr.Must(expr.Select(in, "query"))
				switch aux.Type() {
				case h.reg.Type("ID").Type():
					args.namespaceID = aux.Get().(uint64)
				case h.reg.Type("Handle").Type():
					args.namespaceHandle = aux.Get().(string)
				}
			}

			var results *namespacesLookupResults
			if results, err = h.lookup(ctx, args); err != nil {
				return
			}

			out = &expr.Vars{}
			_ = expr.Set(out, "namespace", expr.Must(h.reg.Type("ComposeNamespace").Cast(results.Namespace)))

			return
		},
	}
}
