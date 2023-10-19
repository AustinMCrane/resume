package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"os"
	"time"

	"github.com/pkg/errors"
)

const (
	TEMPLATE_PATH = "./assets/templates"
)

type ResumeBasics struct {
	Name     string `json:"name"`
	Label    string `json:"label"`
	Image    string `json:"image"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	URL      string `json:"url"`
	Summary  string `json:"summary"`
	Location struct {
		Address     string `json:"address"`
		PostalCode  string `json:"postalCode"`
		City        string `json:"city"`
		CountryCode string `json:"countryCode"`
		Region      string `json:"region"`
	} `json:"location"`
	Profiles []struct {
		Network  string `json:"network"`
		Username string `json:"username"`
		URL      string `json:"url"`
	} `json:"profiles"`
}

type ResumeWork struct {
	Name       string   `json:"name"`
	Position   string   `json:"position"`
	URL        string   `json:"url"`
	StartDate  string   `json:"startDate"`
	EndDate    string   `json:"endDate"`
	Summary    string   `json:"summary"`
	Highlights []string `json:"highlights"`
}

type ResumeEducation struct {
	Institution string   `json:"institution"`
	URL         string   `json:"url"`
	Area        string   `json:"area"`
	StudyType   string   `json:"studyType"`
	StartDate   string   `json:"startDate"`
	EndDate     string   `json:"endDate"`
	Score       string   `json:"score"`
	Courses     []string `json:"courses"`
}

type ResumeSkill struct {
	Name     string   `json:"name"`
	Level    string   `json:"level"`
	Keywords []string `json:"keywords"`
}

const (
	SkillLevelBeginner     = "Beginner"
	SkillLevelIntermediate = "Intermediate"
	SkillLevelMaster       = "Master"
)

type ResumeInterest struct {
	Name     string   `json:"name"`
	Keywords []string `json:"keywords"`
}

type ResumeReference struct {
	Name      string `json:"name"`
	Reference string `json:"reference"`
}

type ResumeProject struct {
	Name        string   `json:"name"`
	StartDate   string   `json:"startDate"`
	EndDate     string   `json:"endDate"`
	Description string   `json:"description"`
	Highlights  []string `json:"highlights"`
	URL         string   `json:"url"`
}

type Resume struct {
	Basics    ResumeBasics `json:"basics"`
	Work      []ResumeWork `json:"work"`
	Volunteer []struct {
		Organization string   `json:"organization"`
		Position     string   `json:"position"`
		URL          string   `json:"url"`
		StartDate    string   `json:"startDate"`
		EndDate      string   `json:"endDate"`
		Summary      string   `json:"summary"`
		Highlights   []string `json:"highlights"`
	} `json:"volunteer"`
	Education []ResumeEducation `json:"education"`
	Awards    []struct {
		Title   string `json:"title"`
		Date    string `json:"date"`
		Awarder string `json:"awarder"`
		Summary string `json:"summary"`
	} `json:"awards"`
	Certificates []struct {
		Name   string `json:"name"`
		Date   string `json:"date"`
		Issuer string `json:"issuer"`
		URL    string `json:"url"`
	} `json:"certificates"`
	Publications []struct {
		Name        string `json:"name"`
		Publisher   string `json:"publisher"`
		ReleaseDate string `json:"releaseDate"`
		URL         string `json:"url"`
		Summary     string `json:"summary"`
	} `json:"publications"`
	Skills    []ResumeSkill `json:"skills"`
	Languages []struct {
		Language string `json:"language"`
		Fluency  string `json:"fluency"`
	} `json:"languages"`
	Interests  []ResumeInterest  `json:"interests"`
	References []ResumeReference `json:"references"`
	Projects   []ResumeProject   `json:"projects"`
}

type HTMLResume struct {
	Resume         Resume
	CalculateYears func(string, string) int
}

func GetBasics() ResumeBasics {
	return ResumeBasics{
		Name:  "Austin Crane",
		Label: "Software Engineer",
		Email: "me@austinmcrane.com",
		URL:   "https://austinmcrane.com",
		Summary: `
TODO
`,
	}
}

func GetWorks() []ResumeWork {
	return []ResumeWork{
		{
			Name:      "National Institute For Aviation Research (NIAR)",
			Position:  "Software Engineer",
			StartDate: "2013-01-01",
			EndDate:   "2015-01-01",
			Summary:   `Robotics and Automation Lab, Software Engineer Mobile Apps.`,
		},
		{
			Name:      "Ennovar",
			Position:  "Software Engineer Lead",
			StartDate: "2015-01-01",
			EndDate:   "2017-01-01",
			Summary:   "Created and maintained websites and apps for clients.",
		},
		{
			Name:      "AccuWeather",
			Position:  "Software Engineer",
			StartDate: "2017-01-01",
			EndDate:   "2021-01-01",
			Summary:   "Worked on an enterprise weather platform for sever weather.",
		},
		{
			Name:      "Ad Hoc",
			Position:  "Software Engineer",
			StartDate: "2021-01-01",
			EndDate:   "Present",
			Summary:   "Work with the Centers for Medicare & Medicaid Services (CMS) to improve the lives of millions of people.",
		},
	}
}

func GetSkills() []ResumeSkill {
	return []ResumeSkill{
		{
			Name:  "Backend Development",
			Level: SkillLevelMaster,
			Keywords: []string{
				"Go",
				"Ruby on Rails",
			},
		},
		{
			Name:  "Data Engineering",
			Level: SkillLevelIntermediate,
			Keywords: []string{
				"PostgreSQL",
				"Redis",
				"AWS Glue",
				"Elasticsearch",
			},
		},
		{
			Name:  "Frontend Development",
			Level: SkillLevelIntermediate,
			Keywords: []string{
				"JavaScript",
				"HTML",
				"CSS",
				"React",
				"Vue",
				"Angular",
			},
		},
		{
			Name:  "DevOps",
			Level: SkillLevelIntermediate,
			Keywords: []string{
				"Terraform",
				"Docker",
				"Ansible",
				"Linux",
			},
		},
	}
}

func GetResume() Resume {
	return Resume{
		Basics: GetBasics(),
		Work:   GetWorks(),
		Skills: GetSkills(),
	}
}

func GetHTMLResume() (string, error) {
	f, err := os.Open(TEMPLATE_PATH + "/index.html")
	if err != nil {
		return "", errors.Wrap(err, "failed to open template")
	}
	defer f.Close()
	templateBytes, err := ioutil.ReadAll(f)
	if err != nil {
		return "", errors.Wrap(err, "failed to read template")
	}

	t, err := template.New("webpage").Parse(string(templateBytes))
	if err != nil {
		return "", errors.Wrap(err, "failed to create template")
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, HTMLResume{
		Resume: GetResume(),
		CalculateYears: func(startDate, endDate string) int {
			start, err := time.Parse("2006-01-02", startDate)
			if err != nil {
				return 0
			}
			end, err := time.Parse("2006-01-02", endDate)
			if err != nil {
				return 0
			}
			return int(end.Sub(start).Hours() / 24 / 365)
		},
	})
	if err != nil {
		return "", errors.Wrap(err, "failed to execute template")
	}

	return buf.String(), nil
}
