package main

import (
	"azimio/controllers"
	"azimio/database"
	"azimio/middlewares"
	"azimio/models"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupRouter() *gin.Engine {
	// r := gin.Default()
	r := gin.New()
	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))

	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE"},
	// 	AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},

	// }))

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	api := r.Group("/api")
	{
		public := api.Group("/public")
		{
			public.POST("/login", controllers.Login)
			public.POST("/signup", controllers.Signup)
		}

		protected := api.Group("/auth").Use(middlewares.Authz())
		{
			protected.GET("/profile", controllers.Profile)
			protected.GET("/pollingdata", controllers.GetPollingData)
			protected.GET("/pollingdata/:pollingstationid", controllers.GetPollingDataById)
			protected.PUT("/presidential/votes/:id", controllers.UpdatePresidentialVotes)
			protected.PUT("/forms/:id", controllers.UpdateForm)
			protected.GET("/forms/:pollingstationid", controllers.GetFormByPollingId)
			protected.GET("/presidential/county/results", controllers.GetResultsByCounty)
			protected.GET("/presidential/country/results", controllers.GetResultsByCountry)
			protected.GET("presidential/country/total", controllers.GetTotalVotes)
			protected.PUT("presidential/alt/:pollingstationid/:altcode", controllers.UpdatePresidentialVotesByCode)

			protected.PUT("parliamentary/alt/:pollingstationid/:altcode", controllers.UpdateParliamentaryVotesByCode)
			protected.PUT("/parliamentary/forms/:id", controllers.UpdateParliamentaryForm)
			protected.GET("/parliamentarydata", controllers.GetParliamentaryData)
			protected.GET("/parliamentary/results", controllers.GetResultsByConstituency)
			protected.GET("/parliamentarydata/:pollingstationid", controllers.GetParliamentaryPollingDataById)
			protected.PUT("/parliamentary/votes/:id", controllers.UpdateParliamentaryVotes)
			protected.GET("/parliamentary/forms/:pollingstationid", controllers.GetParliamentaryFormByPollingId)

			protected.PUT("gubernatorial/alt/:pollingstationid/:altcode", controllers.UpdateGubernatorialVotesByCode)
			protected.PUT("/gubernatorial/forms/:id", controllers.UpdateGubernatorialForm)

			protected.PUT("senatorial/alt/:pollingstationid/:altcode", controllers.UpdateSenatorialVotesByCode)
			protected.PUT("/senatorial/forms/:id", controllers.UpdateSenatorialForm)

			protected.PUT("womenrep/alt/:pollingstationid/:altcode", controllers.UpdateWomenRepVotesByCode)
			protected.PUT("/womenrep/forms/:id", controllers.UpdateWomenRepForm)

			protected.PUT("mca/alt/:pollingstationid/:altcode", controllers.UpdateMcaWomenRepVotesByCode)
			protected.PUT("/mca/forms/:id", controllers.UpdateMcaForm)
		}
	}

	return r
}

