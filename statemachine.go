package statemachine

// StateMachine ...
type StateMachine struct {
	State       State
	Transitions map[string]Transition
	Callbacks   map[string][]Callback
}

// When : sets the transition for a given event
func (s *StateMachine) When(event string, t Transition) {
	s.Transitions[event] = t
}

// On : sets a callback for a given event
func (s *StateMachine) On(event string, c Callback) {
	s.Callbacks[event] = append(s.Callbacks[event], c)
}

// Trigger : trigger a state change
func (s *StateMachine) Trigger(event string) error {
	return nil
}
