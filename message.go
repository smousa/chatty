package main

import (
	"regexp"

	"github.com/goadesign/goa"
	"github.com/smousa/chatty/app"
)

var (
	mentionsRe  = regexp.MustCompile(`@(\w+)`) // TODO: emails?
	emoticonsRe = regexp.MustCompile(`\(([0-9A-z]{1,15})\)`)
	linksRe     = regexp.MustCompile(`(https?://\S+)`) // TODO: less strict for urls
)

// MessageController implements the message resource.
type MessageController struct {
	*goa.Controller
	URLReader URLReader
}

// NewMessageController creates a message controller.
func NewMessageController(service *goa.Service) *MessageController {
	return &MessageController{Controller: service.NewController("MessageController")}
}

// Parse runs the parse action.
func (c *MessageController) Parse(ctx *app.ParseMessageContext) error {
	res := &app.ChattyMessageInfo{
		Mentions:  c.getMentions(ctx.Payload.Input),
		Emoticons: c.getEmoticons(ctx.Payload.Input),
		Links:     c.getLinks(ctx.Payload.Input),
	}
	return ctx.OK(res)
}

func (c *MessageController) getMentions(input string) (mentions []string) {
	out := mentionsRe.FindAllStringSubmatch(input, -1)
	mentions = make([]string, len(out))
	for i, o := range out {
		mentions[i] = o[1]
	}
	return
}

func (c *MessageController) getEmoticons(input string) (emoticons []string) {
	out := emoticonsRe.FindAllStringSubmatch(input, -1)
	emoticons = make([]string, len(out))
	for i, o := range out {
		emoticons[i] = o[1]
	}
	return
}

func (c *MessageController) getLinks(input string) (links app.ChattyLinkCollection) {
	out := linksRe.FindAllStringSubmatch(input, -1)
	for _, o := range out {
		title, err := c.URLReader.PageTitle(o[1])
		if err == nil {
			links = append(links, &app.ChattyLink{
				URL:   o[1],
				Title: title,
			})
		}
	}
	return
}
