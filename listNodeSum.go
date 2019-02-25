package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := new(ListNode)
	//进位的变量
	carry := 0
	cur := result
	for l1 != nil || l2 != nil {
		//考虑位数不同状况下的赋值
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
		}
		if l2 != nil {
			v2 = l2.Val
		}
		//加上上一次的进位
		sum := carry + v1 + v2
		//算出你的进位
		carry = sum / 10
		//设置新节点
		newNode := new(ListNode)
		newNode.Val = sum % 10
		if carry > 0 {
			newNode.Val = carry
			cur.Next = newNode
			cur = cur.Next
			cur.Val = sum % 10
		}
		cur.Next = newNode
		cur = cur.Next
		//继续走下去
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry > 0 {
		newCode := new(ListNode)
		newCode.Val = carry
		cur.Next = newCode
	}
	return result.Next
}
