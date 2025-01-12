package schema

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

type timeSpan struct {
	StartDate string `json:"startDate" validate:"datetime,ltfield=EndDate"`
	EndDate   string `json:"endDate"   validate:"datetime,gtfield=StartDate"`
}

type location struct {
	Address     string `json:"address"`
	PostalCode  string `json:"postalCode"`
	City        string `json:"city"`
	CountryCode string `json:"countryCode"`
	Region      string `json:"region"`
}

type profile struct {
	Network  string `json:"network"`
	Username string `json:"username"`
	URL      string `json:"url"      validate:"url"`
}

type basics struct {
	Name     string    `json:"name"     validate:"required"`
	Label    string    `json:"label"`
	Picture  string    `json:"picture"`
	Email    string    `json:"email"    validate:"email"`
	Phone    string    `json:"phone"`
	URL      string    `json:"url"      validate:"url"`
	Summary  string    `json:"summary"`
	Location location  `json:"location"`
	Profiles []profile `json:"profiles"`
}

type work struct {
	timeSpan
	Company    string   `json:"company"`
	Position   string   `json:"position"`
	Website    string   `json:"website"    validate:"url"`
	Summary    string   `json:"summary"`
	Highlights []string `json:"highlights"`
}

type volunteer struct {
	timeSpan
	Organization string   `json:"organization"`
	Position     string   `json:"position"`
	Website      string   `json:"website"      validate:"url"`
	Summary      string   `json:"summary"`
	Highlights   []string `json:"highlights"`
}

type education struct {
	timeSpan
	Institution string   `json:"institution"`
	Area        string   `json:"area"`
	StudyType   string   `json:"studyType"`
	Score       string   `json:"score"`
	Courses     []string `json:"courses"`
}

type award struct {
	Title   string `json:"title"`
	Date    string `json:"date"    validate:"datetime"`
	Awarder string `json:"awarder"`
	Summary string `json:"summary"`
}

type certificate struct {
	Name   string `json:"name"`
	Date   string `json:"date"   validate:"datetime"`
	Issuer string `json:"issuer"`
	URL    string `json:"url"    validate:"url"`
}

type publication struct {
	Name        string `json:"name"`
	Publisher   string `json:"publisher"`
	ReleaseDate string `json:"releaseDate" validate:"datetime"`
	Website     string `json:"website"     validate:"url"`
	Summary     string `json:"summary"`
}

type skill struct {
	Name     string   `json:"name"`
	Level    string   `json:"level"`
	Keywords []string `json:"keywords"`
}

type language struct {
	Language string `json:"language"`
	Fluency  string `json:"fluency"`
}

type interest struct {
	Name     string   `json:"name"`
	Keywords []string `json:"keywords"`
}

type reference struct {
	Name      string `json:"name"`
	Reference string `json:"reference"`
}

type project struct {
	timeSpan
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Highlights  []string `json:"highlights"`
	URL         string   `json:"url"         validate:"url"`
}

// Taken from https://jsonresume.org/schema/
type Resume struct {
	Basics       basics        `json:"basics"`
	Work         []work        `json:"work"`
	Volunteer    []volunteer   `json:"volunteer"`
	Education    []education   `json:"education"`
	Awards       []award       `json:"awards"`
	Certificates []certificate `json:"certificates"`
	Publications []publication `json:"publications"`
	Skills       []skill       `json:"skills"`
	Languages    []language    `json:"languages"`
	Interests    []interest    `json:"interests"`
	References   []reference   `json:"references"`
	Projects     []project     `json:"project"`
}

func (r *Resume) Validate() error {
	return validate.Struct(r)
}
