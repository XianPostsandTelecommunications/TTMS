/**
 * @Author: lenovo
 * @Description:
 * @File:  movie
 * @Version: 1.0.0
 * @Date: 2023/05/31 11:55
 */

package query

import (
	"fmt"
	"gorm.io/gorm"
	"mognolia/internal/dao"
	"mognolia/internal/model/automigrate"
	"mognolia/internal/model/request"
	"time"
)

type movie struct{}

func NewMovie() *movie {
	return &movie{}
}
func (movie) DeleteMovieByID(movieID uint) error {
	result := dao.Group.DB.Model(&automigrate.Movie{}).Delete(&automigrate.Movie{}, movieID)
	return result.Error
}
func (movie) GetMovieByID(movieID uint) (*automigrate.Movie, error) {
	var u automigrate.Movie
	if result := dao.Group.DB.Model(&automigrate.Movie{}).Where("id = ?", movieID).Find(&u); result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &u, nil
}

func (movie) GetAreaTagsPeriodMovieOrderByPeriod(t request.GetMovieTagsAreaPeriod, size, offset int64) ([]automigrate.Movie, error) {
	type Result struct {
		ID        uint     `json:"id"`
		Actors    []string `json:"actors"`
		Avatar    string   `json:"avatar"`
		Area      string   `json:"area"`
		Name      string   `json:"name"`
		Duration  int64    `json:"duration"`
		ShowTime  string
		Score     float32 `json:"score"`
		BoxOffice float32 `json:"boxOffice"`
		Director  string  `json:"director"`
	}

	var results []Result
	if result := dao.Group.DB.Raw(`
SELECT
	m.id,
	m.actors,
	m.avatar,
	m.area,
	m.NAME,
	m.show_time,
	m.score,
	m.duration,
	m.box_office,
	m.director 
FROM
	movies m,
	(
	SELECT
		movie_id 
	FROM
		tags 
	WHERE
	JSON_CONTAINS( tags, ? )) AS tags 
WHERE
	STR_TO_DATE( m.show_time, '%Y-%m-%d' ) BETWEEN STR_TO_DATE(?, '%Y-%m-%d' ) 
	AND STR_TO_DATE( ?, '%Y-%m-%d' ) 
	AND m.area LIKE ? 
	AND m.show_time LIKE ? 
	AND m.id = tags.movie_id 
ORDER BY
	STR_TO_DATE( m.show_time, '%Y-%m-%d' ) DESC 
	LIMIT ? OFFSET ?
`, fmt.Sprintf("\"%s\"", t.Tag), time.Unix(t.StartTime, 0).Format("2006-01-02"), time.Unix(t.EndTime, 0).Format("2006-01-02"), t.Area, t.Period, size, offset).Scan(&results); result.Error != nil {
		return nil, gorm.ErrRecordNotFound
	}
	var movies []automigrate.Movie
	for _, result := range results {
		movie := automigrate.Movie{
			Model:     gorm.Model{ID: result.ID},
			Name:      result.Name,
			Area:      result.Area,
			Actors:    result.Actors,
			Avatar:    result.Avatar,
			Duration:  result.Duration,
			ShowTime:  result.ShowTime,
			Director:  result.Director,
			Score:     result.Score,
			BoxOffice: result.BoxOffice,
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

func (movie) GetAreaTagsPeriodMovieOrderByReadCount(t request.GetMovieTagsAreaPeriod, size, offset int64) ([]automigrate.Movie, error) {
	type Result struct {
		ID        uint                    `json:"id"`
		Actors    automigrate.ActorString `json:"actors"`
		Avatar    string                  `json:"avatar"`
		Area      string                  `json:"area"`
		Name      string                  `json:"name"`
		Duration  int64                   `json:"duration"`
		ShowTime  string
		Score     float32 `json:"score"`
		BoxOffice float32 `json:"boxOffice"`
		Director  string  `json:"director"`
	}

	var results []Result
	if result := dao.Group.DB.Raw(`
	SELECT
	m.id,
	m.actors,
	m.avatar,
	m.area,
	m.name,
	m.show_time,
	m.score, 
	m.duration,
	m.box_office,
	m.director
FROM
	 movies m,
	( SELECT movie_id FROM tags WHERE JSON_CONTAINS(tags, ?) ) AS   tags
WHERE m.show_time BETWEEN ? AND ?
AND m.area LIKE ?
AND m.id = tags.movie_id
order by m.visit_count desc
limit ? offset ?
`, fmt.Sprintf("\"%s\"", t.Tag), time.Unix(t.StartTime, 0).Format("2006-01-02"), time.Unix(t.EndTime, 0).Format("2006-01-02"), t.Area, size, offset).Scan(&results); result.Error != nil {
		return nil, gorm.ErrRecordNotFound
	}
	var movies []automigrate.Movie
	for _, result := range results {
		movie := automigrate.Movie{
			Model:     gorm.Model{ID: result.ID},
			Name:      result.Name,
			Area:      result.Area,
			Actors:    result.Actors,
			Avatar:    result.Avatar,
			Duration:  result.Duration,
			ShowTime:  result.ShowTime,
			Director:  result.Director,
			Score:     result.Score,
			BoxOffice: result.BoxOffice,
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

func (movie) GetAreaTagsPeriodMovieOrderByScore(t request.GetMovieTagsAreaPeriod, size, offset int64) ([]automigrate.Movie, error) {
	type Result struct {
		ID        uint                    `json:"id"`
		Actors    automigrate.ActorString `json:"actors"`
		Avatar    string                  `json:"avatar"`
		Area      string                  `json:"area"`
		Name      string                  `json:"name"`
		Duration  int64                   `json:"duration"`
		ShowTime  string
		Score     float32 `json:"score"`
		BoxOffice float32 `json:"boxOffice"`
		Director  string  `json:"director"`
	}

	var results []Result
	rows := dao.Group.DB.Raw(`
SELECT
	m.id,
	m.actors,
	m.avatar,
	m.area,
	m.name,
	m.show_time,
	m.score, 
	m.duration, 
	m.box_office,
	m.director
FROM
	 movies m,
	( SELECT movie_id FROM tags WHERE JSON_CONTAINS(tags, ?) ) AS   tags
WHERE m.show_time BETWEEN ? AND ?
AND m.area LIKE ?
AND m.id = tags.movie_id
order by m.score desc
limit ? offset ?
`, fmt.Sprintf("\"%s\"", t.Tag), time.Unix(t.StartTime, 0).Format("2006-01-02"), time.Unix(t.EndTime, 0).Format("2006-01-02"), t.Area, size, offset).Scan(&results)

	if rows.Error != nil {
		return nil, gorm.ErrRecordNotFound
	}
	var movies []automigrate.Movie
	for _, result := range results {
		movie := automigrate.Movie{
			Model:     gorm.Model{ID: result.ID},
			Name:      result.Name,
			Area:      result.Area,
			Actors:    result.Actors,
			Avatar:    result.Avatar,
			Duration:  result.Duration,
			ShowTime:  result.ShowTime,
			Director:  result.Director,
			Score:     result.Score,
			BoxOffice: result.BoxOffice,
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

func (movie) GetKey(key string, page int64) ([]automigrate.Movie, error) {
	var movies []automigrate.Movie
	if result := dao.Group.DB.Model(&automigrate.Movie{}).Scopes(Paginate(page, 0)).Where("name like ?", "%"+key+"%").Or("content like ?", "%"+key+"%").Find(&movies); result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return movies, nil
}

func (movie) GetMoviesOrderbyReadCount(page int64) ([]automigrate.Movie, error) {
	var movies []automigrate.Movie
	if result := dao.Group.DB.Model(&automigrate.Movie{}).Scopes(Paginate(page, 0)).Order("visit_count DESC").Find(&movies); result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return movies, nil
}

func (movie) UpdateMovieInfo(info request.UpdateMovieInfo) error {
	if result := dao.Group.DB.Model(&automigrate.Movie{}).Where("id =?", info.MovieID).Updates(automigrate.Movie{
		Name:     info.Name,
		Area:     info.Area,
		Actors:   info.Actors,
		Content:  info.Content,
		Director: info.Director,
		Avatar:   info.Avatar,
	}); result.RowsAffected == 0 {
		return result.Error
	}
	return nil
}

func (movie) GetAllMoviesInfo() ([]automigrate.Movie, error) {
	var movies []automigrate.Movie
	if result := dao.Group.DB.Model(&automigrate.Movie{}).Find(&movies); result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return movies, nil
}

func (movie) GetMovieOrderByFavorNum(num int) ([]automigrate.Movie, error) {
	type result struct {
		ID        uint                    `json:"id"`
		Actors    automigrate.ActorString `json:"actors"`
		Avatar    string                  `json:"avatar"`
		Area      string                  `json:"area"`
		Name      string                  `json:"name"`
		Duration  int64                   `json:"duration"`
		ShowTime  string
		Score     float32 `json:"score"`
		BoxOffice float32 `json:"boxOffice"`
		Director  string  `json:"director"`
	}

	var Results []result

	t := dao.Group.DB.Raw(`
	SELECT
	m.id,
	m.actors,
	m.avatar,
	m.area,
	m.name,
	m.show_time,
	m.score, 
	m.duration,
	m.box_office,
	m.director
FROM
    movies m
JOIN
    (
        SELECT
            movie_id,
            COUNT(movie_id) AS movie_count
        FROM
            people_movies
        GROUP BY
            movie_id
        ORDER BY
            movie_count DESC
        LIMIT ?
    ) AS pm ON pm.movie_id = m.id;
`, num).Scan(&Results)

	if t.Error != nil {
		return nil, gorm.ErrRecordNotFound
	}

	var movies []automigrate.Movie
	for _, result := range Results {
		movies = append(movies, automigrate.Movie{
			Model:     gorm.Model{ID: result.ID},
			Name:      result.Name,
			Area:      result.Area,
			Actors:    result.Actors,
			Avatar:    result.Avatar,
			Duration:  result.Duration,
			ShowTime:  result.ShowTime,
			Director:  result.Director,
			Score:     result.Score,
			BoxOffice: result.BoxOffice,
		})
	}
	return movies, nil
}
