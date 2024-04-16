/**
 * @Author: lenovo
 * @Description:
 * @File:  tag
 * @Version: 1.0.0
 * @Date: 2023/05/31 18:59
 */

package logic

import (
	"mognolia/internal/dao/mysql/query"
	"mognolia/internal/model/reply"
	"mognolia/internal/myerr"
	"mognolia/internal/pkg/app/errcode"
)

type tag struct{}

func (t *tag) AddTagsForMovie(movieID uint, tags []string) errcode.Err {
	q := query.NewMovie()
	_, err := q.GetMovieByID(movieID)
	if err != nil {
		return myerr.NoRecords
	}
	qt := query.NewTag()
	tagsInfo, err := qt.GetTagsFromMovie(movieID)
	if err != nil {
		return errcode.ErrServer.WithDetails(err.Error())
	}
	for _, tag := range tagsInfo.Tags {
		t := tag
		tags = append(tags, t)
	}
	if err := qt.AddTagForMovie(movieID, tags); err != nil {
		return errcode.ErrServer
	}
	return nil
}

func (t *tag) GetTagsFromMovie(movieID uint) (*reply.GetTagFromMovie, errcode.Err) {
	res := &reply.GetTagFromMovie{}
	q := query.NewTag()
	tagInfo, err := q.GetTagsFromMovie(movieID)
	if err != nil {
		return nil, errcode.ErrServer
	}
	for _, info := range tagInfo.Tags {
		res.Tags = append(res.Tags, info)
	}
	return res, nil
}
