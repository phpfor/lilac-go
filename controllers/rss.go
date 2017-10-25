package controllers

import (
	"fmt"
	"github.com/phpfor/lilac-go/models"
	"github.com/phpfor/lilac-go/system"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/feeds"
	"time"
)

//RssGet handles GET /rss route
func RssGet(c *gin.Context) {
	now := time.Now()
	domain := system.GetConfig().Domain

	feed := &feeds.Feed{
		Title:       "ginblog",
		Link:        &feeds.Link{Href: domain},
		Description: "GIN-powered blog boilerplate",
		Author:      &feeds.Author{Name: "Blog Author", Email: "author@blog.net"}, //hide email from spammers?
		Created:     now,
		Copyright:   "This work is copyright © Ginblog",
	}

	feed.Items = make([]*feeds.Item, 0)
	posts, err := models.GetPublishedPosts()
	if err != nil {
		c.HTML(404, "errors/404", nil)
		return
	}

	for i := range posts {
		feed.Items = append(feed.Items, &feeds.Item{
			Id:          fmt.Sprintf("%s/posts/%d", domain, posts[i].ID),
			Title:       posts[i].Title,
			Link:        &feeds.Link{Href: fmt.Sprintf("%s/posts/%d", domain, posts[i].ID)},
			Description: string(posts[i].Excerpt()),
			Author:      &feeds.Author{Name: posts[i].Author.Name},
			Created:     now,
		})
	}

	rss, err := feed.ToRss()
	if err != nil {
		c.HTML(500, "errors/500", nil)
		return
	}
	c.Data(200, "text/xml", []byte(rss))
}
