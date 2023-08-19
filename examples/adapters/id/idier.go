package id

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/oklog/ulid"
)

type Id interface {
	UUID() string
	ULID() string
}

type _ider struct {
	entropy *rand.Rand
}

func New() Id {
	return &_ider{rand.New(rand.NewSource(time.Now().UnixNano()))}
}

func (i *_ider) UUID() string {
	return uuid.NewString()
}

func (i *_ider) ULID() string {
	ms := ulid.Timestamp(time.Now())
	return ulid.MustNew(ms, i.entropy).String()
}
