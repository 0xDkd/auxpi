<template>
  <el-row :gutter="10">
    <el-col :xs="{span: 24}" :sm="{span: 12}" :md="{span: 8}" :lg="{span: 8}" :xl="{span: 8}" style="margin-bottom:30px;">
      <box-card/>
    </el-col>

    <el-col :xs="{span: 24}" :sm="{span: 12}" :md="{span: 16}" :lg="{span: 16}" :xl="{span: 16}" style="margin-bottom:30px;">
      <el-card shadow="hover" >
        <el-table v-loading="tableLoading" :data="tableData">
          <el-table-column label="缩略图">
            <template slot-scope="scope">
              <img :src="scope.row.link" height="50px" width="50px" >
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

        <pagination v-show="total>0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.limit" @pagination="getList"/>

      </el-card>
    </el-col>
  </el-row>

</template>

<script>

import BoxCard from '@/components/InfoBox'
import Pagination from '@/components/Pagination' // Secondary package based on el-pagination
import { getUserImages } from '@/api/user'

export default {
  name: 'UserInfoView',
  components: {
    BoxCard,
    Pagination
  },
  props: {
    currentModel: {
      required: true,
      type: String
    }
  },
  data() {
    return {
      total: 0,
      // 分页参数
      listQuery: {
        page: 1,
        limit: 20
      },
      tableData: [],
      tableLoading: true
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.tableLoading = true
      console.log(this.listQuery)
      const uid = this.$route.params.id
      getUserImages(uid, this.listQuery).then(resp => {
        this.tableData = resp.list
        this.total = resp.total
        this.tableLoading = false
      })
    }
  }
}
</script>

