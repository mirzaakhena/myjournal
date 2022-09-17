<template>

  <button type="button" class="btn btn-outline-primary" @click="showModalRunJournalCreate">Show</button>

  <ViewModalRunJournalCreate ref="modalRunJournalCreate" @submit="reload"></ViewModalRunJournalCreate>

</template>

<script setup>
import swal from "sweetalert2";
import {ref} from "vue";

import ViewModalRunJournalCreate from "./ModalInput.vue";
const modalRunJournalCreate = ref()
const showModalRunJournalCreate = () => modalRunJournalCreate.value.showModal()

const reload = () => swal.fire({ icon: 'success', title: 'Success', text: 'Usecase executed', })

</script>

<style scoped>

</style>