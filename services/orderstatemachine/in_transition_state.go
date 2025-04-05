package orderstatemachine

import "totesbackend/models"

type InTransitState struct {
	context *OrderStateMachine
	state   *models.OrderStateType
}

func NewInTransitState(context *OrderStateMachine) *InTransitState {
	return &InTransitState{
		context: context,
		state: &models.OrderStateType{
			ID:          2,
			Description: "InTransitState",
		},
	}
}

func (s *InTransitState) ChangeState(target OrderState) error {

	return nil
}

func (s *InTransitState) GetId() int {
	return s.state.ID
}

func (s *InTransitState) GetDescription() string {
	return s.state.Description
}
