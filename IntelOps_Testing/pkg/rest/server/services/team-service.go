package services

import (
	"github.com/bheemeshkammak/IntelOps_Testing/intelops_testing/pkg/rest/server/daos"
	"github.com/bheemeshkammak/IntelOps_Testing/intelops_testing/pkg/rest/server/models"
)

type TeamService struct {
	teamDao *daos.TeamDao
}

func NewTeamService() (*TeamService, error) {
	teamDao, err := daos.NewTeamDao()
	if err != nil {
		return nil, err
	}
	return &TeamService{
		teamDao: teamDao,
	}, nil
}

func (teamService *TeamService) CreateTeam(team *models.Team) (*models.Team, error) {
	return teamService.teamDao.CreateTeam(team)
}

func (teamService *TeamService) UpdateTeam(id int64, team *models.Team) (*models.Team, error) {
	return teamService.teamDao.UpdateTeam(id, team)
}

func (teamService *TeamService) DeleteTeam(id int64) error {
	return teamService.teamDao.DeleteTeam(id)
}

func (teamService *TeamService) ListTeams() ([]*models.Team, error) {
	return teamService.teamDao.ListTeams()
}

func (teamService *TeamService) GetTeam(id int64) (*models.Team, error) {
	return teamService.teamDao.GetTeam(id)
}
