package task

//Defining a type State which holds the integer value, State is an alias for underlying type `int`
type State int

//iota is special identifier in Go, which simplifies definitions of incrementing numbers.
//Pending = 0, Scheduled = 1, Running = 2, Completed = 3, Failed = 4
const (
	Pending State = iota
	Scheduled
	Running
	Completed
	Failed
)