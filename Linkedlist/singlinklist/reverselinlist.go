package singlinklist

func ReverseLinkList(l *LinkList) {
	iLen := l.length
	if iLen == 0 {
		return
	}
	if iLen == 1 {
		return
	}

	var tmp *LinkNode
	pre := l.head.next
	cur := l.head.next.next
	pre.next = nil
	for cur != nil {
		tmp = cur.next
		cur.next = pre
		pre = cur
		cur = tmp
	}
	l.head.next = pre
}
