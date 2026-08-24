package structs

import (
	"fmt"
	"testing"
	"time"
)

func TestAddG(t *testing.T) {
	const interval = 5 * time.Second

	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("Expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("Expected to find value")
				return
			}
		})
	}
}
