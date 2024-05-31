package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func UpdateYaml(name string) error {
	// read the current yaml file
	yamlFile, err := os.ReadFile(Config.filename)
	if err != nil {
		return err
	}
	// parse the yaml file
	var router Routers
	err = yaml.Unmarshal(yamlFile, &router)
	if err != nil {
		return err
	}
	// check if the router exists
	if _, ok := router.HTTP.Routers[name]; !ok {
		// first check is map exists
		if router.HTTP.Routers == nil {
			router.HTTP.Routers = make(map[string]Route)
		}
		// if it doesn't exist, create it
		router.HTTP.Routers[name] = Route{
			EntryPoints: []string{Config.entryPoint},
			Service:     Config.service,
			Rule:        "Host(`" + name + Config.domain + "`)",
			TLS: struct {
				CertResolver string `yaml:"certResolver"`
			}{
				CertResolver: Config.certResolver,
			},
		}
		// write the new router to the yaml file
		newYaml, err := yaml.Marshal(&router)
		if err != nil {
			return err
		}
		err = os.WriteFile(Config.filename, newYaml, 0644)
		if err != nil {
			return err
		}
		log.Println("Router updated for", name)
	}

	return nil

}
