/**
 * @Author: lenovo
 * @Description:
 * @File:  movie
 * @Version: 1.0.0
 * @Date: 2023/05/31 8:43
 */

package routing

import (
	"github.com/gin-gonic/gin"
	v1 "mognolia/internal/api/v1"
	"mognolia/internal/middleware"
)

type movie struct{}

func (m *movie) Init(r *gin.RouterGroup) {

	g := r.Group("/movie")
	g1 := g.Use(middleware.Auth())
	{
		g1.POST("/byTagAreaPeriod", v1.Group.Movie.GetMovieByTagAreaPeriod)
		g1.POST("/byNameOrContent", v1.Group.Movie.GetMovieInfoByNameOrContent)
		g1.GET("/byExpectedNums/:page", v1.Group.Movie.GetMovieOrderByExpectedNum)
		g1.GET("/byReadCount", v1.Group.Movie.GetMoviesByReadCount)
		g1.POST("/details", v1.Group.Movie.GetMovieDetails)
	}

	g2 := g.Use(middleware.Auth(), middleware.AuthManager())
	{
		g2.POST("/create", v1.Group.Movie.CreateMovie)
		g2.DELETE("", v1.Group.Movie.DeleteMovie)
		g2.PUT("/update", v1.Group.Movie.UpdateMovieInfo)
		g1.GET("/list", v1.Group.Movie.GetAllMovie)
	}

}
