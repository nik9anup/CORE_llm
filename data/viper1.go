package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Initialize Viper and set a default value for "key"
func main() {
	viper.SetDefault("key", "defaultValue")
	fmt.Println(viper.GetString("key"))  // Output: defaultValue
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Read configuration from a JSON file
func main() {
	viper.SetConfigName("config") // Name of the config file (without extension)
	viper.SetConfigType("json")   // Type of the config file
	viper.AddConfigPath(".")      // Path to look for the config file

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	fmt.Println(viper.GetString("key"))
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Read configuration from a YAML file
func main() {
	viper.SetConfigName("config") // Name of the config file (without extension)
	viper.SetConfigType("yaml")   // Type of the config file
	viper.AddConfigPath(".")      // Path to look for the config file

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	fmt.Println(viper.GetString("key"))
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Read configuration from a TOML file
func main() {
	viper.SetConfigName("config") // Name of the config file (without extension)
	viper.SetConfigType("toml")   // Type of the config file
	viper.AddConfigPath(".")      // Path to look for the config file

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	fmt.Println(viper.GetString("key"))
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

// Read configuration from environment variables
func main() {
	_ = os.Setenv("APP_KEY", "envValue")
	viper.BindEnv("key", "APP_KEY")

	fmt.Println(viper.GetString("key"))  // Output: envValue
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Set configuration file path explicitly
func main() {
	viper.SetConfigFile("./config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	fmt.Println(viper.GetString("key"))
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Check if a configuration key exists
func main() {
	viper.Set("key", "value")
	if viper.IsSet("key") {
		fmt.Println("Key exists")
	} else {
		fmt.Println("Key does not exist")
	}
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

// Watch for changes in the configuration file
func main() {
	viper.SetConfigFile("./config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	for {
		time.Sleep(time.Second)
	}
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Unmarshal configuration into a struct
type Config struct {
	Key string
}

func main() {
	viper.SetConfigFile("./config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	fmt.Println(config.Key)
}





package main

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
)

// Read configuration from command-line flags
func main() {
	key := flag.String("key", "default", "The key value")
	flag.Parse()

	viper.BindPFlag("key", flag.Lookup("key"))

	fmt.Println(viper.GetString("key"))  // Output: (value provided via --key)
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Write configuration to a file
func main() {
	viper.Set("key", "newValue")
	viper.SetConfigFile("./config.yaml")

	if err := viper.WriteConfig(); err != nil {
		panic(err)
	}

	fmt.Println("Config written to file")
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Merge configuration from multiple files
func main() {
	viper.SetConfigFile("./config1.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	viper.SetConfigFile("./config2.yaml")
	if err := viper.MergeInConfig(); err != nil {
		panic(err)
	}

	fmt.Println(viper.GetString("key"))
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Retrieve all settings as a map
func main() {
	viper.Set("key1", "value1")
	viper.Set("key2", "value2")

	settings := viper.AllSettings()
	fmt.Println(settings)
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Get a configuration value as a boolean
func main() {
	viper.Set("key", true)

	value := viper.GetBool("key")
	fmt.Println(value)  // Output: true
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Get a configuration value as an integer
func main() {
	viper.Set("key", 42)

	value := viper.GetInt("key")
	fmt.Println(value)  // Output: 42
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Get a configuration value as a float
func main() {
	viper.Set("key", 42.42)

	value := viper.GetFloat64("key")
	fmt.Println(value)  // Output: 42.42
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

// Get a configuration value as a time duration
func main() {
	viper.Set("key", "1h")

	value := viper.GetDuration("key")
	fmt.Println(value)  // Output: 1h0m0s
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Get a configuration value as a slice
func main() {
	viper.Set("key", []string{"value1", "value2"})

	value := viper.GetStringSlice("key")
	fmt.Println(value)  // Output: [value1 value2]
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Get a configuration value as a map
func main() {
	viper.Set("key", map[string]string{"subkey": "subvalue"})

	value := viper.GetStringMapString("key")
	fmt.Println(value)  // Output: map[subkey:subvalue]
}





package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
)

// Set configuration from a JSON string
func main() {
	jsonStr := `{"key": "jsonValue"}`
	var data map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &data)
	viper.MergeConfigMap(data)

	fmt.Println(viper.GetString("key"))  // Output: jsonValue
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// Set configuration from a YAML string
func main() {
	yamlStr := `key: yamlValue`
	var data map[string]interface{}
	yaml.Unmarshal([]byte(yamlStr), &data)
	viper.MergeConfigMap(data)

	fmt.Println(viper.GetString("key"))  // Output: yamlValue
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

// Set configuration from environment variables with a prefix
func main() {
	_ = os.Setenv("APP_KEY", "envValue")
	viper.SetEnvPrefix("app")
	viper.BindEnv("key")

	fmt.Println(viper.GetString("key"))  // Output: envValue
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

// Set configuration from environment variables automatically
func main() {
	_ = os.Setenv("APP_KEY", "envValue")
	viper.AutomaticEnv()

	fmt.Println(viper.GetString("APP_KEY"))  // Output: envValue
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Set an alias for a configuration key
func main() {
	viper.Set("originalKey", "value")
	viper.RegisterAlias("aliasKey", "originalKey")

	fmt.Println(viper.GetString("aliasKey"))  // Output: value
}





package main

import (
	"fmt"
	"strings"
	"github.com/spf13/viper"
)

// Use a custom key replacer
func main() {
	viper.Set("custom.key", "value")
	viper.SetKeyDelimiter(".")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	fmt.Println(viper.GetString("custom_key"))  // Output: value
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Handle errors when reading the configuration file
func main() {
	viper.SetConfigFile("./config.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file:", err)
	} else {
		fmt.Println("Config file read successfully")
	}
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Get a configuration value with a fallback
func main() {
	viper.Set("key", "value")

	value := viper.GetString("missingKey")
	if value == "" {
		value = "fallbackValue"
	}

	fmt.Println(value)  // Output: fallbackValue
}





package main

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
)

// Bind a configuration key to a command-line flag
func main() {
	key := flag.String("key", "default", "The key value")
	flag.Parse()

	viper.BindPFlag("key", flag.Lookup("key"))

	fmt.Println(viper.GetString("key"))  // Output: (value provided via --key)
}





package main

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
)

// Bind all flags to Viper
func main() {
	key := flag.String("key", "default", "The key value")
	flag.Parse()

	viper.BindPFlags(flag.CommandLine)

	fmt.Println(viper.GetString("key"))  // Output: (value provided via --key)
}





package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Use Cobra command with Viper
func main() {
	var key string

	var rootCmd = &cobra.Command{
		Use:   "app",
		Short: "A brief description of your application",
		Run: func(cmd *cobra.Command, args []string) {
			viper.BindPFlag("key", cmd.Flags().Lookup("key"))
			fmt.Println(viper.GetString("key"))
		},
	}

	rootCmd.Flags().StringVarP(&key, "key", "k", "default", "The key value")
	cobra.CheckErr(rootCmd.Execute())
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

// Read configuration from a remote source
func main() {
	viper.AddRemoteProvider("consul", "localhost:8500", "path/to/config")
	viper.SetConfigType("json")  // Specify the type of configuration to retrieve

	err := viper.ReadRemoteConfig()
	if err != nil {
		log.Fatalf("Failed to read remote config: %v", err)
	}

	fmt.Println(viper.GetString("key"))
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Add multiple paths for configuration files
func main() {
	viper.AddConfigPath("/etc/appname/")
	viper.AddConfigPath("$HOME/.appname")
	viper.AddConfigPath(".")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	fmt.Println(viper.GetString("key"))
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

// Override configuration with environment variables
func main() {
	_ = os.Setenv("APP_KEY", "envValue")
	viper.BindEnv("key", "APP_KEY")
	viper.Set("key", "configValue")

	fmt.Println(viper.GetString("key"))  // Output: envValue
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Merge multiple configuration maps
func main() {
	viper.Set("key1", "value1")

	moreConfig := map[string]interface{}{
		"key2": "value2",
	}

	viper.MergeConfigMap(moreConfig)

	fmt.Println(viper.GetString("key1"))  // Output: value1
	fmt.Println(viper.GetString("key2"))  // Output: value2
}





package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"time"
)

// Watch environment variables for changes
func main() {
	_ = os.Setenv("APP_KEY", "envValue")
	viper.BindEnv("key", "APP_KEY")

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Environment variable changed")
	})

	for {
		time.Sleep(time.Second)
	}
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Custom unmarshaler function
type Custom struct {
	Key string
}

func (c *Custom) UnmarshalText(text []byte) error {
	c.Key = string(text) + "_unmarshaled"
	return nil
}

// Use a custom unmarshaler
func main() {
	viper.Set("key", "value")

	var custom Custom
	viper.UnmarshalKey("key", &custom)

	fmt.Println(custom.Key)  // Output: value_unmarshaled
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Write a sub-configuration to a file
func main() {
	viper.Set("sub.key", "subValue")
	viper.SetConfigFile("./subconfig.yaml")

	if err := viper.WriteConfigAs("./subconfig.yaml"); err != nil {
		panic(err)
	}

	fmt.Println("Sub-configuration written to file")
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Retrieve all configuration keys
func main() {
	viper.Set("key1", "value1")
	viper.Set("key2", "value2")

	keys := viper.AllKeys()
	fmt.Println(keys)  // Output: [key1 key2]
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Get a nested configuration value
func main() {
	viper.Set("parent.child.key", "nestedValue")

	value := viper.GetString("parent.child.key")
	fmt.Println(value)  // Output: nestedValue
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Unmarshal configuration into nested structs
type Config struct {
	Parent struct {
		Child struct {
			Key string
		}
	}
}

func main() {
	viper.Set("parent.child.key", "nestedValue")

	var config Config
	viper.Unmarshal(&config)

	fmt.Println(config.Parent.Child.Key)  // Output: nestedValue
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

// Bind a configuration key to the content of a file
func main() {
	_ = os.WriteFile("keyfile", []byte("fileContent"), 0644)
	viper.BindConfigFile("key", "keyfile")

	fmt.Println(viper.GetString("key"))  // Output: fileContent
}





package main

import (
	"testing"
	"github.com/spf13/viper"
)

// Use Viper in tests
func TestViper(t *testing.T) {
	viper.Set("key", "testValue")

	if viper.GetString("key") != "testValue" {
		t.Error("Expected testValue")
	}
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Use Viper with JSON tags
type Config struct {
	Key string `json:"key"`
}

func main() {
	viper.SetConfigType("json")
	viper.ReadConfig(strings.NewReader(`{"key": "jsonValue"}`))

	var config Config
	viper.Unmarshal(&config)

	fmt.Println(config.Key)  // Output: jsonValue
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Use Viper with custom tags
type Config struct {
	Key string `mapstructure:"key"`
}

func main() {
	viper.Set("key", "value")

	var config Config
	viper.Unmarshal(&config)

	fmt.Println(config.Key)  // Output: value
}





package main

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
)

// Use Viper with command-line and config file
func main() {
	key := flag.String("key", "default", "The key value")
	flag.Parse()

	viper.SetConfigFile("./config.yaml")
	viper.ReadInConfig()
	viper.BindPFlag("key", flag.Lookup("key"))

	fmt.Println(viper.GetString("key"))
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Set configuration value from raw bytes
func main() {
	viper.Set("key", []byte("rawBytesValue"))

	value := viper.Get("key").([]byte)
	fmt.Println(string(value))  // Output: rawBytesValue
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Reset Viper configuration
func main() {
	viper.Set("key", "value")
	viper.Reset()

	if viper.IsSet("key") {
		fmt.Println("Key still exists")
	} else {
		fmt.Println("Key has been reset")  // Output: Key has been reset
	}
}





package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Log configuration changes
func main() {
	viper.SetConfigFile("./config.yaml")
	viper.ReadInConfig()

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file changed: %s\n", e.Name)
	})

	select {}
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Ignore unexported struct fields
type Config struct {
	Exported   string
	unexported string
}

func main() {
	viper.Set("Exported", "exportedValue")
	viper.Set("unexported", "unexportedValue")

	var config Config
	viper.Unmarshal(&config)

	fmt.Println(config.Exported)    // Output: exportedValue
	fmt.Println(config.unexported)  // Output:
}





package main

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
)

// Load configuration from a byte buffer
func main() {
	var buffer bytes.Buffer
	buffer.WriteString(`key: bufferValue`)

	viper.SetConfigType("yaml")
	viper.ReadConfig(&buffer)

	fmt.Println(viper.GetString("key"))  // Output: bufferValue
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

// Use environment variables with a custom delimiter
func main() {
	_ = os.Setenv("APP__KEY", "envValue")
	viper.SetEnvKeyReplacer(strings.NewReplacer("__", "."))
	viper.AutomaticEnv()

	fmt.Println(viper.GetString("app.key"))  // Output: envValue
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Ignore missing configuration file
func main() {
	viper.SetConfigFile("./missing-config.yaml")
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Config file found and read successfully")
	} else {
		fmt.Println("Config file not found or could not be read, ignoring...")
	}
}





package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Watch configuration files for changes
func main() {
	viper.SetConfigFile("./config.yaml")
	viper.ReadInConfig()

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	select {}  // Block forever
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Get all configuration settings as a map
func main() {
	viper.Set("key1", "value1")
	viper.Set("key2", "value2")

	settings := viper.AllSettings()
	fmt.Println(settings)  // Output: map[key1:value1 key2:value2]
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Set configuration from JSON file
func main() {
	viper.SetConfigFile("./config.json")
	viper.ReadInConfig()

	fmt.Println(viper.GetString("key"))  // Output: (value from config.json)
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Set configuration from YAML file
func main() {
	viper.SetConfigFile("./config.yaml")
	viper.ReadInConfig()

	fmt.Println(viper.GetString("key"))  // Output: (value from config.yaml)
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Use nested struct tags with Viper
type Config struct {
	Parent struct {
		Child struct {
			Key string `mapstructure:"key"`
		} `mapstructure:"child"`
	} `mapstructure:"parent"`
}

func main() {
	viper.Set("parent.child.key", "nestedValue")

	var config Config
	viper.Unmarshal(&config)

	fmt.Println(config.Parent.Child.Key)  // Output: nestedValue
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Check if configuration key exists
func main() {
	viper.Set("key", "value")

	if viper.IsSet("key") {
		fmt.Println("Key exists")  // Output: Key exists
	} else {
		fmt.Println("Key does not exist")
	}
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Get a configuration value as an integer
func main() {
	viper.Set("key", 123)

	value := viper.GetInt("key")
	fmt.Println(value)  // Output: 123
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Get a configuration value as a boolean
func main() {
	viper.Set("key", true)

	value := viper.GetBool("key")
	fmt.Println(value)  // Output: true
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Set configuration from TOML file
func main() {
	viper.SetConfigFile("./config.toml")
	viper.ReadInConfig()

	fmt.Println(viper.GetString("key"))  // Output: (value from config.toml)
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

// Set configuration from a remote source with custom type
func main() {
	viper.AddRemoteProvider("etcd", "localhost:2379", "/config")
	viper.SetConfigType("yaml")

	err := viper.ReadRemoteConfig()
	if err != nil {
		log.Fatalf("Failed to read remote config: %v", err)
	}

	fmt.Println(viper.GetString("key"))
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

// Use key replacer for nested configuration
func main() {
	viper.Set("parent.child.key", "nestedValue")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.AutomaticEnv()

	fmt.Println(viper.GetString("parent__child__key"))  // Output: nestedValue
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Set configuration with function
func main() {
	viper.Set("key", func() string { return "functionValue" }())

	fmt.Println(viper.GetString("key"))  // Output: functionValue
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Use configuration sub-tree
func main() {
	viper.Set("parent.child.key", "nestedValue")

	sub := viper.Sub("parent.child")
	fmt.Println(sub.GetString("key"))  // Output: nestedValue
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Bind to configuration files with different formats
func main() {
	viper.SetConfigFile("./config.json")
	viper.ReadInConfig()
	viper.SetConfigFile("./config.yaml")
	viper.MergeInConfig()

	fmt.Println(viper.GetString("key"))  // Output: (value from either config.json or config.yaml)
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Get configuration value with default
func main() {
	viper.SetDefault("key", "defaultValue")

	fmt.Println(viper.GetString("key"))  // Output: defaultValue
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Set configuration key case insensitive
func main() {
	viper.Set("KEY", "value")

	fmt.Println(viper.GetString("key"))  // Output: value
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Unmarshal configuration into multiple structs
type Config1 struct {
	Key1 string
}

type Config2 struct {
	Key2 string
}

func main() {
	viper.Set("key1", "value1")
	viper.Set("key2", "value2")

	var config1 Config1
	var config2 Config2
	viper.Unmarshal(&config1)
	viper.Unmarshal(&config2)

	fmt.Println(config1.Key1)  // Output: value1
	fmt.Println(config2.Key2)  // Output: value2
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

// Unmarshal configuration from reader
func main() {
	reader := strings.NewReader(`key: value`)

	viper.SetConfigType("yaml")
	viper.ReadConfig(reader)

	fmt.Println(viper.GetString("key"))  // Output: value
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Set configuration using a struct
type Config struct {
	Key string
}

func main() {
	config := Config{Key: "value"}
	viper.Set("config", config)

	fmt.Println(viper.GetStringMap("config"))  // Output: map[Key:value]
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Register custom config type
func main() {
	viper.SetConfigType("custom")

	fmt.Println(viper.GetString("key"))  // Output: (based on custom type handling)
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Handle configuration errors gracefully
func main() {
	viper.SetConfigFile("./invalid-config.yaml")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("Error reading config:", err)
	} else {
		fmt.Println("Config loaded successfully")
	}
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

// Get configuration value as a duration
func main() {
	viper.Set("timeout", "5s")

	duration := viper.GetDuration("timeout")
	fmt.Println(duration)  // Output: 5s
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Get configuration value as a float
func main() {
	viper.Set("key", 123.45)

	value := viper.GetFloat64("key")
	fmt.Println(value)  // Output: 123.45
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Bind configuration to struct with default values
type Config struct {
	Key1 string
	Key2 int
}

func main() {
	viper.SetDefault("key1", "defaultValue")
	viper.SetDefault("key2", 42)

	var config Config
	viper.Unmarshal(&config)

	fmt.Println(config.Key1)  // Output: defaultValue
	fmt.Println(config.Key2)  // Output: 42
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Read configuration with optional key
func main() {
	value := viper.GetString("optionalKey")
	if value == "" {
		fmt.Println("Key not set, using default value")
	} else {
		fmt.Println("Key value:", value)
	}
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Override configuration values
func main() {
	viper.Set("key", "initialValue")
	viper.Set("key", "overriddenValue")

	fmt.Println(viper.GetString("key"))  // Output: overriddenValue
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

// Use environment variables with Viper
func main() {
	_ = os.Setenv("KEY", "envValue")
	viper.BindEnv("key")

	fmt.Println(viper.GetString("key"))  // Output: envValue
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Check for unused keys
type Config struct {
	Key1 string
}

func main() {
	viper.Set("key1", "value1")
	viper.Set("unusedKey", "value")

	var config Config
	viper.Unmarshal(&config)

	for _, key := range viper.AllKeys() {
		if !viper.InConfig(key) {
			fmt.Println("Unused key:", key)  // Output: Unused key: unusedKey
		}
	}
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Get configuration value from nested maps
func main() {
	viper.Set("parent.child.key", "nestedValue")

	value := viper.GetString("parent.child.key")
	fmt.Println(value)  // Output: nestedValue
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
	"strconv"
)

// Set configuration value with custom type conversion
func main() {
	viper.Set("key", "123")

	value, _ := strconv.Atoi(viper.GetString("key"))
	fmt.Println(value)  // Output: 123
}





package main

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

// Validate configuration
func main() {
	viper.Set("key", "value")

	if err := validateConfig(); err != nil {
		fmt.Println("Validation error:", err)
	} else {
		fmt.Println("Configuration is valid")
	}
}

func validateConfig() error {
	if viper.GetString("key") == "" {
		return errors.New("key is required")
	}
	return nil
}





package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
)

// Marshal configuration into JSON
func main() {
	viper.Set("key", "value")

	configJSON, _ := json.Marshal(viper.AllSettings())
	fmt.Println(string(configJSON))  // Output: {"key":"value"}
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Merge multiple configuration maps
func main() {
	viper.Set("key1", "value1")

	otherConfig := map[string]interface{}{"key2": "value2"}
	viper.MergeConfigMap(otherConfig)

	fmt.Println(viper.GetString("key1"))  // Output: value1
	fmt.Println(viper.GetString("key2"))  // Output: value2
}





package main

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
)

// Use Viper with FlagSet
func main() {
	fs := flag.NewFlagSet("example", flag.ContinueOnError)
	fs.String("key", "default", "Description of the key")

	viper.BindFlagSet(fs)

	fs.Parse([]string{"--key", "flagValue"})
	fmt.Println(viper.GetString("key"))  // Output: flagValue
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Add custom configuration file paths
func main() {
	viper.AddConfigPath("./config")
	viper.AddConfigPath("./settings")

	viper.SetConfigName("config")
	viper.ReadInConfig()

	fmt.Println(viper.GetString("key"))  // Output: (value from one of the config paths)
}





package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Handle configuration reload errors
func main() {
	viper.SetConfigFile("./config.yaml")
	viper.ReadInConfig()

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("Error reloading config:", err)
		}
	})

	select {}
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Load configuration with custom name
func main() {
	viper.SetConfigName("custom-config")
	viper.AddConfigPath("./")
	viper.ReadInConfig()

	fmt.Println(viper.GetString("key"))  // Output: (value from custom-config file)
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

// Get configuration value as time.Time
func main() {
	viper.Set("timestamp", "2022-01-01T00:00:00Z")

	timestamp, _ := time.Parse(time.RFC3339, viper.GetString("timestamp"))
	fmt.Println(timestamp)  // Output: 2022-01-01 00:00:00 +0000 UTC
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Set configuration value with slice
func main() {
	viper.Set("key", []string{"value1", "value2"})

	values := viper.GetStringSlice("key")
	fmt.Println(values)  // Output: [value1 value2]
}





package main

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
)

// Override configuration with flags
func main() {
	fs := flag.NewFlagSet("example", flag.ContinueOnError)
	fs.String("key", "defaultValue", "Description of the key")

	viper.BindFlagSet(fs)
	fs.Parse([]string{"--key", "flagValue"})

	fmt.Println(viper.GetString("key"))  // Output: flagValue
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Handle configuration loading errors
func main() {
	viper.SetConfigFile("./invalid-config.yaml")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("Error reading config:", err)
	} else {
		fmt.Println("Config loaded successfully")
	}
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

// Use configuration with time durations
func main() {
	viper.Set("timeout", "5s")

	timeout := viper.GetDuration("timeout")
	fmt.Println(timeout)  // Output: 5s
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

// Log configuration load errors
func main() {
	viper.SetConfigFile("./invalid-config.yaml")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	} else {
		fmt.Println("Config loaded successfully")
	}
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Set configuration value as map
func main() {
	viper.Set("key", map[string]interface{}{"subkey": "value"})

	value := viper.GetStringMap("key")
	fmt.Println(value)  // Output: map[subkey:value]
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Use configuration with custom file name
func main() {
	viper.SetConfigName("my-config")
	viper.AddConfigPath("./")
	viper.ReadInConfig()

	fmt.Println(viper.GetString("key"))  // Output: (value from my-config file)
}





package main

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"github.com/spf13/viper"
)

// Marshal configuration into TOML
func main() {
	viper.Set("key", "value")

	configTOML, _ := toml.Marshal(viper.AllSettings())
	fmt.Println(string(configTOML))  // Output: (TOML representation of configuration)
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Get configuration value as interface{}
func main() {
	viper.Set("key", "value")

	value := viper.Get("key")
	fmt.Println(value)  // Output: value
}





package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Use Viper for command-line application
func main() {
	var rootCmd = &cobra.Command{
		Use: "app",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Config key value:", viper.GetString("key"))
		},
	}

	rootCmd.PersistentFlags().String("key", "defaultValue", "Description of the key")
	viper.BindPFlag("key", rootCmd.PersistentFlags().Lookup("key"))

	rootCmd.Execute()
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

// Use multiple configuration providers
func main() {
	viper.SetConfigFile("./config.yaml")
	viper.ReadInConfig()
	viper.AutomaticEnv()

	_ = os.Setenv("KEY", "envValue")
	viper.BindEnv("key")

	fmt.Println(viper.GetString("key"))  // Output: (value from env or file)
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

// Save configuration to file
func main() {
	viper.Set("key", "value")

	file, _ := os.Create("./saved-config.yaml")
	defer file.Close()

	viper.WriteConfigAs(file.Name())
	fmt.Println("Config saved to", file.Name())
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

// Use custom delimiter for environment variables
func main() {
	_ = os.Setenv("APP__KEY", "envValue")
	viper.SetEnvKeyReplacer(strings.NewReplacer("__", "."))
	viper.AutomaticEnv()

	fmt.Println(viper.GetString("app.key"))  // Output: envValue
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Handle errors in configuration loading
func main() {
	viper.SetConfigFile("./missing-config.yaml")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("Error reading config:", err)
	} else {
		fmt.Println("Config loaded successfully")
	}
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

// Dynamically update configuration
func main() {
	viper.Set("key", "initialValue")

	go func() {
		time.Sleep(2 * time.Second)
		viper.Set("key", "updatedValue")
	}()

	for {
		fmt.Println(viper.GetString("key"))  // Output: initialValue (then updates to updatedValue)
		time.Sleep(1 * time.Second)
	}
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Use configuration with aliases
func main() {
	viper.Set("key", "value")
	viper.RegisterAlias("aliasKey", "key")

	fmt.Println(viper.GetString("aliasKey"))  // Output: value
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

// Read configuration from multiple sources
func main() {
	viper.SetConfigFile("./config.yaml")
	viper.ReadInConfig()
	viper.AutomaticEnv()

	_ = os.Setenv("KEY", "envValue")
	viper.BindEnv("key")

	fmt.Println(viper.GetString("key"))  // Output: (value from env or file)
}





package main

import (
	"encoding/xml"
	"fmt"
	"github.com/spf13/viper"
)

// Marshal configuration into XML
func main() {
	viper.Set("key", "value")

	configXML, _ := xml.Marshal(viper.AllSettings())
	fmt.Println(string(configXML))  // Output: (XML representation of configuration)
}





package main

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
)

// Bind command-line flags to Viper
func main() {
	fs := flag.NewFlagSet("example", flag.ContinueOnError)
	fs.String("key", "defaultValue", "Description of the key")

	viper.BindFlagSet(fs)
	fs.Parse([]string{"--key", "flagValue"})

	fmt.Println(viper.GetString("key"))  // Output: flagValue
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Handle missing configuration file
func main() {
	viper.SetConfigFile("./missing-config.yaml")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("Configuration file not found, proceeding with defaults")
	} else {
		fmt.Println("Config loaded successfully")
	}
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
	"net/http"
	"io/ioutil"
)

// Load configuration from a remote source
func main() {
	resp, err := http.Get("https://example.com/config.yaml")
	if err != nil {
		fmt.Println("Error fetching config:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading config:", err)
		return
	}

	viper.SetConfigType("yaml")
	viper.ReadConfig(bytes.NewBuffer(body))

	fmt.Println(viper.GetString("key"))  // Output: (value from remote config)
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"
)

// Use custom file watcher with Viper
func main() {
	viper.SetConfigFile("./config.yaml")
	viper.ReadInConfig()

	watcher, _ := fsnotify.NewWatcher()
	defer watcher.Close()

	watcher.Add("./config.yaml")

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					viper.ReadInConfig()
					fmt.Println("Config file updated:", viper.AllSettings())
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("Watcher error:", err)
			}
		}
	}()

	select {}
}





package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
)

// Get configuration value as JSON
func main() {
	viper.Set("key", "value")

	valueJSON, _ := json.Marshal(viper.AllSettings())
	fmt.Println(string(valueJSON))  // Output: {"key":"value"}
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Integrate Viper with custom CLI framework
func main() {
	// Example custom CLI framework setup
	// (Custom logic to handle CLI arguments)

	viper.Set("key", "valueFromCLI")

	fmt.Println(viper.GetString("key"))  // Output: valueFromCLI
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Use Viper with custom encoding
func main() {
	viper.SetConfigType("custom")

	// Example: Custom encoding handler (details omitted)
	viper.ReadConfig(customReader)

	fmt.Println(viper.GetString("key"))  // Output: (based on custom encoding)
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

// Dynamic configuration reloading
func main() {
	viper.SetConfigFile("./config.yaml")
	viper.ReadInConfig()

	go func() {
		for {
			time.Sleep(10 * time.Second)
			viper.ReadInConfig()
			fmt.Println("Config reloaded:", viper.AllSettings())
		}
	}()

	select {}
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Custom configuration decoders
func main() {
	viper.SetConfigType("custom")

	// Example: Custom decoder (details omitted)
	viper.ReadConfig(customDecoder)

	fmt.Println(viper.GetString("key"))  // Output: (based on custom decoder)
}





package main

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
)

// Load configuration from multiple remote sources
func main() {
	resp1, _ := http.Get("https://example.com/config1.yaml")
	body1, _ := ioutil.ReadAll(resp1.Body)
	defer resp1.Body.Close()

	resp2, _ := http.Get("https://example.com/config2.yaml")
	body2, _ := ioutil.ReadAll(resp2.Body)
	defer resp2.Body.Close()

	viper.SetConfigType("yaml")
	viper.ReadConfig(bytes.NewBuffer(body1))
	viper.MergeConfig(bytes.NewBuffer(body2))

	fmt.Println(viper.AllSettings())  // Output: (combined values from both remote configs)
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Custom error handling for Viper
func main() {
	viper.SetConfigFile("./config.yaml")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("Custom error handling:", err)
		// Custom handling logic
	} else {
		fmt.Println("Config loaded successfully")
	}
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
	"net/http"
)

// Handle configuration loading in a web server
func main() {
	viper.SetConfigFile("./config.yaml")
	viper.ReadInConfig()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Config key value: %s", viper.GetString("key"))
	})

	http.ListenAndServe(":8080", nil)
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Use Viper with custom configuration formats
func main() {
	viper.SetConfigType("custom")

	// Example: Custom configuration format (details omitted)
	viper.ReadConfig(customFormatReader)

	fmt.Println(viper.GetString("key"))  // Output: (based on custom format)
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

// Monitor configuration changes from multiple sources
func main() {
	viper.SetConfigFile("./config.yaml")
	viper.ReadInConfig()

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	go func() {
		for {
			time.Sleep(10 * time.Second)
			// Check remote source and merge changes
			viper.ReadRemoteConfig() // Example (details omitted)
			fmt.Println("Remote config reloaded:", viper.AllSettings())
		}
	}()

	select {}
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Serialize configuration to custom format
func main() {
	viper.Set("key", "value")

	// Example: Serialize to custom format (details omitted)
	customConfig := serializeToCustomFormat(viper.AllSettings())

	fmt.Println(customConfig)
}





package main

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
)

// Load configuration with custom validators
func main() {
	viper.SetConfigFile("./config.yaml")
	err := viper.ReadInConfig()

	if err == nil {
		err = validateConfig()
	}

	if err != nil {
		fmt.Println("Configuration error:", err)
	} else {
		fmt.Println("Configuration is valid")
	}
}

func validateConfig() error {
	if viper.GetString("key") == "" {
		return errors.New("key is required")
	}
	return nil
}





package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/aws/aws-lambda-go/lambda"
)

// Load configuration in serverless environment
func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context) (string, error) {
	viper.SetConfigFile("/var/task/config.yaml")  // Lambda function file system
	viper.ReadInConfig()

	return fmt.Sprintf("Config key value: %s", viper.GetString("key")), nil
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Use Viper with custom configuration loader
func main() {
	viper.SetConfigType("custom")

	// Example: Custom configuration loader (details omitted)
	viper.ReadConfig(customLoader)

	fmt.Println(viper.GetString("key"))  // Output: (based on custom loader)
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Load configuration from multiple files
func main() {
	viper.SetConfigFile("./config1.yaml")
	viper.ReadInConfig()
	viper.MergeInConfig("./config2.yaml")

	fmt.Println(viper.AllSettings())  // Output: (combined values from both files)
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
	// Assuming some encryption package
	// "github.com/example/encryption"
)

// Secure configuration with encryption
func main() {
	encryptedValue := encrypt("sensitiveValue")

	viper.Set("key", encryptedValue)

	decryptedValue := decrypt(viper.GetString("key"))
	fmt.Println("Decrypted value:", decryptedValue)  // Output: sensitiveValue
}

func encrypt(value string) string {
	// Encryption logic (details omitted)
	return value
}

func decrypt(value string) string {
	// Decryption logic (details omitted)
	return value
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

// Monitor configuration changes in Kubernetes
func main() {
	viper.SetConfigFile("/etc/config/config.yaml")
	viper.ReadInConfig()

	go func() {
		for {
			time.Sleep(10 * time.Second)
			viper.ReadInConfig()
			fmt.Println("Config reloaded:", viper.AllSettings())
		}
	}()

	select {}
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Use Viper with configuration versioning
func main() {
	viper.SetConfigFile("./config-v1.yaml")
	viper.ReadInConfig()

	version := viper.GetString("version")
	if version != "1.0" {
		fmt.Println("Unsupported configuration version")
	} else {
		fmt.Println("Config loaded successfully")
	}
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

// Use Viper with custom loggers
func main() {
	viper.SetConfigFile("./config.yaml")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Custom logger: Error reading config: %v", err)
	} else {
		fmt.Println("Config loaded successfully")
	}
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Handle nested configuration structures
func main() {
	viper.Set("nested.key", "value")

	fmt.Println(viper.GetString("nested.key"))  // Output: value
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
	// Assuming some encryption package
	// "github.com/example/encryption"
)

// Use Viper with custom encryption
func main() {
	viper.Set("key", encrypt("sensitiveValue"))

	value := decrypt(viper.GetString("key"))
	fmt.Println("Decrypted value:", value)  // Output: sensitiveValue
}

func encrypt(value string) string {
	// Encryption logic (details omitted)
	return value
}

func decrypt(value string) string {
	// Decryption logic (details omitted)
	return value
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"
)

// Load configuration with automatic reload
func main() {
	viper.SetConfigFile("./config.yaml")
	viper.ReadInConfig()

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	select {}
}





package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Handle configuration in microservices
func main() {
	viper.SetConfigFile("./service-config.yaml")
	viper.ReadInConfig()

	fmt.Println("Service config key value:", viper.GetString("serviceKey"))
}
