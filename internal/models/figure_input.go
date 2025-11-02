package models

type FigureInput struct {
	FirstName      string   `json:"first_name"`
	LastName       string   `json:"last_name"`
	YearOfBirth    uint     `json:"year_of_birth"`
	YearOfDeath    uint     `json:"year_of_death"`
	CountryISOCode string   `json:"country_iso_code"`
	HistoricalEra  string   `json:"historical_era"`
	Areas          []string `json:"areas"`
	Profession     string   `json:"profession"`
	Achievements   []string `json:"achievements"`
}
