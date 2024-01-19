package service

import (
	"fmt"
	"musiclibrary/models"
)

func CreateUser(firstname string, lastname string) {

	var users models.User = models.User{
		ID:          len(models.ArrayUser) + 1,
		Firstname:   firstname,
		Lastname:    lastname,
		SavedMusics: []models.Music{},
	}

	models.ArrayUser = append(models.ArrayUser, users)

}
func ReadUser() {
	fmt.Println("__________________________________________")
	for i := 0; i < len(models.ArrayUser); i++ {
		fmt.Println("Foydalanuvchi ID:", models.ArrayUser[i].ID)
		fmt.Println("Foydalanuvchi ismi:", models.ArrayUser[i].Firstname)
		fmt.Println("Foydalanuvchi familyasi:", models.ArrayUser[i].Lastname)
		fmt.Println("Foydalanuvchi musiqalari:", models.ArrayUser[i].SavedMusics)
		fmt.Println("___________________________________________")
	}
	//for j:=0; j<len(models.ArrayUser[j].SavedMusics); j++{}
}
func UpdateUser(id int, firstname string, lastname string) {
	for i := 0; i < len(models.ArrayUser); i++ {
		if models.ArrayUser[i].ID == id {
			models.ArrayUser[i].Firstname = firstname
			models.ArrayUser[i].Lastname = lastname
			break
		} else {
			fmt.Println("Bu ID: da foydalanuvchi ma'lumotlari topilmadi", models.ArrayUser[i].ID)
		}
	}
	return
}
func DeleteUser(id int) {
	for i := 0; i < len(models.ArrayUser); i++ {
		if models.ArrayUser[i].ID == id {
			models.ArrayUser = append(models.ArrayUser[:i], models.ArrayUser[i+1:]...)
			break
		} else {
			fmt.Println("Bu ID: da foydalanuvchi ma'lumotlari topilmadi", models.ArrayUser[i].ID)
		}
	}
}
func GetbyidUser(id int) {
	fmt.Println("__________________________________________")
	for i := 0; i < len(models.ArrayUser); i++ {
		if models.ArrayUser[i].ID == id {
			fmt.Println("Foydalanuvchi ID:", models.ArrayUser[i].ID)
			fmt.Println("Foydalanuvchi ismi:", models.ArrayUser[i].Firstname)
			fmt.Println("Foydalanuvchi familyasi:", models.ArrayUser[i].Lastname)
			fmt.Println("Foydalanuvchi musiqalari:", models.ArrayUser[i].SavedMusics)
			fmt.Println("___________________________________________")
			fmt.Println("Amalyot muvofaqiyatli bajarildi!")
			break
		} else {
			fmt.Println("Bu ID: da foydalanuvchi ma'lumotlari topilmadi", models.ArrayUser[i].ID)
		}
	}
	return
}
func TakeMusic() {
	var num1, num2 int
sakra3:
	fmt.Println("Qaysi foydaluvchiga musiqa olmoqchisiz ID sini kiriting")
	fmt.Scan(&num1)
sakra4:
	fmt.Println("Olmoqchi bo'lgan musiqani ID sini kiriting")
	fmt.Scan(&num2)
	for i := 0; i < len(models.ArrayUser); i++ {
		if models.ArrayUser[i].ID == num1 {
			for j := 0; j < len(models.ArrayMusic); j++ {
				if models.ArrayMusic[j].ID == num2 {
					models.ArrayUser[i].SavedMusics = append(models.ArrayUser[i].SavedMusics, models.ArrayMusic[j])
					fmt.Println("Amaliyot muvaffaqiyatli bajarildi!")
					return
				} else {
					fmt.Println("Bu ID dagi musiqa topilmadi qaytadan urinib koring!")
					goto sakra4
				}

			}
		} else {
			fmt.Println("Bu ID dagi foydalanuvchi topilmadi qaytadan urinib koring")
			goto sakra3
		}

	}

}
func DeleteUserMusic() {
	var num1, num2 int
sakra3:
	fmt.Println("Qaysi foydaluvchidagi musiqa o'chirmoqchisiz ID sini kiriting")
	fmt.Scan(&num1)
sakra4:
	fmt.Println("O'chirmoqchi bo'lgan musiqani ID sini kiriting")
	fmt.Scan(&num2)
	for i := 0; i < len(models.ArrayUser); i++ {
		if models.ArrayUser[i].ID == num1 {
			for j := 0; j < len(models.ArrayUser[i].SavedMusics); j++ {
				if models.ArrayMusic[j].ID == num2 {
					models.ArrayUser[i].SavedMusics = append(models.ArrayUser[i].SavedMusics[:i], models.ArrayUser[i].SavedMusics[j+1:]...)
					fmt.Println("Amaliyot muvaffaqiyatli bajarildi!")
					return
				} else {
					fmt.Println("Bu ID dagi musiqa topilmadi qaytadan urinib koring!")
					goto sakra4
				}

			}
		} else {
			fmt.Println("Bu ID dagi foydalanuvchi topilmadi qaytadan urinib koring")
			goto sakra3
		}

	}

}
