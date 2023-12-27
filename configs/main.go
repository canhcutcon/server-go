package configs

import (
	"fmt"
	"reflect"

	"github.com/spf13/viper"
)

// Config is the struct for the configuration of the application
type Config struct {
	Database struct {
		Uri string `mapstructure:"uri" default:"mongodb://localhost:27017"`
	} `mapstructure:"database"`
	Port string `mapstructure:"port" default:":8080"`
	Jwt  struct {
		Secrect string `mapstructure:"secrect" default:"secrect"`
		Expires int64  `mapstructure:"expires" default:"3600"`
	} `mapstructure:"jwt"`

	Documentation struct {
		Enable  bool   `mapstructure:"enable" default:"true"`
		Scheme  string `mapstructure:"schemes" default:"http"`
		Options struct {
			Info struct {
				Title  string `mapstructure:"title" default:"API Documentation"`
				Vesion string `mapstructure:"version" default:"1.0.0"`
			} `mapstructure:"info"`
		} `mapstructure:"options"`
		Groupinq          string `mapstructure:"grouping" default:"tags"`
		DocumentationPage bool   `mapstructure:"documentationPage" default:"true"`
	} `mapstructure:"documentation"`
	Env   string `mapstructure:"env" default:"development"`
	Redis struct {
		Prefix   string `mapstructure:"prefix" default:""`
		Host     string `mapstructure:"host" default:""`
		Port     string `mapstructure:"port" default:""`
		Password string `mapstructure:"password" default:""`
	} `mapstructure:"redis"`
	Email struct {
		Username string `mapstructure:"username" default:""`
		Password string `mapstructure:"password" default:""`
		APIKey   string `mapstructure:"api_key" default:""`
	} `mapstructure:"email"`
	PhoneOtp struct {
		BaseUrl   string `mapstructure:"base_url" default:""`
		APIKey    string `mapstructure:"api_key" default:""`
		MediaType string `mapstructure:"media_type" default:""`
	} `mapstructure:"phone_otp"`
	AWS struct {
		BucketName      string `mapstructure:"bucket_name" default:""`
		CloudFrontUrl   string `mapstructure:"cloud_front_url" default:""`
		Region          string `mapstructure:"region" default:""`
		AccessKey       string `mapstructure:"access_key_id" default:""`
		SecretAccessKey string `mapstructure:"secret_access_key" default:""`
	} `mapstructure:"aws"`
	GeoCode struct {
		RapidApiKey string `mapstructure:"rapid_api_key" default:""`
	} `mapstructure:"geocode"`
}

// Path: configs/main.go
var configuration Config // global variable to store the configuration

// GetConfig returns the configuration of the application
func GetConfig() Config {
	if configuration == (Config{}) { // if the configuration is empty, then load the configuration
		LoadAllConfigurations()
	}
	return configuration
}

func GetAllKeys() []string {
	var jsonKeys []string

	v := reflect.ValueOf(configuration)
	typeOfS := v.Type()
	for i := 0; i < v.NumField(); i++ {
		jsonKeys = append(jsonKeys, typeOfS.Field(i).Tag.Get("mapstructure"))
	}
	return jsonKeys
}

func LoadAllConfigurations() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config") // Register config file name (no extension)
	viper.SetConfigType("json")   // Look for specific type
	viper.ReadInConfig()

	key := GetAllKeys()
	for _, k := range key {
		// log.Info("Load config key::", k)
		fmt.Println("Load config key::", k)
		viper.SetDefault(k, viper.Get(k))
	}
	return viper.Unmarshal(&configuration) // Find and read the config file
}
