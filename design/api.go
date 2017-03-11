package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("chatty", func() {
	Title("Chat message parser")
	Description("Parse all of your chats and other things!")
	Host("localhost:8081")
	Scheme("http")
	BasePath("/api")

	Origin("http://swagger.goa.design", func() {
		Methods("GET", "POST", "PUT", "DELETE", "HEAD")
		MaxAge(600)
		Credentials()
	})

	ResponseTemplate(Created, func(pattern string) {
		Description("Resource created")
		Status(201)
		Headers(func() {
			Header("Location", String, "href to created resource", func() {
				Pattern(pattern)
			})
		})
	})
})
