package homing

import (
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"gopkg.in/ini.v1"
)

type Aide struct {
	kvMap map[string]string
}

// NewAide new an instance.
func NewAide() *Aide {
	return &Aide{
		kvMap: make(map[string]string),
	}
}

// LoadFromIni load key and values from ini config file.
func (aide *Aide) LoadFromIni(filePath string) error {
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
			aide.kvMap[k] = v
		}
	}

	return nil
}

// ReplaceTextFile replace placeholders in text file with config values.
func (aide *Aide) ReplaceTextFile(filePath string) error {
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
		v := aide.getValue(k)
		kvMap[each] = v
	}

	for k, v := range kvMap {
		data = strings.ReplaceAll(data, k, v)
	}

	return ioutil.WriteFile(filePath, []byte(data), os.ModePerm)
}

func (aide *Aide) getValue(key string) string {
	value, ok := aide.kvMap[key]
	if !ok {
		return "_UNSET_"
	}
	return value
}
