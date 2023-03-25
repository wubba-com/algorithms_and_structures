package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}
//l2:= &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return nil
	}

	if l2 == nil {
		return nil
	}
	var (
		first   = &ListNode{}
		current = first
		p       = l1
		q       = l2
		carry   int
	)

	for p != nil || q != nil {
		var x, y = 0, 0
		if p != nil {
			x = p.Val
		}
		if q != nil {
			y = q.Val
		}
		var sum = carry + x + y
		carry = sum / 10
		current.Next = &ListNode{
			Val: sum % 10,
		}
		current = current.Next

		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}
	if carry > 0 {
		current.Next = &ListNode{
			Val: carry,
		}
	}

	return first.Next
}

// Сложность O(3n)
func lengthOfLongestSubstringV1(s string) int {
	bs := []byte(s)
	var check = func(start, end int) bool {
		chars := make(map[uint8]struct{})

		for i := start; i < end+1; i++ {
			c := bs[i]

			if _, ok := chars[c]; ok {
				return false
			}
			chars[c] = struct{}{}
		}

		return true
	}
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(bs)
	res := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if check(i, j) {
				res = max(res, j-i+1)
			}
		}
	}

	return res
}

// Сложность O(n)
func lengthOfLongestSubstringV2(s string) int {
	var (
		m   = make(map[uint8]int)
		ans = 0
		bs  = []byte(s)
		n   = len(bs)
		i   = 0
	)

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for j := 0; j < n; j++ {
		if v, ok := m[bs[j]]; ok {
			i = max(v, i)
		}

		ans = max(ans, j-i+1)
		m[bs[j]] = j + 1
	}

	return ans
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	digits := make([]int, 0)
	for x != 0 {
		digits = append(digits, x%10)
		x /= 10
	}
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		if digits[i] != digits[j] {
			return false
		}
	}

	return true
}

func longestCommonPrefix(strs []string) string {
	var prefix string
	if len(strs) == 0 {
		return prefix
	}
	if len(strs) < 2 {
		prefix = strs[0]
		return prefix
	}
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})
	var check = func(i int) bool {
		p := strs[0][:i+1]
		for j := range strs {
			if strs[j][:i+1] != p {
				return false
			}
		}

		return true
	}
	max := len(strs[0])
	i := 0
	for i < max {
		if check(i) {
			prefix = strs[0][:i+1]
		}
		i++
	}

	return prefix
}

func NewStack() *stack {
	return &stack{
		list: make([]string, 0),
	}
}

type stack struct {
	list []string
}

func (s *stack) push(char string) {
	s.list = append(s.list, char)
}

func (s *stack) pop() (string, bool) {
	if n := len(s.list); n > 0 {
		v := s.list[n-1]
		s.list = s.list[:n-1]
		return v, true
	}

	return "", false
}

func (s *stack) top() (string, bool) {
	if len(s.list) != 0 {
		return s.list[len(s.list)-1], true
	}
	return "", false
}

func (s *stack) len() int { return len(s.list) }

func isValid(s string) bool {
	st := NewStack()
	for i := range s {
		if string(s[i]) == "(" {
			st.push(")")
		} else if string(s[i]) == "[" {
			st.push("]")
		} else if string(s[i]) == "{" {
			st.push("}")
		} else {
			if st.len() == 0 {
				return false
			}
			if v, ok := st.top(); ok {
				if v != string(s[i]) {
					return false
				}
				st.pop()
			}
		}
	}

	return st.len() == 0
}

// ---------- start mergeTwoLists ----------
// mergeTwoLists сортировка связного списка
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	node := head
	for list1 != nil || list2 != nil {
		if list1 != nil {
			node.Next = list1
			list1 = list1.Next
			node = node.Next
		}

		if list2 != nil {
			node.Next = list2
			list2 = list2.Next
			node = node.Next
		}
	}

	return sortList(head.Next)
}

func merge(l1, l2 *ListNode) *ListNode {
	var (
		head = &ListNode{}
		node = head
	)

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}
		node = node.Next
	}

	if l1 != nil {
		node.Next = l1
	}
	if l2 != nil {
		node.Next = l2
	}
	return head.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var (
		prev       *ListNode
		slow, fast = head, head
	)
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil

	f := sortList(head)
	b := sortList(slow)

	return merge(f, b)
}

func printList(ll *ListNode) {
	next := ll
	for next != nil {
		fmt.Print(next.Val)
		next = next.Next
	}
	fmt.Println()
}

// ListNode Definition for singly-linked list.

// ---------- end mergeTwoLists ----------

func removeDuplicates(nums []int) int {
	for i := range nums {
		j := i + 1
		for j < len(nums) {
			if nums[i] == nums[j] {
				copy(nums[i:], nums[i+1:])
				nums = nums[:len(nums)-1]
				j = i

			}
			j++
		}
	}
	return len(nums)
}

// remove element

func removeElementV2(nums []int, val int) int {
	sort.Ints(nums)

	var i int
	for i != -1 {
		i = binary(nums, val)
		if i == -1 {
			break
		}
		copy(nums[i:], nums[i+1:])
		nums = nums[:len(nums)-1]
	}
	return len(nums)
}

