package singlinklist

//func isPalindrome(list *LinkList) bool {
//	if list.length == 0 {
//		return false
//	}
//	if list.length == 1 {
//		return true
//	}
//
//	cur := list.head
//	s := make([]string, 0, list.length/2)
//	for i := 1; i <= list.length; i++ {
//		cur = cur.next
//		if list.length%2 != 0 && i == list.length/2+1 {
//			continue
//		}
//		if i <= list.length/2 {
//			s = append(s, cur.GetValue().String())
//		} else {
//			if s[list.length-i] != cur.GetValue().String() {
//				return false
//			}
//		}
//	}
//
//	return true
//}
//
//func isPalindrome2(l *LinkList) bool {
//	iLen = l.length
//	if iLen == 0 {
//		return false
//	}
//	if iLen == 1 {
//		return true
//	}
//	cur := l.head.next
//	next := l.head.next.next
//	pre := NewLinkNode(0)
//	pre.next = nil
//
//	var tmp *ListNode
//
//	for i := 1; i <= iLen/2; i++ {
//		tmp = cur.next
//		cur.next = pre
//		pre = cur
//		cur = tmp
//	}
//	mid := cur
//
//	var left, right *ListNode = pre, nil
//	if iLen%2 != 0 {
//		right = cur.next
//	} else {
//		right = cur
//	}
//
//	for left != nil && right != nil {
//		if left.GetValue().String() != right.GetValue().String() {
//			return false
//		}
//		left = left.next
//		right = right.next
//	}
//
//	cur = pre
//	pre = mid
//	for cur != nil {
//		tmp = cur.next
//		cur.next = pre
//		pre = cur
//		cur = tmp
//	}
//	l.head.next = pre
//
//	return true
//
//}
