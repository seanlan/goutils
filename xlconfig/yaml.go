package xlconfig

import (
	_config "go.uber.org/config"
)

type YamlProvider struct {
	Config _config.YAML
}

func NewYAMLProvider(path string) (provider Provider, err error) {
	if len(path) == 0 {
		path = "conf.d/conf.yaml"
	}
	yaml, err := _config.NewYAML(_config.File(path))
	if err != nil {
		return nil, err
	}
	provider = YamlProvider{
		Config: *yaml,
	}
	return
}

func (y YamlProvider) GetString(key ...string) (string, error) {
	value := y.Config.Get("")
	for _, k := range key {
		value = value.Get(k)
	}
	return value.String(), nil
}

func (y YamlProvider) GetBool(key ...string) (bool, error) {
	var result bool
	value := y.Config.Get("")
	for _, k := range key {
		value = value.Get(k)
	}
	err := value.Populate(&result)
	if err != nil {
		return false, err
	}
	return result, err
}

func (y YamlProvider) GetInt(key ...string) (int64, error) {
	var result int64
	value := y.Config.Get("")
	for _, k := range key {
		value = value.Get(k)
	}
	err := value.Populate(&result)
	if err != nil {
		return 0, err
	}
	return result, err
}

func (y YamlProvider) GetFloat(key ...string) (float64, error) {
	var result float64
	value := y.Config.Get("")
	for _, k := range key {
		value = value.Get(k)
	}
	err := value.Populate(&result)
	if err != nil {
		return 0, err
	}
	return result, err
}

func (y YamlProvider) Populate(i interface{}, key ...string) error {
	value := y.Config.Get("")
	for _, k := range key {
		value = value.Get(k)
	}
	err := value.Populate(i)
	return err
}
