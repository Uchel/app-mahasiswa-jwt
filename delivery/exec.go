package delivery

import (
	"app-mahasiswa-api2/config"
	"app-mahasiswa-api2/controller"
	"app-mahasiswa-api2/repository"
	"app-mahasiswa-api2/usecase"
	"log"

	jwt_auth "github.com/Uchel/auth-jwt"

	"github.com/gin-gonic/gin"
)

func Exec() {

	r := gin.Default()

	db := config.ConnectDB()
	defer db.Close()
	// user
	userRepo := jwt_auth.NewUserRepo(db)
	userUseCase := jwt_auth.NewUserUsecase(userRepo)
	userctrl := jwt_auth.NewUserController(userUseCase)

	// userRepo := authjwtrussel.NewUserRepo(db)
	// userUseCase := authjwtrussel.NewUserUsecase(userRepo)
	// userctrl := authjwtrussel.NewStudentController(userUseCase)

	//student
	studentRepo := repository.NewStudentRepo(db)
	studentUsecase := usecase.NewStudentUsecase(studentRepo)
	studentctrl := controller.NewStudentController(studentUsecase)

	//register
	studentRegister := r.Group("/api/v1/register")
	studentRegister.POST("", studentctrl.Register)

	//login
	r.POST("/auth/login", userctrl.Login)

	//login => jwt => akses api/v1/students
	studentRouter := r.Group("/api/v1/students")
	studentRouter.Use(jwt_auth.AuthMiddleware())

	studentRouter.GET("", studentctrl.FindUStudents)
	studentRouter.GET("/:id", studentctrl.FindStudentById)
	studentRouter.PUT("", studentctrl.Edit)
	studentRouter.DELETE("/:id", studentctrl.Unreg)

	if err := r.Run(":7000"); err != nil {
		log.Fatal(err)
	}
}
