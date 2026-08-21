package main

import "testing"

func testMain(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{},
	}

	for _, c := range cases {
		t.Errorf("%v", c)
		continue
	}
}
