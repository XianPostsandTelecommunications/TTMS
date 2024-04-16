/**
 * @Author: lenovo
 * @Description:
 * @File:  ticket
 * @Version: 1.0.0
 * @Date: 2023/06/08 2:26
 */

package query

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"mognolia/internal/global"
	"mognolia/internal/pkg/utils"
	"strconv"
)

var (
	ErrTicketHasBeenLocked = errors.New("ticket has been locked")
)

func (q *Queries) LockTicket(ctx context.Context, seatsID []string, planID uint, userID uint) error {
	for _, seatID := range seatsID {
		val := q.redis.HGet(ctx, utils.IDToSting(planID), seatID).Val()
		if val != "" {
			return ErrTicketHasBeenLocked
		}
	}
	for _, seatID := range seatsID {
		if err := q.redis.HSet(ctx, utils.IDToSting(planID), seatID, userID).Err(); err != nil {
			return err
		}
		key := fmt.Sprintf("%d:%s", planID, seatID)
		if err := q.redis.Set(ctx, key, userID, global.Settings.Rule.LockTicketTime).Err(); err != nil {
			return err
		}
	}
	return nil
}

func (q *Queries) SetOrder(ctx context.Context, uuid uuid.UUID, userID uint) error {
	if err := q.redis.Set(ctx, uuid.String(), userID, global.Settings.Rule.LockOrderTime).Err(); err != nil {
		return err
	}
	return nil
}

func (q *Queries) DelSeats(ctx context.Context, seatsID []string, planID uint) error {
	pipe := q.redis.Pipeline()
	for _, seatID := range seatsID {
		pipe.HDel(ctx, utils.IDToSting(planID), seatID)
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (q *Queries) CheckLockedAndIsSelf(ctx context.Context, planID uint, userID uint) bool {
	res, err := q.redis.HGetAll(ctx, utils.IDToSting(planID)).Result()
	if err != nil {
		return false
	}
	// k:seatID
	//v:userID
	for _, v := range res {
		if strconv.Itoa(int(userID)) != v {
			return false
		}
	}
	return true
}

func (q *Queries) CheckIsLock(ctx context.Context, seatID string, planID, userID uint) bool {
	_, err := q.redis.HGet(ctx, utils.IDToSting(planID), seatID).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return true
		}
		return false
	}
	return false
}

func (q *Queries) LockOneTicket(ctx context.Context, seatID string, planID, userID uint) error {
	if err := q.redis.HSet(ctx, utils.IDToSting(planID), seatID, userID).Err(); err != nil {
		return err
	}
	key := fmt.Sprintf("%d:%s", planID, seatID)
	if err := q.redis.Set(ctx, key, userID, global.Settings.Rule.LockTicketTime).Err(); err != nil {
		return err
	}
	return nil
}

func (q *Queries) Expire2Zero(ctx context.Context, key string) error {
	if err := q.redis.Expire(ctx, key, 0).Err(); err != nil {
		return err
	}
	return nil
}
