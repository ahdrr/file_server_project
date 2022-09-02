<template>
  <div>
    <!-- 边侧导航区域-->
    <el-container style="height: 790px">
      <el-container style="height: 800px; border: 1px solid #eceff0">
        <!-- 边侧导航区域-->
        <el-header width="20">
          <p class="one">♩</p>
        </el-header>
        <el-main width="">
          <v-chart class="chart" :option="option" />
          <br /><br />
          <el-upload
            class="upload-demo"
            drag
            :action="getupRealpath()"
            :headers="getupheads()"
            multiple
            :before-upload="onBeforeUpload"
            :limit="100"
            :on-exceed="handleExceed"
            :on-error="handupERR"
            :on-success="handupSuccess"
          >
            <i class="el-icon-upload"></i>
            <div class="el-upload__text">
              将文件拖到此处，或<em>点击上传</em>
            </div>
            <div class="el-upload__tip" slot="tip">文件大小限制10G</div>
          </el-upload>
        </el-main>
      </el-container>

      <!--用户列表区域-->
      <el-container
        style="height: 800px;rgb(152, 191, 33);border: 1px solid #eee"
      >
        <el-main style="text-align: right; font-size: 11px">
          <!--卡片试图区域-->
          <el-card>
            <!--搜索与添加区域-->
            <el-row :gutter="1">
              <el-col :span="11">
                <!-- 面包屑导航区域-->
                <el-breadcrumb separator=">">
                  <el-breadcrumb-item
                    v-for="item in pathurllist"
                    :key="item.index"
                    :to="{ path: item.workdir }"
                    @click.native="breadcrumbhand(item)"
                  >
                    {{ item.dirname }}
                  </el-breadcrumb-item>
                </el-breadcrumb>
              </el-col>
              <el-col :span="2" :push="11">
                <!-- 面包屑导航区域-->
                <el-button
                  class="a666"
                  type="danger"
                  round
                  size="mini"
                  @click="logout()"
                  >退出</el-button
                >
              </el-col>
            </el-row>
          </el-card>

          <!--卡片试图区域-->
          <el-card>
            <!--搜索与添加区域-->
            <el-row :gutter="20">
              <el-col :span="15">
                <el-input
                  placeholder="请输入内容"
                  v-model="queryInfo.query"
                  clearable
                  @clear="getFileList"
                >
                  <el-button
                    slot="append"
                    icon="el-icon-search"
                    @click="searchFileList"
                  ></el-button>
                </el-input>
              </el-col>
              <el-col :span="4" :push="3">
                <el-button type="primary" @click="addnewdir()"
                  >新建文件夹</el-button
                >
              </el-col>
            </el-row>
          </el-card>
          <el-table
            :data="fileList"
            height="590px"
            size="mini"
            border
            stripe
            style="width: 100%"
            :default-sort="{ prop: 'filetype', order: 'ascending' }"
          >
            <el-table-column type="index" label="序"></el-table-column>
            <el-table-column label="文件名" prop="name" width="">
              <template slot-scope="scope">
                <el-button
                  size="medium"
                  type="text"
                  :style="{
                    color: scope.row.filetype == 'dir' ? '#409EFF' : '#474242'
                  }"
                  @click="handleTopath(scope.row)"
                  >{{ scope.row.name }}</el-button
                >
              </template>
            </el-table-column>

            <el-table-column
              label="类型"
              prop="filetype"
              width="70px"
            ></el-table-column>
            <el-table-column
              label="创建日期"
              prop="time"
              width="160px"
            ></el-table-column>
            <el-table-column
              label="大小"
              prop="size"
              width="90px"
            ></el-table-column>
            <el-table-column label="操作" width="190px">
              <template slot-scope="scope">
                <!-- 修改按钮 -->
                <el-button
                  type="primary"
                  icon="el-icon-edit"
                  size="mini"
                  @click="renamefile(scope.row)"
                ></el-button>
                <!-- 下载按钮 -->
                <el-button
                  type="primary"
                  icon="el-icon-download"
                  size="mini"
                  @click="downloadfile(scope.row)"
                ></el-button>
                <!-- 删除按钮 -->
                <el-button
                  type="danger"
                  icon="el-icon-delete"
                  size="mini"
                  @click="removefile(scope.row)"
                ></el-button>
              </template>
            </el-table-column>
          </el-table>
          <!--分页区域-->
          <el-footer style="text-align: center; font-size: 0px">
            <el-pagination
              @size-change="handleSizeChange"
              @current-change="handleCurrentChange"
              :current-page="queryInfo.pagenum"
              :page-sizes="[1, 2, 5, 10]"
              :page-size="queryInfo.pagesize"
              layout="total, sizes, prev, pager, next, jumper"
              :total="total"
            >
            </el-pagination>
          </el-footer>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script>
