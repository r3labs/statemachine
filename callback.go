package statemachine

// Callback ...
type Callback func(state string, t interface{}) error
