package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	"errors"
	"database/sql"
	"gorm.io/driver/postgres"
  	"gorm.io/gorm"

)

type Exercise struct {
	gorm.Model
	Name string `gorm:"uniqueIndex; not null"`
	Descr string
	MuscleGroupBroadName string
	PrimaryMuscles []PrimaryMuscle `gorm:"many2many:exercise_primarymuscles;"`
	SecondaryMuscles []SecondaryMuscle `gorm:"many2many:exercise_secondarymuscles;"`
	Equipment []Equipment `gorm:"many2many:exercise_equipments;"`
	Advanced string
	Superset string
	Intensity string
	CompoundMovement string
	ExerciseImages *[]string `gorm:"type:varchar(128)[]"`

}

type PrimaryMuscle struct {
	Name string `gorm:"primaryKey"`
	Name2 string //name_en needs to map to this one
	IsFront bool
	Image string
	Image2 string

}

type SecondaryMuscle struct {
	Name string `gorm:"primaryKey"`
	Name2 string //name_en needs to map to this one
	IsFront bool
	Image string
	Image2 string

}

type Equipment struct {
	Name string `gorm:"primaryKey"`
}

type MuscleGroupBroad struct {
	Name string `gorm:"primaryKey"`
	Exercises []Exercise `gorm:"foreignKey:MuscleGroupBroadName;association_foreignkey:Name"`
}



func init() {
	if _, err := os.Stat("./.env"); errors.Is(err, os.ErrNotExist) {
		// .env file does not exist (production)
		fmt.Println("cant find env")
	} else {
		// .env file exists (development)
		err2 := godotenv.Load(".env")

		if err2 != nil {
			log.Fatal("Error loading .env file")
		}
	}

	

}

func testingStuff(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hey the test is working")
}

func main() {
	fmt.Println("App running")
	port:= os.Getenv("PORT")
	fmt.Println(port)

	//db is a pointer to a sql.DB 
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Error opening database")
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("error connecting to database")
	}

	//gormDB is a pointer to a gorm.DB
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	//gormDB.Migrator().DropTable(&MuscleGroupBroad{}, &Exercise{}, &PrimaryMuscle{}, &SecondaryMuscle{}, &Equipment{}, "exercise_primarymuscles", "exercise_secondarymuscles", "exercise_equipments")

	gormDB.AutoMigrate(&MuscleGroupBroad{}, &Exercise{}, &PrimaryMuscle{}, &SecondaryMuscle{}, &Equipment{})

	exercises := GetData()
	//fmt.Println((*exercises)[6])

	/*
	m1 := MuscleGroupBroad{Name: "Abs"}
	m2 := MuscleGroupBroad{Name: "Arms"}
	m3 := MuscleGroupBroad{Name: "Back"}
	m4 := MuscleGroupBroad{Name: "Calves"}
	m5 := MuscleGroupBroad{Name: "Chest"}
	m6 := MuscleGroupBroad{Name: "Legs"}
	m7 := MuscleGroupBroad{Name: "Shoulders"}
	gormDB.Create(&m1)
	gormDB.Create(&m2)
	gormDB.Create(&m3)
	gormDB.Create(&m4)
	gormDB.Create(&m5)
	gormDB.Create(&m6)
	gormDB.Create(&m7)
	*/
	
	//gormDB.Create(&(*exercises))

	
	for i := range *exercises {
		gormDB.Create(&((*exercises)[i]))
	}
	
	
	mux := http.NewServeMux()
	mux.HandleFunc("/test", testingStuff)


	log.Fatal(http.ListenAndServe(":"+port, mux))
}