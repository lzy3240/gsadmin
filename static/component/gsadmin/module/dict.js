layui.define(['element','jquery'], function(exports) {
    "use strict";

    var MOD_NAME = 'dict',
        element = layui.element,
        $ = layui.jquery;

    var dict = function () {};

    dict.prototype.renderDictAll = function (formId, b) {
        var element = layui.element,
            $ = layui.$;

        $.ajaxSettings.async = false; //ajax设置为同步
        $('.selDict').each(function(){
            var _this = $(this);
            var dict = _this.attr("dict");
            $.ajax()
            $.get("/system/dict/data/json?dict_type="+dict, function(data){
                _this.append("<option value=''>请选择</option>");
                if(data.data!=null){
                    $.each(data.data, function(index, item){
                        _this.append("<option value='"+item.dict_value+"'>"+item.dict_label+"</option>");
                    });
                }
            });
        })

        $.ajaxSettings.async = true; //ajax恢复为异步
    }

    exports(MOD_NAME, new dict());
});