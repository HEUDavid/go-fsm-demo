package internal

import (
	"encoding/json"
	. "github.com/HEUDavid/go-fsm/pkg/metadata"
	"log"
)

var (
	New = GenState("New", false, newHandler)
	Pay = GenState("Pay", false, payHandler)
	End = State[*MyExtData]{Name: "End", IsFinal: true, Handler: nil}
)

var (
	New2Pay = GenTransition(New, Pay)
	Pay2End = GenTransition(Pay, End)
	End2End = GenTransition(End, End)
)

var PayFsm = func() FSM[*MyExtData] {
	fsm := GenFSM(New)
	fsm.RegisterState(New, Pay, End)
	fsm.RegisterTransition(New2Pay, Pay2End, End2End)
	return fsm
}()

func newHandler(task *Task[*MyExtData]) error {
	log.Printf("State: %s, ExtData: %s", task.State, _pretty(task.GetExtData()))

	task.ExtData.Comment = "Modified by newHandler" // Update ExtData
	task.State = Pay.GetName()                      // Switch to next state
	return nil
}

func payHandler(task *Task[*MyExtData]) error {
	log.Printf("State: %s, ExtData: %s", task.State, _pretty(task.GetExtData()))

	// Invoke RPC interfaces to perform certain operations
	// ...

	task.State = End.GetName() // Switch to next state
	return nil
}

func _pretty(v interface{}) string {
	b, _ := json.MarshalIndent(v, "", "  ")
	return string(b)
}
