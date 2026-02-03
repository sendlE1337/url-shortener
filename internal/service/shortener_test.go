package shortener_test

import (
	shortener "SHORTNERED_URL/internal/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortener(t *testing.T) {
	t.Run("returns an alphanumberic short identifier", func(t *testing.T) {
		type testCase struct {
			id       uint32
			expected string
		}

		testCases := []testCase{
			{
				id:       1024,
				expected: "Ce",
			},
			{
				id:       0,
				expected: "",
			},
			{
				id:       3215,
				expected: "G8r",
			},
		}

		for _, tc := range testCases {
			actual := shortener.Shortener(tc.id)
			assert.Equal(t, tc.expected, actual)
		}
	})

	t.Run("is idempotent", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			assert.Equal(t, "Ce", shortener.Shortener(1024))
		}
	})

}
