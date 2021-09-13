package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturns1KittenWhenSearchGarfield(t *testing.T) {
	store := MemoryStore{}
	kittens := store.Search("Garfield")

	assert.Equal(t, 1, len(kittens)) // Garfield 라는 data가 있으므로 len 값은 1이여야 한다.
}

func TestReturns0KittenWhenSearchTom(t *testing.T) {
	store := MemoryStore{}
	kittens := store.Search("Tom")

	assert.Equal(t, 0, len(kittens)) // TOM이라는 data가 없으므로 len 값은 0이여야 한다.
}
