package logic

import (
	"github.com/Dylanps05/Capacity-Scanner/internal/logic/capacity"
	"github.com/Dylanps05/Capacity-Scanner/internal/logic/sensor"
	"github.com/Dylanps05/Capacity-Scanner/internal/storage"
)

type DefaultController struct {
	sensor.SensorModule
	capacity.CapacityModule
}

func NewDefaultController(s storage.SiteStorage) Controller {
	return &DefaultController{
		SensorModule:   sensor.NewDefaultSensorModule(s.GetSensorAuthStorage(), s.GetCapacityStorage()),
		CapacityModule: capacity.NewDefaultCapacityModule(s.GetCapacityStorage()),
	}
}
