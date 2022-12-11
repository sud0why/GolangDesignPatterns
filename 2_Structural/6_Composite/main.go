package main

import "fmt"

//组合模式
//树形结构

// 抽象文件

type FileNode interface {
	Search(string) bool
}

// 目录实现

type Folder struct {
	name  string
	nodes []FileNode
}

func (f *Folder) Search(searchName string) bool {
	if f.name == searchName {
		fmt.Printf("Find %s\n", searchName)
		return true
	}
	for _, node := range f.nodes {
		if node.Search(searchName) {
			return true
		}
	}
	return false
}

func (f *Folder) Add(node FileNode) {
	f.nodes = append(f.nodes, node)
}

// 文件实现

type File struct {
	name string
}

func (f *File) Search(searchName string) bool {
	if f.name == searchName {
		fmt.Printf("Find %s\n", searchName)
		return true
	} else {
		return false
	}
}

//业务

func main() {
	file := &File{name: "c"}
	folder1 := &Folder{name: "admin"}
	folder1.Add(&File{name: "a1"})
	folder1.Add(&File{name: "a2"})
	folder1.Add(&Folder{
		name: "a3",
		nodes: []FileNode{
			&File{name: "b1"},
			&File{name: "b2"},
			file,
		},
	})
	fmt.Println(file.Search("c"))
	fmt.Println(file.Search("b"))
	fmt.Println(folder1.Search("c"))
	fmt.Println(folder1.Search("b3"))

}
