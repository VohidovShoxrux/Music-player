package service

import (
	"fmt"
	"musiclibrary/models"
)

func CreateMusic(name, author, genre string) {

	var musics models.Music = models.Music{
		ID:     len(models.ArrayMusic) + 1,
		Name:   name,
		Author: author,
		Genre:  genre,
	}

	models.ArrayMusic = append(models.ArrayMusic, musics)

}
func ReadMusic() {
	fmt.Println("__________________________________________")
	for i := 0; i < len(models.ArrayMusic); i++ {
		fmt.Println("Musiqa ID:", models.ArrayMusic[i].ID)
		fmt.Println("Musiqa nomi:", models.ArrayMusic[i].Name)
		fmt.Println("Musiqa Ijrochisi:", models.ArrayMusic[i].Author)
		fmt.Println("Musiqa turi:", models.ArrayMusic[i].Genre)
		fmt.Println("___________________________________________")
	}
}
func UpdateMusic(id int, name, author, genre string) {
	for i := 0; i < len(models.ArrayMusic); i++ {
		if models.ArrayMusic[i].ID == id {
			models.ArrayMusic[i].Name = name
			models.ArrayMusic[i].Author = author
			models.ArrayMusic[i].Genre = genre
		} else {
			fmt.Println("Bu ID: dagi musiqa ma'lumotlari topilmadi", models.ArrayMusic[i].ID)
		}
	}
	return
}
func DeleteMusic(id int) {
	for i := 0; i < len(models.ArrayMusic); i++ {
		if models.ArrayMusic[i].ID == id {
			models.ArrayMusic = append(models.ArrayMusic[:i], models.ArrayMusic[i+1:]...)
		} else {
			fmt.Println("Bu ID: dagi musiqa ma'lumptlari topilmadi", models.ArrayMusic[i].ID)
		}
	}
}
func GetbyidMusic(id int) {
	fmt.Println("__________________________________________")
	for i := 0; i < len(models.ArrayMusic); i++ {
		if models.ArrayMusic[i].ID == id {
			fmt.Println("Musiqa ID:", models.ArrayMusic[i].ID)
			fmt.Println("Musiqa nomi:", models.ArrayMusic[i].Name)
			fmt.Println("Musiqa Ijrochisi:", models.ArrayMusic[i].Author)
			fmt.Println("Musiqa turi:", models.ArrayMusic[i].Genre)
			fmt.Println("___________________________________________")
			break
		}else {
			fmt.Println("Bu ID: dagi musiqa ma'lumotlari topilmadi", models.ArrayMusic[i].ID)
		}
	} 
	return
}
func GetbynameMusic(name string) {
	fmt.Println("__________________________________________")
	for i := 0; i < len(models.ArrayMusic); i++ {
		if models.ArrayMusic[i].Name == name {
			fmt.Println("Musiqa ID:", models.ArrayMusic[i].ID)
			fmt.Println("Musiqa nomi:", models.ArrayMusic[i].Name)
			fmt.Println("Musiqa Ijrochisi:", models.ArrayMusic[i].Author)
			fmt.Println("Musiqa turi:", models.ArrayMusic[i].Genre)
			fmt.Println("___________________________________________")
			break
		}else {
			fmt.Println("Bu nomdagi musiqa ma'lumotlari topilmadi", models.ArrayMusic[i].Name)
		}
	} 
	return
}
func GetbyauthorMusic(author string) {
	fmt.Println("__________________________________________")
	for i := 0; i < len(models.ArrayMusic); i++ {
		if models.ArrayMusic[i].Author == author {
			fmt.Println("Musiqa ID:", models.ArrayMusic[i].ID)
			fmt.Println("Musiqa nomi:", models.ArrayMusic[i].Name)
			fmt.Println("Musiqa Ijrochisi:", models.ArrayMusic[i].Author)
			fmt.Println("Musiqa turi:", models.ArrayMusic[i].Genre)
			fmt.Println("___________________________________________")
			break
		}else {
			fmt.Println("Bu ijrochi nomidagi musiqa ma'lumotlari topilmadi", models.ArrayMusic[i].Author)
		}
	} 
	return
}
func GetbygenreMusic(genre string) {
	fmt.Println("__________________________________________")
	for i := 0; i < len(models.ArrayMusic); i++ {
		if models.ArrayMusic[i].Genre == genre {
			fmt.Println("Musiqa ID:", models.ArrayMusic[i].ID)
			fmt.Println("Musiqa nomi:", models.ArrayMusic[i].Name)
			fmt.Println("Musiqa Ijrochisi:", models.ArrayMusic[i].Author)
			fmt.Println("Musiqa turi:", models.ArrayMusic[i].Genre)
			fmt.Println("___________________________________________")
			break
		}else {
			fmt.Println("Bu turdagi musiqa ma'lumotlari topilmadi", models.ArrayMusic[i].Genre)
		}
	} 
	return
}