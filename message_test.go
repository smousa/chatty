package main_test

import (
	. "github.com/smousa/chatty"
	"github.com/smousa/chatty/app"
	"github.com/smousa/chatty/app/test"
	"github.com/smousa/chatty/mocks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Message", func() {
	var (
		t          = GinkgoT()
		controller *MessageController

		input  string
		result *app.ChattyMessageInfo
		m      *mocks.URLReader
	)

	BeforeEach(func() {
		t = GinkgoT()
		m = &mocks.URLReader{}
		controller = NewMessageController(svc)
		controller.URLReader = m
	})

	JustBeforeEach(func() {
		_, result = test.ParseMessageOK(t, ctx, svc, controller, &app.ParseMessagePayload{
			Input: input,
		})
	})

	Context("@chris you around?", func() {
		BeforeEach(func() {
			input = "@chris you around"
		})

		It("should show that chris was mentioned", func() {
			Expect(result.Mentions).To(ConsistOf("chris"))
			Expect(result.Emoticons).To(BeEmpty())
			Expect(result.Links).To(BeEmpty())
		})
	})

	Context("Olympics are starting soon; http://www.nbcolympics.com", func() {
		BeforeEach(func() {
			input = "Olympics are starting soon; http://www.nbcolympics.com"
			m.On("PageTitle", "http://www.nbcolympics.com").Return("2016 Rio Olympic Games | NBC Olympics", nil)
		})

		It("should show the title and url of the link", func() {
			Expect(result.Mentions).To(BeEmpty())
			Expect(result.Emoticons).To(BeEmpty())
			Expect(result.Links).To(ConsistOf(&app.ChattyLink{
				URL:   "http://www.nbcolympics.com",
				Title: "2016 Rio Olympic Games | NBC Olympics",
			}))
		})
	})

	Context("@bob @john (success) such a cool feature; https://twitter.com/jdorfman/status/430511497475670016", func() {
		BeforeEach(func() {
			input = "@bob @john (success) such a cool feature; https://twitter.com/jdorfman/status/430511497475670016"
			m.On("PageTitle", "https://twitter.com/jdorfman/status/430511497475670016").Return("Justin Dorfman on Twitter: &quot;nice @littlebigdetail from @HipChat (shows hex colors when pasted in chat). http://t.co/7cI6Gjy5pq&quot;", nil)
		})

		It("should show bob, john, success, and the url", func() {
			Expect(result.Mentions).To(ConsistOf("bob", "john"))
			Expect(result.Emoticons).To(ConsistOf("success"))
			Expect(result.Links).To(ConsistOf(&app.ChattyLink{
				URL:   "https://twitter.com/jdorfman/status/430511497475670016",
				Title: "Justin Dorfman on Twitter: &quot;nice @littlebigdetail from @HipChat (shows hex colors when pasted in chat). http://t.co/7cI6Gjy5pq&quot;",
			}))
		})
	})
})
