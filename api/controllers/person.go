package controllers

import (
	"api/httputil"
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PersonController ...
type PersonController struct{}

var personModel = new(models.Person)
var familtyTreeModel = new(models.FamilyTree)

// Get godoc
// @Summary Get Person
// @Description Gets Person by ID
// @Tags person
// @Accept  json
// @Produce  json
// @Param id path string true "Person ID"
// @Success 200 {object} models.Person
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
// @Success 200 {object} models.Person
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

// Update godoc
// @Summary Update Person
// @Description Updates a Person
// @Tags person
// @Accept  json
// @Produce  json
// @Param id path string true "Person ID"
// @Param person body models.Person true "Person"
// @Success 200 {object} models.Person
// @Router /person/{id} [put]
func (p PersonController) Update(c *gin.Context) {
	var person models.Person
	err := c.ShouldBindJSON(&person)
	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err, "Invalid body")
	}
	person.ID = c.Param("id")
	gg, err := person.Update()
	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err, "Shit.")
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request", "body": gg})
	c.Abort()
	return
}

// CreateParent godoc
// @Summary Create Parent link
// @Description Creates a Person
// @Tags person
// @Accept  json
// @Produce  json
// @Param id path string true "Person ID"
// @Param person body models.Person true "Person"
// @Success 200 {object} models.Person
// @Router /person/parent/{id} [post]
func (p PersonController) CreateParent(c *gin.Context) {
	var newPerson models.Person
	err := c.ShouldBindJSON(&newPerson)
	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err, "Invalid body")
	}
	gg, err := newPerson.CreateParent(c.Param("id"))
	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err, "Shit.")
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request", "body": gg})
	c.Abort()
	return
}

// CreateChild godoc
// @Summary Create Child link
// @Description Creates a Person
// @Tags person
// @Accept  json
// @Produce  json
// @Param id path string true "Person ID"
// @Param person body models.Person true "Person"
// @Success 200 {object} models.Person
// @Router /person/child/{id} [post]
func (p PersonController) CreateChild(c *gin.Context) {
	var newPerson models.Person
	err := c.ShouldBindJSON(&newPerson)
	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err, "Invalid body")
	}
	newChild, err := newPerson.CreateChild(c.Param("id"))
	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err, "Shit.")
	}
	c.JSON(http.StatusOK, gin.H{"newChild": newChild})
	c.Abort()
	return
}

// CreateLink godoc
// @Summary Create Child link
// @Description Creates a Person
// @Tags person
// @Accept  json
// @Produce  json
// @Param childId path string true "Child ID"
// @Param parentId path string true "Parent ID"
// @Success 200 {object} string
// @Router /person/link/{childId}/{parentId} [post]
func (p PersonController) CreateLink(c *gin.Context) {
	var newPerson models.Person
	gg, err := newPerson.CreateLink(c.Param("childId"), c.Param("parentId"))
	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err, "Shit.")
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request", "body": gg})
	c.Abort()
	return
}
