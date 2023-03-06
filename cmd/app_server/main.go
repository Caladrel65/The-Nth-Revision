package main

import (
	"nth_revision/entities"
)

func main() {
	// TODO: stuff :)

	// setup things

	// etc

	// use tiles for info? what about storage?
	// blergh

	/**

	  what do we need?

	  Also...a map/grid?

	  enemy and character should be able to attack each other

	  let's start with just the character/enemy and their interactions

	  // utils for the interactions? That way I can pass a character, enemy, etc, and as long as it implements the interface we're good? Kinda like polymorphism! I learned about that in college, yeah!

	*/

	// Gracefully exit :D
	// Stop all incoming processes
	// Resolve all incoming processes
	// Return all results
	// Store all states
	// Turn off

	// TESTING
	// Combat test 1

	character, _ := entities.NewCharacter("character", 1, 11, 2)
	character.PrintHealth()
	enemy, _ := entities.NewCharacter("enemy", 1, 5, 2)

	for character.Health() > 0 {
		character.TakeDamage(enemy.Damage())
		character.PrintHealth()
	}
}
