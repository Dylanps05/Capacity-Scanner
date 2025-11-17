package logic

import (
	"github.com/Dylanps05/Capacity-Scanner/internal/logic/capacity"
	"github.com/Dylanps05/Capacity-Scanner/internal/logic/sensor"
)

type Controller interface {
	sensor.SensorModule
	capacity.CapacityModule
}
