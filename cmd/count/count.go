package main

import (
	"flag"
	"log"

	"github.com/kozyarskaya/laba-10/internal/count/api"
	"github.com/kozyarskaya/laba-10/internal/count/config"
	"github.com/kozyarskaya/laba-10/internal/count/provider"
	"github.com/kozyarskaya/laba-10/internal/count/usecase"
	_ "github.com/lib/pq"
)

func main() {
	// Считываем аргументы командной строки
	configPath := flag.String("config-path", "D:\\Progects go\\laba-10\\\\configs\\count.yaml", "путь к файлу конфигурации")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	//Инициализация провайдера
	prv := provider.NewProvider(cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBname)
	//Инициализация бизнес-логики
	use := usecase.NewUsecase(cfg.Usecase.DefaultCount, prv)
	//Инициализация сервера
	srv := api.NewServer(cfg.IP, cfg.Port, cfg.API.MaxNum, use)

	srv.Run()
}
