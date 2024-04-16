/**
 * @Author: lenovo
 * @Description:
 * @File:  email
 * @Version: 1.0.0
 * @Date: 2023/05/29 12:13
 */

package request

type SendEmailParams struct {
	Email string `json:"Email" binding:"required"`
}
