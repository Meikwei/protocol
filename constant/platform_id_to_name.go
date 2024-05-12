/*
 * @Author: zhangkaiwei 1126763237@qq.com
 * @Date: 2024-05-12 08:10:57
 * @LastEditors: zhangkaiwei 1126763237@qq.com
 * @LastEditTime: 2024-05-12 11:52:41
 * @FilePath: \protocol\constant\platform_id_to_name.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package constant

// 定义平台相关的常量和转换函数。
// 平台ID与平台字符串的映射，用于在不同上下文间转换。


const (
	// 平台ID。
	IOSPlatformID        = 1 // iOS平台
	AndroidPlatformID    = 2 // Android平台
	WindowsPlatformID    = 3 // Windows平台
	OSXPlatformID        = 4 // macOS平台
	WebPlatformID        = 5 // Web平台
	MiniWebPlatformID    = 6 // MiniWeb平台
	LinuxPlatformID      = 7 // Linux平台
	AndroidPadPlatformID = 8 // Android平板平台
	IPadPlatformID       = 9 // iPad平台
	AdminPlatformID      = 10 // 管理员平台

	// 平台字符串。
	IOSPlatformStr        = "IOS"        // iOS平台字符串
	AndroidPlatformStr    = "Android"    // Android平台字符串
	WindowsPlatformStr    = "Windows"    // Windows平台字符串
	OSXPlatformStr        = "OSX"        // macOS平台字符串
	WebPlatformStr        = "Web"        // Web平台字符串
	MiniWebPlatformStr    = "MiniWeb"    // MiniWeb平台字符串
	LinuxPlatformStr      = "Linux"      // Linux平台字符串
	AndroidPadPlatformStr = "APad"       // Android平板平台字符串
	IPadPlatformStr       = "IPad"       // iPad平台字符串
	AdminPlatformStr      = "Admin"      // 管理员平台字符串

	// 终端类型。
	TerminalPC     = "PC"     // 个人电脑终端
	TerminalMobile = "Mobile" // 移动设备终端
)

// 平台ID到平台名称的映射。
var PlatformID2Name = map[int]string{
	IOSPlatformID:        IOSPlatformStr,
	AndroidPlatformID:    AndroidPlatformStr,
	WindowsPlatformID:    WindowsPlatformStr,
	OSXPlatformID:        OSXPlatformStr,
	WebPlatformID:        WebPlatformStr,
	MiniWebPlatformID:    MiniWebPlatformStr,
	LinuxPlatformID:      LinuxPlatformStr,
	AndroidPadPlatformID: AndroidPadPlatformStr,
	IPadPlatformID:       IPadPlatformStr,
	AdminPlatformID:      AdminPlatformStr,
}

// 平台名称到平台ID的映射。
var PlatformName2ID = map[string]int{
	IOSPlatformStr:        IOSPlatformID,
	AndroidPlatformStr:    AndroidPlatformID,
	WindowsPlatformStr:    WindowsPlatformID,
	OSXPlatformStr:        OSXPlatformID,
	WebPlatformStr:        WebPlatformID,
	MiniWebPlatformStr:    MiniWebPlatformID,
	LinuxPlatformStr:      LinuxPlatformID,
	AndroidPadPlatformStr: AndroidPadPlatformID,
	IPadPlatformStr:       IPadPlatformID,
	AdminPlatformStr:      AdminPlatformID,
}

// 平台名称到终端类型的映射。
var PlatformName2class = map[string]string{
	IOSPlatformStr:     TerminalMobile,
	AndroidPlatformStr: TerminalMobile,
	MiniWebPlatformStr: WebPlatformStr,
	WebPlatformStr:     WebPlatformStr,
	WindowsPlatformStr: TerminalPC,
	OSXPlatformStr:     TerminalPC,
	LinuxPlatformStr:   TerminalPC,
}

// 平台ID到终端类型的映射。
var PlatformID2class = map[int]string{
	IOSPlatformID:     TerminalMobile,
	AndroidPlatformID: TerminalMobile,
	MiniWebPlatformID: WebPlatformStr,
	WebPlatformID:     WebPlatformStr,
	WindowsPlatformID: TerminalPC,
	OSXPlatformID:     TerminalPC,
	LinuxPlatformID:   TerminalPC,
}

// PlatformIDToName 根据平台ID返回对应的平台名称。
// 参数：
//   num int - 平台ID
// 返回值：
//   string - 平台名称
func PlatformIDToName(num int) string {
	return PlatformID2Name[num]
}

// PlatformNameToID 根据平台名称返回对应的平台ID。
// 参数：
//   name string - 平台名称
// 返回值：
//   int - 平台ID
func PlatformNameToID(name string) int {
	return PlatformName2ID[name]
}

// PlatformNameToClass 根据平台名称返回对应的终端类型。
// 参数：
//   name string - 平台名称
// 返回值：
//   string - 终端类型
func PlatformNameToClass(name string) string {
	return PlatformName2class[name]
}

// PlatformIDToClass 根据平台ID返回对应的终端类型。
// 参数：
//   num int - 平台ID
// 返回值：
//   string - 终端类型
func PlatformIDToClass(num int) string {
	return PlatformID2class[num]
}