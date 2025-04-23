layui.define(['jquery', 'element'], function(exports) {
    "use strict";

    var MOD_NAME = 'mdEditor',
        $ = layui.jquery,
        element = layui.element;

    var mdEditor = new function (){
        let editor;
        this.initEditor = function (content){
            editor = editormd("desc", {
                width   : "100%",
                height  : 810,
                syncScrolling : "single",
                saveHTMLToTextarea : false,
                watch:true,
                path    : "/static/component/editormd/lib/",
                toolbarIcons:function (){
                    return ["save","undo","redo","|","bold","del","italic","quote","ucwords","uppercase","lowercase","|","h1","h2","h3","h4","h5","h6","|","list-ul","list-ol","hr","|","link","reference-link","image","code","preformatted-text","code-block","table","datetime","html-entities","pagebreak","|","goto-line","watch","preview","fullscreen","clear","search"]
                },
                toolbarIconsClass:{
                    save: "fa-save"
                },
                toolbarHandlers:{
                    save:function (){
                        window.save();
                    }},
                lang:{
                    toolbar:{
                        save:"保存（Ctrl+S）"
                    }},
                imageUpload    : true,
                imageFormats   : ["jpg", "jpeg", "gif", "png", "bmp", "webp"],
                imageUploadURL : "/system/upload",
                onload : function() {
                    var keyMap = {
                        "Ctrl-S":function (){
                            window.save();
                        }
                    };
                    this.addKeyMap(keyMap);
                    this.setMarkdown(content);
                    initPasteDragImg(this);
                }
            });
        }

        this.setMarkdownValue = function (content){
            editor.setMarkdown(content)
        };

        this.getMarkdownValue = function (){
            return editor.getMarkdown()
        };

        this.getPreviewedHTMLValue = function (){
            return editor.getPreviewedHTML()
        };

    };

    function initPasteDragImg(Editor) {
        var doc = document.getElementById(Editor.id)
        doc.addEventListener('paste', function (event) {
            var items = (event.clipboardData || window.clipboardData).items;
            var file = null;
            if (items && items.length) {
                // 搜索剪切板items
                for (var i = 0; i < items.length; i++) {
                    if (items[i].type.indexOf('image') !== -1) {
                        file = items[i].getAsFile();
                        break;
                    }
                }
            } else {
                console.log("当前浏览器不支持");
                return;
            }
            if (!file) {
                console.log("粘贴内容非图片");
                return;
            }
            uploadImg(file,Editor);
        });
        var dashboard = document.getElementById(Editor.id)
        dashboard.addEventListener("dragover", function (e) {
            e.preventDefault()
            e.stopPropagation()
        })
        dashboard.addEventListener("dragenter", function (e) {
            e.preventDefault()
            e.stopPropagation()
        })
        dashboard.addEventListener("drop", function (e) {
            e.preventDefault()
            e.stopPropagation()
            var files = this.files || e.dataTransfer.files;
            uploadImg(files[0],Editor);
        })
    }

    function uploadImg(file,Editor){
        var formData = new FormData();
        var fileName=new Date().getTime()+"."+file.name.split(".").pop();
        formData.append('file', file, fileName);
        $.ajax({
            //配置图片上传的接口路径
            url: Editor.settings.imageUploadURL,
            type: 'post',
            data: formData,
            processData: false,
            contentType: false,
            dataType: 'json',
            success: function (msg) {
                var code=msg['code'];
                if(code===200){
                    var url=msg["data"];
                    if(/\.(png|jpg|jpeg|gif|bmp|ico)$/.test(url)){
                        Editor.insertValue("![image](/"+url+")");
                    }else{
                        Editor.insertValue("[下载附件](/"+url+")");
                    }
                }else{
                    console.log(msg);
                    alert("上传失败");
                }
            },
            error: function () {
                console.log('请求失败');
            }
        });
    }

    exports(MOD_NAME, mdEditor);
});