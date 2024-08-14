package env

import (
	"os"
	"strings"

	"github.com/pelletier/go-toml/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	Dev     = "dev"
	Staging = "staging"
	Prod    = "prod"
)

func (c *Config) IsProd() bool {
	return c.App.Env == Prod
}

func (c *Config) IsDev() bool {
	return c.App.Env == Dev
}

func (c *Config) IsStaging() bool {
	return c.App.Env == Staging
}

type (
	app = struct {
		Name            string
		Port            string
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
			Addr     string
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
			Port     string
			BasePath string `mapstructure:"base_path"`
		}
		Admin struct {
			Enable   bool
			Port     string
			BasePath string `mapstructure:"base_path"`
		}
		Bot struct {
			Enable   bool
			Port     string
			BasePath string `mapstructure:"base_path"`
		}
	}
	logger = struct {
		TimeFormat string `mapstructure:"time_format"`
		Level      zerolog.Level
		Prettier   bool
	}
	middleware = struct {
		Swagger struct {
			Enable bool
			Path   string
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
}

func ParseEnv(name string) (*Config, error) {
	var contents *Config

	file, err := os.ReadFile("./env/." + name + "_env.toml")
	if err != nil {
		return &Config{}, err
	}

	err = toml.Unmarshal(file, &contents)
	return contents, err
}

func NewConfig() *Config {
	filename := "dev"
	envName := os.Getenv("APP_ENV")

	switch envName {
	case "local", "dev", "staging", "prod":
		filename = envName
	}

	config, err := ParseEnv(filename)
	if err != nil {
		log.Panic().Err(err).Msg("failed to parse config")
	}

	return config
}

func ParseAddr(raw string) (host, port string) {
	if i := strings.LastIndex(raw, ":"); i != -1 {
		return raw[:i], raw[i+1:]
	}
	return raw, ""
}
