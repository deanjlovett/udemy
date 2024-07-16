package timeparse

import "testing"

// * Write some unit tests to check your work
//   - Run tests with `go test ./exercise/errors`
func TestTimeParse(t *testing.T) {
	table := []struct {
		time string
		ok   bool
	}{
		{"0:0:0", false},
		{"0:00:01", true},
		{"1:20:30", true},
		{"25:20:30", false},
		{"1:60:30", false},
		{"1:20:60", false},
		{"-1:20:59", false},
		{"19:00:12", true},
		{"1:3:44", true},
		{"bad", false},
		{"1:-3:44", false},
		{"0:59:59", true},
		{"", false},
		{"11:22", false},
		{"aa:bb:cc", false},
		{"5:23:", false},
	}

	for _, data := range table {
		_, err := TimeParse(data.time)
		if data.ok && err != nil {
			t.Errorf("%v: %v, error should be nil", data.time, err)
		}
	}
}
