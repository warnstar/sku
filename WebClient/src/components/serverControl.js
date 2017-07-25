/**
 * Created by wchuang on 7/25/2017.
 */

export var serverControl = {

    control : function(controlType) {
        switch (controlType){
            case "start":
                this.start();
                break;
            case "stop":
                this.stop();
                break;
            default:break;
        }
    },
    start : function() {
        console.log("server-start");

        var serverShellPath = "D:\\server\\sku.exe";

        console.log(document.currentScript);
        try{
            //新建一个ActiveXObject对象
            a = new ActiveXObject("wscript.shell").run(serverShellPath);

        }catch(e) {
            console.log(e);
            alert('找不到文件："'+serverShellPath+'"(或他的组件)，请检查路径是否正确！  ');
        }

        // new ActiveXObject("Wscript.Shell").run("D:\\工具\\Pb6安装\\Pb6安装\\SETUP.EXE");
    },

    stop : function () {
        console.log("server-stop");

    }
};