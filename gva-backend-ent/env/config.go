package env

import (
	"fmt"
	"os"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

const (
	Dev  = "dev"
	Stag = "stag"
	Prod = "prod"
)

func (c *Config) IsProd() bool {
	return c.App.Env == Prod
}

func (c *Config) IsDev() bool {
	return c.App.Env == Dev
}

func (c *Config) IsStag() bool {
	return c.App.Env == Stag
}

type (
	app = struct {
		Name            string
		Address         string
		PrintRoutes     bool `mapstructure:"print_routes"`
		Env             string
		IdleTimeout     int64 `mapstructure:"idle_timeout"`
		ShutdownTimeout int64 `mapstructure:"shutdown_timeout"`
		TLS             struct {
			Auto     bool
			Enable   bool
			CertFile string `mapstructure:"cert_file"`
			KeyFile  string `mapstructure:"key_file"`
		}
	}
	db = struct {
		Mysql struct {
			DSN string
		}
		Redis struct {
			Address  string
			Password string
		}
	}
	seed struct {
		Enable     bool
		SuperAdmin struct {
			Username string
			Password string
		} `mapstructure:"super_admin"`
	}
	api struct {
		Web struct {
			Enable   bool
			Address  string
			BasePath string `mapstructure:"base_path"`
		}
		Admin struct {
			Enable   bool
			Address  string
			BasePath string `mapstructure:"base_path"`
		}
		Bot struct {
			Enable   bool
			Address  string
			BasePath string `mapstructure:"base_path"`
		}
	}
	logger = struct {
		TimeFormat string `mapstructure:"time_format"`
		Level      int8
		Prettier   bool
	}
	middleware = struct {
		Swagger struct {
			Enable bool
			Path   string
			Host   string
		}
		Compress struct {
			Enable bool
			Level  int
		}
		Monitor struct {
			Enable bool
			Path   string
		}
		Pprof struct {
			Enable bool
		}
		Limiter struct {
			Enable            bool
			Max               int
			ExpirationSeconds int64 `mapstructure:"expiration_seconds"`
		}
	}

	jwt struct {
		Secret string
	}
	password struct {
		HashCost int `mapstructure:"hash_cost"`
	}

	s3 struct {
		BucketName string `mapstructure:"bucket_name"`
		Address    string
		AccessKey  string `mapstructure:"access_key"`
		SecretKey  string `mapstructure:"secret_key"`
	}
)

type Config struct {
	API        api
	App        app
	DB         db
	Seed       seed
	Logger     logger
	Middleware middleware
	Jwt        jwt
	Password   password
	S3         s3
}

func ParseEnv(name string) (*Config, error) {
	var contents *Config

	file, err := os.ReadFile("./env/" + name + ".toml")
	if err != nil {
		return &Config{}, err
	}

	err = toml.Unmarshal(file, &contents)
	return contents, err
}

func ParsePath(path string) (*Config, error) {
	var contents *Config

	file, err := os.ReadFile(path)
	if err != nil {
		return &Config{}, err
	}

	err = toml.Unmarshal(file, &contents)
	return contents, err
}

func NewConfig() *Config {
	// generate .env if not exist
	config, err := ReadEnvOrGenerate()
	if err != nil {
		panic(fmt.Errorf("ReadEnvOrGenerate %v", err))
	}

	return config
}

func ParseAddress(raw string) (host, port string) {
	if i := strings.LastIndex(raw, ":"); i != -1 {
		return raw[:i], raw[i+1:]
	}
	return raw, ""
}
