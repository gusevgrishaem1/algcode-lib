package algcode_lib

import "testing"

func TestStackPush(t *testing.T) {
	// init s
	s := NewStack(2)

	// push 1 --> s.Size() == 1
	s.Push(1)
	if s.Size() != 1 {
		t.Errorf("stack Size is wrong. Expected 1, got %d", s.Size())
	}

	// push 2 --> s.Size() == 2
	s.Push(2)
	if s.Size() != 2 {
		t.Errorf("stack Size is wrong. Expected 2, got %d", s.Size())
	}
}

func TestStackPop(t *testing.T) {
	// init s
	s := NewStack(4)

	// push stack
	s.Push(5)
	s.Push(2)
	s.Push(1)

	// fill array
	arr := make([]int, 0, 3)
	for s.Size() != 0 {
		el, ok := s.Pop().(int)
		if !ok {
			t.Error("Error cast element to int")
		}
		arr = append(arr, el)
	}

	// check len of the array
	if len(arr) != 3 {
		t.Error("len arr != 3")
	}

	// check stack size
	if s.Size() != 0 {
		t.Error("stack size != 0")
	}

	// check elements in array
	expected := []int{1, 2, 5}
	for i, v := range expected {
		if arr[i] != v {
			t.Error("wrong element in array")
		}
	}

	// push 4 to stack
	s.Push(4)

	// check size of the stack
	if s.Size() != 1 {
		t.Error("size stack != 1")
	}

	// pop el
	el, ok := s.Pop().(int)
	if !ok {
		t.Error("error cast el to int")
	}

	// check el == 4
	if el != 4 {
		t.Error("el != 4")
	}

	// check size of the stack
	if s.Size() != 0 {
		t.Error("size stack != 0")
	}
}
