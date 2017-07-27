/**
 * Created by wchuang on 7/14/2017.
 */
import { Alert } from 'element-ui'

export var tsiStatus = {
    vue : null,

    TSI_START_POINT : 500,
    TSI_ALERT_START_CLEANER : 200,
    TSI_STOP_POINT : 14,
    TSI_FLAG_TIMES : 3, //多少次出现时，才认为有效

    isRunning: false,
    startPointTimes : 0,
    startCalculateTimes : 0,
    isAlertCloseSmoke : false,
    isAlertStartCleaner : false,
    isAlertCloseCleaner : false,
    init : function(vueObject) {
        this.vue = vueObject;

        this.isRunning=false;
        this.startPointTimes =0;
        this.startCalculateTimes = 0;
        this.isAlertCloseSmoke = false;
        this.isAlertStartCleaner = false;
        this.isAlertCloseCleaner = false;
    },

    nowData : function(pm25) {
        if (this.isRunning) {
            if (pm25 > this.TSI_START_POINT) {
                if (pm25 > 530) {
                    if (!this.isAlertCloseSmoke) {
                        this.isAlertCloseSmoke = true;
                        this.vue.$alert('pm25超过500，请灭烟！',{
                            confirmButtonText: '确定'
                        });
                    }
                }

                //大于500  舍弃
            } else if (pm25 > this.TSI_STOP_POINT) {
                //14 ~ 500 之间的数据 统计
                if (this.startCalculateTimes >= this.TSI_FLAG_TIMES) {
                    //判定有效

                    if (pm25 <= this.TSI_ALERT_START_CLEANER) {
                        if (!this.isAlertStartCleaner) {
                            this.isAlertStartCleaner = true;
                            this.vue.$alert('pm25已降到200，请开启空气净化器！',{
                                confirmButtonText: '确定'
                            });
                        }
                    }

                } else {
                    this.startCalculateTimes++;
                }
            } else if (pm25 > 0 ) {
                if (!this.isAlertCloseCleaner) {
                    this.isAlertCloseCleaner = true;
                    this.vue.$alert('pm25已降到14，请关闭空气净化器！',{
                        confirmButtonText: '确定'
                    });
                }
            } else {
                //小于等于0 弃用
            }
        } else {
            if (pm25 > this.TSI_START_POINT) {
                // 准备开始记录，待下降到500
                if (this.startPointTimes >= this.TSI_FLAG_TIMES) {
                    this.isRunning = true;
                } else {
                    this.startPointTimes++;
                }
            } else {
                //小于500的舍弃
            }
        }

    },
};