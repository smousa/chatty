// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/smousa/chatty/design
// --out=$(GOPATH)/src/github.com/smousa/chatty
// --version=v1.1.0-dirty
//
// API "chatty": Application Media Types
//
// The content of this file is auto-generated, DO NOT MODIFY

package client

import (
	"github.com/goadesign/goa"
	"net/http"
)

// Info about a link (default view)
//
// Identifier: application/chatty.link+json; view=default
type ChattyLink struct {
	// Title of the url
	Title string `form:"title" json:"title" xml:"title"`
	// Parsed url
	URL string `form:"url" json:"url" xml:"url"`
}

// Validate validates the ChattyLink media type instance.
func (mt *ChattyLink) Validate() (err error) {
	if mt.URL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "url"))
	}
	if mt.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title"))
	}
	return
}

// DecodeChattyLink decodes the ChattyLink instance encoded in resp body.
func (c *Client) DecodeChattyLink(resp *http.Response) (*ChattyLink, error) {
	var decoded ChattyLink
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// ChattyLinkCollection is the media type for an array of ChattyLink (default view)
//
// Identifier: application/chatty.link+json; type=collection; view=default
type ChattyLinkCollection []*ChattyLink

// Validate validates the ChattyLinkCollection media type instance.
func (mt ChattyLinkCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeChattyLinkCollection decodes the ChattyLinkCollection instance encoded in resp body.
func (c *Client) DecodeChattyLinkCollection(resp *http.Response) (ChattyLinkCollection, error) {
	var decoded ChattyLinkCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// Info about a message (default view)
//
// Identifier: application/chatty.message_info+json; view=default
type ChattyMessageInfo struct {
	// Emoticons
	Emoticons []string `form:"emoticons,omitempty" json:"emoticons,omitempty" xml:"emoticons,omitempty"`
	// Links
	Links ChattyLinkCollection `form:"links,omitempty" json:"links,omitempty" xml:"links,omitempty"`
	// Mentions
	Mentions []string `form:"mentions,omitempty" json:"mentions,omitempty" xml:"mentions,omitempty"`
}

// Validate validates the ChattyMessageInfo media type instance.
func (mt *ChattyMessageInfo) Validate() (err error) {
	if err2 := mt.Links.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// DecodeChattyMessageInfo decodes the ChattyMessageInfo instance encoded in resp body.
func (c *Client) DecodeChattyMessageInfo(resp *http.Response) (*ChattyMessageInfo, error) {
	var decoded ChattyMessageInfo
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
