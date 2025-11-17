package sensor

import (
	"context"
	"fmt"
	"github.com/Dylanps05/Capacity-Scanner/internal/cstype"
	"github.com/jackc/pgx/v5"
	"reflect"
)

type NvSensorAuthStorage struct {
	db *pgx.Conn
}

func NewNvSensorAuthStorage(db *pgx.Conn) SensorAuthStorage {
	return &NvSensorAuthStorage{
		db: db,
	}
}

func (s *NvSensorAuthStorage) StoreToken(ctx context.Context, t cstype.ScannerToken) error {
	stmt := `insert into sensor
        values (token) = ($1)`
	tx_any := ctx.Value("pgxtx")
	if tx_any == nil {
		return fmt.Errorf("%T: was passed nill tx", reflect.TypeOf(s))
	}
	tx := tx_any.(pgx.Tx)
	_, err := tx.Exec(ctx, stmt, t)
	if err != nil {
		err = fmt.Errorf("%T: failed to add token to tx\n%w", reflect.TypeOf(s), err)
	}

	return err
}

func (s *NvSensorAuthStorage) HaveToken(ctx context.Context, t cstype.ScannerToken) (bool, error) {
	stmt := `select *
        from sensor
        where token = ($1)`
	rows, err := s.db.Query(ctx, stmt)
	if err != nil {
		err = fmt.Errorf("%T: failed to query for token\n%w", reflect.TypeOf(s), err)
	}
	return rows.Next(), err
}

func (s *NvSensorAuthStorage) RevokeToken(ctx context.Context, t cstype.ScannerToken) error {
	stmt := `delete from sensor
        where token = ($1)`
	tx_any := ctx.Value("pgxtx")
	if tx_any == nil {
		return fmt.Errorf("%T: was passed nill tx", reflect.TypeOf(s))
	}
	tx := tx_any.(pgx.Tx)
	_, err := tx.Exec(ctx, stmt, t)
	if err != nil {
		err = fmt.Errorf("%T: failed to add token revocation to tx\n%w", reflect.TypeOf(s), err)
	}

	return err
}
