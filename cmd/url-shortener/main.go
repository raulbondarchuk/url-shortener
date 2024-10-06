package main

import (
	"fmt"
	"os"
	"url-shortener/internal/config"
	"url-shortener/internal/lib/logger"
	"url-shortener/internal/lib/logger/sl"
	"url-shortener/internal/storage/mysql"
)

func main() {

	// ROADMAP of project

	config.SetupConfigPath() // TODO: init CONFIG_PATH: godotenv --> .env

	cfg := config.MustLoad() // TODO: init config: cleanenv --> config.MustLoad(env)
	fmt.Println(cfg)         //  remove on dev/prod

	log := logger.SetupLogger(cfg.Env) // TODO: init Logger: slog (biblioteca para trabajar con logs --> text, json, etc)
	fmt.Println(log)

	// Инициализация хранилища
	storage, err := mysql.New(cfg) // TODO: init storage: mysql
	if err != nil {
		log.Error("failed to initialize storage", sl.Err(err))
		os.Exit(1) // Exit code 1 - error
	}

	_ = storage

	// TODO:init router: chi, chi-render

	// TODO: run server
}
