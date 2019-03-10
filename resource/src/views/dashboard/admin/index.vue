<template>
  <div class="dashboard-editor-container">

    <github-corner style="position: absolute; top: 0px; border: 0; right: 0;"/>

    <panel-group
      :users-report-count="usersReportCount"
      :all-images-report-count="allImagesReportCount"
      :local-images-report-count="localImagesReportCount"
      :api-report-count="apiReportCount"
      @handleSetLineChartData="handleSetLineChartData"
    />

    <el-row v-loading="lineLoading" style="background:#fff;padding:16px 16px 0;margin-bottom:32px;">

      <line-chart v-if="showline" :chart-y-data="chartYData" :chart-x-data="chartXData"/>

      <div v-else class="lineDiv">
        <el-alert
          title="此部分目前还没有任何数据哟~"
          type="success"
        />
      </div>

    </el-row>

    <el-row :gutter="32">
      <el-col :xs="24" :sm="24" :lg="24">
        <div class="chart-wrapper">
          <pie-chart/>
        </div>
      </el-col>
    </el-row>

    <el-row :gutter="8">
      <el-col :xs="{span: 24}" :sm="{span: 12}" :md="{span: 12}" :lg="{span: 12}" :xl="{span: 12}" style="margin-bottom:30px;">
        <box-card/>
      </el-col>
    </el-row>

    <!-- <el-card class="box-card">
      <div slot="header" class="clearfix">
        <span>GoRoutines</span>
      </div>
      <div v-html="goroutines"/>
    </el-card> -->

  </div>
</template>

<script>
import GithubCorner from '@/components/GithubCorner'
import PanelGroup from './components/PanelGroup'
import LineChart from './components/LineChart'
import PieChart from './components/PieChart'
import BoxCard from './components/BoxCard'
import { getUserSevenReport, getApiSevenReport, getAllImageSevenReport, getLocalImageSevenReport } from '@/api/dashbroad'
const lineChartData = {}

const chartYData = {}

const chartXData = {}

export default {
  name: 'DashboardAdmin',
  components: {
    GithubCorner,
    PanelGroup,
    LineChart,
    PieChart,
    BoxCard
  },
  data() {
    return {
      lineChartData: lineChartData.newVisitis,
      chartYData: [],
      chartXData: [],
      showline: false,
      lineLoading: true,
      userTime: [],
      apiTime: [],
      allImageTime: [],
      localImageTime: [],
      userReport: [],
      apiReport: [],
      allImageReport: [],
      localImageReport: [],
      usersReportCount: 0,
      allImagesReportCount: 0,
      localImagesReportCount: 0,
      apiReportCount: 0,
      heap: '',
      goroutines: '',
      gc: '',
      thread: ''

    }
  },
  created() {
    this.getAllReport()
    // this.getAuxpiInfo()
  },
  methods: {
    handleSetLineChartData(type) {
      this.chartYData = chartYData[type]
      this.chartXData = chartXData[type]

      if (this.chartYData.length === 0 || this.chartXData.length === 0) {
        this.showline = false
        return
      }
      this.showline = true
    },
    getAllReport() {
      const actionUser = getUserSevenReport().then(resp => {
        resp.list.forEach((e, i) => {
          this.userTime[i] = e.date
          this.userReport[i] = e.number
          this.usersReportCount += e.number
        })
        chartYData.userReport = this.userReport
        chartXData.userReport = this.userTime
      })

      const actionApi = getApiSevenReport().then(resp => {
        resp.list.forEach((e, i) => {
          this.apiTime[i] = e.date
          this.apiReport[i] = e.number
          this.apiReportCount += e.number
        })
        chartYData.apiReport = this.apiReport
        chartXData.apiReport = this.apiTime
      })

      const actionAllImage = getAllImageSevenReport().then(resp => {
        resp.list.forEach((e, i) => {
          this.allImageTime[i] = e.date
          this.allImageReport[i] = e.number
          this.allImagesReportCount += e.number
        })
        chartYData.allImageReport = this.allImageReport
        chartXData.allImageReport = this.allImageTime
      })

      const actionLocalImage = getLocalImageSevenReport().then(resp => {
        resp.list.forEach((e, i) => {
          this.localImageTime[i] = e.date
          this.localImageReport[i] = e.number
          this.localImagesReportCount += e.number
        })
        chartYData.localImageReport = this.localImageReport
        chartXData.localImageReport = this.localImageTime
      })

      Promise.all([actionUser, actionApi, actionAllImage, actionLocalImage]).then(() => {
        // 设定到第一个上面
        this.chartYData = this.allImageReport
        this.chartXData = this.allImageTime
        this.showline = true
        if (this.chartYData.length === 0 || this.chartXData.length === 0) {
          this.showline = false
        }

        this.lineLoading = false
      })
    }
    // getAuxpiInfo() {
    //   getAuxpiSystemInfo().then((resp) => {
    //     this.goroutines = resp.data.goroutines
    //     this.heap = resp.data.heap
    //     console.log(this.heap)

    //     this.thread = resp.data.thread
    //     this.gc = resp.data.gc
    //   })
    // }
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
</style>
