<template>
  <div class="dashboard-editor-container">
    <el-row :gutter="8">
      <el-col :xs="{span: 24}" :sm="{span: 24}" :md="{span: 24}" :lg="{span: 24}" :xl="{span: 24}" style="margin-bottom:30px;">
        <el-card>
          <div >
            <el-upload
              :on-success="handleUploadSuccess"
              :headers="myHeader"
              :data="myData"
              drag
              name="image"
              action="http://localhost:2333/api/v2/upload/"
              multiple
            >
              <i class="el-icon-upload">
                <div class="el-upload__text">将图片拖到此处，或<em>点击上传</em></div>
            </i></el-upload>
          </div>

        </el-card>
        <el-card>
          <el-button type="primary" round @click="sortUploadLink">排序</el-button>
          <el-tabs v-model="activeName" @tab-click="handleClick">
            <el-tab-pane label="卡片式展示" name="card">
              <el-col v-for="(item,index) in testArr" :xs="12" :sm="6" :md="4" :lg="4" :xl="4" :key="index" style="padding:0px 5px 10px;">
                <el-card :body-style="{ padding: '0px'}" shadow="hover" >
                  <img :src="item.file" class="image">
                  <div style="padding: 14px;">
                    <div class="bottom clearfix">
                      <el-row style="margin-top:2px;">
                        <el-button type="info" icon="el-icon-zoom-in" plain size="mini" @click="zoom(item.file)"/>
                        <el-button :data-clipboard-text ="item.url" type="success" icon="el-icon-tickets" size="mini" plain class="tag-read" @click="copy()"/>
                        <el-checkbox :key="index" :label="item" style="margin-left:20px" >&nbsp;</el-checkbox>
                      </el-row>
                    </div>
                  </div>
                </el-card>
              </el-col>
            </el-tab-pane>
            <el-tab-pane label="链接式展示" name="link">
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
                    :autosize="{ minRows: 2}"
                    type="textarea"
                  />
                </el-tab-pane>

                <el-tab-pane label="HTML">
                  <el-input
                    :rows="linksRows"
                    v-model="linksArea"
                    :autosize="{ minRows: 2}"
                    type="textarea"
                  />
                </el-tab-pane>
                <el-tab-pane label="MarkDown" >
                  <el-input
                    :rows="linksRows"
                    v-model="linksArea"
                    :autosize="{ minRows: 2}"
                    type="textarea"
                  />
                </el-tab-pane>

                <el-tab-pane label="BBCode">
                  <el-input
                    :rows="linksRows"
                    v-model="linksArea"
                    :autosize="{ minRows: 2}"
                    type="textarea"
                  />
                </el-tab-pane>
              </el-tabs>
            </el-tab-pane>
          </el-tabs>
        </el-card>
      </el-col>
    </el-row>

    <!-- 看大图 -->
    <el-dialog :visible.sync="imageDialog" title="图片原图">
      <img :src="dialogLink" width="100%" alt="">
    </el-dialog>

    <!-- 单个图片信息 -->
    <!-- <el-dialog :visible.sync="singleInfo" title="图片信息">
      <el-row>
        <el-col>
          <el-card :body-style="{ padding: '0px' }">
            <img :src=" info.url" class="image">
            <div style="padding: 14px;">

              <el-button :data-clipboard-text="info.url" type="success" size="medium" class="tag-read" plain round @click="copy">复制链接</el-button>
              <el-button size="medium" plain round @click="zoom(info.url,this)">查看原图</el-button>
              <el-button size="medium" type="warning" plain round @click="jump(info.url)">新窗口查看</el-button>
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
    </el-dialog> -->
  </div>
</template>

<script>
import Clipboard from 'clipboard'
import { getToken } from '@/utils/auth'
export default {
  name: 'Uploader',
  data() {
    return {
      testArr: [],
      imageDialog: false,
      showtext: false,
      sortfile: [],
      dialogLink: '',
      activeName: 'card',
      info: [],
      singleInfo: false,
      linksArea: '',
      linksRows: 0,
      linksBackup: '',
      authHeader: Object
    }
  },
  computed: {
    myHeader() {
      return { 'X-Token': getToken() }
    }
  },
  methods: {
    handleUploadSuccess(resp, file, fileList) {
      this.testArr.push({ 'url': resp.data.url, 'file': file.url })
      this.linksArea += resp.data.url + '\n'
      this.linksBackup = this.linksArea
      this.sortfile = fileList
    },
    sortUploadLink() {
      this.testArr = []
      this.sortfile.forEach((element, index) => {
        this.testArr[index] = { 'url': element.response.data.url, 'file': element.url }
      })
    },
    handleClick(tab, event) {
      console.log(tab, event)
    },
    zoom(url) {
      this.dialogLink = ''
      this.dialogLink = url
      this.imageDialog = true
    },
    showSingleInfo(info) {
      this.info = info
      this.singleInfo = true
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
          this.testArr.forEach((e, index) => {
            this.linksArea += '![' + index + '](' + e.url + ')' + '\n'
          })
          break
        case 'HTML':
          // this.linksArea = ''
          this.testArr.forEach((e, index) => {
            this.linksArea += '<img  alt="' + index + '" src="' + e.url + '" />' + '\n'
          })
          break
        case 'BBCode':
          // this.linksArea = ''
          this.testArr.forEach(e => {
            this.linksArea += '[img]' + e.url + '[/img]' + '\n'
          })
          break
      }
      console.log(this.linksRows)
    }
  }
}
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
.dashboard-editor-container {
  padding: 20px 20px 400px 20px;
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
.pre-scrollable {
	max-height: 340px;
	overflow-y: scroll;
}
pre {
	display: block;
	line-height: 1.38461538;
	color: #333;
	word-break: break-all;
	word-wrap: break-word;
	background-color: #f5f5f5;
	border: 1px solid #ccc;
	border-radius: 4px;
  text-align: center;
  min-height: 100px;
}
.image {
    width: 100%;
    height: 230px;
    display: block;
  }
</style>
