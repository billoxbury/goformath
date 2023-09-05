package main

// method to insert a new item into a tree
func (n *BinaryNode) insert(data itemType) {

	if n == nil {
		return
	} else if data == n.data {
		n.count++
	} else if data < n.data {
		if n.left == nil {
			n.left = &BinaryNode{data: data, left: nil, right: nil, count: 1}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryNode{data: data, left: nil, right: nil, count: 1}
		} else {
			n.right.insert(data)
		}
	}
}

func (t *BinaryTree) insert(data itemType) *BinaryTree {
	if t.root == nil {
		t.root = &BinaryNode{data: data, left: nil, right: nil, count: 1}
	} else {
		t.root.insert(data)
	}
	return t
}
