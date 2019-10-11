package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

var (
	envs = map[string]bool{
		"global": true,
		"qa":     true,
		"local":  true,
	}
	envValue string
)

func init() {
	envValue = getEnv()
}

type Config struct {
	Project struct {
		Name  string
		Debug bool
	}
	Web struct {
		Port int
	}
	Service struct {
		Google struct {
			Host    string
			Enabled bool
		}
	}
}

func Get() *Config {
	cfg := &Config{}
	loadConfig(envValue, cfg)
	return cfg
}

func loadConfig(env string, cfg *Config) {
	path := getPath(env)

	dep := getDep(path)
	if len(dep) > 0 {
		loadConfig(dep, cfg)
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		panic(err)
	}

	d, err := yaml.Marshal(cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("Config dump:\n%s", string(d))
}

func getEnv() string {
	args := os.Args
	if len(args) >= 2 && envs[args[1]] {
		return args[1]
	}

	return "global"
}

func getPath(env string) string {
	return "cli/config/project_" + env + ".yaml"
}
