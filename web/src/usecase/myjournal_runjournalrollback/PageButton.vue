<template>

  <button type="button" class="btn btn-outline-primary" @click="showModalRunJournalRollback">Show</button>

  <ViewModalRunJournalRollback ref="modalRunJournalRollback" @submit="reload"></ViewModalRunJournalRollback>

</template>

<script setup>
import swal from "sweetalert2";
import {ref} from "vue";

import ViewModalRunJournalRollback from "./ModalInput.vue";
const modalRunJournalRollback = ref()
const showModalRunJournalRollback = () => modalRunJournalRollback.value.showModal()

const reload = () => swal.fire({ icon: 'success', title: 'Success', text: 'Usecase executed', })

</script>

<style scoped>

</style>