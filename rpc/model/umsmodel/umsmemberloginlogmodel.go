package umsmodel

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ UmsMemberLoginLogModel = (*customUmsMemberLoginLogModel)(nil)

type (
	// UmsMemberLoginLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUmsMemberLoginLogModel.
	UmsMemberLoginLogModel interface {
		umsMemberLoginLogModel
		Count(ctx context.Context) (int64, error)
		FindAll(ctx context.Context, Current int64, PageSize int64) (*[]UmsMemberLoginLog, error)
		DeleteByIds(ctx context.Context, ids []int64) error
	}

	customUmsMemberLoginLogModel struct {
		*defaultUmsMemberLoginLogModel
	}
)

// NewUmsMemberLoginLogModel returns a model for the database table.
func NewUmsMemberLoginLogModel(conn sqlx.SqlConn) UmsMemberLoginLogModel {
	return &customUmsMemberLoginLogModel{
		defaultUmsMemberLoginLogModel: newUmsMemberLoginLogModel(conn),
	}
}

func (m *customUmsMemberLoginLogModel) FindAll(ctx context.Context, Current int64, PageSize int64) (*[]UmsMemberLoginLog, error) {

	query := fmt.Sprintf("select %s from %s limit ?,?", umsMemberLoginLogRows, m.table)
	var resp []UmsMemberLoginLog
	err := m.conn.QueryRows(&resp, query, (Current-1)*PageSize, PageSize)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customUmsMemberLoginLogModel) Count(ctx context.Context) (int64, error) {
	query := fmt.Sprintf("select count(*) as count from %s", m.table)

	var count int64
	err := m.conn.QueryRow(&count, query)

	switch err {
	case nil:
		return count, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}

func (m *customUmsMemberLoginLogModel) DeleteByIds(ctx context.Context, ids []int64) error {
	query := fmt.Sprintf("delete from %s where `id` in (?)", m.table)
	_, err := m.conn.ExecCtx(ctx, query, strings.Replace(strings.Trim(fmt.Sprint(ids), "[]"), " ", ",", -1))
	return err
}
