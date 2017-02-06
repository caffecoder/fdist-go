package fdist

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

type FileDistribution struct {
	ext    string
	prefix string
	path   string
}

func NewFileDistribution(prefix string) *FileDistribution {
	fd := FileDistribution{".dat", prefix, prefix}

	return &fd
}

func (fd *FileDistribution) HexPath(id int) {
	var hex string
	hex = fmt.Sprintf("%x", id)

	if len(hex)%2 != 0 {
		hex = fmt.Sprintf("0%s", hex)
	}

	for j := len(hex) - 2; j > 1; j -= 2 {
		hex = path.Join(hex[:j], hex[j:])
	}

	fd.path = path.Join(fd.prefix, hex)
	fd.path += fd.ext
}

func (fd *FileDistribution) GetPath() string {
	return fd.path
}

func (fd *FileDistribution) SetExtension(ext string) {
	if ext != "" && ext[0] != '.' {
		fd.ext = "." + ext
	} else {
		fd.ext = ext
	}
}

func (fd *FileDistribution) RenameFrom(path string) error {
	var err error

	dst := filepath.Dir(fd.path)

	if _, err := os.Stat(dst); os.IsNotExist(err) {
		err = os.MkdirAll(dst, os.ModeDir|os.ModePerm)
	}

	if err == nil {
		err = os.Rename(path, fd.path)
	}

	return err
}
