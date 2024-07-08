package internal

import (
	. "github.com/HEUDavid/go-fsm/pkg/metadata"
)

var NEW = New{State: State{Name: "New", IsFinal: false}}
var PAY = Pay{State: State{Name: "Pay", IsFinal: false}}

var END = State{Name: "End", IsFinal: true}

var New2Pay = Transition{From: NEW, To: PAY}
var Pay2End = Transition{From: PAY, To: END}

var PayFsm = FSM{
	InitialState: NEW,
	States: map[string]IState{
		NEW.GetName(): NEW,
		PAY.GetName(): PAY,
		END.GetName(): END,
	},
	Transitions: map[string]Transition{
		New2Pay.GetName(): New2Pay,
		Pay2End.GetName(): Pay2End,
	},
}
