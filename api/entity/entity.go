package entity

type Entity interface {
	GetId() string
	SetId(id string)
}

type base struct {
	Id string
}

func (b *base) SetId(id string) {
	b.Id = id
}
