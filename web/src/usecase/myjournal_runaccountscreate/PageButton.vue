<template>

  <button type="button" class="btn btn-outline-primary" @click="showModalRunAccountsCreate">Show</button>

  <ViewModalRunAccountsCreate ref="modalRunAccountsCreate" @submit="reload"></ViewModalRunAccountsCreate>

</template>

<script setup>
import swal from "sweetalert2";
import {ref} from "vue";

import ViewModalRunAccountsCreate from "./ModalInput.vue";
const modalRunAccountsCreate = ref()
const showModalRunAccountsCreate = () => modalRunAccountsCreate.value.showModal()

const reload = () => swal.fire({ icon: 'success', title: 'Success', text: 'Usecase executed', })

</script>

<style scoped>

</style>