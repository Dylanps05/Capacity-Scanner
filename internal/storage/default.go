package storage

import (
	"github.com/Dylanps05/Capacity-Scanner/internal/storage/capacity"
	"github.com/Dylanps05/Capacity-Scanner/internal/storage/sensor"
	"github.com/jackc/pgx/v5"
)

type DefaultSiteStorage struct {
	capacity.CapacityStorage
	sensor.SensorAuthStorage
}

func (s DefaultSiteStorage) New(db *pgx.Conn) SiteStorage {
	s.CapacityStorage = capacity.NewNvCapacityStorage(db)
	//s.SensorAuthStorage = sensor.NewNvSensorAuthStorage(db)

	return &s
}

func (s *DefaultSiteStorage) GetCapacityStorage() capacity.CapacityStorage {
	return s.CapacityStorage
}

func (s *DefaultSiteStorage) GetSensorAuthStorage() sensor.SensorAuthStorage {
	return s.SensorAuthStorage
}
