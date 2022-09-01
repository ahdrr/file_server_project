<!-- eslint-disable indent -->
<!-- eslint-disable indent -->
<!-- eslint-disable indent -->
<!-- eslint-disable indent -->
<template>
  <div class="login">
    <img src="../assets/logo.png" />
    <el-tabs v-model="activeName" @tab-click="handleClick">
      <el-tab-pane label="" name="first">
        <el-form
          :model="ruleForm"
          :rules="rules"
          ref="ruleForm"
          label-width="60px"
          class="demo-ruleForm"
        >
          <el-form-item label="名称" prop="username">
            <el-input v-model="ruleForm.username"></el-input>
          </el-form-item>
          <el-form-item label="密码" prop="password">
            <el-input
              type="password"
              v-model="ruleForm.password"
              auto-complete="off"
            ></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="submitForm('ruleForm')"
              >登录</el-button
            >
            <el-button @click="resetForm('ruleForm')">重置</el-button>
          </el-form-item>
        </el-form>
      </el-tab-pane>
      <!--el-tab-pane label="注册" name="second">
        <register></register>
      </el-tab-pane-->
    </el-tabs>
  </div>
</template>
<script>
import register from '@/components/register'
import axios from '../axios.js'
export default {
  data () {
    var validatePass = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入密码'))
      } else {
        if (this.ruleForm.checkPass !== '') {
          this.$refs.ruleForm.validateField('checkPass')
        }
        callback()
      }
    }
    return {
      activeName: 'first',
      ruleForm: {
        username: '',
        password: '',
        checkPass: ''
      },
      rules: {
        username: [
          { required: true, message: '请输入您的名称', trigger: 'blur' },
          { min: 2, max: 100, message: '长度在 2 到 5 个字符', trigger: 'blur' }
        ],
        password: [{ required: true, validator: validatePass, trigger: 'blur' }]
      }
    }
  },
  methods: {
    // 选项卡切换
    handleClick (tab, event) { },
    // 重置表单
    resetForm (formName) {
      this.$refs[formName].resetFields()
    },
    // 提交表单
    submitForm (formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          axios.userLogin(this.ruleForm).then(({ data }) => {
            // 账号不存在

            if (data.code !== 200) {
              this.$message({
                type: 'info',
                message: '账号不存在'
              })
              return
            }
            if (data.code === 200) {
              this.$message({
                type: 'success',
                message: '登录成功'
              })
              // 拿到返回的token和username，并存到store
              let token = data.data.token
              let username = data.data.username
              this.$store.dispatch('UserLogin', token)
              this.$store.dispatch('UserName', username)
              // 跳到目标页
              this.$router.push('frontpage')
            }
          }).catch(() => {
            this.$message({
              type: 'error',
              message: '后端请求异常'
            })
          })
        } else {
          console.log('error submit!!')
          return false
        }
      })
    }
  },
  components: {
    register
  }
}
</script>
<style rel="stylesheet/scss" lang="scss">
.login {
  width: 400px;
  margin: 0 auto;
}

#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 70px;
}

.el-tabs__item {
  text-align: center;
  width: 70px;
}
.el-tabs__active-bar.is-top {
  color: #00050a !important;
  background-color: #e4e7ed !important;
}
</style>
