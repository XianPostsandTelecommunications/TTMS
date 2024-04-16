/**
 * @Author: lenovo
 * @Description:
 * @File:  file
 * @Version: 1.0.0
 * @Date: 2023/05/31 11:07
 */

package logic

import (
	"mognolia/internal/model/reply"
	"mognolia/internal/model/request"
	"mognolia/internal/pkg/app/errcode"
	"mognolia/internal/pkg/oss"
)

type file struct{}

func (file) UploadFile(param request.FileParam) (*reply.FileRly, errcode.Err) {
	ossClient := oss.NewOss()
	URL, _, err := ossClient.UploadFile(param.File)
	if err != nil {
		return nil, errcode.ErrServer
	}
	return &reply.FileRly{URL: URL}, nil
}
