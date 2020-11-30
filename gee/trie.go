package gee

import (
	"strings"
	"github.com/davecgh/go-spew/spew"
)

type node struct {
	pattern  string
	part     string
	children []*node
	isWild   bool
}

// 第一个匹配的节点用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// 所有匹配成功的节点
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	spew.Dump(n.children, part)
	for _, child := range n.children {
		print("here \n")
		print(child.part, "\n")
		print(part, "\n")
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	print("what is nodes \n")
	print(nodes)
	return nodes
}

func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)


	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}
	print("1****\n")
	spew.Dump(parts, height)
	print("2****\n")
	part := parts[height]
	spew.Dump(part, n)
	print("3****\n")
	children := n.matchChildren(part)
	spew.Dump(n)
	for _, child := range children {
		print("range children\n")
		print(child)
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}

func (n *node) showAllPath(){
	print(n.part, "------", n.pattern, "\n")

	for _, node := range n.children{
		node.showAllPath()
	}
}