package memory

import (
	"github.com/AustinMCrane/resume/pkg/model"
)

// ResumeStore is a simple in-memory store for Resume objects.
type ResumeStore struct{}

func (mem *ResumeStore) GetResume() (*model.Resume, error) {
	return getResume(), nil
}

func getBasics() model.ResumeBasics {
	return model.ResumeBasics{
		Name:  "Austin Crane",
		Label: "Software Engineer",
		Email: "me@austinmcrane.com",
		URL:   "https://austinmcrane.com",
		Summary: `
TODO
`,
	}
}

func getWorks() []model.ResumeWork {
	return []model.ResumeWork{
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

func getSkills() []model.ResumeSkill {
	return []model.ResumeSkill{
		{
			Name:  "Backend Development",
			Level: model.SkillLevelMaster,
			Keywords: []string{
				"Go",
				"Ruby on Rails",
			},
		},
		{
			Name:  "Data Engineering",
			Level: model.SkillLevelIntermediate,
			Keywords: []string{
				"PostgreSQL",
				"Redis",
				"AWS Glue",
				"Elasticsearch",
			},
		},
		{
			Name:  "Frontend Development",
			Level: model.SkillLevelIntermediate,
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
			Level: model.SkillLevelIntermediate,
			Keywords: []string{
				"Terraform",
				"Docker",
				"Ansible",
				"Linux",
			},
		},
	}
}

func GetResumeStore() model.ResumeStore {
	return &ResumeStore{}
}

func getResume() *model.Resume {
	return &model.Resume{
		Basics: getBasics(),
		Work:   getWorks(),
		Skills: getSkills(),
	}
}
