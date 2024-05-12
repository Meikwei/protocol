/*
 * @Author: zhangkaiwei 1126763237@qq.com
 * @Date: 2024-05-11 23:47:53
 * @LastEditors: zhangkaiwei 1126763237@qq.com
 * @LastEditTime: 2024-05-11 23:48:26
 * @FilePath: \protocol\sdkws\sdkws.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package sdkws

import (
	"errors"
)

func (x *MsgData) Check() error {
	if x.SendID == "" {
		return errors.New("sendID is empty")
	}
	if x.Content == nil {
		return errors.New("content is empty")
	}
	return nil
}