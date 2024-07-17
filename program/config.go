package program

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Programs map[string]Program

func Load() Programs {
	b, err := os.ReadFile("./program.yml")
	if err != nil {
		return nil
	}
	var config Programs
	if err := yaml.Unmarshal(b, &config); err != nil {
		return nil
	}
	for s, program := range config {
		program.Name = s
		config[s] = program
	}
	return config
}

func List() []Program {
	b, err := os.ReadFile("./program.yml")
	if err != nil {
		panic(err)
	}
	var node yaml.Node
	if err := yaml.Unmarshal(b, &node); err != nil {
		panic(err)
	}
	var programs []Program
	if node.Kind == yaml.DocumentNode && len(node.Content) > 0 {
		root := node.Content[0]
		if root.Kind == yaml.MappingNode {
			for i := 0; i < len(root.Content); i += 2 {
				var name string
				var program Program
				if err := root.Content[i].Decode(&name); err != nil {
					panic(err)
				}
				if err := root.Content[i+1].Decode(&program); err != nil {
					panic(err)
				}
				program.Name = name
				programs = append(programs, program)
			}
		}
	}
	return programs
}
