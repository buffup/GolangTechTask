package handlers

type BuffStore interface {
	GetBuff(id uint64) (*Buff, error)
	SetBuff(*Buff) (id uint64, err error)
}

type Store interface {
	BuffStore
}
