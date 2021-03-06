package lobby

var (
	//ErrorString 返回错误字符串
	ErrorString = map[int32]string{
		int32(MsgError_ErrSuccess):                                 "操作成功",
		int32(MsgError_ErrDecode):                                  "解码错误",
		int32(MsgError_ErrEncode):                                  "编码错误",
		int32(MsgError_ErrRoomExist):                               "房间已经存在",
		int32(MsgError_ErrNoRoomConfig):                            "没有房间配置",
		int32(MsgError_ErrServerIsFull):                            "服务器已经满",
		int32(MsgError_ErrDecodeRoomConfig):                        "解码房间配置错误",
		int32(MsgError_ErrRoomNotExist):                            "房间不存在",
		int32(MsgError_ErrDatabase):                                "数据库错误",
		int32(MsgError_ErrRequestGameServerTimeOut):                "请求游戏服务器超时",
		int32(MsgError_ErrWaitGameServerSN):                        "分配的序列号错误",
		int32(MsgError_ErrRoomIDIsEmpty):                           "房间ID为空",
		int32(MsgError_ErrNotRoomCreater):                          "你不是房间创建者",
		int32(MsgError_ErrGameIsPlaying):                           "游戏正在进行中",
		int32(MsgError_ErrTokenIsEmpty):                            "token为空",
		int32(MsgError_ErrUserIdIsEmpty):                           "用户ID为空",
		int32(MsgError_ErrRoomCountIsOutOfLimit):                   "房间数量已经达到上限",
		int32(MsgError_ErrRoomNumberNotExist):                      "你输入的房间号不存在，请确认",
		int32(MsgError_ErrGameServerIDNotExist):                    "游戏服务器ID不存在",
		int32(MsgError_ErrRoomNumberIsEmpty):                       "房间号为空",
		int32(MsgError_ErrRequestInvalidParam):                     "请求的参数无效",
		int32(MsgError_ErrTakeoffDiamondFailedNotEnough):           "剩余钻石不足",
		int32(MsgError_ErrTakeoffDiamondFailedIO):                  "数据库IO出错",
		int32(MsgError_ErrTakeoffDiamondFailedRepeat):              "已经扣取钻石",
		int32(MsgError_ErrGameServerUnsupportRoomType):             "游戏服务器不支持房间类型",
		int32(MsgError_ErrGameServerRoomExist):                     "游戏服务器已经存在这个房间",
		int32(MsgError_ErrGameServerNoRoomConfig):                  "游戏服务器没有房间配置",
		int32(MsgError_ErrGameServerDecodeRoomConfig):              "游戏服务器解码房间配置错误",
		int32(MsgError_ErrGameServerRoomNotExist):                  "游戏服务器不存在这个房间",
		int32(MsgError_ErrUserInOtherRoom):                         "用户正在别的房间",
		int32(MsgError_ErrRoomIsFull):                              "你输入的房间已满，无法加入",
		int32(MsgError_ErrUserInBlacklist):                         "你已经在黑名单中",
		int32(MsgError_ErrClubIDIsEmtpy):                           "牌友群ID为空",
		int32(MsgError_ErrRoomPriceCfgNotExist):                    "服务器房间价格配置不存在",
		int32(MsgError_ErrUserCreateRoomLock):                      "正在创建房间中，请稍等...",
		int32(MsgError_ErrGenerateRoomNumber):                      "生成房间号失败",
		int32(MsgError_ErrIsNeedUpdate):                            "需要更新",
		int32(MsgError_ErrOnlyClubCreatorOrManagerAllowCreateRoom): "只有群主或者管理员才允许创建房间",
		int32(MsgError_ErrOnlyClubCreatorOrManagerAllowDeleteRoom): "只有群主或者管理员才允许解散房间",
		int32(MsgError_ErrNotClubMember):                           "不是牌友圈成员",
	}

	// LoginString 登录回复
	LoginString = map[int32]string{
		int32(LoginState_Faild):           "登录失败",
		int32(LoginState_Success):         "登录成功",
		int32(LoginState_UserInBlacklist): "用户已经被封号",
		int32(LoginState_ParseTokenError): "解析token出错",
	}
)
