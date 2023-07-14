package daos

import (
	"database/sql"
	"errors"
	"github.com/bheemeshkammak/IntelOps_Testing/intelops_testing/pkg/rest/server/daos/clients/sqls"
	"github.com/bheemeshkammak/IntelOps_Testing/intelops_testing/pkg/rest/server/models"
	log "github.com/sirupsen/logrus"
)

type TeamDao struct {
	sqlClient *sqls.SQLiteClient
}

func migrateTeams(r *sqls.SQLiteClient) error {
	query := `
	CREATE TABLE IF NOT EXISTS teams(
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
        
		Width TEXT NOT NULL,
		Bandwidth INTEGER NOT NULL,
		Length TEXT NOT NULL,
        CONSTRAINT id_unique_key UNIQUE (Id)
	)
	`
	_, err1 := r.DB.Exec(query)
	return err1
}

func NewTeamDao() (*TeamDao, error) {
	sqlClient, err := sqls.InitSqliteDB()
	if err != nil {
		return nil, err
	}
	err = migrateTeams(sqlClient)
	if err != nil {
		return nil, err
	}
	return &TeamDao{
		sqlClient,
	}, nil
}

func (teamDao *TeamDao) CreateTeam(m *models.Team) (*models.Team, error) {
	insertQuery := "INSERT INTO teams(Width, Bandwidth, Length)values(?, ?, ?)"
	res, err := teamDao.sqlClient.DB.Exec(insertQuery, m.Width, m.Bandwidth, m.Length)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	m.Id = id

	log.Debugf("team created")
	return m, nil
}

func (teamDao *TeamDao) UpdateTeam(id int64, m *models.Team) (*models.Team, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	team, err := teamDao.GetTeam(id)
	if err != nil {
		return nil, err
	}
	if team == nil {
		return nil, sql.ErrNoRows
	}

	updateQuery := "UPDATE teams SET Width = ?, Bandwidth = ?, Length = ? WHERE Id = ?"
	res, err := teamDao.sqlClient.DB.Exec(updateQuery, m.Width, m.Bandwidth, m.Length, id)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, sqls.ErrUpdateFailed
	}

	log.Debugf("team updated")
	return m, nil
}

func (teamDao *TeamDao) DeleteTeam(id int64) error {
	deleteQuery := "DELETE FROM teams WHERE Id = ?"
	res, err := teamDao.sqlClient.DB.Exec(deleteQuery, id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sqls.ErrDeleteFailed
	}

	log.Debugf("team deleted")
	return nil
}

func (teamDao *TeamDao) ListTeams() ([]*models.Team, error) {
	selectQuery := "SELECT * FROM teams"
	rows, err := teamDao.sqlClient.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	var teams []*models.Team
	for rows.Next() {
		m := models.Team{}
		if err = rows.Scan(&m.Id, &m.Width, &m.Bandwidth, &m.Length); err != nil {
			return nil, err
		}
		teams = append(teams, &m)
	}
	if teams == nil {
		teams = []*models.Team{}
	}

	log.Debugf("team listed")
	return teams, nil
}

func (teamDao *TeamDao) GetTeam(id int64) (*models.Team, error) {
	selectQuery := "SELECT * FROM teams WHERE Id = ?"
	row := teamDao.sqlClient.DB.QueryRow(selectQuery, id)
	m := models.Team{}
	if err := row.Scan(&m.Id, &m.Width, &m.Bandwidth, &m.Length); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sqls.ErrNotExists
		}
		return nil, err
	}

	log.Debugf("team retrieved")
	return &m, nil
}
