<template>
  <div class="dashboard-editor-container" >
    <el-row>
      <el-card>
        <div slot="header" class="clearfix">
          <span style="font-size:25px;"> <svg-icon icon-class="config"/>  站点设置</span>
        </div>
        <div v-loading="loading" v-if="loading">
          <br>
          <br>
          <br>
          <br>
        </div>
        <transition name="el-fade-in-linear">
          <el-row v-if="show" :gutter="8">

            <el-col :xs="24" :sm="24" :md="12" :lg="12" :xl="12" class="box">
              <!-- 基本信息 -->
              <el-card class="box-card">
                <div slot="header" class="clearfix" >
                  <span>基本信息</span>
                  <el-button style="float: right;" type="success" plain="" round="" icon="el-icon-success" @click="reset(site,'site_base')">保存</el-button>
                </div>
                <div class="box">
                  <el-row>
                    <el-input v-model="site.site_name" placeholder="请输入您的站点名称">
                      <template slot="prepend" >
                        <el-tooltip placement="top">
                          <div slot="content">此项会显示在您的首页和其他地方</div>
                          <el-button type="primary">站点名称</el-button>
                        </el-tooltip>
                      </template>
                    </el-input>
                  </el-row>

                  <el-row>
                    <el-input v-model="site.site_url" placeholder="请输入你的站点域名，需要带有 http或 https">
                      <template slot="prepend" >
                        <el-tooltip placement="top">
                          <div slot="content">最后必须要有 "/",例如<br>http://abc.com/<br>此项决定了您的本地图床的链接</div>
                          <el-button type="primary">站点链接</el-button>
                        </el-tooltip>
                      </template>
                    </el-input>
                  </el-row>

                  <el-row>
                    <el-input v-model="site.site_footer" placeholder="">
                      <template slot="prepend" >
                        <el-tooltip placement="top">
                          <div slot="content">此项将在您网站的页面底部显示</div>
                          <el-button type="primary">页脚文字</el-button>
                        </el-tooltip>
                      </template>
                    </el-input>
                  </el-row>

                  <el-row>
                    <el-input v-model="site.logo" placeholder="">
                      <template slot="prepend" >
                        <el-tooltip placement="top">
                          <div slot="content">此项将于登录等地方显示</div>
                          <el-button type="primary">Logo链接</el-button>
                        </el-tooltip>
                      </template>
                    </el-input>
                  </el-row>

                  <el-row>
                    一次性最大上传 :
                    <el-input-number v-model="site.site_upload_max_number" style="margin-left:25px;"/>
                    张
                  </el-row>

                  <el-row>
                    允许的最大图片体积 :
                    <el-input-number v-model="site.site_upload_max_size" style="margin-left:25px;"/>
                    MB
                  </el-row>

                  <el-row>
                    是否开启注册 :
                    <el-switch
                      v-model="site.allow_register"
                      active-color="#13ce66"
                      inactive-color="#ff4949"
                      style="margin-left:25px"
                    />
                  </el-row>

                  <el-row>
                    是否允许游客上传 :
                    <el-switch
                      v-model="site.allow_tourists"
                      active-color="#13ce66"
                      inactive-color="#ff4949"
                      style="margin-left:25px"
                    />
                  </el-row>

                </div>
              </el-card>
              <!-- 分发设置 -->
              <el-card class="box-card">
                <div slot="header" class="clearfix">
                  <span style="font-size:25px;">分发配置</span>
                  <el-button style="float: right;" type="success" plain="" round="" icon="el-icon-success" @click="reset(dispatch,'dispatch')">保存</el-button>
                </div>
                <div >
                  <el-row>
                    根节点图床 :
                    <el-select v-model="dispatch.root" placeholder="请选择" style="margin-left:25px">
                      <el-option
                        v-for="item in rootList"
                        :key="item.value"
                        :label="item.info"
                        :value="item.value"/>
                    </el-select>
                  </el-row>

                </div>
              </el-card>
              <!-- 邮件服务器配置 -->
              <el-card class="box-card">
                <div slot="header" class="clearfix">
                  <span style="font-size:25px;">邮件配置</span>
                  <el-button style="float: right;" type="success" plain="" round="" icon="el-icon-success" @click="reset(site,'site_base')">保存</el-button>
                </div>
                <div >
                  <el-row>
                    是否开启 SMTP 发邮件:
                    <el-switch
                      v-model="site.mail_config.status"
                      active-color="#13ce66"
                      inactive-color="#ff4949"
                      style="margin-left:25px"
                    />
                  </el-row>
                  <!-- TODO:增加淡入淡出的特效 -->
                  <transition name="el-zoom-in-center">
                    <el-card v-show="site.mail_config.status" class="box-card">
                      <el-row>
                        <el-input v-model="site.mail_config.host" placeholder="邮件服务器地址">
                          <template slot="prepend" >
                            <el-tooltip placement="top">
                              <div slot="content">例如 stmp.qq.com</div>
                              <el-button type="primary">服务器地址</el-button>
                            </el-tooltip>
                          </template>
                        </el-input>
                      </el-row>

                      <el-row>
                        <el-input v-model="site.mail_config.from" placeholder="发件昵称">
                          <template slot="prepend" >
                            <el-tooltip placement="top">
                              <div slot="content">即 来自 xxxx 的邮件</div>
                              <el-button type="primary">发件昵称</el-button>
                            </el-tooltip>
                          </template>
                        </el-input>
                      </el-row>

                      <el-row>
                        <el-input v-model="site.mail_config.user" placeholder="用户名">
                          <template slot="prepend" >
                            <el-tooltip placement="top">
                              <div slot="content">您 SMTP 服务的用户名</div>
                              <el-button type="primary">用户名</el-button>
                            </el-tooltip>
                          </template>
                        </el-input>
                      </el-row>

                      <el-row>
                        <el-input v-model="site.mail_config.pass" placeholder="密码">
                          <template slot="prepend" >
                            <el-tooltip placement="top">
                              <div slot="content">您 SMTP 服务的密码</div>
                              <el-button type="primary">密码</el-button>
                            </el-tooltip>
                          </template>
                        </el-input>
                      </el-row>
                      <el-row>
                        发送端口 :
                        <el-select v-model="site.mail_config.port" placeholder="请选择" style="margin-left:25px">
                          <el-option
                            v-for="item in mailPortList"
                            :key="item.value"
                            :label="item.info"
                            :value="item.value"/>
                        </el-select>
                      </el-row>

                    </el-card>
                  </transition>
                </div>
              </el-card>
            </el-col>

            <el-col :xs="24" :sm="24" :md="12" :lg="12" :xl="12" class="box">
              <!-- 图床基本配置 -->
              <el-card class="box-card">
                <div slot="header" class="clearfix">
                  <span>图床配置</span>
                  <el-button style="float: right;" type="success" plain="" round="" icon="el-icon-success" @click="reset(api_config,'api_option')">保存</el-button>
                </div>
                <el-row>
                  是否启用 API 上传:
                  <el-switch
                    v-model="api_config.status"
                    active-color="#13ce66"
                    inactive-color="#ff4949"
                    style="margin-left:25px"
                  />
                </el-row>
                <transition name="el-zoom-in-center">
                  <el-card v-show="api_config.status" class="box-card">
                    <el-row>
                      默认图床API :
                      <el-select v-model="api_config.api_default" placeholder="请选择" style="margin-left:25px">
                        <el-option
                          v-for="item in storeList"
                          :key="item.id"
                          :label="item.api"
                          :value="item.api"/>
                      </el-select>
                    </el-row>
                    <el-row>
                      是否开启 API 认证(推荐开启) :
                      <el-switch
                        v-model="api_config.auth"
                        active-color="#13ce66"
                        inactive-color="#ff4949"
                        style="margin-left:25px"
                      />
                    </el-row>
                  </el-card>
                </transition>
              </el-card>
              <!-- 限制设置 -->
              <el-card class="box-card">
                <div slot="header" class="clearfix">
                  <span>上传限制</span>
                  <el-button style="float: right;" type="success" plain="" round="" icon="el-icon-success" @click="reset(ip_limit,'ip_limit')">保存</el-button>
                </div>
                <div class="box">
                  <el-row>
                    是否限制上传:
                    <el-switch
                      v-model="ip_limit.status"
                      active-color="#13ce66"
                      inactive-color="#ff4949"
                      style="margin-left:25px"
                    />
                  </el-row>
                  <!-- TODO:增加淡入淡出的特效 -->
                  <transition name="el-zoom-in-center">
                    <el-card v-show="ip_limit.status" class="box-card">
                      <el-row>
                        <el-row>
                          在
                          <el-input-number v-model="ip_limit.allow_time" style="margin:0px 5px 0px 5px;"/>
                          秒内允许上传
                          <el-input-number v-model="ip_limit.allow_num" style="margin:0px 5px 0px 5px;"/>
                          张
                        </el-row>
                      </el-row>

                      <el-row>
                        如果用户上传超额，则封 IP
                        <el-input-number v-model="ip_limit.block_time" style="margin:0px 5px 0px 5px;"/>
                        秒
                      </el-row>

                      <el-row>
                        如果用户被封
                        <el-input-number v-model="ip_limit.dead_line" style="margin:0px 5px 0px 5px;"/>
                        次以后, ip 会被加入黑名单，永远不允许上传
                      </el-row>
                    </el-card>
                  </transition>
                </div>
              </el-card>
              <!-- JWT 的设置 -->
              <el-card class="box-card">
                <div slot="header" class="clearfix">
                  <span>JWT 设置</span>
                  <el-button style="float: right;" type="success" plain="" round="" icon="el-icon-success" @click="reset(site,'site_base')">保存</el-button>
                </div>
                <div class="box">
                  <el-row>
                    <el-input v-model="site.jwt_secret" placeholder="jwt_secret">
                      <template slot="prepend" >
                        <el-tooltip placement="top">
                          <div slot="content">请随意输入字符串用于认证加密</div>
                          <el-button type="primary">jwt_secret</el-button>
                        </el-tooltip>
                      </template>
                    </el-input>
                  </el-row>

                  <el-row>
                    jwt签发时长 :
                    <el-input-number v-model="site.jwt_due_time" style="margin-left:25px;"/> 小时
                  </el-row>
                </div>
              </el-card>
            </el-col>
            <!-- 证书显示 -->
            <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="box">
              <el-card class="box-card">
                <div slot="header" class="clearfix">
                  <span style="font-size:25px;">RSA证书</span>
                </div>
                <div >
                  <el-row>
                    <el-col :xs="24" :sm="24" :md="11" :lg="11" :xl="11" style="margin:10px;">
                      <el-card>
                        <div slot="header" class="clearfix">
                          <span>公钥</span>
                        </div>
                        <el-input
                          :autosize="{ minRows: 10, maxRows: 20}"
                          :disabled="true"
                          v-model="publicKey"
                          type="textarea"
                        />
                      </el-card>

                    </el-col>
                    <el-col :xs="24" :sm="24" :md="11" :lg="11" :xl="11" style="margin:10px;">
                      <el-card>
                        <div slot="header" class="clearfix">
                          <span>私钥</span>
                        </div>
                        <el-input
                          :autosize="{ minRows: 10, maxRows: 20}"
                          :disabled="true"
                          v-model="privateKey"
                          type="textarea"
                        />
                      </el-card>

                    </el-col>
                  </el-row>

                </div>
              </el-card>
            </el-col>
            <!-- 版本号显示 -->
            <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="box">
              <el-card class="box-card">
                <div slot="header" class="clearfix">
                  <span style="font-size:25px;">AUXPI 程序信息</span>
                  <el-button style="float: right;" type="primary" plain="" round="" icon="el-icon-time" >检查更新</el-button>
                </div>
                <div >
                  <el-row>
                    <el-table
                      :data="tableData"
                      stripe
                      style="width: 100%">
                      <el-table-column
                        prop="author"
                        label="作者"
                        width="180"/>
                      <el-table-column
                        prop="version"
                        label="版本"
                        width="180"/>
                      <el-table-column
                        prop="branch"
                        label="branch"/>
                      <el-table-column
                        prop="repositories"
                        label="项目地址"/>
                    </el-table>
                  </el-row>

                </div>
              </el-card>
            </el-col>
          </el-row>
        </transition>
      </el-card>
    </el-row>
  </div>
