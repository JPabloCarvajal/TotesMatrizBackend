package orderstatemachine

import (
	"fmt"
	"totesbackend/models"
)

type OrderStateMachine struct {
	currentState  OrderState
	purchaseOrder *models.PurchaseOrder
}

// NewStateMachine construye la máquina y setea el estado actual según el estado de la orden
func NewStateMachine(po *models.PurchaseOrder) (*OrderStateMachine, error) {
	sm := &OrderStateMachine{
		purchaseOrder: po,
	}

	// Determinar estado inicial en base al OrderStateID de la orden
	switch po.OrderStateID {
	case 1:
		sm.currentState = NewIssuedState(sm)
	case 2:
		sm.currentState = NewInTransitState(sm)
	case 3:
		sm.currentState = NewCancelledState(sm)
	case 4:
		sm.currentState = NewApprovedState(sm)
	default:
		return nil, fmt.Errorf("unknown state: %d", po.OrderStateID)
	}

	return sm, nil
}

func (sm *OrderStateMachine) GetCurrentState() OrderState {
	return sm.currentState
}

func (sm *OrderStateMachine) ChangeState(target OrderState) error {
	return sm.currentState.ChangeState(target)
}
