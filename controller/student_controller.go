package controller

import (
	"app-mahasiswa-api2/model"
	"app-mahasiswa-api2/usecase"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type StudentController struct {
	usecase usecase.StudentUsecase
}

func (c *StudentController) FindUStudents(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	username := claims["username"].(string)

	res := c.usecase.FindStudents()
	ctx.JSON(http.StatusOK, gin.H{
		"data":     res,
		"username": username,
	})
}

func (c *StudentController) FindStudentById(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	username := claims["username"].(string)

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid Student id",
		})
		return
	}

	res := c.usecase.FindStudentById(id)
	ctx.JSON(http.StatusOK, gin.H{
		"data":     res,
		"username": username,
	})
}

func (c *StudentController) Register(ctx *gin.Context) {
	var newStudent model.Student

	if err := ctx.BindJSON(&newStudent); err != nil {
		ctx.JSON(http.StatusBadRequest, "invalid input")
		return
	}

	res := c.usecase.Register(&newStudent)

	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

func (c *StudentController) Edit(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	username := claims["username"].(string)

	var student model.Student

	if err := ctx.BindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		return
	}

	res := c.usecase.Edit(&student)
	ctx.JSON(http.StatusOK, gin.H{
		"data":     res,
		"username": username,
	})
}

func (c *StudentController) Unreg(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	username := claims["username"].(string)

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid student id")
		return
	}
	c.usecase.UnregUser(username)

	res := c.usecase.Unreg(id)
	ctx.JSON(http.StatusOK, gin.H{
		"data":     res,
		"username": username,
	})
}

func NewStudentController(u usecase.StudentUsecase) *StudentController {
	controller := StudentController{
		usecase: u,
	}

	return &controller
}
