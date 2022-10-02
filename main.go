package main

// 47 min

import (
	"golang/arrays"
	"golang/conditional"
	"golang/datatype"
	"golang/functions"
	"golang/loops"
	"golang/maths"
	"golang/pointers"
	"golang/runes"
	"golang/slices"
	"golang/strings"
	"golang/terminal"
	"golang/times"
)

func main() {
	// Arrays funcs
	arrays.Arrays()

	// Datatype funcs
	datatype.MainDataTypes()
	datatype.ConvertDataTypes()
	datatype.PrintDataTypes()

	// Functions funcs
	functions.Main()

	// Conditional funcs
	conditional.IfConditionals()

	// Loops funcs
	loops.Loops()
	loops.While()
	loops.Ranges()

	// Maths funcs
	maths.Maths()
	maths.RandomNumbers()

	// Pointers funcs
	pointers.Pointers()

	// Runes funcs
	runes.Runes()

	// Slices funcs
	slices.SlicesRunesAndStrings()
	slices.SlicesDataTypes()

	// Strings funcs
	strings.FStrings()

	// Terminal funcs
	terminal.GetUsername()

	// Times funcs
	times.Times()
}
