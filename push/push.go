/*
 * @Author: zhangkaiwei 1126763237@qq.com
 * @Date: 2024-05-12 18:46:16
 * @LastEditors: zhangkaiwei 1126763237@qq.com
 * @LastEditTime: 2024-05-12 18:46:29
 * @FilePath: \protocol\push\push.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
// Copyright © 2023 aetim. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package push

import "errors"

func (x *PushMsgReq) Check() error {
	if x.MsgData == nil {
		return errors.New("MsgData is empty")
	}
	if err := x.MsgData.Check(); err != nil {
		return err
	}
	if x.ConversationID == "" {
		return errors.New("ConversationID is empty")
	}
	return nil
}

func (x *DelUserPushTokenReq) Check() error {
	if x.UserID == "" {
		return errors.New("UserID is empty")
	}
	if x.PlatformID < 1 || x.PlatformID > 9 {
		return errors.New("PlatformID is invalid")
	}
	return nil
}
