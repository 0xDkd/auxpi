<template>
  <el-row :gutter="40" class="panel-group">
    <!-- 一个星期内的上
    传链接统计 -->
    <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
      <div class="card-panel" @click="handleSetLineChartData('allImageReport')">
        <div class="card-panel-icon-wrapper icon-records">
          <svg-icon icon-class="data-base" class-name="card-panel-icon" />
        </div>
        <div class="card-panel-description">
          <div class="card-panel-text">七天内新记录 总数</div>
          <count-to :start-val="0" :end-val="allImagesReportCount" :duration="3000" class="card-panel-num"/>
        </div>
      </div>
    </el-col>
    <!-- 一星期内的本地图片增加统计 -->
    <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
      <div class="card-panel" @click="handleSetLineChartData('localImageReport')">
        <div class="card-panel-icon-wrapper icon-Images">
          <svg-icon icon-class="images" class-name="card-panel-icon" />
        </div>
        <div class="card-panel-description">
          <div class="card-panel-text">七天内本地新图片 总数</div>
          <count-to :start-val="0" :end-val="localImagesReportCount" :duration="3200" class="card-panel-num"/>
        </div>
      </div>
    </el-col>
    <!-- 一个星期内用户统计 -->
    <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
      <div class="card-panel" @click="handleSetLineChartData('userReport')">
        <div class="card-panel-icon-wrapper icon-users">
          <svg-icon icon-class="users" class-name="card-panel-icon" />
        </div>
        <div class="card-panel-description">
          <div class="card-panel-text">七天新用户 总数</div>
          <count-to :start-val="0" :end-val="usersReportCount" :duration="2600" class="card-panel-num"/>
        </div>
      </div>
    </el-col>
    <!-- 一星期内的api 调用统计 -->
    <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
      <div class="card-panel" @click="handleSetLineChartData('apiReport')">
        <div class="card-panel-icon-wrapper icon-api">
          <svg-icon icon-class="cloud" class-name="card-panel-icon" />
        </div>
        <div class="card-panel-description">
          <div class="card-panel-text">七天内 api 调用 总数</div>
          <count-to :start-val="0" :end-val="apiReportCount" :duration="3600" class="card-panel-num"/>
        </div>
      </div>
    </el-col>

    <!-- 30 天内人口总数统计 -->
    <!-- <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
      <div class="card-panel" @click="handleSetLineChartData('shoppings')">
        <div class="card-panel-icon-wrapper icon-api">
          <svg-icon icon-class="cloud" class-name="card-panel-icon" />
        </div>
        <div class="card-panel-description">
          <div class="card-panel-text">api 调用次数</div>
          <count-to :start-val="0" :end-val="13600" :duration="3600" class="card-panel-num"/>
        </div>
      </div>
    </el-col> -->

    <!-- 30 天内本地图片总数统计 -->
    <!-- <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
      <div class="card-panel" @click="handleSetLineChartData('shoppings')">
        <div class="card-panel-icon-wrapper icon-api">
          <svg-icon icon-class="cloud" class-name="card-panel-icon" />
        </div>
        <div class="card-panel-description">
          <div class="card-panel-text">api 调用次数</div>
          <count-to :start-val="0" :end-val="13600" :duration="3600" class="card-panel-num"/>
        </div>
      </div>
    </el-col> -->
    <!-- star 图片图片30天内统计，显示目前总数 -->
    <!-- <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
      <div class="card-panel" @click="handleSetLineChartData('shoppings')">
        <div class="card-panel-icon-wrapper icon-api">
          <svg-icon icon-class="cloud" class-name="card-panel-icon" />
        </div>
        <div class="card-panel-description">
          <div class="card-panel-text">api 调用次数</div>
          <count-to :start-val="0" :end-val="13600" :duration="3600" class="card-panel-num"/>
        </div>
      </div>
    </el-col> -->

    <!-- auxpi 版本号 -->
    <!-- <el-col :xs="12" :sm="12" :lg="6" class="card-panel-col">
      <div class="card-panel" @click="handleSetLineChartData('shoppings')">
        <div class="card-panel-icon-wrapper icon-api">
          <svg-icon icon-class="cloud" class-name="card-panel-icon" />
        </div>
        <div class="card-panel-description">
          <div class="card-panel-text">api 调用次数</div>
          <count-to :start-val="0" :end-val="13600" :duration="3600" class="card-panel-num"/>
        </div>
      </div>
    </el-col> -->

  </el-row>
</template>

<script>
import CountTo from 'vue-count-to'

export default {
  components: {
    CountTo
  },
  props: {
    usersReportCount: {
      type: Number,
      required: true,
      default: 0
    },
    allImagesReportCount: {
      type: Number,
      required: true,
      default: 0
    },
    localImagesReportCount: {
      type: Number,
      required: true,
      default: 0
    },
    apiReportCount: {
      type: Number,
      required: true,
      default: 0
    }
  },
  methods: {
    handleSetLineChartData(type) {
      this.$emit('handleSetLineChartData', type)
    }
  }
}
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
.panel-group {
  margin-top: 18px;
  .card-panel-col{
    margin-bottom: 32px;
  }
  .card-panel {
    height: 108px;
    cursor: pointer;
    font-size: 12px;
    position: relative;
    overflow: hidden;
    color: #666;
    background: #fff;
    box-shadow: 4px 4px 40px rgba(0, 0, 0, .05);
    border-color: rgba(0, 0, 0, .05);
    &:hover {
      .card-panel-icon-wrapper {
        color: #fff;
      }
      .icon-users {
         background: #40c9c6;
      }
      .icon-records {
        background: #36a3f7;
      }
      .icon-Images {
        background: #f4516c;
      }
      .icon-api {
        background: #34bfa3
      }
    }
    .icon-users {
      color: #40c9c6;
    }
    .icon-records {
      color: #36a3f7;
    }
    .icon-Images {
      color: #f4516c;
    }
    .icon-api {
      color: #34bfa3
    }
    .card-panel-icon-wrapper {
      float: left;
      margin: 14px 0 0 14px;
      padding: 16px;
      transition: all 0.38s ease-out;
      border-radius: 6px;
    }
    .card-panel-icon {
      float: left;
      font-size: 48px;
    }
    .card-panel-description {
      float: right;
      font-weight: bold;
      margin: 26px;
      margin-left: 0px;
      .card-panel-text {
        line-height: 18px;
        color: rgba(0, 0, 0, 0.45);
        font-size: 16px;
        margin-bottom: 12px;
      }
      .card-panel-num {
        font-size: 20px;
      }
    }
  }
}
</style>
