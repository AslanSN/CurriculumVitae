package handlers

import (
	"github.com/AslanSN/CurriculumVitae/db/constants"
	"github.com/AslanSN/CurriculumVitae/db/models"
	"github.com/AslanSN/CurriculumVitae/views"
	"github.com/labstack/echo/v4"
)

type HomeHandler struct {
}

func (h HomeHandler) RenderIndex(c echo.Context) error {
	return render(c, views.Index())
}

func (h HomeHandler) HandleShowInfo(c echo.Context) error {
	return show(c, constants.SocialMediaList)
}

func (h HomeHandler) HandleShowExperience(c echo.Context) error {
	var experiences []models.ExperienceStruct
	return showInfoDBConstructor(&experiences, c)
}

func (h HomeHandler) HandleShowAboutMe(c echo.Context) error {
	return show(c, constants.AboutMe)
}

func (h HomeHandler) HandleShowSkills(c echo.Context) error {
	return show(c, constants.SkillsList)
}
