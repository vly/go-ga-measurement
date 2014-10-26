package measurement

import (
	"testing"
)

// Testing struct to url string encoding
func TestFlatten(t *testing.T) {
	message := PageviewMessage{
		"pageview",
		"test.com",
		"/test.html",
		"Testing page",
		1}
	expected := "dh=test.com&dp=%2Ftest.html&dt=Testing+page&linkid=1&t=pageview"
	out := Flatten(&message)
	if expected != out.Encode() {
		t.Errorf("Failed to Flatten struct: expected (%s), got (%s)\n", expected, out.Encode())
	}
}
