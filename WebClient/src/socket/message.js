export var message = {
    WEB_USER : 'web_user',
    WEB_TSI_CHECK : 'web_tsi_check',
    WEB_TSI_NOW_DATA : 'web_tsi_now_data',
    WEB_TSI_TEST : 'web_tsi_test',
    WEB_TSI_TEST_PRE : 'web_tsi_test_pre',
    WEB_CLIENT_CONNECT_COMPLETE : 'web_client_connect_complete',
    WEB_CLIENT_TIME_SYNC_COMPLETE : 'web_client_time_sync_complete',
    WEB_CLIENT_CONNECT_AND_TIME_SYNC_CHECK : 'web_client_connect_and_time_sync_check',
    WEB_CLIENT_TREE_DATA : 'web_client_tree_data',
    WEB_CLIENT_LOG : 'web_client_log',
    WEB_CLIENT_EXIT : 'web_client_exit',
    WEB_CAN_START_TSI_TEST : 'web_can_start_tsi_test',
    WEB_TSI_TEST_MODULE_RESULT : 'web_tsi_test_module_result',

    socket : null,
    
    option : {
        host : "localhost",
        tsi_host : "172.16.15.214",
        port : 9502,
        client_num : 8,
        client_module_num : 16
    },
    getOption : function() {
        var host = localStorage.getItem("option.host");
        this.option.host = host ? host : this.option.host;
        
        var tsi_host = localStorage.getItem("option.tsi_host");
        this.option.tsi_host = tsi_host ? tsi_host : this.option.tsi_host;
        
        var port = localStorage.getItem("option.port");
        this.option.port = port ? port : this.option.port;

        var client_num = localStorage.getItem("option.client_num");
        this.option.client_num = client_num ? client_num : this.option.client_num;

        var client_module_num = localStorage.getItem("option.client_module_num");
        this.option.client_module_num = client_module_num ? client_module_num : this.option.client_module_num;

        return this.option;
    },
    setServerAddress : function(option) {
        localStorage.setItem("option.host", option.host);
        localStorage.setItem("option.tsi_host", option.tsi_host);
        localStorage.setItem("option.port", option.port);
        localStorage.setItem("option.client_num", option.client_num);
        localStorage.setItem("option.client_module_num", option.client_module_num);
    },
    getServerAddress : function() {
        var host, port;

        host = localStorage.getItem("option.host");
        host = host ? host : this.option.host;

        port = localStorage.getItem("option.port");
        port = port ? port : this.option.port;

        return  'ws:' +  host + ":" + port;
    },
    getMsg : function(type, content) {
        var data = {
            type : type,
            content : content
        };
        return JSON.stringify(data);
    },
    connect : function() {
        var _this = this;
        this.socket = new WebSocket(this.getServerAddress());
        var a = new WebSocket(this.getServerAddress());

        var socket = this.socket;
        socket.onopen = function(event) {
            console.log('已成功连接上服务器！');
            // 上报浏览器信息
            var msg = _this.getMsg(_this.WEB_USER, _this.getOption());

            socket.send(msg);

            var i = 0;
            socket.onclose = function(code,reason) {
                //重启页面
                setTimeout('location.reload()',1500);
            };
        };
    },
    isConnect : function() {
        var state = false;
        if (this.socket.readyState == 1) {
            state = true;
        }
        return state;
    },
    sendTsiCheck : function() {
        var msg = this.getMsg(this.WEB_TSI_CHECK, '');
        this.socket.send(msg)
    },
    sendClientConnectAndTimeSyncCheck : function() {
        var msg = this.getMsg(this.WEB_CLIENT_CONNECT_AND_TIME_SYNC_CHECK, '');
        this.socket.send(msg)
    },
    sendTsiCollectPre : function() {
        var msg = this.getMsg(this.WEB_TSI_TEST_PRE, '');
        this.socket.send(msg)
    },
    sendTsiCollect : function() {
        var msg = this.getMsg(this.WEB_TSI_TEST, '');
        this.socket.send(msg)
    },
    sendClientExit : function() {
        var msg = this.getMsg(this.WEB_CLIENT_EXIT, '');
        this.socket.send(msg)
    }
};
