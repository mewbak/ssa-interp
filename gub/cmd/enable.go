// Copyright 2013 Rocky Bernstein.

package gubcmd
import (
	"fmt"
	"github.com/rocky/ssa-interp/gub"
)

func init() {
	name := "enable"
	gub.Cmds[name] = &gub.CmdInfo{
		Fn: EnableCommand,
		Help: `enable [bpnum1 ...]

Enable a breakpoint by the number assigned to it.`,

		Min_args: 0,
		Max_args: -1,
	}
	gub.AddToCategory("breakpoints", name)
}

// FIXME: DRY with Disable and Delete?

// EnableCommand implements the debugger command:
//    enable [bpnum1 ...]
// which enables a breakpoint by its breakpoint number.
//
// See also "disable", "delete", and "info break".
func EnableCommand(args []string) {
	for i:=1; i<len(args); i++ {
		msg := fmt.Sprintf("breakpoint number for argument %d", i)
		val, err := gub.GetUInt(args[i], msg, 0, uint64(len(gub.Breakpoints)-1))
		if err != nil { continue }
		bpnum := gub.BpId(val)
		if gub.BreakpointExists(bpnum) {
			if gub.BreakpointIsEnabled(bpnum) {
				gub.Msg("Breakpoint %d is already enabled", bpnum)
				continue
			}
			if gub.BreakpointEnable(bpnum) {
				gub.Msg("Breakpoint %d enabled", bpnum)
			} else {
				gub.Errmsg("Trouble enabling breakpoint %d", bpnum)
			}
		} else {
			gub.Errmsg("Breakpoint %d doesn't exist", bpnum)
		}
	}
}
