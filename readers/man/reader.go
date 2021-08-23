package man

import (
	"errors"
	"io/ioutil"
	"path"
	"path/filepath"
)

type Reader struct {
	directories []string
}

// NewReader new Reader instance
func NewReader(dir string) *Reader {
	dirs := []string{
		dir,
		"/usr/share/man",
		"/usr/local/share/man",
	}

	return &Reader{
		directories: dirs,
	}
}

func (r *Reader) Read(page string) (string, error) {
	file := r.findPage(page)

	if len(file) == 0 {
		return "", errors.New("Failed to find man page for: " + page)
	}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return "", errors.New("Failed to read man page for " + page + " in " + file)
	}

	return string(data), nil
}

func (r *Reader) findPage(page string) string {
	var pageLocation string

	// loop through top level man directories
	for _, dir := range r.directories {
		pageDirs, err := listInnerDirectories(dir)
		if err != nil {
			continue
		}

		// loop through inner man directories
		for _, pageDir := range pageDirs {
			files, err := ioutil.ReadDir(pageDir)
			if err != nil {
				continue
			}

			// loop through files
			for _, file := range files {
				if !file.IsDir() && trimExtension(file.Name()) == page {
					pageLocation = path.Join(pageDir, file.Name())

					return pageLocation
				}
			}
		}
	}

	return pageLocation
}

func listInnerDirectories(dir string) (dirs []string, err error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		if f.IsDir() {
			dirs = append(dirs, path.Join(dir, f.Name()))
		}
	}

	return dirs, nil
}

func trimExtension(file string) string {
	var extension = filepath.Ext(file)
	var name = file[0 : len(file)-len(extension)]

	return name
}
