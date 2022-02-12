// package zebra implements routines to solve the zebra-puzzle
package zebra

type Solution struct {
	DrinksWater string
	OwnsZebra   string
}

// define constants to represent information on given facts
const Cigar, Citizenship, Drink, HouseColor, Pet = 0, 1, 2, 3, 4
const Chesterfields, Kools, LuckyStrike, OldGold, Parliaments = 0, 1, 2, 3, 4
const Englishman, Japanese, Norwegian, Spaniard, Ukrainian = 0, 1, 2, 3, 4
const Coffee, Milk, OrangeJuice, Tea, Water = 0, 1, 2, 3, 4
const Blue, Green, Ivory, Red, Yellow = 0, 1, 2, 3, 4
const Dog, Fox, Horse, Snail, Zebra = 0, 1, 2, 3, 4
const Unknown, PropertiesCount, HousesCount = -1, 5, 5
const FirstHouse, MiddleHouse = 0, 2

// database
var db [HousesCount][PropertiesCount]int

// NextTo checks if the house adjacent to given house has the given property value
func NextTo(house, property, value int) bool {
	if house > 0 && (db[house-1][property] == Unknown || db[house-1][property] == value) {
		return true
	}
	if house < HousesCount-1 && (db[house+1][property] == Unknown || db[house+1][property] == value) {
		return true
	}
	return false
}

// Supported checks if the pair of property values are true
func Supported(prop1, value1, prop2, value2 int) bool {
	for i := 0; i < HousesCount; i++ {
		if db[i][prop1] == value1 && db[i][prop2] != value2 && db[i][prop2] != Unknown {
			return false
		}
	}
	return true
}

// SupportedNextTo checks if the given property value and the property value of adjacent house are true
func SupportedNextTo(prop1, value1, prop2, value2 int) bool {
	for i := 0; i < HousesCount; i++ {
		if db[i][prop1] == value1 && !NextTo(i, prop2, value2) {
			return false
		}
	}
	return true
}

// Valid checks if all the given hints are satisfied
func Valid() bool {
	// 1. There are five houses.
	for i := 0; i < HousesCount; i++ {
		for j := 0; j < PropertiesCount; j++ {
			for k := i + 1; k < HousesCount; k++ {
				if db[i][j] != Unknown && db[k][j] != Unknown && db[i][j] == db[k][j] {
					return false
				}
			}
		}
	}
	// 2. The Englishman lives in the red house.
	if !Supported(Citizenship, Englishman, HouseColor, Red) {
		return false
	}
	// 3. The Spaniard owns the dog.
	if !Supported(Citizenship, Spaniard, Pet, Dog) {
		return false
	}
	// 4. Coffee is drunk in the green house.
	if !Supported(Drink, Coffee, HouseColor, Green) {
		return false
	}
	// 5. The Ukrainian drinks tea.
	if !Supported(Citizenship, Ukrainian, Drink, Tea) {
		return false
	}
	// 6. The green house is immediately to the right of the ivory house.
	if db[FirstHouse][HouseColor] == Green {
		return false
	}
	for i := 1; i < HousesCount; i++ {
		if db[i][HouseColor] == Green && db[i-1][HouseColor] != Ivory && db[i-1][HouseColor] != Unknown {
			return false
		}
	}
	// 7. The Old Gold smoker owns Snail.
	if !Supported(Cigar, OldGold, Pet, Snail) {
		return false
	}
	// 8. Kools are smoked in the yellow house.
	if !Supported(Cigar, Kools, HouseColor, Yellow) {
		return false
	}
	// 9. Milk is drunk in the middle house.
	if db[MiddleHouse][Drink] != Milk && db[MiddleHouse][Drink] != Unknown {
		return false
	}
	// 10. The Norwegian lives in the first house.
	if db[FirstHouse][Citizenship] != Norwegian && db[FirstHouse][Citizenship] != Unknown {
		return false
	}
	// 11. The man who smokes Chesterfields lives in the house next to the man with the fox.
	if !SupportedNextTo(Cigar, Chesterfields, Pet, Fox) {
		return false
	}
	// 12. Kools are smoked in the house next to the house where the horse is kept.
	if !SupportedNextTo(Cigar, Kools, Pet, Horse) {
		return false
	}
	// 13. The Lucky Strike smoker drinks orange juice.
	if !Supported(Cigar, LuckyStrike, Drink, OrangeJuice) {
		return false
	}
	// 14. The Japanese smokes Parliaments.
	if !Supported(Citizenship, Japanese, Cigar, Parliaments) {
		return false
	}
	// 15. The Norwegian lives next to the blue house.
	if !SupportedNextTo(Citizenship, Norwegian, HouseColor, Blue) {
		return false
	}
	return true
}

// HasSolution determines if there's a solution to the puzzle
func HasSolution(house, property int) bool {
	oldVal := db[house][property]
	for i := 0; i < 5; i++ {
		db[house][property] = i
		if Valid() {
			newHouse := (house + 1) % HousesCount
			newProperty := property
			if newHouse == 0 {
				newProperty++
			}
			if newProperty >= PropertiesCount {
				return true
			}
			if HasSolution(newHouse, newProperty) {
				return true
			}
		}
	}
	db[house][property] = oldVal
	return false
}

// SolvePuzzle invokes the actual solver and returns the Solution
func SolvePuzzle() Solution {
	Name := map[int]string{
		Englishman: "Englishman",
		Japanese:   "Japanese",
		Norwegian:  "Norwegian",
		Spaniard:   "Spaniard",
		Ukrainian:  "Ukrainian",
	}
	for h := 0; h < HousesCount; h++ {
		for p := 0; p < PropertiesCount; p++ {
			db[h][p] = Unknown
		}
	}
	S := Solution{}
	if HasSolution(0, 0) {
		for h := 0; h < HousesCount; h++ {
			if db[h][Drink] == Water {
				S.DrinksWater = Name[db[h][Citizenship]]
			}
			if db[h][Pet] == Zebra {
				S.OwnsZebra = Name[db[h][Citizenship]]
			}
		}
	}
	return S
}
