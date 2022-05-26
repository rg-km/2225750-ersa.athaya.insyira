package main

import "fmt"

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk membuat interface dan mengimplementasikan interface.
// Buatlah interface Employee yang memiliki method signature GetBonus() int
// Buatlah implementasi interface Employee pada objek Manager, SeniorEngineer, dan JuniorEngineer.
// Pada objek Manager, SeniorEngineer, dan JuniorEngineer memiliki satu atribut yaitu BaseSalary.
// Untuk menghitung bonus terdapat tiga aturan sebagai berikut:
// Bonus untuk Manager adalah 3 * BaseSalary
// Bonus untuk SeniorEngineer adalah 2 * BaseSalary
// Bonus untuk JuniorEngineer adalah 1 * BaseSalary

// TODO: answer here
type Employee interface {
	GetBonus() int
}

// deklarasi objek
type Manager struct {
	BaseSalary int
}

type SeniorEngineer struct {
	BaseSalary int
}

type JuniorEngineer struct {
	BaseSalary int
}

// deklarasi method
func (m Manager) GetBonus() int {
	return 3 * m.BaseSalary
}

func (s SeniorEngineer) GetBonus() int {
	return 2 * s.BaseSalary
}

func (j JuniorEngineer) GetBonus() int {
	return 1 * j.BaseSalary
}

func TotalEmployeeBonus(employees []Employee) (total int) {
	// Hitunglah total bonus yang dikeluarkan dari list of Employee
	// TODO: answer here
	total = 0
	for _, item := range employees {
		total += item.GetBonus()
	}
	return
}

func ShowTotalBonus(employees []Employee) {
	total := TotalEmployeeBonus(employees)
	fmt.Println(total)

}

func main() {
	// Buatlah objek konkret untuk masing-masing objek dan panggil function TotalEmployeeBonus. Print total bonus untuk semua employee.
	// TODO: answer here

	// inisiasi objek
	var m Manager
	// 3*
	m.BaseSalary = 20000000

	var s SeniorEngineer
	// 2*
	s.BaseSalary = 15000000

	var j JuniorEngineer
	// 1*
	j.BaseSalary = 10000000

	// hitung bonus
	TotalEmployeeBonus([]Employee{m, s, j})

	// tampilkan total bonus
	ShowTotalBonus([]Employee{m, s, j})

}
