package capacity

import (
	"context"
	"fmt"

	"reflect"

	"github.com/jackc/pgx/v5"
)

type NvCapacityStorage struct {
	db *pgx.Conn
}

func NewNvCapacityStorage(db *pgx.Conn) CapacityStorage {
	return &NvCapacityStorage{
		db: db,
	}
}

func (s *NvCapacityStorage) StoreCapacity(ctx context.Context, c int) error {
	stmt := `insert into capacity
        values (pop) = ($1)`
	tx_any := ctx.Value("pgxtx")
	if tx_any == nil {
		return fmt.Errorf("%T: was passed nill tx", reflect.TypeOf(s))
	}
	tx := tx_any.(pgx.Tx)
	_, err := tx.Exec(ctx, stmt, c)
	if err != nil {
		err = fmt.Errorf("%T: failed to add capacity delta to tx\n%w", reflect.TypeOf(s), err)
	}

	return err
}

func (s *NvCapacityStorage) GetCapacity(ctx context.Context) (int, error) {
	stmt := `select pop
        from capacity
        order by time desc
        limit 1`
	row := s.db.QueryRow(ctx, stmt)
	pop := 0
	err := row.Scan(&pop)
	if err != nil {
		err = fmt.Errorf("%T: failed to scan pop\n%w", reflect.TypeOf(s), err)
	}
	return pop, err
}
