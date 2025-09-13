package middle

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"teaching_evaluation_backend/biz/model/base"
	eva "teaching_evaluation_backend/biz/model/teaching_evaluation"
	"teaching_evaluation_backend/consts"
	"teaching_evaluation_backend/utils"
	"time"
)

var JwtKey = []byte("hdu_itmo_teaching_evaluation")

type Claims struct {
	Username string       `json:"username"`
	ID       int64        `json:"id"`
	Role     eva.UserRole `json:"role"`
	CreateAt int64        `json:"create_at"`
	jwt.RegisteredClaims
}

var (
	NotNeedAuthPathSuffix = []string{"/login", "/ping"}
	AdminPrefix           = "/admin/"
)

func JWTAuthMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		path := string(c.Path())

		lastPathStart := strings.LastIndex(path, "/")
		pathEnd := path[lastPathStart:]
		// 白名单（不需要鉴权的接口）
		method := strings.ToUpper(string(c.Method()))
		if utils.Contains(NotNeedAuthPathSuffix, pathEnd) || method == "OPTIONS" {
			c.Next(ctx)
			return
		}

		authHeader := c.GetHeader("Authorization")
		if len(authHeader) == 0 {
			c.AbortWithStatusJSON(http.StatusOK, &base.BaseResp{
				StatusCode:    consts.NoTokenErrorCode,
				StatusMessage: "missing token",
			})
			return
		}

		claims, err := ParseToken(string(authHeader))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, &base.BaseResp{
				StatusCode:    consts.InvalidTokenErrorCode,
				StatusMessage: fmt.Sprintf("invalid token: %s", err.Error()),
			})
			return
		}

		if strings.Contains(path, AdminPrefix) && claims.Role != eva.UserRole_Admin {
			c.AbortWithStatusJSON(http.StatusOK, &base.BaseResp{
				StatusCode:    consts.NoAuthErrorCode,
				StatusMessage: "no auth, need admin role",
			})
			return
		}

		ctx = utils.SetCurrentUserInfo(ctx, eva.UserInfo{
			ID:       claims.ID,
			Name:     claims.Username,
			Role:     claims.Role,
			CreateAt: claims.CreateAt,
		})

		c.Next(ctx)
	}
}

func GenerateToken(expireTime time.Time, userInfo *eva.UserInfo) (string, error) {
	claims := &Claims{
		Username: userInfo.Name,
		ID:       userInfo.ID,
		Role:     userInfo.Role,
		CreateAt: userInfo.CreateAt,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "teaching_evaluation",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", fmt.Errorf("token signing error: %s", err.Error())
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
