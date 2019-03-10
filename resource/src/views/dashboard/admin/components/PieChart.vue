<template>
  <div v-loading="pieChartLoading" :class="className" :style="{height:height,width:width}"/>
</template>

<script>
import echarts from 'echarts'
require('echarts/theme/macarons') // echarts theme
import { debounce } from '@/utils'
import { getStorePercent } from '@/api/dashbroad'

export default {
  props: {
    className: {
      type: String,
      default: 'chart'
    },
    width: {
      type: String,
      default: '100%'
    },
    height: {
      type: String,
      default: '300px'
    }
  },
  data() {
    return {
      chart: null,
      storeName: [],
      storeValue: [],
      fakerData: [
        { value: 320, name: '本地图床' },
        { value: 240, name: '新浪图床' },
        { value: 149, name: '搜狗图床' },
        { value: 100, name: 'SMMS' },
        { value: 59, name: 'CC 图床' },
        { value: 240, name: '京东图床' },
        { value: 240, name: '百度图床' },
        { value: 240, name: '阿里图床' },
        { value: 240, name: '掘金图床' }
      ],
      pieChartLoading: true
    }
  },
  mounted() {
    getStorePercent().then(resp => {
      this.storeValue = resp.list
      // 获取储存方式
      this.storeValue.forEach((e, i) => {
        this.storeName[i] = e.name
      })

      this.initChart()
      this.__resizeHandler = debounce(() => {
        if (this.chart) {
          this.chart.resize()
        }
      }, 100)
      window.addEventListener('resize', this.__resizeHandler)
    })
  },
  beforeDestroy() {
    if (!this.chart) {
      return
    }
    window.removeEventListener('resize', this.__resizeHandler)
    this.chart.dispose()
    this.chart = null
  },
  methods: {
    initChart() {
      this.chart = echarts.init(this.$el, 'macarons')

      this.chart.setOption({
        tooltip: {
          trigger: 'item',
          formatter: '{a} <br/>{b} : {c} ({d}%)'
        },
        legend: {
          left: 'center',
          bottom: '10',
          data: this.storeName
        },
        calculable: true,
        series: [
          {
            name: 'WEEKLY WRITE ARTICLES',
            type: 'pie',
            roseType: 'radius',
            radius: [15, 95],
            center: ['50%', '38%'],
            data: this.storeValue,
            animationEasing: 'cubicInOut',
            animationDuration: 2600
          }
        ]
      })
    }
  }
}
</script>
