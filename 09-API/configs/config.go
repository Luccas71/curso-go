package configs

import (
	"github.com/go-chi/jwtauth/v5"
	"github.com/spf13/viper"
)

// var de inicialização
var cfg *conf

// variaveis de ambiente
// configs db
type conf struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
	JWTExperesIn  int    `mapstructure:"JWT_EXPERES_IN"`
	TokenAuth     *jwtauth.JWTAuth
}

// função para carregar confs de env
func LoadConfig(path string) (*conf, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	// gerar code jwt
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)
	return cfg, nil
}
