package example

//boardgame:codegen
const (
	//Because the naked Phase exists, this will be a TreeEnum. See package doc for "boardgame/enum" for more.
	Phase = iota
	PhaseSetUp
	PhaseNormal
)
