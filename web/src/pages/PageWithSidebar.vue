<template>
  <Sidebar title="App" :sections="sections"/>
  <div style="margin-left: 300px; margin-right: 20px; margin-top: 20px">
    <router-view/>
  </div>
</template>

<script setup>

import Sidebar from "../components/sidebar/Sidebar.vue";

const sections = [
  {
    header: "Account", menus: [
      {href: '/account', title: 'Account', icon: '#speedometer2',},
      {href: '/subaccount', title: 'SubAccount', icon: '#tools',},
    ]
  },
  {
    header: "Balance", menus: [
      {href: '/balancesubaccount', title: 'SubAccount Balance', icon: '#tools',},
    ]
  },
  {
    header: "Journal", menus: [
      {href: '/journal', title: 'Journal', icon: '#tools',},
    ]
  },
  // {
  //   header: "Section", menus: [
  //     {href: '/order', title: 'Order', icon: '#speedometer2',},
  //     {href: '/payment', title: 'Payment', icon: '#tools',},
  //   ]
  // },
]

</script>
