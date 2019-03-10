<template>
  <div class="dashboard-editor-container">

    <el-row style="margin-left:22px;padding:6px 0px 6px 18px;">
      <el-switch
        v-model="currentModel"
        style="display: block"
        active-color="#13ce66"
        inactive-color="#ff4949"
        active-text="编辑模式"
        inactive-text="总览模式"
        active-value="userInfoEdit"
        inactive-value="userInfoView"
        @change="mod"
    /></el-row>

    <!-- <component :is="currentModel" /> -->
    <userInfoEdit v-if="abc" />
    <userInfoView v-else-if="abc"/>
  </div>
</template>

<script>
import userInfoEdit from './components/edit'
import userInfoView from './components/view'

export default {
  name: 'UserInfoTest',
  components: { userInfoEdit, userInfoView },
  data() {
    return {
      currentModel: 'userInfoView',
      abc: false
    }
  },
  methods: {
    mod() {
      this.abc = true
    }
  }
}
</script>

