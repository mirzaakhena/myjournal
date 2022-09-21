<template>

  <button type="button" class="btn btn-outline-primary" @click="showModalRunSubAccountsCreate">Show</button>

  <ViewModalRunSubAccountsCreate ref="modalRunSubAccountsCreate" @submit="reload"></ViewModalRunSubAccountsCreate>

</template>

<script setup>
import swal from "sweetalert2";
import {ref} from "vue";

import ViewModalRunSubAccountsCreate from "./ModalInput.vue";
const modalRunSubAccountsCreate = ref()
const showModalRunSubAccountsCreate = () => modalRunSubAccountsCreate.value.showModal()

const reload = () => swal.fire({ icon: 'success', title: 'Success', text: 'Usecase executed', })

</script>

<style scoped>

</style>