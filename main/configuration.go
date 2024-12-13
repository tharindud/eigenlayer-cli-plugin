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