// Сложность O(n)
func removeElement(nums []int, val int) int {
	var i = 0
	for i < len(nums) {
		if nums[i] == val {
			nums[i] = nums[len(nums)-1]
			nums = nums[:len(nums)-1]
			i = 0
		} else {
			i++
		}
	}
	return len(nums)
}

// Сложность O(logN(n))
func binary(sl []int, v int) int {
	start := 0
	end := len(sl) - 1
	middle := 0
	for start <= end {
		middle = (start + end) / 2
		if sl[middle] == v {
			return middle
		} else if sl[middle] < v {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}

	return -1
}

// Сложность O(logN(n))
func searchInsert(nums []int, target int) int {
	sort.Ints(nums)
	if v := binary(nums, target); v != -1 {
		return v
	} else {
		nums = append(nums, target)
		sort.Ints(nums)
		return binary(nums, target)
	}
}

// Сложность O(1)
func lengthOfLastWord(s string) int {
	flow := strings.Fields(s)

	return len([]byte(flow[len(flow)-1]))
}

// Сложность O(n)
func plusOne(digits []int) []int {
	if digits[len(digits)-1] == 9 {
		var (
			carry     int
			i         = len(digits) - 1
			newDigits = make([]int, len(digits))
		)

		for i >= 0 {
			var sum int
			if i == len(digits)-1 {
				sum = carry + digits[i] + 1
			} else {
				sum = carry + digits[i]
			}

			carry = sum / 10
			newDigits[i] = sum % 10
			i--
		}
		if carry != 0 {
			newDigits = append(newDigits, carry)
			newDigits[0], newDigits[len(newDigits)-1] = newDigits[len(newDigits)-1], newDigits[0]
		}

		return newDigits
	} else {
		digits[len(digits)-1]++

		return digits
	}
}

func addBinary(a string, b string) string {

	i := len(a) - 1
	j := len(b) - 1

	var output string
	var sum int
	carry := 0
	for i >= 0 && j >= 0 {
		first := int(a[i] - '0')
		second := int(b[j] - '0')

		sum, carry = binarySum(first, second, carry)

		output = strconv.Itoa(sum) + output
		i--
		j--
	}

	for i >= 0 {
		first := int(a[i] - '0')

		sum, carry = binarySum(first, 0, carry)

		output = strconv.Itoa(sum) + output
		i = i - 1

	}

	for j >= 0 {
		second := int(b[j] - '0')

		sum, carry = binarySum(0, second, carry)

		output = strconv.Itoa(sum) + output
		j = j - 1
	}

	if carry > 0 {
		output = strconv.Itoa(1) + output
	}

	return output
}

func binarySum(a, b, carry int) (int, int) {
	output := a + b + carry

	if output == 0 {
		return 0, 0
	}

	if output == 1 {
		return 1, 0
	}

	if output == 2 {
		return 0, 1
	}

	if output == 3 {
		return 1, 1
	}

	return 0, 0
}

func romanToInt(s string) int {
	m := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	var sum int
	for i, v := range s {
		if value, ok := m[string(v)]; ok {
			var next uint8
			if len(s)-1 != i {
				next = s[i+1]
			} else {
				next = s[i]
			}

			nextVal := m[string(next)]
			if value < nextVal {
				sum -= value
			} else {
				sum += value
			}

		}
	}

	return sum
}

func Download(sizes ...int) <-chan int {
	var c = make(chan int, 2)
	go func() {
		wg := sync.WaitGroup{}
		for _, size := range sizes {
			wg.Add(1)
			go func(size int) {
				defer wg.Done()

				var sum int
				for i := 0; i <= size; i++ {
					sum += i
				}
				c <- sum

			}(size)
		}
		wg.Wait()
		close(c)
	}()
	return c
}

func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	n := &ListNode{}
	next := n
	for current != nil {
		if !In(current.Val, n.Next) {
			next.Next = &ListNode{Val: current.Val}
			next = next.Next
		}
		current = current.Next
	}

	return n.Next
}

func In(val int, other *ListNode) bool {
	for other != nil {
		if other.Val == val {
			return true
		}
		other = other.Next
	}

	return false
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else {
		if (p == nil && q != nil) || (q == nil && p != nil) {
			return false
		}
		if p.Val == q.Val {
			isLeft := isSameTree(p.Left, q.Left)
			isRight := isSameTree(p.Right, q.Right)

			if isLeft && isRight {
				return true
			}
		}
	}

	return false
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	nums := make([]int, 0)
	nums = append(nums, inorderTraversal(root.Left)...)
	nums = append(nums, root.Val)
	nums = append(nums, inorderTraversal(root.Right)...)

	return nums
}

func isSymmetrics(x, y *TreeNode) bool {
	if x == nil && y == nil {
		return true
	}

	return (x != nil && y != nil) && isSymmetrics(x.Left, y.Right) && isSymmetrics(x.Right, y.Left)
}
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetrics(root.Left, root.Right)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	return int(float64(1) + math.Max(float64(leftDepth), float64(rightDepth)))
}
