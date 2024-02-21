package common

import (
	"bytes"
	"encoding/binary"
	"math/rand"

	"github.com/hitoshi44/go-uid64"
)

type Id struct {
	Client uint64
	Clock  uint64
}

func NewId(client, clock uint64) *Id {
	return &Id{
		Client: client,
		Clock:  clock,
	}
}

func GenerateId() (*Id, error) {
	g, err := uid64.NewGenerator(rand.Intn(4))
	if err != nil {
		return nil, err
	}
	id, err := g.Gen()
	if err != nil {
		return nil, err
	}
	return &Id{
		Client: uint64(id),
		Clock:  0,
	}, nil
}

// was compareIDs in yjs
func (i *Id) Equal(other *Id) bool {
	return i == other || (i.Client == other.Client && i.Clock == other.Clock)
}

// Write appends encoded ID to the buffer
func (i *Id) Write(buffer *bytes.Buffer) error {
	temp := make([]byte, 8)
	bno := binary.PutUvarint(temp, i.Client)
	_, err := buffer.Write(temp[:bno])
	if err != nil {
		return err
	}
	bno = binary.PutUvarint(temp, i.Clock)
	_, err = buffer.Write(temp[:bno])
	if err != nil {
		return err
	}
	return nil
}

// Read reads ID from buffer
func (i *Id) Read(buffer *bytes.Buffer) error {
	value, err := binary.ReadUvarint(buffer)
	if err != nil {
		return err
	}
	i.Client = value
	value, err = binary.ReadUvarint(buffer)
	if err != nil {
		return err
	}
	i.Clock = value
	return nil
}
