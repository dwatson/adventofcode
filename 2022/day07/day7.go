package day07

import (
	"strconv"
	"strings"
)

type Node struct {
	name     string
	nodetype string
	parent   *Node
	children []*Node
	size     int
}

func BuildTree(values []string) *Node {
	// create root node
	root := &Node{name: "/"}
	currentNode := root
	for _, val := range values {
		if strings.HasPrefix(val, "$") {
			cmd := strings.Split(val, " ")
			switch cmd[1] {
			case "cd":
				if cmd[2] == ".." {
					if currentNode.name != "/" {
						currentNode = currentNode.parent
					}
				} else {
					if currentNode.name != cmd[2] {
						for _, c := range currentNode.children {
							if c.name == cmd[2] {
								currentNode = c
								break
							}
						}
					}
				}
			case "ls":
				//fmt.Println("LS")
			}
		} else {
			output := strings.Split(val, " ")
			switch output[0] {
			case "dir":
				// add dir to currentnode
				currentNode.children = append(currentNode.children, &Node{name: output[1], nodetype: "dir", parent: currentNode})
			default:
				// should be a file node
				filesize, _ := strconv.Atoi(output[0])
				currentNode.children = append(currentNode.children, &Node{name: output[1], nodetype: "file", size: filesize})
			}
		}
	}
	return root
}

func AddDirSizes(root *Node) int {
	for _, i := range root.children {
		if i.nodetype == "file" {
			root.size += i.size
		} else if i.nodetype == "dir" {
			root.size += AddDirSizes(i)
		}
	}
	return root.size
}

func GetTargetDirSizes(root *Node) int {
	total := 0
	targetsize := 100000
	if root.size <= targetsize {
		total += root.size
	}
	for _, i := range root.children {
		if i.nodetype == "dir" {
			total += GetTargetDirSizes(i)
		}
	}

	return total
}

func GetDirs(root *Node) []*Node {
	dirs := []*Node{}
	for _, i := range root.children {
		if i.nodetype == "dir" {
			dirs = append(dirs, i)
			dirs = append(dirs, GetDirs(i)...)
		}
	}
	return dirs
}

func GetDirToDelete(root *Node) int {
	totaldisk := 70000000
	required := 30000000
	target := required - (totaldisk - root.size)
	dirs := GetDirs(root)

	candidate := root
	for _, i := range dirs {
		if i.size >= target {
			if i.size < candidate.size {
				candidate = i
			}
		}
	}
	return candidate.size
}

func Day7_1(values []string) int {
	root := BuildTree(values)
	AddDirSizes(root)
	return GetTargetDirSizes(root)
}

func Day7_2(values []string) int {
	root := BuildTree(values)
	AddDirSizes(root)
	return GetDirToDelete(root)
}
