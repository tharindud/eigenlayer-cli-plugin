# EigenLayer CLI Plugin

This repository contains a demo plugin for EigenLayer CLI for custom off-chain code execution for AVS registration flows.

## Creating a Minimal Plugin

A minimal plugin for EigenLayer CLI requires a shared library to be built that includes a symbol named `PluginCoordinator` that implement the following interface.

```go
type Coordinator interface {
    Type() string
	Register() error
    OptIn() error
    OptOut() error
    Deregister() error
    Status() (string, error)
}
```

In the above interface, the `Type()` function should return the type of the coordinator, which can be a non-empty string that identifies the coordinator implementation.

The `Register()`, `OptIn()`, `OptOut()` and `Deregister()` functions are invoked when the plugin is to execute the logic for the corresponding AVS registration workflows.

The `Status()` function is called to query the registration status of an operator for an AVS.

### Create the Project

Create a new directory to hold the project (following example uses `eigenlayer-cli-plugin-demo`).

```bash
$ mkdir eigenlayer-plugin-demo
$ cd eigenlayer-plugin-demo
```

Initialize the project as a module (following example uses `github.com/eigenlayer/eigenlayer-cli-demo`).

```bash
$ go mod init github.com/eigenlayer/eigenlayer-cli-demo
```

### Add a Coordinator

Create the plugin coordinator implementation (following is added as `coordinator.go`).

```go
package main

import (
    "fmt"
)

type Coordinator struct {
}

func (coordinator Coordinator) Type() string {
    return "eigenlayer-cli-demo"
}

func (coordinator Coordinator) Register() error {
    fmt.Println("eigenlayer-cli-demo:register")
    return nil
}

func (coordinator Coordinator) OptIn() error {
    fmt.Println("eigenlayer-cli-demo:opt-in")
    return nil
}

func (coordinator Coordinator) OptOut() error {
    fmt.Println("eigenlayer-cli-demo:opt-out")
    return nil
}

func (coordinator Coordinator) Deregister() error {
    fmt.Println("eigenlayer-cli-demo:deregister")
    return nil
}

func (coordinator Coordinator) Status() (string, error) {
    fmt.Println("eigenlayer-cli-demo:status")
    return "", nil
}

var PluginCoordinator Coordinator
```

### Build the Plugin

Build the project as a plugin (following example build `eigenlayer-cli-demo.so`).

```bash
    $ go build -buildmode=plugin -o eigenlayer-cli-demo.so
```

### Host the Plugin

Host the plugin shared library so that it can be accessible from a public URL.

### Use it in a Specification

The plugin can be used in a specification by setting the specification's `coordinator` and `library_url` attributes.

The following example shows an `avs.json` of a specification that uses a plugin hosted at `https://download.eigenlayer.xys/cli/eigenlayer-cli-demo.so`.

```json
{
    "name": "eigenlayer-cli-demo",
    "network": "mainnet",
    "contract_address": "0x870679e138bcdf293b7ff14dd44b70fc97e12fc0",
    "coordinator": "plugin",
    "remote_signing": false,
    "library_url": "https://download.eigenlayer.xys/cli/eigenlayer-cli-demo.so"
}
```

## Accessing the Specification

The plugin can access the specification used for launching the coordinator by including a symbol named `PluginSpecification` that implements the following interface.

```go
type Specification interface {
    Type() string
    Validate() error
}
```

The `Type()` function should return the type of the specification implementation. The `Validate()` function is called after loading the specification in order to check its validity.

Note that this plugin specification can also include custom properties included in the corresponding `avs.json` file.

For example, consider the following `avs.json`.

```json
{
    "name": "eigenlayer-cli-demo",
    "description": "Specification for CLI Plugin Demo",
    "network": "mainnet",
    "contract_address": "0x870679e138bcdf293b7ff14dd44b70fc97e12fc0",
    "coordinator": "plugin",
    "remote_signing": false,
    "library_url": "https://download.eigenlayer.xys/cli/eigenlayer-cli-demo.so",
    "foo": "bar"
}
```

The plugin can access the specification as follows.

```go
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

func (spec Specification) Type() string {
    return spec.Coordinator
}

func (spec Specification) Validate() error {
    if spec.Foo == "" {
        return errors.New("specification: foo is required")
    }

    return nil
}

var PluginSpecification Specification
```

## Accessing Configuration Parameters

The plugin can access the configuration parameters used when workflows are invoked by including a symbol named `PluginConfiguration` that implements the following interface.

```go
type Configuration interface {
    Get(key string) interface{}
    Set(key string, value interface{})
}
```

The following is a full example.

```go
package main

type Configuration struct {
	registry map[string]interface{}
}

func (config Configuration) Get(key string) interface{} {
	return config.registry[key]
}

func (config Configuration) Set(key string, value interface{}) {
	config.registry[key] = value
}

var PluginConfiguration Configuration
```
