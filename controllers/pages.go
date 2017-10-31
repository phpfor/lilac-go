package controllers

import (
	"net/http"

	"html/template"

	"github.com/Sirupsen/logrus"
	"github.com/phpfor/lilac-go/helpers"
	"github.com/phpfor/lilac-go/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	//"github.com/russross/blackfriday"
)

//PageGet handles GET /pages/:id route
func PageGet(c *gin.Context) {
	page, err := models.GetPageBySlug(c.Param("slug"))
	if err != nil || !page.Published {
		c.HTML(http.StatusNotFound, "errors/404", nil)
		return
	}
	h := helpers.DefaultH(c)
	h["Title"] = page.Name
	h["Description"] = template.HTML(page.Description)
	h["Active"] = "pages"
	c.HTML(http.StatusOK, "pages/show", h)
}

//PageIndex handles GET /admin/pages route
func PageIndex(c *gin.Context) {
	list, err := models.GetPages()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "errors/500", nil)
		logrus.Error(err)
		return
	}
	h := helpers.DefaultH(c)
	h["Title"] = "List of pages"
	h["List"] = list
	h["Active"] = "pages"
	c.HTML(http.StatusOK, "admin/pages/index", h)
}

//PageNew handles GET /admin/new_page route
func PageNew(c *gin.Context) {
	h := helpers.DefaultH(c)
	h["Title"] = "New page"
	h["Active"] = "pages"
	session := sessions.Default(c)
	h["Flash"] = session.Flashes()
	session.Save()

	c.HTML(http.StatusOK, "admin/pages/form", h)
}

//PageCreate handles POST /admin/new_page route
func PageCreate(c *gin.Context) {
	page := &models.Page{}
	if err := c.Bind(page); err != nil {
		session := sessions.Default(c)
		session.AddFlash(err.Error())
		session.Save()
		c.Redirect(http.StatusSeeOther, "/admin/new_page")
		return
	}

	if err := page.Insert(); err != nil {
		c.HTML(http.StatusInternalServerError, "errors/500", nil)
		logrus.Error(err)
		return
	}
	c.Redirect(http.StatusFound, "/admin/pages")
}

//PageEdit handles GET /admin/pages/:id/edit route
func PageEdit(c *gin.Context) {
	page, _ := models.GetPage(c.Param("id"))
	if page.ID == "" {
		c.HTML(http.StatusNotFound, "errors/404", nil)
		return
	}
	h := helpers.DefaultH(c)
	h["Title"] = "Edit page"
	h["Active"] = "pages"
	h["Page"] = page
	session := sessions.Default(c)
	h["Flash"] = session.Flashes()
	session.Save()
	c.HTML(http.StatusOK, "admin/pages/form", h)
}

//PageUpdate handles POST /admin/pages/:id/edit route
func PageUpdate(c *gin.Context) {
	page := &models.Page{}
	if err := c.Bind(page); err != nil {
		session := sessions.Default(c)
		session.AddFlash(err.Error())
		session.Save()
		c.Redirect(http.StatusSeeOther, "/admin/pages")
		return
	}
	if err := page.Update(c.Param("id")); err != nil {
		c.HTML(http.StatusInternalServerError, "errors/500", nil)
		logrus.Error(err)
		return
	}
	c.Redirect(http.StatusFound, "/admin/pages")
}

//PageDelete handles POST /admin/pages/:id/delete route
func PageDelete(c *gin.Context) {
	page, _ := models.GetPage(c.Param("slug"))
	if err := page.Delete(c.Param("slug")); err != nil {
		c.HTML(http.StatusInternalServerError, "errors/500", nil)
		logrus.Error(err)
		return
	}
	c.Redirect(http.StatusFound, "/admin/pages")
}
