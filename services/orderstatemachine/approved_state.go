package orderstatemachine

import "totesbackend/models"

type ApprovedState struct {
	context *OrderStateMachine
	state   *models.OrderStateType
}

func NewApprovedState(context *OrderStateMachine) *ApprovedState {
	return &ApprovedState{
		context: context,
		state: &models.OrderStateType{
			ID:          4,
			Description: "ApprovedState",
		},
	}
}

func (s *ApprovedState) ChangeState(target OrderState) error {
	// lógica de transición
	return nil
}

func (s *ApprovedState) GetId() int {
	return s.state.ID
}

func (s *ApprovedState) GetDescription() string {
	return s.state.Description
}
