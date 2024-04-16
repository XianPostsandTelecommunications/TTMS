/**
 * @Author: lenovo
 * @Description:
 * @File:  movie_test
 * @Version: 1.0.0
 * @Date: 2023/05/31 17:26
 */

package query

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueries_AddReadCountToMovie(t *testing.T) {
	data, err := client.AddReadCountToMovie(context.Background(), 7)
	assert.NoError(t, err)
	t.Log(data)
}

func TestFlushAllData(t *testing.T) {
	err := client.FlushAllData(context.Background())
	assert.NoError(t, err)
}

func TestQueries_FlushDataByMovieID(t *testing.T) {
	data, err := client.FlushDataByMovieID(context.Background(), []uint{1})
	assert.NoError(t, err)
	t.Log(data)
}

func TestGetAllDataAndFlushThem(t *testing.T) {
	m, err := client.GetAllDataAndFlushThem(context.Background())
	assert.NoError(t, err)
	for k, v := range m {
		t.Log(k, v)
	}
}

func TestFlushDataByMovieID(t *testing.T) {
	data, err := client.FlushDataByMovieID(context.Background(), []uint{1, 2, 3})
	assert.NoError(t, err)
	t.Log(data)
}
