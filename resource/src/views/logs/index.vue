<template>
  <div class="app-container">

    <el-table v-loading="listLoading" :data="list" border fit highlight-current-row style="width: 100%">

      <el-table-column align="center" label="ID" width="80">
        <template slot-scope="scope">
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>

      <el-table-column align="center" label="日志时间">
        <template slot-scope="scope">
          <span>{{ scope.row.created_on | parseTime('{y}-{m}-{d} {h}:{i}') }}</span>
        </template>
      </el-table-column>

      <el-table-column align="center" label="日志级别">
        <template slot-scope="scope">
          <el-tag :type="scope.row.level | statusFilter">{{ scope.row.level }}</el-tag>

        </template>
      </el-table-column>
      <el-table-column align="center" label="日志内容">
        <template slot-scope="scope">
          <span>{{ scope.row.content }}</span>
        </template>
      </el-table-column>

      <el-table-column class-name="status-col" label="日志所属" width="130px">
        <template slot-scope="scope">
          <span> {{ scope.row.type }}</span>
        </template>
      </el-table-column>

      <el-table-column class-name="status-col" label="执行者" width="130px">
        <template slot-scope="scope">
          <span> {{ scope.row.part }}</span>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total>0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.limit" @pagination="getList"/>

  </div>
</template>

<script>
import { getLogs } from '@/api/log'
import Pagination from '@/components/Pagination' // Secondary package based on el-pagination

export default {
  name: 'LogsTable',
  components: { Pagination },
  filters: {
    statusFilter(level) {
      const levelMap = {
        INFO: 'info',
        ERROR: 'danger',
        NONE: 'success',
        DEBUG: 'primary',
        WARNING: 'warning'
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
      getLogs(this.listQuery).then(resp => {
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
