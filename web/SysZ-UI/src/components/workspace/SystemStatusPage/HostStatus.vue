<template>
  <div class="HostStatus_container">
    <div class="HostStatus_card">
      <el-card class="box-card">
        <template #header>
          <div class="card-header">
            <span class="Host_info_title">Host information</span>
          </div>
        </template>
        <div class="Host_info_name">Host Name:
          <div class="Host_info_value">{{ props.value.host_name }}</div>
        </div>
        <hr/>
        <div class="Host_info_name">Host OS:
          <div class="Host_info_value">{{ props.value.host_os }}</div>
        </div>
        <hr/>
        <div class="Host_info_name">Host OS Version:
          <div class="Host_info_value">{{ props.value.host_os_version }}</div>
        </div>
        <hr/>
        <div class="Host_info_name">Host Kernel Version:
          <div class="Host_info_value">{{ props.value.host_kernel_version }}</div>
        </div>
        <hr/>
        <div class="Host_info_name">Host Kernel Arch:
          <div class="Host_info_value">{{ props.value.host_kernel_arch }}</div>
        </div>
        <hr/>
        <div class="Host_info_name">Host ID:
          <div class="Host_info_value">{{ props.value.host_id }}</div>
        </div>
        <hr/>
      </el-card>
    </div>
  </div>
</template>

<script lang="ts" setup>

const props = defineProps(['value'])
</script>

<style scoped>
/*总布局*/
.HostStatus_container {
  align-items: center;
  padding: 10px;
  margin: auto;
}

/*卡片*/
.HostStatus_card {
  width: 600px;
  margin: auto;
}

/*卡片布局*/
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  font-size: 20px;
}

.box-card {
  width: 600px;
  height: 400px;
  border-radius: 6px;
  transition: all 0.3s;
}

.box-card:hover {
  transition: all 0.3s;
  border-radius: 6px;
}

.light .box-card:hover {
  background: #ffffff;
  box-shadow: 7px 7px 14px #c4c4c4,
  -7px -7px 14px #ffffff;
  transition: all 0.3s;
}

.dark .box-card:hover {
  background: #1d1e1f;
  box-shadow: 7px 7px 14px #161718,
  -7px -7px 14px #242526;
  transition: all 0.3s;
}

/*名称*/
.Host_info_title {
  font-size: 20px;
  text-align: center;
  width: 100%;
  font-weight: bold;
}

/*Host信息字段*/
.Host_info_name {
  font-size: 14px;
  padding: 6px 7px;
  text-align: left;
  line-height: 35px;
  transition: all 0.3s;
}

.dark .Host_info_name:hover {
  background-color: #edede9;
  color: #303133;
  font-size: 16px;
  border-radius: 15px;
  cursor: pointer;
  transition: all 0.3s;
}

.light .Host_info_name:hover {
  background-color: #caf0f8;
  font-size: 16px;
  border-radius: 15px;
  cursor: pointer;
  transition: all 0.3s;
}

.Host_info_value {
  font-weight: normal;
  text-align: center;
  display: inline-block;
  width: 320px;
  float: right;
}
</style>