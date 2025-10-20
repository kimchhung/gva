package env

import (
	"backend/utils"
	"fmt"
	"reflect"
	"sort"

	"os"
	"strings"

	"github.com/creasty/defaults"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"golang.org/x/exp/maps"
)

func DefaultConfig() (config *Config, err error) {
	config = new(Config)
	if err := defaults.Set(config); err != nil {
		return nil, err
	}

	return config, nil
}

func Environ() map[string]string {
	m := make(map[string]string)
	for _, s := range os.Environ() {
		eqIndex := strings.Index(s, "=")
		if eqIndex < 0 {
			continue
		}
		key, value := s[:eqIndex], s[eqIndex+1:]
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

	envConfig := viper.New()
	if err := envFile.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			for k, v := range Environ() {
				newK := strings.ReplaceAll(k, seperator, ".")
				envConfig.Set(newK, v)
			}

		} else if utils.IsEmpty(config) {
			return nil, nil, fmt.Errorf("config is empty, failed to read env %v", err)
		}
	}

	settings := envFile.AllSettings()
	for k, v := range settings {
		newK := strings.ReplaceAll(k, seperator, ".")
		envConfig.Set(newK, v)
	}

	if err := envConfig.Unmarshal(&config); err != nil {
		return nil, nil, err
	}
	return envConfig, config, nil
}

func ReadEnvFromFile() (*Config, error) {
	_, config, err := ReadEnv(envFilePath, envFileName)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func structToKeyValue(s interface{}) (map[string]any, error) {
	m := make(map[string]any)
	if err := mapstructure.Decode(s, &m); err != nil {
		return nil, err
	}
	return flatNestMap(m), nil
}

func GenerateEnvFromDefaultConfig(overwrite bool) error {
	_, err := os.Stat(envFileName)
	isDotEnvFileExist := err == nil

	config, err := DefaultConfig()
	if err != nil {
		return err
	}

	keyvalues, err := structToKeyValue(config)
	if err != nil {
		return err
	}

	keys := maps.Keys(keyvalues)
	sort.Strings(keys)

	envString := ""
	for i, envKey := range keys {
		prefix := strings.Split(envKey, seperator)[0]

		if i == 0 {
			envString += fmt.Sprintf("# %s\n", strings.ToUpper(prefix))
		}

		if i > 0 {
			oldPrefix := strings.Split(keys[i-1], seperator)[0]
			if prefix != oldPrefix {
				envString += "\n"
				envString += fmt.Sprintf("# %s\n", strings.ToUpper(prefix))
			}
		}

		envVal := keyvalues[envKey]
		envStr := fmt.Sprintf("%v", envVal)

		switch envVal.(type) {
		case string, bool, int, int8, int16, int32, int64, float32, float64:
		default:
			valType := reflect.TypeOf(envVal)
			val := reflect.ValueOf(envVal)

			switch valType.Kind() {
			case reflect.Slice:
				var sliceVal []any
				sliceVal, ok := val.Interface().([]any)
				if !ok {
					sliceVal = []any{}
				}

				strSlice := make([]string, len(sliceVal))
				for i, v := range sliceVal {
					strSlice[i] = fmt.Sprintf("%v", v)
				}
				envStr = strings.Join(strSlice, ",")
			}

		}

		envString += fmt.Sprintf("%s=%v \n", strings.ToUpper(envKey), envStr)
	}

	if !isDotEnvFileExist || overwrite {
		if err := os.WriteFile(envFileName, []byte(envString), os.ModePerm); err != nil {
			return err
		}
		fmt.Println("env/config.go -> .env : file generated")
	} else {
		fmt.Println("env/config.go -> .env : file already exists, skipping generation")
	}

	donttouch := "# DO NOT TOUCH THIS FILE, please change it in env/config.go then run `make env.create`\n"

	if err := os.WriteFile(envFileName+".example", []byte(donttouch+envString), os.ModePerm); err != nil {
		return err
	}

	fmt.Println("env/config.go -> .env.example : file generated")
	return nil
}

func flatNestMap(nested map[string]any) map[string]any {
	result := make(map[string]any)

	for k, v := range nested {
		switch v := v.(type) {
		case map[string]any:
			subMap := flatNestMap(v)
			for subK, subV := range subMap {
				path := k + seperator + subK
				result[path] = subV
			}
		default:
			result[k] = v
		}

	}

	return result
}

func ParseAddress(raw string) (host, port string) {
	if i := strings.LastIndex(raw, ":"); i != -1 {
		return raw[:i], raw[i+1:]
	}
	return raw, ""
}
