package listutil

import "testing"

func TestArrToList(t *testing.T) {
	var tests = []struct {
		input []int
		want  *ListNode
	}{
		{[]int{5, 6, 7}, &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7, Next: nil}}}},
		{[]int{}, nil},
		{[]int{5}, &ListNode{Val: 5, Next: nil}},
	}
	for _, test := range tests {
		if got := ArrToList(test.input); !ListsEqual(got, test.want) {
			t.Errorf("arrToList(%v) = %v - want %v", test.input, ListToArr(got), ListToArr(test.want))
		}
	}
}

func TestListToArr(t *testing.T) {
	var tests = []struct {
		input *ListNode
		want  []int
	}{
		{&ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7, Next: nil}}}, []int{5, 6, 7}},
		{nil, []int{}},
		{&ListNode{Val: 5, Next: nil}, []int{5}},
	}
	for _, test := range tests {
		if got := ListToArr(test.input); !ArrsEqual(got, test.want) {
			t.Errorf("arrToList(%v) = %v - want %v", ListToArr(test.input), got, test.want)
		}
	}
}

func ArrsEqual(l1, l2 []int) bool {
	if len(l1) != len(l2) {
		return false
	}
	for i := range l1 {
		if l1[i] != l2[i] {
			return false
		}
	}
	return true
}

func ListsEqual(l1, l2 *ListNode) bool {
	return ArrsEqual(ListToArr(l1), ListToArr(l2))
}
