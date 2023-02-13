package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
)

func extraSetsHelper(restTimeHypertrophy int, remainingTime int) int {
	var maxNumExtraSets int = remainingTime / (restTimeHypertrophy + 90)
	fmt.Println("Remaining time: " + strconv.Itoa(remainingTime))
	fmt.Println("Max num of extra sets: " + strconv.Itoa(maxNumExtraSets))
	if maxNumExtraSets >= 2 {
		rand.Seed(time.Now().UnixNano())
		actualNumExtraSets := rand.Intn(maxNumExtraSets - 2 + 1) + 2 //random number from 2 - max # sets
		return actualNumExtraSets
	}
	return 0
}

func strengthLoopTesting(workoutLength int, restTimeStrength int, restTimeHypertrophy int, setsPerExercise int, warmUp int) {
	rand.Seed(time.Now().UnixNano())

	strengthExerciseLength := (restTimeStrength + 90) * setsPerExercise + (warmUp * 240)
	hypertrophyCompoundExerciseLength := (restTimeHypertrophy + 90) * setsPerExercise + (warmUp * 240)
	hypertrophyNonCompoundExerciseLength := (restTimeHypertrophy + 90) * setsPerExercise + (warmUp * 90)

	typeExercise := "strength"
	numStrength := 0
	numHypertrophyCompound := 0
	numHypertrophyNonCompound := 0
	desiredNumStrength := rand.Intn(3) + 1 //random number 1 - 3
	//desiredNumStrength := 1
	fmt.Println(desiredNumStrength)
	lastWasHypertrophyCompound := false

	//actualLength := 0



	i := workoutLength
	for true {
		if typeExercise == "strength" {
			i -= strengthExerciseLength
			numStrength++


		} else if typeExercise == "hypertrophyCompound" {
			i -= hypertrophyCompoundExerciseLength
			numHypertrophyCompound++
			lastWasHypertrophyCompound = true
		} else if typeExercise == "hypertrophyNonCompound" {
			i -= hypertrophyNonCompoundExerciseLength
			numHypertrophyNonCompound++
			lastWasHypertrophyCompound = false
		}
		
		if numStrength == desiredNumStrength {
			if lastWasHypertrophyCompound {
				if i - hypertrophyNonCompoundExerciseLength >= -350 {
					typeExercise = "hypertrophyNonCompound"
				} else {
					extraSets := extraSetsHelper(restTimeHypertrophy, i)
					if (extraSets > 0) {
						numHypertrophyNonCompound++
					}
					fmt.Println("extra sets: " + strconv.Itoa(extraSets))
					break
				} 
			} else {
				if i - hypertrophyCompoundExerciseLength >= -350 { 
					typeExercise = "hypertrophyCompound"
				} else if i - hypertrophyNonCompoundExerciseLength >= -350 {
					typeExercise = "hypertrophyNonCompound"
				} else {
					extraSets := extraSetsHelper(restTimeHypertrophy, i)
					if (extraSets > 0) {
						numHypertrophyNonCompound++
					}
					fmt.Println("extra sets: " + strconv.Itoa(extraSets))
					break
				} 
			}
		} else {
			if i - strengthExerciseLength >= -240 { //4 mins variation
				typeExercise = "strength"
			} else if i - hypertrophyCompoundExerciseLength >= -350 { 
				typeExercise = "hypertrophyCompound"
			} else if i - hypertrophyNonCompoundExerciseLength >= -350 {
				typeExercise = "hypertrophyNonCompound"
			} else {
				extraSets := extraSetsHelper(restTimeHypertrophy, i)
				if (extraSets > 0) {
					numHypertrophyNonCompound++
				}
				fmt.Println("extra sets: " + strconv.Itoa(extraSets))
				break
			} 
		}

	}

	fmt.Printf("# Strength: %d # Hypertrophy Compound %d # Hypertrophy NonCompound %d \n", numStrength, numHypertrophyCompound, numHypertrophyNonCompound)
}