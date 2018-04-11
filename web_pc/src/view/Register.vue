<style scoped lang="scss">
    .form-box{
        width: 50%;
        margin: 0 auto;
        padding-top: 30%;

        .el-button{
            width: 100%;
        }
    }
</style>

<template>
    <div class="form-box">
        <el-form ref="form" :model="form" :rules="rules" label-width="80px">
            <el-form-item label="邮箱" prop="username">
                <el-input v-model="form.username"></el-input>
            </el-form-item>
            <el-form-item label="密码" prop="password">
                <el-input v-model="form.password" type="password"></el-input>
            </el-form-item>
            <el-form-item size="large">
                <el-button type="primary" @click="register">立即创建</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                form: {
                    username: "",
                    password: "",
                },

                rules: {
                    username: [
                        {type: 'email', required: true, message: '请输入正确邮箱', trigger: 'blur' },
                        { min: 1, max: 30, message: '用户名超长', trigger: 'blur' }
                    ],
                    password: [
                        { required: true, message: '请输入密码', trigger: 'blur' },
                        { min: 6, max: 30, message: '密码长度 6-30', trigger: 'blur' }
                    ],
                }
            }
        },

        created() {

        },

        mounted() {

        },

        methods: {
            register(){
                this.$refs["form"].validate((valid) => {
                    if (!valid) {
                        return false;
                    }

                    this.$post("sendRegisterMail", {
                        username: this.form.username,
                        password: this.$md5(this.form.password),
                    }).then(res => {

                    })

                });

            }
        },

        components: {}
    }
</script>
