package middlewares

import (
	"blog/api/helpers"
	"blog/api/helpers/auth_helper"
	"blog/api/helpers/common"
	"blog/pkg/auth_manager"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserAuthMiddleware struct {
	AuthManager auth_manager.AuthManager
	AuthHelper auth_helper.AuthHeaderHelper
}

var (
	ErrYouAreUnAuthorized = errors.New("you are unauthorized")
)
func NewUserAuthMiddelware(authManger auth_manager.AuthManager,authHelper auth_helper.AuthHeaderHelper)*UserAuthMiddleware{
	return &UserAuthMiddleware{
		AuthManager:authManger,
		AuthHelper:authHelper,
	}
}
func (m *UserAuthMiddleware)SetUserStatus() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorizetion")
		if token == "" {
			ctx.Set("is_logged", false)
			ctx.Next()
		} else {
			cliams, err := m.AuthManager.DecodeToken(token,auth_manager.AccessToken)
			if err != nil{
				ctx.AbortWithStatusJSON(http.StatusInternalServerError,
					helpers.NewHttpResponse(
						http.StatusInternalServerError, err.Error(), nil))
				return
			}
			ctx.Set("id",cliams.ID)
			ctx.Set("role",cliams.Role)
			ctx.Set("is_logged", true)
			ctx.Set("is_admin", common.IsAdmin(cliams.Role))
			ctx.Next()
		}
	}
}
func (m *UserAuthMiddleware)EnsureLoggedIn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		is_logged := common.GetUserStatus(ctx)
		if is_logged {
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				helpers.NewHttpResponse(
					http.StatusUnauthorized, ErrYouAreUnAuthorized.Error(), nil))
			return
		}
	}
}
func (m *UserAuthMiddleware)EnsureNotLoggedIn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		is_logged := common.GetUserStatus(ctx)
		if !is_logged {
			ctx.Next()
		} else {
			token,_ := m.AuthHelper.GetHeader(ctx)
			m.AuthHelper.DeleteHeader(ctx)
			m.AuthManager.Destroy(token)
			ctx.Next()
		}
	}
}
func (m *UserAuthMiddleware)EnsureAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.GetBool("is_admin"){
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				helpers.NewHttpResponse(
					http.StatusUnauthorized, ErrYouAreUnAuthorized.Error(), nil))
			return
		} else {
			ctx.Next()
		}
	}
}
