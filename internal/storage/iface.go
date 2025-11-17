package storage

import (
	"github.com/Dylanps05/Capacity-Scanner/internal/storage/capacity"
	"github.com/Dylanps05/Capacity-Scanner/internal/storage/sensor"
)

type SiteStorage interface {
	GetCapacityStorage() capacity.CapacityStorage
	GetSensorAuthStorage() sensor.SensorAuthStorage
}
