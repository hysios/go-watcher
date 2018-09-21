package watcher

import (
	"testing"
)

// func TestPrepareArgs(t *testing.T) {
// 	args := []string{"watcher", "-run", "balcony", "-p", "11880", "--watch", "show", "--host", "localhost"}

// 	params := ParseArgs(args)
// 	if len(params.Package) != 4 {
// 		t.Fatalf("Expected 2 parameters with their values in Package parameters but got %d", len(params.Package))
// 	}

// 	if params.Package[0] != "-p" {
// 		t.Errorf("Expected -p as package parameter but got %s", params.Package[0])
// 	}

// 	if params.Package[2] != "--host" {
// 		t.Errorf("Expected --host as package parameter but got %s", params.Package[0])
// 	}

// 	if len(params.Watcher) != 2 {
// 		t.Fatalf("Expected 2 parameter with their values in System parameters but got %d", len(params.Watcher))
// 	}

// 	if params.Watcher["run"] != "balcony" {
// 		t.Errorf("Expected balcony but got %s", params.Watcher["run"])
// 	}

// 	// TODO check this fatal error case
// 	// args = []string{"watcher", "-run", "balcony", "-p", "11880", "--watch"}
// 	// params = PrepareArgs(args)

// }

func TestExistIn(t *testing.T) {
	input := []string{"a", "b", "c"}

	tests := []struct {
		name   string
		search string
		input  []string
		want   bool
	}{
		{"c", "c", input, true},
		{"f", "f", input, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := existIn(tt.search, tt.input)
			if got != tt.want {
				t.Errorf("existIn() = %v, want %v", got, tt.want)
			}
		})
	}
}
