//go:generate goagen bootstrap -d github.com/smousa/chatty/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/smousa/chatty/app"
)

func main() {
	// Create service
	service := goa.New("chatty")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "message" controller
	c := NewMessageController(service)
	app.MountMessageController(service, c)
	c.URLReader = &DefaultURLReader{}

	// Start service
	if err := service.ListenAndServe(":8081"); err != nil {
		service.LogError("startup", "err", err)
	}
}
