// Copyright 2013 Rocky Bernstein.
package gubcmd

import (
	"github.com/rocky/ssa-interp/gub"
	"github.com/rocky/ssa-interp/interp"
)

func init() {
	name := "eval"
	gub.Cmds[name] = &gub.CmdInfo{
		Fn: EvalCommand,
		Help: `eval *expr*

Evaluate go expression *expr*.
`,

		Min_args: 1,
		Max_args: -1,
	}
	gub.AddToCategory("data", name)
}

func EvalCommand(args []string) {

	// Don't use args, but gub.CmdArgstr which preserves blanks inside quotes
	if expr, err := gub.EvalExpr(gub.CmdArgstr); err == nil {
		if expr == nil {
			gub.Section("Kind=nil")
			gub.Msg("nil")
		} else {
			expr := *expr
			if len(expr) == 1 {
				value := expr[0]
				kind := value.Kind().String()
				typ  := value.Type().String()
				if typ != kind {
					gub.Section("Kind = %v", kind)
					gub.Section("Type = %v", typ)
				} else {
					gub.Section("Kind = Type = %v", kind)
				}
				gub.Msg("%s", interp.ToInspect(value.Interface()))
			} else {
				if len(expr) == 0 {
					gub.Section("Kind=Slice")
					gub.Msg("void")
				} else {
					gub.Section("Kind = Multi-Value")
					size := len(expr)
					for i, v := range expr {
						if v.Interface() == nil {
							gub.MsgNoCr("nil")
						} else {
							gub.MsgNoCr("%v", v.Interface())
						}
						if i < size-1 { gub.MsgNoCr(", ") }
					}
					gub.Msg("")
				}
			}
		}
	}
}
