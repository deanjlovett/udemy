//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Player struct {
	name      string
	health    uint
	maxHealth uint
	energy    uint
	maxEnergy uint
}
// func NewPlayer(name string)Player{
// 	player := Player{"",0,100,0,100}
// 	player.name = name
// 	return player
// }
func NewPlayer(name string,maxHealth, maxEnergy uint)Player{
	p := Player{name,0,maxHealth,0,maxEnergy}
	fmt.Println("New Player...")
	p.PrintPlayerState()
	return p
}

func (p *Player)PrintPlayerState(){
	fmt.Println("   Player Name:",p.name)
	fmt.Println("    Max Health:",p.maxHealth)
	fmt.Println("Current Health:",p.health)
	fmt.Println("    Max Energy:",p.maxEnergy)
	fmt.Println("Current Energy:",p.energy)
	fmt.Println()
}

func (p *Player)SetHealth(newHealth uint){
	if newHealth > p.maxHealth {
		p.health = p.maxHealth
	} else{
		p.health = newHealth
	}
	fmt.Println("Player:",p.name, "Set health to ", p.health)
}
func (p *Player)ModHealth(mod int){
	// fmt.Printf("func (p *Player)ModHealth(mod int:%v){\n",mod)
	// p.PrintPlayerState()

	if mod >= 0 {
		p.health += uint(mod)
		if p.health > p.maxHealth {
			p.health += p.maxHealth
		}
		fmt.Println("Player:",p.name, "Add", mod, "health -> ", p.health)
	}else{
		if -uint(mod) > p.health {
			p.health = 0
		}else{
			p.health -= -uint(mod)
		}
		fmt.Println("Player:",p.name, "Dec", -uint(mod), "health -> ", p.health)
	}
} 
func (p *Player)SetEnergy(newEnergy uint){
	if newEnergy > p.maxEnergy {
		p.energy = p.maxEnergy
	} 
	p.energy = newEnergy
	fmt.Println("Player:",p.name,"Set energy to ", p.energy)
}
func (p *Player)ModEnergy(mod int){
	if mod >= 0 {
		p.energy += uint(mod)
		if p.energy > p.maxEnergy {
			p.energy += p.maxEnergy
		}
		fmt.Println("Player:",p.name, "Add", mod, "energy -> ", p.energy)
	}else{
		if -uint(mod) > p.energy {
			p.energy = 0
		}else{
			p.energy -= -uint(mod)
		}
		fmt.Println("Player:",p.name, "Dec", -uint(mod), "energy -> ", p.energy)
	}
} 

func main() {
	aPlayer := NewPlayer("fred",100,1000)
	aPlayer.SetHealth(90)
	aPlayer.SetEnergy(900)
	aPlayer.PrintPlayerState()

	aPlayer.ModHealth(10)
	aPlayer.ModEnergy(100)
	aPlayer.PrintPlayerState()

	aPlayer.ModHealth(-1)
	aPlayer.ModEnergy(-1)
	aPlayer.PrintPlayerState()
}
