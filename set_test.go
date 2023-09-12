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

func TestFromMap(t *testing.T) {
	x := map[string]int{
		"hello": 1,
		"world": 2,
	}

	s := NewSetFromMap(x)
	if !s.Has("hello") {
		t.Errorf("Set is missing an expected value: \"hello\"")
		return
	}
	if !s.Has("world") {
		t.Errorf("Set is missing an expected value: \"world\"")
		return
	}
}

func TestUnion(t *testing.T) {
	a := NewSet("a", "b", "c")
	b := NewSet("c", "d", "e")
	c := Union(a, b)
	if c.Size() != 5 {
		t.Errorf("Set size is incorrect: %v", c.Items())
		return
	}
}

func TestIntersect(t *testing.T) {
	a := NewSet("a", "b", "c")
	b := NewSet("c", "d", "e")
	c := Intersect(a, b)
	if c.Size() != 1 {
		t.Errorf("Set size is incorrect: %v", c.Items())
		return
	}
}

func TestDifference(t *testing.T) {
	a := NewSet("a", "b", "c")
	b := NewSet("c", "d", "e")
	c := a.Difference(b)
	if c.Size() != 2 {
		t.Errorf("Set size is incorrect: %v", c.Items())
		return
	}
	if !c.Has("a") {
		t.Errorf("missing set element: a")
		return
	}
	if !c.Has("b") {
		t.Errorf("missing set element: b")
		return
	}
}

func TestDifferenceEmpty(t *testing.T) {
	a := NewSet("a", "b", "c")
	b := NewSet[string]()
	c := a.Difference(b)
	if c.Size() != 3 {
		t.Errorf("Set size is incorrect: %v", c.Items())
		return
	}
	if !c.Has("a") {
		t.Errorf("missing set element: a")
		return
	}
	if !c.Has("b") {
		t.Errorf("missing set element: b")
		return
	}
	if !c.Has("c") {
		t.Errorf("missing set element: b")
		return
	}
}

func TestPopEmpty(t *testing.T) {
	s := NewSet[string]()
	val := s.Pop()
	if val != nil {
		t.Errorf("return value should be nil")
	}
}

func TestPopNonEmpty(t *testing.T) {
	s := NewSet[string]("a")
	val := s.Pop()
	if val != "a" {
		t.Errorf("return value should be nil")
	}
}

func TestCopy(t *testing.T) {
	s := NewSet[string]("a", "b", "c")
	u := s.Copy()
	if u.Size() != s.Size() {
		t.Errorf("Size of new set does not match size of old set")
		return
	}

	for _, item := range s.Items() {
		if !u.Has(item) {
			t.Errorf("new set is missing elements from old")
			return
		}
	}

}
