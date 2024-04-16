/**
 * @Author: lenovo
 * @Description:
 * @File:  movie
 * @Version: 1.0.0
 * @Date: 2023/05/31 8:53
 */

package reply

type CreateMovieRly struct {
	MovieID uint `json:"movieID"`
}
type GetMovieRly struct {
	MovieInfo []MovieInfo `json:"movieInfo"`
	Total     int64       `json:"total"`
}

type MovieInfo struct {
	MovieID   uint     `json:"movieID"`
	Avatar    string   `json:"avatar"`
	Name      string   `json:"name"`
	Content   string   `json:"content"`
	Actors    []string `json:"actors"`
	Director  string   `json:"director"`
	Duration  int64    `json:"duration"`
	ShowTime  string   `json:"showTime"`
	Score     float32  `json:"score"`
	BoxOffice float32  `json:"boxOffice"`
}

type GetMovieDetails struct {
	MovieInfoRow MovieInfo `json:"movieInfo"`   //电影的基本信息
	IsComment    bool      `json:"isComment" `  //用户是否评论过
	IsFavorite   bool      `json:"isFavorite" ` //用户是否兴趣
	Tag          []string  `json:"tag"`         //电影的标签
}
