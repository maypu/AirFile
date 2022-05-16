webpackJsonp([1],{"75l9":function(t,e){t.exports={_args:[["axios@0.21.3","C:\\Users\\Administrator\\Documents\\Projects\\VSCode\\AirFile_UI"]],_from:"axios@0.21.3",_id:"axios@0.21.3",_inBundle:!1,_integrity:"sha1-+F2bdH+bZtWcpGNgXO3xhEhyuC4=",_location:"/axios",_phantomChildren:{},_requested:{type:"version",registry:!0,raw:"axios@0.21.3",name:"axios",escapedName:"axios",rawSpec:"0.21.3",saveSpec:null,fetchSpec:"0.21.3"},_requiredBy:["/"],_resolved:"https://registry.nlark.com/axios/download/axios-0.21.3.tgz?cache=0&sync_timestamp=1630782409101&other_urls=https%3A%2F%2Fregistry.nlark.com%2Faxios%2Fdownload%2Faxios-0.21.3.tgz",_spec:"0.21.3",_where:"C:\\Users\\Administrator\\Documents\\Projects\\VSCode\\AirFile_UI",author:{name:"Matt Zabriskie"},browser:{"./lib/adapters/http.js":"./lib/adapters/xhr.js"},bugs:{url:"https://github.com/axios/axios/issues"},bundlesize:[{path:"./dist/axios.min.js",threshold:"5kB"}],dependencies:{"follow-redirects":"^1.14.0"},description:"Promise based HTTP client for the browser and node.js",devDependencies:{coveralls:"^3.0.0","es6-promise":"^4.2.4",grunt:"^1.3.0","grunt-banner":"^0.6.0","grunt-cli":"^1.2.0","grunt-contrib-clean":"^1.1.0","grunt-contrib-watch":"^1.0.0","grunt-eslint":"^23.0.0","grunt-karma":"^4.0.0","grunt-mocha-test":"^0.13.3","grunt-ts":"^6.0.0-beta.19","grunt-webpack":"^4.0.2","istanbul-instrumenter-loader":"^1.0.0","jasmine-core":"^2.4.1",karma:"^6.3.2","karma-chrome-launcher":"^3.1.0","karma-firefox-launcher":"^2.1.0","karma-jasmine":"^1.1.1","karma-jasmine-ajax":"^0.1.13","karma-safari-launcher":"^1.0.0","karma-sauce-launcher":"^4.3.6","karma-sinon":"^1.0.5","karma-sourcemap-loader":"^0.3.8","karma-webpack":"^4.0.2","load-grunt-tasks":"^3.5.2",minimist:"^1.2.0",mocha:"^8.2.1",sinon:"^4.5.0","terser-webpack-plugin":"^4.2.3",typescript:"^4.0.5","url-search-params":"^0.10.0",webpack:"^4.44.2","webpack-dev-server":"^3.11.0"},homepage:"https://axios-http.com",jsdelivr:"dist/axios.min.js",keywords:["xhr","http","ajax","promise","node"],license:"MIT",main:"index.js",name:"axios",repository:{type:"git",url:"git+https://github.com/axios/axios.git"},scripts:{build:"NODE_ENV=production grunt build",coveralls:"cat coverage/lcov.info | ./node_modules/coveralls/bin/coveralls.js",examples:"node ./examples/server.js",fix:"eslint --fix lib/**/*.js",postversion:"git push && git push --tags",preversion:"npm test",start:"node ./sandbox/server.js",test:"grunt test",version:"npm run build && grunt version && git add -A dist && git add CHANGELOG.md bower.json package.json"},typings:"./index.d.ts",unpkg:"dist/axios.min.js",version:"0.21.3"}},"7QVd":function(t,e){},NHnr:function(t,e,o){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var s=o("7+uW"),i={render:function(){var t=this.$createElement,e=this._self._c||t;return e("div",{attrs:{id:"app"}},[e("router-view")],1)},staticRenderFns:[]};var a=o("VU/8")({name:"App"},i,!1,function(t){o("Y9+7")},null,null).exports,n=o("/ocq"),r={name:"Home",data:function(){return{productName:"AirFile",containHeight:(window.innerHeight||document.documentElement.clientHeight)-61+"px",topHeight:((window.innerHeight||document.documentElement.clientHeight)-580)/2+"px",uuid:"",verifiCode:"",password:"",frequency:2,frequencyMax:10,limithours:24,limithoursMax:168,file:null,uploadMaxSize:102400,uploadText:null,progressNum:0,loadingStatus:!1,showMoreUpload:!1,showCopy:!1,showDownload:!1,showPwdModel:!1,disableLimitTip:!0,limitTipContent:"",isCopyed:!1,historyCol:[{type:"index",width:60,align:"center"},{title:"文件名",key:"FileName",width:250},{title:"文件码",key:"RandomCode",width:100},{title:"可下载次数",key:"LimitTimes",width:120},{title:"已下载次数",key:"NumDownloads",width:120},{title:"过期时间",key:"ExpiryTime"}],history:[],isShowHistory:!1,isShowAbout:!1}},watch:{limithours:function(t,e){this.disableLimitTip=!(t>24),this.limitTipContent=(Math.round(t/24)>0?(t/24).toFixed(2):0)+"天"}},computed:{nightTime:function(){return(this.utils.CompareTime(new Date,"23:00:00")>0||this.utils.CompareTime(new Date,"07:00:00")<0)&&(document.body.style.backgroundColor="#22303f",!0)}},created:function(){var t=this,e=this.$route.path,o=this.$route.params.fileCode;0==e.indexOf("/fileCode")&&o&&(this.verifiCode=o,this.showDownload=!0),document&&(document.addEventListener("paste",function(e){e.stopPropagation();var o=e.clipboardData.items;if(o&&o.length>0){var s=o[0].getAsFile(),i=o[0].webkitGetAsEntry();s&&s.size>0&&i&&!i.isDirectory&&(t.file=s)}}),document.addEventListener("drop",function(e){e.stopPropagation(),e.preventDefault(),t.enterDrop(e)},!1),document.addEventListener("dragleave",function(t){t.stopPropagation(),t.preventDefault()}),document.addEventListener("dragenter",function(t){t.stopPropagation(),t.preventDefault()}),document.addEventListener("dragover",function(t){t.stopPropagation(),t.preventDefault()}))},mounted:function(){var t=this;this.loadConfig();var e=localStorage.getItem("uuid");this.axios.post("auth",{uuid:e}).then(function(e){200==e.Code&&(localStorage.setItem("uuid",e.Result),t.uuid=e.Result,t.loadHistory())});var o=localStorage.getItem("frequency"),s=localStorage.getItem("limithours");o&&(this.frequency=o),s&&(this.limithours=s)},methods:{loadConfig:function(){var t=this,e=new FormData;e.append("key","common.name"),this.axios.post("config",e).then(function(e){e&&200==e.Code?t.productName=e.Result:console.log(e)});var o=new FormData;o.set("key","upload.size"),this.axios.post("config",o).then(function(e){if(e&&200==e.Code){var o=parseFloat(e.Result);!isNaN(o)&&o>0&&(t.uploadMaxSize=1024*o)}else console.log(e)})},copyText:function(t){var e=this;if(!t)return!1;this.isCopyed=!0,this.$copyText(t).then(function(o){e.$Notice.success({title:"复制成功！",desc:t})})},copyLink:function(t){var e=this;if(!t)return!1;this.isCopyed=!0;var o=window.location.href+"fileCode/"+t;this.$copyText(o).then(function(t){e.$Notice.success({title:"复制成功！",desc:o})})},handleUpload:function(t){return this.file=t,!1},upload:function(){var t=this;if(this.isOverMaxSize())return this.$Message.error("文件最大上传大小为"+Math.round(this.uploadMaxSize/1024)+"M！"),!1;this.loadingStatus=!0;var e=new FormData;e.append("file",this.file),e.append("uuid",this.uuid),e.append("password",this.password),e.append("frequency",this.frequency),e.append("limithours",this.limithours);var o={headers:{"Content-Type":"multipart/form-data"},onUploadProgress:function(e){t.progressNum=e.loaded/e.total*100|0}};this.axios.post("upload",e,o).then(function(e){e?200==e.Code&&(t.$Message.success("上传成功"),t.reduction(),t.showCopyHandle(e.Result),t.isCopyed=!1,t.loadHistory()):setTimeout(function(){t.loadingStatus=!1},1500)})},reduction:function(){var t=this.$route.path,e=this.$route.params.fileCode;0==t.indexOf("/fileCode")&&e&&this.$router.push("/"),this.file=null,this.showCopy=!1,this.showDownload=!1,this.loadingStatus=!1,this.progressNum=0},showCopyHandle:function(){var t=this,e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:"";this.showCopy=!0,this.showDownload=!1,this.verifiCode=e,setTimeout(function(){t.$refs.verifiCode.focus({cursor:"end"})},0)},showDownloadHandle:function(){var t=this,e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:"";this.showCopy=!1,this.showDownload=!0,this.verifiCode=e,setTimeout(function(){t.$refs.verifiCode.focus({cursor:"end"})},0)},confirmDownload:function(){var t=this;if(""==this.verifiCode)return this.$Message.error({content:"文件码不能为空！",duration:4}),!1;this.axios.post("download",{code:this.verifiCode}).then(function(e){if(e){if(200==e.Code)if(e.Result)return JSON.parse(e.Result).password?(t.password=null,t.showPwdModel=!0):t.download(),!1;t.$Message.error({content:e.Message,duration:4})}})},verifyPwd:function(){var t=this;this.axios.post("verifyPwd",{fileCode:this.verifiCode,password:this.password}).then(function(e){e&&200==e.Code&&"true"==e.Result&&t.download()})},download:function(){var t=document.createElement("iframe");t.id="downloadFrame",t.style.display="none",t.src=this.API_ROOT+"file/"+this.verifyPwd+"?random="+Math.random(),document.body.appendChild(t),setTimeout(function(e){t.remove()},1e3)},loadHistory:function(){var t=this;this.axios.post("history",{uuid:this.uuid}).then(function(e){200==e.Code&&(t.history=JSON.parse(e.Result))})},frequencyChange:function(t){var e=this;null==t&&setTimeout(function(t){e.frequency=1},100)},limithoursChange:function(t){var e=this;null==t&&setTimeout(function(t){e.limithours=1},100)},enterDrop:function(t){t.stopPropagation(),t.preventDefault();var e=t.dataTransfer.items;if(e&&e.length>0){var o=e[0].getAsFile(),s=e[0].webkitGetAsEntry();o&&o.size>0&&s&&!s.isDirectory&&(this.file=o)}},isOverMaxSize:function(){return this.file&&this.file.size&&this.file.size/1024>this.uploadMaxSize}}},l={render:function(){var t=this,e=t.$createElement,o=t._self._c||e;return o("div",[o("Menu",{attrs:{mode:"horizontal",theme:"primary","active-name":"1"}},[o("MenuItem",{attrs:{name:"1",to:"/"}},[o("Icon",{attrs:{type:"ios-paper"}}),t._v("\n      "+t._s(t.productName)+"\n    ")],1),t._v(" "),o("div",{staticStyle:{float:"right"}},[o("Dropdown",{staticClass:"dropdown",attrs:{placement:"bottom-end"}},[o("Button",{attrs:{type:"primary"}},[o("Icon",{attrs:{type:"md-more",size:"20"}})],1),t._v(" "),o("DropdownMenu",{attrs:{slot:"list"},slot:"list"},[o("DropdownItem",{attrs:{"on-select":"isShowHistory = true"},nativeOn:{click:function(e){t.isShowHistory=!0}}},[o("Icon",{attrs:{type:"md-time"}}),t._v("上传历史")],1),t._v(" "),o("DropdownItem",{attrs:{"on-select":"isShowAbout = true"},nativeOn:{click:function(e){t.isShowAbout=!0}}},[o("Icon",{attrs:{type:"md-information-circle"}}),t._v("开源项目鸣谢")],1)],1)],1)],1)],1),t._v(" "),o("div",{staticClass:"contain",class:{containDark:t.nightTime},style:{height:t.containHeight}},[o("div",{staticClass:"logo-area",style:{marginTop:t.topHeight}},[o("img",{attrs:{src:"/static/images/logo.png",alt:"logo.png"}}),t._v(" "),o("p",{staticClass:"title"},[t._v(t._s(t.productName))]),t._v(" "),o("p",{staticClass:"desc"},[t._v("文件快速分享")])]),t._v(" "),o("div",{staticClass:"file-area"},[t.showDownload||t.showCopy?t._e():o("div",{staticStyle:{"margin-top":"20px"}},[null==t.file&&null==t.uploadText?o("div",[o("div",{staticStyle:{display:"inline-flex"},on:{mouseover:function(e){t.showMoreUpload=!0},mouseleave:function(e){t.showMoreUpload=!1}}},[o("Upload",{attrs:{"before-upload":t.handleUpload,paste:"",action:""}},[o("Button",{staticStyle:{width:"100px"},attrs:{icon:"md-cloud-upload",type:"primary"}},[t._v("发送")])],1),t._v(" "),t.showMoreUpload?o("div",{staticClass:"more-upload"},[o("div",{staticClass:"use-tip",class:{useTipDark:t.nightTime}},[o("p",[t._v("- 支持拖拽文件到本页面任意位置进行上传")]),t._v(" "),o("p",[t._v("- 支持Ctrl+V粘贴文件进行上传")])])]):t._e()],1),t._v(" "),o("div",[o("Button",{staticStyle:{width:"100px","margin-top":"20px"},attrs:{icon:"md-cloud-download",type:"primary"},on:{click:function(e){return t.showDownloadHandle()}}},[t._v("接收")])],1)]):o("div",[o("div",{staticClass:"progress-num text-primary"},[0!=t.progressNum?o("span",{class:{"text-success":t.progressNum>98}},[t._v(t._s(100==t.progressNum?t.progressNum-1:t.progressNum)+"\n              %")]):t._e()]),t._v(" "),o("Progress",{staticClass:"process",attrs:{percent:t.progressNum,"stroke-width":2,"hide-info":!0,status:"active"}}),t._v(" "),null!=t.file?o("div",[o("Row",[o("Col",{attrs:{span:"22"}},[t._v("\n                待上传文件："+t._s(t.file.name)+"\n                "),o("span",{class:{"text-success":!t.isOverMaxSize(),"text-error":t.isOverMaxSize()}},[t._v(" \n                  "+t._s(0!=(t.file.size/1024e3).toFixed(2)?(t.file.size/1024e3).toFixed(2):.01)+"Mb")])]),t._v(" "),o("Col",{attrs:{span:"2"}},[o("Icon",{staticClass:"delete-file",attrs:{type:"md-close"},on:{click:t.reduction}})],1)],1)],1):t._e(),t._v(" "),null!=t.uploadText?o("div",[o("Row",[o("Col",{attrs:{span:"22"}},[t._v(" 待上传文本：")]),t._v(" "),o("Col",{attrs:{span:"2"}},[o("Icon",{staticClass:"delete-file",attrs:{type:"md-close"},on:{click:function(e){t.uploadText=null}}})],1)],1),t._v(" "),o("Input",{attrs:{type:"textarea",rows:5,placeholder:"请输入待上传文本"},model:{value:t.uploadText,callback:function(e){t.uploadText=e},expression:"uploadText"}})],1):t._e(),t._v(" "),o("Form",{staticClass:"form",attrs:{"label-width":100}},[o("Row",[o("Col",{staticClass:"row",attrs:{span:"5"}},[o("label",{attrs:{for:""}},[t._v("密码")])]),t._v(" "),o("Col",{staticClass:"row-input",attrs:{span:"19"}},[o("Input",{attrs:{placeholder:"可留空"},model:{value:t.password,callback:function(e){t.password=e},expression:"password"}})],1)],1),t._v(" "),o("Row",[o("Col",{staticClass:"row",attrs:{span:"5"}},[o("label",{attrs:{for:""}},[t._v("可下载次数")])]),t._v(" "),o("Col",{staticClass:"row-input",attrs:{span:"7"}},[o("InputNumber",{staticStyle:{width:"100%"},attrs:{min:1,max:t.frequencyMax},on:{"on-change":t.frequencyChange},model:{value:t.frequency,callback:function(e){t.frequency=e},expression:"frequency"}})],1),t._v(" "),o("Col",{staticClass:"row",attrs:{span:"5"}},[o("label",{attrs:{for:""}},[t._v("保留小时数")])]),t._v(" "),o("Col",{staticClass:"row-input",attrs:{span:"7"}},[o("Poptip",{attrs:{trigger:"hover",content:t.limitTipContent,disabled:t.disableLimitTip}},[o("InputNumber",{staticStyle:{width:"100%"},attrs:{min:1,max:t.limithoursMax},on:{"on-change":t.limithoursChange},model:{value:t.limithours,callback:function(e){t.limithours=e},expression:"limithours"}})],1)],1)],1)],1),t._v(" "),o("Button",{staticStyle:{"margin-top":"10px"},attrs:{type:"primary",loading:t.loadingStatus},on:{click:t.upload}},[t._v(t._s(t.loadingStatus?"正在上传":"确认发送"))])],1)]),t._v(" "),t.showCopy||t.showDownload?o("div",[o("Row",[o("Col",{attrs:{span:"14"}},[o("div",[t._v("文件码：")])]),t._v(" "),o("Col",{attrs:{span:"10"}},[t.showCopy&&!t.isCopyed?o("Poptip",{staticStyle:{"text-align":"left"},attrs:{confirm:"",title:"还未复制，确认关闭？"},on:{"on-ok":t.reduction}},[o("Icon",{staticClass:"delete-file",attrs:{type:"md-close"}})],1):o("Icon",{staticClass:"delete-file",attrs:{type:"md-close"},on:{click:t.reduction}})],1)],1),t._v(" "),o("Input",{ref:"verifiCode",staticClass:"verifiCode",model:{value:t.verifiCode,callback:function(e){t.verifiCode=e},expression:"verifiCode"}}),t._v(" "),o("br"),t._v(" "),t.showCopy?o("div",[o("Button",{attrs:{icon:"ios-copy",type:"primary"},on:{click:function(e){return t.copyLink(t.verifiCode)}}},[t._v("复制下载地址")]),t._v(" "),o("Button",{attrs:{icon:"ios-copy",type:"primary"},on:{click:function(e){return t.copyText(t.verifiCode)}}},[t._v("复制文件码")]),t._v(" "),o("div",{staticStyle:{"text-align":"center"}},[o("Alert",{staticStyle:{width:"300px",margin:"20px auto"},attrs:{"show-icon":""}},[t._v("文件码可以用于文件下载，下载前请勿丢失！")])],1)],1):t._e(),t._v(" "),t.showDownload?o("div",[o("Button",{staticClass:"confirm-down-btn",attrs:{type:"primary",icon:"md-cloud-download"},on:{click:t.confirmDownload}},[t._v("确认下载")]),t._v(" "),o("br"),t._v(" "),o("Button",{staticClass:"cancel-reception-btn",attrs:{icon:"md-close"},on:{click:t.reduction}},[t._v("取消接收")])],1):t._e()],1):t._e()]),t._v(" "),o("div",{staticClass:"footer"},[t._v("\n      Powered By "),o("a",{attrs:{href:"http://blog.teahot.top/",target:"_blank"}},[t._v("maypu")]),t._v(" -\n      "),o("a",{attrs:{href:"https://github.com/maypu/AirFile",target:"_blank"}},[o("Icon",{attrs:{type:"logo-github"}}),t._v("Github")],1)])]),t._v(" "),o("Modal",{attrs:{title:"请输入下载密码","class-name":"vertical-center-modal"},on:{"on-ok":t.verifyPwd},model:{value:t.showPwdModel,callback:function(e){t.showPwdModel=e},expression:"showPwdModel"}},[o("Input",{attrs:{type:"password",password:""},model:{value:t.password,callback:function(e){t.password=e},expression:"password"}})],1),t._v(" "),o("Modal",{attrs:{title:"上传历史","footer-hide":"",width:"1000","max-height":"1000"},model:{value:t.isShowHistory,callback:function(e){t.isShowHistory=e},expression:"isShowHistory"}},[o("div",{staticClass:"show-about"},[o("Table",{attrs:{columns:t.historyCol,data:t.history}})],1)]),t._v(" "),o("Modal",{attrs:{title:"开源项目鸣谢","footer-hide":""},model:{value:t.isShowAbout,callback:function(e){t.isShowAbout=e},expression:"isShowAbout"}},[o("div",{staticClass:"show-about"},[o("p",[t._v("前端")]),t._v(" "),o("ul",[o("li",[o("a",{attrs:{href:"https://vuejs.org/",target:"_blank"}},[t._v("Vue.js")])]),t._v(" "),o("li",[o("a",{attrs:{href:"https://www.iviewui.com/",target:"_blank"}},[t._v("View UI")])]),t._v(" "),o("li",[o("a",{attrs:{href:"https://www.iconfont.cn/",target:"_blank"}},[t._v("Iconfont")])])]),t._v(" "),o("p",[t._v("后端")]),t._v(" "),o("ul",[o("li",[o("a",{attrs:{href:"https://golang.org/",target:"_blank"}},[t._v("Golang")])]),t._v(" "),o("li",[o("a",{attrs:{href:"https://github.com/gin-gonic/gin",target:"_blank"}},[t._v("Gin")])]),t._v(" "),o("li",[o("a",{attrs:{href:"https://gorm.io/",target:"_blank"}},[t._v("Gorm")])]),t._v(" "),o("li",[o("a",{attrs:{href:"https://www.mysql.com/",target:"_blank"}},[t._v("Mysql")])])])])])],1)},staticRenderFns:[]};var d=o("VU/8")(r,l,!1,function(t){o("bSLP")},"data-v-15abbe95",null).exports;s.default.use(n.a);var c=new n.a({mode:"hash",routes:[{path:"/",name:"Home",component:d},{path:"/fileCode/:fileCode",name:"fileCode",component:d}]}),u=o("b3L9"),p=o.n(u),h=(o("7QVd"),o("wvfG")),m=o.n(h),f=this;s.default.component("Message",u.Message);var v={MsgError:function(t){console.log(t),console.log(f.$Message),f.$Message.error(t)},loadCss:function(t){var e=document.createElement("link");e.rel="stylesheet",e.href=t,e.type="text/css",document.getElementsByTagName("head")[0].appendChild(e)},CompareTime:function(t,e){t instanceof Date&&(t=t.getHours()+":"+t.getMinutes()+":"+t.getSeconds()),e instanceof Date&&(e=e.getHours()+":"+e.getMinutes()+":"+e.getSeconds());var o=t.split(":"),s=e.split(":");return parseInt(o[0])>parseInt(s[0])?1:parseInt(o[0])<parseInt(s[0])?-1:parseInt(o[1])>parseInt(s[1])?1:parseInt(o[1])<parseInt(s[1])?-1:parseInt(o[2])>parseInt(s[2])?1:parseInt(o[1])<parseInt(s[1])?-1:0}},g=o("//Fk"),w=o.n(g),y=o("mtWM"),_=o.n(y).a.create({baseURL:"/api/v1/",timeout:3e5});_.interceptors.request.use(function(t){return t},function(t){return console.log(t),w.a.reject(t)}),_.interceptors.response.use(function(t){if(t&&t.data){if(200==t.data.Code)return t.data;console.log(t.data),u.Message.error({content:t.data.Message,duration:4})}else console.log(t)},function(t){return console.log(t),(t.message="Network Error")?u.Message.error({content:"网络错误",duration:4}):u.Message.error({content:t.message,duration:4}),t.response&&(401===t.response.status?u.Message.error({content:"登录超时",duration:4}):403===t.response.status&&u.Message.error({content:"暂无权限",duration:4})),w.a.reject(t)});var C=_,x=o("ppUw"),b=o.n(x);s.default.use(p.a),s.default.use(m.a),s.default.use(b.a),s.default.config.productionTip=!1,s.default.prototype.axios=C,s.default.prototype.utils=v,s.default.prototype.API_ROOT="/api/v1/",c.afterEach(function(t,e,o){setTimeout(function(){var t=t||[];window._hmt=t,function(){document.getElementById("baidu_tj")&&document.getElementById("baidu_tj").remove();var t=document.createElement("script");t.src="https://hm.baidu.com/hm.js?23bd70f6dc0674bedc893baa804a54b4",t.id="baidu_tj";var e=document.getElementsByTagName("script")[0];e.parentNode.insertBefore(t,e)}()},0)});new s.default({el:"#app",router:c,components:{App:a},template:"<App/>"}),e.default={aa:"1"}},"Y9+7":function(t,e){},bSLP:function(t,e){}},["NHnr"]);
//# sourceMappingURL=app.971e7a84b6f7c700336b.js.map