package config

import (
	"log"
	"os"
	"time"
	"url-shortener/internal/storage"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	// yaml:env - name of variable in yaml file
	// env:ENV - название параметра из переменной окружения
	// env-default:"local" - значение по умолчанию, в этом случае, если нет переменной окружения, то будет использоваться local. Можно настроить prod, dev и т.д.
	// env-required:"true" - параметр, который обязательно должен быть установлен, либо приложение не запустится (чтобы точно не забыть установить значение)
	Env        string                `yaml:"env"  env-default:"local"` // env-required:"true", либо env-default:"local"
	Storage    storage.StorageConfig `yaml:"storage"`
	HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

// Приставка Must означает, что функция должна паниковать, если не сможет прочитать файл конфигурации.
// Используется для инициализации конфигурации в main.go
func MustLoad() *Config {

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatalf("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	return &cfg
}
