package homing

import (
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"gopkg.in/ini.v1"
)

type Aider struct {
	kvMap map[string]string
}

// NewAider new an instance.
func NewAider() *Aider {
	return &Aider{
		kvMap: make(map[string]string),
	}
}

// LoadFromIni load key and values from ini config file.
func (aider *Aider) LoadFromIni(filePath string) error {
	file, err := ini.Load(filePath)
	if err != nil {
		return err
	}

	sections := file.Sections()
	for _, section := range sections {
		sectionName := section.Name()
		prefix := ""
		if sectionName != "DEFAULT" {
			prefix = sectionName + "."
		}

		keys := section.Keys()
		for _, key := range keys {
			k := prefix + key.Name()
			v := key.Value()
			aider.kvMap[k] = v
		}
	}

	return nil
}

// ReplaceTextFile replace placeholders in text file with config values.
func (aider *Aider) ReplaceTextFile(filePath string) error {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	data := string(bytes)

	kvMap := make(map[string]string)
	re := regexp.MustCompile("{{(.+?)}}")
	matches := re.FindAllString(data, -1)
	for _, each := range matches {
		k := each[2 : len(each)-2]
		v := aider.getValue(k)
		kvMap[each] = v
	}

	for k, v := range kvMap {
		data = strings.ReplaceAll(data, k, v)
	}

	return ioutil.WriteFile(filePath, []byte(data), os.ModePerm)
}

func (aider *Aider) getValue(key string) string {
	value, ok := aider.kvMap[key]
	if !ok {
		return "_UNSET_"
	}
	return value
}
