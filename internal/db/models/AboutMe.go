package models

import "gorm.io/gorm"

type AboutMeType struct {
	gorm.Model
	Id       string `gorm:"not null;unique_index"`
	Language string `gorm:"not null"`
	AboutMe  string `gorm:"not null"`
}

var AboutMeEnglish = AboutMeType{
	Id: "aboutMeEn",
	Language: "en",
	AboutMe:  `Passionate about programming.  Specialized in React, JavaScript, TypeScript, React Query and Redux.  Training in React Native, Golang, Templ and HTMX. Great rigor in best practices (SOLID). Committed to teamwork and effective communication. My attention to detail guarantees the quality of the projects in which I participate.`,
}
