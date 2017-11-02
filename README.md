# StateMachine
A minimalistic state machine implementation, loosely based on github.com/soveran/micromachine.

## Usage

### State

To create a new state machine, a type that satisfies the State interface must be set. i.e.

```go
type State struct {
	state string
}

func New(state string) *State {
	return &State{state: state}
}

func (s *State) GetState() string {
	return s.state
}

func (s *State) SetState(state string) {
	s.state = state
}
```

### Transitions, Callbacks and Errors

Transitions are assigned using the `When` method

```go
import (
    "fmt"
    "github.com/r3labs/statemachine"
)

func main() {
    s := New("initializing")
	sm := statemachine.New(s)

    // Assign all transitions for a given event
	sm.When("accept", statemachine.Transitions{"initializing": "accepted"})
	sm.When("reject", statemachine.Transitions{"initializing": "rejected"})
	sm.When("reset", statemachine.Transitions{"accepted": "initializing", "rejected": "initializing"})
}
```

Callbacks can be added for execution before transitioning to a state. Multiple callbacks for an event are supported. In the event of a callback failure, the transition will be halted
```go
    // Additional data can be passed into the callback via an interface
	sm.On("accepted", func(state string, t interface{}) error {
        // produces new state: accepted
		fmt.Println("new state: " + state)   

		return nil
	})
```

State transitions can be triggered with the `Trigger` method
```go
    // Returns an error
    err := sm.Trigger("unknown", nil)

    // Returns nil and sets state to 'accepted'
    err := sm.Trigger("accept", nil)
```

If you wish to pass additional data into the callbacks, you can do so as following
```go
    sm.On("accepted", func(state string, t interface{}) error {
        data := t.(*map[string]string)
        fmt.Printf("welcome back %s!" + data["user"])   

        return nil
    })

    // Will print 'welcome back john!'
    sm.Trigger("accept", &map[string]string{"user": "john"})
```

Default errors can be specified for when a transition fails
```go
    sm.Error("accepted", errors.New("unable to transition as state is currently 'accepted'"))

    err := sm.Trigger("reject", nil)
```

## Contributing

Please read through our
[contributing guidelines](CONTRIBUTING.md).
Included are directions for opening issues, coding standards, and notes on
development.

Moreover, if your pull request contains patches or features, you must include
relevant unit tests.


## Copyright and License

Code and documentation copyright since 2015 r3labs.io authors.

Code released under
[the Mozilla Public License Version 2.0](LICENSE).
