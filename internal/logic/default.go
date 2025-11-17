package logic

import (
	"github.com/Dylanps05/Capacity-Scanner/internal/logic/sensor"
	"github.com/Dylanps05/Capacity-Scanner/internal/storage"
)

type DefaultController struct {
	sensor.SensorModule
}

func NewDefaultController(s storage.SiteStorage) Controller {
	return &DefaultController{
		SensorModule: sensor.NewDefaultSensorModule(s.GetSensorAuthStorage(), s.GetCapacityStorage()),
	}
}
