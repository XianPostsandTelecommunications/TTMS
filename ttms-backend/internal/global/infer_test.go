/**
 * @Author: lenovo
 * @Description:
 * @File:  infer_test
 * @Version: 1.0.0
 * @Date: 2023/05/29 8:32
 */

package global

import "testing"

func TestInferRootDir(t *testing.T) {
	inferRootDir()
	t.Log(RootDir) // D:\workspace\go\src\mognolia
}
