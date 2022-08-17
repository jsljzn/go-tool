package middlewarex

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/rest/token"
	"net/http"
)

const (
	jwtAudience    = "aud"
	jwtExpire      = "exp"
	jwtId          = "jti"
	jwtIssueAt     = "iat"
	jwtIssuer      = "iss"
	jwtNotBefore   = "nbf"
	jwtSubject     = "sub"
	noDetailReason = "no detail reason"
)

type AuthMiddleware struct {
	Secret string `json:"secret"`
}

func NewAuthMiddleware(secret string) *AuthMiddleware {
	return &AuthMiddleware{Secret: secret}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		logx.Info("局部中间件前")
		// Passthrough to next handler if need
		parser := token.NewTokenParser()
		tok, err := parser.ParseToken(r, m.Secret, "")

		if err != nil {
			logx.Errorf("[jwt]验证异常:%s",err.Error())
			httpx.OkJson(w, resx.ResFail("请登录"))
			return
		}
		if !tok.Valid {
			logx.Errorf("[jwt]验证异常:%s",err.Error())
			httpx.OkJson(w, resx.ResFail("请登录"))
			return
		}
		claims, ok := tok.Claims.(jwt.MapClaims)
		if !ok {
			logx.Errorf("[jwt]验证异常:%s",err.Error())
			httpx.OkJson(w, resx.ResFail("请登录"))
			return
		}
		ctx := r.Context()
		for k, v := range claims {
			switch k {
			case jwtAudience, jwtExpire, jwtId, jwtIssueAt, jwtIssuer, jwtNotBefore, jwtSubject:
				// ignore the standard claims
			default:
				ctx = context.WithValue(ctx, k, v)
			}
		}
		next.ServeHTTP(w, r.WithContext(ctx))
		logx.Info("局部中间件后")
	}
}
