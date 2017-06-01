package main

import (
	"fmt"
	"errors"
	"github.com/satori/go.uuid"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	log "github.com/sirupsen/logrus"
)

func addUsage() error {
	log.Info("Usage: gadget [flags] add [type] [name]")
	log.Info("               *opt        *req   *req ")
	log.Info("Type: service | onboot                 ")
	log.Info("Name: friendly name for container      ")
	
	return errors.New("Incorrect add usage")
}

// Process the build arguments and execute build
func GadgetAdd(args []string, g *GadgetContext) error {

	//~ g.LoadConfig()

	addUu := uuid.NewV4()
	
	if len(args) != 2 {
		return addUsage()
	}
	
	log.Info(fmt.Sprintf("Adding new %s: \"%s\" ", args[0], args[1]))
	
	addGadgetContainer := GadgetContainer {	
		Name: 	args[1], 
		Image: 	fmt.Sprintf("%s/%s", g.Config.Name, args[1]),
		UUID: 	fmt.Sprintf("%s", addUu),
	}
	
	// parse arguments
	switch args[0] {
	case "service":
		g.Config.Services = append(g.Config.Services, addGadgetContainer)
	case "onboot":
		g.Config.Onboot = append(g.Config.Onboot, addGadgetContainer)
	default:
		log.Error(fmt.Sprintf("%q is not valid command.\n", args[0]))
		return addUsage()
	}
	
	g.Config = CleanConfig(g.Config)

	outBytes, err := yaml.Marshal(g.Config)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fmt.Sprintf("%s/gadget.yml", g.WorkingDirectory), outBytes, 0644)
	if err != nil {
		return err
	}
	
	//~ fmt.Printf("✔\n")
	
	return err
}
