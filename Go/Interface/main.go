package main

import (
	"fmt"
)

type ninjaWeapon interface {
	attack()
}

type ninjaStar struct {
	owner string
}

type ninjaSword struct {
	owner string
}

func (ninjaStar) attack() {
	fmt.Println("Throwing ninja star")
}

func (ninjaSword) attack() {
	fmt.Println("Swining ninja star")
}

func attack(weapon ninjaWeapon) {
	weapon.attack()
}

func main() {
	weapons := []ninjaWeapon{ninjaStar{"Wallace"}, ninjaSword{"Wallace"}}

	for _, weapon := range weapons {
		attack(weapon)
	}

}
