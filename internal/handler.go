package internal

import (
	"encoding/json"
	. "github.com/HEUDavid/go-fsm/pkg/metadata"
	"log"
)

var (
	New = GenState("New", false, newHandler)
	Pay = GenState("Pay", false, payHandler)
	End = State[*MyData]{Name: "End", IsFinal: true, Handler: nil}
)

var (
	New2Pay = GenTransition(New, Pay)
	Pay2End = GenTransition(Pay, End)
	End2End = GenTransition(End, End)
)

var PayFSM = func() FSM[*MyData] {
	fsm := GenFSM(New)
	fsm.RegisterState(New, Pay, End)
	fsm.RegisterTransition(New2Pay, Pay2End, End2End)
	return fsm
}()

func newHandler(task *Task[*MyData]) error {
	log.Printf("[FSM] State: %s, Task.Data: %s", task.State, _pretty(task.GetData()))

	// It may be necessary to perform some checks.
	// It may be necessary to pre-record the request to the database to ensure idempotency.
	// For example, generating some request IDs.
	// ...

	task.Data.Comment = "Modified by newHandler" // Update Data
	task.State = Pay.GetName()                   // Switch to next state
	return nil
}

func payHandler(task *Task[*MyData]) error {
	log.Printf("[FSM] State: %s, Task: %s", task.State, _pretty(task))

	// Invoke RPC interfaces to perform certain operations.
	// ...

	task.Data.Operator = "system"
	task.State = End.GetName() // Switch to next state
	return nil
}

func _pretty(v interface{}) string {
	s, _ := json.MarshalIndent(v, "", "  ")
	return string(s)
}
