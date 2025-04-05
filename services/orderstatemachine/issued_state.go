package orderstatemachine

import "totesbackend/models"

type IssuedState struct {
	context *OrderStateMachine
	state   *models.OrderStateType
}

func NewIssuedState(context *OrderStateMachine) *IssuedState {
	return &IssuedState{
		context: context,
		state: &models.OrderStateType{
			ID:          1,
			Description: "IssuedState",
		},
	}
}

func (s *IssuedState) ChangeState(target OrderState) error {
	// lógica de transición
	return nil
}

func (s *IssuedState) GetId() int {
	return s.state.ID
}

func (s *IssuedState) GetDescription() string {
	return s.state.Description
}