</template>

<script>
import { getRsaKey, updateOptions, getOptions } from '@/api/siteconfig'
import { getStoreList } from '@/api/image'
export default {
  name: 'SiteConfig',
  data() {
    return {
      show: false,
      tableData: [],
      storeList: [],
      selectValue: '',
      mailPortList: [
        { info: '465 (ssl)', value: '465' },
        { info: '587 (tls)', value: '587' },
        { info: '22 (无加密)', value: '22' }
      ],
      rootList: [
        { info: '本地图床', value: 'local' },
        { info: 'Gitee 仓库', value: 'gitee' },
        { info: 'Github 仓库', value: 'github' }],
      privateKey: '',
      publicKey: '',
      loading: true,
      site: null,
      api_config: null,
      ip_limit: null,
      dispatch: null
    }
  },
  watch: {
    siteBase: function() {
      console.log('hello')
    },
    config: function() {
      console.log('change')
    }
  },
  created() {
    this.loadConfig()
  },
  methods: {
    loadConfig() {
      // 获取站点配置
      const site = getOptions({ 'key': 'site_base', 'group': 'conf' }).then(resp => {
        this.site = JSON.parse(resp.data)
      })
      // api_config
      const api_config = getOptions({ 'key': 'api_option', 'group': 'conf' }).then(resp => {
        this.api_config = JSON.parse(resp.data)
      })
      // ip_limit
      const ip_limit = getOptions({ 'key': 'ip_limit', 'group': 'conf' }).then(resp => {
        this.ip_limit = JSON.parse(resp.data)
      })
      const storeList = getStoreList().then(resp => {
        this.storeList = resp.list
      })
      const dispatch = getOptions({ 'key': 'dispatch', 'group': 'conf' }).then(resp => {
        this.dispatch = JSON.parse(resp.data)
      })
      Promise.all([site, api_config, ip_limit, storeList, dispatch]).then(() => {
        console.log('allDone---->')
        console.log(this.ip_limit)

        this.show = true
        this.loading = false
      })
      getRsaKey().then(resp => {
        this.privateKey = resp.data.private_key
        this.publicKey = resp.data.public_key
      })
    },
    reset(v, p) {
      updateOptions(v, { 'key': p, 'group': 'conf' }).then(r => {
        this.$message({
          message: '修改成功',
          type: 'success'
        })
      })
    }
  }
}
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
.dashboard-editor-container {
  padding: 32px 32px 50px 32px ;
  background-color: rgb(240, 242, 245);
  .chart-wrapper {
    background: #fff;
    padding: 16px 16px 0;
    margin-bottom: 32px;
  }
}
.lineDiv{
  height: 120px;
}
.box{
  margin-top: 10px;
}
.box .el-row{
  margin: 30px;
}
.box .el-card{
  margin-top:20px;
}
</style>
