package routes

import (
	"demo-project/controllers"
	"demo-project/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// Public routes
	r.POST("/api/register", controllers.RegisterUser)
	r.POST("/api/login", controllers.LoginUser)

	// Protected routes
	protected := r.Group("/")
	protected.Use(middleware.JwtAuthMiddleware())
	protected.GET("/api/students", controllers.GetStudents)
	protected.GET("/api/students/:id", controllers.GetStudentByID)
	protected.POST("/api/students", controllers.CreateStudent)
	protected.PUT("/api/students/:id", controllers.UpdateStudent)
	protected.DELETE("/api/students/:id", controllers.DeleteStudent)
}
