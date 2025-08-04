package main

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)	
		AssertNotEqual(t, 1, 2)	
	})
	t.Run("asserting on string", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")	
		AssertNotEqual(t, "hello", "grace")	
	})
	// AssertEqual(t, "1", 1)	
}

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		myStackOfInts := NewStack[int]()

		// check stack is empty
		AssertTrue(t, myStackOfInts.IsEmpty())
		
		// add a thing and check it's not empty
		myStackOfInts.Push(123)
		AssertFalse(t, myStackOfInts.IsEmpty())

		// add another thing, pop it back again
		myStackOfInts.Push(456)
		value, _ := myStackOfInts.Pop()
		AssertEqual(t, value, 456)
		value, _ = myStackOfInts.Pop()
		AssertEqual(t, value, 123)
		AssertTrue(t, myStackOfInts.IsEmpty())
	})
}

func AssertEqual(t *testing.T, got, want any) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v want %+v", got, want)
	}
}

func AssertNotEqual(t *testing.T, got, want any) {
	t.Helper()
	if got == want {
		t.Errorf("got %+v want %+v", got, want)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}
