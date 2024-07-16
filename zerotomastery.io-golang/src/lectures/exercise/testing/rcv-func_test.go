//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests

package main
import "testing"
// import "fmt"

func TestAbove(t *testing.T){

	tp := NewPlayer("test_player",100,1000)
	if tp.health != 0 {
		t.Errorf("incorrect health: %v, want %v", tp.health, 0)
	}else{
		tp.SetHealth(tp.maxHealth)
		if tp.health != tp.maxHealth {
			t.Errorf("incorrect health: %v, want %v", tp.health, tp.maxHealth)
		}
		tp.ModHealth(int(tp.maxHealth) + 1)
		if tp.health != tp.maxHealth {
			t.Errorf("incorrect health: %v, want %v", tp.health, tp.maxHealth)
		}
	}
	if tp.energy != 0 {
		t.Errorf("incorrect energy: %v, want %v", tp.energy, 0)
	}else{
		tp.SetEnergy(tp.maxEnergy)
		if tp.energy != tp.maxEnergy {
			t.Errorf("incorrect energy: %v, want %v", tp.energy, tp.maxEnergy)
		}
		tp.ModEnergy(int(tp.maxEnergy) + 1)
		if tp.energy != tp.maxEnergy {
			t.Errorf("incorrect energy: %v, want %v", tp.energy, tp.maxEnergy)
		}
	}
}

func TestBelow(t *testing.T){

	tp := NewPlayer("test_player",100,1000)
	if tp.health != 0 {
		t.Errorf("incorrect health: %v, want %v", tp.health, 0)
	}else{
		tp.SetHealth(0)
		if tp.health != 0 {
			t.Errorf("incorrect health: %v, want %v", tp.health, 0)
		}
		tp.ModHealth(-1)
		if tp.health != 0 {
			t.Errorf("incorrect health: %v, want %v", tp.health, 0)
		}
	}
	if tp.energy != 0 {
		t.Errorf("incorrect energy: %v, want %v", tp.energy, 0)
	}else{
		tp.SetEnergy(0)
		if tp.energy != 0 {
			t.Errorf("incorrect energy: %v, want %v", tp.energy, 0)
		}
		tp.ModEnergy(-1)
		if tp.energy != 0 {
			t.Errorf("incorrect energy: %v, want %v", tp.energy, 0)
		}
	}
}

func TestEnergy(t *testing.T){
	// test min
	// fmt.Println("testing for excedding min energy")

	tp := NewPlayer("test_player",100,1000)
	if tp.energy != 0 {
		t.Errorf("incorrect energy: %v, want %v", tp.energy, 0)
	}else{
		tp.SetEnergy(0)
		if tp.energy != 0 {
			t.Errorf("incorrect energy: %v, want %v", tp.energy, 0)
		}
		tp.ModEnergy(-1)
		if tp.energy != 0 {
			t.Errorf("incorrect energy: %v, want %v", tp.energy, 0)
		}
	}

	// test max
	// fmt.Println("testing for excedding max energy")
	tp = NewPlayer("test_player",100,1000)

	if tp.energy != 0 {
		t.Errorf("incorrect energy: %v, want %v", tp.energy, 0)
	}else{
		tp.SetEnergy(tp.maxEnergy)
		if tp.energy != tp.maxEnergy {
			t.Errorf("incorrect energy: %v, want %v", tp.energy, tp.maxEnergy)
		}
		tp.ModEnergy(int(tp.maxEnergy) + 1)
		if tp.energy != tp.maxEnergy {
			t.Errorf("incorrect energy: %v, want %v", tp.energy, tp.maxEnergy)
		}
	}

}
func TestHealth(t *testing.T){
	// test min
	// fmt.Println("test min health")

	tp := NewPlayer("test_player",100,1000)
	if tp.health != 0 {
		t.Errorf("incorrect health: %v, want %v", tp.health, 0)
	}else{
		tp.SetHealth(0)
		if tp.health != 0 {
			t.Errorf("incorrect health: %v, want %v", tp.health, 0)
		}
		tp.ModHealth(-1)
		if tp.health != 0 {
			t.Errorf("incorrect health: %v, want %v", tp.health, 0)
		}
	}

	// test max
	// fmt.Println("test max health")

	tp = NewPlayer("test_player",100,1000)
	if tp.health != 0 {
		t.Errorf("incorrect health: %v, want %v", tp.health, 0)
	}else{
		tp.SetHealth(tp.maxHealth)
		if tp.health != tp.maxHealth {
			t.Errorf("incorrect health: %v, want %v", tp.health, tp.maxHealth)
		}
		tp.ModHealth(int(tp.maxHealth) + 1)
		if tp.health != tp.maxHealth {
			t.Errorf("incorrect health: %v, want %v", tp.health, tp.maxHealth)
		}
	}
	
}

