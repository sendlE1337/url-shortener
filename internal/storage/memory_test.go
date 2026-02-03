package storage_test

import (
	"testing"

	"SHORTNERED_URL/internal/model"
	storage "SHORTNERED_URL/internal/storage"

	"github.com/stretchr/testify/assert"
)

func TestInMemory(t *testing.T) {
	mem := storage.NewInMemory()

	short := model.Shortening{
		Identifier:  "abc123",
		OriginalURL: "https://youtube.com",
	}

	t.Run("put stores a shortening", func(t *testing.T) {
		got, err := mem.Put(short)
		assert.NoError(t, err)
		assert.NotNil(t, got)
	})

	t.Run("put duplicate returns error", func(t *testing.T) {
		got, err := mem.Put(short)
		assert.Error(t, err)
		assert.Nil(t, got)
	})

	t.Run("get returns stored shortening", func(t *testing.T) {
		got, err := mem.Get(short.Identifier)
		assert.NoError(t, err)
		assert.NotNil(t, got)
		assert.Equal(t, short.OriginalURL, got.OriginalURL)
	})

	t.Run("get non-existing returns error", func(t *testing.T) {
		got, err := mem.Get("nonexistent")
		assert.Error(t, err)
		assert.Nil(t, got)
	})

	t.Run("increment visits increases counter", func(t *testing.T) {
		err := mem.IncrementVisits(short.Identifier)
		assert.NoError(t, err)

		got, _ := mem.Get(short.Identifier)
		assert.Equal(t, 1, got.Visits)
	})

	t.Run("increment visits non-existing returns error", func(t *testing.T) {
		err := mem.IncrementVisits("nonexistent")
		assert.Error(t, err)
	})
}
