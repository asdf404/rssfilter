package rssfilter

import "gopkg.in/yaml.v2"

type StringValues []string

// https://github.com/go-yaml/yaml/issues/100#issuecomment-324964723
func (values *StringValues) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var multi []string
	err := unmarshal(&multi)
	if err != nil {
		var single string
		err := unmarshal(&single)
		if err != nil {
			return err
		}
		*values = []string{single}
	} else {
		*values = multi
	}
	return nil
}

type rule struct {
	Author   StringValues `yaml:"author"`
	Category StringValues `yaml:"category"`
	Title    StringValues `yaml:"title"`
	Not      *rule        `yaml:"not"`
}

type Feed struct {
	Url   string `yaml:"url"`
	Path  string `yaml:"path"`
	Title string `yaml:"title"`
	Keep []rule `yaml:"keep"`
	Skip []rule `yaml:"skip"`
}

type settings struct {
	Port  int `yaml:"port"`
	Cache int `yaml:"cache"`
}

type Config struct {
	Settings settings `yaml:"settings"`
	Feeds    []Feed   `yaml:"feeds"`
}

func ParseConfig(data []byte) (*Config, error) {
	config := Config{} // TODO: default settings
	err := yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
