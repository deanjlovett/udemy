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

var isVerbose = true
func MyvPrintln(a ...interface{}){
	if isVerbose {
		fmt.Println(a...)
	}
}

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
	MyvPrintln("New Player...")
	p.PrintPlayerState()
	return p
}

func (p *Player)PrintPlayerState(){
	MyvPrintln("   Player Name:",p.name)
	MyvPrintln("    Max Health:",p.maxHealth)
	MyvPrintln("Current Health:",p.health)
	MyvPrintln("    Max Energy:",p.maxEnergy)
	MyvPrintln("Current Energy:",p.energy)
	MyvPrintln()
}

func (p *Player)SetHealth(newHealth uint){
	if newHealth > p.maxHealth {
		p.health = p.maxHealth
	} else{
		p.health = newHealth
	}
	MyvPrintln("Player:",p.name, "Set health to ", p.health)
}
func (p *Player)ModHealth(mod int){
	// fmt.Printf("func (p *Player)ModHealth(mod int:%v){\n",mod)
	// p.PrintPlayerState()

	if mod >= 0 {
		p.health += uint(mod)
		if p.health > p.maxHealth {
			p.health = p.maxHealth
		}
		MyvPrintln("Player:",p.name, "Add", mod, "health -> ", p.health)
	}else{
		if -uint(mod) > p.health {
			p.health = 0
		}else{
			p.health -= -uint(mod)
		}
		MyvPrintln("Player:",p.name, "Dec", -uint(mod), "health -> ", p.health)
	}
} 
func (p *Player)SetEnergy(newEnergy uint){
	if newEnergy > p.maxEnergy {
		p.energy = p.maxEnergy
	} 
	p.energy = newEnergy
	MyvPrintln("Player:",p.name,"Set energy to ", p.energy)
}
func (p *Player)ModEnergy(mod int){
	if mod >= 0 {
		p.energy += uint(mod)
		if p.energy > p.maxEnergy {
			p.energy = p.maxEnergy
		}
		MyvPrintln("Player:",p.name, "Add", mod, "energy -> ", p.energy)
	}else{
		if -uint(mod) > p.energy {
			p.energy = 0
		}else{
			p.energy -= -uint(mod)
		}
		MyvPrintln("Player:",p.name, "Dec", -uint(mod), "energy -> ", p.energy)
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
