package controllers

import (
	"github.com/bheemeshkammak/IntelOps_Testing/intelops_testing/pkg/rest/server/models"
	"github.com/bheemeshkammak/IntelOps_Testing/intelops_testing/pkg/rest/server/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type TeamController struct {
	teamService *services.TeamService
}

func NewTeamController() (*TeamController, error) {
	teamService, err := services.NewTeamService()
	if err != nil {
		return nil, err
	}
	return &TeamController{
		teamService: teamService,
	}, nil
}

func (teamController *TeamController) CreateTeam(context *gin.Context) {
	// validate input
	var input models.Team
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger team creation
	if _, err := teamController.teamService.CreateTeam(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Team created successfully"})
}

func (teamController *TeamController) UpdateTeam(context *gin.Context) {
	// validate input
	var input models.Team
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger team update
	if _, err := teamController.teamService.UpdateTeam(id, &input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Team updated successfully"})
}

func (teamController *TeamController) FetchTeam(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger team fetching
	team, err := teamController.teamService.GetTeam(id)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, team)
}

func (teamController *TeamController) DeleteTeam(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger team deletion
	if err := teamController.teamService.DeleteTeam(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Team deleted successfully",
	})
}

func (teamController *TeamController) ListTeams(context *gin.Context) {
	// trigger all teams fetching
	teams, err := teamController.teamService.ListTeams()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, teams)
}

func (*TeamController) PatchTeam(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

func (*TeamController) OptionsTeam(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "OPTIONS",
	})
}

func (*TeamController) HeadTeam(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "HEAD",
	})
}
