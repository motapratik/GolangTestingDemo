package function

import "testing"

func TestAddition(t *testing.T) {
	got := Addition(2, 2)
	expected := 4

	if got != expected {
		t.Errorf("Error in test, got: %v and expected %v", got, expected)
	}
}

func TestGreeting(t *testing.T) {
	got := Greeting("Gopher")
	expected := "Hello Gopher!"

	if got != expected {
		t.Errorf("Error in test, got: %q and expected %q", got, expected)
	}
}

func TestGreetingTableDriven(t *testing.T) {
	scenarios := []struct {
		input    string
		expected string
	}{
		{input: "Gopher", expected: "Hello Gopher!"},
		{input: "Buddy", expected: "Hello Buddy!"},
	}

	for _, s := range scenarios {
		got := Greeting(s.input)
		if got != s.expected {
			t.Errorf("Error in test, got: %q and expected %q", got, s.expected)
		}
	}
}

func TestDisplayMsg(t *testing.T) {
	got := DisplayMsg("Gopher")
	expected := "Msg: Gopher"

	if got != expected {
		t.Errorf("Error in test, got: %q and expected %q", got, expected)
	}
}
