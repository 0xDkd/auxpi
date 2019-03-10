<template>
  <div class="app-container">

    <el-table v-loading="listLoading" :data="list" border fit highlight-current-row style="width: 100%">

      <el-table-column align="center" label="ID" width="80">
        <template slot-scope="scope">
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>

      <el-table-column align="center" label="用户名">
        <template slot-scope="scope">
          <span>{{ scope.row.username }}</span>
        </template>
      </el-table-column>

      <el-table-column align="center" label="注册邮箱">
        <template slot-scope="scope">
          <span> {{ scope.row.email }}</span>
        </template>
      </el-table-column>
      <el-table-column align="center" label="注册时间">
        <template slot-scope="scope">
          <span>{{ scope.row.created_on | parseTime('{y}-{m}-{d} {h}:{i}') }}</span>
        </template>
      </el-table-column>

      <el-table-column class-name="status-col" label="用户组" width="130px">
        <template slot-scope="scope">
          <el-tag :type="scope.row.role.name | statusFilter">{{ scope.row.role.display_name }}</el-tag>
        </template>
      </el-table-column>

      <el-table-column align="center" label="Actions" >
        <template slot-scope="scope">
          <el-button type="success" size="small" icon="el-icon-circle-check-outline" @click="jumpToUserDetailInfo(scope.row.id)">查看详情</el-button>
          <el-button type="danger" size="small" icon="el-icon-delete" @click="deleteU(scope.row.id,scope.$index)" >Delete</el-button>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total>0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.limit" @pagination="getList"/>

  </div>
</template>

<script>
import { getUser, DeleteUser } from '@/api/user'
import Pagination from '@/components/Pagination' // Secondary package based on el-pagination

export default {
  name: 'InlineEditTable',
  components: { Pagination },
  filters: {
    statusFilter(role) {
      const roleMap = {
        admin: 'success',
        normalUser: 'info',
        blockUser: 'danger'
      }
      return roleMap[role]
    }
  },
  data() {
    return {
      list: [],
      listLoading: true,
      listQuery: {
        page: 1,
        limit: 10
      },
      total: 0,
      canDelete: true
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.listLoading = true
      getUser(this.listQuery).then(resp => {
        this.list = resp.list
        this.total = resp.total
        this.listLoading = false
      })
    },
    cancelEdit(row) {
      row.title = row.originalTitle
      row.edit = false
      this.$message({
        message: 'The title has been restored to the original value',
        type: 'warning'
      })
    },
    confirmEdit(row) {
      row.edit = false
      row.originalTitle = row.title
      this.$message({
        message: 'The title has been edited',
        type: 'success'
      })
    },
    //
    jumpToUserDetailInfo(uid) {
      this.$router.push({ name: 'usersInfoView', params: { id: uid }})
    },
    deleteU(uid, index) {
      this.$confirm('此操作将永久删除该用户, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.list.forEach(e => {
          if (e.role_id === 1) {
            this.canDelete = false
            return
          }
        })

        if (!this.canDelete) {
          this.$message({
            type: 'error',
            message: '管理员无法删除管理员'
          })
          return
        }

        this.list.splice(index, 1)

        DeleteUser(uid).then(resp => {
          this.$message({
            type: 'success',
            message: '删除成功!'
          })
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        })
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
