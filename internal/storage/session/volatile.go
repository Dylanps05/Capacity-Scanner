package session

import (
	"context"
	"errors"
	"sync"

	"github.com/Dylanps05/Capacity-Scanner/internal/cstype"
)

type SyncedSessionStorage struct {
	sessions *sync.Map
}

func NewSyncedSessionStorage() SessionStorage {
	return &SyncedSessionStorage{
		sessions: &sync.Map{},
	}
}

func (s *SyncedSessionStorage) StoreSession(ctx context.Context, session_id string) error {
	_, occupied := s.sessions.Load(session_id)
	if occupied {
		return errors.New("Session already exists.")
	}

	s.sessions.Store(session_id, cstype.UserID(""))

	return nil
}

func (s *SyncedSessionStorage) GetSessionUUID(ctx context.Context, token string) (cstype.UserID, error) {
	uuid_any, ok := s.sessions.Load(token)
	if !ok {
		err_msg := "SyncedSessionStorage: Session " + token + " is not active."
		err := errors.New(err_msg)
		return "", err
	}
	uuid := uuid_any.(cstype.UserID)
	return uuid, nil
}

func (s *SyncedSessionStorage) SetSessionUUID(ctx context.Context, token string, uuid cstype.UserID) error {
	_, ok := s.sessions.Load(token)
	if !ok {
		err_msg := "Session " + token + " is not active."
		err := errors.New(err_msg)
		return err
	}
	s.sessions.Store(token, uuid)
	return nil
}

func (s *SyncedSessionStorage) DeleteSession(ctx context.Context, token string) error {
	_, ok := s.sessions.LoadAndDelete(token)
	if !ok {
		err_msg := "Session " + token + " is not active."
		err := errors.New(err_msg)
		return err
	}
	return nil
}
