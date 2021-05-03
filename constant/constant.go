package constant

// 状态码, 由5位构成，服务级别错误码：（eg：由1位表示，1表示系统级错误，2表示普通错误）
// 模块级错误码：由第二位和第三位两位构成（eg：01表示用户模块，02表示文章模块）
// 具体错误码：由于第四位和第五位表示，表示具体的错误（eg：01表示用户名错误，02表示密码错误等）
const (
	// 服务级别模块
	Success = 10000

	// 帐号模块
	QrCodeEncodeError      = 20101
	QrCodeCacheError       = 20102
	UidNotExistError       = 20103
	QrCodeConvertJsonError = 20104
	TemplateParseError     = 20105
	TemplateExecuteError   = 20106
	CacheNotSetError       = 20107

	// 实用工具模块
	UrlJoinError          = 20201
	LocalIpNotObtainError = 20202
)

var CodeMsg = map[int]string{
	// 服务级别状态码
	Success: "ok",
	// 帐号模块
	UrlJoinError:           "URL拼接错误",
	QrCodeEncodeError:      "二维码生成失败",
	QrCodeCacheError:       "二维码缓存失败",
	UidNotExistError:       "uid不存在",
	QrCodeConvertJsonError: "二维码信息转换为json格式失败",
	TemplateParseError:     "模板解析失败",
	TemplateExecuteError:   "模板执行失败",
	LocalIpNotObtainError:  "本地IP获取失败",
	CacheNotSetError:       "缓存设置失败",
}
