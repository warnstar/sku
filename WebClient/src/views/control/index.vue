<template>
    <div class="control_main">

        <el-row type="flex" class="row-bg control_top" justify="center" style="margin-top: 75px;">
            <el-col :span="5" class="row-bg" style="margin: 10px;border: thin;">
                <span><b>客户端：</b></span>

                <router-link to="/setting"><el-button class="el-icon-setting"></el-button></router-link>
                <el-button @click="alertToReset">重置</el-button>
            </el-col>

            <el-col :span="15" class="row-bg" style="margin: 10px;border: thin;">
                <!--<span><b>服务器：</b></span>-->

                <!--<el-button @click="serverCtl">启动</el-button>-->
            </el-col>
        </el-row>

        <el-row type="flex" class="row-bg control_body" justify="center" style="margin-top: 75px;">
            <el-col :span="18" class="row-bg">
                <el-row>
                    <el-col :span="5" class="row-bg control_body_left">
                        <el-row class="row-bg">
                            <el-steps :active="active" direction="vertical" finish-status="success">
                                <el-step title="检查TSI状态"></el-step>
                                <el-step style="margin-top: 15px;" title="客户端连接"></el-step>
                                <el-step style="margin-top: 15px;" title="客户端时间同步"></el-step>
                                <el-step style="margin-top: 15px;" title="启动TSI校准"></el-step>
                                <el-step style="margin-top: 15px;" title="启动TSI测量"></el-step>
                            </el-steps>
                        </el-row>

                        <el-row v-loading="isLoading" element-loading-text="服务端处理中" :span="24" class="row-bg">
                            <el-button :disabled="unNext" type="warning" style="margin-top: 15px;" @click="next">下一步</el-button>
                        </el-row>
                    </el-col>

                    <el-col :span="16" class="row-bg control_body_center">
                        <!-- TSI 实时数据 -->
                        <div style="width:78%;padding: 10px 10px 10px 50px" :hidden="tsi_data_hide">
                            <canvas id="tsi_data_body"></canvas>
                        </div>

                        <!-- 测试结果分析 -->
                        <div :hidden="test_result_body_hide">
                            <el-row class="test_result_body" v-for="pi in piLight" :key="pi.pid">
                                <el-col :span="2">{{pi.pid}}</el-col>

                                <div v-for="module in pi.modules">
                                    <el-col :span="1">
                                        <el-popover
                                                trigger="hover">
                                            <el-row type="flex" class="row-bg" justify="center">
                                                <el-col style="text-align: center" class="row-bg">
                                                    <h2>派：{{pi.pid}}&nbsp&nbspUSB：{{module.module_id}}</h2>
                                                </el-col>
                                            </el-row>
                                            <el-table :data="module.info" title="采集数据信息">
                                                <el-table-column width="100" property="stage" label="范围"></el-table-column>
                                                <el-table-column width="100" property="total" label="总采集点"></el-table-column>
                                                <el-table-column width="100" property="error" label="错误数"></el-table-column>
                                                <el-table-column width="100" property="proportion" label="错误率"></el-table-column>
                                            </el-table>
                                            <span slot="reference">
                                                <span v-if="module.status == 'success'" style="color: #13CE66;" class="el-icon-star-on"></span>
                                                <span v-else-if="module.status == 'error'" style="color: #FF4949;" class="el-icon-star-on"></span>
                                                <span v-else="module.status == 'off'" style="color: #D3DCE6;" class="el-icon-star-on"></span>
                                            </span>
                                        </el-popover>
                                    </el-col>
                                </div>

                            </el-row>
                        </div>

                    </el-col>
                </el-row>

                <el-row type="flex" class="row-bg control_footer" justify="left">
                    <el-col :span="3" class="row-bg"  style="margin-top: 15px;">
                        <b>消息日志：</b>
                    </el-col>
                    <el-col :span="18" class="row-bg log_info_box" style="margin-top: 25px;">
                        <el-row v-for="logOne in logData" :key="logOne.id"  class="log_info_msg" justify="left">
                            <el-col :span="4">
                                <span> {{ logOne.time }} </span>
                            </el-col>
                            <el-col :span="2">
                                <span> {{ logOne.type }} </span>
                            </el-col>
                            <el-col :span="18" style="word-wrap:break-word ;">
                                <span style=""> {{ logOne.content }} </span>
                            </el-col>
                        </el-row>
                    </el-col>
                </el-row>
            </el-col>


            <el-col :span="3" class="row-bg control_body_right" >
                <el-tree :data="treeData" :props="defaultProps" @node-click="handleNodeClick"></el-tree>
            </el-col>
        </el-row>

    </div>

