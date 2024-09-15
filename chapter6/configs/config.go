package configs

import "github.com/spf13/viper"

type ConfigurationEnvironment struct {
	DBDriver     string `mapstructure:"DB_POSTGRES_DRIVER"`
	DBHost       string `mapstructure:"DB_POSTGRES_HOST"`
	DBPort       string `mapstructure:"DB_POSTGRES_PORT"`
	DBUser       string `mapstructure:"DB_POSTGRES_USER"`
	DBPass       string `mapstructure:"DB_POSTGRES_PASS"`
	DBName       string `mapstructure:"DB_POSTGRES_NAME"`
	DBCacheName  string `mapstructure:"DB_CACHE_NAME"`
	JWTSecret    string `mapstructure:"JWT_SECRET"`
	JWTExpiresIn int32  `mapstructure:"JWT_EXPIRES_IN"`
}

func LoadConfigurationEnvironment(path string) (*ConfigurationEnvironment, error) {

	var cfg *ConfigurationEnvironment

	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	err := viper.ReadInConfig()

	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)

	if err != nil {
		return nil, err
	}

	return cfg, nil
}
