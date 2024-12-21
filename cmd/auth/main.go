package main

import (
	"flag"
	"log"

	"github.com/kozyarskaya/laba-11/internal/auth/api"
	"github.com/kozyarskaya/laba-11/internal/auth/config"
	"github.com/kozyarskaya/laba-11/internal/auth/provider"
	"github.com/kozyarskaya/laba-11/internal/auth/usecase"
	_ "github.com/lib/pq"
)

func main() {
	// Считываем аргументы командной строки
	configPath := flag.String("config-path", "D:\\Go\\FINL\\lr11\\\\configs\\auth.yaml", "путь к файлу конфигурации")
	flag.Parse()
	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	//Инициализация провайдера
	prv := provider.NewProvider(cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBname)
	//Инициализация бизнес-логики
	use := usecase.NewUsecase(prv)
	//Инициализация сервера
	srv := api.NewServer(cfg.IP, cfg.Port, use)
	srv.Run()
}
