package errcode

var (
	ErrNameHasExist       = NewErr(3001, "名字已经存在")
	ErrUpdateDataSame     = NewErr(3002, "更新的数据一致")
	ErrTimeConflict       = NewErr(3003, "时间冲突")
	ErrSendTooMany        = NewErr(3004, "请求过多")
	ErrRepeatOpt          = NewErr(3005, "重复操作")
	ErrSelfOpt            = NewErr(3006, "不能为自己操作")
	ErrOutTimeVerify      = NewErr(3007, "验证码错误或过期")
	ErrOutTimeInvite      = NewErr(3008, "邀请码不存在或者已过期")
	ErrFileOutSize        = NewErr(3009, "文件超过限制")
	ErrPlanNotExist       = NewErr(3010, "演出计划已过时或不存在")
	ErrSeatUnlock         = NewErr(3011, "未经锁定不允许购票")
	ErrCinemaHasPlans     = NewErr(3012, "影厅存在未过期演出计划")
	ErrMovieHasPlans      = NewErr(3013, "电影存在演出计划")
	ErrPlanHasSoldTickets = NewErr(3014, "演出计划存在已经售出或者锁定的票")
	ErrUserNotExist       = NewErr(3015, "用户不存在")
	ErrFileOpen           = NewErr(3016, "文件打开错误")
	ErrUUIDParse          = NewErr(3017, "uuid解析错误")
	ErrNameOrEmailExist   = NewErr(3018, "名字或邮箱已经存在")
	ErrLogin              = NewErr(3019, "用户名或密码错误")
)

var (
	ErrRedis = NewErr(4001, "redis错误")
)
