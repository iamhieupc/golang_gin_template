package main

import (
	"github.com/joho/godotenv"
	"log"
	// "fmt"
	// "runtime"
	"os"
	"github.com/gin-gonic/gin"
	db "hustchihieu/todolist-golang/database"
	models "hustchihieu/todolist-golang/models"
	routes "hustchihieu/todolist-golang/routes"
	// repositories "hustchihieu/todolist-golang/repositories"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}


func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error: failed to load the env file")
	}
	
	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	//Start the default gin server
	r := gin.Default()

	r.Use(CORSMiddleware())

	
	db := db.ConnectDB()

	db.AutoMigrate(&models.User{})

	// defer db.Close()


	// v1 := r.Group("/api")
	// {
	// 	/*** START USER ***/
	// 	user := new(controllers.UserController)

	// 	v1.GET("/", user.HealthCheck)

	// 	v1.GET("/user/", user.FindAllUsers)
		
	// }

	r.LoadHTMLGlob("./public/html/*")

	r.Static("/public", "./public")

	// userRepository := repositories.NewUserRepository(db)

	route := routes.SetupRoutes()


	// r.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.html", gin.H{
	// 		"ginBoilerplateVersion": "v0.03",
	// 		"goVersion":             runtime.Version(),
	// 	})
	// })

	// r.NoRoute(func(c *gin.Context) {
	// 	c.HTML(404, "404.html", gin.H{})
	// })

	port := os.Getenv("PORT")

	log.Printf("\n\n PORT: %s \n ENV: %s \n SSL: %s \n Version: %s \n\n", port, os.Getenv("ENV"), os.Getenv("SSL"), os.Getenv("API_VERSION"))

	route.Run(":8000")

}