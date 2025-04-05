package orderstatemachine

import "totesbackend/models"

type CancelledState struct {
	context *OrderStateMachine
	state   *models.OrderStateType
}

func NewCancelledState(context *OrderStateMachine) *CancelledState {
	return &CancelledState{
		context: context,
		state: &models.OrderStateType{
			ID:          3,
			Description: "CancelledState",
		},
	}
}

func (s *CancelledState) ChangeState(target OrderState) error {
	// lógica de transición
	return nil
}

func (s *CancelledState) GetId() int {
	return s.state.ID
}

func (s *CancelledState) GetDescription() string {
	return s.state.Description
}
