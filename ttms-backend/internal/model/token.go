/**
 * @Author: lenovo
 * @Description:
 * @File:  token
 * @Version: 1.0.0
 * @Date: 2023/05/29 18:54
 */

package model

import (
	"encoding/json"
	"mognolia/internal/model/automigrate"
)

type Content struct {
	ID   uint
	Role automigrate.Roler
}

func NewContent(id uint, role automigrate.Roler) *Content {
	return &Content{
		ID:   id,
		Role: role,
	}
}

func (c *Content) Marshal() ([]byte, error) {
	return json.Marshal(c)
}

func (c *Content) UnMarshal(data []byte) error {
	return json.Unmarshal(data, &c)
}
