package unsafe

import (
	"testing"
)

type User struct {
	// 0
	Name string
	// 16
	Age int32
	// 20
	Age2 int32
	// 24
	Address []string
}

func TestPrintFieldOffset(t *testing.T) {
	testCases := []struct {
		name   string
		entity any
	}{
		{
			name:   "user",
			entity: User{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			PrintFieldOffset(tc.entity)
		})
	}

}
