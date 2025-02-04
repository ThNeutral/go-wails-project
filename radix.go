package main

type RadixNode struct {
	next *RadixNode
	char rune
}

func generateRadixTree(str string) *RadixNode {
	if len(str) == 0 {
		return nil
	}

	head := &RadixNode{char: rune(str[0])}
	current := head

	for _, char := range str[1:] {
		node := &RadixNode{char: char}
		current.next = node
		current = node
	}

	return head
}

func (tree *RadixNode) search(value string) bool {
	for index := range value {
		currentIndex := index
		currentNode := tree
		for {
			if currentIndex >= len(value) {
				break
			}

			if currentNode.char == rune(value[currentIndex]) {
				if currentNode.next == nil {
					return true
				} else {
					currentIndex += 1
					currentNode = currentNode.next
				}
			} else {
				break
			}
		}
	}

	return false
}
