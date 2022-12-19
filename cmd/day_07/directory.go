package main

import (
	"errors"
	"fmt"
)

type directory struct {
	name        string
	size        int
	files       []*item
	directories map[string]*directory
	parent      *directory
}

type item struct {
	name string
	size int
}

func (d *directory) createFile(fileName string, fileSize int) {
	d.updateSize(fileSize)
	d.files = append(d.files, &item{name: fileName, size: fileSize})
}

func (d *directory) createDirectory(directoryName string) {
	d.directories[directoryName] = &directory{
		name:        directoryName,
		size:        0,
		parent:      d,
		directories: map[string]*directory{},
	}
}

func (d *directory) updateSize(size int) {
	d.size += size
	if d.parent != nil {
		d.parent.updateSize(size)
	}
}

func (d *directory) crawlDirectory(directoryName string) (pointer *directory, err error) {
	if directoryName == ".." {
		pointer, err = d.parent, nil
	} else if d.directories[directoryName] != nil {
		pointer, err = d.directories[directoryName], nil
	} else {
		pointer, err = nil, errors.New(fmt.Sprintf("error: directory %s not found", directoryName))
	}

	return
}

/*
* findDirectoriesSmallerThan will take a limit and a pointer to an item array.  Then find all directories smaller than
* the limit.
 */
func (d *directory) findDirectoriesSmallerThan(limit int, directories *[]item) {
	if d.size <= limit {
		*directories = append(*directories, item{name: d.name, size: d.size})
	}
	if d.directories != nil {
		for _, dir := range d.directories {
			dir.findDirectoriesSmallerThan(limit, directories)
		}
	}
}

/*
* findDirectoriesGreaterThan will take a limit and a pointer to an item array.  Then find all directories greater than
* the limit.
 */
func (d *directory) findDirectoriesGreaterThan(limit int, directories *[]item) {
	if d.size >= limit {
		*directories = append(*directories, item{name: d.name, size: d.size})
	}
	if d.directories != nil {
		for _, dir := range d.directories {
			dir.findDirectoriesGreaterThan(limit, directories)
		}
	}
}
