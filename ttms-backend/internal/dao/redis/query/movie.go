/**
 * @Author: lenovo
 * @Description:
 * @File:  movie
 * @Version: 1.0.0
 * @Date: 2023/05/31 16:55
 */

package query

import (
	"context"
	"github.com/go-redis/redis/v8"
	"mognolia/internal/pkg/utils"
)

// 使用zset作为排行榜
const (
	MovieReadCountRank = "movieReadCountRank"
)

// AddReadCountToMovie 增加并返回电影的访问量
func (q *Queries) AddReadCountToMovie(ctx context.Context, movieID uint) (int64, error) {
	res, err := q.redis.ZIncrBy(ctx, MovieReadCountRank, 1, utils.IDToSting(movieID)).Result()
	return int64(res), err
}

// FlushDataByMovieID 批量获取电影的访问量清零
func (q *Queries) FlushDataByMovieID(ctx context.Context, movieIDs []uint) ([]float64, error) {
	keys := make([]string, 0)
	for _, movieID := range movieIDs {
		keys = append(keys, utils.IDToSting(movieID))
	}
	res := make([]float64, 0)
	for _, key := range keys {
		score, err := q.redis.ZScore(ctx, MovieReadCountRank, key).Result()
		if err != nil {
			return nil, err
		}
		res = append(res, score)
		// 清零操作
		err = q.redis.ZRem(ctx, MovieReadCountRank, key).Err()
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

// FlushAllData 将全部电影的访问量清零
func (q *Queries) FlushAllData(ctx context.Context) error {
	members := q.redis.ZRange(ctx, MovieReadCountRank, 0, -1).Val()
	for _, member := range members {
		if err := q.redis.ZRem(ctx, MovieReadCountRank, member).Err(); err != nil {
			return err
		}
	}
	return nil
}

// GetAllDataAndFlushThem 获取每个电影的访问量然后将其删除
func (q *Queries) GetAllDataAndFlushThem(ctx context.Context) (map[string]int, error) {
	pipe := q.redis.TxPipeline()
	nums := pipe.ZRevRangeByScoreWithScores(ctx, MovieReadCountRank, &redis.ZRangeBy{
		Min: "-1",
		Max: "+inf",
	})
	pipe.Expire(ctx, MovieReadCountRank, 0)
	_, err := pipe.Exec(ctx)
	if err != nil {
		return nil, err
	}
	rank := make(map[string]int)
	for _, v := range nums.Val() {
		rank[v.Member.(string)] = int(v.Score)
	}
	return rank, nil
}

func (q *Queries) GetMovieIDsOrderByReadCount(ctx context.Context, offset, count int) ([]int64, error) {
	nums := q.redis.ZRevRangeByScore(ctx, MovieReadCountRank, &redis.ZRangeBy{
		Min:    "-1",
		Max:    "+inf",
		Offset: int64(offset),
		Count:  int64(count),
	})
	m := make([]int64, 0)
	for _, num := range nums.Val() {
		m = append(m, utils.StringToIDMust(num))
	}
	return m, nil
}
