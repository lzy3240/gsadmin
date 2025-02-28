layui.define(['table', 'jquery', 'element'], function(exports) {
	"use strict";

	var MOD_NAME = 'message',
		$ = layui.jquery,
		element = layui.element;

	var message = function(opt) {
		this.option = opt;
	};

	message.prototype.render = function(opt) {
		//默认配置值
		var option = {
			elem: opt.elem,
			url: opt.url ? opt.url : false,
			height: opt.height,
			data: opt.data,
		}
		if (option.url !== false) {
			option.data = getData(option.url);
			var notice = createHtml(option);
			$(option.elem).html(notice);
		}
		setTimeout(function(){
			element.init();
		},300);
		return new message(option);
	}
	
	message.prototype.click = function(callback){
		$("*[notice-id]").click(function(event) {
			event.preventDefault();
			var id = $(this).attr("notice-id");
			var title = $(this).attr("notice-title");
			var context = $(this).attr("notice-context");
			var from = $(this).attr("notice-from");
			var created_at =$(this).attr("notice-created_at");
			callback(id, title, context, from, created_at);
		})
	}
	
	/** 同 步 请 求 获 取 数 据 */
	function getData(url) {
		$.ajaxSettings.async = false;
		var data = null;
		$.get(url, function(result) {
			data = result;
		});
		$.ajaxSettings.async = true;
		return data.data;
	}

	function refresh(){
		console.log(this.option)
	}

	function createHtml(option) {
        var notice = '<li class="layui-nav-item" lay-unselect="">' +
			'<a href="#" class="notice layui-icon layui-icon-notice">'
		if ((option.data.length) == 0) {
			notice += '<span class="layui-badge-dot" style="display: none"></span> </a>'
		} else {
			notice += '<span class="layui-badge-dot" ></span> </a>'
		}

		// notice += `</li>`;
		notice += '<div class="layui-nav-child layui-tab gsadmin-notice" style="margin-top: 0px; max-height:360px; left: -200px; overflow: auto;">';

		var noticeList = '<ul class="list">';
		if ((option.data.length) == 0) {
			noticeList += '<li class="list-item">';
			noticeList += '<div class="gsadmin-notice-item">' +
				'<span>' + "没有未读消息" + '</span> </div> </li>' ;
		} else {
			// 根据 data 遍历数据
			$.each(option.data, function(i, item) {
				noticeList += '<li class="list-item">';
				noticeList += '<div class="gsadmin-notice-item" notice-from="' + item.notice_from + '" notice-context="' +item.notice_context +
					'" notice-title="' + item.notice_title + '" notice-id="' + item.id + '" notice-created_at="' + item.created_at + '">' +
					'<span>' + '[' + item.notice_type + ']' + '</span>' +
					'<span>' + item.notice_title.substring(0,8) + '</span>' +
					'<span>' + item.created_at.substring(0,10) + '</span>' +
					'</div>'
				noticeList += '</li>';
			})
		}

		noticeList += '</ul>';
		notice += noticeList;
		notice += '</div></li>';
		return notice;
	}

	exports(MOD_NAME, new message());
})
