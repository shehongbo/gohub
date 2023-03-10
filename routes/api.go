package routes

import (
	"github.com/gin-gonic/gin"
	"gohub/app/http/controllers/api/v1/auth"
	"gohub/app/http/middlewares"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	v1 := r.Group("/v1")

	// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来。
	// 作为参考 Github API 每小时最多 60 个请求（根据 IP）。
	// 测试时，可以调高一点。
	v1.Use(middlewares.LimitIP("200-H"))

	{
		authGroup := v1.Group("/auth")

		// 限流中间件：每小时限流，作为参考 Github API 每小时最多 60 个请求（根据 IP）
		// 测试时，可以调高一点
		authGroup.Use(middlewares.LimitIP("1000-H"))

		{
			suc := new(auth.SignupController)
			//判断手机是否已注册
			authGroup.POST("/signup/phone/exist", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), suc.IsPhoneExist)
			//判断Email是否已注册
			authGroup.POST("/signup/email/exist", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), suc.IsEmailExist)
			//注册手机号
			authGroup.POST("/signup/using-phone", middlewares.GuestJWT(), suc.SignupUsingPhone)
			//注册邮箱
			authGroup.POST("/signup/using-email", middlewares.GuestJWT(), suc.SignupUsingEmail)

			// 发送验证码
			vcc := new(auth.VerifyCodeController)
			// 图片验证码，需要加限流
			authGroup.POST("/verify-code/captcha", middlewares.LimitPerRoute("50-H"), vcc.ShowCaptcha)
			//发送手机验证码
			authGroup.POST("/verify-code/phone", middlewares.LimitPerRoute("20-H"), vcc.SendUsingPhone)
			//发送邮件验证码
			authGroup.POST("/verify-code/email", middlewares.LimitPerRoute("20-H"), vcc.SendUsingEmail)

			lgc := new(auth.LoginController)
			//使用手机号，短信验证码进行登录
			authGroup.POST("/login/using-phone", middlewares.GuestJWT(), lgc.LoginByPhone)
			// 支持手机号，Email 和 用户名
			authGroup.POST("/login/using-password", middlewares.GuestJWT(), lgc.LoginByPassword)
			// 刷新token
			authGroup.POST("/login/refresh-token", middlewares.AuthJWT(), lgc.RefreshToken)

			//重置密码
			pwc := new(auth.PasswordController)
			authGroup.POST("/password-reset/using-phone", middlewares.GuestJWT(), pwc.ResetByPhone)
			authGroup.POST("/password-reset/using-email", middlewares.GuestJWT(), pwc.ResetByEmail)
		}
	}
}
