<template>

    <el-row type="flex" class="row-bg control_top" justify="center" style="margin-top: 75px;">
        <el-col :span="15" class="row-bg">
            <el-form ref="form"  :model="form" :rules="formRule" label-width="120px" style="margin:20px;width:20%;min-width:400px;">
                <el-form-item label="服务器地址" prop="host">
                    <el-input v-model="form.host"></el-input>
                </el-form-item>

                <el-form-item label="服务器端口" prop="port">
                    <el-input v-model.number="form.port"></el-input>
                </el-form-item>

                <el-form-item label="测试派数量" prop="client_num">
                    <el-input v-model.number="form.client_num"></el-input>
                </el-form-item>

                <el-form-item label="最大模块数量" prop="client_num">
                    <el-input v-model.number="form.client_module_num"></el-input>
                </el-form-item>

                <el-form-item label="TSI服务器地址" prop="tsi_host">
                    <el-input v-model="form.tsi_host"></el-input>
                </el-form-item>

                <el-form-item>
                    <el-button type="primary" @click.native.prevent="onSubmit">设置</el-button>
                    <el-button @click="$router.back(-1)">返回</el-button>
                </el-form-item>
            </el-form>
        </el-col>
    </el-row>

</template>

<script>
    import { message } from '../../socket/message';

    export default {
        data() {
            return {
                form: {
                    host: "",
                    tsi_host:"",
                    port: 0,
                    client_num: 0,
                    client_module_num : 0,
                },
                formRule: {
                    host : [
                        { required: true }
                    ],
                    tsi_host : [
                        { required: true }
                    ],
                    port : [
                        { required: true },
                        { type: 'number' }
                    ],
                    client_num : [
                        { required: true },
                        { type: 'number' }
                    ],
                    client_module_num : [
                        { required: true },
                        { type: 'number' }
                    ]
                }
            }
        },
        mounted() {
            var option = message.getOption();
            this.form.host = option.host;
            this.form.tsi_host = option.tsi_host;
            this.form.port = parseInt(option.port);
            this.form.client_num = parseInt(option.client_num);
            this.form.client_module_num = parseInt(option.client_module_num);
        },
        methods: {
            onSubmit() {
                var _this = this;
                this.$refs.form.validate(function(valid){
                    if (valid) {
                        message.setServerAddress(_this.form);
                        _this.$message.success("操作成功");
                    }
                });
            }
        }
    }

</script>