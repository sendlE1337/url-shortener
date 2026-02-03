package shortener_test

import (
	"testing"

	shortener "SHORTNERED_URL/internal/service"
	"SHORTNERED_URL/internal/storage"

	"github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {
	mem := storage.NewInMemory()
	s := shortener.NewService(mem)

	t.Run("shorten creates a short URL", func(t *testing.T) {
		url := "https://youtube.com"
		short, err := s.Shorten(url)
		assert.NoError(t, err)
		assert.Equal(t, url, short.OriginalURL)
		assert.NotEmpty(t, short.Identifier)
	})

	t.Run("get returns existing shortening", func(t *testing.T) {
		url := "https://example.com"
		short, _ := s.Shorten(url)

		got, err := s.Get(short.Identifier)
		assert.NoError(t, err)
		assert.NotNil(t, got)
		assert.Equal(t, url, got.OriginalURL)
	})

	t.Run("redirect returns original URL and increments visits", func(t *testing.T) {
		url := "https://golang.org"
		short, _ := s.Shorten(url)

		orig, err := s.Redirect(short.Identifier)
		assert.NoError(t, err)
		assert.Equal(t, url, orig)

		got, _ := s.Get(short.Identifier)
		assert.Equal(t, 1, got.Visits)
	})

	t.Run("get non-existing returns error", func(t *testing.T) {
		got, err := s.Get("nonexistent")
		assert.Error(t, err)
		assert.Nil(t, got)
	})

	t.Run("redirect non-existing returns error", func(t *testing.T) {
		orig, err := s.Redirect("nonexistent")
		assert.Error(t, err)
		assert.Empty(t, orig)
	})
}
