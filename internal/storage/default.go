package storage

import (
	"github.com/Dylanps05/Capacity-Scanner/internal/storage/capacity"
	"github.com/Dylanps05/Capacity-Scanner/internal/storage/login"
	"github.com/Dylanps05/Capacity-Scanner/internal/storage/sensor"
	"github.com/Dylanps05/Capacity-Scanner/internal/storage/session"
	"github.com/jackc/pgx/v5"
)

type DefaultSiteStorage struct {
	capacity.CapacityStorage
	sensor.SensorAuthStorage
	session.SessionStorage
	login.LoginStorage
}

func NewDefaultSiteStorage(db *pgx.Conn) SiteStorage {
	s := &DefaultSiteStorage{
		CapacityStorage:   capacity.NewNvCapacityStorage(db),
		SensorAuthStorage: sensor.NewNvSensorAuthStorage(db),
		LoginStorage:      login.NewCachedLoginStorage(db),
		SessionStorage:    session.NewSyncedSessionStorage(),
	}
	return s
}

func (s *DefaultSiteStorage) GetCapacityStorage() capacity.CapacityStorage {
	return s.CapacityStorage
}

func (s *DefaultSiteStorage) GetSensorAuthStorage() sensor.SensorAuthStorage {
	return s.SensorAuthStorage
}
