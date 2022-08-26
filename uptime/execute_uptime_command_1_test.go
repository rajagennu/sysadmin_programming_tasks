package uptime

import (
	"strings"
	"testing"
)

func TestUptime(t *testing.T) {
	expectedShouldHave := "load average"
	result := Uptime()
	if !strings.Contains(result, expectedShouldHave) {
		t.Errorf("Expecting result : '%s' to have substring '%s' ", result, expectedShouldHave)
	}

}
