package io

type Command int

const (
	Unknown Command = iota
	Forward
	Back
	TogglePlay
	Quit
)
