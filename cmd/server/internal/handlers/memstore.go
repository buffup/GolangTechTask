package handlers

import (
	"errors"
	"sync"
)

type inMemStore struct {
	mu        *sync.RWMutex
	idCounter uint64
	buffs     map[uint64]*Buff
}

func (i *inMemStore) GetBuff(id uint64) (*Buff, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()

	b, ok := i.buffs[id]
	if !ok {
		return nil, errors.New("buff not found")
	}
	return b, nil
}

func (i *inMemStore) SetBuff(b *Buff) (uint64, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.idCounter++
	b.ID = i.idCounter
	i.buffs[i.idCounter] = b
	return i.idCounter, nil
}

func NewInMemStore() Store {
	return &inMemStore{
		mu:        &sync.RWMutex{},
		idCounter: 0,
		buffs:     make(map[uint64]*Buff),
	}
}
