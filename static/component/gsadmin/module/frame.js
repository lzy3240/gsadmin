layui.define(['table', 'jquery', 'element'], function (exports) {
    "use strict";

    var MOD_NAME = 'frame',
        $ = layui.jquery,
		element = layui.element;
		
    var gsadminFrame = function (opt) {
        this.option = opt;
    };

    gsadminFrame.prototype.render = function (opt) {
		var option = {
			elem:opt.elem,
			url:opt.url,
			title:opt.title,
			width:opt.width,
			height:opt.height,
			done:opt.done ? opt.done: function(){ console.log("菜单渲染成功");}
		}
	    createFrameHTML(option);
	    $("#"+option.elem).width(option.width);
	    $("#"+option.elem).height(option.height);
		return new gsadminFrame(option);
    } 
	
	gsadminFrame.prototype.changePage = function(url,title,loading){
		if(loading){
			var loading = $("#"+this.option.elem).find(".gsadmin-frame-loading");	
			loading.css({display:'block'});
		}
		$("#"+this.option.elem+" iframe").attr("src",url);
	    $("#"+this.option.elem+" .title").html(title);
	     if(loading){
	     	var loading = $("#"+this.option.elem).find(".gsadmin-frame-loading");
			setTimeout(function(){
				loading.css({display:'none'});
			},800)	
	     }
	}
	
	gsadminFrame.prototype.changePageByElement = function(elem,url,title,loading){
		if(loading){
			var loading = $("#"+elem).find(".gsadmin-frame-loading");	
			loading.css({display:'block'});
		}
		$("#"+elem+" iframe").attr("src",url);
	    $("#"+elem+" .title").html(title);
	     if(loading){
	     	var loading = $("#"+elem).find(".gsadmin-frame-loading");
			setTimeout(function(){
				loading.css({display:'none'});
			},800)	
	     }
	}
	
	gsadminFrame.prototype.refresh = function (time) {
		if(time!=false){
			var loading = $("#"+this.option.elem).find(".gsadmin-frame-loading");
			loading.css({display:'block'});
			if(time!=0){
				setTimeout(function(){
					loading.css({display:'none'});
				},time)
			}
		}
		$("#"+this.option.elem).find("iframe")[0].contentWindow.location.reload(true);
	}
	
	function createFrameHTML(option){
		 var header = "<div class='gsadmin-frame-title'><div class='dot'></div><div class='title'>"+option.title+"</div></div>"
		 var iframe = "<iframe class='gsadmin-frame-content' style='width:100%;height:100%;'  scrolling='auto' frameborder='0' src='"+option.url+"' ></iframe>";
	     var loading = '<div class="gsadmin-frame-loading">'+
			       '<div class="ball-loader">'+
				      '<span></span><span></span><span></span><span></span>'+
			       '</div>'+
		        '</div></div>';
	     $("#"+option.elem).html("<div class='gsadmin-frame'>"+header+iframe+loading+"</div>");	
	}
	exports(MOD_NAME,new gsadminFrame());
})