<template>
  <MirzaModal id="modalFilter" ref="modalFilter" title="Filter Items" @submit="submitDataFilter">
    <div class="mb-3">
      <label class="form-label">Name start with</label>
      <input type="text" class="form-control" placeholder="Name Like" v-model="state.filter.name_like">
    </div>

    <div class="mb-3">
      <label class="form-label">Date Start</label>
      <input type="date" class="form-control" placeholder="Date Start" v-model="state.filter.date_start">
    </div>

    <div class="mb-3">
      <label class="form-label">Date End</label>
      <input type="date" class="form-control" placeholder="Date End" v-model="state.filter.date_end">
    </div>
  </MirzaModal>
</template>

<script setup>
import MirzaModal from "../../components/modal/MirzaModal.vue";
import {state} from "./state.js";
import {ref} from "vue";

const modalFilter = ref()

const emit = defineEmits(["submit"])

const submitDataFilter = () => {
  state.filter.page = 1
  emit("submit")
  hideModal()
}

const showModal = () => {
  modalFilter.value.showModal()
}

const hideModal = () => {
  modalFilter.value.hideModal()
}

defineExpose({showModal, hideModal})

</script>

<style scoped>

</style>