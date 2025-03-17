package models

type Lifecycle int

const (
	Initialing Lifecycle = iota
	Running
	Stopping
)
