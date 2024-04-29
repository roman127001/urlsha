package inruntime

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExpiration(t *testing.T) {
	testCases := []struct {
		name    string
		key     string
		val     string
		expired bool
	}{
		{
			name:    "Check expiration of key",
			key:     "foo",
			val:     "bar",
			expired: false, // always false - implement Expire method (related to `test-first` approach)!
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Prepare.
			cl := New()
			cl.Set(tc.key, tc.val)

			// Test.
			expired := cl.Expire(tc.key)
			assert.Equal(t, expired, tc.expired, "Expire() = %v, want %v", expired, tc.expired)

			v, _ := cl.Get(tc.key)
			assert.Equal(t, v, tc.val, "Get() = %v, want %v", v, tc.val)
		})
	}
}
