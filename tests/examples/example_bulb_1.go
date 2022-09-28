package examples

import (
	"fmt"
	"strings"

	"github.com/Quinn-Fang/quinne/quinne"
	"github.com/Quinn-Fang/quinne/uspace"
)

func CreateBulbDiagram() {
	eventHandler := quinne.NewEventHandler("samples/sample_001.go")
	event, err := eventHandler.GetNextEvent()
	for err == nil {
		fmt.Printf("%+v\n", event)
		if event.GetEventType() == uspace.EventTypeFunction {
			// deal with how function show be executed here
			// and provide the return value
			fFunction := event.GetFunction(event.GetEventContext())
			if fFunction.GetFunctionName() == "create_battery" {

				fFunction.SetReturnValue(true)
				fmt.Printf("%+v %+v\n", fFunction, fFunction.GetReturnValue())
			} else if fFunction.GetFunctionName() == "create_switch" {
				fFunction.SetReturnValue("success")
				fmt.Printf("%+v %+v\n", fFunction, fFunction.GetReturnValue())
			} else if fFunction.GetFunctionName() == "create_bulb" {
				// Get Function arguments, maybe put some check on them ...
				params := fFunction.GetParams()
				for _, v := range params {
					fmt.Printf("%+v ", v.GetVariableValue())
				}

				// Set the return value for this particular function
				fFunction.SetReturnValue("success")
				fmt.Printf("%+v %+v\n\n", fFunction, fFunction.GetReturnValue())
			}
		} else if event.GetEventType() == uspace.EventTypeIfElseExpr {
			// Get the if expression or else-if expression and variables within that you
			// should provide value or assigned automatically if has defined previously

			ifElseExpr, ifElseExprVarNames := event.GetExpr(event.GetEventContext())
			if strings.Contains(ifElseExpr, "SWITCH_ON") {
				fmt.Printf("%+v %+v\n", ifElseExpr, ifElseExprVarNames)
				varMap := make(map[string]interface{}, 8)
				varMap["SWITCH_ON"] = false
				event.SetExpr(varMap)

				// by calling SetExpr with emptyMap would filled variables automatically
				// if already defined and can be accessed by scope rules

				// varMap := make(map[string]interface{}, 8)
				// event.SetExpr()

			}

		}

		event, err = eventHandler.GetNextEvent()

		fmt.Println()
	}
}
