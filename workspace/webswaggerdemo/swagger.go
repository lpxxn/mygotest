package main

import (
	"flag"
	"net/http"
	"strings"
	"text/template"

	"github.com/gin-gonic/gin"
)

var SWAGDIR = "./swagger-ui"
var staticContent = flag.String("staticPath", SWAGDIR, "Path to folder with Swagger UI")
var apiurl = flag.String("api", "http://localhost:5065", "The base path URI of the API service")

func swaggify(router *gin.Engine) {

	// Swagger Routes
	router.GET("/", IndexHandler)
	router.Static("/swagger-ui", *staticContent)
	for apiKey := range apiDescriptionsJson {
		router.GET("/"+apiKey+"/", ApiDescriptionHandler)
	}

	// API json data
	router.ApiDescriptionsJson = apiDescriptionsJson
}

func IndexHandler(c *gin.Context) {
	w := c.Writer
	r := c.Request

	isJsonRequest := false

	if acceptHeaders, ok := r.Header["Accept"]; ok {
		for _, acceptHeader := range acceptHeaders {
			if strings.Contains(acceptHeader, "json") {
				isJsonRequest = true
				break
			}
		}
	}

	if isJsonRequest {
		t, e := template.New("desc").Parse(resourceListingJson)
		if e != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		t.Execute(w, *apiurl)
	} else {
		http.Redirect(w, r, "/swagger-ui/", http.StatusFound)
	}
}

func ApiDescriptionHandler(c *gin.Context) {
	w := c.Writer
	r := c.Request

	apiKey := strings.Trim(r.RequestURI, "/")

	if json, ok := apiDescriptionsJson[apiKey]; ok {
		t, e := template.New("desc").Parse(json)
		if e != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		t.Execute(w, *apiurl)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
