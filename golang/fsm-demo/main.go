package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type State string
type Event string

type Transition struct {
	From  State
	Event Event
	To    State
}

type FSM struct {
	current State
	states  map[State]bool
	events  map[Event]bool
	trans   []Transition
	onTrans func(Transition)
}

func NewFSM(initial State) *FSM {
	return &FSM{
		current: initial,
		states:  map[State]bool{initial: true},
		events:  map[Event]bool{},
	}
}

func (f *FSM) AddTransition(from State, event Event, to State) {
	f.states[from] = true
	f.states[to] = true
	f.events[event] = true
	f.trans = append(f.trans, Transition{From: from, Event: event, To: to})
}

func (f *FSM) OnTransition(cb func(Transition)) { f.onTrans = cb }

func (f *FSM) Current() State { return f.current }

func (f *FSM) Trigger(event Event) error {
	for _, t := range f.trans {
		if t.From == f.current && t.Event == event {
			f.current = t.To
			if f.onTrans != nil {
				f.onTrans(t)
			}
			return nil
		}
	}
	return errors.New("illegal transition")
}

func (f *FSM) Can(event Event) bool {
	for _, t := range f.trans {
		if t.From == f.current && t.Event == event {
			return true
		}
	}
	return false
}

func (f *FSM) Dump() string {
	states := make([]State, 0, len(f.states))
	for s := range f.states {
		states = append(states, s)
	}

	events := make([]Event, 0, len(f.events))
	for e := range f.events {
		events = append(events, e)
	}

	type dump struct {
		Current     State        `json:"current"`
		States      []State      `json:"states"`
		Events      []Event      `json:"events"`
		Transitions []Transition `json:"transitions"`
	}

	b, _ := json.MarshalIndent(dump{
		Current:     f.current,
		States:      states,
		Events:      events,
		Transitions: f.trans,
	}, "", "  ")
	return string(b)
}

func orderDemo() {
	fmt.Println("--- Order Lifecycle ---")

	fsm := NewFSM("created")
	fsm.AddTransition("created", "pay", "paid")
	fsm.AddTransition("paid", "ship", "shipped")
	fsm.AddTransition("shipped", "complete", "completed")
	fsm.AddTransition("created", "cancel", "cancelled")
	fsm.AddTransition("paid", "cancel", "cancelled")

	fsm.OnTransition(func(t Transition) {
		fmt.Printf("  [%s] --(%s)--> [%s]\n", t.From, t.Event, t.To)
	})

	fmt.Printf("  start: %s\n", fsm.Current())
	fsm.Trigger("pay")
	fsm.Trigger("ship")
	fsm.Trigger("complete")
	fmt.Printf("  end:   %s\n", fsm.Current())

	// illegal transition
	fsm2 := NewFSM("created")
	fsm2.AddTransition("created", "pay", "paid")
	fmt.Printf("  ship from created: %v\n", fsm2.Trigger("ship"))

	// can query
	fsm3 := NewFSM("paid")
	fsm3.AddTransition("paid", "ship", "shipped")
	fmt.Printf("  paid can ship? %v, can pay? %v\n", fsm3.Can("ship"), fsm3.Can("pay"))
	fmt.Println()
}

func trafficLightDemo() {
	fmt.Println("--- Traffic Light ---")

	fsm := NewFSM("red")
	fsm.AddTransition("red", "tick", "green")
	fsm.AddTransition("green", "tick", "yellow")
	fsm.AddTransition("yellow", "tick", "red")

	fsm.OnTransition(func(t Transition) {
		fmt.Printf("  %s -> %s\n", t.From, t.To)
	})

	for range 6 {
		fsm.Trigger("tick")
	}
	fmt.Println()
}

func main() {
	orderDemo()
	trafficLightDemo()

	fsm := NewFSM("created")
	fsm.AddTransition("created", "pay", "paid")
	fsm.AddTransition("paid", "ship", "shipped")
	fsm.AddTransition("shipped", "complete", "completed")
	fsm.AddTransition("created", "cancel", "cancelled")
	fsm.AddTransition("paid", "cancel", "cancelled")

	fmt.Println("--- FSM Structure ---")
	fmt.Println(fsm.Dump())
}
