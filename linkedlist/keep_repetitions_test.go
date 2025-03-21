package linkedlist

import "testing"

/*
TestKeepRepetitions tests solution(s) with the following signature and problem description:

	func KeepRepetitions(head *Node) *Node

Given a linked list of sorted integers, create a copy of the list that contains one example of
each repeated item.

For example given 1->1->4->4->6->6->7, return 1->4->6 because
1,4,6 are items that are repeated in this list and 7 is not repeated.
*/
func TestKeepRepetitions(t *testing.T) {
	tests := []struct {
		list, solution string
	}{
		{"", ""},
		{"1", ""},
		{"1->1", "1"},
		{"1->4->6", ""},
		{"1->4->4->4->6", "4"},
		{"1->1->4->4->4->6", "1->4"},
		{"1->1->4->4->4->6->6", "1->4->6"},
		{"1->1->4->4->6->6->7", "1->4->6"},
		{"1->1->4->4->4->6->7->7", "1->4->7"},
	}

	for i, test := range tests {
		got := Serialize(KeepRepetitions(Deserialize(test.list)))
		if got != test.solution {
			t.Fatalf("Failed test case #%d. Want %s got %s", i, test.solution, got)
		}
	}
}
