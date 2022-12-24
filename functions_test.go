package nimble

import "testing"

func TestUnique(t *testing.T) {
	a := Unique("hello", "world", "banana", "world", "hello")
	if len(a) != 3 {
		t.Errorf("Expected 3 elements, got %d\n", len(a))
		return
	}
}

func TestFilter(t *testing.T) {
	results := Filter(func(v int) bool {
		return v%2 == 0
	}, 1, 2, 3, 4, 5, 6, 7)
	if len(results) != 3 {
		t.Errorf("Expected 3 elements, got %d\n", len(results))
		return
	}
}

func TestMap(t *testing.T) {
	results := Map(func(v int) string {
		switch v {
		case 1:
			return "one"
		case 2:
			return "two"
		case 3:
			return "three"
		default:
			return "default"
		}
	}, 1, 2, 3, 7, 1)
	if len(results) != 5 {
		t.Errorf("Expected 3 elements, got %d\n", len(results))
		return
	}

	if results[0] != "one" {
		t.Errorf("Got an incorrect result on the first element")
	}

	if results[3] != "default" {
		t.Errorf("Got an incorrect result on the fourth element")
	}
}
