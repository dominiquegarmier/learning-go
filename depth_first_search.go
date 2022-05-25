package main

type DiGraph struct {
	nodes map[int]*Node
}

type Node struct {
	id  int
	out []*Edge
	in  []*Edge
}

type Edge struct {
	id   int
	from *Node
	to   *Node
}

type StackNode struct {
	value *Node
	next  *StackNode
}

type Stack struct {
	top *StackNode
}

func (s *Stack) push(value *Node) {
	s.top = &StackNode{value, s.top}
}

func (s *Stack) pop() *Node {
	if s.top == nil {
		return nil
	}
	ret := s.top.value
	s.top = s.top.next
	return ret
}

func (g *DiGraph) depth_first_search(from *Node, to *Node) bool {
	var stack Stack
	var visited map[int]bool
	stack.push(from)
	for stack.top != nil {
		curr := stack.pop()
		for _, edge := range curr.out {
			if !visited[edge.to.id] {
				stack.push(edge.to)
			}
			if edge.to == to {
				return true
			}
		}
		visited[curr.id] = true
	}
	return false
}
