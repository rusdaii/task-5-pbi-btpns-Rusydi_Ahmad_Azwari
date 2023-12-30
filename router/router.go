package router

import (
	"go-project/controllers"
	"go-project/middlewares"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Port() int {
	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		port = 5000
	}

	return port
}

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.SetTrustedProxies([]string{"127.0.0.1"})

	public := r.Group("/api")

	public.POST("/users/register", controllers.Register)
	public.POST("/users/login", controllers.Login)

	secured := r.Group("/api").Use(middlewares.Auth())
	{
		secured.PUT("/users/:userId", controllers.UpdateUser)
		secured.DELETE("/users/:userId", controllers.DeleteUser)

		secured.GET("/photos", controllers.GetAllPhotos)
		secured.GET("/photos/:photoId", controllers.GetPhotoById)
		secured.POST("/photos", controllers.CreatePhoto)
		secured.PUT("/photos/:photoId", controllers.UpdatePhoto)
		secured.DELETE("/photos/:photoId", controllers.DeletePhoto)
	}

	return r
}
