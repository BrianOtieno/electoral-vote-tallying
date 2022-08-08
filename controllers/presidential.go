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

func GetParliamentaryData(c *gin.Context) {
	var pollingdata []models.Parliamentary
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

func GetParliamentaryPollingDataById(c *gin.Context) {
	pollingstationid := c.Params.ByName("pollingstationid")
	var pollingdata []models.Parliamentary

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

func UpdateParliamentaryVotes(c *gin.Context) {
	var pollingdata models.Parliamentary
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

func UpdateParliamentaryVotesByCode(c *gin.Context) {
	var parliamentarydata models.Parliamentary
	q := make(map[string]interface{})
	q["altcode"] = c.Params.ByName("altcode")
	q["pollingstationid"] = c.Params.ByName("pollingstationid")

	if err := database.DBCon.Where(q).First(&parliamentarydata).Error; err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.BindJSON(&parliamentarydata)
	database.DBCon.Save(&parliamentarydata)
	c.IndentedJSON(http.StatusOK, parliamentarydata)
}

func UpdateGubernatorialVotesByCode(c *gin.Context) {
	var gubernatorial models.Gubernatorial
	q := make(map[string]interface{})
	q["altcode"] = c.Params.ByName("altcode")
	q["pollingstationid"] = c.Params.ByName("pollingstationid")

	if err := database.DBCon.Where(q).First(&gubernatorial).Error; err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.BindJSON(&gubernatorial)
	database.DBCon.Save(&gubernatorial)
	c.IndentedJSON(http.StatusOK, gubernatorial)
}

func UpdateSenatorialVotesByCode(c *gin.Context) {
	var senatorial models.Senatorial
	q := make(map[string]interface{})
	q["altcode"] = c.Params.ByName("altcode")
	q["pollingstationid"] = c.Params.ByName("pollingstationid")

	if err := database.DBCon.Where(q).First(&senatorial).Error; err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.BindJSON(&senatorial)
	database.DBCon.Save(&senatorial)
	c.IndentedJSON(http.StatusOK, senatorial)
}

func UpdateWomenRepVotesByCode(c *gin.Context) {
	var womenrep models.Womenrep
	q := make(map[string]interface{})
	q["altcode"] = c.Params.ByName("altcode")
	q["pollingstationid"] = c.Params.ByName("pollingstationid")

	if err := database.DBCon.Where(q).First(&womenrep).Error; err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.BindJSON(&womenrep)
	database.DBCon.Save(&womenrep)
	c.IndentedJSON(http.StatusOK, womenrep)
}

func UpdateMcaWomenRepVotesByCode(c *gin.Context) {
	var mca models.Mca
	q := make(map[string]interface{})
	q["altcode"] = c.Params.ByName("altcode")
	q["pollingstationid"] = c.Params.ByName("pollingstationid")

	if err := database.DBCon.Where(q).First(&mca).Error; err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.BindJSON(&mca)
	database.DBCon.Save(&mca)
	c.IndentedJSON(http.StatusOK, mca)
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

func UpdateParliamentaryForm(c *gin.Context) {
	var forms models.ParliamentaryForms
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

func UpdateSenatorialForm(c *gin.Context) {
	var forms models.SenatorialForms
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

func UpdateGubernatorialForm(c *gin.Context) {
	var forms models.GubernatorialForms
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

func UpdateWomenRepForm(c *gin.Context) {
	var forms models.WomenRepForms
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

func UpdateMcaForm(c *gin.Context) {
	var forms models.McaForms
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

func GetResultsByConstituency(c *gin.Context) {
	type Result struct {
		Candidate string `json:"candidate"`
		Votes     int    `json:"votes"`
		// Percentage int    `json:"percentage"`
	}

	var result []Result
	// round(votes*100/ sum(votes),2) as percentage
	if err := database.DBCon.Table("parliamentaries").Select("candidate, sum(votes) as votes").Group("candidate").Order("votes desc").Find(&result).Error; err != nil {
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

func GetParliamentaryFormByPollingId(c *gin.Context) {
	var forms models.ParliamentaryForms

	m := make(map[string]interface{})
	m["pollingstationid"] = c.Params.ByName("pollingstationid")

	if err := database.DBCon.Where(m).Find(&forms).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.IndentedJSON(http.StatusOK, forms)
}

func GetSenatorialFormByPollingId(c *gin.Context) {
	var forms models.SenatorialForms

	m := make(map[string]interface{})
	m["pollingstationid"] = c.Params.ByName("pollingstationid")

	if err := database.DBCon.Where(m).Find(&forms).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.IndentedJSON(http.StatusOK, forms)
}

func GetGubernatorialFormByPollingId(c *gin.Context) {
	var forms models.GubernatorialForms

	m := make(map[string]interface{})
	m["pollingstationid"] = c.Params.ByName("pollingstationid")

	if err := database.DBCon.Where(m).Find(&forms).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.IndentedJSON(http.StatusOK, forms)
}

func GetMcaFormByPollingId(c *gin.Context) {
	var forms models.McaForms

	m := make(map[string]interface{})
	m["pollingstationid"] = c.Params.ByName("pollingstationid")

	if err := database.DBCon.Where(m).Find(&forms).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.IndentedJSON(http.StatusOK, forms)
}

func GetWomenRepFormByPollingId(c *gin.Context) {
	var forms models.WomenRepForms

	m := make(map[string]interface{})
	m["pollingstationid"] = c.Params.ByName("pollingstationid")

	if err := database.DBCon.Where(m).Find(&forms).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.IndentedJSON(http.StatusOK, forms)
}
