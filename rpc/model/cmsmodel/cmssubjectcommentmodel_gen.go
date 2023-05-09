// Code generated by goctl. DO NOT EDIT.

package cmsmodel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	cmsSubjectCommentFieldNames          = builder.RawFieldNames(&CmsSubjectComment{})
	cmsSubjectCommentRows                = strings.Join(cmsSubjectCommentFieldNames, ",")
	cmsSubjectCommentRowsExpectAutoSet   = strings.Join(stringx.Remove(cmsSubjectCommentFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	cmsSubjectCommentRowsWithPlaceHolder = strings.Join(stringx.Remove(cmsSubjectCommentFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	cmsSubjectCommentModel interface {
		Insert(ctx context.Context, data *CmsSubjectComment) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*CmsSubjectComment, error)
		Update(ctx context.Context, data *CmsSubjectComment) error
		Delete(ctx context.Context, id int64) error
	}

	defaultCmsSubjectCommentModel struct {
		conn  sqlx.SqlConn
		table string
	}

	CmsSubjectComment struct {
		Id             int64     `db:"id"`
		SubjectId      int64     `db:"subject_id"`
		MemberNickName string    `db:"member_nick_name"`
		MemberIcon     string    `db:"member_icon"`
		Content        string    `db:"content"`
		CreateTime     time.Time `db:"create_time"`
		ShowStatus     int64     `db:"show_status"`
	}
)

func newCmsSubjectCommentModel(conn sqlx.SqlConn) *defaultCmsSubjectCommentModel {
	return &defaultCmsSubjectCommentModel{
		conn:  conn,
		table: "`cms_subject_comment`",
	}
}

func (m *defaultCmsSubjectCommentModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultCmsSubjectCommentModel) FindOne(ctx context.Context, id int64) (*CmsSubjectComment, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", cmsSubjectCommentRows, m.table)
	var resp CmsSubjectComment
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultCmsSubjectCommentModel) Insert(ctx context.Context, data *CmsSubjectComment) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, cmsSubjectCommentRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.SubjectId, data.MemberNickName, data.MemberIcon, data.Content, data.ShowStatus)
	return ret, err
}

func (m *defaultCmsSubjectCommentModel) Update(ctx context.Context, data *CmsSubjectComment) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, cmsSubjectCommentRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.SubjectId, data.MemberNickName, data.MemberIcon, data.Content, data.ShowStatus, data.Id)
	return err
}

func (m *defaultCmsSubjectCommentModel) tableName() string {
	return m.table
}
