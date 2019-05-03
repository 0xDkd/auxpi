<template>
  <div class="dashboard-editor-container">
    <el-row :gutter="8">
      <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" class="box">
        <el-card class="">
          <div slot="header" class="clearfix">
            <span style="font-size:25px;"><svg-icon icon-class="menu"/> 菜单管理</span>
            <el-button style="float: right;" type="success" plain="" round="" icon="el-icon-success" @click="resetMenu">保存</el-button>
          </div>
          <div v-loading="loading" v-if="loading">
            <br>
            <br>
            <br>
            <br>
          </div>
          <transition name="el-fade-in-linear">
            <div v-if="show" >
              <el-row >
                <el-col :xs="24" :sm="24" :md="11" :lg="11" :xl="11" style="margin:10px;">
                  <el-card class="box-card">
                    <div slot="header" class="clearfix">
                      <span>启用的图床</span>
                    </div>
                    <!-- 组1 -->
                    <draggable :list="used" class="list-group" group="people" @change="log">
                      <el-card
                        v-for="(element,index) in used"
                        :key="element.name"
                        class="list-card-group-item"
                        shadow="hover"
                      >
                        <span><svg-icon :icon-class="element.icon" style="color:#3fb911"/>
                          {{ element.name }} <el-tag type="danger" size="mini">{{ element.rank }}</el-tag>&nbsp;<el-tag type="success" size="mini">{{ index+1 }}</el-tag></span>

                      </el-card>

                      <el-card
                        slot="header"
                        shadow="never"
                        role="group"
                        style="background:#fafafa;"
                      >
                        <svg-icon icon-class="start" style="color:#3fb911;" /> 开放使用的图床拖动于此
                      </el-card>
                    </draggable>
                  </el-card>
                </el-col>
                <el-col :xs="24" :sm="24" :md="11" :lg="11" :xl="11" style="margin:10px;">
                  <el-card class="box-card">
                    <div slot="header" class="clearfix">
                      <span>未启用的图床</span>
                    </div>
                    <!-- 组2 -->
                    <draggable :list="cancel" class="list-group" group="people" @change="log">
                      <el-card
                        v-for="element in cancel"
                        :key="element.name"
                        class="list-card-group-item"
                        shadow="hover"
                      >
                        <span><svg-icon :icon-class="element.icon" style="color:#f4516c;"/>
                          {{ element.name }} <el-tag type="success" size="mini">{{ element.id }}</el-tag></span>
                      </el-card>
                      <el-card
                        slot="header"
                        shadow="never"
                        role="group"
                        style="background:#fafafa;"
                      >
                        <svg-icon icon-class="stop" style="color:#f4516c;" /> 停止使用的图床拖动于此
                      </el-card>
                    </draggable>
                  </el-card>

                </el-col>
              </el-row>
            </div>
          </transition>
        </el-card>
      </el-col>
    </el-row>

  </div>
</template>

<script>
import { update, updateMenu } from '@/api/siteconfig'
import { getStoreList } from '@/api/image'
import draggable from 'vuedraggable'
export default {
  name: 'MenuConfig',
  display: 'Two Lists',
  order: 1,
  components: {
    draggable
  },
  data() {
    return {
      show: false,
      tableData: [],
      used: [],
      cancel: [],
      loading: true
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
    this.loadStore()
  },
  methods: {
    loadStore() {
      getStoreList().then(resp => {
        var i = 0
        var j = 0
        resp.list.forEach((element) => {
          if (element.status) {
            this.used[i] = element
            i++
          } else {
            this.cancel[j] = element
            j++
          }
        })
      }).then(() => {
        this.show = true
        this.loading = false
        console.log(this.used)
      }
      )
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
    },
    add: function() {
      this.list.push({ name: 'Juan' })
    },
    replace: function() {
      this.list = [{ name: 'Edgard' }]
    },
    clone: function(el) {
      return {
        name: el.name + ' cloned'
      }
    },
    log: function(evt) {
      window.console.log(evt)
    },
    resetMenu() {
      const menu = { 'enable': this.used, 'disable': this.cancel }
      updateMenu(menu).then(resp => {
        this.$message({
          message: '修改成功',
          type: 'success'
        })
      })
    }
  }
}
</script>

<style>
.dashboard-editor-container{
	padding: 32px;
	background-color: #f0f2f5;
}
.flip-list-move {
  transition: transform 0.5s;
}

.no-move {
  transition: transform 0s;
}

.ghostUse {
  opacity: 0.5;
  background: #c8ebfb;
}

.list-group {
  min-height: 20px;
}

.el-card__body{
  padding: 10px;
}

.list-card-group-item{
cursor: move;
}

.box{
  margin-top: 10px;
}
/* .box .el-row{
  margin: 30px;
} */
.box .el-card{
  margin-top:5px;
}
</style>
