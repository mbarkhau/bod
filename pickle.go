package bod

import (
	"errors"
	"os"
)

type pickle struct {
	id   string
	size uint
	file *os.File
}

func newPickle(id string, size int) (*pickle, error) {
	file, err := os.OpenFile(id, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return nil, err
	}
	return &pickle{id, uint(size), file}, nil
}

func (p *pickle) getBytes(offset uint, length uint) ([]byte, error) {
	bytes := make([]byte, length)
	_, err := p.file.ReadAt(bytes, int64(offset))
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (p *pickle) setBytes(offset uint, bytes []byte) error {
	if offset >= p.size {
		return errors.New("Pos out of bound")
	}
	_, err := p.file.WriteAt(bytes, int64(offset))
	return err
}
