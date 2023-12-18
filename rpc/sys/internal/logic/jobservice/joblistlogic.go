package jobservicelogic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logc"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// JobListLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:05
*/
type JobListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJobListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JobListLogic {
	return &JobListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// JobList 岗位列表
func (l *JobListLogic) JobList(in *sysclient.JobListReq) (*sysclient.JobListResp, error) {
	all, err := l.svcCtx.JobModel.FindAll(l.ctx, in)
	count, _ := l.svcCtx.JobModel.Count(l.ctx, in)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logc.Errorf(l.ctx, "查询岗位列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*sysclient.JobListData
	for _, job := range *all {
		list = append(list, &sysclient.JobListData{
			Id:             job.Id,
			JobName:        job.JobName,
			OrderNum:       job.OrderNum,
			CreateBy:       job.CreateBy,
			CreateTime:     job.CreateTime.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   job.UpdateBy.String,
			LastUpdateTime: job.UpdateTime.Time.Format("2006-01-02 15:04:05"),
			DelFlag:        job.DelFlag,
			Remarks:        job.Remarks.String,
		})
	}

	logc.Infof(l.ctx, "查询岗位列表信息,参数：%+v,响应：%+v", in, list)
	return &sysclient.JobListResp{
		Total: count,
		List:  list,
	}, nil
}
