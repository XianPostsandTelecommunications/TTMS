/**
 * @Author: lenovo
 * @Description:
 * @File:  auto
 * @Version: 1.0.0
 * @Date: 2023/06/03 17:41
 */

package logic

import (
	"context"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"mognolia/internal/dao"
	"mognolia/internal/dao/mysql/query"
	"mognolia/internal/global"
	"mognolia/internal/model/automigrate"
	t "mognolia/internal/pkg/task"
	"mognolia/internal/pkg/utils"
)

type auto struct{}

var tasks []t.Task

func (a *auto) Work() {
	ctx := context.Background()
	movieReadCountFlush2DBTask := t.Task{
		Name:            "movieReadCountFlush2DBTask",
		Ctx:             ctx,
		TaskDuration:    global.Settings.Auto.AutoFlushReadCount2DBTime,
		TimeoutDuration: global.Settings.Serve.DefaultContextTimeout,
		F:               movieVisitCountFlush(),
	}
	peopleFavorToCache := t.Task{
		Name:            "peopleFavorToCache",
		Ctx:             ctx,
		TaskDuration:    global.Settings.Auto.AutoFlushReadCount2DBTime,
		TimeoutDuration: global.Settings.Serve.DefaultContextTimeout,
		F:               peopleFavorToCache(),
	}
	deleteTimeOutPlans := t.Task{
		Name:            "deleteTimeOutPlans",
		Ctx:             ctx,
		TaskDuration:    global.Settings.Auto.DeleteOutTimeTime,
		TimeoutDuration: global.Settings.Serve.DefaultContextTimeout,
		F:               deleteTimeOutPlans(),
	}

	tasks = append(tasks, movieReadCountFlush2DBTask, peopleFavorToCache, deleteTimeOutPlans)
	a.Init()
}

func (a *auto) Init() {
	for i := range tasks {
		t.NewTickerTask(tasks[i])
	}
}

func movieVisitCountFlush() t.DoFunc {
	return func(parentCtx context.Context) {
		global.Logger.Info("start movieVisitCountFlush task....")
		ctx, cancle := context.WithTimeout(parentCtx, global.Settings.Serve.DefaultContextTimeout)
		defer cancle()
		readMap, err := dao.Group.Redis.GetAllDataAndFlushThem(ctx)
		if err != nil {
			global.Logger.Info("GetAllDataAndFlushThem failed...")
			return
		}
		tx := dao.Group.DB.Begin()
		for movieID, count := range readMap {
			i := utils.StringToIDMust(movieID)
			result := tx.Model(&automigrate.Movie{}).Where("id = ?", i).UpdateColumn("visit_count", gorm.Expr("visit_count + ?", count))
			if result.Error != nil {
				tx.Rollback()
				global.Logger.Info("GetMovie failed")
				return
			}
		}
		tx.Commit()
	}
}

func peopleFavorToCache() t.DoFunc {
	return func(parentCtx context.Context) {
		global.Logger.Info("start peopleFavorToCache task ...")
		ctx, cancel := context.WithTimeout(parentCtx, global.Settings.Serve.DefaultContextTimeout)
		defer cancel()
		page, size := global.Settings.Rule.DefaultUserFavorPage, global.Settings.Rule.DefaultUserFavorSize
		q := query.NewMovie()
		movieInfos, err := q.GetMovieOrderByFavorNum(page * size)
		if err != nil {
			zap.S().Info("get nothing ...")
			return
		}

		for i := 1; i <= page; i++ {
			start := (i - 1) * size
			end := start + size
			ok := false
			if end >= len(movieInfos) {
				end = len(movieInfos)
				ok = true
			}
			if err := dao.Group.Redis.Set(ctx, utils.IDToSting(uint(i)), movieInfos[start:end]); err != nil {
				zap.S().Infof("set to cache failed,err:%v", err)
				return
			}
			if ok {
				break
			}
		}
	}
}

func deleteTimeOutPlans() t.DoFunc {
	return func(parentCtx context.Context) {
		global.Logger.Info("start deleteTimeOutPlans ...")
		var ids []uint

		if result := dao.Group.DB.Raw(`
		select p.id
		from plans p
		where p.end_at <= now()	
		AND p.deleted_at IS NULL
`).Scan(&ids); result.Error != nil {
			zap.S().Info("select p.id failed err:%v", result.Error)
			return
		}
		if result := dao.Group.DB.Model(&automigrate.Ticket{}).Where("plan_id in ?", ids).Delete(&automigrate.Ticket{}); result.RowsAffected == 0 {
			zap.S().Infof("delete ticket failed err:%v", result.Error)
			return
		}
		if err := dao.Group.DB.Model(&automigrate.Plan{}).Where("end_at <= NOW()").Delete(&automigrate.Plan{}); err != nil {
			zap.S().Info("delete timeOutPlans failed,err:%v", err)
			return
		}
	}
}
