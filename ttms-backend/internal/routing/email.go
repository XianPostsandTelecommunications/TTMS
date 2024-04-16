/**
 * @Author: lenovo
 * @Description:
 * @File:  email
 * @Version: 1.0.0
 * @Date: 2023/05/29 12:54
 */

package routing

import (
	v1 "mognolia/internal/api/v1"

	"github.com/gin-gonic/gin"
)

type email struct{}

func (email) Init(r *gin.RouterGroup) {
	g := r.Group("/email")
	{
		g.POST("/send", v1.Group.Email.SendEmail)
	}

}
