<template>
  <div>
    <SpTheme>
      <SpNavbar :links="navbarLinks" :outerlinks="outerLinks" :active-route="router.currentRoute.value.path" />
      <router-view />
    </SpTheme>
  </div>
</template>

<script lang="ts">
import { SpNavbar, SpTheme } from './starportvue'
import { computed, onBeforeMount } from 'vue'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'

export default {
  components: { SpTheme, SpNavbar },

  setup() {
    // store
    let $s = useStore()

    // router
    let router = useRouter()

    // state
    let navbarLinks = [
      { name: 'Wallet', url: '/' },
    ]

    // outer linkes
    let outerLinks = [
      { name: 'Explorer', url: 'https://explorer.bluwallet.app' },
      { name: 'Faucet', url: 'https://faucet.bluwallet.app' },
    ]

    // computed
    let address = computed(() => $s.getters['common/wallet/address'])

    // lh
    onBeforeMount(async () => {
      await $s.dispatch('common/env/init')

      // There is only wallet page for now, ignore router push
      // router.push('portfolio')
    })

    return {
      navbarLinks,
      outerLinks,
      // router
      router,
      // computed
      address
    }
  }
}
</script>

<style scoped lang="scss">
body {
  margin: 0;
}
</style>
