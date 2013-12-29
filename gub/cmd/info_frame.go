// Copyright 2013 Rocky Bernstein.

package gubcmd

import (
	"github.com/rocky/ssa-interp/gub"
)

func init() {
	parent := "info"
	gub.AddSubCommand(parent, &gub.SubcmdInfo{
		Fn: InfoProgramSubcmd,
		Help: `info frame

Prints information about the selected frame including:
*  goroutine number
*  location
*  function and parameter names

See also backtrace.
`,
		Min_args: 0,
		Max_args: 0,
		Short_help: "Show information about the selected frame",
		Name: "frame",
	})
}

// InfoFrameSubcmd implements the debugger command:
//   info frame
// which prints frame information including:
//    goroutine number
//    location
//    function and parameter names
func InfoFrameSubcmd(args []string) {
	fr := gub.CurFrame()
	gub.Msg("goroutine number: %d", fr.GoNum())
	gub.Msg("location: %s", fr.PositionRange())
	gub.Msg("frame: %s", fr.FnAndParamString())
}
