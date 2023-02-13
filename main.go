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
	"encoding/json"
	"github.com/lib/pq"

)

type Exercise struct {
	gorm.Model
	Name string `gorm:"uniqueIndex; not null" json:"name"`
	Descr string `json:"descr"`
	MuscleGroupBroadName string `json:"muscleGroupBroadName"`
	PrimaryMuscles []PrimaryMuscle `gorm:"many2many:exercise_primarymuscles;" json:"primaryMuscles"`
	SecondaryMuscles []SecondaryMuscle `gorm:"many2many:exercise_secondarymuscles;" json:"secondaryMuscles"`
	Equipment []Equipment `gorm:"many2many:exercise_equipments;" json:"equipment"`
	Advanced string `json:"advanced"`
	Superset string `json:"superset"`
	Intensity string `json:"intensity"`
	CompoundMovement string `json:"compoundMovement"`
	ExerciseImages pq.StringArray `gorm:"type:varchar(128)[]" json:"exerciseImages"`

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

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
}

func testingStuff(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hey the test is working")
}




func getExerciseDataHandler(w http.ResponseWriter, req *http.Request, db *gorm.DB) {
	if req.Method == http.MethodGet {
		var exercises []Exercise
		fmt.Println("handling get all tasks at %s\n", req.URL.Path)
		db.Find(&exercises)
		js, err := json.Marshal(exercises)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

func getPrimaryMuscleDataHandler(w http.ResponseWriter, req *http.Request, db *gorm.DB) {
	if req.Method == http.MethodGet {
		var primaryMuscles []PrimaryMuscle
		fmt.Println("handling get all tasks at %s\n", req.URL.Path)
		db.Find(&primaryMuscles)
		js, err := json.Marshal(primaryMuscles)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

func getSecondaryMuscleDataHandler(w http.ResponseWriter, req *http.Request, db *gorm.DB) {
	if req.Method == http.MethodGet {
		var secondaryMuscles []SecondaryMuscle
		fmt.Println("handling get all tasks at %s\n", req.URL.Path)
		db.Find(&secondaryMuscles)
		js, err := json.Marshal(secondaryMuscles)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

func getEquipmentDataHandler(w http.ResponseWriter, req *http.Request, db *gorm.DB) {
	if req.Method == http.MethodGet {
		var equipment []Equipment
		fmt.Println("handling get all tasks at %s\n", req.URL.Path)
		db.Find(&equipment)
		js, err := json.Marshal(equipment)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

func getMuscleGroupBroadDataHandler(w http.ResponseWriter, req *http.Request, db *gorm.DB) {
	if req.Method == http.MethodGet {
		var muscleGroupBroad []MuscleGroupBroad
		fmt.Println("handling get all tasks at %s\n", req.URL.Path)
		db.Find(&muscleGroupBroad)
		js, err := json.Marshal(muscleGroupBroad)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}


func getAllExerciseData(w http.ResponseWriter, req *http.Request, db *gorm.DB) {
	enableCors(&w)
	if req.Method == http.MethodGet {
		var ex []Exercise
		fmt.Println("handling get all tasks at %s\n", req.URL.Path)
		db.Preload("PrimaryMuscles").Preload("SecondaryMuscles").Preload("Equipment").Find(&ex)
		//fmt.Println(ex)
		js, err := json.Marshal(ex)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
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

	/* this line gets all the exercises from API for insertion */
	//exercises := GetData()

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

	/* this creates entires in the database for each exercise */
	/*
	for i := range *exercises {
		gormDB.Create(&((*exercises)[i]))
	}
	*/
	
	mux := http.NewServeMux()
	mux.HandleFunc("/test", testingStuff)
	mux.HandleFunc("/GetExerciseData", func(w http.ResponseWriter, req *http.Request) {
		getExerciseDataHandler(w, req, gormDB)
	})
	mux.HandleFunc("/GetPrimaryMuscleData", func(w http.ResponseWriter, req *http.Request) {
		getPrimaryMuscleDataHandler(w, req, gormDB)
	})
	mux.HandleFunc("/GetSecondaryMuscleData", func(w http.ResponseWriter, req *http.Request) {
		getSecondaryMuscleDataHandler(w, req, gormDB)
	})
	mux.HandleFunc("/GetEquipmentData", func(w http.ResponseWriter, req *http.Request) {
		getEquipmentDataHandler(w, req, gormDB)
	})
	mux.HandleFunc("/GetMuscleGroupBroadData", func(w http.ResponseWriter, req *http.Request) {
		getMuscleGroupBroadDataHandler(w, req, gormDB)
	})

	mux.HandleFunc("/GetAllExerciseData", func(w http.ResponseWriter, req *http.Request) {
		getAllExerciseData(w, req, gormDB)
	})

	strengthLoopTesting(7080, 180, 90, 4, 1)
	//log.Fatal(http.ListenAndServe(":"+port, mux))
}