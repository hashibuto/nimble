package nimble

import "testing"

func TestSetString(t *testing.T) {
	s := NewSet("hello", "world", "world", "dog")
	if s.Size() != 3 {
		t.Errorf("Expected 3 elements but got %d", s.Size())
		return
	}

	if !s.Has("world") {
		t.Errorf("Expected item to be present")
		return
	}

	if s.Has("cat") {
		t.Errorf("Item should not be present")
		return
	}

	s.Remove("dog")

	if s.Size() != 2 {
		t.Errorf("Wronge size after removal")
	}

	s.Remove("cat")
	if s.Size() != 2 {
		t.Errorf("Wronge size after noop removal")
	}
}

func TestSetInt(t *testing.T) {
	s := NewSet(1, 3, 3, 5)
	if s.Size() != 3 {
		t.Errorf("Expected 3 elements but got %d", s.Size())
		return
	}

	if !s.Has(3) {
		t.Errorf("Expected item to be present")
		return
	}

	if s.Has(4) {
		t.Errorf("Item should not be present")
		return
	}

	s.Remove(3)

	if s.Size() != 2 {
		t.Errorf("Wronge size after removal")
	}

	s.Remove(4)
	if s.Size() != 2 {
		t.Errorf("Wronge size after noop removal")
	}
}

func TestSetItems(t *testing.T) {
	s := NewSet(1, 2)
	items := s.Items()
	if len(items) != 2 {
		t.Errorf("Expected 2 items in set, got %d", len(items))
		return
	}
	if (items[0] == 1 && items[1] == 2) || (items[0] == 2 && items[1] == 1) {
	} else {
		t.Errorf("Incorrect items in set")
		return
	}
}
