package wechat

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// LoadAccessTokenFromWeChatServer 从wechat服务器加载access token
// 微信公众平台拉取AccessToken
func LoadAccessTokenFromWeChatServer(wechatCode string) (*LoadAccessTokenReply, error) {
	urlGetAccessToken := fmt.Sprintf(urlWeChatGetAccessToken, weChatAPPID, weChatAPPSecret, wechatCode)
	log.Println("loadAccessTokenFromWeChatServer, full url:", urlGetAccessToken)

	reply := &LoadAccessTokenReply{}
	err := loadDataUseHTTPGet(urlGetAccessToken, reply)
	if err != nil {
		return nil, err
	}

	log.Printf("loadAccessTokenFromWeChatServer, reply json:%+v\n", reply)
	return reply, nil
}

// LoadNativeAccessTokenFromWeChatServer 微信开放平台拉取accessToken
func LoadNativeAccessTokenFromWeChatServer(wechatCode string) (*LoadAccessTokenReply, error) {
	urlGetAccessToken := fmt.Sprintf(urlWeChatNativeCodeGetAccessToken, nativeWeChatAPPID, nativeWeChatAPPSecret, wechatCode)
	log.Println("LoadNativeAccessTokenFromWeChatServer, full url:", urlGetAccessToken)

	reply := &LoadAccessTokenReply{}
	err := loadDataUseHTTPGet(urlGetAccessToken, reply)
	if err != nil {
		return nil, err
	}

	log.Printf("loadAccessTokenFromWeChatServer, reply json:%+v\n", reply)
	return reply, nil
}

// GetWeiXinUserInfoFromAccessToken 微信开放平台拉取用户信息
func GetWeiXinUserInfoFromAccessToken(accessToken string, openID string) (*LoadUserInfoReply, error) {
	urlGetUserInfo := fmt.Sprintf(urlWeChatGetUserInfo, accessToken, openID)
	log.Println("GetWeiXinUserInfoFromAccessToken, full url:", urlGetUserInfo)

	reply := &LoadUserInfoReply{}
	err := loadDataUseHTTPGet(urlGetUserInfo, reply)
	if err != nil {
		return nil, err
	}

	log.Printf("loadAccessTokenFromWeChatServer, reply json:%+v\n", reply)
	return reply, nil
}


// GetWeiXinPlusUserInfo 获取用户信息
func GetWeiXinPlusUserInfo(sessionkey string, encrypteddata string, iv string) (*WeiXinUserPlusInfo, error) {
	// log := common.NewLogger(0, nil)
	log.Printf("GetWeiXinPlusUserInfo, sessionkey:%s, encrypteddata:%s, iv:%s", sessionkey, encrypteddata, iv)

	skey, err := base64.StdEncoding.DecodeString(sessionkey)
	if err != nil {
		log.Error("GetWeiXinPlusUserInfo, decode sessionkey error:", err)
		return nil, err
	}

	sdata, err := base64.StdEncoding.DecodeString(encrypteddata)
	if err != nil {
		log.Error("GetWeiXinPlusUserInfo, decode encrypteddata error:", err)
		return nil, err
	}

	siv, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		log.Error("GetWeiXinPlusUserInfo, decode iv error:", err)
		return nil, err
	}

	databyte := pswDecrypt(string(sdata), string(skey), string(siv))

	userplusinfo := &WeiXinUserPlusInfo{}
	err = json.Unmarshal(databyte, &userplusinfo)
	if err != nil {
		log.Error("GetWeiXinPlusUserInfo, Unmarshal WeiXinUserPlusInfo error:", err)
		return nil, err
	}

	return userplusinfo, nil
}

func pswDecrypt(src string, skey string, siv string) []byte {
	key := []byte(skey)
	iv := []byte(siv)
	data := []byte(src)

	var err error

	origData, err := aes128Decrypt(data, key, iv)
	if err != nil {
		panic(err)
	}
	return origData
}

func aes128Decrypt(crypted, key []byte, IV []byte) ([]byte, error) {
	if key == nil || len(key) != 16 {
		return nil, nil
	}
	if IV != nil && len(IV) != 16 {
		return nil, nil
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, IV[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = pkcs5UnPadding(origData)
	return origData, nil
}

func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func loadDataUseHTTPGet(url string, jsonStruct interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	// 确保body关闭
	defer resp.Body.Close()

	// htmlData, err := ioutil.ReadAll(resp.Body) //<--- here!

	// log.Println("loadDataUseHTTPGet, body:", string(htmlData))
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(jsonStruct)
	return err

}
