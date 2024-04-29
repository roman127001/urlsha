package shorter

import (
	"testing"
)

func TestLen(t *testing.T) {
	testCases := []struct {
		name string
		want int
	}{
		{
			name: "Check length of generated string",
			want: 6,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sh := New()
			if got := sh.Generate(); len(got) != tc.want {
				t.Errorf("Shorter.Generate() = %v, want %v", got, tc.want)
			}
		})
	}
}
