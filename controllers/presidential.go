package controllers

import (
	"azimio/database"
	"azimio/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPollingData(c *gin.Context) {
	var pollingdata []models.Pollingdata
	if err := database.DBCon.Find(&pollingdata).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	c.IndentedJSON(http.StatusOK, pollingdata)
}

func GetPollingDataById(c *gin.Context) {
	pollingstationid := c.Params.ByName("pollingstationid")
	var pollingdata []models.Pollingdata

	m := make(map[string]interface{})
	m["pollingstationid"] = pollingstationid

	if err := database.DBCon.Where(m).Find(&pollingdata).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.IndentedJSON(http.StatusOK, pollingdata)
}

func UpdatePresidentialVotes(c *gin.Context) {
	var pollingdata models.Pollingdata
	q := make(map[string]interface{})
	q["id"] = c.Params.ByName("id")

	if err := database.DBCon.Where(q).First(&pollingdata).Error; err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.BindJSON(&pollingdata)
	database.DBCon.Save(&pollingdata)
	c.IndentedJSON(http.StatusOK, pollingdata)
}

func UpdatePresidentialVotesByCode(c *gin.Context) {
	var pollingdata models.Pollingdata
	q := make(map[string]interface{})
	q["altcode"] = c.Params.ByName("altcode")
	q["pollingstationid"] = c.Params.ByName("pollingstationid")

	if err := database.DBCon.Where(q).First(&pollingdata).Error; err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.BindJSON(&pollingdata)
	database.DBCon.Save(&pollingdata)
	c.IndentedJSON(http.StatusOK, pollingdata)
}

func GetUsersById(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	if err := database.DBCon.Where("id", id).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.IndentedJSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	database.DBCon.Create(&user)
	c.IndentedJSON(http.StatusOK, user)
}

func UpdateForm(c *gin.Context) {
	var forms models.Forms
	q := make(map[string]interface{})
	q["pollingstationid"] = c.Params.ByName("id")
	x := c.Params.ByName("id")
	fmt.Println(x)

	if err := database.DBCon.Where(q).Take(&forms).Error; err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		fmt.Println(err)
	}
	c.BindJSON(&forms)
	database.DBCon.Save(&forms)
	c.IndentedJSON(http.StatusOK, forms)
}

func GetResultsByCounty(c *gin.Context) {
	type Result struct {
		Candidate string `json:"candidate"`
		Cname     string `json:"cname"`
		Votes     int    `json:"votes"`
		// Percentage int    `json:"percentage"`
	}

	var result []Result
	// round(votes*100/ sum(votes),2) as percentage
	if err := database.DBCon.Table("pollingdata").Select("cname, candidate, sum(votes) as votes").Group("cname, candidate").Order("votes desc").Find(&result).Error; err != nil {
		c.AbortWithStatus(404)
	}
	c.IndentedJSON(http.StatusOK, result)
}

func GetResultsByCountry(c *gin.Context) {
	type Result struct {
		Candidate string `json:"candidate"`
		Votes     int    `json:"votes"`
		// Percentage int    `json:"percentage"`
	}

	var result []Result
	// round(votes*100/ sum(votes),2) as percentage
	if err := database.DBCon.Table("pollingdata").Select("candidate, sum(votes) as votes").Group("candidate").Order("votes desc").Find(&result).Error; err != nil {
		c.AbortWithStatus(404)
	}
	c.IndentedJSON(http.StatusOK, result)
}

func GetTotalVotes(c *gin.Context) {
	type Result struct {
		Votes int `json:"votes"`
		// Percentage int    `json:"percentage"`
	}

	var result []Result
	// round(votes*100/ sum(votes),2) as percentage
	if err := database.DBCon.Table("pollingdata").Select("sum(votes) as votes").Take(&result).Error; err != nil {
		c.AbortWithStatus(404)
	}
	c.IndentedJSON(http.StatusOK, result)
}

func GetFormByPollingId(c *gin.Context) {
	var forms models.Forms

	m := make(map[string]interface{})
	m["pollingstationid"] = c.Params.ByName("pollingstationid")

	if err := database.DBCon.Where(m).Find(&forms).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.IndentedJSON(http.StatusOK, forms)
}
