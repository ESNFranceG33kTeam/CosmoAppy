package cocas

import (
	"fmt"
	"net/http"

	"gopkg.in/cas.v2"
)

// swagger:model Profile
type Profile struct {
	Username      string   `json:"username"`
	First         string   `json:"firstname"`
	Last          string   `json:"lastname"`
	Email         string   `json:"email"`
	Country       string   `json:"country"`
	Nationality   string   `json:"nationality"`
	PictureURL    string   `json:"picture_url"`
	Birthdate     string   `json:"dateofbirth"`
	Gender        string   `json:"gender"`
	Phone         string   `json:"phone"`
	Sc            string   `json:"section_code"`
	Section       string   `json:"association"`
	Roles         []string `json:"roles"`
	ExtendedRoles []string `json:"extended_roles"`
	Admin         bool     `json:"admin"`
}

var ESNer Profile

func GetProfile(r *http.Request) *Profile {
	ESNer = Profile{
		Username:      cas.Username(r),
		First:         cas.Attributes(r)["first"][0],
		Last:          cas.Attributes(r)["last"][0],
		Email:         cas.Attributes(r)["mail"][0],
		Country:       cas.Attributes(r)["country"][0],
		Nationality:   cas.Attributes(r)["nationality"][0],
		PictureURL:    cas.Attributes(r)["picture"][0],
		Birthdate:     cas.Attributes(r)["birthdate"][0],
		Gender:        cas.Attributes(r)["gender"][0],
		Phone:         cas.Attributes(r)["telephone"][0],
		Sc:            cas.Attributes(r)["sc"][0],
		Section:       cas.Attributes(r)["section"][0],
		Roles:         cas.Attributes(r)["roles"],
		ExtendedRoles: cas.Attributes(r)["extended_roles"],
		Admin:         false,
	}

	// Rules for Super Admin
	for _, role := range ESNer.ExtendedRoles {
		if role == "IT.it_support" || role == fmt.Sprintf("National.webmaster:%s", GalaxyUsers.Country) {
			ESNer.Admin = true
		}
	}
	for _, geek := range GalaxyUsers.G33kTeam {
		if ESNer.Username == geek {
			ESNer.Admin = true
		}
	}
	for _, extra := range GalaxyUsers.ExtraUsername {
		if ESNer.Username == extra {
			ESNer.Admin = true
		}
	}

	return &ESNer
}
