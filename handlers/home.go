package handlers

import (
	"context"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/AslanSN/CurriculumVitae/db"
	"github.com/AslanSN/CurriculumVitae/db/constants"
	"github.com/AslanSN/CurriculumVitae/db/models"
	"github.com/AslanSN/CurriculumVitae/views"
	"github.com/labstack/echo/v4"
)

type HomeHandler struct {
}

func (h HomeHandler) CreateHTMLFiles() {
	rootPath := "public"
	name := path.Join(rootPath, "index.html")
	f, err := os.Create(name)
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}

	err = views.Index().Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write index page: %v", err)
	}
}

func (h HomeHandler) RenderIndex(c echo.Context) error {
	return render(c, views.Index())
}

func (h HomeHandler) HandleShowInfo(c echo.Context) error {
	return show(c, constants.SocialMediaList)
}

func (h HomeHandler) HandleShowExperience(c echo.Context) error {
	var experiences []models.ExperienceStruct
	result := db.DB.Find(&experiences)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": result.Error.Error(),
		})
	}
	return c.JSON(http.StatusOK, experiences)
}

func (h HomeHandler) HandleShowAboutMe(c echo.Context) error {
	return show(c, constants.AboutMe)
}

func (h HomeHandler) HandleShowSkills(c echo.Context) error {
	return show(c, constants.SkillsList)
}
