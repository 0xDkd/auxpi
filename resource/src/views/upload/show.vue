<template>
  <div class="app-container">

    <el-table v-loading="listLoading" :data="list" border highlight-current-row style="width: 100%">

      <el-table-column
        type="selection"
        width="55"/>

      <el-table-column align="center" label="上传时间" width="120">
        <template slot-scope="scope">
          <span>{{ scope.row.created_on | parseTime('{y}-{m}-{d}') }}</span>
        </template>
      </el-table-column>

      <el-table-column align="center" label="hash">
        <template slot-scope="scope">
          {{ scope.row.hash }}
        </template>
      </el-table-column>
      <el-table-column align="center" label="根链接">
        <template slot-scope="scope">
          <span>{{ scope.row.root_url }}</span>
        </template>
      </el-table-column>

      <el-table-column class-name="status-col" label="目前归属" width="130px">
        <template slot-scope="scope">
          <el-tag type="success">{{ scope.row.store.name }}</el-tag>
        </template>
      </el-table-column>

      <el-table-column class-name="status-col" label="分发状态" width="130px">
        <template slot-scope="scope">
          <el-tag :type="scope.row.url | typeFilter">{{ scope.row.store.name===""?'正在分发':'已完成' }}</el-tag>
        </template>
      </el-table-column>

      <el-table-column class-name="status-col" label="操作" width="130px" fixed="true">
        <template slot-scope="scope">
          <span> {{ scope.row.part }}</span>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total>0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.limit" @pagination="getList"/>

  </div>
</template>

<script>
import { getDispatchList } from '@/api/dispatch'
import Pagination from '@/components/Pagination' // Secondary package based on el-pagination

export default {
  name: 'LogsTable',
  components: { Pagination },
  filters: {
    typeFilter(level) {
      const levelMap = {
        '': 'warning'
      }
      return levelMap[level]
    },
    statusFilter(level) {
      const levelMap = {
        '': '正在分发'
      }
      return levelMap[level]
    }
  },
  data() {
    return {
      list: null,
      listLoading: true,
      listQuery: {
        page: 1,
        limit: 12
      },
      total: 0
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.listLoading = true
      getDispatchList(this.listQuery).then(resp => {
        this.list = resp.list
        this.total = resp.total
        this.listLoading = false
      })
    }

  }
}
</script>

<style scoped>
.edit-input {
  padding-right: 100px;
}
.cancel-btn {
  position: absolute;
  right: 15px;
  top: 10px;
}
</style>
