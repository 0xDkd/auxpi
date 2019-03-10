<template>
  <div class="dashboard-editor-container">

    <el-row v-if="loadConfig" :gutter="8">

      <el-col :xs="24" :sm="24" :md="12" :lg="12" :xl="12" class="box">
        <!-- 基本信息 -->
        <el-card class="box-card">
          <div slot="header" class="clearfix" >
            <span>基本信息</span>
            <el-button style="float: right;" type="success" plain="" round="" icon="el-icon-success" @click="reset(siteBase)">保存</el-button>
          </div>
          <div class="box">
            <el-row>
              <el-input v-model="siteBase.site_name" placeholder="请输入您的站点名称">
                <template slot="prepend" >
                  <el-tooltip placement="top">
                    <div slot="content">此项会显示在您的首页和其他地方</div>
                    <el-button type="primary">站点名称</el-button>
                  </el-tooltip>
                </template>
              </el-input>
            </el-row>

            <el-row>
              <el-input v-model="siteBase.site_url" placeholder="请输入你的站点域名，需要带有 http或 https">
                <template slot="prepend" >
                  <el-tooltip placement="top">
                    <div slot="content">最后必须要有 "/",例如<br>http://abc.com/<br>此项决定了您的本地图床的链接</div>
                    <el-button type="primary">站点链接</el-button>
                  </el-tooltip>
                </template>
              </el-input>
            </el-row>

            <el-row>
              <el-input v-model="siteBase.site_footer" placeholder="">
                <template slot="prepend" >
                  <el-tooltip placement="top">
                    <div slot="content">此项将在您网站的页面底部显示</div>
                    <el-button type="primary">页脚文字</el-button>
                  </el-tooltip>
                </template>
              </el-input>
            </el-row>

            <el-row>
              <el-input v-model="siteBase.logo" placeholder="">
                <template slot="prepend" >
                  <el-tooltip placement="top">
                    <div slot="content">此项将于登录等地方显示</div>
                    <el-button type="primary">Logo链接</el-button>
                  </el-tooltip>
                </template>
              </el-input>
            </el-row>

            <el-row>
              是否开启注册 :
              <el-switch
                v-model="siteBase.allow_register"
                active-color="#13ce66"
                inactive-color="#ff4949"
                style="margin-left:25px"
              />
            </el-row>

            <el-row>
              是否允许游客上传 :
              <el-switch
                v-model="siteBase.allow_tourists"
                active-color="#13ce66"
                inactive-color="#ff4949"
                style="margin-left:25px"
              />
            </el-row>

          </div>
        </el-card>
        <!-- 本地图床配置 -->
        <el-card class="box-card">
          <div slot="header" class="clearfix">

            <span style="font-size:25px;">本地图床配置</span>
            <el-button style="float: right;" type="success" plain="" round="" icon="el-icon-success" @click="reset(local)">保存</el-button>
          </div>
          <div >
            <el-row>
              是否开启本地图床:
              <el-switch
                v-model="local.site_upload_way.local_store.open"
                active-color="#13ce66"
                inactive-color="#ff4949"
                style="margin-left:25px"
              />
            </el-row>
            <!-- TODO:增加淡入淡出的特效 -->
            <transition name="el-zoom-in-center">
              <el-card v-show="local.site_upload_way.local_store.open" class="box-card">
                <el-row>
                  <el-button round>点我查看配置教程</el-button>
                </el-row>
                <el-row>
                  <el-input v-model="local.site_upload_way.local_store.link" placeholder="请输入您的默认链接地址">
                    <template slot="prepend" >
                      <el-tooltip placement="top">
                        <div slot="content">开头必须要有"/",最后不能有 "/"</div>
                        <el-button type="primary">软路径</el-button>
                      </el-tooltip>
                    </template>
                  </el-input>
                </el-row>

                <el-row>
                  <el-input v-model="local.site_upload_way.local_store.storage_location" placeholder="请输入储存路径">
                    <template slot="prepend" >
                      <el-tooltip placement="top">
                        <div slot="content">开头不能有"/",最后也不能有"/"<br>也可以直接填写绝对地址</div>
                        <el-button type="primary">储存路径</el-button>
                      </el-tooltip>
                    </template>
                  </el-input>
                </el-row>
              </el-card>
            </transition>
          </div>
        </el-card>
        <!-- 数据库配置 -->
        <el-card class="box-card">
          <div slot="header" class="clearfix">

            <span style="font-size:25px;">数据库配置</span>
            <el-button style="float: right;" type="success" plain="" round="" icon="el-icon-success" @click="reset(database)">保存</el-button>
          </div>
          <div >
            <el-row>
              <el-alert
                title="请注意，如果您不清楚数据库的信息请不要随意修改，这可能导致无法访问后台"
                type="warning"
                center
                show-icon/>
            </el-row>
            <el-card class="box-card">

              <el-row>
                <el-input v-model="database.db_option.db_type" :disabled="true">
                  <template slot="prepend" >
                    <el-tooltip placement="top">
                      <div slot="content">暂时仅支持 MySQL</div>
                      <el-button type="primary">数据库驱动</el-button>
                    </el-tooltip>
                  </template>
                </el-input>
              </el-row>

              <el-row>
                <el-input v-model="database.db_option.db_host" placeholder="数据库服务器地址">
                  <template slot="prepend" >
                    <el-tooltip placement="top">
                      <div slot="content">数据库服务器地址</div>
                      <el-button type="primary">服务器地址</el-button>
                    </el-tooltip>
                  </template>
                </el-input>
              </el-row>

              <el-row>
                <el-input v-model="database.db_option.db_name" placeholder="数据库名称">
                  <template slot="prepend" >
                    <el-tooltip placement="top">
                      <div slot="content">您安装程序所用的数据库</div>
                      <el-button type="primary">数据库名称</el-button>
                    </el-tooltip>
                  </template>
                </el-input>
              </el-row>

              <el-row>
                <el-input v-model="database.db_option.db_user" placeholder="用户名">
                  <template slot="prepend" >
                    <el-tooltip placement="top">
                      <div slot="content">链接数据库所使用的用户名</div>
                      <el-button type="primary">用户名</el-button>
                    </el-tooltip>
                  </template>
                </el-input>
              </el-row>

              <el-row>
                <el-input v-model="database.db_option.db_user" placeholder="密码">
                  <template slot="prepend">
                    <el-tooltip placement="top">
                      <div slot="content">链接数据库所使用的密码</div>
                      <el-button type="primary">密码</el-button>
                    </el-tooltip>
                  </template>
                </el-input>
              </el-row>

              <el-row>
                <el-input v-model="database.db_option.table_prefix" placeholder="表前缀">
                  <template slot="prepend" >
                    <el-tooltip placement="top">
                      <div slot="content">数据表前缀</div>
                      <el-button type="primary">表前缀</el-button>
                    </el-tooltip>
                  </template>
                </el-input>
              </el-row>

            </el-card>

          </div>
        </el-card>
        <!-- 邮件服务器配置 -->
        <el-card class="box-card">
          <div slot="header" class="clearfix">

            <span style="font-size:25px;">邮件配置</span>
            <el-button style="float: right;" type="success" plain="" round="" icon="el-icon-success" @click="reset(local)">保存</el-button>
          </div>
          <div >
            <el-row>
              是否开启 SMTP 发邮件:
              <el-switch
                v-model="local.mail_config.status"
                active-color="#13ce66"
                inactive-color="#ff4949"
                style="margin-left:25px"
              />
            </el-row>
            <!-- TODO:增加淡入淡出的特效 -->
            <transition name="el-zoom-in-center">
              <el-card v-show="local.mail_config.status" class="box-card">
                <el-row>
                  <el-input v-model="local.mail_config.host" placeholder="邮件服务器地址">
                    <template slot="prepend" >
                      <el-tooltip placement="top">
                        <div slot="content">例如 stmp.qq.com</div>
                        <el-button type="primary">服务器地址</el-button>
                      </el-tooltip>
                    </template>
                  </el-input>
                </el-row>

                <el-row>
                  <el-input v-model="local.mail_config.from" placeholder="发件昵称">
                    <template slot="prepend" >
                      <el-tooltip placement="top">
                        <div slot="content">即 来自 xxxx 的邮件</div>
                        <el-button type="primary">发件昵称</el-button>
                      </el-tooltip>
                    </template>
                  </el-input>
                </el-row>

                <el-row>
                  <el-input v-model="local.mail_config.user" placeholder="用户名">
                    <template slot="prepend" >
                      <el-tooltip placement="top">
                        <div slot="content">您 SMTP 服务的用户名</div>
                        <el-button type="primary">用户名</el-button>
                      </el-tooltip>
                    </template>
                  </el-input>
                </el-row>

                <el-row>
                  <el-input v-model="local.mail_config.pass" placeholder="密码">
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
                  <el-select v-model="local.mail_config.port" placeholder="请选择" style="margin-left:25px">
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
            <el-button style="float: right;" type="success" plain="" round="" icon="el-icon-success" @click="reset(upload)">保存</el-button>
          </div>
          <div>
            <el-row>
              默认图床API :
              <el-select v-model="upload.api_default" placeholder="请选择" style="margin-left:25px">
                <el-option
                  v-for="item in storeList"
                  :key="item.id"
                  :label="item.api"
                  :value="item.api"/>
              </el-select>
            </el-row>
            <!-- TODO:可拖动排序对侧栏进行排序 -->

            <el-row>
              一次性最大上传 :
              <el-input-number v-model="upload.site_upload_max_number" style="margin-left:25px;"/> 张
            </el-row>

            <el-row>
              允许的最大图片体积 :
              <el-input-number v-model="upload.site_upload_max_size" style="margin-left:25px;"/> MB
            </el-row>
          </div>

        </el-card>
        <!-- 新浪图床配置 -->

        <el-card class="box-card">
          <div slot="header" class="clearfix">

            <span style="font-size:25px;">微博图床配置</span>
            <el-button style="float: right;" type="success" plain="" round="" icon="el-icon-success" @click="reset(sina)">保存</el-button>
          </div>
          <div >
            <el-row>
              是否开启微博图床:
              <el-switch
                v-model="sina.site_upload_way.open_sina_pic_store"
                active-color="#13ce66"
                inactive-color="#ff4949"
                style="margin-left:25px"
              />
            </el-row>
            <transition name="el-zoom-in-center">
              <!-- TODO:增加淡入淡出的特效 -->
              <el-card v-if="sina.site_upload_way.open_sina_pic_store" class="box-card">
                <el-row>
                  <el-input v-model="sina.site_upload_way.sina_account.user_name" placeholder="请输入用户名">
                    <template slot="prepend" >
                      <el-tooltip placement="top">
                        <div slot="content">您的微博登录用来登录的用户名</div>
                        <el-button type="primary">微博用户名</el-button>
                      </el-tooltip>
                    </template>
                  </el-input>
                </el-row>

                <el-row>
                  <el-input v-model="sina.site_upload_way.sina_account.pass_word" placeholder="请输入密码">
                    <template slot="prepend" >
                      <el-tooltip placement="top">
                        <div slot="content">您的微博登录用来登录的密码</div>
                        <el-button type="primary">微博密码</el-button>
                      </el-tooltip>
                    </template>
                  </el-input>
                </el-row>

                <el-row>
                  默认图片大小 :
                  <el-select v-model="sina.site_upload_way.sina_account.defult_pic_size" placeholder="请选择" style="margin-left:25px" @change="resetLink(sina.site_upload_way.sina_account.defult_pic_size)">
                    <el-option
                      v-for="item in sinaOptions"
                      :key="item.size"
                      :label="item.size"
                      :value="item.size"/>
                  </el-select>
                  <el-card shadow="always">
                    链接效果 (点击可以直接查看图片) :<br><br>
                    <a :href="sinaEx" target="_blank">{{ sinaEx }}</a>
                  </el-card>
                </el-row>

                <el-row>
                  <span> Cookie 缓存时间:</span>

                  <el-select v-model="sina.site_upload_way.sina_account.reset_sina_cookie_time" clearable placeholder="请选择" style="margin-left:25px">
                    <el-option
                      v-for="item in sinaTimeOptions"
                      :key="item.value"
                      :label="item.size"
                      :value="item.value"/>
                  </el-select>
                </el-row>
              </el-card>
            </transition>
          </div>
        </el-card>
        <!-- Flickr 图床配置 -->
        <el-card class="box-card">
          <div slot="header" class="clearfix">

            <span style="font-size:25px;">Flickr图床配置</span>
            <el-button style="float: right;" type="success" plain="" round="" icon="el-icon-success" @click="reset(flickr)">保存</el-button>
          </div>
          <div >
            <el-row>
              是否开启Flickr图床:
              <el-switch
                v-model="flickr.site_upload_way.open_flickr_store"
                active-color="#13ce66"
                inactive-color="#ff4949"
                style="margin-left:25px"
              />
            </el-row>
            <!-- TODO:增加淡入淡出的特效 -->
            <transition name="el-zoom-in-center">
              <el-card v-if="flickr.site_upload_way.open_flickr_store" class="box-card">

                <el-row>
                  <el-input v-model="flickr.site_upload_way.flickr_account.id" placeholder="Flickr ID">
                    <template slot="prepend" >
                      <el-tooltip placement="top">
                        <div slot="content">Flickr ID 及 Flickr 分配给您的用户 ID</div>
                        <el-button type="primary">ID</el-button>
                      </el-tooltip>
                    </template>
                  </el-input>
                </el-row>

                <el-row>
                  <el-input v-model="flickr.site_upload_way.flickr_account.api_key" placeholder="api_key">
                    <template slot="prepend" >
                      <el-tooltip placement="top">
                        <div slot="content">向 Flickr 申请到的 api_key</div>
                        <el-button type="primary">Api_key</el-button>
                      </el-tooltip>
                    </template>
                  </el-input>
                </el-row>

                <el-row>
                  <el-input v-model="flickr.site_upload_way.flickr_account.api_secret" placeholder="api_secret">
                    <template slot="prepend" >
                      <el-tooltip placement="top">
                        <div slot="content">向 Flickr 申请到的 api_secret</div>
                        <el-button type="primary">Api_secret</el-button>
                      </el-tooltip>
                    </template>
                  </el-input>
                </el-row>

                <el-row>
                  <el-input v-model="flickr.site_upload_way.flickr_account.oauth_token" placeholder="Oauth_token">
                    <template slot="prepend" >
                      <el-tooltip placement="top">
                        <div slot="content">通过 Flickr 工具获取到的Oauth_token</div>
                        <el-button type="primary">Oauth_token</el-button>
                      </el-tooltip>
                    </template>
                  </el-input>
                </el-row>

                <el-row>
                  <el-input v-model="flickr.site_upload_way.flickr_account.oauth_token_secret" placeholder="Oauth_token_secret">
                    <template slot="prepend" >
                      <el-tooltip placement="top">
                        <div slot="content">通过 Flickr 工具获取到的Oauth_token_secret</div>
                        <el-button type="primary">Oauth_token_secret</el-button>
                      </el-tooltip>
                    </template>
                  </el-input>
                </el-row>

                <el-row>
                  默认图片大小 :
                  <el-select v-model="flickr.site_upload_way.flickr_account.default_size" clearable placeholder="请选择" style="margin-left:25px">
                    <el-option
                      v-for="item in flickrOptions"
                      :key="item.value"
                      :label="item.size"
                      :value="item.value"/>
                  </el-select>
                </el-row>
                <el-row>
                  <el-button round style="float:right;">查看教程</el-button>
                </el-row>

              </el-card>
            </transition>

          </div>
        </el-card>
        <!-- JWT 的设置 -->
        <el-card class="box-card">
          <div slot="header" class="clearfix">
            <span>JWT 设置</span>
            <el-button style="float: right;" type="success" plain="" round="" icon="el-icon-success" @click="reset(jwt)">保存</el-button>
          </div>
          <div class="box">
            <el-row>
              <el-input v-model="jwt.jwt_secret" placeholder="jwt_secret">
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
              <el-input-number v-model="jwt.jwt_due_time" style="margin-left:25px;"/> 小时
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

  </div>
