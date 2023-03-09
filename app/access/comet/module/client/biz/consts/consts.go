package consts

const (
	ChannelAttrUid                     = "uid"                     // 当前用户ID
	ChannelAttrProductId               = "productId"               // 当前应用ID
	ChannelAttrLastHeartBeatTime       = "lastHeartBeatTime"       // 最近一次客户端向comet服务发起心跳的时间
	ChannelAttrServerHeartBeatInterval = "serverHeartBeatInterval" // comet服务向后端服务发起心跳的间隔
	ChannelAttrLastServerHeartBeatTime = "lastServerHeartBeatTime" // 最近一次comet服务向后端服务发起心跳的时间

	ChannelAttrTimerAuthCheck      = "timer.authCheck"      // 定时器-检查是否成功授权
	ChannelAttrTimerHeartBeatCheck = "timer.heartBeatCheck" // 定时器-检查客户端发送心跳的间隔
)
