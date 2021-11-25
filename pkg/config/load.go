package config

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// LoadFiles load and parse config files
func (c *Config) LoadFiles(sourceFiles ...string) (err error) {
	for _, file := range sourceFiles {
		if err = c.loadFile(file); err != nil {
			return
		}
	}
	return
}

func (c *Config) loadFile(file string) (err error) {
	var format string
	if format == "" {
		format = strings.Trim(filepath.Ext(file), ".")
	}

	fileContent, err := ioutil.ReadFile(file)

	if err != nil {
		return err
	}

	if err = c.parseData(format, fileContent); err != nil {
		return
	}

	return nil
}

func (c *Config) parseData(extension string, byteArray []byte) (err error) {
	decoder := c.decoders[extension]
	if decoder == nil {
		return errors.New("no decoder present for given file format")
	}
	data, _ := decoder(byteArray)

	if len(c.data) == 0 {
		c.data = data
	} else {
		c.data = merge(c.data, data)
	}
	return nil
}
