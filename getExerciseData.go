package main

import (
	//"fmt"
	"net/http"
	"log"
	"io"
	"encoding/json"
)

type ExercisesGet struct {
	ExercisesGet []ExerciseGet`json:"results"`
}

type ExerciseGet struct {
	Name string `json:"name"`
	Descr string `json:"description"`
	Muscle_group_broad Muscle_group_broad `json:"category"`
	Primary_muscles_targeted []MuscleGet `json:"muscles"`
	Secondary_muscles_targeted []MuscleGet `json:"muscles_secondary"`
	Equipment []EquipmentGet `json:"equipment"`
	Language Language `json:"language"`
	Exercise_images []Image `json:"images"`
}

type MuscleGet struct {
	Muscle_group_specific_technical_name string `json:"name"`
	Muscle_group_specific string `json:"name_en"`
	Is_front bool `json:"is_front"`
	Muscle_group_specific_image string `json:"image_url_main"`
	Muscle_group_specific_image2 string `json:"image_url_secondary"`

}

type Image struct {
	Exercise_image string `json:"image"`
	Is_main bool `json:"is_main"`
}

type EquipmentGet struct {
	Equipment_name string `json:"name"`
}

type Muscle_group_broad struct {
	Muscle_group_broad_name string `json:"name"`
}

type Language struct {
	Id int `json:"id"`
	Short_name string `json:"short_name"`
}

func makeImagesStrings(curr []Image) *[]string {
	ret := make([]string, len(curr))
	for i := range ret {
		ret[i] = curr[i].Exercise_image
	}
	return &ret
}

func GetData() *[]Exercise {
	resp, err := http.Get("https://wger.de/api/v2/exerciseinfo/?limit=419")
	if err != nil {
		log.Fatal("error retrieving from API")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("failed to read data received from API")
	}

	var ex ExercisesGet
	json.Unmarshal(body, &ex)

	exerciseList := make([]Exercise, len(ex.ExercisesGet))

	for i := 0; i < len(ex.ExercisesGet); i++ {
		
		if ex.ExercisesGet[i].Language.Id == 2 {
			currPrimaryMuscles := make([]PrimaryMuscle, len(ex.ExercisesGet[i].Primary_muscles_targeted))
			currSecondaryMuscles := make([]SecondaryMuscle, len(ex.ExercisesGet[i].Secondary_muscles_targeted))
			currEquipment := make([]Equipment, len(ex.ExercisesGet[i].Equipment))

			for j := range ex.ExercisesGet[i].Primary_muscles_targeted {
				currPrimaryMuscles[j] = PrimaryMuscle{Name: ex.ExercisesGet[i].Primary_muscles_targeted[j].Muscle_group_specific_technical_name,
													Name2: ex.ExercisesGet[i].Primary_muscles_targeted[j].Muscle_group_specific,
													IsFront: ex.ExercisesGet[i].Primary_muscles_targeted[j].Is_front,
													Image: ex.ExercisesGet[i].Primary_muscles_targeted[j].Muscle_group_specific_image,
													Image2: ex.ExercisesGet[i].Primary_muscles_targeted[j].Muscle_group_specific_image2}
			}
			
			for j := range ex.ExercisesGet[i].Secondary_muscles_targeted {
				currSecondaryMuscles[j] = SecondaryMuscle{Name: ex.ExercisesGet[i].Secondary_muscles_targeted[j].Muscle_group_specific_technical_name,
													Name2: ex.ExercisesGet[i].Secondary_muscles_targeted[j].Muscle_group_specific,
													IsFront: ex.ExercisesGet[i].Secondary_muscles_targeted[j].Is_front,
													Image: ex.ExercisesGet[i].Secondary_muscles_targeted[j].Muscle_group_specific_image,
													Image2: ex.ExercisesGet[i].Secondary_muscles_targeted[j].Muscle_group_specific_image2}
			}

			for j := range ex.ExercisesGet[i].Equipment {
				currEquipment[j] = Equipment{Name: ex.ExercisesGet[i].Equipment[j].Equipment_name}
			}

			exerciseList[i] = Exercise{Name: ex.ExercisesGet[i].Name,
									Descr: ex.ExercisesGet[i].Descr,
									MuscleGroupBroadName: ex.ExercisesGet[i].Muscle_group_broad.Muscle_group_broad_name,
									PrimaryMuscles: currPrimaryMuscles,
									SecondaryMuscles: currSecondaryMuscles,
									Equipment: currEquipment,
									ExerciseImages: makeImagesStrings(ex.ExercisesGet[i].Exercise_images)} 
		}
		
	}



	return &exerciseList
	//fmt.Println(ex.ExercisesGet[11])

}
