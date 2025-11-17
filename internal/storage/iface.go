package storage

import (
	"git.11spades.net/CivilMatch/civilmatch/internal/storage/capacity"
	"git.11spades.net/CivilMatch/civilmatch/internal/storage/sensor"
	"github.com/jackc/pgx/v5"
)

type SiteStorage interface {
	New(db *pgx.Conn) SiteStorage
	GetCapacityStorage() capacity.CapacityStorage
	GetSensorAuthStorage() sensor.SensorAuthStorage
}
