<template>
  <div id="app" style="border-bottom: 0;">
    <div style="padding-bottom: 15px; padding-top: 0;" class="bg-dark text-light container-fluid">
      <b-container fluid>
      <b-row align-h="center" style="padding-top: 10px;">
        <b-col class="col-4">
          <b-button class="btn-dark" v-on:click="$router.push('names')">Names / Domains</b-button>
          <b-button class="btn-dark" v-on:click="$router.push('fees')">Fee Votes</b-button>
          <b-button class="btn-dark" v-on:click="$router.push('tables')">Search tables</b-button>
        </b-col>
        <b-col class="col-2">
          <td>
                <img src="./assets/anchor.svg" alt="connect anchor wallet" style="padding-right: 5px; padding-left: 40px; height: 32px;"> &nbsp;
          </td>
          <td>
                <b-button v-on:click="anchorLogin()" :disabled="anchorConnected" class="btn-sm bg-transparent text-white btn-outline-dark">
                  <b-icon v-if="!anchorConnected"  v-b-popover.hover.top="'Connect to Anchor wallet'" class="text-secondary" icon="toggle2-off" font-scale="1.5"></b-icon>
                  <b-icon v-if="anchorConnected" v-b-popover.hover.top="'Anchor is connected'" class="text-success" icon="toggle-on" font-scale="1.5"></b-icon>
                </b-button>
          </td>
        </b-col>
        <b-col class="col-2 text-light">
          <div style="padding-top: 6px;">
          <b-select v-model="network" v-on:change="setNetwork(network)" style="border-color: grey;" class="custom-select-sm bg-dark text-white-50">
            <b-select-option value="mainnet">mainnet</b-select-option>
            <b-select-option value="testnet">testnet</b-select-option>
          </b-select>
          </div>
        </b-col>
      </b-row>
      </b-container>
    </div>

    <div style="border-bottom: 14px;">
      <router-view class="view"></router-view>
    </div>
    <div style="border-top: 14px;"><hr></div>

    <b-modal id="about-modal" size="lg"  ok-only ok-title="Close" title="About">
      todo
    </b-modal>
    <div class="fixed-bottom">
      <div class="bg-dark text-white-50" id="p" style="border-bottom: 0; border-top: 8px; font-size: medium;">
        ≁ <b-link href="https://blockpane.com" class="light-link">&copy; 2021<img alt="Block Pane, LLC." src="./assets/bp-small.png" height="18"></b-link> &nbsp;
        ≁ <b-link v-b-modal.about-modal class="light-link">About</b-link> ≁ <a class="light-link" href="https://blockpane.com/privacy.html">Privacy Policy</a> ≁ &nbsp;
        <b-link href="https://github.com/blockpane/fio-web-utils">
          <img src="./assets/GitHub-Mark-Light-32px.png" alt="source on Github" height="18">
        </b-link>
      </div>
    </div>

  </div>
</template>

<script>
import {mapActions, mapGetters, mapMutations} from "vuex";

export default {
  name: 'App',
  components: {
  },

  data () {
    return {
      showNames: true,
      showVote: false,
      showTables: false,
      network: "mainnet",
    }
  },

  methods: {
    ...mapActions([
      'anchorLogin',
    ]),
    ...mapMutations([
        'setNetwork'
    ]),

    hideAll: function () {
      this.showVote = false
      this.showNames = false
      this.showTables = false
    },

    enableVotes: function () {
      this.hideAll()
      this.showVote = true
    },

    enableNames: function () {
      this.hideAll()
      this.showNames = true
    },

    enableTables: function () {
      this.hideAll()
      this.showTables = true
    },

  },

  computed: {
    ...mapGetters({
        anchorConnected: 'getAnchorConnected',
    })
  },

  watch: {
    network: function () {
      this.setNetwork(this.network)
    },
  }

}
</script>

<style>
#app {
  margin-top: 0;
  text-align: center;
}
#p {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  font-size: small;
}

a.light-link:link {
  color: #BABCD0;
  font-family: Avenir, Helvetica, Arial, sans-serif;
}
a.light-link:visited {
  color:#BABCD0;
  font-family: Avenir, Helvetica, Arial, sans-serif;
}
a.light-link:hover {
  color:#BABCD0;
  font-family: Avenir, Helvetica, Arial, sans-serif;
}

</style>
<!--
  color: #2c3e50;
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
-->
