package auth

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"net/http"
	"time"

	"github.com/Dylanps05/Capacity-Scanner/internal/cstype"
	"golang.org/x/crypto/argon2"
)

type DefaultAuthModuleRepo interface {
	StoreLogin(ctx context.Context, login cstype.Login) (cstype.Login, error)
	GetLoginByUsername(ctx context.Context, username string) (cstype.Login, error)
	GetLoginByUUID(ctx context.Context, uuid cstype.UserID) (cstype.Login, error)

	StoreSession(ctx context.Context, s string) error
	GetSessionUUID(ctx context.Context, token string) (cstype.UserID, error)
	SetSessionUUID(ctx context.Context, s string, u cstype.UserID) error
}

type DefaultAuthModule struct {
	repo DefaultAuthModuleRepo
}

func NewDefaultAuthModule(r DefaultAuthModuleRepo) AuthModule {
	m := &DefaultAuthModule{
		repo: r,
	}

	return m
}

func (m *DefaultAuthModule) NewSession(ctx context.Context) string {
	err := errors.New("")
	new_session := ""
	for err != nil {
		new_session = rand.Text()
		err = m.repo.StoreSession(ctx, new_session)
	}

	return new_session
}

func (m *DefaultAuthModule) hashPassword(ctx context.Context, p string, s string) (string, error) {
	raw_pass := []byte(p)
	raw_salt, err := hex.DecodeString(s)
	if err != nil {
		return "", errors.New("DefaultAuthModule: Failed to hash password, salt couldn't be decoded:\n" + err.Error())
	}
	time := uint32(1)
	mem := uint32(64 * 1024)
	threads := uint8(4)
	key_len := uint32(32)

	hash := hex.EncodeToString(argon2.IDKey(raw_pass, raw_salt, time, mem, threads, key_len))
	return hash, nil
}

func (m *DefaultAuthModule) CreateLogin(ctx context.Context, u string, p string) error {
	salt := hex.EncodeToString([]byte(rand.Text()))
	hash, err := m.hashPassword(ctx, p, salt)
	password_challenge := cstype.Password{
		Version: 0,
		Salt:    salt,
		Hash:    hash,
	}

	new_login := cstype.Login{
		Username: u,
		Password: password_challenge,
	}
	_, err = m.repo.StoreLogin(ctx, new_login)
	if err != nil {
		return errors.New("DefaultAuthModule: Failed to create login, LoginStorage returned error:\n" + err.Error())
	}

	return nil
}

func (m *DefaultAuthModule) GetSessionFromCookie(ctx context.Context, c *http.Cookie) (string, error) {
	// TODO: Move this and more to cookie validation function
	expired := c.Expires.After(time.Now())
	if expired {
		return "", errors.New("Session cookie is expired.")
	}

	session := c.Value
	return session, nil
}

// TODO: Make this into a submodule with hash upgrading capability
func (m *DefaultAuthModule) validatePassword(ctx context.Context, p string, l cstype.Login) (bool, error) {
	attempted_hash, err := m.hashPassword(ctx, p, l.Password.Salt)
	if err != nil {
		return false, errors.New("DefaultAuthModule: Failed to hash password for validation:\n" + err.Error())
	}

	canon_hash := l.Password.Hash
	password_valid := attempted_hash == canon_hash
	return password_valid, nil
}

func (m *DefaultAuthModule) SetSessionAuthByLogin(ctx context.Context, s string, u string, p string) error {
	canon_login, err := m.repo.GetLoginByUsername(ctx, u)
	if err != nil {
		return err
	}

	password_valid, err := m.validatePassword(ctx, p, canon_login)
	if err != nil {
		return errors.New("DefaultAuthModule: Failed to validate password:\n" + err.Error())
	}
	if !password_valid {
		return errors.New("DefaultAuthModule: Password invalid.")
	}

	m.repo.SetSessionUUID(ctx, s, canon_login.UUID)

	return nil
}

func (m *DefaultAuthModule) GetUsernameFromSession(ctx context.Context, s string) (string, error) {
	uuid, err := m.repo.GetSessionUUID(ctx, s)
	if err != nil {
		errmsg := "DefaultAuthModule: Failed to extract username from session:\n" + err.Error()
		return "", errors.New(errmsg)
	}

	login, err := m.repo.GetLoginByUUID(ctx, uuid)
	if err != nil {
		errmsg := "DefaultAuthModule: Failed to get username from session UUID:\n" + err.Error()
		return "", errors.New(errmsg)
	}

	return login.Username, nil
}

func (m *DefaultAuthModule) GetUserIDFromSession(ctx context.Context, s string) (cstype.UserID, error) {
	uuid, err := m.repo.GetSessionUUID(ctx, s)
	if err != nil {
		errmsg := "DefaultAuthModule: Failed to extract ID from session:\n" + err.Error()
		return "", errors.New(errmsg)
	}

	return uuid, nil
}

func (m *DefaultAuthModule) ValidateSession(ctx context.Context, s string) bool {
	_, err := m.repo.GetSessionUUID(ctx, s)
	if err != nil {
		return false
	}
	return true
}

func (m *DefaultAuthModule) IsSessionAuthed(ctx context.Context, s string) bool {
	id, err := m.repo.GetSessionUUID(ctx, s)
	no_uuid := id == ""
	if err != nil || no_uuid {
		return false
	}
	return true
}

func (m *DefaultAuthModule) DeauthSession(ctx context.Context, s string) error {
	invalid_session := !m.ValidateSession(ctx, s)
	if invalid_session {
		errmsg := "DefaultAuthModule: Attempted to deauthorize invalid session."
		return errors.New(errmsg)
	}

	m.repo.SetSessionUUID(ctx, s, "")

	return nil
}
