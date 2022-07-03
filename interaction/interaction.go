package interaction

import "time"

const (
	INTERACTION_BUTTON_DATA InteractionComponentDataType = iota + 1
)

var InteractionComponentDataList []InteractionComponentData

type InteractionComponentDataType uint8
type InteractionComponentDataDefault struct {
	Received  bool
	CreatedOn time.Duration
	DeleteFor time.Duration
}

type InteractionComponentData interface {
	Type() InteractionComponentDataType
	Delete() func() string
}

type InteractionButtonData struct {
	InteractionComponentDataDefault
}

func (InteractionButtonData) Type() InteractionComponentDataType {
	return INTERACTION_BUTTON_DATA
}

func (InteractionButtonData) Delete() string {
	return "tata"
}