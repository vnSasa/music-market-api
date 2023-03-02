package main

import(
	"github.com/vnSasa/music-market-api/pkg/repository"
	"github.com/vnSasa/music-market-api/pkg/service"
	"github.com/vnSasa/music-market-api/pkg/handler"
	api "github.com/vnSasa/music-market-api"
	"github.com/spf13/viper"
	"github.com/sirupsen/logrus"

	"os"
	"os/signal"
	"syscall"
)

func main()  {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	
	db, err := repository.NewMySQLDB(repository.Config{
		UserName:	viper.GetString("db.dbusername"),
		Password:	viper.GetString("db.dbpassword"),
		Host:	viper.GetString("db.dbhost"),
		Port:	viper.GetString("db.dbport"),
		DBName:	viper.GetString("db.dbname"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(api.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoute()); err != nil {
			logrus.Fatalf("error occurred running http server: %s", err.Error())
		}
	}()

	logrus.Print("Market Started")
	
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Market Closed")
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}