package main

import (
	"database/sql"
	"dental_clinic_go/cmd/server/handler"
	"dental_clinic_go/docs"
	"dental_clinic_go/internal/appointment"
	"dental_clinic_go/internal/dentist"
	"dental_clinic_go/internal/patient"
	"dental_clinic_go/pkg/middleware"
	"dental_clinic_go/pkg/store"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Dental API
// @version 1.0
// @description This API handle a Dental Clinic .
// @termsOfService https://developers.ctd.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.ctd.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	/* -------------------- Cargamos las variables de entorno ------------------- */
	// if err := godotenv.Load(".env"); err != nil {
	// 	panic("error loading .env file")
	// }
	DB_URL := os.Getenv("DB_URL")
	HOST := os.Getenv("HOST")
	PORT := os.Getenv("PORT")

	/* ----------------------- Levantamos la base de datos ---------------------- */
	db, err := sql.Open("mysql", DB_URL)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	errPing := db.Ping()
	if errPing != nil {
		fmt.Println("DB_URL " + DB_URL)
		fmt.Println("HOST " + HOST)
		fmt.Println("PORT " + PORT)
		panic(errPing.Error())
	}

	/* -------------------- Instanciamos gin-gonic y swagger -------------------- */
	r := gin.Default()
	docs.SwaggerInfo.Host = HOST
	r.GET("/docs/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler))

	/* --------------------------------- Dentists ------------------------------- */
	dentistStorage := store.NewDentistSqlStore(db)
	dentistRepo := dentist.NewDentistRepository(dentistStorage)
	dentistService := dentist.NewDentistService(dentistRepo)
	dentistHandler := handler.NewDentistHandler(dentistService)

	dentists := r.Group("/dentists")
	{
		dentists.POST("", middleware.Authentication(), dentistHandler.Post())
		dentists.GET(":id", dentistHandler.GetByID())
		dentists.PUT(":id", middleware.Authentication(), dentistHandler.Put())
		dentists.PATCH(":id", middleware.Authentication(), dentistHandler.Patch())
		dentists.DELETE(":id", middleware.Authentication(), dentistHandler.Delete())
	}

	/* --------------------------------- Patients ------------------------------- */
	patientStorage := store.NewPatientSqlStore(db)
	patientRepo := patient.NewPatientRepository(patientStorage)
	patientService := patient.NewPatientService(patientRepo)
	patientHandler := handler.NewPatientHandler(patientService)

	patients := r.Group("/patients")
	{
		patients.POST("", middleware.Authentication(), patientHandler.Post())
		patients.GET(":id", patientHandler.GetByID())
		patients.PUT(":id", middleware.Authentication(), patientHandler.Put())
		patients.PATCH(":id", middleware.Authentication(), patientHandler.Patch())
		patients.DELETE(":id", middleware.Authentication(), patientHandler.Delete())
	}

	/* ------------------------------- Appointment ------------------------------ */
	appointmentStorage := store.NewAppointmentSqlStore(db)
	appointmentRepo := appointment.NewAppointmentRepository(appointmentStorage, patientStorage, dentistStorage)
	appointmentService := appointment.NewAppointmentService(appointmentRepo)
	appointmentHandler := handler.NewAppointmentHandler(appointmentService)

	appointments := r.Group("/appointments")
	{
		appointments.POST("", middleware.Authentication(), appointmentHandler.Post())
		appointments.POST("/dni/license", middleware.Authentication(), appointmentHandler.PostByDniAndLicense())
		appointments.GET(":id", appointmentHandler.GetByID())
		appointments.GET("/dni/:dni", appointmentHandler.GetByDni())
		appointments.PUT(":id", middleware.Authentication(), appointmentHandler.Put())
		appointments.PATCH(":id", middleware.Authentication(), appointmentHandler.Patch())
		appointments.DELETE(":id", middleware.Authentication(), appointmentHandler.Delete())
	}

	r.Run(fmt.Sprintf(":%s", PORT))
}
