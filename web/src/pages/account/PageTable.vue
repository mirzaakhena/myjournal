<template>

  <div class="d-flex justify-content-between">
    <div>
      <div class="btn-group">
        <button type="button" class="btn btn-primary" @click="reload">Reload</button>
        <button type="button" class="btn btn-dark" @click="showModalFilter">Filter</button>
      </div>
    </div>
    <div>
      <div class="btn-group float-end">
        <button type="button" class="btn btn-success" @click="previousPage">Prev Page</button>
        <button type="button" class="btn btn-primary" @click="showModalPaging">{{state.pagePerTotalRecord}}</button>
        <button type="button" class="btn btn-dark" @click="nextPage">Next Page</button>
      </div>
    </div>
  </div>

  <table class="table table-sm">

    <thead>
    <tr>
      <th>Action</th>
      <th>Side</th>
      <th>Name</th>
      <th>Code</th>
    </tr>
    </thead>
    <tbody>
    <tr v-for="item in state.items" :key="item.id" :class="item.level===1?'table-secondary':''">
      <td>
        <div class="btn-group">
          <button type="button" class="btn btn-warning btn-sm" @click="showModalDetail(item)">Detail</button>
        </div>
      </td>
      <td>
        <span class="badge" :class="item.side==='ACTIVA'?'text-bg-primary': 'text-bg-danger'">{{ item.side }}</span>
      </td>
      <td>
        <template v-for="(n,index) in item.level-1">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</template> {{item.name}}
      </td>
      <td>{{ item.code }}</td>
    </tr>
    </tbody>
  </table>

<!--  <MirzaTable :fields="fields" :items="state.items">-->
<!--    <template #action="{item}">-->
<!--      <div class="btn-group">-->
<!--        <button type="button" class="btn btn-warning btn-sm" @click="showModalDetail(item)">Detail</button>-->
<!--      </div>-->
<!--    </template>-->
<!--    <template #name="{item}">-->
<!--      <template v-for="(n,index) in item.level-1">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</template> {{item.name}}-->
<!--    </template>-->
<!--    <template #side="{item}">-->
<!--      <span class="badge" :class="item.side==='ACTIVA'?'text-bg-primary': 'text-bg-danger'">{{ item.side }}</span>-->
<!--    </template>-->

<!--  </MirzaTable>-->

  <ViewModalDetail ref="modalDetail" @submit="reload"></ViewModalDetail>

  <ViewModalFilter ref="modalFilter" @submit="reload"></ViewModalFilter>

  <ViewModalPaging ref="modalPaging" @submit="reload"></ViewModalPaging>

</template>

<script setup>
import MirzaTable from "../../components/table/MirzaTable.vue";
import {BASE_URL} from "../shared.js";
import {state, getNumberOfPage} from "./state.js";
import {ref} from "vue";
import to from "await-to-js";
import axios from "axios";
import swal from "sweetalert2";

import ViewModalDetail from "./ModalDetail.vue";
const modalDetail = ref()
const showModalDetail = (payload) => modalDetail.value.showModal(payload)

import ViewModalPaging from "./ModalPaging.vue";
const modalPaging = ref()
const showModalPaging = () => modalPaging.value.showModal()

import ViewModalFilter from "./ModalFilter.vue";
const modalFilter = ref()
const showModalFilter = () => modalFilter.value.showModal()

const nextPage = () => {
  if (state.filter.page + 1 <= getNumberOfPage()) {
    state.filter.page++
    reload()
  }
}

const previousPage = () => {
  if (state.filter.page - 1 > 0) {
    state.filter.page--
    reload()
  }
}

const reload = async () => {

  const url = `${BASE_URL}/wallet/WLT01/account`

  const requestConfig = { params: { ...state.filter } }

  const [err, res] = await to(axios.get(url, requestConfig).catch((err) => Promise.reject(err)))

  if (err) {
    await swal.fire({ icon: 'error', title: 'Oops...', text: err.response.data.errorMessage, })
    return
  }

  state.items = res.data.data.items
  state.totalItems = res.data.data.count
}

const fields = [
  {header: "Action", fieldName: "action",},
  {header: "Side", fieldName: "side",},
  {header: "Name", fieldName: "name",},
  {header: "Code", fieldName: "code",},
]

reload()
</script>

<style scoped>

</style>