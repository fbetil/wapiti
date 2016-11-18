package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/renders/multitemplate"

	"github.com/gin-gonic/gin"
)

func main() {
	//define main router
	router := gin.Default()

	//define templates
	templates := multitemplate.New()
	templates.AddFromFiles("index", "pages/layout.html", "pages/index.html")
	templates.AddFromFiles("logs", "pages/layout.html", "pages/logs.html")
	templates.AddFromFiles("packages", "pages/layout.html", "pages/packages.html")
	templates.AddFromFiles("parameters", "pages/layout.html", "pages/parameters.html")
	templates.AddFromFiles("rule", "pages/layout.html", "pages/rule.html")
	templates.AddFromFiles("rules", "pages/layout.html", "pages/rules.html")
	templates.AddFromFiles("script", "pages/layout.html", "pages/script.html")
	templates.AddFromFiles("scripts", "pages/layout.html", "pages/scripts.html")
	router.HTMLRender = templates

	//set static
	router.Static("/assets", "./assets")
	//router.StaticFS("/more_static", http.Dir("my_file_system"))

	//Routes
	router.GET("/index", renderIndex)
	router.GET("/logs", renderLogs)
	router.GET("/packages", renderPackages)
	router.GET("/parameters", renderParameters)
	router.GET("/rule/:id", renderRule)
	router.GET("/rules", renderRules)
	router.GET("/script/:id", renderScript)
	router.GET("/scripts", renderScripts)

	//Default route
	router.NoRoute(func(c *gin.Context) { c.Redirect(http.StatusMovedPermanently, "/index") })

	router.Run(":3000")
}

func renderIndex(c *gin.Context) {
	c.HTML(200, "index", gin.H{
		"title": "Accueil",
	})
}

func renderLogs(c *gin.Context) {
	c.HTML(200, "logs", gin.H{
		"title": "Journaux",
	})
}

func renderPackages(c *gin.Context) {
	c.HTML(200, "packages", gin.H{
		"title": "Packages",
	})
}

func renderParameters(c *gin.Context) {
	c.HTML(200, "parameters", gin.H{
		"title": "Paramètres",
	})
}

func renderRule(c *gin.Context) {
	c.HTML(200, "rule", gin.H{
		"title": "Règle",
	})
}

func renderRules(c *gin.Context) {
	c.HTML(200, "rules", gin.H{
		"title": "Règles",
	})
}

func renderScript(c *gin.Context) {
	c.HTML(200, "script", gin.H{
		"title": "Script",
	})
}

func renderScripts(c *gin.Context) {
	c.HTML(200, "scripts", gin.H{
		"title": "Scripts",
	})
}
