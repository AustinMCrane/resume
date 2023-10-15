package main

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

type ResumeWork struct {
	Name       string   `json:"name"`
	Position   string   `json:"position"`
	URL        string   `json:"url"`
	StartDate  string   `json:"startDate"`
	EndDate    string   `json:"endDate"`
	Summary    string   `json:"summary"`
	Highlights []string `json:"highlights"`
}

func GetWorks() []ResumeWork {
	return []ResumeWork{
		{
			Name:      "National Institute For Aviation Research (NIAR)",
			Position:  "Software Engineer",
			StartDate: "2013-01-01",
			EndDate:   "2015-01-01",
			Summary:   `Robotics and Automation Lab, Software Engineer Mobile Apps`,
		},
	}
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

func GetResume() Resume {
	return Resume{
		Basics: GetBasics(),
		Work:   GetWorks(),
		Skills: GetSkills(),
	}
}
