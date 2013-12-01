// Copyright 2013 Rocky Bernstein.

// info program
//
// Prints program information

package gubcmd

import (
	"github.com/rocky/ssa-interp"
	"github.com/rocky/ssa-interp/gub"
)

func init() {
	parent := "info"
	gub.AddSubCommand(parent, &gub.SubcmdInfo{
		Fn: InfoProgramSubcmd,
		Help: `info program

Prints information about the program including:
*  instruction number
*  block number
*  function number
*  stop event
*  source-code position
`,
		Min_args: 0,
		Max_args: 0,
		Short_help: "Information about debugged program",
		Name: "program",
	})
}

func InfoProgramSubcmd(args []string) {
	fr := gub.CurFrame()
	gub.Msg("instruction number: %d", fr.PC())
	block := fr.Block()
	if block == nil {
		gub.Msg("unknown block")
	} else {
		gub.Msg("basic block: %d", block.Index)
		if block.Scope != nil {
			gub.Msg("scope: %d", block.Scope.ScopeId())
		} else {
			gub.Msg("unknown scope")
		}
	}
	gub.Msg("function: %s", fr.Fn().Name())
	gub.Msg("program stop event: %s", ssa2.Event2Name[gub.TraceEvent])
	gub.Msg("position: %s", gub.CurFrame().PositionRange())
}
