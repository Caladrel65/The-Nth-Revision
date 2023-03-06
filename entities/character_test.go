package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Some default values we will use in testing
const (
	Character_Name       = "Test Name"
	Character_Level      = 1
	Character_Max_Health = 10
	Character_Health     = Character_Max_Health
	Character_Damage     = 2
)

// TestNewCharacter calls the NewCharacter function with some
// default values. If any of these tests break, changes were
// made that affect creating characters.

func TestNewCharacter_Success(t *testing.T) {
	expected := Character{
		name:      Character_Name,
		level:     Character_Level,
		maxHealth: Character_Max_Health,
		health:    Character_Health,
		damage:    Character_Damage,
		gear:      &EquippedGear{},
	}

	actual, err := NewCharacter(Character_Name, Character_Level, Character_Max_Health, Character_Damage)
	assert.NoError(t, err)

	assert.Equal(t, expected, *actual)
}

func TestNewCharacter_BadLevel(t *testing.T) {
	actual, err := NewCharacter(Character_Name, 0, Character_Max_Health, Character_Damage)
	assert.Error(t, err)

	assert.Nil(t, actual)
}

func TestNewCharacter_BadMaxHealth(t *testing.T) {
	actual, err := NewCharacter(Character_Name, Character_Level, 0, Character_Damage)
	assert.Error(t, err)

	assert.Nil(t, actual)
}
