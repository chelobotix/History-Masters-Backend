package models

// Basic Response
type YearResponse struct {
	ID   uint `json:"id"`
	Year int  `json:"year"`
	BC   bool `json:"bc"`
}

type CountryResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	IsoCode string `json:"iso_code"`
}

type HistoricalEraResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type AreaResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type ProfessionResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type AchievementResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type FigureResponse struct {
	ID            uint                  `json:"id"`
	FirstName     string                `json:"first_name"`
	LastName      string                `json:"last_name"`
	YearOfBirth   YearResponse          `json:"year_of_birth"`
	YearOfDeath   YearResponse          `json:"year_of_death"`
	Country       CountryResponse       `json:"country"`
	HistoricalEra HistoricalEraResponse `json:"historical_era"`
	Areas         []AreaResponse        `json:"areas"`
	Profession    ProfessionResponse    `json:"profession"`
	Achievements  []AchievementResponse `json:"achievements"`
}
