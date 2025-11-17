package storage

import (
	"github.com/Dylanps05/Capacity-Scanner/internal/storage/capacity"
	"github.com/Dylanps05/Capacity-Scanner/internal/storage/sensor"
	"github.com/jackc/pgx/v5"
)

type SiteStorage interface {
	New(db *pgx.Conn) SiteStorage
	GetCapacityStorage() capacity.CapacityStorage
	GetSensorAuthStorage() sensor.SensorAuthStorage
}
