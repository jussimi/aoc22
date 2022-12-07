package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed data.txt
var data []byte

func main() {
	fmt.Println("answer1", runPart1(data))
	fmt.Println("answer2", runPart2(data))
}

type file struct {
	name string
	size int32
}

type directory struct {
	parent *directory
	name   string
	files  []*file
	dirs   []*directory
}

func (directory *directory) getSize() int32 {
	var sum int32 = 0

	for _, file := range directory.files {
		sum += file.size
	}
	for _, dir := range directory.dirs {
		sum += dir.getSize()
	}
	return sum
}

type fileSystem struct {
	root    *directory
	current *directory
}

func (fs *fileSystem) setCurrent(dir *directory) {
	fs.current = dir
}

func (fs *fileSystem) back() {
	if back := fs.current.parent; back != nil {
		fs.setCurrent(back)
	}
}

func (fs *fileSystem) goTo(path string) {
	for _, dir := range fs.current.dirs {
		if dir.name == path {
			fs.setCurrent(dir)
		}
	}
	fmt.Println("  - no such directory")
}

func (fs *fileSystem) Move(path string) {
	switch path {
	case "/":
		fs.setCurrent(fs.root)
	case "..":
		fs.back()
	default:
		fs.goTo(path)
	}
}

func (fs *fileSystem) print(str string, indent int) {
	indentation := ""
	for i := 0; i < indent; i += 1 {
		indentation += "  "
	}
	indentation += "- "
	fmt.Println(indentation + str)
}

func (fs *fileSystem) treeAtDepth(directory *directory, depth int) {
	fs.print(fmt.Sprintf("%s (dir)", directory.name), depth)
	for _, dir := range directory.dirs {
		fs.treeAtDepth(dir, depth+1)
	}
	for _, file := range directory.files {
		fs.print(fmt.Sprintf("%s (file, size=%d)", file.name, file.size), depth+1)
	}
}

func (fs *fileSystem) Tree() {
	fs.treeAtDepth(fs.root, 0)
}

func createFileSystem(input []byte) fileSystem {
	rootDir := directory{
		parent: nil,
		name:   "/",
		files:  []*file{},
		dirs:   []*directory{},
	}

	fs := fileSystem{
		root:    &rootDir,
		current: &rootDir,
	}

	for _, line := range strings.Split(string(input), "\n") {
		params := strings.Fields(line)
		if params[0] == "$" {
			switch params[1] {
			case "cd":
				fs.Move(params[2])
			case "ls":
			}
			continue
		}
		if params[0] == "dir" {
			fs.current.dirs = append(fs.current.dirs, &directory{
				parent: fs.current,
				name:   params[1],
				files:  []*file{},
				dirs:   []*directory{},
			})
			continue
		}
		// Is file
		size, _ := strconv.Atoi(params[0])
		fs.current.files = append(fs.current.files, &file{
			size: int32(size),
			name: params[1],
		})
	}
	return fs
}

func addDirs(dir *directory) int32 {
	var sum int32 = 0
	size := dir.getSize()
	if size <= 100000 {
		sum += size
	}
	for _, d := range dir.dirs {
		sum += addDirs(d)
	}
	return sum
}

func findDirsToDelete(neededSpace int32, dir *directory, dirMap map[string]int32) map[string]int32 {
	size := dir.getSize()
	if size >= neededSpace {
		dirMap[dir.name] = size
	}
	for _, d := range dir.dirs {
		dirMap = findDirsToDelete(neededSpace, d, dirMap)
	}
	return dirMap
}

func runPart1(input []byte) int32 {
	fs := createFileSystem(input)
	// fs.Tree()

	return addDirs(fs.root)
}

func runPart2(input []byte) int32 {
	fs := createFileSystem(input)

	availableSpace := 70000000 - fs.root.getSize()
	neededSpace := 30000000 - availableSpace
	dirMap := make(map[string]int32)

	dirMap = findDirsToDelete(neededSpace, fs.root, dirMap)

	smallest := "/"
	for key, size := range dirMap {
		if size < dirMap[smallest] {
			smallest = key
		}
	}
	// fmt.Printf("Smallest %s, size %d\n", smallest, dirMap[smallest])
	return dirMap[smallest]
}
