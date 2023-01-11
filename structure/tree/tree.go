package tree

import (
	"errors"
)

func New(root int) *BinaryTree {
	return &BinaryTree{
		val: root,
	}
}

type (
	BinaryTree struct {
		val   int
		left  *BinaryTree
		right *BinaryTree
	}
)

// Add Добавляет элемент в дерево на корректную позицию.
// Сложность: O(log n) в среднем; O(n) в худшем случае.
func (t *BinaryTree) Add(value int) error {
	if t == nil {
		return errors.New("tree == nil")
	}

	if t.val == value {
		return errors.New("this is value already exists")
	}

	if value > t.val {
		if t.right == nil {
			newTree := &BinaryTree{val: value}
			t.right = newTree

			return nil
		}

		return t.right.Add(value)
	}

	if value < t.val {
		if t.left == nil {
			newTree := &BinaryTree{val: value}
			t.left = newTree

			return nil
		}

		return t.left.Add(value)
	}

	return nil
}

// Inorder Инфиксный обход используется тогда,
//когда нам надо обойти дерево в порядке,
// соответствующем значениям узлов
func (t *BinaryTree) Inorder(callback func(value int)) {
	if t == nil {
		return
	}

	t.left.Inorder(callback)
	callback(t.val)
	t.right.Inorder(callback)
}

// Preorder При префиксном обходе алгоритм
// получает значение текущего узла перед тем,
// как перейти сначала в левое поддерево, а затем в правое
func (t *BinaryTree) Preorder(callback func(value int)) {
	if t == nil {
		return
	}

	callback(t.val)
	t.left.Preorder(callback)
	t.right.Preorder(callback)
}

// Postorder При постфиксном обходе мы посещаем
// левое поддерево, правое поддерево, а потом,
// после обхода всех детей, переходим к самому узлу.
func (t *BinaryTree) Postorder(callback func(value int)) {
	if t == nil {
		return
	}

	t.left.Postorder(callback)
	t.right.Postorder(callback)
	callback(t.val)
}

func (t *BinaryTree) Find(val int) bool {

	if t == nil {
		return false
	}

	switch {
	case val == t.val:
		return true
	case val > t.val:
		return t.right.Find(val)
	default:
		return t.left.Find(val)
	}
}

func (t *BinaryTree) Delete(val int) {
	t.remove(val)
}

// Идея удаления элемента делится на несколько случаев:
//
//    1. у узла нет дочерних узлов;
//    2. у узла есть левый дочерний узел;
//    3. у узла есть правый дочерний узел;
//    4. у узла есть оба ребёнка.
//
// В случае 1 просто удаляем узел, дополнительная работа не требуется.
// В случае 2 и 3 заменяем удаляемый узел на его
// потомка, на этом удаление заканчивается. Случай
// 4 самый сложный, но концепция такого удаления
// довольно проста — найти в правом поддереве
// минимальный элемент и переместить его на место удаляемого узла.
func (t *BinaryTree) remove(value int) *BinaryTree {
	if t == nil {
		return nil
	}

	if value < t.val {
		t.left = t.left.remove(value)
		return t
	}

	if value > t.val {
		t.right = t.right.remove(value)
		return t
	}

	if t.left == nil && t.right == nil {
		t = nil
		return nil
	}

	if t.left == nil {
		t = t.right
		return t
	}

	if t.right == nil {
		t = t.left
		return t
	}

	smallestValOnRight := t.right
	for {
		if smallestValOnRight != nil && smallestValOnRight.left != nil {
			smallestValOnRight = smallestValOnRight.left
		} else {
			break
		}
	}

	t.val = smallestValOnRight.val
	t.right = t.right.remove(t.val)

	return t
}
