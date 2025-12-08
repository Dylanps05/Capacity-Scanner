package storage

import (
	"github.com/Dylanps05/Capacity-Scanner/internal/storage/capacity"
	"github.com/Dylanps05/Capacity-Scanner/internal/storage/login"
	"github.com/Dylanps05/Capacity-Scanner/internal/storage/sensor"
	"github.com/Dylanps05/Capacity-Scanner/internal/storage/session"
)

type SiteStorage interface {
	capacity.CapacityStorage
	sensor.SensorAuthStorage
	session.SessionStorage
	login.LoginStorage
	GetCapacityStorage() capacity.CapacityStorage
	GetSensorAuthStorage() sensor.SensorAuthStorage
}
