package js


import (
	"gitee.com/youbeiwuhuan/knife4go/gin-openapi3-knife/constant"
	"gitee.com/youbeiwuhuan/knife4go/gin-openapi3-knife/utils"
	"github.com/gin-gonic/gin"

)


const (
	CHUNK_260D712A_390177FE_JS_RELATIVE_PATH = constant.ROOT_PATH + "/webjars/js/chunk-260d712a.390177fe.js"
	// 文件内容的16进制表示
)

func AddRouterOfChunk260d712a390177feJs(router *gin.Engine) {
    
    utils.GetJs(router, CHUNK_260D712A_390177FE_JS_RELATIVE_PATH, CHUNK_260D712A_390177FE_JS_HEX_CONTENT)
    
}