import axios from '../axios.js'
import VChart, { THEME_KEY } from 'vue-echarts'

export default {
  components: {
    VChart
  },
  provide: {
    [THEME_KEY]: 'dark'
  },
  data () {
    return {
      option: {
        backgroundColor: '#FFFFFF', // rgba设置透明度0.1
        title: {
          text: 'Disk Info',
          left: 'center',
          textStyle: {
            fontSize: 20,
            fontWeight: 'bolder',
            color: '#333' // 主标题文字颜色
          },
          subtextStyle: {
            fontSize: 18,
            fontWeight: 'bolder',
            color: '#333' // 主标题文字颜色
          }

        },
        tooltip: {
          trigger: 'item',
          formatter: '{a} <br/>{b} : {c} ({d}%)'
        },
        legend: {
          orient: 'vertical',
          left: 'left',
          textStyle: {
            'color': 'inherit'
          },
          data: [
            'used',
            'free'
          ]
        },
        series: [
          {
            name: 'Disk info',
            type: 'pie',
            radius: '60%',
            center: ['50%', '65%'],
            data: [{
              value: '', name: 'used', convertvalue: '', Percent: '', itemStyle: { color: '#F56C6C' }
            },
            {
              value: '', name: 'free', convertvalue: '', Percent: '', itemStyle: { color: '#409EFF' }
            }], // .sort(function (a, b) { return a.value - b.value }), // lable排序
            label: {
              'formatter': function formatterFunc (params) {
                const values = params.data // 内容
                const formatter = [`{rect|}{name|${values.convertvalue}} {value|${values.Percent}%}`, `{value|${values.Percent}%} {name|${values.convertvalue}}{rect|}`]
                const midAngle = (values._startArc + values._endArc) / 2
                if (midAngle <= Math.PI) {
                  return formatter[0]
                } else {
                  return formatter[1]
                }
              },
              'rich': {
                'value': {
                  'color': 'inherit'
                },
                'name': {
                  'color': 'inherit',
                  'borderColor': 'inherit',
                  'borderWidth': 1,
                  'padding': [2, 2],
                  'height': 10,
                  'width': '55%'
                },
                'rect': {
                  'height': 8,
                  // 'width': '5%',
                  'backgroundColor': 'inherit'
                }
              }
            },

            labelLine: {
              lineStyle: {
                color: '#b4adad'
              },
              smooth: 0.3,
              length: 5,
              length2: 15
            },
            emphasis: {
              itemStyle: {
                shadowBlur: 10,
                shadowOffsetX: 0,
                shadowColor: 'rgba(0, 0, 0, 0.5)'
              }
            }
          }
        ]
      },
      upfileList: [],
      avtivecolor: '#409eff',
      pathurllist: [
        {
          'workdir': '/',
          'dirname': 'basedir'
        }
      ],
      // path 定位
      requesturl: this.$route.path,
      // 获取用户列表的参数对象
      queryInfo: {
        query: '',
        // 当前的页数
        pagenum: 1,
        // 当前每页显示的条数
        pagesize: 10
      },
      fileList: [],
      total: 0,
      editForm: {}
    }
  },
  watch: {
    $route () {
      let thisrequesturl = this.getRealpath()
      this.getFileList(thisrequesturl)
    }
  },
  created () {
    this.getFileList()
    this.setChart()
    this.parshPathurl()
  },
  methods: {
    // searchFileList

    searchFileList () {
      this.$message.success('搜索功能暂未实现')
    },
    // 解析路由地址
    parshPathurl () {
      let realpathlist = this.getRealpath().split('/')
      let fdir = ''
      for (let i in realpathlist) {
        if (!(realpathlist[i])) {
          continue
        }
        fdir = fdir + '/' + realpathlist[i]
        this.pathurllist.push({
          'dirname': realpathlist[i],
          'workdir': fdir
        })
      }
    },

    // 存储单位转换
    byteConvert (bytes) {
      if (isNaN(bytes)) {
        return ''
      }
      var symbols = ['B', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB']
      var exp = Math.floor(Math.log(bytes) / Math.log(2))
      if (exp < 1) {
        exp = 0
      }
      var i = Math.floor(exp / 10)
      bytes = bytes / Math.pow(2, 10 * i)

      if (bytes.toString().length > bytes.toFixed(2).toString().length) {
        bytes = bytes.toFixed(0)
      }
      return bytes + symbols[i]
    },
    // 设置饼形图属性
    setChart () {
      axios.getdiskinfo().then((response) => {
        if (response.data.code === 200) {
          let diskused = response.data.data.diskuse
          let diskfree = response.data.data.diskfree
          let usedPercent = response.data.data.usedPercent
          this.option.series[0].data = [
            {
              value: diskused, name: 'used', convertvalue: this.byteConvert(diskused), Percent: usedPercent, itemStyle: { color: '#F56C6C' }
            },
            {
              value: diskfree, name: 'free', convertvalue: this.byteConvert(diskfree), Percent: 100 - usedPercent, itemStyle: { color: '#409EFF' }
            }
          ]
        }
      }
      )
    },
    // 新建文件夹
    async addnewdir () {
      this.$prompt('请输入文件名称', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        inputPattern: /^\w+$/,
        inputErrorMessage: '文件格式不正确'
      }).then(({ value }) => {
        let realpath = this.getRealpath()
        let data = JSON.stringify({
          'newname': value
        })
        axios.createnewdir(realpath, data
        ).then((response) => {
          if (response.data.code === 200) {
            this.$message.success('文件夹新建成功')
            this.getFileList()
          } else {
            this.$message.error('文件夹新建失败: ' + response.data.msg)
          }
        }
        )
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '取消输入'
        })
      })
    },
    // 导航栏
    breadcrumbhand (breaditem) {
      let dirindex = this.pathurllist.indexOf(breaditem)
      this.pathurllist = this.pathurllist.slice(0, dirindex + 1)
    },
    // 获取url路径
    getRealpath () {
      let realpath = this.$route.path.replace(/\/frontpage\/*/g, '')
      return realpath
    },
    // 获取文件列表
    getFileList (thisrequesturl) {
      if (!thisrequesturl) {
        thisrequesturl = this.getRealpath()
      }
      axios.getFiles(thisrequesturl).then((response) => {
        if (response.status === 200) {
          this.fileList = response.data.data
        }
        // if (response.status === 401) {
        // 不成功跳转回登录页
        //  this.$router.push('/login')
        // 并且清除掉这个token
        // this.$store.dispatch('UserLogout')
        // }
      }
      ).catch((response) => {
        if (response.status !== 401) {
          this.$message({
            type: 'error',
            message: '后端请求异常'
          })
        }
      })
    },
    // 监听 pagesize 改变的事件
    handleSizeChange (newSize) {
      this.queryInfo.pagesize = newSize
      this.getFileList()
      console.log(newSize)
    },
    // 监听页码值改变的事件
    handleCurrentChange (newPage) {
      this.queryInfo.pagenum = newPage
      this.getFileList()
    },
    // 导航栏跳转地址
    handleTopath (row) {
      if (row.filetype === 'dir') {
        let pushurl = this.$route.path
        if (pushurl === '/') {
          pushurl = pushurl + row.name
        } else {
          pushurl = pushurl + '/' + row.name
        }
        this.pathurllist.push({
          'dirname': row.name,
          'workdir': pushurl
        })
        this.$router.push(pushurl)
      } else {
        console.log('下载文件')
      }
    },
    // 下载文件
    async downloadfile (row) {
      if (row.filetype === 'dir') {
        this.$message.error('暂时不支持目录下载')
        return
      }
      let realpath = this.getRealpath(row.name)
      axios.download(realpath, { responseType: 'blob' }).then((response) => {
        const { data, headers } = response
        // const fileName = headers['content-disposition'].replace(/\w+;filename=(.*)/, '$1')
        const fileName = row.name
        // 此处当返回json文件时需要先对data进行JSON.stringify处理，其他类型文件不用做处理
        // const blob = new Blob([JSON.stringify(data)], ...)
        const blob = new Blob([data], { type: headers['content-type'] })
        let dom = document.createElement('a')
        let url = window.URL.createObjectURL(blob)
        dom.href = url
        dom.download = decodeURI(fileName)
        dom.style.display = 'none'
        document.body.appendChild(dom)
        dom.click()
        dom.parentNode.removeChild(dom)
        window.URL.revokeObjectURL(url)
      }
      )
    },
    // 删除文件
    async renamefile (row) {
      await this.$prompt('请输入新名称', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        inputPattern: /^\w+$/,
        inputErrorMessage: '邮箱格式不正确'
      }).then(({ value }) => {
        let realpath = this.getRealpath() + '/' + row.name
        let data = JSON.stringify({
          'newname': value
        })
        axios.rename(realpath, data
        ).then((response) => {
          if (response.data.code === 200) {
            this.$message.success('重命名成功')
            this.getFileList()
          } else {
            this.$message.error('重命名失败 ' + response.data.msg)
          }
        }
        )
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '取消输入'
        })
      })
    },
    // 删除文件
    async removefile (row) {
      // 弹框提示用户是否删除
      const confirmResult = await this.$confirm(
        '此操作将永久删除' + row.name + ', 是否继续?',
        '提示',
        {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }
      ).catch(err => err)
      // 如果用户确认删除，则返回字符串 confirm
      // 如果用户取消删除，则返回字符串 cancel
      if (confirmResult !== 'confirm') {
        return this.$message.info('已取消删除')
      }
      // 发送删除请求
      let realpath = this.getRealpath() + '/' + row.name
      axios.removefile(realpath).then((response) => {
        if (response.data.code === 200) {
          this.$message.success('删除成功')
          this.getFileList()
        } else {
          this.$message.error('删除失败: ' + response.data.msg)
        }
      }
      )
    },
    // 上传文件前校验
    onBeforeUpload (file) {
      const isLt10G = file.size / 1024 / 1024 / 1024 < 10
      if (!isLt10G) {
        this.$message.error('上传文件大小不能超过 10G!')
      }
      return isLt10G
    },
    // 文件上传数量超出触发函数
    handleExceed (files, fileList) {
      this.$message.warning(`当前限制选择 100 个文件，本次选择了 ${files.length} 个文件，共选择了 ${files.length + fileList.length} 个文件`)
    },
    // 获取上传文件接口
    getupRealpath () {
      let realpath = this.$route.path.replace(/\/frontpage\/*/g, '')
      return '/api/up' + realpath + '/'
    },
    // 获取当前token
    getupheads () {
      let token = this.$store.state.token
      return { token: token }
    },
    // 上传失败回调
    handupERR () {
      this.$message.error('上传失败')
    },
    // 上传成功回调
    handupSuccess (a, b, c) {
      this.getFileList()
      if (c.length > 3) {
        c.shift()
      }
    },
    logout () {
      this.$router.push('/login')
      // 并且清除掉这个token
      this.$store.dispatch('UserLogout')
    }
  }
}
</script>

<style>
.el-header {
  background-color: #eceff0;
  --color: #b4adad;
  --line-height: 70px;
}

.el-breadcrumb__inner.is-link {
  font-size: 17px;
  font-weight: 700 !important;
  --background: 设置背景色;
  --background: #000000;
  --color: rgb(255, 255, 255) !important;
  color: rgb(40, 40, 218);
}
p.one {
  border-style: groove;
  border-width: 5px;
  border-color: #cbccc7;
  background: #eceff0;
}

.chart {
  height: 210px;
  width: 320px;
}

.el-upload-dragger {
  height: 200px;
  width: 290px;
}

.a666 {
  padding-top: 0px;
  padding-right: 0px;
  padding-bottom: 0px;
  padding-left: 0px;
}
</style>
