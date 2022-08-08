package controllers

import (
	"azimio/auth"
	"azimio/database"
	"azimio/models"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Signup creates a user in db
func Signup(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Println(err)

		c.JSON(400, gin.H{
			"msg": "invalid json",
		})
		c.Abort()

		return
	}

	err = user.HashPassword(user.Password)
	if err != nil {
		log.Println(err.Error())

		c.JSON(500, gin.H{
			"msg": "error hashing password",
		})
		c.Abort()

		return
	}

	err = user.CreateUserRecord()
	if err != nil {
		log.Println(err)

		c.JSON(409, gin.H{
			"msg": "User Exist",
		})
		c.Abort()

		return
	}

	c.JSON(200, user)
}

// LoginPayload login body
type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse token response
type LoginResponse struct {
	Firstname        string `json:"firstname"`
	LastName         string `json:"lastname"`
	Username         string `json:"username"`
	Token            string `json:"token"`
	Role             int    `json:"role"`
	Pollingstationid string `json:"pollingstationid"`
	Phonenumber      string `json:"phonenumber"`
	Pollingstation   string `json:"pollingstation"`
	Registerdvoters  int    `json:"registeredvoters"`
}

// Login logs users in
func Login(c *gin.Context) {
	var payload LoginPayload
	var user models.User

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid json",
		})
		c.Abort()
		return
	}

	result := database.DBCon.Where("username = ?", payload.Username).First(&user)

	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(401, gin.H{
			"msg": "invalid user credentials",
		})
		c.Abort()
		return
	}

	err = user.CheckPassword(payload.Password)
	if err != nil {
		log.Println(err)
		c.JSON(401, gin.H{
			"msg": "invalid user credentials",
		})
		c.Abort()
		return
	}

	jwtWrapper := auth.JwtWrapper{
		SecretKey:       "verysecretkey",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	signedToken, err := jwtWrapper.GenerateToken(user.Username)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"msg": "error signing token",
		})
		c.Abort()
		return
	}

	tokenResponse := LoginResponse{
		Firstname:        user.Firstname,
		LastName:         user.Lastname,
		Username:         user.Username,
		Token:            signedToken,
		Role:             user.Role,
		Pollingstationid: user.Pollingstationid,
		Phonenumber:      user.Phonenumber,
		Pollingstation:   user.Pollingstation,
		Registerdvoters:  user.Registerdvoters,
	}

	c.JSON(200, tokenResponse)

	return
}
