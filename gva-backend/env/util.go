package env

import (
	"fmt"
	"sort"

	"os"
	"strings"

	"github.com/spf13/viper"
	"golang.org/x/exp/maps"
)

func ReadToml(filename string, path ...string) (tomlFile *viper.Viper, config *Config, err error) {
	config = new(Config)
	tomlFile = viper.New()
	tomlFile.SetConfigName(filename)
	tomlFile.SetConfigType("toml")

	for _, p := range path {
		tomlFile.AddConfigPath(p)
	}

	if err := tomlFile.ReadInConfig(); err != nil {
		return nil, nil, err
	}

	if err := tomlFile.Unmarshal(&config); err != nil {
		return nil, nil, err
	}

	return tomlFile, config, nil
}

const (
	tomlFilePath = "./env"
	tomlFileName = "config"

	envFilePath = "."
	envFileName = ".env"
)

func Environ() map[string]string {
	m := make(map[string]string)
	for _, s := range os.Environ() {
		eqIndex := strings.Index(s,"=")
		if eqIndex < 0{
			continue
		}
		key,value := s[:eqIndex],s[eqIndex:]
		if key != "" {
			m[key] = value
		}
	}
	return m
}

func ReadEnv(filename string, path ...string) (envFile *viper.Viper, config *Config, err error) {
	config = new(Config)
	envFile = viper.New()
	envFile.SetConfigName(filename)
	envFile.SetConfigType("env")

	for _, p := range path {
		envFile.AddConfigPath(p)
	}

	if err := envFile.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			for k, v := range Environ() {
				envFile.Set(k, v)
			}
		} else {
			return nil, nil, fmt.Errorf("failed to read env %v", err)
		}
	}

	if err := envFile.Unmarshal(&config); err != nil {
		return nil, nil, err
	}

	return envFile, config, nil
}

func ReadEnvOrGenerate() (*Config, error) {
	_, config, err := ReadEnv(envFilePath, envFileName)
	if err != nil {
		return nil, err
	}

	if config.App.Env == "" {
		fmt.Println("env not found, generating env from env/config.toml...")
		if err := GenerateEnvFromToml(false); err != nil {
			return nil, fmt.Errorf("GenerateEnvFromToml %v", err)
		}

		_, config, err = ReadEnv(envFilePath, envFileName)
		if err != nil {
			return nil, err
		}
	}

	return config, nil
}

func GenerateEnvFromToml(overwrite bool) error {
	_, err := os.Stat(envFileName)
	isExist := err == nil
	if isExist && !overwrite {
		return nil
	}

	tomlFile, _, err := ReadToml(tomlFileName, tomlFilePath)
	if err != nil {
		return err
	}

	flats := flatNestMap(tomlFile.AllSettings())

	keys := maps.Keys(flats)
	sort.Strings(keys)

	envString := ""

	for i, envKey := range keys {
		prefix := strings.Split(envKey, ".")[0]

		if i == 0 {
			envString += fmt.Sprintf("# %s\n", prefix)
		}

		if i > 0 {
			oldPrefix := strings.Split(keys[i-1], ".")[0]
			if prefix != oldPrefix {
				envString += "\n"
				envString += fmt.Sprintf("# %s\n", prefix)
			}
		}

		envVal := flats[envKey]
		envString += fmt.Sprintf("%s=%v \n", strings.ToUpper(envKey), envVal)
	}

	if err := os.WriteFile(envFileName, []byte(envString), 0644); err != nil {
		return err
	}

	return nil
}

func GroupMapAny(nested map[string]any) map[string]map[string]any {
	result := make(map[string]map[string]any)

	for k, v := range nested {
		path := strings.SplitAfter(k, ".")[0]
		_, ok := result[path]
		if !ok {
			result[path] = map[string]any{}
		}

		result[path][k] = v
	}

	return result
}

func flatNestMap(nested map[string]any) map[string]any {
	result := make(map[string]any)

	for k, v := range nested {
		switch v := v.(type) {
		case map[string]any:
			subMap := flatNestMap(v)
			for subK, subV := range subMap {
				path := k + "." + subK
				result[path] = subV
			}
		default:
			result[k] = v
		}

	}

	return result
}
