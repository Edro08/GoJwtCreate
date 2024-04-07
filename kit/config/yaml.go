package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type IConfig interface {
	GetString(keys string) (string, bool)
	GetInt(keys string) (int, bool)
	GetBool(keys string) (bool, bool)
	GetMapInterface(keys string) (map[string]interface{}, bool)
	GetMapString(keys string) (map[string]string, bool)
}

type Config struct {
	data map[string]interface{}
}

func NewConfig(file string) *Config {
	// Lee el archivo YAML
	yamlFile, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	// Declara una variable para almacenar el YAML deserializado
	var data map[string]interface{}

	// Deserializa el YAML en la variable
	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		panic(err)
	}

	return &Config{data: data}
}

func (c *Config) GetString(keys string) (string, bool) {
	value, found := c.getValue(keys)
	return fmt.Sprintf("%v", value), found
}

func (c *Config) GetInt(keys string) (int, bool) {
	value, found := c.getValue(keys)
	if intValue, ok := value.(int); ok {
		return intValue, found
	}
	return 0, false
}

func (c *Config) GetBool(keys string) (bool, bool) {
	value, found := c.getValue(keys)
	if boolValue, ok := value.(bool); ok {
		return boolValue, found
	}
	return false, false
}

func (c *Config) GetMapInterface(keys string) (map[string]interface{}, bool) {
	value, found := c.getMap(keys, "interface")
	if value, ok := value.(map[string]interface{}); ok {
		return value, found
	}
	return nil, false
}

func (c *Config) GetMapString(keys string) (map[string]string, bool) {
	value, found := c.getMap(keys, "string")
	if value, ok := value.(map[string]string); ok {
		return value, found
	}
	return nil, false
}

// Funcion para obtener un valor
func (c *Config) getValue(keys string) (interface{}, bool) {
	keysSlice := strings.Split(keys, ".")
	current := c.data
	for _, key := range keysSlice {
		val, ok := current[key]
		if !ok {
			return nil, false
		}
		if next, ok := val.(map[string]interface{}); ok {
			current = convertMap(next)
		} else {
			return val, true
		}
	}
	return current, true
}

// Función para convertir un mapa de claves de tipo interface{} a un mapa de claves de tipo string
func convertMap(input map[string]interface{}) map[string]interface{} {
	output := make(map[string]interface{})
	for key, val := range input {
		output[fmt.Sprintf("%v", key)] = val
	}
	return output
}

// Funcion para obtener Map
func (c *Config) getMap(keys string, typeMap string) (interface{}, bool) {
	keysSlice := strings.Split(keys, ".")
	current := c.data
	for _, key := range keysSlice {
		val, ok := current[key]
		if !ok {
			return nil, false
		}

		if next, ok := val.(map[string]interface{}); ok {
			current = convertMap(next)
		}
	}

	if typeMap == "interface" {
		return current, true
	} else if typeMap == "string" {
		stringMap := make(map[string]string)
		for k, v := range current {
			stringMap[k] = fmt.Sprintf("%v", v)
		}
		return stringMap, true
	}

	return nil, false
}
