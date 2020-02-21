package database

import (
	"context"
	"fmt"

	"github.com/frozentech/comment/model"
)

const (
	// SQLDropTableComment drop table
	SQLDropTableComment = `DROP TABLE IF EXISTS comments;`

	// SQLCreateTableComment create table
	SQLCreateTableComment = `CREATE TABLE IF NOT EXISTS comments (
			id VARCHAR(36) PRIMARY KEY,
		  member VARCHAR(255),
		  organization VARCHAR(50),
		  message TEXT,
		  visible INTEGER,
		  created_at VARCHAR(30),
		  updated_at VARCHAR(30),
		  deleted_at VARCHAR(30));`

	// SQLFetchComment  select comment
	SQLFetchComment = `SELECT id, member, organization, message, visible, created_at, updated_at, deleted_at FROM %s WHERE organization = '%s' AND visible = 1 ORDER BY created_at DESC;`

	// SQLDeleteComment delete comments
	SQLDeleteComment = `UPDATE %s SET visible=0 WHERE organization = '%s';`

	// SQLInsertComment  insert comment
	SQLInsertComment = `INSERT INTO %s (id,member,organization,message,visible,created_at,updated_at,deleted_at)
		VALUES ('%s','%s','%s','%s',%d,'%s','%s','%s');`
)

// DeleteComment delete comments
func DeleteComment(ctx context.Context, name string) (err error) {

	sql := fmt.Sprintf(SQLDeleteComment, model.TableComment, name)

	_, err = Connection.Exec(ctx, sql)

	return
}

// GetComment query comments
func GetComment(ctx context.Context, name string) (comments []model.Comment, err error) {

	sql := fmt.Sprintf(SQLFetchComment, model.TableComment, name)

	rows, err := Connection.Query(ctx, sql)
	defer rows.Close()
	if err != nil {
		return
	}

	rowsRead := 0

	c := model.Comment{}
	for rows.Next() {
		rows.Scan(&c.ID, &c.MemberAsString, &c.Organization, &c.Message, &c.Visible, &c.CreatedAt, &c.UpdatedAt, &c.DeletedAt)
		rowsRead++
		comments = append(comments, c)
	}

	if rowsRead == 0 {
		err = fmt.Errorf("No record found")
	}

	return
}

// PostComment post comment
func PostComment(ctx context.Context, c model.Comment) (err error) {

	sql := fmt.Sprintf(SQLInsertComment, model.TableComment, c.ID, c.MemberAsString, c.Organization, c.Message, c.Visible, c.CreatedAt, c.UpdatedAt, c.DeletedAt)
	_, err = Connection.Exec(ctx, sql)

	return
}
