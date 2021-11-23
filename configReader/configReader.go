package configReader

import (
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
    "os"
)

type Conf struct {
    DatabaseURL string `yaml:"database_url"`
    DatabasePort int32 `yaml:"database_port"`
    DatabaseName string `yaml:"database_name"`
    DatabaseUser string `yaml:"database_user"`
    DatabasePassword string `yaml:"database_password"`
}

func (c *Conf) GetConf(fileLocation string) *Conf {

    yamlFile, err := ioutil.ReadFile(fileLocation)
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
        os.Exit(1)
    }
    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
        os.Exit(1)
    }

    return c
}
