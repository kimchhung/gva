package config

import (
	"os"
	"strings"

	"github.com/pelletier/go-toml/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type (
	app = struct {
		Name            string `toml:"name"`
		Port            string `toml:"port"`
		PrintRoutes     bool   `toml:"print_routes"`
		Production      bool   `toml:"production"`
		IdleTimeout     int64  `toml:"idle_timeout"`
		ShutdownTimeout int64  `toml:"shutdown_timeout"`
		TLS             struct {
			Auto     bool
			Enable   bool
			CertFile string `toml:"cert_file"`
			KeyFile  string `toml:"key_file"`
		}
	}
	db = struct {
		Mysql struct {
			DSN string `toml:"dsn"`
		}
	}
	seed struct {
		Enable     bool
		SuperAdmin struct {
			Username string `toml:"username"`
			Password string `toml:"password"`
		} `toml:"super_admin"`
	}
	api struct {
		Web struct {
			Enable   bool
			Port     string
			BasePath string `toml:"base_path"`
		}
		Admin struct {
			Enable   bool
			Port     string
			BasePath string `toml:"base_path"`
		}
	}
	logger = struct {
		TimeFormat string        `toml:"time_format"`
		Level      zerolog.Level `toml:"level"`
		Prettier   bool          `toml:"prettier"`
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
			Enable  bool
			Max     int
			ExpSecs int64 `toml:"expiration_seconds"`
		}
	}

	jwt struct {
		Secret string
	}
	password struct {
		HashCost int `toml:"hash_cost"`
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

func ParseConfig(name string) (*Config, error) {
	var contents *Config

	file, err := os.ReadFile("./" + name + ".toml")
	if err != nil {
		return &Config{}, err
	}

	err = toml.Unmarshal(file, &contents)
	return contents, err
}

func NewConfig() *Config {
	config, err := ParseConfig(".env")
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
