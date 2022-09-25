import {computed, reactive} from "vue";

export const state = reactive({
    items: [],
    filter: {
        page: 1,
        size: 20,
        sub_account_name: '',
    },
    totalItems: 0,
    pagePerTotalRecord: computed(()=> `${state.filter.page} / ${getNumberOfPage()}` )
})

export const getNumberOfPage = () => Math.ceil(state.totalItems / state.filter.size)