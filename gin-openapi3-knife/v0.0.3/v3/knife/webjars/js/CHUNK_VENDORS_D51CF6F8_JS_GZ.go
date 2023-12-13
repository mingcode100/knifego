package js


import (
	"gitee.com/youbeiwuhuan/knife4go/gin-openapi3-knife/constant"
	"gitee.com/youbeiwuhuan/knife4go/gin-openapi3-knife/utils"
	"github.com/gin-gonic/gin"

)


const (
	CHUNK_VENDORS_D51CF6F8_JS_GZ_RELATIVE_PATH = constant.ROOT_PATH + "/webjars/js/chunk-vendors.d51cf6f8.js.gz"
	// 文件内容的16进制表示
)

func AddRouterOfChunkVendorsD51cf6f8JsGz(router *gin.Engine) {
    
	utils.GetOther(router, CHUNK_VENDORS_D51CF6F8_JS_GZ_RELATIVE_PATH, CHUNK_VENDORS_D51CF6F8_JS_GZ_HEX_CONTENT)
	
}






