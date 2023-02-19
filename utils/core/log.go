package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func CustomB5Log(param gin.LogFormatterParams) string {
	//过滤静态资源请求
	if param.Method == "GET" {
		if strings.Index(param.Path,"/static/") == 0 || strings.Index(param.Path,"/uploads/") == 0 {
			return ""
		}
	}
	return fmt.Sprintf("[b5gocmf] %s   |  %d   |   %s   |   %s   |   %s   |   %s \n---------------------------------------------------------------------------------------------------------------------------------------------------\n",
		param.TimeStamp.Format(G_TIME),
		param.StatusCode,
		param.Latency,
		param.ClientIP,
		param.Method,
		param.Path,
	)
}

