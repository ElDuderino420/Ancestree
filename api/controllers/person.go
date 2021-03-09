package controllers

import (
	"api/httputil"
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PersonController struct{}

var personModel = new(models.Person)

// Get godoc
// @Summary Get Person
// @Description Gets Person by ID
// @Tags person
// @Accept  json
// @Produce  json
// @Param id path string true "Person ID"
// @Success 200 {object} string
// @Router /person/{id} [get]
func (p PersonController) Get(c *gin.Context) {
	if c.Param("id") != "" {
		user, err := personModel.GetByID(c.Param("id"))
		if err != nil {
			httputil.NewError(c, http.StatusInternalServerError, err, "Error while getting person")
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Person found!", "person": user})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
	return
}

// Create godoc
// @Summary Create Person
// @Description Creates a Person
// @Tags person
// @Accept  json
// @Produce  json
// @Param person body models.Person true "Person"
// @Success 200 {object} string
// @Router /person/ [post]
func (p PersonController) Create(c *gin.Context) {
	var newPerson models.Person
	err := c.ShouldBindJSON(&newPerson)
	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err, "Invalid body")
	}
	gg, err := newPerson.Create()
	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err, "Shit.")
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request", "body": gg})
	c.Abort()
	return
}

// CreateLink godoc
// @Summary Create Person link
// @Description Creates a Person
// @Tags person
// @Accept  json
// @Produce  json
// @Param id path string true "Person ID"
// @Param person body models.Person true "Person"
// @Success 200 {object} string
// @Router /person/{id} [post]
func (p PersonController) CreateLink(c *gin.Context) {
	var newPerson models.Person
	err := c.ShouldBindJSON(&newPerson)
	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err, "Invalid body")
	}
	gg, err := newPerson.LinkMother(c.Param("id"))
	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err, "Shit.")
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request", "body": gg})
	c.Abort()
	return
}
