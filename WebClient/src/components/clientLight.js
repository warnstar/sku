/**
 * Created by wchuang on 7/7/2017.
 */

import { message } from '../socket/message';

export var clientLight = {
    piAll : [],
    piOne : {
            pid : "无",
            modules : []
        }
    ,
    moduleOne : {
        module_id : null,
        status : "off",
        info : [
            {
                stage : "",
                total : 0,
                error : 0,
                proportion : 0
            }
        ]
    },
    init : function() {

        let option = message.getOption();

        // let piOne = this.piOne;
        //
        // for (let j = 1; j <= option.client_module_num; j++) {
        //     let moduleOne = this.moduleOne;
        //     piOne.modules.push(moduleOne);
        // }
        //
        // for (let i = 1; i <= option.client_num; i++) {
        //     this.piAll.push(piOne);
        // }

    },
    updateOne : function(pi) {
        console.log(this.piAll)
        var _this = this;
        this.piAll.push(pi)
    },
    arrayObjSort : function(arr) {
        for (let i = 1; i < arr.length ; i++ ) {
            for (let j = i+1; j < arr.length; j++) {

            }
        }
    }
}