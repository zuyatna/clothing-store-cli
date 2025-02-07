package tables

import (
	"clothing-pair-project/internal/models"

	"strconv"

	"github.com/olekukonko/tablewriter"
)

type UserTable struct {
	writer *tablewriter.Table
}

func UsersTablePresenter(writer *tablewriter.Table) *UserTable {
	writer.SetHeader([]string{"ID", "Username", "Email", "Role", "Created At", "Active"})
	writer.SetRowLine(true)
	return &UserTable{writer: writer}
}

func AddUserTablePresenter(writer *tablewriter.Table) *UserTable {
	writer.SetHeader([]string{"Username", "Email", "Password", "Role"})
	writer.SetRowLine(true)
	return &UserTable{writer: writer}
}

func (t *UserTable) DisplayUsers(users []models.User) {
	t.writer.ClearRows()
	for _, user := range users {
		t.writer.Append([]string{
			strconv.Itoa(user.UserID),
			user.Username,
			user.Email,
			user.Role,
			user.CreatedAt.Format("2006-01-02 15:04:05"),
			strconv.FormatBool(user.Active),
		})
	}
	t.writer.Render()
}

func (t *UserTable) DisplayAddUser(user models.User) {
	t.writer.ClearRows()
	t.writer.Append([]string{
		user.Username,
		user.Email,
		user.Password,
		user.Role,
	})
	t.writer.Render()
}
