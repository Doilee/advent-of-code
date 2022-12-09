package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const TOTAL_DISK_SPACE = 70000000
const NEEDED_SPACE = 30000000

type File struct {
	IsDirectory bool
	Parent      *File
	Children    []*File
	Depth       int
	Name        string
	Size        uint64
}

func main() {
	bytes, _ := os.ReadFile("2022/day7.txt")
	puzzleInput := string(bytes)
	data := strings.Split(puzzleInput, "\n")

	// a [ ]
	filesystem := &File{Name: "/"}
	currentDirectory := filesystem

	for _, terminal_line := range data {
		parts := strings.Split(terminal_line, " ")

		// is command
		if parts[0] == "$" {
			if parts[1] == "cd" {

				if parts[2] == "/" {
					continue
				} else if parts[2] == ".." {
					currentDirectory.Size += getFileSizeOfFile(currentDirectory)

					currentDirectory = currentDirectory.Parent
				} else {
					for _, child := range currentDirectory.Children {
						if child.IsDirectory && child.Name == parts[2] {
							currentDirectory = child
						}
					}
				}

			}
		} else if parts[0] == "dir" {
			currentDirectory.Children = append(currentDirectory.Children, &File{IsDirectory: true, Name: parts[1], Size: 0, Parent: currentDirectory, Depth: currentDirectory.Depth + 1})
		} else {
			// is a file
			size, _ := strconv.ParseUint(parts[0], 10, 64)
			currentDirectory.Children = append(currentDirectory.Children, &File{IsDirectory: false, Name: parts[1], Size: size, Parent: currentDirectory, Depth: currentDirectory.Depth + 1})
		}
	}
	// go back to root to set filesystem
	fmt.Println("Looping back to get to filesystem..")

	for currentDirectory.Parent != nil {
		//currentDirectory.Parent.Size += currentDirectory.Size
		currentDirectory.Size += getFileSizeOfFile(currentDirectory)
		currentDirectory = currentDirectory.Parent
	}

	filesystem = currentDirectory

	filesystem.Size = getFileSizeOfFile(filesystem)

	fmt.Println(filesystem.Size)

	fmt.Println("Part 1", getPartOneAnswer(filesystem))

	unusedSpace := TOTAL_DISK_SPACE - filesystem.Size
	minDirSpace := NEEDED_SPACE - unusedSpace
	candidates := findCandidates(filesystem, minDirSpace)

	// Sort from min size to max size
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].Size < candidates[j].Size
	})
	fmt.Println("Part 2", candidates[0])
}

func getFileSizeOfFile(directory *File) uint64 {
	var totalSize uint64

	for _, child := range directory.Children {
		if child.IsDirectory {
			totalSize += getFileSizeOfFile(child)
		} else {
			totalSize += child.Size
		}
	}

	return totalSize
}

func getPartOneAnswer(file *File) uint64 {
	var totalSize uint64

	if file.Size <= 100000 {
		totalSize += file.Size
	}

	for _, child := range file.Children {
		if !child.IsDirectory {
			continue
		}

		totalSize += getPartOneAnswer(child)
	}

	return totalSize
}

func findCandidates(file *File, minSize uint64) []*File {
	directories := make([]*File, 0)

	if file.Size >= minSize { // this is the size we need to free up
		directories = append(directories, file)
	}

	for _, child := range file.Children {
		directories = append(directories, findCandidates(child, minSize)...)
	}

	return directories
}
