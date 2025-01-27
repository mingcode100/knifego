package css

import (
	"github.com/gin-gonic/gin"
	"github.com/mingcode100/knife4go/gin-swagger-knife/constant"
	"github.com/mingcode100/knife4go/gin-swagger-knife/utils"
)

const (
	CHUNK_51277DBE_57225F85_CSS_RELATIVE_PATH = constant.ROOT_PATH + "/webjars/css/chunk-51277dbe.57225f85.css"
	// 文件内容的16进制表示
	CHUNK_51277DBE_57225F85_CSS_HEX_CONTENT = `.api-tab[data-v-7da2228c]{margin-top:15px}.api-tab .ant-tag[data-v-7da2228c]{height:32px;line-height:32px}.api-basic[data-v-7da2228c]{padding:11px}.api-basic-title[data-v-7da2228c]{font-size:14px;font-weight:700}.api-basic-body[data-v-7da2228c]{font-size:14px;font-family:-webkit-body}.api-description[data-v-7da2228c]{border-left:4px solid #ddd;line-height:30px}.api-body-desc[data-v-7da2228c]{padding:10px;min-height:35px;-webkit-box-sizing:border-box;box-sizing:border-box;border:1px solid #e8e8e8}.ant-card-body[data-v-7da2228c]{padding:5px}.api-title[data-v-7da2228c]{margin-top:10px;margin-bottom:5px;font-size:16px;font-weight:600;height:30px;line-height:30px;border-left:4px solid #00ab6d;text-indent:8px}.content-line[data-v-7da2228c]{height:25px;line-height:25px}.content-line-count[data-v-7da2228c]{height:35px;line-height:35px}.divider[data-v-7da2228c]{margin:4px 0}`
)

func AddRouterOfChunk51277dbe57225f85Css(router *gin.Engine) {

	utils.GetCss(router, CHUNK_51277DBE_57225F85_CSS_RELATIVE_PATH, CHUNK_51277DBE_57225F85_CSS_HEX_CONTENT)

}
