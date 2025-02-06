package tables

import (
	"clothing-pair-project/internal/models"

	"strconv"

	"github.com/olekukonko/tablewriter"
)

type TableUserDisplayer struct {
	writer *tablewriter.Table
}

func NewTableUsersDisplayer(writer *tablewriter.Table) *TableUserDisplayer {
	writer.SetHeader([]string{"ID", "Username", "Email", "Role", "Created At", "Active"})
	writer.SetRowLine(true)
	return &TableUserDisplayer{writer: writer}
}

func NewTableAddUserDisplayer(writer *tablewriter.Table) *TableUserDisplayer {
	writer.SetHeader([]string{"Username", "Email", "Password", "Role"})
	writer.SetRowLine(true)
	return &TableUserDisplayer{writer: writer}
}

func (t *TableUserDisplayer) DisplayUsers(users []models.User) {
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

func (t *TableUserDisplayer) DisplayAddUser(user models.User) {
	t.writer.ClearRows()
	t.writer.Append([]string{
		user.Username,
		user.Email,
		user.Password,
		user.Role,
	})
	t.writer.Render()
}
