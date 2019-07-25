package classpath

import (
	"archive/zip"
	"github.com/getlantern/errors"
	"io/ioutil"
	"path/filepath"
)

type EntryZip struct {
	absPath string
}

func newZipEntry(path string) *EntryZip {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &EntryZip{absPath: absPath}
}

func (self *EntryZip) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()
	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (self *EntryZip) String() string {
	return self.absPath
}
