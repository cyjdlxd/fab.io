package controllers

import (
	"testing"

	"github.com/karlseguin/expect"
)

type TesterController struct {
}

// RegisterActions used to register actions
func (controller *TesterController) RegisterActions(actionsHandler *ActionsHandler) {
	actionsHandler.RegisterConnectedAction(controller.connected)
	actionsHandler.RegisterDisconnectedAction(controller.connected)
	actionsHandler.RegisterErrorAction(controller.error)
}

func (controller *TesterController) connected(context *Context) {
}

func (controller *TesterController) disconnected(context *Context) {
}

func (controller *TesterController) error(context *Context) {
}

type Tester struct {
	manager *Manager
}

func (tester *Tester) RegistorActions() {
	var err error

	tester.manager.RegisterController("test", &TesterController{})

	_ = err
}

func TestController(t *testing.T) {
	manager := &Manager{}

	manager.Setup(nil)

	tester := &Tester{
		manager: manager,
	}

	expect.Expectify(tester, t)
}
