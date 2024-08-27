package main

import (
    "github.com/Afomiat/Loan-Tracker-API/config"
    "github.com/Afomiat/Loan-Tracker-API/delivery/routers"
    "log"
)

func main() {
    env := config.NewEnv()
    router := routers.NewRouter(env)

    if err := router.Run(env.ServerAddress); err != nil {
        log.Fatal("Server failed to start: ", err)
    }
}
