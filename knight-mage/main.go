package main

import "fmt"

type Character interface {
	TakeDamage(amount int)
	Status() string
}

type Warrior struct {
	Name  string
	HP    int
	Armor int
}

type Mage struct {
	Name string
	HP   int
}

func (m *Mage) TakeDamage(amount int) {
	m.HP -= amount
	if m.HP < 0 {
		m.HP = 0
		fmt.Println("ХП не может быть меньше 0")
	}
}

func (m Mage) Status() string {
	return fmt.Sprintf("%s is %d HP", m.Name, m.HP)
}

func (w *Warrior) TakeDamage(amount int) {
	dmg := max(amount-w.Armor, 0)
	w.HP -= dmg
	if w.HP < 0 {
		w.HP = 0
		fmt.Println("ХП не может быть меньше 0")
	}
}

func (w Warrior) Status() string {
	return fmt.Sprintf("%s is %d HP", w.Name, w.HP)
}

func attackAll(chars []Character, damage int) {

	for _, char := range chars {
		char.TakeDamage(damage)
		fmt.Println(char.Status())
	}
}

func main() {
	chars := []Character{
		&Mage{"Mage", 60},
		&Warrior{"Knight", 100, 10},
	}
	attackAll(chars, 30)
	attackAll(chars, 50)
}
