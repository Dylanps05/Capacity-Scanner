package storage

import (
	"git.11spades.net/CivilMatch/civilmatch/internal/storage/login"
	"git.11spades.net/CivilMatch/civilmatch/internal/storage/race"
	"github.com/jackc/pgx/v5"
)

type DefaultSiteStorage struct {
	capacity.CapacityStorage
	sensor.SensorAuthStorage
}

func (s DefaultSiteStorage) New(db *pgx.Conn) SiteStorage {
	capacity.CapacityStorage = capacity.NewNonvolatileCapacityStorage(db)
	sensor.SensorAuthStorage = sensor.NewNonvolatileSensorAuthStorage(db)

	return &s
}

func (s *DefaultSiteStorage) GetCapacityStorage() capacity.CapacityStorage {
	return s.CapacityStorage
}

func (s *DefaultSiteStorage) GetSensorAuthStorage() sensor.SensorAuthStorage {
	return s.SensorAuthStorage
}
