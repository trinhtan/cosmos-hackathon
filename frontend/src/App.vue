<template>
  <div id="app">
    <header-page />
    <router-view />
    <footer-page />
  </div>
</template>

<script>
import HeaderPage from '@/components/Header';
import FooterPage from '@/components/Footer';
import { mapState, mapActions } from 'vuex';
export default {
  name: 'app',
  components: {
    HeaderPage,
    FooterPage
  },
  computed: { ...mapState('ethereum', ['web3']) },
  methods: {
    ...mapActions('ethereum', ['setWeb3']),
    ...mapActions('cosmos', ['setCosmosAccount'])
  },
  async mounted() {
    // await this.setWeb3();
    await this.setCosmosAccount();
  }
};
</script>

<style>
body {
  margin: 0;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Oxygen', 'Ubuntu',
    'Cantarell', 'Fira Sans', 'Droid Sans', 'Helvetica Neue', sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

.App-logo {
  height: 40vmin;
  pointer-events: none;
}

@media (prefers-reduced-motion: no-preference) {
  .App-logo {
    animation: App-logo-spin infinite 20s linear;
  }
}

.App-header {
  background-color: #282c34;
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  font-size: calc(10px + 2vmin);
  color: white;
}

.App-link {
  color: #61dafb;
}

.text-align-left {
  text-align: left !important;
}

.text-align-right {
  text-align: right !important;
}
.text-align-center {
  text-align: center !important;
}
.text-align-justify {
  text-align: justify !important;
}

.cursor-pointer {
  cursor: pointer;
}

.width-100 {
  width: 100%;
}

@keyframes App-logo-spin {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}

section {
  width: 100vw;
  height: auto;
  display: flex;
  justify-content: center;
}

.box-shadow {
  padding: 20px;
  border-radius: 10px;
  box-shadow: 0 10px 16px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
}
</style>
