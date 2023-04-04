package main

import (
	"fmt"

	"github.com/amajakai14/admin-renewal/internal/channel"
	"github.com/amajakai14/admin-renewal/internal/db"
	httpTransport "github.com/amajakai14/admin-renewal/internal/transport/http"
	appUser "github.com/amajakai14/admin-renewal/internal/user"
)

func Run() error {
	fmt.Println("starting an application")
	db, err := db.NewDatabase()
	if err != nil {
		return err
	}

	if err := db.MigrateDB(); err != nil {
		return err
	}

	services := initializeServices(db)
	httpHandler := httpTransport.NewHandler(services)
	if err := httpHandler.Serve(); err != nil {
		return err
	}
	
	return nil
}

func initializeServices(db *db.Database) httpTransport.Services {
	userService := appUser.NewService(db)
	channelService := channel.NewService(db)
	services := httpTransport.Services{
		UserService: userService,
		ChannelService: channelService,
	}
	return services
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
