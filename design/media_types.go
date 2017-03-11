package design

import (
	"fmt"

	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

func mimetype(name string) string {
	return fmt.Sprintf("application/chatty.%s+json", name)
}

var (
	MessageInfo *MediaTypeDefinition
	LinkInfo    *MediaTypeDefinition
)

func init() {
	MessageInfo = MediaType(mimetype("message_info"), func() {
		Description("Info about a message")
		Attributes(func() {
			Attribute("mentions", ArrayOf(String), "Mentions")
			Attribute("emoticons", ArrayOf(String), "Emoticons")
			Attribute("links", CollectionOf(LinkInfo), "Links")
		})
		View("default", func() {
			Attribute("mentions")
			Attribute("emoticons")
			Attribute("links")
		})
	})

	LinkInfo = MediaType(mimetype("link"), func() {
		Description("Info about a link")
		Attributes(func() {
			Attribute("url", String, "Parsed url")
			Attribute("title", String, "Title of the url")
			Required("url", "title")
		})
		View("default", func() {
			Attribute("url")
			Attribute("title")
		})
	})
}
