package login

import (
	"context"
	"errors"
	"sync"

	"github.com/Dylanps05/Capacity-Scanner/internal/cstype"
	"github.com/jackc/pgx/v5"
)

const LOGIN_TABLE = "login"

type LoginCache struct {
	mux              *sync.RWMutex
	loginsByUUID     map[cstype.UserID]cstype.Login
	loginsByUsername map[string]cstype.Login
}

func (c LoginCache) New() LoginCache {
	c.mux = &sync.RWMutex{}
	c.loginsByUUID = map[cstype.UserID]cstype.Login{}
	c.loginsByUsername = map[string]cstype.Login{}

	return c
}

func (c *LoginCache) StoreLogin(l cstype.Login) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.loginsByUUID[l.UUID] = l
	c.loginsByUsername[l.Username] = l
}

func (c *LoginCache) GetLoginByUUID(id cstype.UserID) (cstype.Login, bool) {
	c.mux.RLock()
	defer c.mux.RUnlock()
	l, ok := c.loginsByUUID[id]
	return l, ok
}

func (c *LoginCache) GetLoginByUsername(u string) (cstype.Login, bool) {
	c.mux.RLock()
	defer c.mux.RUnlock()
	l, ok := c.loginsByUsername[u]
	return l, ok
}

func (c *LoginCache) DeleteLogin(id cstype.UserID) cstype.Login {
	c.mux.Lock()
	defer c.mux.Unlock()
	l, ok := c.loginsByUUID[id]
	if !ok {
		return cstype.Login{}
	}

	delete(c.loginsByUsername, l.Username)
	delete(c.loginsByUUID, id)

	return l
}

type CachedLogins struct {
	db    *pgx.Conn
	cache LoginCache
}

func NewCachedLoginStorage(db *pgx.Conn) LoginStorage {
	return &CachedLogins{
		db:    db,
		cache: LoginCache{}.New(),
	}
}

func (s *CachedLogins) StoreLogin(ctx context.Context, l cstype.Login) (cstype.Login, error) {
	stmt := `insert into ` + LOGIN_TABLE + ` (username, password_version, password_salt, password_hash)
        values ($1, $2, $3, $4)
        returning *`
	rows, err := s.db.Query(ctx, stmt, l.Username, l.Password.Version, l.Password.Salt, l.Password.Hash)
	if err != nil {
		errmsg := "CachedLogins: Failed to store new login to db:\n" + err.Error()
		return cstype.Login{}, errors.New(errmsg)
	}
	l, err = pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[cstype.Login])
	if err != nil {
		errmsg := "CachedLogins: Failed to collect or deserialize login after storing to db:\n" + err.Error()
		return l, errors.New(errmsg)
	}

	s.cache.StoreLogin(l)
	return l, nil
}

func (s *CachedLogins) GetLoginByUsername(ctx context.Context, u string) (cstype.Login, error) {
	l, ok := s.cache.GetLoginByUsername(u)
	if ok {
		return l, nil
	}

	stmt := `
        select *
        from ` + LOGIN_TABLE + `
        where username = $1`
	rows, err := s.db.Query(ctx, stmt, u)
	if err != nil {
		errmsg := "CachedLogins: Failed to query for login with username " + u + ":\n" + err.Error()
		return cstype.Login{}, errors.New(errmsg)
	}
	l, err = pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[cstype.Login])
	if err != nil {
		errmsg := "CachedLogins: Failed to collect query result when getting login with username " + u + ":\n" + err.Error()
		return cstype.Login{}, errors.New(errmsg)
	}

	return l, nil
}

func (s *CachedLogins) GetLoginByUUID(ctx context.Context, id cstype.UserID) (cstype.Login, error) {
	l, ok := s.cache.GetLoginByUUID(id)
	if ok {
		return l, nil
	}

	stmt := `
        select *
        from ` + LOGIN_TABLE + `
        where id = $1`
	rows, err := s.db.Query(ctx, stmt, id)
	if err != nil {
		errmsg := "CachedLogins: Failed to query for login with id " + string(id) + ":\n" + err.Error()
		return cstype.Login{}, errors.New(errmsg)
	}
	l, err = pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[cstype.Login])
	if err != nil {
		errmsg := "CachedLogins: Failed to collect query result when getting login with id " + string(id) + ":\n" + err.Error()
		return cstype.Login{}, errors.New(errmsg)
	}

	return l, nil
}

func (s *CachedLogins) SetLoginByUUID(ctx context.Context, l cstype.Login) (cstype.Login, error) {
	stmt := `update ` + LOGIN_TABLE + `
        set username = $1,
        password_version = $2,
        password_salt = $3,
        password_hash = $4,
        where id = $5
        returning *`
	rows, err := s.db.Query(ctx, stmt, l.Username, l.Password.Version, l.Password.Salt, l.Password.Hash, l.UUID)
	if err != nil {
		errmsg := "CachedLogins: Failed to update login from db by id:\n" + err.Error()
		return cstype.Login{}, errors.New(errmsg)
	}
	l, err = pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[cstype.Login])
	if err != nil {
		errmsg := "CachedLogins: Failed to collect and deserialize login from db after update:\n" + err.Error()
		return l, errors.New(errmsg)
	}
	s.cache.StoreLogin(l)

	return l, nil
}

func (s *CachedLogins) DeleteLoginByUUID(ctx context.Context, id cstype.UserID) (cstype.Login, error) {
	stmt := `
        delete
        from ` + LOGIN_TABLE + `
        where id = $1
        returning *`
	rows, err := s.db.Query(ctx, stmt, id)
	if err != nil {
		errmsg := "CachedLogins: Failed to delete login id " + string(id) + ":\n" + err.Error()
		return cstype.Login{}, errors.New(errmsg)
	}
	l, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[cstype.Login])
	if err != nil {
		errmsg := "CachedLogins: Failed to collect query result when deleting login id " + string(id) + ":\n" + err.Error()
		return l, errors.New(errmsg)
	}

	s.cache.DeleteLogin(id)

	return l, nil
}
