package tables

import (
	"clothing-pair-project/internal/models"

	"strconv"

	"github.com/olekukonko/tablewriter"
)

type UserTable struct {
	table *tablewriter.Table
}

func UsersTablePresenter(table *tablewriter.Table) *UserTable {
	table.SetHeader([]string{"ID", "Username", "Email", "Role", "Created At", "Active"})
	table.SetRowLine(true)
	return &UserTable{table: table}
}

func AddUserTablePresenter(table *tablewriter.Table) *UserTable {
	table.SetHeader([]string{"Username", "Email", "Password", "Role"})
	table.SetRowLine(true)
	return &UserTable{table: table}
}

func (t *UserTable) DisplayUsers(users []models.User) {
	t.table.ClearRows()
	for _, user := range users {
		t.table.Append([]string{
			strconv.Itoa(user.UserID),
			user.Username,
			user.Email,
			user.Role,
			user.CreatedAt.Format("2006-01-02 15:04:05"),
			strconv.FormatBool(user.Active),
		})
	}
	t.table.Render()
}

func (t *UserTable) DisplayAddUser(user models.User) {
	t.table.ClearRows()
	t.table.Append([]string{
		user.Username,
		user.Email,
		user.Password,
		user.Role,
	})
	t.table.Render()
}
