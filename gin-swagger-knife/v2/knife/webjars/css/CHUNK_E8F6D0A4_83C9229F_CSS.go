package css


import (
	"gitee.com/youbeiwuhuan/knife4go/gin-swagger-knife/constant"
	"gitee.com/youbeiwuhuan/knife4go/gin-swagger-knife/utils"
	"github.com/gin-gonic/gin"

)


const (
	CHUNK_E8F6D0A4_83C9229F_CSS_RELATIVE_PATH = constant.ROOT_PATH + "/webjars/css/chunk-e8f6d0a4.83c9229f.css"
	// 文件内容的16进制表示
	CHUNK_E8F6D0A4_83C9229F_CSS_HEX_CONTENT = `.api-tab[data-v-2e1642ac]{margin-top:15px}.api-tab .ant-tag[data-v-2e1642ac]{height:32px;line-height:32px}.api-basic[data-v-2e1642ac]{padding:11px}.api-basic-title[data-v-2e1642ac]{font-size:14px;font-weight:700}.api-basic-body[data-v-2e1642ac]{font-size:14px;font-family:-webkit-body}.api-description[data-v-2e1642ac]{border-left:4px solid #ddd;line-height:30px}.api-body-desc[data-v-2e1642ac]{padding:10px;min-height:35px;-webkit-box-sizing:border-box;box-sizing:border-box;border:1px solid #e8e8e8}.ant-card-body[data-v-2e1642ac]{padding:5px}.api-title[data-v-2e1642ac]{margin-top:10px;margin-bottom:5px;font-size:16px;font-weight:600;height:30px;line-height:30px;border-left:4px solid #00ab6d;text-indent:8px}.content-line[data-v-2e1642ac]{height:25px;line-height:25px}.content-line-count[data-v-2e1642ac]{height:35px;line-height:35px}.divider[data-v-2e1642ac]{margin:4px 0}`
)

func AddRouterOfChunkE8f6d0a483c9229fCss(router *gin.Engine) {
    
    utils.GetCss(router, CHUNK_E8F6D0A4_83C9229F_CSS_HEX_CONTENT, CHUNK_E8F6D0A4_83C9229F_CSS_RELATIVE_PATH)
	
}






