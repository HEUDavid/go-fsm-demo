package internal

import (
	"github.com/HEUDavid/go-fsm-demo/internal/pkg"
	. "github.com/HEUDavid/go-fsm/pkg/metadata"
	"github.com/HEUDavid/go-fsm/pkg/util"
	"log"
)

var NEW = New{State: State{Name: "New", IsFinal: false}}
var PAY = New{State: State{Name: "Pay", IsFinal: false}}

var END = State{Name: "End", IsFinal: true}

type New struct{ State }

func (s New) Handle(task *Task[ExtDataEntity]) error {
	extData, _ := util.Assert[*MyExtData](task.ExtData)
	log.Printf("State: %s, ExtData: %s", task.State, pkg.Pretty(extData))

	extData.Symbol = "BNB"     // Update ExtData
	task.State = PAY.GetName() // Switch to next state
	return nil
}

type Pay struct{ State }

func (s Pay) Handle(task *Task[ExtDataEntity]) error {
	extData, _ := util.Assert[*MyExtData](task.ExtData)
	log.Printf("State: %s, ExtData: %s", task.State, pkg.Pretty(extData))

	// Invoke RPC interfaces to perform certain operations
	// ...

	task.State = END.GetName() // Switch to next state
	return nil
}
