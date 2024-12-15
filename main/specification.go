package main

import (
	"errors"
)

type Specification struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	Network         string `json:"network"`
	ContractAddress string `json:"contract_address"`
	Coordinator     string `json:"coordinator"`
	RemoteSigning   bool   `json:"remote_signing"`
	LibraryURL      string `json:"library_url"`
	Foo             string `json:"foo"`
}

func (spec Specification) Validate() error {
	if spec.Foo == "" {
		return errors.New("specification: foo is required")
	}

	return nil
}

var PluginSpecification Specification
