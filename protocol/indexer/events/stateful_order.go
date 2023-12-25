package events

import (
	v1 "github.com/jinxprotocol/v4-chain/protocol/indexer/protocol/v1"
	"github.com/jinxprotocol/v4-chain/protocol/indexer/shared"
	clobtypes "github.com/jinxprotocol/v4-chain/protocol/x/clob/types"
)

func NewLongTermOrderPlacementEvent(
	order clobtypes.Order,
) *StatefulOrderEventV1 {
	indexerOrder := v1.OrderToIndexerOrder(order)
	orderPlace := StatefulOrderEventV1_LongTermOrderPlacementV1{
		Order: &indexerOrder,
	}
	return &StatefulOrderEventV1{
		Event: &StatefulOrderEventV1_LongTermOrderPlacement{
			LongTermOrderPlacement: &orderPlace,
		},
	}
}

func NewStatefulOrderRemovalEvent(
	removedOrderId clobtypes.OrderId,
	reason shared.OrderRemovalReason,
) *StatefulOrderEventV1 {
	orderId := v1.OrderIdToIndexerOrderId(removedOrderId)
	orderRemoval := StatefulOrderEventV1_StatefulOrderRemovalV1{
		RemovedOrderId: &orderId,
		Reason:         reason,
	}
	return &StatefulOrderEventV1{
		Event: &StatefulOrderEventV1_OrderRemoval{
			OrderRemoval: &orderRemoval,
		},
	}
}

func NewConditionalOrderPlacementEvent(
	order clobtypes.Order,
) *StatefulOrderEventV1 {
	indexerOrder := v1.OrderToIndexerOrder(order)
	orderPlace := StatefulOrderEventV1_ConditionalOrderPlacementV1{
		Order: &indexerOrder,
	}
	return &StatefulOrderEventV1{
		Event: &StatefulOrderEventV1_ConditionalOrderPlacement{
			ConditionalOrderPlacement: &orderPlace,
		},
	}
}

func NewConditionalOrderTriggeredEvent(
	orderId clobtypes.OrderId,
) *StatefulOrderEventV1 {
	indexerOrder := v1.OrderIdToIndexerOrderId(orderId)
	orderTriggered := StatefulOrderEventV1_ConditionalOrderTriggeredV1{
		TriggeredOrderId: &indexerOrder,
	}
	return &StatefulOrderEventV1{
		Event: &StatefulOrderEventV1_ConditionalOrderTriggered{
			ConditionalOrderTriggered: &orderTriggered,
		},
	}
}
