package main

import (
	"fmt"
	"github.com/babulal107/go-kubernetes-poc/internal/config"
	"github.com/babulal107/go-kubernetes-poc/internal/router"
	"github.com/google/martian/log"
)

func main() {

	// initialize application log service name and message field name
	var configs, err = config.Init()
	if err != nil {
		log.Errorf(config.ErrMsgUnableToInitConfig, err)
	}

	// initialize router
	ginRouter := router.Init(configs)
	if err := ginRouter.Run(fmt.Sprintf(":%d", configs.APP.GetPort())); err != nil {
		log.Errorf("Unable to start application", err)
	}
}
