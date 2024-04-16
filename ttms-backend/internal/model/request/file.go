/**
 * @Author: lenovo
 * @Description:
 * @File:  file
 * @Version: 1.0.0
 * @Date: 2023/05/31 11:05
 */

package request

import "mime/multipart"

type FileParam struct {
	File *multipart.FileHeader `json:"File"`
}
