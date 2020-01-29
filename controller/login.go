package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gsxhnd/BangumiServer/common"
	"github.com/gsxhnd/BangumiServer/logger"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"time"
)

type accessTokenReq struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
	RedirectUri  string `json:"redirect_uri"`
	State        string `json:"state"`
}

type accessTokenRes struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    string `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
	UserId       string `json:"user_id"`
}

func GetCode(ctx *gin.Context) {
	// get temp code
	code := ctx.Query("code")
	logger.Debug("get code", zap.String("code", code))

	var reqBody = &accessTokenReq{
		GrantType:    "authorization_code",
		ClientId:     "bgm13805e31595be3d6d",
		ClientSecret: "ac8f2e08a60eee7c3f9ec7b9efec1f73",
		Code:         code,
		RedirectUri:  "http://192.168.4.238:8080/wait_login",
		State:        time.Now().String(),
	}
	var resBody accessTokenRes

	b, err := json.Marshal(reqBody)
	if err != nil {
	}
	req := bytes.NewBuffer(b)

	res, _ := http.Post("https://bgm.tv/oauth/access_token", "application/json;charset=utf-8", req)

	resB, _ := ioutil.ReadAll(res.Body)
	_ = json.Unmarshal(resB, &resBody)
	logger.Debug("", zap.Any("res", resBody))

	common.SendResponse(ctx, nil, resBody)
}

func WaitLogin(ctx *gin.Context) {
	common.SendResponse(ctx, nil, "waiting login")
}
