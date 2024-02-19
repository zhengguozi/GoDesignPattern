package Composite

import "fmt"

// 组合模式其实就是一种树形结构
type Component interface {
	Search(pre, keyword string)
}

type File struct {
	Name string
}

func (f *File) Search(pre, keyword string) {
	fmt.Println(pre + "在" + f.Name + "搜索" + keyword)
}

type Folder struct {
	Name      string
	Component []Component
}

func (f *Folder) Search(pre, keyword string) {
	fmt.Println(pre + "在" + f.Name + "文件夹搜索" + keyword)
	for _, composite := range f.Component {
		composite.Search(" "+pre, keyword)
	}
}

func (f *Folder) AddComponent(component Component) {
	f.Component = append(f.Component, component)
}
