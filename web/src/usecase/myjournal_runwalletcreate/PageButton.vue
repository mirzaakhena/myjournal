<template>

  <button type="button" class="btn btn-outline-primary" @click="showModalRunWalletCreate">Show</button>

  <ViewModalRunWalletCreate ref="modalRunWalletCreate" @submit="reload"></ViewModalRunWalletCreate>

</template>

<script setup>
import swal from "sweetalert2";
import {ref} from "vue";

import ViewModalRunWalletCreate from "./ModalInput.vue";
const modalRunWalletCreate = ref()
const showModalRunWalletCreate = () => modalRunWalletCreate.value.showModal()

const reload = () => swal.fire({ icon: 'success', title: 'Success', text: 'Usecase executed', })

</script>

<style scoped>

</style>