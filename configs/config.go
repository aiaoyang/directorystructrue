package configs

import (
	"flag"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var filename string
var GCONF *Config

func init() {
	flag.StringVar(&filename, "c", "config.yaml", "specifice config file")
	flag.Parse()
	if ok := flag.Parsed(); !ok {
		flag.Usage()
	}

	GCONF = LoadConfig(filename)
	if GCONF == nil {
		log.Fatal("config parse error")
	}

}

type Config struct {
	Addr string `yaml:"addr"`
	Port int    `yaml:"port"`
}

func LoadConfig(filename string) (cfg *Config) {
	cfg = &Config{}
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(f, cfg)
	if err != nil {
		log.Fatal(err)
	}
	return
}
