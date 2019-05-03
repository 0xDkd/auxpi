<template>
  <div class="dashboard-editor-container">
    <el-row>
      <el-card>
        <div slot="header" class="clearfix">
          <span style="font-size:25px;"><svg-icon icon-class="img-manager" style="font-size:30px;"/> 图床设置</span>
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
              <!-- Github -->
              <el-card class="box-card">
                <div slot="header" class="clearfix" >
                  <span style="font-size:25px;">Github</span>
                  <el-button style="float: right;" type="success" plain="" round="" icon="el-icon-success" @click="reset(github,github_info)">保存</el-button>
                </div>
                <div class="box">
                  <el-row>
                    是否开启Github 存储:
                    <el-switch
                      v-model="github_info.status"
                      active-color="#13ce66"
                      inactive-color="#ff4949"
                      style="margin-left:25px"
                    />
                  </el-row>
                  <transition name="el-zoom-in-center">
                    <el-card v-if="github_info.status" class="box-card">
                      <el-row>
                        <el-input v-model="github.access_token" placeholder="assess_token">
                          <template slot="prepend" >
                            <el-tooltip placement="top">
                              <div slot="content">您在 github 申请到的access_token</div>
                              <el-button type="primary">Access_Token</el-button>
                            </el-tooltip>
                          </template>
                        </el-input>
                      </el-row>

                      <el-row>
                        <el-input v-model="github.owner" placeholder="owner">
                          <template slot="prepend" >
                            <el-tooltip placement="top">
                              <div slot="content">仓库拥有者，一般填写您本人的 github 用户名即可</div>
                              <el-button type="primary">Owner</el-button>
                            </el-tooltip>
                          </template>
                        </el-input>
                      </el-row>

                      <el-row>
                        <el-input v-model="github.repo" placeholder="repo">
                          <template slot="prepend" >
                            <el-tooltip placement="top">
                              <div slot="content">填写您要使用的仓库即可</div>
                              <el-button type="primary">Repo</el-button>
                            </el-tooltip>
                          </template>
                        </el-input>
                      </el-row>

                      <el-row>
                        <el-input v-model="github.email" placeholder="email">
                          <template slot="prepend" >
                            <el-tooltip placement="top">
                              <div slot="content">填写您注册github 所使用的 Email 即可</div>
                              <el-button type="primary">Email</el-button>
                            </el-tooltip>
                          </template>
                        </el-input>
                      </el-row>
                      <el-row>
                        是否开启代理:
                        <el-switch
                          v-model="github.proxy.status"
                          active-color="#13ce66"
                          inactive-color="#ff4949"
                          style="margin-left:25px"
                        />
                        <transition name="el-fade-in-linear">
                          <el-card v-if="github.proxy.status">
                            <el-input v-model="github.proxy.node" placeholder="proxy">
                              <template slot="prepend" >
                                <el-tooltip placement="top">
                                  <div slot="content">代理链接</div>
                                  <el-button type="primary">Proxy</el-button>
                                </el-tooltip>
                              </template>
                            </el-input>
                          </el-card>
                        </transition>

                      </el-row>
                    </el-card>
                  </transition>

                </div>
              </el-card>
              <!-- 本地图床配置 -->
              <el-card class="box-card">
                <div slot="header" class="clearfix">

                  <span style="font-size:25px;">本地图床配置</span>
                  <el-button style="float: right;" type="success" plain="" round="" icon="el-icon-success" @click="reset(local,local_info)">保存</el-button>
                </div>
                <div >
                  <el-row>
                    是否开启本地图床:
                    <el-switch
                      v-model="local_info.status"
                      active-color="#13ce66"
                      inactive-color="#ff4949"
                      style="margin-left:25px"
                    />
                  </el-row>
                  <transition name="el-zoom-in-center">
                    <el-card v-show="local_info.status" class="box-card">
                      <el-row>
                        <el-button round>点我查看配置教程</el-button>
                      </el-row>
                      <el-row>
                        <el-input v-model="local.link" placeholder="请输入您的默认链接地址">
                          <template slot="prepend" >
                            <el-tooltip placement="top">
                              <div slot="content">开头必须要有"/",最后不能有 "/"</div>
                              <el-button type="primary">软路径</el-button>
                            </el-tooltip>
                          </template>
                        </el-input>
                      </el-row>

                      <el-row>
                        <el-input v-model="local.storage_location" placeholder="请输入储存路径">
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
              <!-- Imgur -->
              <el-card class="box-card">
                <div slot="header" class="clearfix">

                  <span style="font-size:25px;">Imgur</span>
                  <el-button style="float: right;" type="success" plain="" round="" icon="el-icon-success" @click="reset(imgur,imgur_info)">保存</el-button>
                </div>
                <div >
                  <el-row>
                    是否开启Imgur图床:
                    <el-switch
                      v-model="imgur_info.status"
                      active-color="#13ce66"
                      inactive-color="#ff4949"
                      style="margin-left:25px"
                    />
                  </el-row>
                  <!-- TODO:增加淡入淡出的特效 -->
                  <transition name="el-zoom-in-center">
                    <el-card v-show="imgur_info.status" class="box-card">
                      <el-row>
                        <el-button round>点我查看配置教程</el-button>
                      </el-row>
                      <el-row>
                        <el-input v-model="imgur.imgur_id" placeholder="Client ID">
                          <template slot="prepend" >
                            <el-tooltip placement="top">
                              <div slot="content">从 imgur 申请到的 id</div>
                              <el-button type="primary">Client ID</el-button>
                            </el-tooltip>
                          </template>
                        </el-input>
                      </el-row>
                      <el-row>
                        是否开启代理:
                        <el-switch
                          v-model="imgur.proxy.status"
                          active-color="#13ce66"
                          inactive-color="#ff4949"
                          style="margin-left:25px"
                        />
                        <transition name="el-fade-in-linear">
                          <el-card v-if="imgur.proxy.status">
                            <el-input v-model="imgur.proxy.node" placeholder="proxy">
                              <template slot="prepend" >
                                <el-tooltip placement="top">
                                  <div slot="content">代理链接</div>
                                  <el-button type="primary">Proxy</el-button>
                                </el-tooltip>
                              </template>
                            </el-input>
                          </el-card>
                        </transition>
                      </el-row>
                    </el-card>
                  </transition>
                </div>
              </el-card>
            </el-col>

            <el-col :xs="24" :sm="24" :md="12" :lg="12" :xl="12" class="box">

              <!-- Gitee -->
              <el-card class="box-card">
                <div slot="header" class="clearfix" >
                  <span style="font-size:25px;">Gitee</span>
                  <el-button style="float: right;" type="success" plain="" round="" icon="el-icon-success" @click="reset(gitee,gitee_info)">保存</el-button>
                </div>
                <div class="box">
                  <el-row>
                    是否开启Gitee 存储:
                    <el-switch
                      v-model="gitee_info.status"
                      active-color="#13ce66"
                      inactive-color="#ff4949"
                      style="margin-left:25px"
                    />
                  </el-row>
                  <transition name="el-zoom-in-center">
                    <el-card v-if="gitee_info.status" class="box-card">
                      <el-row>
                        <el-input v-model="gitee.access_token" placeholder="assess_token">
                          <template slot="prepend" >
                            <el-tooltip placement="top">
                              <div slot="content">您在 gitee 申请到的access_token</div>
                              <el-button type="primary">Access_Token</el-button>
                            </el-tooltip>
                          </template>
                        </el-input>
                      </el-row>

                      <el-row>
                        <el-input v-model="gitee.owner" placeholder="owner">
                          <template slot="prepend" >
                            <el-tooltip placement="top">
                              <div slot="content">仓库拥有者，一般填写您本人的 github 用户名即可</div>
                              <el-button type="primary">Owner</el-button>
                            </el-tooltip>
                          </template>
                        </el-input>
                      </el-row>

                      <el-row>
                        <el-input v-model="gitee.repo" placeholder="repo">
                          <template slot="prepend" >
                            <el-tooltip placement="top">
                              <div slot="content">填写您要使用的仓库即可</div>
                              <el-button type="primary">Repo</el-button>
                            </el-tooltip>
                          </template>
                        </el-input>
                      </el-row>
                    </el-card>
                  </transition>

                </div>
              </el-card>
              <!-- 新浪图床配置 -->
              <el-card class="box-card">
                <div slot="header" class="clearfix">

                  <span style="font-size:25px;">微博图床配置</span>
                  <el-button style="float: right;" type="success" plain="" round="" icon="el-icon-success" @click="reset(sina,sina_info)">保存</el-button>
                </div>
                <div >
                  <el-row>
                    是否开启微博图床:
                    <el-switch
                      v-model="sina_info.status"
                      active-color="#13ce66"
                      inactive-color="#ff4949"
                      style="margin-left:25px"
                    />
                  </el-row>
                  <transition name="el-zoom-in-center">

                    <el-card v-if="sina_info.status" class="box-card">
                      <el-row>
                        <el-input v-model="sina.user_name" placeholder="请输入用户名">
                          <template slot="prepend" >
                            <el-tooltip placement="top">
                              <div slot="content">您的微博登录用来登录的用户名</div>
                              <el-button type="primary">微博用户名</el-button>
                            </el-tooltip>
                          </template>
                        </el-input>
                      </el-row>

                      <el-row>
                        <el-input v-model="sina.pass_word" placeholder="请输入密码">
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
                        <el-select v-model="sina.defult_pic_size" placeholder="请选择" style="margin-left:25px" @change="resetLink(sina.defult_pic_size)">
                          <el-option
                            v-for="item in sinaOptions"
                            :key="item.size"
                            :label="item.size"
                            :value="item.size"/>
                        </el-select>
                        <el-card shadow="always">
                          链接效果 (点击可以直接查看图片) :<br><br>
                          <a :href="'https://images.weserv.nl/?url='+sinaEx" target="_blank">{{ sinaEx }}</a>
                        </el-card>
                      </el-row>

                      <el-row>
                        <span> Cookie 缓存时间:</span>

                        <el-select v-model="sina.reset_sina_cookie_time" clearable placeholder="请选择" style="margin-left:25px">
                          <el-option
                            v-for="item in sinaTimeOptions"
                            :key="item.value"
                            :label="item.size"
                            :value="item.value"/>
                        </el-select>
                      </el-row>
                      <el-row>
                        是否开启代理:
                        <el-switch
                          v-model="sina.proxy.status"
                          active-color="#13ce66"
                          inactive-color="#ff4949"
                          style="margin-left:25px"
                        />
                        <transition name="el-fade-in-linear">
                          <el-card v-if="sina.proxy.status">
                            <el-input v-model="sina.proxy.node" placeholder="proxy">
                              <template slot="prepend" >
                                <el-tooltip placement="top">
                                  <div slot="content">代理链接</div>
                                  <el-button type="primary">Proxy</el-button>
                                </el-tooltip>
                              </template>
                            </el-input>
                          </el-card>
                        </transition>

                      </el-row>
                    </el-card>
                  </transition>
                </div>
              </el-card>
              <!-- Flickr 图床配置 -->
              <el-card class="box-card">
                <div slot="header" class="clearfix">

                  <span style="font-size:25px;">Flickr图床配置</span>
                  <el-button style="float: right;" type="success" plain="" round="" icon="el-icon-success" @click="reset(flickr,flickr_info)">保存</el-button>
                </div>
                <div >
                  <el-row>
                    是否开启Flickr图床:
                    <el-switch
                      v-model="flickr_info.status"
                      active-color="#13ce66"
                      inactive-color="#ff4949"
                      style="margin-left:25px"
                    />
                  </el-row>
                  <!-- TODO:增加淡入淡出的特效 -->
                  <transition name="el-zoom-in-center">
                    <el-card v-if="flickr_info.status" class="box-card">

                      <el-row>
                        <el-input v-model="flickr.id" placeholder="Flickr ID">
                          <template slot="prepend" >
                            <el-tooltip placement="top">
                              <div slot="content">Flickr ID 及 Flickr 分配给您的用户 ID</div>
                              <el-button type="primary">ID</el-button>
                            </el-tooltip>
                          </template>
                        </el-input>
                      </el-row>

                      <el-row>
                        <el-input v-model="flickr.api_key" placeholder="api_key">
                          <template slot="prepend" >
                            <el-tooltip placement="top">
                              <div slot="content">向 Flickr 申请到的 api_key</div>
                              <el-button type="primary">Api_key</el-button>
                            </el-tooltip>
                          </template>
                        </el-input>
                      </el-row>

                      <el-row>
                        <el-input v-model="flickr.api_secret" placeholder="api_secret">
                          <template slot="prepend" >
                            <el-tooltip placement="top">
                              <div slot="content">向 Flickr 申请到的 api_secret</div>
                              <el-button type="primary">Api_secret</el-button>
                            </el-tooltip>
                          </template>
                        </el-input>
                      </el-row>

                      <el-row>
                        <el-input v-model="flickr.oauth_token" placeholder="Oauth_token">
                          <template slot="prepend" >
                            <el-tooltip placement="top">
                              <div slot="content">通过 Flickr 工具获取到的Oauth_token</div>
                              <el-button type="primary">Oauth_token</el-button>
                            </el-tooltip>
                          </template>
                        </el-input>
                      </el-row>

                      <el-row>
                        <el-input v-model="flickr.oauth_token_secret" placeholder="Oauth_token_secret">
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
                        <el-select v-model="flickr.default_size" clearable placeholder="请选择" style="margin-left:25px">
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
            </el-col>
          </el-row>
        </transition>
      </el-card>

    </el-row>

  </div>
