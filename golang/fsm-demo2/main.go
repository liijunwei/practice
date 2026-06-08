package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

func main() {
	orchestrated()
}

// --- Demo ---

type orderPayment struct {
	order   *FSM
	payment *FSM
	och     *Orchestrator
}

func newOrderPayment() *orderPayment {
	order := NewFSM("created")
	order.AddTransition("created", "pay", "paid")
	order.AddTransition("paid", "ship", "shipped")
	order.AddTransition("shipped", "complete", "completed")
	order.AddTransition("created", "cancel", "cancelled")
	order.AddTransition("paid", "cancel", "cancelled")

	payment := NewFSM("pending")
	payment.AddTransition("pending", "process", "processing")
	payment.AddTransition("processing", "succeed", "completed")
	payment.AddTransition("processing", "fail", "failed")
	payment.AddTransition("completed", "refund", "refunded")

	och := NewOrchestrator()
	och.Wire("order", "pay", "payment", "process", nil)
	och.Wire("payment", "succeed", "order", "ship", nil)
	och.Wire("payment", "fail", "order", "cancel", nil)
	och.Wire("order", "cancel", "payment", "refund", func() bool {
		return payment.Current() == "completed"
	})

	och.AddFSM("order", order)
	och.AddFSM("payment", payment)
	och.Observe(func(subject string, t Transition) {
		fmt.Printf("  [%s] %s --(%s)--> %s\n", subject, t.From, t.Event, t.To)
	})

	return &orderPayment{order, payment, och}
}

func orchestrated() {
	fmt.Println("--- Orchestrated: Order + Payment ---")

	// Happy path
	fmt.Println("= Happy path =")
	op := newOrderPayment()
	fmt.Print(op.och.Mermaid())
	fmt.Printf("  order=%-10s payment=%s\n", op.order.Current(), op.payment.Current())
	op.order.Trigger("pay")
	fmt.Printf("  order=%-10s payment=%s\n", op.order.Current(), op.payment.Current())
	op.payment.Trigger("succeed")
	fmt.Printf("  order=%-10s payment=%s\n", op.order.Current(), op.payment.Current())

	// Failure path
	fmt.Println("\n= Failure path =")
	op2 := newOrderPayment()
	fmt.Printf("  order=%-10s payment=%s\n", op2.order.Current(), op2.payment.Current())
	op2.order.Trigger("pay")
	fmt.Printf("  order=%-10s payment=%s\n", op2.order.Current(), op2.payment.Current())
	op2.payment.Trigger("fail")
	fmt.Printf("  order=%-10s payment=%s\n", op2.order.Current(), op2.payment.Current())

	fmt.Println()
}

// --- Orchestrator ---

type wire struct {
	srcFSM   string
	srcEvent Event
	dstFSM   string
	dstEvent Event
	guard    func() bool
}

type Orchestrator struct {
	fsms     map[string]*FSM
	wires    []wire
	observer func(string, Transition)
}

func NewOrchestrator() *Orchestrator {
	return &Orchestrator{fsms: map[string]*FSM{}}
}

// Wire registers a cross-FSM trigger: when srcFSM fires srcEvent,
// dstFSM receives dstEvent. An optional guard can veto the trigger.
func (o *Orchestrator) Wire(srcFSM string, srcEvent Event, dstFSM string, dstEvent Event, guard func() bool) {
	o.wires = append(o.wires, wire{srcFSM, srcEvent, dstFSM, dstEvent, guard})
}

func (o *Orchestrator) Observe(cb func(string, Transition)) {
	o.observer = cb
}

func (o *Orchestrator) AddFSM(name string, fsm *FSM) {
	o.fsms[name] = fsm
	fsm.OnTransition(func(t Transition) {
		if o.observer != nil {
			o.observer(name, t)
		}
		for _, w := range o.wires {
			if w.srcFSM == name && w.srcEvent == t.Event {
				if w.guard != nil && !w.guard() {
					continue
				}
				if dst, ok := o.fsms[w.dstFSM]; ok {
					dst.Trigger(w.dstEvent)
				}
			}
		}
	})
}

// Mermaid returns a mermaid stateDiagram-v2 showing all FSMs and cross-FSM wires.
// Copy the output into any mermaid renderer.
func (o *Orchestrator) Mermaid() string {
	var b strings.Builder
	b.WriteString("```mermaid\n")
	b.WriteString("stateDiagram-v2\n")

	names := make([]string, 0, len(o.fsms))
	for name := range o.fsms {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		f := o.fsms[name]
		fmt.Fprintf(&b, "  state %q as %s {\n", name, name)
		for _, t := range f.trans {
			from := name + "_" + string(t.From)
			to := name + "_" + string(t.To)
			fmt.Fprintf(&b, "    %s --> %s: %s\n", from, to, t.Event)
		}
		b.WriteString("  }\n\n")
	}

	if len(o.wires) > 0 {
		for _, w := range o.wires {
			srcFSM := o.fsms[w.srcFSM]
			dstFSM := o.fsms[w.dstFSM]

			var srcState, dstState string
			for _, t := range srcFSM.trans {
				if t.Event == w.srcEvent {
					srcState = w.srcFSM + "_" + string(t.To)
					break
				}
			}
			for _, t := range dstFSM.trans {
				if t.Event == w.dstEvent {
					dstState = w.dstFSM + "_" + string(t.To)
					break
				}
			}

			guard := ""
			if w.guard != nil {
				guard = " [guard]"
			}

			if srcState != "" && dstState != "" {
				fmt.Fprintf(&b, "  %s --> %s: %s.%s ⇢ %s.%s%s\n",
					srcState, dstState,
					w.srcFSM, w.srcEvent, w.dstFSM, w.dstEvent, guard)
			}
		}
	}

	b.WriteString("```\n")
	return b.String()
}

// --- FSM ---

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
