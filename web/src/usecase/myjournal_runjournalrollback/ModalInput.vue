<template>
  <MirzaModal id="modalRunJournalRollback" ref="modalRunJournalRollback" title="RunJournalRollback" @submit="submitRunJournalRollback">
    
  </MirzaModal>
</template>

<script setup>
import MirzaModal from "../../components/modal/MirzaModal.vue";
import {BASE_URL} from "../shared.js";
import {ref} from "vue";
import to from "await-to-js";
import axios from "axios";
import swal from "sweetalert2";
import {reactive} from "vue";

const modalRunJournalRollback = ref()

const emit = defineEmits(["submit"])

const payload = reactive({
  data: { 
  }
})

const submitRunJournalRollback = async () => {

  const url = `${BASE_URL}/runjournalrollback`

  const [err, res] = await to(axios.post(`${url}`, payload.data).catch((err) => Promise.reject(err)))

  if (err) {
    await swal.fire({ icon: 'error', title: 'Oops...', text: err.response.data.errorMessage, })
    return
  }

  console.log(res.data.data)
  emit("submit")
  hideModal()
}

const showModal = () => {
  modalRunJournalRollback.value.showModal()
}

const hideModal = () => {
  modalRunJournalRollback.value.hideModal()
}

defineExpose({showModal, hideModal})

</script>

<style scoped>

</style>