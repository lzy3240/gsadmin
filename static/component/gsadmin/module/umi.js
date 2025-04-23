layui.define(['jquery', 'element'], function(exports) {
    "use strict";

    // import put from "./context";
    // import get from "./context";

    var MOD_NAME = 'umi',//用户标识信息
        $ = layui.jquery,
        element = layui.element;

    var umi = new function (){
        // this.initMark = function (){
        //     generateLocalMark();
        // };
        //
        // this.initWelcome = function (){
        //     var data = getSetting();
        //     var mark = generateLocalMark();
        //     flushWelcome(data,mark);
        // }
        $(function (){
            var data = getSetting();
            var mark = generateLocalMark();
            flushWelcome(data,mark);
        })
    }

    function flushWelcome(data,mark){
        //修改欢迎词
        if (mark.flag){
            $("#title").html('<span>欢迎回来!</span>')
            $("#last_time").html('<span>上次访问时间:</span><span> '+mark.last_time+'</span>')
        }else {
            $("#title").html('<span>访客您好~</span>')
            $("#last_time").html('<span>上次访问时间:</span><span> '+''+'</span>')
        }

        //修改欢迎信息
        $("#local_ip").html('<span>当前IP源地址:</span><span> '+data.local.ip+'</span>')
        $("#local_addr").html('<span>当前IP归属地:</span><span> '+data.local.addr+'</span>')

        //修改介绍信息
        $("#description").html('<span>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;'+data.site.description+'</span>')
    }

    function generateLocalMark(){
        var flag = false;
        var last_time = "";
        var mark = localStorage.getItem("web_marker");
        if (mark === null){
            var ts = Date.now()
            var tm = generateFingerprint()+":"+ts
            localStorage.setItem("web_marker",tm)
            return {"flag":flag,"last_time":last_time}
        }else {
            var tmp =mark.split(":")
            var tm0 = tmp[0]
            var ts1 = Number(tmp[1])
            var ts2 = Date.now()
            // var day = Math.abs(ts2 - ts1)/(1000 * 60 * 60 * 24)
            // if (day > 1){
            //     var description = data.site.Description;
            //     var copyright = data.site.Copyright;
            //     var icp = data.site.Icp;
            //
            //     localStorage.setItem("web_desc",description);
            //     localStorage.setItem("web_copyright",copyright);
            //     localStorage.setItem("web_icp",icp);

            var date = new Date(ts1)
            let options = {
                timeZone: "Asia/Shanghai",
                year: 'numeric',
                month: '2-digit',
                day: '2-digit',
                hour: '2-digit',
                minute: '2-digit',
                second: '2-digit'
            };

            localStorage.setItem("web_marker",tm0+":"+ts2)
            let formatter = new Intl.DateTimeFormat("zh-CN", options);
            let date_str = formatter.format(date).replace(/\//g,"-") //格式化并替换/
            return {"flag":true,"last_time":date_str}
        }
    }

    function getSetting(){
        var data = "";
        $.ajax({
            url:"/welcome",
            type: "get",
            async:false,
            success:function (res){
                data = res.data;
            }
        });

        return data;
    }

    function generateFingerprint() {
        var screenSize = getScreenSize();
        var viewportSize = getViewportSize();
        var colorDepth = screen.colorDepth;
        var userAgent = navigator.userAgent;
        var platform = navigator.platform;
        var language = navigator.language;
        var plugins = Array.from(navigator.plugins).map(plugin => plugin.name).join(',');
        var doNotTrack = navigator.doNotTrack;
        var timeZone = Intl.DateTimeFormat().resolvedOptions().timeZone;
        var timeZoneOffset = new Date().getTimezoneOffset();
        var fonts = getFonts();
        var hardwareConcurrency = navigator.hardwareConcurrency;
        var deviceMemory = navigator.deviceMemory;
        var webGLRenderer = getWebGLRenderer();
        var canvasFingerprint = getCanvasFingerprint();
        var touchSupport = 'ontouchstart' in window || navigator.maxTouchPoints > 0;

        var fingerprint = [
            screenSize, viewportSize, colorDepth, userAgent, platform, language,
            plugins, doNotTrack, timeZone, timeZoneOffset, fonts, hardwareConcurrency,
            deviceMemory, webGLRenderer, canvasFingerprint, touchSupport
        ].join('|');

        // console.log(fingerprint)
        return hashString(fingerprint);
    }

    function hashString(str) {
        let hash = 0;
        for (let i = 0; i < str.length; i++) {
            const char = str.charCodeAt(i);
            hash = ((hash << 5) - hash) + char;
            hash |= 0; // Convert to 32bit integer
        }
        return hash.toString(16);
    }

    function getScreenSize() {
        return `${screen.width}x${screen.height}`;
    }

    function getViewportSize() {
        return `${window.innerWidth}x${window.innerHeight}`;
    }

    function getFonts() {
        const fontList = [
            'Arial', 'Verdana', 'Times New Roman', 'Courier New', 'Georgia', 'Comic Sans MS', 'Trebuchet MS', 'Arial Black', 'Impact'
        ];
        const availableFonts = [];

        const canvas = document.createElement('canvas');
        const context = canvas.getContext('2d');
        const text = 'abcdefghijklmnopqrstuvwxyz0123456789';

        fontList.forEach((font) => {
            context.font = `16px ${font}`;
            const width = context.measureText(text).width;
            context.font = `16px monospace`;
            if (context.measureText(text).width !== width) {
                availableFonts.push(font);
            }
        });

        return availableFonts.join(',');
    }

    function getWebGLRenderer() {
        const canvas = document.createElement('canvas');
        const gl = canvas.getContext('webgl') || canvas.getContext('experimental-webgl');
        if (gl) {
            const debugInfo = gl.getExtension('WEBGL_debug_renderer_info');
            if (debugInfo) {
                return gl.getParameter(debugInfo.UNMASKED_RENDERER_WEBGL);
            }
        }
        return null;
    }

    function getCanvasFingerprint() {
        const canvas = document.createElement('canvas');
        const context = canvas.getContext('2d');
        context.textBaseline = 'top';
        context.font = '14px Arial';
        context.textBaseline = 'alphabetic';
        context.fillStyle = '#f60';
        context.fillRect(125, 1, 62, 20);
        context.fillStyle = '#069';
        context.fillText('Based on GsAdmin.', 2, 15);
        context.fillStyle = 'rgba(102, 204, 0, 0.7)';
        context.fillText('Based on GsAdmin.', 4, 17);
        return canvas.toDataURL();
    }

    exports(MOD_NAME, umi);
})