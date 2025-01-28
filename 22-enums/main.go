package main

import (
	"fmt"
)

type serverState int

const (
	StateIdle serverState = iota
	StateConnected
	StateRetrying
	StateError
)

var stateName = map[serverState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateRetrying:  "retrying",
	StateError:     "error",
}

// implement fmt.Sprintf
func (ss serverState) String() string {
	return stateName[ss]
}

func transition(s serverState) serverState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown serverState %s\n", s))
	}
}
func main() {
	nextS := transition(StateIdle)
	fmt.Println(nextS)

	nextS2 := transition(nextS)
	fmt.Println(nextS2)
}
