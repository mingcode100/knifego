package js


import (
	"gitee.com/youbeiwuhuan/knife4go/gin-swagger-knife/constant"
	"gitee.com/youbeiwuhuan/knife4go/gin-swagger-knife/utils"
	"github.com/gin-gonic/gin"

)


const (
	CHUNK_2D0D0B98_CB1DEA78_JS_RELATIVE_PATH = constant.ROOT_PATH + "/webjars/js/chunk-2d0d0b98.cb1dea78.js"
	// 文件内容的16进制表示
	CHUNK_2D0D0B98_CB1DEA78_JS_HEX_CONTENT = `2877696e646f772e7765627061636b4a736f6e703d77696e646f772e7765627061636b4a736f6e707c7c5b5d292e70757368285b5b226368756e6b2d3264306430623938225d2c7b2236386363223a66756e6374696f6e286e2c652c74297b2275736520737472696374223b742e722865293b76617220733d7428223065353422292c723d742e6e2873293b722e612e7365744f7074696f6e73287b67666d3a21302c7461626c65733a21302c627265616b733a21312c706564616e7469633a21312c73616e6974697a653a21312c736d6172744c697374733a21302c736d6172747970616e74733a21317d293b76617220613d7b6e616d653a224d61726b646f776e222c70726f70733a7b736f757263653a7b747970653a537472696e677d7d2c636f6d70757465643a7b6d61726b646f776e536f757263653a66756e6374696f6e28297b72657475726e2072282928746869732e736f75726365297d7d7d2c6f3d7428223238373722292c633d4f626a656374286f2e612928612c2866756e6374696f6e28297b766172206e3d746869732c653d6e2e24637265617465456c656d656e743b72657475726e286e2e5f73656c662e5f637c7c65292822646976222c7b737461746963436c6173733a226b6e696665346a2d6d61726b646f776e222c646f6d50726f70733a7b696e6e657248544d4c3a6e2e5f73286e2e6d61726b646f776e536f75726365297d7d297d292c5b5d2c21312c6e756c6c2c6e756c6c2c6e756c6c293b652e64656661756c743d632e6578706f7274737d7d5d293b`
)

func AddRouterOfChunk2d0d0b98Cb1dea78Js(router *gin.Engine) {
    
    utils.GetJs(router, CHUNK_2D0D0B98_CB1DEA78_JS_RELATIVE_PATH, CHUNK_2D0D0B98_CB1DEA78_JS_HEX_CONTENT)
    
}






