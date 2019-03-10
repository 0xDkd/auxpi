<template>
  <el-container>
    <el-main >
      <!-- 顶部操作栏 -->
      <el-row style="margin-left:22px;padding:6px 0px 6px 18px;">
        <el-checkbox-group :indeterminate="isIndeterminate" v-model="checkAll" size="mini" @change="handleCheckAllChange">
          <el-checkbox label="全选" border />
          <el-select v-model="selectValue" placeholder="选择图床" size="mini" style="margin-left:40px;" @change="getSelectImages(selectValue)" >
            <el-option
              v-for="item in options"
              :key="item.id"
              :label="item.name"
              :value="item.id"/>
          </el-select>
          <el-dropdown size="mini" trigger="click" split-button type="primary" style="margin-left:40px;">
            操作
            <el-dropdown-menu slot="dropdown">
              <el-dropdown-item @click.native="getCheckedImageLinks"><i class="el-icon-circle-check"/> 提取链接</el-dropdown-item>
              <el-dropdown-item @click.native="showInfo"><i class="el-icon-date"/> 查看信息</el-dropdown-item>
              <el-dropdown-item @click.native="syncToLocal"><i class="el-icon-sort"/> 同步到本地</el-dropdown-item>

              <el-dropdown-item divided @click.native="deleteChecked" ><span style="color:red;"><i class="el-icon-delete" /> 删除选中</span></el-dropdown-item>

            </el-dropdown-menu>
          </el-dropdown>
          <el-switch
            v-if="showUser"
            v-model="model"
            style="display: inline;margin-left:10px;"
            active-color="#13ce66"
            inactive-color="#ff4949"
            active-text="操作模式"
            inactive-text="查看模式"
            active-value="edit"
            inactive-value="view"
            @change="changeModel"
          />
        </el-checkbox-group>
      </el-row>
      <!-- 被操作用户 -->
      <el-row v-if="showUser" style="margin:20px">
        <el-alert
          title=""
          type="success"
          center
        >您当前正在操作的用户为 : <b> {{ user }}</b>
          <br>
          <br>
          AUXPI 图床提醒您: 数据千万条,备份第一条。数据不备份,站长两行泪
        </el-alert>
      </el-row>
      <!-- 进度条 -->
      <el-row v-show="showProgress" style="margin:20px">
        <el-col >
          <transition name="el-fade-in-linear">
            <el-progress :stroke-width="14" :percentage.sync="progress" status="success"/>
          </transition>
        </el-col>

      </el-row>
      <!-- 图片 -->
      <el-row v-loading="mainLoading" v-if="isExist" style="margin:20px">
        <el-checkbox-group v-model="checkedImages" @change="handleCheckedImageChange">
          <el-col v-for="(item,index) in imgLists" :xs="12" :sm="6" :md="4" :lg="4" :xl="4" :key="item.id" style="padding:0px 5px 10px;">
            <el-card :body-style="{ padding: '0px'}" shadow="hover" >
              <div :style="{backgroundImage:'url(' + item.link + ')' ,backgroundRepeat:'no-repeat', backgroundPosition:'center top'}" class="image" />
              <div style="padding: 14px;">
                <div class="bottom clearfix">
                  <el-row style="margin-top:2px;">
                    <el-button type="primary" icon="el-icon-info" plain size="mini" @click="showSingleInfo(item)"/>
                    <el-button type="info" icon="el-icon-zoom-in" plain size="mini" @click="zoom(item.link)"/>
                    <el-button type="danger" icon="el-icon-delete" plain size="mini" @click="deleteImg(index,item.id)"/>
                    <el-checkbox :key="index" :label="item" style="margin-left:20px" >&nbsp;</el-checkbox>
                  </el-row>
                </div>
              </div>
            </el-card>
          </el-col>
        </el-checkbox-group>
      </el-row>
      <!-- 提示  -->
      <el-row v-if="!isExist" style="margin:100px">
        <el-col :xs="24" :sm="6" :md="4" :lg="4" :xl="4" :offset="10">
          <el-alert
            title="这里什么都没有~快去补充一点图片吧 []~(￣▽￣)~*"
            type="success" />
        </el-col>

      </el-row>
      <!-- 分页 -->
      <!--  -->
      <pagination v-show="total>0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.limit" @pagination="getList"/>

    </el-main>

    <!-- 图片信息 -->
    <el-dialog :visible.sync="imageInfoDialog" :fullscreen="true" title="图片信息">
      <el-table :data="imgInfo">
        <el-table-column property="id" label="ID" width="150"/>
        <el-table-column property="name" label="名称" width="200"/>
        <el-table-column label="上传者">
          <template slot-scope="scope">
            <el-tag size="medium">{{ scope.row.user.username }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="图床">
          <template slot-scope="scope">
            <el-tag size="medium" type="success">{{ scope.row.store.name }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column property="link" label="图片链接"/>
        <el-table-column label="上传时间">
          <template slot-scope="scope">
            {{ scope.row.created_on | parseTime('{y}-{m}-{d} {h}:{i}') }}
          </template>
        </el-table-column>
      </el-table>
    </el-dialog>

    <!-- 链接合集 -->
    <el-dialog :visible.sync="imageLinkDialog" :fullscreen="fullScreen" title="链接合集">
      <div style="margin-bottom: 20px;">
        <el-button
          :data-clipboard-text ="linksArea"
          size="small"
          class="tag-read"
          @click="copy()"
        >
          Copy All
        </el-button>
      </div>

      <el-tabs type="border-card" style="margin:0px 0px 20px;" @tab-click="transformTo">
        <el-tab-pane label="URL">
          <el-input
            :rows="linksRows"
            v-model="linksArea"

            type="textarea"
          />
        </el-tab-pane>

        <el-tab-pane label="HTML">
          <el-input
            :rows="linksRows"
            v-model="linksArea"
            type="textarea"
          />
        </el-tab-pane>
        <el-tab-pane label="MarkDown" >
          <el-input
            :rows="linksRows"
            v-model="linksArea"
            type="textarea"
          />
        </el-tab-pane>

        <el-tab-pane label="BBCode">
          <el-input
            :rows="linksRows"
            v-model="linksArea"
            type="textarea"
          />
        </el-tab-pane>
      </el-tabs>

      <el-table :data="checkedImages">
        <el-table-column label="图床">
          <template slot-scope="scope">
            <el-tag size="medium" type="success">{{ scope.row.store.name }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column property="link" label="图片链接"/>
        <el-table-column label="操作">
          <template slot-scope="scope" >
            <el-button :data-clipboard-text="scope.row.link" class="tag-read" size="medium" plain @click="copy">复制链接</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-dialog>

    <!-- 看大图 -->
    <el-dialog :visible.sync="imageDialog" :fullscreen="true" title="图片原图">
      <img :src="dialogLink">
    </el-dialog>

    <!-- 单个图片信息 -->
    <el-dialog :visible.sync="singleInfo" title="图片信息">
      <el-row>
        <el-col>
          <el-card :body-style="{ padding: '0px' }">
            <img :src=" info.link" class="image">
            <div style="padding: 14px;">

              <el-button :data-clipboard-text="info.link" type="success" size="medium" class="tag-read" plain round @click="copy">复制链接</el-button>
              <el-button size="medium" plain round @click="zoom(info.link)">查看原图</el-button>
              <el-button size="medium" type="warning" plain round @click="jump(info.link)">新窗口查看</el-button>
          </div></el-card>
        </el-col>
      </el-row>
      <el-table :data="[info]" style="margin-top:20px">
        <el-table-column property="id" label="ID" width="150"/>
        <el-table-column property="name" label="名称" width="200"/>
        <el-table-column label="上传者">
          <template slot-scope="scope">
            <el-tag size="medium">{{ scope.row.user.username }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="图床">
          <template slot-scope="scope">
            <el-tag size="medium" type="success">{{ scope.row.store.name }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column property="link" label="图片链接"/>
        <el-table-column label="上传时间">
          <template slot-scope="scope">
            {{ scope.row.created_on | parseTime('{y}-{m}-{d} {h}:{i}') }}
          </template>
        </el-table-column>
      </el-table>
    </el-dialog>
  </el-container>
</template>

<script>

import Pagination from '@/components/Pagination' // Secondary package based on el-pagination
import { getImageList, delImage, getStoreList, syncImage } from '@/api/image'
import { getUserImages } from '@/api/user'
import Clipboard from 'clipboard'

export default {
  components: { Pagination },
  data() {
    return {
      imgLists: [],
      checkAll: false,
      isIndeterminate: true,
      checkedImages: [],
      // 分页参数
      listQuery: {
        page: 1,
        limit: 18,
        type: 0,
        sort: '+id'
      },
      total: 0,
      mainLoading: false,
      options: [],
      selectValue: '',
      firstOption: {
        id: 0,
        name: '全部'
      },
      selectImages: [],
      imagesBackup: [],
      isExist: true,
      imageInfoDialog: false,
      imageLinkDialog: false,
      imgInfo: [],
      linksArea: '',
      linksRows: 0,
      linksBackup: '',
      fullScreen: false,
      dialogLink: '',
      imageDialog: false,
      singleInfo: false,
      info: [],
      syncImageInfo: [],
      ws: null,
      wsMsg: '',
      progress: 0,
      showProgress: false,
      user: '',
      uid: 0,
      showUser: false,
      model: 'edit'
    }
  },
  created() {
    this.getStore()
    this.getList()
  },
  methods: {
    deleteImg(index, id) {
      this.$confirm('此操作将永久删除该文件, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        delImage([id, 0])
        this.imgLists.splice(index, 1)
        this.$message({
          type: 'success',
          message: '删除成功!'
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        })
      })
    },
    zoom(url) {
      this.dialogLink = ''
      this.dialogLink = url
      this.imageDialog = true
      // window.open(url)
    },
    jump(url) {
      window.open(url)
    },
    handleCheckAllChange(val) {
      this.checkedImages = val ? this.imgLists : []
      this.isIndeterminate = false
    },
    handleCheckedImageChange(v) {
      const checkCount = v.length
      this.checkAll = checkCount === this.imgLists.length
      this.isIndeterminate = checkCount > 0 && checkCount < this.imgLists.length
    },
    deleteChecked() {
      const checkCount = this.checkedImages.length
      if (checkCount === 0) {
        this.$message({
          type: 'warning',
          message: '您没有选择任何图片'
        })
      } else {
        this.$confirm('此操作将永久删除记录, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          const ids = []
          const checkCount = this.checkedImages.length
          for (let index = 0; index < checkCount; index++) {
            console.log(this.checkedImages[index].id)
            ids[index] = this.checkedImages[index].id
          }
          // 删除 数据库数据
          delImage(ids)

          // 删去数组元素
          const remain = []
          this.imgLists.forEach(element => {
            const c = this.checkedImages.findIndex(b => element === b)
            if (c > -1) {
              delete this.checkedImages[c]
            } else {
              remain.push(element)
            }
          })
          // 覆盖数组
          this.imgLists = remain
          this.$message({
            type: 'success',
            message: '删除成功!'
          })
          // 重置
          this.checkedImages = []
        }).catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除'
          })
        })
      }
      // 操作完成之后重置
    },
    getList() {
      this.mainLoading = true
      this.uid = this.$route.params.id
      if (this.uid) {
        this.showUser = true
      }

      if (this.showUser) {
        getUserImages(this.uid, this.listQuery).then(response => {
          this.imgLists = response.list
          this.imagesBackup = response.list
          this.total = response.total
          this.mainLoading = false
          if (this.imgLists === undefined || this.imgLists.length === 0) {
            this.isExist = false
            return
          }
          this.isExist = true
          this.user = response.list[0].user.username
          console.log(response.list[0].user.username)
        })
      } else {
        getImageList(this.listQuery).then(response => {
          this.imgLists = response.list
          this.imagesBackup = response.list
          this.total = response.total
          this.mainLoading = false
          if (this.imgLists === undefined || this.imgLists.length === 0) {
            this.isExist = false
            return
          }
          this.isExist = true
        })
      }

      this.checkedImages = []
      this.checkAll = false
    },
    getAll() {
      console.log(this.checkedImages)
    },
    getStore() {
      getStoreList().then(resp => {
        this.options = resp.list
        this.options.unshift(this.firstOption)
        // console.log(this.options)
      })
    },
    getSelectImages(query) {
      this.listQuery.type = query
      this.imgLists = this.getList()
    },
    showInfo() {
      if (this.checkedImages.length === 0 || this.checkedImages === undefined) {
        this.$message({
          type: 'warning',
          message: '没有选中任何图片'
        })
        return
      }
      this.imgInfo = this.checkedImages
      this.imageInfoDialog = true
    },
    getCheckedImageLinks() {
      if (this.checkedImages.length === 0 || this.checkedImages === undefined) {
        this.$message({
          type: 'warning',
          message: '您没有选择任何图片'
        })
        return
      }
      this.linksRows = 0
      this.linksBackup = ''
      this.linksArea = ''
      this.fullScreen = false
      this.checkedImages.forEach(e => {
        this.linksArea += e.link + '\n'
        this.linksRows++
      })
      if (this.linksRows > 10) {
        this.fullScreen = true
      }
      this.linksBackup = this.linksArea
      this.imageLinkDialog = true
    },
    copy() {
      var clipboard = new Clipboard('.tag-read')
      console.log(clipboard)
      clipboard.on('success', e => {
        this.$message({
          type: 'success',
          message: '复制成功'
        })
        // 释放内存
        clipboard.destroy()
      })
      clipboard.on('error', e => {
        // 不支持复制
        console.log('该浏览器不支持自动复制')
        // 释放内存
        clipboard.destroy()
      })
    },
    transformTo(tab, event) {
      this.linksArea = ''
      switch (tab.label) {
        case 'URL':
          this.linksArea = this.linksBackup
          break
        case 'MarkDown':
          this.checkedImages.forEach(e => {
            this.linksArea += '![' + e.name + '](' + e.link + ')' + '\n'
          })
          break
        case 'HTML':
          // this.linksArea = ''
          this.checkedImages.forEach(e => {
            this.linksArea += '<img  alt="' + e.name + '" src="' + e.link + '" />' + '\n'
          })
          break
        case 'BBCode':
          // this.linksArea = ''
          this.checkedImages.forEach(e => {
            this.linksArea += '[img]' + e.link + '[/img]' + '\n'
          })
          break
      }
      console.log(this.linksRows)
    },
    showSingleInfo(info) {
      this.info = info

      this.singleInfo = true
    },
    syncToLocal() {
      if (this.checkedImages.length === 0 || this.checkedImages === undefined) {
        this.$message({
          type: 'warning',
          message: '没有选中任何图片'
        })
        return
      }

      this.ws = new WebSocket(process.env.WS_LOCATION)

      this.syncImageInfo = []
      const Count = this.checkedImages.length
      this.checkedImages.forEach((e, i) => {
        this.syncImageInfo[i] = { id: e.id, url: e.link }
      })

      var self = this
      this.showProgress = true
      this.ws.addEventListener('message', (e) => {
        var msg = JSON.parse(e.data)
        console.log('进入监听')
        // console.log(msg)
        switch (msg.status) {
          case 'running':
            // console.log('running')
            self.progress = ((msg.data + 1) / Count) * 100
            console.log(self.progress)
            break
          case 'success':
            console.log('同步完成')
            self.$notify({
              title: msg.title,
              message: msg.msg,
              type: msg.status
            })
            // self.ws.close()
            self.showProgress = false
            break
          case 'timeout':
            console.log('任务超时')
            self.$notify({
              title: msg.title,
              message: msg.msg,
              type: 'danger'
            })
            self.ws.close()
            break
          default:
            break
        }
      })

      syncImage(this.syncImageInfo)
      // this.send(this.syncImageInfo)
    },
    notification(msg) {
      this.$notify({
        title: '同步成功',
        message: msg,
        type: 'success'
      })
    },
    changeModel() {
      setTimeout(() => {
        this.$router.push({ name: 'usersInfoView', params: { id: this.uid }})
      }, 100)
    }
  }
}

</script>

<style>
  .image {
    width: 100%;
    height: 230px;
    display: block;
  }

  .el-row {
    margin: 0px -5px;
  }
  .header{
    height: 130px;
    line-height: 127px;
    display: block;
    margin-bottom: 20px
  }
  .el-switch{
    padding-bottom: 10px;
  }

</style>
