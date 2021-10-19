package knife


import (
	"gitee.com/youbeiwuhuan/knife4go/gin-swagger-knife/constant"
	"gitee.com/youbeiwuhuan/knife4go/gin-swagger-knife/utils"
	"github.com/gin-gonic/gin"

)


const (
	FAVICON_ICO_RELATIVE_PATH = constant.ROOT_PATH + "/favicon.ico"
	// 文件内容的16进制表示
	FAVICON_ICO_HEX_CONTENT = `0000010001002020000001002000a81000001600000028000000200000004000000001002000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c0cda902c0ddcf28c0e5df48c0e2d84ec0e3db3ebfd7c0120000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c1e9eb00c0e8e626c0e8e792c0e8e8e8c1eaebffc1eaebffc1eaebffc1eaebffc0e9eafcc0e7e5c6c0e3dd5ebec49b0400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c0eaeb0ec1eaeb92c1e9ebf8c1eaebffc1eaebffc1eaebffc0e4dfffbfd4bcffbecba7ffbecba9ffbfd8c3ffc0e7e5ffc0e8e6cec0dfd32e0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c1eaeb26c1eaebd8c1eaebffc1eaebffc0e9eaffbec8a1ffbb9c40ffb98207ffb88000ffb98000ffb98000ffb88000ffb9850cffbba14bffbecba8f0c0e4d748000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c1e9ea26c0eaebe6c1eaebffc1e9eaffbfdccdffba9838ffb88000ffb98100ffb98100ffb98100ffb98100ffb98100ffb98100ffb98100ffb98100ffb88000ffba9837f4bca6583a00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c0eaea10c0e9ead8c1eaebffc1e9eaffbfd3b9ffb98611ffb88002ffb88002ffb88002ffb88001ffb98100ffb98100ffb88001ffb88002ffb88002ffb88002ffb88001ffb98100ffb98000e0b980000e0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000bfeaea00c0eaea92c1eaebffc1eaebffbfdbcbffb98712ffa8750fff281c77ff0b088aff0e0989ff6a4a42ffb98100ffb98100ff885f2bff110c87ff0b088aff160f84ff825b30ffb98100ffb87f00ffbcaa5e8e0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c0e9ea22c0e9eaf8c1eaebffc1eaebffba9837ffb98100ffb98100ff97691cff000091ff3a2967ffb98100ffb98100ffb98100ffaf7a09ff0a078aff03028fff926622ffb98100ffb98307ffbcb170ffbfd5bef4bfdecf3200000000000000000000000000000000000000000000000000000000000000000000000000000000c0eaea88c1eaebffc1eaebffbecba7ffb88000ffb98100ffb98100ffa6740fff000092ff49335affb98100ffb98100ffb98100ff5b404dff000091ff573c51ffb98100ffb98100ffb98000ffb98610ffbcb06dffb8800272000000000000000000000000000000000000000000000000000000000000000000000000c2ebed00c1eaead8c1eaebffc1eaebffbba14cffb98100ffb98100ffb98100ffa8750eff000093ff4c3558ffb98100ffb98100ffa07015ff06048dff171081ffb27c07ffb98100ffb98100ffb98100ffb88000ffb9860dffb98000c0000000000000000000000000000000000000000000000000000000000000000000000000c1e8e914c0eaebffc1eaebffc0e8e8ffb9860fffb98100ffb98100ffb98100ffa8750eff000093ff4c3558ffb98100ffb98100ff382768ff000091ff825a2effb88000ffb98102ffba9734ffbdb578ffbec8a2ffbecfb2ffbec69ff4c0e8e84cc0eaea0000000000000000000000000000000000000000000000000000000000c1eaeb3ac1eaebffc1eaebffc0ddcfffb88000ffb98100ffb98100ffb98100ffa8750eff000093ff4c3558ffb88100ff845c2dff000091ff3b2966ffb98205ffbba758ffbfd7c3ffc1eaebffc1eaeaffc1eaebffc1eaebffc1eaebffc1eaebffc1e9eb7c00000000000000000000000000000000000000000000000000000000c1eaeb46c1eaebffc1eaebffbfd6c2ffb88000ffb98100ffb98100ffb98100ffa8750eff000093ff4a345affa27115ff150f82ff08058bffa58949ffbfd6c1ffc1eaebffc1eaebffc1eaebffc1eaecffc1eaecffc1eaeaffc1eaebffc1eaebffc0e9eaf8c0ebeb16000000000000000000000000000000000000000000000000c0eaeb46c1eaebffc1eaebffc0e0d5ffb98102ffb98100ffb98100ffb98100ffa8750eff000093ff120d84ff130d85ff050692ff7989b2ffc0e9e9ffc1eaebffc0e9eaffbec8a2ffbb9d42ffb9840affb9860fffbdbf8fffc1eaebffc1eaebffc1eaebffc1eaeb42000000000000000000000000000000000000000000000000c1eaeb2ac1eaebffc1eaebffc1eaebffbba24effb98100ffb98100ffb98100ffa8750eff000093ff4c3558ffb69642ffa4c4d8ffc1eaebffc1eaebffbfd7c3ffba9736ffb88000ffb98100ffb98100ffb98100ffb98206ffc0e0d7ffc1eaebffc1eaebffc0eaeb44000000000000000000000000000000000000000000000000bfe9ea02c1eaebe4c1eaebffc1eaebffc0e4dfffbba658ffb98207ffb87f00ffa97917ff3a46adff96b1caffc1e9eaffc1eaebffbfe7eaff99a09effb98307ffb98100ffb98100ffb98100ffb98100ffb98100ffb88000ffbfd3bbffc1eaebffc0e9eaffc1eaeb1c00000000000000000000000000000000000000000000000000000000c0e9eb72c1eaebffc1eaebffc1eaebffc1eaebffc0e6e2ffc0e1d7ffc0e7e5ffc1eaebffc1eaebffc1eaebffbfd7c2ffa58949ff100b87ff906523ffb88000ffb98100ffb98100ffb98100ffb98100ffb98a18ffc0e7e4ffc1eaebffc1e9ebb80000000000000000000000000000000000000000000000000000000000000000beceb104becaa6dec1eaeaffc1eaebffc1eaebffc1eaebffc1eaebffc1eaebffc1eaebffb3cdc9ffbba657ffb98102ffb88100ff755139ff130d84ffa8750fffb98100ffba8f22ffba9228ffbb9a3dffbfd5bfffc1eaeaffc0e9ebe2c0e9ea20000000000000000000000000000000000000000000000000000000000000000000000000b880007aba993affbdbd8affbecdacffbecfb0ffbec49affa7a17bff21289fff412e61ffb98100ffb98100ffb98000ffb67f03ff0d0988ff291c74ffb37d05ffbdc192ffc1eaecffc1eaebffc0eaebfac1e9eaa6c1eaeb1c00000000000000000000000000000000000000000000000000000000000000000000000000000000b9810012b88000e8b98100ffb98000ffb98000ffad780bff37266cff000093ff0a078dff7d5734ffb98205ffb98c1bff7e5833ff030291ff000093ff524f87ffc1e7e4ffbfd5bfffbba759ffba95328e00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000b8810060b88100ffb98100ffb98100ffb57e04ffa17016ffa17016ffa17016ffa8750fffb98000ffbdbc87ffaca57effa49366ffa5aa98ffbbe2e7ffc1eaeaffbb9f47ffb98100e0b981000e00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000bb800000b880009ab98100ffb98100ffb98100ffb98100ffb98100ffb98100ffb98100ffb98100ffb98812ffbfd2b8ffc1eaebffc1eaebffc0e6e2ffbcad66ffb88000f2b8810038000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000b8800004b98100a0b88000ffb98100ffb98100ffb98100ffb98100ffb98100ffb98100ffb98100ffb98102ffba9632ffbb9c3effb98610ffb88000f0b88000440000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c3940002b880006eb98100f2b98100ffb98100ffb98100ffb98100ffb98100ffb98100ffb98100ffb98100ffb98100ffb88000cab980002a0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000b881001cb981008ab88100e2b98100ffb98100ffb98100ffb98100ffb98100fab88000c0b880005ab88200040000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000b9810002b9800022b9810044b880004cb981003ab98000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffffffffffffffffffffffffe01fffff8007ffff0003fffe0001fffc0000fff800007ff800007ff000007ff000003ff000003ff000001ff000000ff000000ff000000ff000000ff800000ff800001ffc00003ffc00007ffe0000fffe0001ffff0003ffffc007ffffe01fffffffffffffffffffffffffffffffffff`
)

func AddRouterOfFaviconIco(router *gin.Engine) {
    
	utils.GetOther(router, FAVICON_ICO_HEX_CONTENT, FAVICON_ICO_RELATIVE_PATH)
	
}






