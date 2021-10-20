package composite

import "fmt"

type component interface {
	search(string)
}

// file
type file struct {
	name string
}

func (f *file) search(keyword string) {
	fmt.Printf("Searching in %s\n", f.name)
	if keyword == f.name {
		fmt.Println("find!!")
	}
}

func (f *file) getName() string {
	return f.name
}

// folder
type folder struct {
	components []component
	name       string
}

func (f *folder) search(keyword string) {
	fmt.Printf("Searching in folder %s\n", f.name)
	for _, composite := range f.components {
		composite.search(keyword)
	}
}

func (f *folder) add(c component) {
	f.components = append(f.components, c)
}
