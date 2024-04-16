/**
 * @Author: lenovo
 * @Description:
 * @File:  query
 * @Version: 1.0.0
 * @Date: 2023/05/29 10:27
 */

package query

import (
	"github.com/go-redis/redis/v8"
)

var client = &Queries{}

type Queries struct {
	redis *redis.Client
}

func New(client *redis.Client) *Queries {
	return &Queries{
		redis: client,
	}
}
