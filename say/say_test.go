package say

import (
	"strings"
	"testing"
)

func TestSay(t *testing.T) {
	message := "Hello, World!"
	got, err := Say(message)
	if err != nil {
		t.Fatalf("Say(%q) returned an error: %v", message, err)
	}

	if !strings.Contains(got, "< "+message+" >") {
		t.Errorf("Say(%q) = %q, want output containing the message in a speech bubble", message, got)
	}
	if !strings.Contains(got, "(oo)") {
		t.Errorf("Say(%q) = %q, want output containing the default cow", message, got)
	}
}
