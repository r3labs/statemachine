package statemachine

// State : stores internal state
type State interface {
	GetState() string
	SetState(string)
}
