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
      <th class="text-start">Action</th>
      <th class="text-start">Date</th>
      <th class="text-start">Description</th>
      <th class="text-start">Account</th>
      <th class="text-start">SubAccount</th>
      <th class="text-end">Amount</th>
      <th class="text-end">Balance</th>
    </tr>
    <tr>
      <th class="text-start"></th>
      <th class="text-start"></th>
      <th class="text-start"></th>
      <th class="text-start"></th>
      <th class="text-start">
        <input type="text" class="form-control" placeholder="Name Like" v-model="state.filter.sub_account_name">
      </th>
      <th class="text-end"></th>
      <th class="text-end"></th>
    </tr>
    </thead>
    <tbody>
    <tr v-for="item in state.items" :key="item.id">

      <td>
        <div class="btn-group">
          <button type="button" class="btn btn-warning btn-sm" @click="showModalDetail(item)">Detail</button>
        </div>
      </td>
      <td>{{item.date}}</td>
      <td>{{item.journal.description}}</td>
      <td>{{item.subAccount.parentAccount.name}}</td>
      <td>{{item.subAccount.name}}</td>
      <td class="text-end">{{item.amount}}</td>
      <td class="text-end">{{item.balance}}</td>

    </tr>
    </tbody>
  </table>


<!--  <MirzaTable :fields="fields" :items="state.items">-->
<!--    <template #action="{item}">-->
<!--      <div class="btn-group">-->
<!--        <button type="button" class="btn btn-warning btn-sm" @click="showModalDetail(item)">Detail</button>-->
<!--      </div>-->
<!--    </template>-->
<!--    <template #description="{item}">-->
<!--      {{item.journal.description}}-->
<!--    </template>-->
<!--    <template #subAccount="{item}">-->
<!--      ({{item.subAccount.code}}) {{item.subAccount.name}}-->
<!--    </template>-->
<!--    <template #amount="{item}">-->
<!--      <div class="text-end">{{item.amount}}</div>-->
<!--    </template>-->
<!--    <template #balance="{item}">-->
<!--      <div class="text-end">{{item.balance}}</div>-->
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

  const url = `${BASE_URL}/wallet/WLT01/subaccountbalance`

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
  {header: "Desc", fieldName: "description",},
  {header: "Date", fieldName: "date",},
  {header: "Subaccount", fieldName: "subAccount",},
  {header: "Amount", fieldName: "amount",},
  {header: "Balance", fieldName: "balance",},
]

reload()
</script>

<style scoped>

</style>