</template>

<script>
import { getStoreOption, updateStoreAccount, updateStoreStatus } from '@/api/siteconfig'
export default {
  name: 'SiteConfig',
  data() {
    return {
      show: false,
      tableData: [],
      sina: null,
      flickr: null,
      local: null,
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
      loading: true,
      github: null,
      gitee: null,
      imgur: null,
      gb: { 'key': 'github', 'group': 'conf' },
      ge: { 'key': 'gitee', 'group': 'conf' },
      ll: { 'key': 'local', 'group': 'conf' },
      ir: { 'key': 'imgur', 'group': 'conf' },
      fr: { 'key': 'flickr', 'group': 'conf' },
      sa: { 'key': 'sina', 'group': 'conf' },
      github_info: null,
      gitee_info: null,
      local_info: null,
      flickr_info: null,
      imgur_info: null,
      sina_info: null

    }
  },
  created() {
    this.loadConfig()
  },
  methods: {
    loadConfig() {
      // 获取github 和 gitee 的配置
      const getGithub = getStoreOption(this.gb).then(resp => {
        this.github = JSON.parse(resp.data.account)
        this.github_info = resp.data.store
      })
      const getGitee = getStoreOption(this.ge).then(resp => {
        this.gitee = JSON.parse(resp.data.account)
        this.gitee_info = resp.data.store
      })
      const getLocal = getStoreOption(this.ll).then(resp => {
        this.local = JSON.parse(resp.data.account)
        this.local_info = resp.data.store
      })
      const getImgur = getStoreOption(this.ir).then(resp => {
        this.imgur = JSON.parse(resp.data.account)
        this.imgur_info = resp.data.store
        console.log(this.imgur.proxy.status)
      })
      const getFlickr = getStoreOption(this.fr).then(resp => {
        this.flickr = JSON.parse(resp.data.account)
        this.flickr_info = resp.data.store
      })
      const getSina = getStoreOption(this.sa).then(resp => {
        this.sina = JSON.parse(resp.data.account)
        this.sina_info = resp.data.store
      })
      Promise.all([getGithub, getGitee, getLocal, getImgur, getFlickr, getSina]).then(() => {
        console.log('allDone')
        this.sinaEx = 'https://ww2.sinaimg.cn/' + this.sina.defult_pic_size + '/006A66c0ly1g2l330kewvj31hc0u0azk'

        this.show = true
        this.loading = false
      })
    },
    resetLink(v) {
      this.sinaEx = 'https://ww2.sinaimg.cn/' + v + '/006A66c0ly1g2l330kewvj31hc0u0azk'
    },
    reset(account, status) {
      updateStoreAccount(account, status.api).then(r => {
        this.$message({
          message: '修改成功',
          type: 'success'
        })
      })
      updateStoreStatus(status)
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
