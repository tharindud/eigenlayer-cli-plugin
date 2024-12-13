package main

import (
	"fmt"
)

type Coordinator struct {
}

func (coordinator Coordinator) Type() string {
	return "eigenlayer-cli-plugin"
}

func (coordinator Coordinator) Register() error {
	fmt.Println("eigenlayer-cli-plugin:register")
	fmt.Printf("specification:name=%s\n", PluginSpecification.Name)
	fmt.Printf("specification:foo=%s\n", PluginSpecification.Foo)
	fmt.Printf("configuration:bar=%s\n", PluginConfiguration.Get("bar"))
	return nil
}

func (coordinator Coordinator) OptIn() error {
	fmt.Println("eigenlayer-cli-plugin:opt-in")
	return nil
}

func (coordinator Coordinator) OptOut() error {
	fmt.Println("eigenlayer-cli-plugin:opt-out")
	return nil
}

func (coordinator Coordinator) Deregister() error {
	fmt.Println("eigenlayer-cli-plugin:deregister")
	return nil
}

func (coordinator Coordinator) Status() (string, error) {
	fmt.Println("eigenlayer-cli-plugin:status")
	return "", nil
}

var PluginCoordinator Coordinator
