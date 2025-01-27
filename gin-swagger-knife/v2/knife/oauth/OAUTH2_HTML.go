package oauth

import (
	"github.com/gin-gonic/gin"
	"github.com/mingcode100/knife4go/gin-swagger-knife/constant"
	"github.com/mingcode100/knife4go/gin-swagger-knife/utils"
)

const (
	OAUTH2_HTML_RELATIVE_PATH = constant.ROOT_PATH + "/oauth/oauth2.html"
	// 文件内容的16进制表示
	OAUTH2_HTML_HEX_CONTENT = `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.0 Transitional//EN">
<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1"/>
    <title>Knife4j-OAuth2</title>
    <script src="jquery.min.js"></script>
</head>
<body>
<script type="text/javascript">
$(function(){
    function OAuth2(url){
        this.url=url;
        this.code=null;
        this.accessToken=null;
        this.tokenType=null;
        this.state=null;
        //缓存在localStorage中的对象
        this.cacheValue=null;
    }
    OAuth2.prototype.init=function(){
        var local=this.url;
        this.code=this.getKey("code",local,"");
        this.accessToken=this.getKey("access_token",local,"");
        this.tokenType=this.getKey("token_type",local,"Bearer");
        this.state=this.getKey("state",local);
        if(window.localStorage){
            var value=window.localStorage.getItem(this.state);
            if(this.strNotBlank(value)){
                this.cacheValue=JSON.parse(value);
            }
        }
    }
    OAuth2.prototype.auth=function(){
        if(this.strNotBlank(this.code)){
            this.authorizationCode();
        }else{
            this.implicit();
        }
    }
    OAuth2.prototype.getKey=function(key,str,defaultValue){
        var reg=new RegExp(".*?"+key+"=(.*?)(&.*)?$","ig");
        var val=defaultValue;
        if(reg.test(str)){
            val=RegExp.$1;
        }
        return val;
    }
    OAuth2.prototype.strNotBlank=function(str){
        var flag = false;
        if ( str != undefined &&str != null &&str != "") {
            flag = true;
        }
        return flag;
    }

    OAuth2.prototype.implicit=function(){
        this.cacheValue.accessToken=this.tokenType+" "+this.accessToken;
        this.cacheValue.tokenType=this.tokenType;
        this.cacheValue.granted=true;
        window.localStorage.setItem(this.state,JSON.stringify(this.cacheValue))
        window.close();
    }
    OAuth2.prototype.authorizationCode=function(){
        var that=this;
        console.log(this.cacheValue);
        var url=this.cacheValue.tokenUrl;
        var params={
            "grant_type":"authorization_code",
            "code":this.code,
            "redirect_uri":decodeURIComponent(this.cacheValue.redirectUri),
            "client_id":this.cacheValue.clientId,
            "client_secret":this.cacheValue.clientSecret
        }
        $.post(url,params,function(data){
            if(data!=null&&data!=undefined) {
                that.cacheValue.accessToken=data.token_type+" "+data.access_token;
                that.cacheValue.tokenType=data.token_type;
                that.cacheValue.granted=true;
                window.localStorage.setItem(that.state,JSON.stringify(that.cacheValue))
                window.close();
            }
        })
    }
    var oauth=new OAuth2(window.location.href);
    oauth.init();
    oauth.auth();
})
</script>

</body>
</html>`
)

func AddRouterOfOauth2Html(router *gin.Engine) {

	utils.GetHtml(router, OAUTH2_HTML_RELATIVE_PATH, OAUTH2_HTML_HEX_CONTENT)

}
