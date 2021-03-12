<template>
  <div id="app" style="border-bottom: 0;">
    <div style="padding-bottom: 15px; padding-top: 0;" class="bg-dark text-light">
      <b-row style="padding-top: 10px;">
        <b-col class="col-sm-4">&nbsp;</b-col>
        <b-col class="col-sm-4">
        <b-button class="btn-dark" v-on:click="enableNames()">Names / Domains</b-button>
        <b-button class="btn-dark" v-on:click="enableVotes()">Fee Votes</b-button>
        <img src="./assets/anchor.svg" alt="connect anchor wallet" style="padding-left: 40px; height: 32px;"> &nbsp;
        <b-button v-on:click="anchorLogin()" :disabled="anchorConnected" class="btn-md bg-transparent text-white">
          <b-icon v-if="!anchorConnected"  v-b-popover.hover.top="'Connect to Anchor wallet'" class="text-secondary" icon="toggle2-off" font-scale="1.5"></b-icon>
          <b-icon v-if="anchorConnected" v-b-popover.hover.top="'Anchor is connected'" class="text-success" icon="toggle-on" font-scale="1.2"></b-icon>
        </b-button>
        </b-col>
        <b-col class="col-sm-1 text-light">
          <div style="padding-top: 6px;">
        <b-select v-model="network" v-on:change="setNetwork(network)" style="border-color: grey;" class="custom-select-sm bg-dark text-white-50">
          <b-select-option value="mainnet">mainnet</b-select-option>
          <b-select-option value="testnet">testnet</b-select-option>
        </b-select>
          </div>
        </b-col>
        <b-col class="col-sm-3">&nbsp;</b-col>
      </b-row>
    </div>

    <div v-if="showVote" style="border-bottom: 14px;">
      <FeeVote/>
    </div>
    <div v-if="showNames" style="border-bottom: 14px;">
      <Names/>
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
        </b-link> ≁
      </div>
    </div>

  </div>
</template>

<script>
import {mapActions, mapGetters, mapMutations} from "vuex";
import FeeVote from './components/Fees.vue'
import Names from './components/Names.vue'

export default {
  name: 'App',
  components: {
    FeeVote,
    Names,
  },

  data () {
    return {
      showVote: false,
      showNames: true,
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
    },

    enableVotes: function () {
      this.hideAll()
      this.showVote = true
    },

    enableNames: function () {
      this.hideAll()
      this.showNames = true
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
