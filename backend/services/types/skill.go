package serviceTypes

type Skill struct {
	ID          uint     `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Industry    Industry `json:"industry"`
}
