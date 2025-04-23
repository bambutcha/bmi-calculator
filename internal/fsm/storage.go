package fsm

import (
    "sync"
)

type StateStorage struct {
    users map[int64]*UserState
    mu    sync.RWMutex
}

func NewStateStorage() *StateStorage {
    return &StateStorage{
        users: make(map[int64]*UserState),
    }
}

func (s *StateStorage) GetState(userID int64) *UserState {
    s.mu.RLock()
    defer s.mu.RUnlock()
    return s.users[userID]
}

func (s *StateStorage) SetState(userID int64, state *UserState) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.users[userID] = state
}

func (s *StateStorage) ClearState(userID int64) {
    s.mu.Lock()
    defer s.mu.Unlock()
    delete(s.users, userID)
}