</template>


<style>
    .control_top {
        margin: 10px;
    }
    .control_body {
        margin: 10px;
    }
    .control_body_left {
        margin: 25px;
        border: solid;border-width: thin;
        padding: 20px;
    }
    .control_body_center {
        min-height: 313px;
        margin: 25px;border: solid;border-width: thin;
    }
    .control_body_right {
        padding: 15px;
        margin: 25px;border: solid;border-width: thin;
    }
    .control_footer{
        height: 150px;
        margin-left: 45px;
    }
    .control_footer .log_info_box {
        overflow-y:scroll;
        margin: 10px 0 10px 0;
    }
    .control_footer .log_info_msg {
        margin-top: 5px;
    }

    .test_result_body {
        margin: 15px 10px 15px 25px;
    }
</style>

<script>
    import { message } from '../../socket/message';
    import { clientLight } from '../../components/clientLight';
    import { serverControl } from '../../components/serverControl';
    import { tsiStatus } from '../../components/tsiStatus';
    import popover from "element-ui/packages/popover/src/directive";
    import"Chart.js"

    const  ACTIVE_INIT = 0;
    const  ACTIVE_TSI_CHECK = 1;
    const  ACTIVE_CLIENT_CONNECT = 2;
    const  ACTIVE_CLIENT_TIME_SYNC = 3;
    const  ACTIVE_TSI_COLLECT_PRE = 4;
    const  ACTIVE_TSI_COLLECT = 5;
    export default {
        components: {popover},
		data() {
			return {
                treeData: [],
                defaultProps: {
                    children: 'children',
                    label: 'label'
                },
                isLoading : false,
                socket : null,
				active: ACTIVE_INIT,
                isAllClientTimeSync : false,
                isAllClientConnected : false,
                logData: [],
                unNext : false,
                piLight : [],
                test_result_body_hide : true,
                tsi_data_hide : false,
                tsi_data : {
                    labels : [],
                    datasets : [
                        {
                            label: "TSI 实时数据",
                            backgroundColor: 'rgba(54, 162, 235, 0.5)',
                            borderColor: 'rgba(54, 162, 235, 0.5)',
                            fill: false,
                            pointRadius: 0,
                            pointHoverRadius: 1,
                            data : []
                        }
                    ]
                }
			};
		},
		methods: {
            serverCtl() {
                serverControl.start();
            },
            alertToReset() {
                var _this = this;
                this.$alert('请确认是否重置应用!!!',{
                    title: "警告",
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    callback: function(msg) {
                        if (msg == 'confirm') {
                            //通知终端退出重启
                            message.sendClientExit();

                            setTimeout('location.reload()',1000);
                        }
                    }
                });
            },
            alertTsiTestPre(isPre) {
                var _this = this;
                this.$alert('即将启动TSI校准或测量，请点烟！',{
                    confirmButtonText: '确定',
                    callback: function(msg) {
                        if (msg == 'confirm') {
                            //启动tsi校准
                            _this.isLoading = true;
                            if (isPre === false) {
                                message.sendTsiCollect();
                            } else {
                                message.sendTsiCollectPre();
                            }

                            //没有下一步了
                            _this.unNext = true;
                        }
                    }
                });
            },
            handleNodeClick(data) {
                //console.log(data);
            },
			next() {
                //检查是否已连上服务器
                if (!message.isConnect()) {
                    this.$message.error('与服务器断开连接');
                    setTimeout('location.reload()',1000);
                    return ;
                }

                switch(this.active) {
                    case ACTIVE_INIT :
                        //启动tsi检查
                        this.isLoading = true;
                        message.sendTsiCheck();
                        break;
                    case ACTIVE_TSI_CHECK :
                        // 检查设备是否已链接
                        this.$message.info("等待终端设备连接！");
                        message.sendClientConnectAndTimeSyncCheck();
                        break;
                    case ACTIVE_CLIENT_TIME_SYNC :
                        //准备启动tsi校准---弹窗确认要点烟
                        this.alertTsiTestPre(true);

                        this.unNext = true;

                        //初始化tsi状态检测器
                        tsiStatus.init(this);
                        break;
                    case ACTIVE_TSI_COLLECT_PRE :
                        //准备启动tsi测量---弹窗确认要点烟
                        this.alertTsiTestPre(false);


                        //初始化tsi状态检测器
                        tsiStatus.init(this);
                        break;
                    default :
                        break;
                }
			}
		},
		mounted() {
            var _this = this;
            //启动socket 连接
            message.connect();

            //初始化客户端显示信息
            clientLight.init();

            // 监听消息
            message.socket.onmessage = function(event) {
                var data = JSON.parse(event.data);
                console.log(data);

                switch(data.type) {
                    case message.WEB_TSI_CHECK :
                        _this.isLoading = false;
                        if (data.content == 'true') {
                            if (_this.isAllClientTimeSync) {
                                _this.active = ACTIVE_CLIENT_TIME_SYNC;
                            } else if (_this.isAllClientConnected) {
                                _this.active = ACTIVE_CLIENT_CONNECT;
                            } else {
                                _this.active = ACTIVE_TSI_CHECK;
                            }
                            _this.$message.success("TSI检查通过！");
                        } else {
                            _this.$message.error("TSI检查不通过！");
                        }
                        break;
                    case message.WEB_CLIENT_CONNECT_COMPLETE :
                        if (_this.active < ACTIVE_CLIENT_CONNECT) {
                            if (_this.active == ACTIVE_TSI_CHECK) {
                                _this.active = ACTIVE_CLIENT_CONNECT;
                            }
                            _this.$message.success("全部终端设备已连接！");
                            _this.isAllClientConnected = true;
                        }
                        break;
                    case message.WEB_CLIENT_TIME_SYNC_COMPLETE :
                        if (_this.active == ACTIVE_CLIENT_CONNECT || _this.active == ACTIVE_TSI_CHECK) {
                            _this.active = ACTIVE_CLIENT_TIME_SYNC;
                        }
                        _this.isAllClientTimeSync = true;
                        _this.$message.success("全部终端设备已时间同步！");
                        break;
                    case message.WEB_TSI_TEST_PRE :
                        _this.isLoading = false;
                        _this.active = ACTIVE_TSI_COLLECT_PRE;
                        _this.$alert('TSI校验完成！待KB写入完成，即可进行测量。',{
                            confirmButtonText: '确定'
                        });

                        break;
                    case message.WEB_CAN_START_TSI_TEST :
                        _this.$alert('全部终端KB已写入完成，可以开始TSI测量！', '写入KB', {
                            confirmButtonText: '确定'
                        });
                        console.log("可以开始tsi测试")
                        _this.unNext = false;
                        break;
                    case message.WEB_TSI_TEST :
                        _this.isLoading = false;
                        _this.active = ACTIVE_TSI_COLLECT;
                        _this.$alert('TSI测量完成，等待分析结果',{
                            confirmButtonText: '确定'
                        });
                        break;
                    case message.WEB_TSI_TEST_MODULE_RESULT :
                        _this.test_result_body_hide = false;
                        _this.tsi_data_hide = true;

                        clientLight.updateOne(data.content);
                        _this.piLight = clientLight.piAll;

                        break;
                    case message.WEB_CLIENT_LOG :
                        // 日志接受处理
                        _this.logData.unshift(data.content);
                        break;
                    case message.WEB_TSI_NOW_DATA :
                        //检测当前pm25的值
                        tsiStatus.nowData(parseInt(data.content));

                        var myDate = new Date();
                        var hour = myDate.getHours();
                        var min = myDate.getMinutes();
                        var sec = myDate.getSeconds();

                        if (sec%5 == 0) {
                            if (min%5 == 0 && myDate.getSeconds() == 1) {
                                _this.tsi_data.labels.push(hour + ":" + min);
                            } else {
                                _this.tsi_data.labels.push("")
                            }

                            // Tsi 实时数据显示
                            _this.tsi_data.datasets[0].data.push(parseInt(data.content));

                            window.tsi_data_line.update();
                        }

                        break;
                    case message.WEB_CLIENT_TREE_DATA :
                        // 终端树数据更新
                        if(data.content != null) {
                            _this.treeData = data.content;
                        }
                        break;
                    default:
                        break;
                }
            };


            //画图
            var options = {
                responsive: true,
                hover: {
                    mode: 'nearest',
                    intersect: true
                },
                scales: {
                    xAxes: [{
                        display: false,
                        scaleLabel: {
                            display: true,
                            labelString: '时间',
                            gridLines: {
                                drawOnChartArea: false, // only want the grid lines for one axis to show up
                            },
                        },
                    }],
                    yAxes: [{
                        display: true,
                        scaleLabel: {
                            display: true,
                            labelString: 'Tsi数值'
                        },
                        position: "right",

                    }]
                }
            };

            var config = {
                type: 'line',
                data: _this.tsi_data,
                options: options
            };
            var ctx = document.getElementById("tsi_data_body").getContext("2d");
            window.tsi_data_line = new Chart(ctx, config);
		}
	}

</script>