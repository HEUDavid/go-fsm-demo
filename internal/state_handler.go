package internal

import (
	"github.com/HEUDavid/go-fsm-demo/internal/pkg"
	. "github.com/HEUDavid/go-fsm/pkg/metadata"
	"log"
)

var (
	New = GenState("New", false, newHandler)
	Pay = GenState("Pay", false, payHandler)
	End = State[*MyExtData]{Name: "End", IsFinal: true, ReHandle: nil}
)

func newHandler(task *Task[*MyExtData]) error {
	log.Printf("State: %s, ExtData: %s", task.State, pkg.Pretty(task.GetExtData()))

	task.ExtData.Comment = "Modified by newHandler" // Update ExtData
	task.State = Pay.GetName()                      // Switch to next state
	return nil
}

func payHandler(task *Task[*MyExtData]) error {
	log.Printf("State: %s, ExtData: %s", task.State, pkg.Pretty(task.GetExtData()))

	// Invoke RPC interfaces to perform certain operations
	// ...

	task.State = End.GetName() // Switch to next state
	return nil
}
