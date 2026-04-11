package utils

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Configuration struct {
	AppName       string
	Port          string
	Debug         bool
	Limit         int
	PathLogging   string
	JWTSecret     string
	Issuer        string
	DB            DatabaseConfig
	Redis         RedisConfig
	SMTP          SMTPConfig
	BaseURL       string
	BusinessRules BusinessRules
	RajaOngkirConfig RajaOngkirConfig
	CloudinaryConfig CloudinaryConfig
	ClientHost string
}

type DatabaseConfig struct {
	Name     string
	Username string
	Password string
	Host     string
	Port     string
	MaxConn  int32
	ConnString string
}

type SMTPConfig struct {
	Host     string
	Port     int
	Email    string
	Password string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
	Conn string
}

type BusinessRules struct {
	TaxRate           int
	ProfitMargin      int
	DefaultShiftStart string
	DefaultShiftEnd   string
}

type RajaOngkirConfig struct {
	APIKey string
	APILimit int
}

type CloudinaryConfig struct {
	CloudName string
	APIKey    string
	APISecret string
}

func ReadConfiguration() (Configuration, error) {
	// 1. flags (highest priority)
	pflag.Int("port", 0, "port for app golang")
	pflag.Parse()
	_ = viper.BindPFlags(pflag.CommandLine)

	// 2. environment variables
	viper.AutomaticEnv()

	// 3. optional .env file (lowest priority)
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	_ = viper.ReadInConfig() // DO NOT fail if missing

	return Configuration{
		AppName:     viper.GetString("APP_NAME"),
		Port:        viper.GetString("PORT"),
		Debug:       viper.GetBool("DEBUG"),
		Limit:       viper.GetInt("LIMIT"),
		PathLogging: viper.GetString("PATH_LOGGING"),
		JWTSecret:   viper.GetString("JWT_SECRET"),
		Issuer:      viper.GetString("JWT_ISSUER"),
		ClientHost: viper.GetString("CLIENT_HOST"),

		DB: DatabaseConfig{
			Name:     viper.GetString("DB_NAME"),
			Username: viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASSWORD"),
			Host:     viper.GetString("DB_HOST"),
			Port:     viper.GetString("DB_PORT"),
			MaxConn:  viper.GetInt32("DB_MAX_CONN"),
			ConnString: viper.GetString("DB_CONN"),
		},

		SMTP: SMTPConfig{
			Host:     viper.GetString("SMTP_HOST"),
			Port:     viper.GetInt("SMTP_PORT"),
			Email:    viper.GetString("SMTP_EMAIL"),
			Password: viper.GetString("SMTP_PASSWORD"),
		},

		Redis: RedisConfig{
			Host:     viper.GetString("REDIS_HOST"),
			Port:     viper.GetString("REDIS_PORT"),
			Password: viper.GetString("REDIS_PASSWORD"),
			DB:       viper.GetInt("REDIS_DB"),
			Conn: viper.GetString("REDIS_CONN"),
		},

		BaseURL: viper.GetString("APP_URL"),

		BusinessRules: BusinessRules{
			TaxRate:           viper.GetInt("TAX_RATE"),
			ProfitMargin:      viper.GetInt("PROFIT_MARGIN"),
			DefaultShiftStart: viper.GetString("DEFAULT_SHIFT_START"),
			DefaultShiftEnd:   viper.GetString("DEFAULT_SHIFT_END"),
		},

		RajaOngkirConfig: RajaOngkirConfig{
			APIKey:   viper.GetString("RAJAONGKIR_API_KEY"),
			APILimit: viper.GetInt("RAJAONGKIR_API_LIMIT"),
		},

		CloudinaryConfig: CloudinaryConfig{
			CloudName: viper.GetString("CLOUDINARY_CLOUD_NAME"),
			APIKey:    viper.GetString("CLOUDINARY_API_KEY"),
			APISecret: viper.GetString("CLOUDINARY_API_SECRET"),
		},
	}, nil
}

