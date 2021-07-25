// package allergies implements routines to that determines
// the allergens given a score and whether a score includes
// an allergen
package allergies

var allergenList = []string {
    "eggs", "peanuts", "shellfish", "strawberries",
    "tomatoes", "chocolate", "pollen", "cats" }

var allergenScore = map[string]uint {
    "eggs":1, "peanuts":2, "shellfish":4, "strawberries":8,
    "tomatoes":16, "chocolate":32, "pollen":64, "cats":128 }

// Allergic takes in a score and allergen and 
// checks if the score includes the given allergen
func AllergicTo(score uint, allergen string) bool {
    return allergenScore[allergen] & score != 0
}

// Allergies takes in a score and returns a list
// of allergens indicated by the score
func Allergies(score uint) []string {
    allergens := []string{}
    for _, a := range allergenList {
        if AllergicTo(score, a) {
            allergens = append(allergens, a)
        }
    }
    return allergens
}

