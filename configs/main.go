package configs

import (
	"reflect"

	"github.com/spf13/viper"
)

// Config is the struct for the configuration of the application
type Config struct {
	DatabaseUri string `mapstructure:"database.uri" default:"mongodb://localhost:27017"`
	Port        string `mapstructure:"port" default:":8080"`
	Jwt         struct {
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
	}
	PhoneOtp struct {
		BaseUrl   string `mapstructure:"base_url" default:""`
		APIKey    string `mapstructure:"api_key" default:""`
		MediaType string `mapstructure:"media_type" default:""`
	}
	AWS struct {
		BucketName      string `mapstructure:"bucket_name" default:""`
		CloudFrontUrl   string `mapstructure:"cloud_front_url" default:""`
		Region          string `mapstructure:"region" default:""`
		AccessKey       string `mapstructure:"access_key_id" default:""`
		SecretAccessKey string `mapstructure:"secret_access_key" default:""`
	}
	GeoCode struct {
		RapidApiKey string `mapstructure:"rapid_api_key" default:""`
	} `mapstructure:"geocode"`
}

// Path: configs/main.go
var configuration Config // global variable to store the configuration

func init() {
	// initialize the configuration
	configuration = Config{}

	// load the configuration from environment variables
	// err := envconfig.Process("", &configuration)
}

// GetConfig returns the configuration of the application
func GetConfig() Config {
	return configuration
}

func LoadAllConfigurations() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.ReadInConfig()

	config := Config{}
	key := reflect.TypeOf(config)

	for i := 0; i < key.NumField(); i++ {
		field := key.Field(i)
		viper.BindEnv(field.Tag.Get("mapstructure"))
	}
}
