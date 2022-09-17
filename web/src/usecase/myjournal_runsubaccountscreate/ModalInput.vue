<template>
  <MirzaModal id="modalRunSubAccountsCreate" ref="modalRunSubAccountsCreate" title="RunSubAccountsCreate" @submit="submitRunSubAccountsCreate">
    
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

const modalRunSubAccountsCreate = ref()

const emit = defineEmits(["submit"])

const payload = reactive({
  data: { 
  }
})

const submitRunSubAccountsCreate = async () => {

  const url = `${BASE_URL}/runsubaccountscreate`

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
  modalRunSubAccountsCreate.value.showModal()
}

const hideModal = () => {
  modalRunSubAccountsCreate.value.hideModal()
}

defineExpose({showModal, hideModal})

</script>

<style scoped>

</style>