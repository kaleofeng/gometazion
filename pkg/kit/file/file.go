package file

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

func Exists(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func IsDir(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), nil
}

func ListFiles(dir string, filePattern string) ([]string, error) {
	var filePaths []string

	fis, err := ioutil.ReadDir(dir)
	if err != nil {
		return filePaths, err
	}

	reg := regexp.MustCompile(filePattern)
	for _, fi := range fis {
		filePath := filepath.Clean(filepath.Join(dir, fi.Name()))
		if fi.IsDir() {
			fps, err := ListFiles(filePath, filePattern)
			if err != nil {
				return filePaths, err
			}

			filePaths = append(filePaths, fps...)
			continue
		}

		if reg.MatchString(fi.Name()) {
			filePaths = append(filePaths, filePath)
		}
	}

	return filePaths, err
}
