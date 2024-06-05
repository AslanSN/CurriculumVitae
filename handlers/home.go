package handlers

import (
	"context"
	"log"
	"os"
	"path"

	"github.com/AslanSN/CV/db/constants"
	"github.com/AslanSN/CV/views"
	"github.com/labstack/echo/v4"
)

type HomeHandler struct {
}

func (h HomeHandler) CreateHTMLFiles(){
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
	return show(c, constants.Workplaces)
}

func (h HomeHandler) HandleShowAboutMe(c echo.Context) error {
	return show(c, constants.AboutMe)
}

func (h HomeHandler) HandleShowSkills(c echo.Context) error {
	return show(c, constants.SkillsList)
}
