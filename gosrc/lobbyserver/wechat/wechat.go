package wechat

import (
	"gconst"
	"lobbyserver/lobby"

	"github.com/gomodule/redigo/redis"
	log "github.com/sirupsen/logrus"
)

var (
	weChatAPPID     = ""
	weChatAPPSecret = ""
	// 微信开放平台appid
	nativeWeChatAPPID = ""
	// 微信开放平台appSecret = ""
	nativeWeChatAPPSecret = ""
	// 微信开放平台，获取access_token
	urlWeChatNativeCodeGetAccessToken = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
	// urlWeChatGetAccessToken = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
	// 微信公众平台，获取access_token
	urlWeChatGetAccessToken = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	// 微信开放平台，获取用户信息
	urlWeChatGetUserInfo    = "https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN"
	// urlWeChatGetJSTicket    = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi"
	// urlWeChatRefreshToken   = "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s"
	// urlWeChatSysAccessToken = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
)

// MyWeChatConfig 前面两个是微信公众平台，后两个是微信开放平台
type MyWeChatConfig struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`

	// 微信开放平台的appID和appSecret
	NativeAppID string `json:"native_app_id"`
	NativeAppSecret string `json:"native_app_secret"`
}

func loadWeChatConfig() *MyWeChatConfig {
	conn := lobby.Pool().Get()
	defer conn.Close()

	fields, err := redis.Strings(conn.Do("HMGET", gconst.LobbyWeChatConfig, "appID", "appAcret", "nativeAppID", "nativeAppSecret"))
	if err != nil {
		log.Println("loadUserInfoFromRedis, error", err)
		return nil
	}

	wechatCfg := &MyWeChatConfig{}
	wechatCfg.AppID = fields[0]
	wechatCfg.AppSecret = fields[1]
	wechatCfg.NativeAppID = fields[2]
	wechatCfg.NativeAppSecret = fields[3]

	return wechatCfg
}

// InitWechat 初始化wechat
func InitWechat() {
	weChatConfig := loadWeChatConfig()

	weChatAPPID = weChatConfig.AppID
	weChatAPPSecret = weChatConfig.AppSecret

	nativeWeChatAPPID = weChatConfig.NativeAppID
	nativeWeChatAPPSecret = weChatConfig.NativeAppSecret

	if weChatAPPID == "" {
		log.Error("wechat.InitWechat, weChatAPPID is empty, please config hset l:wc:cfg appID xxx")
	}

	if weChatAPPSecret == "" {
		log.Error("wechat.InitWechat, weChatAPPSecret is empty, please config hset l:wc:cfg appAcret xxx")
	}

	if nativeWeChatAPPID == "" {
		log.Error("wechat.InitWechat, nativeWeChatAPPID is empty, please config hset l:wc:cfg nativeAppID xxx")
	}

	if nativeWeChatAPPSecret == "" {
		log.Error("wechat.InitWechat, nativeWeChatAPPSecret is empty, please config hset l:wc:cfg nativeAppSecret xxx")
	}
}

// GetAppID 获得APP ID
func GetAppID() string {
	return weChatAPPID
}
