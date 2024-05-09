package models

import "fmt"

type Auto struct {
	Brand string `json:"brand"`
	Model string `json:"model"`
	Color string `json:"color"`
	Year  int    `json:"year"`
	Price int    `json:"price"`
}

var autos = make(map[uint]Auto)

func GetAutos() (autosList []Auto) {
	for _, a := range autos {
		autosList = append(autosList, a)
	}
	return
}

func GetAuto(id uint) (auto Auto, err error) {
	if a, ok := autos[id]; ok {
		return a, nil
	}
	return Auto{}, fmt.Errorf("Auto with id %d not found", id)
}

func CreateAuto(auto Auto) (status int) {
	autos[uint(len(autos)+1)] = auto
	return 1
}

func UpdateAuto(id uint, autoToUpdate Auto) (auto Auto, err error) {
	if _, ok := autos[id]; ok {
		auto = autoToUpdate
		autos[id] = auto
		return auto, nil
	}
	return Auto{}, fmt.Errorf("Auto with id %d not found", id)
}

func DeleteAuto(id uint) (status int) {
	if _, ok := autos[id]; ok {
		delete(autos, id)
		return 1
	}
	return 0
}
