package middlewares

// JWT is jwt middleware
//func JWT() gin.HandlerFunc {
	//return func(c *gin.Context) {
	//	var code int
	//	var data interface{}
	//
	//	code = response.SUCCESS
	//	token := c.Query("token")
	//	if token == "" {
	//		code = response.INVALID_PAYLOAD
	//	} else {
	//		_, err := util.ParseToken(token)
	//		if err != nil {
	//			switch err.(*jwt.ValidationError).Errors {
	//			case jwt.ValidationErrorExpired:
	//				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
	//			default:
	//				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
	//			}
	//		}
	//	}
	//
	//	if code != e.SUCCESS {
	//		c.JSON(http.StatusUnauthorized, gin.H{
	//			"code": code,
	//			"msg":  e.GetMsg(code),
	//			"data": data,
	//		})
	//
	//		c.Abort()
	//		return
	//	}
	//
	//	c.Next()
	//}
//}
