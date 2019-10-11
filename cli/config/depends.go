package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func getDep(path string) string {

	dep := &struct {
		Depends string
	}{}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, dep)
	if err != nil {
		panic(err)
	}
	return dep.Depends
}
