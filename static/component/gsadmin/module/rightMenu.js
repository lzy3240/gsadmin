layui.define(['element'], function (exports) {
    let element = layui.element;
    const $ = layui.jquery;

    let MOD_NAME = 'rightMenu';
    let RIGHTMENUMOD = function () {
        this.v = '1.0.0';
        this.author = 'raowenjing';

    };
    String.prototype.format = function () {
        if (arguments.length == 0) return this;
        let param = arguments[0];
        let s = this;
        if (typeof(param) == 'object') {
            for (var key in param) s = s.replace(new RegExp("\\{" + key + "\\}", "g"), param[key]);
            return s;
        } else {
            for (var i = 0; i < arguments.length; i++) s = s.replace(new RegExp("\\{" + i + "\\}", "g"), arguments[i]);
            return s;
        }
    }

    function createStyle(ulClassName) {
        let style = '.{name} {position: absolute;width: 60px;z-index: 9999;display: none;background-color: #fff;padding: 2px;color: #333;border: 1px solid #c2c2c2;border-radius: 2px;cursor: pointer;}.{name} li {text-align: center;display: block;height: 30px;line-height: 32px;}.{name} li:hover {background-color: #666;color: #fff;}'
            .format({name: ulClassName});
        return style;
    }

    /**
     * 初始化
     */
    RIGHTMENUMOD.prototype.render = function (opt) {
        createStyle();

        if (!opt.container) {
            console.error("[ERROR]使用rightmenu组件需要制定'container'属性！");
            return;
        }
        let defaultNavArr = [

        ];
        opt = opt || {};
        opt.triggerDom = opt.triggerDom || "li";
        opt.boxClassName = opt.boxClassName || "right-click-menu-container";
        opt.navArr = opt.navArr || defaultNavArr;
        CreateRightMenu(opt,"");
        _CustomRightClick(opt);
    };


    /**
     * 创建右键菜单项目
     * @param rightMenuConfig
     * @constructor
     */
    function CreateRightMenu(rightMenuConfig,currentData,callback) {
        if($('.'+rightMenuConfig.boxClassName).length>0) $('.'+rightMenuConfig.boxClassName).remove();

        $("<style></style>").text(createStyle(rightMenuConfig.boxClassName)).appendTo($("head"));
        let li = '';
        $.each(rightMenuConfig.navArr, function (index, conf) {
            if(!!currentData && typeof conf.showFormat != "undefined"){ // 控制
                if(typeof conf.showFormat == "function"){
                    var isShow = conf.showFormat(currentData);
                    isShow = !!isShow?true:false;
                    if(isShow){
                        li += '<li data-type="{eventName}"><i class="layui-icon {icon}"></i>{title}</li>'
                            .format({eventName:conf.eventName,icon:conf.icon?conf.icon:"",title:conf.title});
                    }
                }else{
                    if(!!conf.showFormat){
                        li += '<li data-type="{eventName}"><i class="layui-icon {icon}"></i>{title}</li>'
                            .format({eventName:conf.eventName,icon:conf.icon?conf.icon:"",title:conf.title});
                    }
                }
            }else{
                li += '<li data-type="{eventName}"><i class="layui-icon {icon}"></i>{title}</li>'
                    .format({eventName:conf.eventName,icon:conf.icon?conf.icon:"",title:conf.title});
            }
        })
        let tmpHtml = '<ul class="{className}">{liStr} </ul>'.format({liStr: li, className: rightMenuConfig.boxClassName})
        $(rightMenuConfig.container).after(tmpHtml);

        setTimeout(function(){
            registerMenuClick(rightMenuConfig);

            if(!!callback && typeof callback == "function") callback();
        },10)
    }

    /**
     * 绑定右键菜单
     * @constructor
     */
    function _CustomRightClick(rightMenuConfig) {
        $(rightMenuConfig.container).off("click");
        $(rightMenuConfig.container).on("click", rightMenuConfig.triggerDom,function () {
            $('.'+rightMenuConfig.boxClassName).hide();
        });
        $(rightMenuConfig.container).off("contextmenu")
        $(rightMenuConfig.container).on("contextmenu", rightMenuConfig.triggerDom, function (e) {
            // 阻止默认的右键菜单
            e.preventDefault();

            // 重构菜单
            CreateRightMenu(rightMenuConfig,$(this).data(),function(){
                let popupmenu = $("."+rightMenuConfig.boxClassName);
                let leftValue = ($(document).width() - e.clientX) < popupmenu.width() ? (e.clientX - popupmenu.width()) : e.clientX;
                let topValue = ($(document).height() - e.clientY) < popupmenu.height() ? (e.clientY - popupmenu.height()) : e.clientY;
                popupmenu.css({left: leftValue, top: topValue}).show();
            });


            // 将该元素的所有数据复制到
            $.each($(this).data(), function(key, value) {
                $("."+rightMenuConfig.boxClassName).data(key, value);
            });

            return false;
        });
        // 点击空白处隐藏弹出菜单
        $(document).click(function (event) {
            event.stopPropagation();
            $("."+rightMenuConfig.boxClassName).hide();
        });


    }

    function registerMenuClick(rightMenuConfig){
        $('.' + rightMenuConfig.boxClassName + ' li').prop("onclick",null).off("click");
        /**
         * 注册tab右键菜单点击事件
         */
        $('.' + rightMenuConfig.boxClassName + ' li').click(function () {
            rightMenuConfig.registMethod[$(this).attr("data-type")]($(this).parents("."+rightMenuConfig.boxClassName).data());

            $("."+rightMenuConfig.boxClassName).hide();
        })
    }


    let rightmenuObj = new RIGHTMENUMOD();
    exports(MOD_NAME, rightmenuObj);
})