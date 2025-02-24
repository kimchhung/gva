package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/gva/env"
	"github.com/gva/internal/logging"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var generateEnvCmd = &cobra.Command{
	Use:   "gen.env",
	Short: "push routes from json to database, delete and recreate base on file",
	Long:  `This command pulls routes from the database and performs necessary operations.`,
	Run: func(cmd *cobra.Command, args []string) {
		tomlfile := viper.New()
		tomlfile.SetConfigFile("./env/config.toml")
		if err := tomlfile.ReadInConfig(); err != nil { // Handle errors reading the config file
			panic(fmt.Errorf("fatal error config file: %w", err))
		}

		var config env.Config
		if err := tomlfile.Unmarshal(&config); err != nil {
			fmt.Printf("Error occurred during unmarshalling. Err: %s\n", err)
			os.Exit(1)
		}

		flated := flattenNestedMaps(tomlfile.AllSettings())
		groupPrefix := map[string]map[string]any{}

		for key, value := range flated {
			prefix := strings.Split(key, ".")[0]
			_, ok := groupPrefix[prefix]
			if !ok {
				groupPrefix[prefix] = map[string]any{}
			}

			groupPrefix[prefix][key] = value
		}

		envFormat := ""
		i := 0
		for prefix, keyvalue := range groupPrefix {
			if i > 0 {
				envFormat += "\n"
			}

			envFormat += fmt.Sprintf("# %v", prefix)
			envFormat += "\n"
			for key, value := range keyvalue {
				envFormat += strings.ToUpper(key)
				envFormat += "="
				envFormat += fmt.Sprintf("%v", value)
				envFormat += "\n"
			}
			i++
		}

		os.WriteFile(".env", []byte(envFormat), 0644)
		envfile := viper.NewWithOptions()
		envfile.AddConfigPath(".")
		envfile.SetConfigFile(".env")
		if err := envfile.ReadInConfig(); err != nil { // Handle errors reading the config file
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
		var envConfig env.Config
		envfile.Unmarshal(&envConfig)
		logging.Log(envConfig)
	},
}

func flattenNestedMaps(inputMap map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	for key, value := range inputMap {
		switch v := value.(type) {
		case map[string]interface{}:
			// If the value is another map, recursively flatten it
			nestedMap := flattenNestedMaps(v)
			for nestedKey, nestedValue := range nestedMap {
				result[key+"."+nestedKey] = nestedValue
			}
		default:
			// If the value is not a map, add it directly to the result map
			result[key] = value
		}
	}

	return result
}

func init() {
	rootCmd.AddCommand(generateEnvCmd)
}
