/**
 * @Author: lenovo
 * @Description:
 * @File:  tag
 * @Version: 1.0.0
 * @Date: 2023/05/31 19:23
 */

package query

import (
	"mognolia/internal/dao"
	"mognolia/internal/model/automigrate"
)

type tag struct{}

func NewTag() *tag {
	return &tag{}
}

func (tag) AddTagForMovie(movieID uint, tags []string) error {
	result := dao.Group.DB.Model(&automigrate.Tag{}).Where("movie_id = ?", movieID).Updates(&automigrate.Tag{Tags: tags})
	return result.Error
}

func (tag) GetTagsFromMovie(movieID uint) (automigrate.Tag, error) {
	var u automigrate.Tag
	result := dao.Group.DB.Model(&automigrate.Tag{}).Where("movie_id = ?", movieID).Find(&u)
	return u, result.Error
}
