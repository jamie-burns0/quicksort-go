package stack

import(
	"testing"
)

func TestStack( t *testing.T ) {
	t.Run("Test 1", func(t *testing.T) {
		
		var s Stack[int]

		s.Push(1)
		s.Push(2)
		s.Push(3)

		// pop 3
		if item, ok := s.Pop() ; !ok {
			t.Errorf( "Expected to pop item 3 off the stack. Item was not present")
		} else if item != 3 {
			t.Errorf( "Expected to pop 3 off the stack. Got %v", item )
		}

		// pop 2
		if item, ok := s.Pop() ; !ok {
			t.Errorf( "Expected to pop item 2 off the stack. Item was not present")
		} else if item != 2 {
			t.Errorf( "Expected to pop 2 off the stack. Got %v", item )
		}

		s.Push(4)
		
		// pop 4
		if item, ok := s.Pop() ; !ok {
			t.Errorf( "Expected to pop item 4 off the stack. Item was not present")
		} else if item != 4 {
			t.Errorf( "Expected to pop 4 off the stack. Got %v", item )
		}

		// pop 1
		if item, ok := s.Pop() ; !ok {
			t.Errorf( "Expected to pop item 1 off the stack. Item was not present")
		} else if item != 1 {
			t.Errorf( "Expected to pop 1 off the stack. Got %v", item )
		}

		if _, ok := s.Pop() ; ok {
			t.Errorf( "Expected to pop no item off the stack. Stack was not empty")
		}
	})
}