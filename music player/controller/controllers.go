package controller

import (
	"fmt"
	"musiclibrary/service"
)

func Display() {

	fmt.Println("___________Music Player sistem___________")
sakra:
	fmt.Println("")
	fmt.Println("Menyuni tanlang")
	fmt.Println("1.Foydalanuvchi bo'limiga o'tish")
	fmt.Println("2.Musiqa bo'limiga o'tish")
	fmt.Println("3.Chiqish")
	fmt.Println("")
	fmt.Println("Raqamlar orqali amalyot turini tanlang va raqam kiriting")
	fmt.Println("")
	var num int
	fmt.Scan(&num)
	switch num {
	case 1:
	sakra1:
		fmt.Println("___________Foydalanuvchi___________")
		fmt.Println("")
		fmt.Println("1.Foydalanuvchi yaratish")
		fmt.Println("2.Foydalanuvchi ma'lumotlarini o'zgartirish")
		fmt.Println("3.Foydalanuvchi malumotlarini o'chirish")
		fmt.Println("4.Foydalanuvchi malumotlarini ko'rish")
		fmt.Println("5.Foydalanuvchiga musiqa olish")
		fmt.Println("6.Foydalanuvchidagi musiqani o'chirish")
		fmt.Println("7.Foydalanuvchini ID boyicha qidirish")
		fmt.Println("")
		fmt.Println("____________________________________________")
		fmt.Println("Raqamlar orqali amalyot turini tanlang va raqam kiriting")
		fmt.Println("")
		var num1 int
		fmt.Scan(&num1)
		switch num1 {
		case 1:
			fmt.Println("1.Foydalanuvchi ismini kiriting")
			fmt.Println("2.Foydalanuvchi familyasini kiriting")
			var a, b string
			fmt.Scan(&a, &b)
			service.CreateUser(a, b)
			fmt.Println("Foydalanuvchi muvofaqiyatli yartildi!")
			fmt.Println("1.Bosh menuga qaytish uchun 1 ni bosing")
			fmt.Println("2.Orqaga qaytish uchun 2 ni bosing")
			var c int
			fmt.Scan(&c)
			switch c {
			case 1:
				goto sakra
			case 2:
				goto sakra1
			}

		case 2:
			fmt.Println("1.Foydalanuvchi ID kiriting")
			fmt.Println("2.Foydalanuvchi ismini kiriting")
			fmt.Println("3.Foydalanuvchi familyasini kiriting")
			var a, b string
			var d int
			fmt.Scan(&d, &a, &b)
			service.UpdateUser(d, a, b)
			fmt.Println("Foydalanuvchi muvofaqiyatli o'zgartirildi!")
			fmt.Println("1.Bosh menuga qaytish uchun 1 ni bosing")
			fmt.Println("2.Orqaga qaytish uchun 2 ni bosing")
			var c int
			fmt.Scan(&c)
			switch c {
			case 1:
				goto sakra
			case 2:
				goto sakra1
			}
		case 3:
			fmt.Println("ID kiriting")
			var l int
			fmt.Scan(&l)
			service.DeleteUser(l)
			fmt.Println("Foydalanuvchi muvofaqiyatli o'chirildi!")
			fmt.Println("1.Bosh menuga qaytish uchun 1 ni bosing")
			fmt.Println("2.Orqaga qaytish uchun 2 ni bosing")
			var c int
			fmt.Scan(&c)
			switch c {
			case 1:
				goto sakra
			case 2:
				goto sakra1
			}
		case 4:
			service.ReadUser()

			fmt.Println("Amalyot muvofaqiyatli bajarildi!")
			fmt.Println("1.Bosh menuga qaytish uchun 1 ni bosing")
			fmt.Println("2.Orqaga qaytish uchun 2 ni bosing")
			var c int
			fmt.Scan(&c)
			switch c {
			case 1:
				goto sakra
			case 2:
				goto sakra1
			}
		case 5:
			service.TakeMusic()
			fmt.Println("1.Bosh menuga qaytish uchun 1 ni bosing")
			fmt.Println("2.Orqaga qaytish uchun 2 ni bosing")
			var c int
			fmt.Scan(&c)
			switch c {
			case 1:
				goto sakra
			case 2:
				goto sakra1
			}

		case 6:
			service.DeleteUserMusic()
			fmt.Println("1.Bosh menuga qaytish uchun 1 ni bosing")
			fmt.Println("2.Orqaga qaytish uchun 2 ni bosing")
			var c int
			fmt.Scan(&c)
			switch c {
			case 1:
				goto sakra
			case 2:
				goto sakra1
			}
		case 7:
			fmt.Println("Qidirayotgan foydalanuvchi ID sini kiriting")
			var id int
			fmt.Scan(&id)
			service.GetbyidUser(id)
			fmt.Println("1.Bosh menuga qaytish uchun 1 ni bosing")
			fmt.Println("2.Orqaga qaytish uchun 2 ni bosing")
			var c int
			fmt.Scan(&c)
			switch c {
			case 1:
				goto sakra
			case 2:
				goto sakra1
			}

		default:
			fmt.Println("1.Bosh menuga qaytish uchun 1 ni bosing")
			fmt.Println("2.Orqaga qaytish uchun 2 ni bosing")
			var c int
			fmt.Scan(&c)
			switch c {
			case 1:
				goto sakra
			case 2:
				goto sakra1
			}
		}
	case 2:
	sakra2:
		fmt.Println("___________Music___________")
		fmt.Println("")
		fmt.Println("1.Musiqa yaratish")
		fmt.Println("2.Musiqa nomini o'zgartirish")
		fmt.Println("3.Musiqa malumotlarini o'chirish")
		fmt.Println("4.Musiqa malumotlarini ko'rish")
		fmt.Println("5.Musiqa ID boyicha ko'rish")
		fmt.Println("6.Musiqa nomi boyicha ko'rish")
		fmt.Println("7.Musiqa ijrochisi boyicha ko'rish")
		fmt.Println("8.Musiqa turi boyicha ko'rish")
		fmt.Println("")
		fmt.Println("____________________________________________")
		fmt.Println("Raqamlar orqali amalyot turini tanlang va raqam kiriting")
		fmt.Println("")

		var qw int
		fmt.Scan(&qw)
		switch qw {
		case 1:
			fmt.Println("1.Musiqa nomini kiriting")
			fmt.Println("2.Musiqa ijrochi nomini kiriting")
			fmt.Println("3.Musiqa turini kiriting")
			var a, b, d string

			fmt.Scan(&a, &b, &d)
			service.CreateMusic(a, b, d)
			fmt.Println("Musiqa muvofaqiyatli yartildi!")
			fmt.Println("1.Bosh menuga qaytish uchun 1 ni bosing")
			fmt.Println("2.Orqaga qaytish uchun 2 ni bosing")
			var c int
			fmt.Scan(&c)
			switch c {
			case 1:
				goto sakra
			case 2:
				goto sakra2
			}
		case 2:
			fmt.Println("1.Musiqa ID kiriting")
			fmt.Println("2.Musiqa nomini kiriting")
			fmt.Println("3.Musiqa ijrochi nomini kiriting")
			fmt.Println("4.Musiqa turini kiriting")
			var a, b, k string
			var d int
			fmt.Scan(&d, &a, &b, &k)
			service.UpdateMusic(d, a, b, k)
			fmt.Println("Musiqa muvofaqiyatli o'zgartirildi!")
			fmt.Println("1.Bosh menuga qaytish uchun 1 ni bosing")
			fmt.Println("2.Orqaga qaytish uchun 2 ni bosing")
			var c int
			fmt.Scan(&c)
			switch c {
			case 1:
				goto sakra
			case 2:
				goto sakra2
			}
		case 3:
			fmt.Println("Musiqa ID kiriting")
			var l int
			fmt.Scan(&l)
			service.DeleteMusic(l)
			fmt.Println("Musiqa muvofaqiyatli o'chirildi!")
			fmt.Println("1.Bosh menuga qaytish uchun 1 ni bosing")
			fmt.Println("2.Orqaga qaytish uchun 2 ni bosing")
			var c int
			fmt.Scan(&c)
			switch c {
			case 1:
				goto sakra
			case 2:
				goto sakra2
			}
		case 4:
			service.ReadMusic()
			fmt.Println("Amalyot muvofaqiyatli bajarildi!")
			fmt.Println("1.Bosh menuga qaytish uchun 1 ni bosing")
			fmt.Println("2.Orqaga qaytish uchun 2 ni bosing")
			var c int
			fmt.Scan(&c)
			switch c {
			case 1:
				goto sakra
			case 2:
				goto sakra2
			}
		case 5:
			fmt.Println("Qidirilayotgan musiqa ID sini kiriting")
			var id int
			fmt.Scan(&id)
			service.GetbyidMusic(id)
			fmt.Println("Amalyot muvofaqiyatli bajarildi!")
			fmt.Println("1.Bosh menuga qaytish uchun 1 ni bosing")
			fmt.Println("2.Orqaga qaytish uchun 2 ni bosing")
			var c int
			fmt.Scan(&c)
			switch c {
			case 1:
				goto sakra
			case 2:
				goto sakra2
			}

		case 6:
			fmt.Println("Qidirilayotgan musiqa nomi kiriting")
			var namer string
			fmt.Scan(&namer)
			service.GetbynameMusic(namer)
			fmt.Println("Amalyot muvofaqiyatli bajarildi!")
			fmt.Println("1.Bosh menuga qaytish uchun 1 ni bosing")
			fmt.Println("2.Orqaga qaytish uchun 2 ni bosing")
			var c int
			fmt.Scan(&c)
			switch c {
			case 1:
				goto sakra
			case 2:
				goto sakra2
			}
		case 7:
			fmt.Println("Qidirilayotgan musiqa ijrochi nomini kiriting")
			var author string
			fmt.Scan(&author)
			service.GetbyauthorMusic(author)
			fmt.Println("Amalyot muvofaqiyatli bajarildi!")
			fmt.Println("1.Bosh menuga qaytish uchun 1 ni bosing")
			fmt.Println("2.Orqaga qaytish uchun 2 ni bosing")
			var c int
			fmt.Scan(&c)
			switch c {
			case 1:
				goto sakra
			case 2:
				goto sakra2
			}
		case 8:
			fmt.Println("Qidirilayotgan musiqa turini kiriting")
			var genre string
			fmt.Scan(&genre)
			service.GetbygenreMusic(genre)
			fmt.Println("Amalyot muvofaqiyatli bajarildi!")
			fmt.Println("1.Bosh menuga qaytish uchun 1 ni bosing")
			fmt.Println("2.Orqaga qaytish uchun 2 ni bosing")
			var c int
			fmt.Scan(&c)
			switch c {
			case 1:
				goto sakra
			case 2:
				goto sakra2
			}
		default:
			fmt.Println("1.Bosh menuga qaytish uchun 1 ni bosing")
			fmt.Println("2.Orqaga qaytish uchun 2 ni bosing")
			var c int
			fmt.Scan(&c)
			switch c {
			case 1:
				goto sakra
			case 2:
				goto sakra2
			}

		}

	case 3:
		return
	default:
		fmt.Println("Bunday amalyot mavjud emas!")
		return
	}

}
