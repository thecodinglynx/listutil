package listutil

type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	Head *ListNode
}

func ListToArr(ln *ListNode) []int {
	res := []int{}
	ptr := ln
	for ptr != nil {
		res = append(res, ptr.Val)
		ptr = ptr.Next
	}

	return res
}

func ArrToList(input []int) *ListNode {
	var res List
	for _, v := range input {
		res.Insert(v)
	}
	return res.head
}

func (l *List) Insert(val int) *List {
	entry := &ListNode{Val: val, Next: nil}
	if l.head == nil {
		l.head = entry
	} else {
		p := l.head
		for p.Next != nil {
			p = p.Next
		}
		p.Next = entry
	}
	return l
}
