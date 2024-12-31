package internal

import (
	. "github.com/HEUDavid/go-fsm/pkg/metadata"
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
	fsm := GenFSM[*MyData]("PayFSM")
	fsm.RegisterState(New, Pay, End)
	fsm.RegisterTransition(New2Pay, Pay2End, End2End)
	return fsm
}()

func newHandler(task *Task[*MyData]) error {
	// It may be necessary to perform some checks.
	// It may be necessary to pre-record the request to the database to ensure idempotency.
	// For example, generating some request IDs.
	// ...

	task.Data.Comment = "Modified by newHandler" // Update Data
	task.State = Pay.GetName()                   // Switch to next state
	return nil
}

func payHandler(task *Task[*MyData]) error {
	// Invoke RPC interfaces to perform certain operations.
	// ...

	task.Data.Comment = "Modified by payHandler"
	task.Data.Operator = "system" // Update Data
	task.State = End.GetName()    // Switch to next state
	return nil
}
