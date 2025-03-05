package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"urmi-arch/common/middlewares"
	"urmi-arch/common/settings"
	"urmi-arch/user/handler"
	"urmi-arch/user/repositories"
	"urmi-arch/user/service"

	"github.com/go-redis/redis"
)

func main() {
	isApiKey := flag.Bool("is_api_key", false, "Set for use API-Key instead JWT")
	flag.Parse()

	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	pong, err := rdb.Ping().Result()
	if err != nil {
		log.Fatalf("Could not ping Redis: %v", err)
	}
	fmt.Println("Redis Ping Response:", pong)
	middlewares.APIKEYS_INIT()
	rl := middlewares.NewRLimit(rdb)
	fmt.Println(rl)
	userRepo := repositories.NewUserRepo()
	userSvc := service.NewUserService(userRepo)
	fmt.Println("APIKEY ", *isApiKey)
	r := handler.UserRoutes(userSvc, rl, *isApiKey)
	server := http.Server{
		Addr:    settings.UserServerPort,
		Handler: r,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
