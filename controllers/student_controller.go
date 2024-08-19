package controllers

import (
	"demo-project/config"
	"demo-project/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetStudents(c *gin.Context) {
	var students []models.Student
	config.DB.Find(&students)
	c.JSON(http.StatusOK, students)
}

func GetStudentByID(c *gin.Context) {
	id := c.Param("id")
	var student models.Student
	if err := config.DB.Where("id = ?", id).First(&student).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, student)
}

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, student)
}

func UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student
	if err := config.DB.Where("id = ?", id).First(&student).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&student)
	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student
	if err := config.DB.Where("id = ?", id).First(&student).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	config.DB.Delete(&student)
	c.JSON(http.StatusOK, gin.H{"message": "Student deleted"})
}
