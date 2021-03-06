/************************************
 *
 * This file contains auto-generated methods to help configure enums.
 * It was generated by the codegen package via 'boardgame-util codegen'.
 *
 * DO NOT EDIT by hand.
 *
 ************************************/

package example

import (
	"github.com/jkomoros/boardgame/enum"
)

var Enums = enum.NewSet()

//ConfigureEnums simply returns Enums, the auto-generated Enums variable. This
//is output because gameDelegate appears to be a struct that implements
//boardgame.GameDelegate, and does not already have a ConfigureEnums
//explicitly defined.
func (g *gameDelegate) ConfigureEnums() *enum.Set {
	return Enums
}

var PhaseEnum = Enums.MustAddTree("Phase", map[int]string{
	Phase:       "",
	PhaseNormal: "Normal",
	PhaseSetUp:  "Set Up",
}, map[int]int{
	Phase:       Phase,
	PhaseNormal: Phase,
	PhaseSetUp:  Phase,
})
