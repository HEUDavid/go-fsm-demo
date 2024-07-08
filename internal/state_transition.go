package internal

import (
	. "github.com/HEUDavid/go-fsm/pkg/metadata"
)

var (
	New2Pay = Transition{From: NEW, To: PAY}
	Pay2End = Transition{From: PAY, To: END}
	End2End = Transition{From: END, To: END}
)

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
		End2End.GetName(): End2End,
	},
}