func main() {

	dsn := database.DB_USERNAME + ":" + database.DB_PASSWORD + "@tcp" + "(" + database.DB_HOST + ":" + database.DB_PORT + ")/" + database.DB_NAME + "?" + "parseTime=true&loc=Local"
	database.DBCon, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{PrepareStmt: true})
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	database.DBCon.AutoMigrate(&models.User{})
	database.DBCon.AutoMigrate(&models.Forms{})
	database.DBCon.AutoMigrate(&models.GubernatorialForms{})
	database.DBCon.AutoMigrate(&models.SenatorialForms{})
	database.DBCon.AutoMigrate(&models.ParliamentaryForms{})
	database.DBCon.AutoMigrate(&models.WomenRepForms{})
	database.DBCon.AutoMigrate(&models.McaForms{})
	database.DBCon.AutoMigrate(&models.Pollingdata{})
	database.DBCon.AutoMigrate(&models.Gubernatorial{})
	database.DBCon.AutoMigrate(&models.Parliamentary{})
	database.DBCon.AutoMigrate(&models.Senatorial{})
	database.DBCon.AutoMigrate(&models.Mca{})
	database.DBCon.AutoMigrate(&models.Womenrep{})

	// // CREATE TEST DATA

	// // User Data
	// database.DBCon.Create(&models.User{Username: "OBrien", Firstname: "Brian", Lastname: "Otieno", Phonenumber: "+254723328969", Email: "gebryo@intelligencia.com", Password: "$2a$14$SaSgFyNhW9ncAmMf19BTg.wSlAV2dctl/MXNsSdpYKupJE6AWAhpy", Pollingstationid: "2", Pollingstation: "MAWEGO BOARDING PRIMARY", Registerdvoters: 1350, Role: 1, Approved: true})
	// database.DBCon.Create(&models.User{Username: "jdoe", Firstname: "Jane", Lastname: "Doe", Phonenumber: "+254713846445", Email: "jane.doe.com", Password: "$2a$14$hTW2RgNLwKpnG4tNVKMZLed2thwKqtSuB7OOJhsJmKg8HD8/FRjkS", Pollingstationid: "1", Pollingstation: "KOKWANYO PRIMARY SCHOOL", Registerdvoters: 1350, Role: 1, Approved: true})

	// // Polling Data
	// database.DBCon.Create(&models.Pollingdata{Pollingstationid: "2", Candidate: "Raila Odinga", Altcode: "C01", Scid: "403", Ccode: "043", Cname: "HOMABAY", Scname: "RACHUONYO SOUTH", Pollingstation: "MAWEGO BOARDING PRIMARY", Votes: 825})
	// database.DBCon.Create(&models.Pollingdata{Pollingstationid: "2", Candidate: "William Ruto", Altcode: "C02", Scid: "403", Ccode: "043", Cname: "HOMABAY", Scname: "RACHUONYO SOUTH", Pollingstation: "MAWEGO BOARDING PRIMARY", Votes: 2})
	// database.DBCon.Create(&models.Pollingdata{Pollingstationid: "2", Candidate: "Wajackoyah", Altcode: "C03", Scid: "403", Ccode: "043", Cname: "HOMABAY", Scname: "RACHUONYO SOUTH", Pollingstation: "MAWEGO BOARDING PRIMARY", Votes: 1})
	// database.DBCon.Create(&models.Pollingdata{Pollingstationid: "2", Candidate: "Waihiga", Altcode: "C04", Scid: "403", Ccode: "043", Cname: "HOMABAY", Scname: "RACHUONYO SOUTH", Pollingstation: "MAWEGO BOARDING PRIMARY", Votes: 0})

	// database.DBCon.Create(&models.Pollingdata{Pollingstationid: "1", Candidate: "Raila Odinga", Altcode: "C01", Scid: "403", Ccode: "043", Cname: "HOMABAY", Scname: "RACHUONYO SOUTH", Pollingstation: "KOKWANYO PRIMARY SCHOOL", Votes: 920})
	// database.DBCon.Create(&models.Pollingdata{Pollingstationid: "1", Candidate: "William Ruto", Altcode: "C02", Scid: "403", Ccode: "043", Cname: "HOMABAY", Scname: "RACHUONYO SOUTH", Pollingstation: "KOKWANYO PRIMARY SCHOOL", Votes: 1})
	// database.DBCon.Create(&models.Pollingdata{Pollingstationid: "1", Candidate: "Wajackoyah", Altcode: "C03", Scid: "403", Ccode: "043", Cname: "HOMABAY", Scname: "RACHUONYO SOUTH", Pollingstation: "KOKWANYO PRIMARY SCHOOL", Votes: 1})
	// database.DBCon.Create(&models.Pollingdata{Pollingstationid: "1", Candidate: "William Ruto", Altcode: "C04", Scid: "403", Ccode: "043", Cname: "HOMABAY", Scname: "RACHUONYO SOUTH", Pollingstation: "KOKWANYO PRIMARY SCHOOL", Votes: 0})

	// database.DBCon.Create(&models.Pollingdata{Pollingstationid: "3", Candidate: "Raila Odinga", Altcode: "C01", Scid: "403", Ccode: "043", Cname: "HOMABAY", Scname: "RACHUONYO SOUTH", Pollingstation: "UMAI PRIMARY SCHOOL", Votes: 730})
	// database.DBCon.Create(&models.Pollingdata{Pollingstationid: "3", Candidate: "William Ruto", Altcode: "C02", Scid: "403", Ccode: "043", Cname: "HOMABAY", Scname: "RACHUONYO SOUTH", Pollingstation: "UMAI PRIMARY SCHOOL", Votes: 1})
	// database.DBCon.Create(&models.Pollingdata{Pollingstationid: "3", Candidate: "Wajackoyah", Altcode: "C03", Scid: "403", Ccode: "043", Cname: "HOMABAY", Scname: "RACHUONYO SOUTH", Pollingstation: "UMAI PRIMARY SCHOOL", Votes: 1})
	// database.DBCon.Create(&models.Pollingdata{Pollingstationid: "3", Candidate: "Waihiga", Altcode: "C04", Scid: "403", Ccode: "043", Cname: "HOMABAY", Scname: "RACHUONYO SOUTH", Pollingstation: "UMAI PRIMARY SCHOOL", Votes: 1})

	// // Form 34A
	// database.DBCon.Create(&models.Forms{Pollingstationid: "1"})
	// database.DBCon.Create(&models.Forms{Pollingstationid: "2"})
	// database.DBCon.Create(&models.Forms{Pollingstationid: "3"})

	r := setupRouter()
	r.Run("127.0.0.1:8080")
}
