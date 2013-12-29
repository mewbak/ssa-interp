// Copyright 2013 Rocky Bernstein.
// Debugger subcommands

package gub

import (
	"sort"
	"strings"
	"code.google.com/p/go-columnize"
)

type SubcmdFunc func([]string)

type SubcmdInfo struct {
	Help string
	Short_help string
	Min_args int
	Max_args int
	Fn SubcmdFunc
	Name string
}

type SubcmdMap map[string]*SubcmdInfo

// SubcmdMgr contains the name of a debugger command that has
// subcommands, "info", and "set" are examples of such commmands,and
// information about each of the subcommands.
type SubcmdMgr struct {
	Name string
	Subcmds SubcmdMap
}

// Subcmds is a map of a debugger subcommand name to information about
// that subcommand. For example, the "info" command has subcommands
// "break", "program", "line" and so on.
var Subcmds map[string]*SubcmdInfo  = make(map[string]*SubcmdInfo)

// AddSubCommand adds to the subcommand manager mgrName, subcommand
// subcmdInfo.
func AddSubCommand(mgrName string, subcmdInfo *SubcmdInfo) {
	Subcmds[mgrName] = subcmdInfo
	mgr := Cmds[mgrName]
	if mgr != nil {
		mgr.SubcmdMgr.Subcmds[subcmdInfo.Name] = subcmdInfo
	} else {
		Errmsg("Internal error: can't find command '%s' to add to", subcmdInfo.Name)
	}
}

func ListSubCommandArgs(mgr *SubcmdMgr) {
	Section("List of " + mgr.Name + " commands")
	subcmds := mgr.Subcmds

	names := make([]string, len(subcmds))
	i := 0
	for name, _ := range subcmds {
		names[i] = name
		i++
	}
	sort.Strings(names)

	for _, name := range names {
		Msg("%-10s -- %s", name, subcmds[name].Short_help)
	}
}

func HelpSubCommand(subcmdMgr *SubcmdMgr, args []string) {
	if len(args) == 2 {
		Msg(Cmds[subcmdMgr.Name].Help)
	} else {
		what := args[2]
		if what == "*" {
			var names []string
			for name, _ := range subcmdMgr.Subcmds {
				names = append(names, name)
			}
			Section("All %s subcommand names:", subcmdMgr.Name)
			sort.Strings(names)
			opts := columnize.DefaultOptions()
			opts.DisplayWidth = Maxwidth
			mems := strings.TrimRight(columnize.Columnize(names, opts),
				"\n")
			Msg(mems)
		} else if info := subcmdMgr.Subcmds[what]; info != nil {
			Msg(info.Help)
		} else {
			Errmsg("Can't find help for subcommand '%s' in %s", what, subcmdMgr.Name)
		}
	}
}
