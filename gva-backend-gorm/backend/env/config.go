package env

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

const (
	Dev  = "dev"
	Stag = "stag"
	Prod = "prod"
)

const (
	envFilePath = "."
	envFileName = ".env"

	seperator = "__"
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
		Name            string `default:"dev"`
		Port            string `default:":4000"`
		PrintRoutes     bool   `mapstructure:"print_routes" default:"true"`
		Env             string `validate:"oneof=dev stag prod" default:"dev"`
		IdleTimeout     int64  `mapstructure:"idle_timeout" default:"5"`
		ShutdownTimeout int64  `mapstructure:"shutdown_timeout" default:"10"`
		TLS             struct {
			Auto     bool
			Enable   bool
			CertFile string `mapstructure:"cert_file"`
			KeyFile  string `mapstructure:"key_file"`
		}
		Host string `mapstructure:"host" default:"localhost:4000"`
	}

	api struct {
		Admin struct {
			Enable   bool   `default:"true"`
			Port     string `default:"4001"`
			BasePath string `mapstructure:"base_path" default:"/admin/v1"`
			Auth     struct {
				JwtSecret        string `mapstructure:"jwt_secret" default:"secret"`
				PasswordHashCost int    `mapstructure:"password_hash_cost" default:"10"`
				TotpTestCode     string `mapstructure:"totp_test_code" default:"666666"`
				AccessTokenTTL   int    `mapstructure:"access_token_ttl" default:"7200"`
				RefreshTokenTTL  int    `mapstructure:"refresh_token_ttl" default:"86400"`
			}
		}
		Bot struct {
			Enable   bool   `default:"true"`
			Port     string `default:"4002"`
			BasePath string `mapstructure:"base_path" default:"/bot/v1"`
		}
		Web struct {
			Enable   bool   `default:"true"`
			Port     string `default:"4000"`
			BasePath string `mapstructure:"base_path" default:"/web/v1"`
		}
	}

	db = struct {
		Url    string `default:"mysql://root:password@db-service/gva?parseTime=true"`
		UrlDev string `default:"mysql://root:password@db-service/gva_dev?parseTime=true" mapstructure:"url_dev"`
		Redis  struct {
			Enable bool  `default:"true"`
			Url    string `default:"rediss://:password@cache-service"`
		}
		Seed struct {
			// enable true, always run on app start
			Enable bool `default:"true"`
			// skip auto run on specific types
			BlacklistTypes []string `mapstructure:"blacklist_types" default:"[]"`
			// seed default super admin
			SuperAdmin struct {
				Username string `default:"admin"`
				Password string `default:"123456"`
			} `mapstructure:"super_admin"`
		}
	}

	middleware = struct {
		Swagger struct {
			Enable bool   `default:"true"`
			Path   string `default:"/docs"`
		}
		Compress struct {
			Enable bool `default:"true"`
			Level  int  `default:"1"`
		}
		Monitor struct {
			Enable bool `default:"true"`
			Path   string
		}
		Pprof struct {
			Enable bool `default:"true"`
		}
		Limiter struct {
			Enable        bool  `default:"true"`
			Max           int   `default:"20"`
			ExpirationTTL int64 `mapstructure:"expiration_ttl" default:"60"`
		}
		Translation struct {
			Enable bool `default:"true"`
		}
	}

	s3 struct {
		Region   string
		Key      string `mapstructure:"access_key_id"`
		Secret   string `mapstructure:"secret_access_key"`
		Bucket   string `mapstructure:"bucket_name"`
		Endpoint string
		Session  string `mapstructure:"session_token"`
	}

	logger = struct {
		TimeFormat string `mapstructure:"time_format"`
		Level      int8   `default:"0"`
		Prettier   bool   `default:"true"`
	}

	google struct {
		Enable         bool
		ChatWebhookURL string `mapstructure:"chat_webhook_url"`
	}
)

type CtxKey struct{}

type Config struct {
	App        app
	API        api
	DB         db
	Middleware middleware
	S3         s3
	Logger     logger
	Google     google
}

func NewConfig() *Config {
	// generate .env if not exist
	config, err := ReadEnvFromFile()
	if err != nil {
		panic(fmt.Errorf("false to read env %v", err))
	}

	if err := validator.New().Struct(config); err != nil {
		panic(fmt.Errorf("false to validate env validator: %v", err))
	}

	return config
}
