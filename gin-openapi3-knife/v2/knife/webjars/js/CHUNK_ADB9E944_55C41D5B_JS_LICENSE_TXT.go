package js


import (
	"gitee.com/youbeiwuhuan/knife4go/gin-swagger-knife/constant"
	"gitee.com/youbeiwuhuan/knife4go/gin-swagger-knife/utils"
	"github.com/gin-gonic/gin"

)


const (
	CHUNK_ADB9E944_55C41D5B_JS_LICENSE_TXT_RELATIVE_PATH = constant.ROOT_PATH + "/webjars/js/chunk-adb9e944.55c41d5b.js.LICENSE.txt"
	// 文件内容的16进制表示
	CHUNK_ADB9E944_55C41D5B_JS_LICENSE_TXT_HEX_CONTENT = `2f2a210d0a202a20636c6970626f6172642e6a732076322e302e360d0a202a2068747470733a2f2f636c6970626f6172646a732e636f6d2f0d0a202a200d0a202a204c6963656e736564204d495420c2a9205a656e6f20526f6368610d0a202a2f0d0a`
)

func AddRouterOfChunkAdb9e94455c41d5bJsLICENSETxt(router *gin.Engine) {
    
	utils.GetOther(router, CHUNK_ADB9E944_55C41D5B_JS_LICENSE_TXT_RELATIVE_PATH, CHUNK_ADB9E944_55C41D5B_JS_LICENSE_TXT_HEX_CONTENT)
	
}






