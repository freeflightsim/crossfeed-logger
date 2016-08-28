
package server

import (
	"io/ioutil"
	"path/filepath"
	"fmt"

	"gopkg.in/yaml.v2"
)


// Module config
var Config ConfigOpts

// Yaml file reader
type ConfigOpts struct {
	Secret string ` yaml:"secret" `
	HTTPAddress string ` yaml:"http_address" `
	CSVDir string ` yaml:"csv_dir" `
	Db  DBConfig  ` yaml:"db" `
}

type DBConfig struct {
	User string ` yaml:"user" `
	Password string ` yaml:"password" `
	Database string ` yaml:"database" `

}


// LoadConfig and all modules
func LoadConfig(file_path string) (ConfigOpts, error) {
	abs_path, _ := filepath.Abs(file_path)
	yaml_bites, err := ioutil.ReadFile(abs_path)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yaml_bites, &Config)
	if err != nil {
		panic(err)
	}
	fmt.Println("Config=", Config)
	return Config, nil
}