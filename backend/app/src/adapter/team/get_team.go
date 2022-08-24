package team

import (
	"app/entity/model"
	"app/infra/database"
	"app/usecases"
	"database/sql"
)

type GetTeamDependencies struct {
	sqlAdapter *database.SqlAdapter
}

func (dep GetTeamDependencies) Do(teamId *int) (*model.Team, error) {

	stmt := `
	select
		ts.team_id
		ts.team_name
		ts.team_icon
		ts.team_memo
	from
		teams as ts
	on
		ts.team_id = jut.team_id
	inner join
		roles as rs
	on
		rs.role_id = us.role_id
	where
		ts.team_id = ? and

		teams.deleted_at is null
	`
	defer dep.sqlAdapter.DB.Close()

	rows, err := dep.sqlAdapter.DB.Query(stmt, teamId)
	if err != nil {
		return nil, err
	}

	rows.Next()
	d := struct {
		TeamID   string
		TeamName string
		TeamIcon sql.NullString
	}{}

	rows.Scan(
		&d.TeamID,
		&d.TeamName,
		&d.TeamIcon,
	)

	blankToNil := func(column sql.NullString) *string {
		if column.Valid {
			return &column.String
		}
		return nil
	}

	users, err := NewGetTeamUsersAdapter(dep.sqlAdapter).Do(teamId)
	if err != nil {
		return nil, err
	}

	team := &model.Team{
		TeamID:   d.TeamID,
		TeamName: d.TeamName,
		TeamIcon: blankToNil(d.TeamIcon),
		Users:    users,
	}

	return team, nil
}

func NewGetTeamAdapter(sqlAdapter *database.SqlAdapter) usecases.GetTeamAdapter {
	return &GetTeamDependencies{sqlAdapter}
}
