package tables

import (
	"clothing-pair-project/internal/models"

	"strconv"

	"github.com/olekukonko/tablewriter"
)

type TableUserDisplayer struct {
	writer *tablewriter.Table
}

func NewTableAllUsersDisplayer(writer *tablewriter.Table) *TableUserDisplayer {
	writer.SetHeader([]string{"ID", "Username", "Email", "Role", "Created At", "Active"})
	writer.SetRowLine(true)
	return &TableUserDisplayer{writer: writer}
}

func (t *TableUserDisplayer) DisplayAllUser(users []models.User) {
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
