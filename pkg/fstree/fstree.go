package fstree

import (
	"bufio"
	"strings"
)

type FsNode struct {
	Name     string    `json:"name"`
	Type     string    `json:"type"`
	Parent   string    `json:"parent"`
	Path     string    `json:"path"`
	Children []*FsNode `json:"children"`
}

type FsTree struct {
	Root FsNode `json:"root"`
}

// ConstructFsTree generate file system hierarchy from find command output.
// Command sample: `find . -exec bash -c 'x=""; if [ -d "{}" ]; then x="/"; fi; printf "{}$x\n"' \; | sort`
func ConstructFsTree(text string) *FsTree {
	tree := &FsTree{}

	nodeMap := make(map[string]*FsNode, 0)
	nodeMap[""] = &tree.Root

	r := strings.NewReader(text)
	buf := bufio.NewScanner(r)
	for {
		if !buf.Scan() {
			break
		}

		line := buf.Text()
		if strings.Contains(line, FileNotFound) {
			continue
		}

		elements := strings.Split(line, "/")
		size := len(elements)
		if size < 2 {
			continue
		}

		leafIndex := size - 1
		leafText := elements[leafIndex]
		nodeType := TypeFile
		nameIndex := leafIndex
		parentIndex := nameIndex - 1
		if leafText == "" {
			nodeType = TypeDirectory
			nameIndex = leafIndex - 1
			parentIndex = nameIndex - 1
		}

		nodeName := elements[nameIndex]
		parentPath := ""
		selfPath := nodeName
		if parentIndex >= 0 {
			parentPath = strings.Join(elements[:nameIndex], "/")
			selfPath = strings.Join(elements[:nameIndex+1], "/")
		}

		node := &FsNode{
			Name:     nodeName,
			Type:     nodeType,
			Parent:   parentPath,
			Path:     selfPath,
			Children: make([]*FsNode, 0),
		}

		parent, ok := nodeMap[parentPath]
		if !ok {
			parent = nodeMap[""]
		}
		parent.Children = append(parent.Children, node)

		if node.Type == TypeDirectory {
			nodeMap[node.Path] = node
		}
	}

	return tree
}
