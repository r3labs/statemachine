package statemachine

import "errors"

// StateMachine ...
type StateMachine struct {
	State       State
	Transitions map[string]Transitions
	Callbacks   map[string][]Callback
	Errors      map[string]error
}

// New : returns a new statemachine
func New(state State) *StateMachine {
	return &StateMachine{
		State:       state,
		Transitions: make(map[string]Transitions),
		Callbacks:   make(map[string][]Callback),
		Errors:      make(map[string]error),
	}
}

// When : sets the transition for a given event
func (s *StateMachine) When(event string, t Transitions) {
	s.Transitions[event] = t
}

// On : sets a callback for a given event
func (s *StateMachine) On(event string, c Callback) {
	s.Callbacks[event] = append(s.Callbacks[event], c)
}

// Trigger : trigger a state change
func (s *StateMachine) Trigger(event string, t interface{}) error {
	return s.change(event, t)
}

// Error : sets an error to return when there are no valid transitions for the current state
func (s *StateMachine) Error(state string, err error) {
	s.Errors[state] = err
}

// ValidateTransition : returns an error if there are no transitions for an event
func (s *StateMachine) ValidateTransition(event, state string) error {
	if s.Transitions[event] == nil {
		return errors.New("invalid event: " + event)
	}

	if state == "" || s.Transitions[event][state] == "" {
		err := s.Errors[state]
		if err != nil {
			return err
		}
		return errors.New("invalid state transition")
	}

	return nil
}

func (s *StateMachine) change(event string, t interface{}) error {
	err := s.ValidateTransition(event, s.State.GetState())
	if err != nil {
		return err
	}

	state := s.Transitions[event][s.State.GetState()]

	for _, cb := range s.Callbacks[state] {
		err := cb(t)
		if err != nil {
			return err
		}
	}

	s.State.SetState(state)
	return nil
}
