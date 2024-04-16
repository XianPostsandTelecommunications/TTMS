/**
 * @Author: lenovo
 * @Description:
 * @File:  common
 * @Version: 1.0.0
 * @Date: 2023/05/31 12:39
 */

package query

import (
	"context"
	"encoding/json"
	"mognolia/internal/global"
	"mognolia/internal/pkg/singleflight"
	"time"
)

func (q *Queries) Set(ctx context.Context, key string, val interface{}) error {
	data, err := json.Marshal(val)
	if err != nil {
		return err
	}
	return q.redis.Set(ctx, key, data, global.Settings.Redis.CacheTime).Err()
}

func (q *Queries) SetTimeOut(ctx context.Context, key string, val interface{}, duration time.Duration) error {
	data, err := json.Marshal(val)
	if err != nil {
		return err
	}
	return q.redis.Set(ctx, key, data, duration).Err()
}
func (q *Queries) Get(ctx context.Context, key string, val interface{}) error {

	//保证原子性,只执行一次
	result, err := singleflight.Group.Do(key, func() (interface{}, error) {
		result, err := q.redis.Get(ctx, key).Result()
		if err != nil {
			return nil, err
		}
		return result, err
	})
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(result.(string)), &val)
}