</template>

<script>
import { getConfig, getRsaKey, update } from '@/api/siteconfig'
import { getStoreList } from '@/api/image'
export default {
  name: 'SiteConfig',
  data() {
    return {
      tableData: [],
      config: null,
      siteBase: Object,
      sina: Object,
      flickr: Object,
      local: Object,
      upload: Object,
      database: Object,
      jwt: Object,
      storeList: [],
      selectValue: '',
      sinaOptions: [
        { size: 'square' },
        { size: 'thumb150' },
        { size: 'orj360' },
        { size: 'orj480' },
        { size: 'mw690' },
        { size: 'mw1024' },
        { size: 'mw2048' },
        { size: 'small' },
        { size: 'bmiddle' },
        { size: 'large' }
      ],
      sinaTimeOptions: [
        { size: '一小时', value: 3600 },
        { size: '半天', value: 43200 },
        { size: '一天', value: 86400 }
      ],
      sinaEx: '',
      flickrOptions: [
        { size: '小正方形', value: 's' },
        { size: '大正方形', value: 'q' },
        { size: '缩略图', value: 't' },
        { size: '小型图片', value: 'm' },
        { size: '稍微小的图片', value: 'n' },
        { size: '中等图片', value: '' },
        { size: '中等偏大', value: 'z' },
        { size: '中等大', value: 'c' },
        { size: '大尺寸', value: 'b' },
        { size: '高清尺寸', value: 'h' },
        { size: '超大尺寸', value: 'k' },
        { size: '原始图片', value: 'o' }
      ],
      mailPortList: [
        { info: '465 (ssl)', value: '465' },
        { info: '587 (tls)', value: '587' },
        { info: '22 (无加密)', value: '22' }
      ],
      privateKey: '',
      publicKey: '',
      configLoading: false
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
    this.loadStore()
  },
  methods: {
    loadConfig() {
      getConfig().then(resp => {
        this.config = resp.data
        this.siteBase = resp.data
        this.sina = resp.data
        this.flickr = resp.data
        this.upload = resp.data
        this.local = resp.data
        this.database = resp.data
        this.jwt = resp.data
        this.tableData = [resp.data.auxpi_info]
        this.selectValue = resp.data.api_default
        this.sinaEx = 'https://ws3.sinaimg.cn/' + this.sina.site_upload_way.sina_account.defult_pic_size + '/0072Vf1pgy1foxk7j6oxuj31hc0u0k6u'
        this.flickr.site_upload_way.flickr_account.default_size = resp.data.site_upload_way.flickr_account.default_size

        console.log(this.local.site_upload_way.local_store.open)
      }).then(
        this.loadConfig = true
      )
      getRsaKey().then(resp => {
        this.privateKey = resp.data.private_key
        this.publicKey = resp.data.public_key
      })
    },
    loadStore() {
      getStoreList().then(resp => {
        this.storeList = resp.list
        console.log(this.storeList)
      })
    },
    resetLink(v) {
      this.sinaEx = 'https://ws3.sinaimg.cn/' + v + '/0072Vf1pgy1foxk7j6oxuj31hc0u0k6u'
    },
    reset(v) {
      update(v).then(r => {
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
  padding: 32px;
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
