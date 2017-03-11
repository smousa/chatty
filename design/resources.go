package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("message", func() {
	DefaultMedia(MessageInfo)
	BasePath("/messages")

	Action("parse", func() {
		Routing(PUT(""))
		Description("Extract message info from message")
		Payload(func() {
			Member("input")
			Required("input")
		})
		Response(OK, MessageInfo)
	})
})
