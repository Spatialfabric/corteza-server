package expr

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestVars(t *testing.T) {
	var (
		req = require.New(t)

		vars = RVars{
			"int": Must(NewInteger(42)),
			"sub": RVars{
				"foo": Must(NewString("foo")),
			}.Vars(),
			"three": RVars{
				"two": RVars{
					"one": RVars{
						"go": Must(NewString("!")),
					}.Vars(),
				}.Vars(),
			}.Vars(),
		}.Vars()
	)

	//req.NoError(Set(vars, "int", 123))
	//req.Equal(int64(123), vars["int"].(TypedValue).Get().(int64))
	//
	//req.NoError(Set(vars, "sub.foo", "bar"))
	//req.Equal("bar", (*(vars["sub"]).(*Vars))["foo"].Get().(string))

	req.NoError(Set(vars, "kv", &KV{}))
	req.NoError(Set(vars, "kv.foo", "bar"))
	req.Equal("bar", Must(Select(vars, "kv.foo")).Get().(string))

	//req.NoError(Set(vars, "three.two.one.go", "!!!"))
	//req.Equal("!!!", Must(Select(vars, "three.two.one.go")).Get().(string))
}
