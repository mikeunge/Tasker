package models

import "github.com/gofrs/uuid"

type Model interface {
	GetId() uuid.UUID
	SetId(id uuid.UUID)
}

type UID struct {
	Id uuid.UUID
}

func (m *UID) SetId(id uuid.UUID) {
	m.Id = id
}
