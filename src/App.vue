<template>
  <div id="app" style="border-bottom: 0;">
    <div style="padding-bottom: 15px; padding-top: 0;" class="bg-dark text-light container-fluid">
      <b-container fluid="true">
      <b-row align-h="center" style="padding-top: 10px;">
        <b-col class="col-3">
          <b-dropdown variant="dark" text="FIO Tools">
            <b-dropdown-item-button v-on:click="navNames()">Name / Domain Lookup</b-dropdown-item-button>
            <!-- b-dropdown-item-button v-on:click="navProxies()">Voting: Proxies</b-dropdown-item-button -->
            <b-dropdown-item-button v-on:click="navVote()">Voting: BP Rankings</b-dropdown-item-button>
          </b-dropdown>
          <b-dropdown variant="dark" text="BP Tools">
            <b-dropdown-item-button v-on:click="navFees()">Fee Votes</b-dropdown-item-button>
          </b-dropdown>
          <b-dropdown variant="dark" text="Dev Tools">
            <b-dropdown-item-button variant="dark" v-on:click="navTables()">Search Tables</b-dropdown-item-button>
          </b-dropdown>
        </b-col>
        <b-col class="col2"> &nbsp; </b-col>
        <b-col class="col-2"><div style="padding-top: 10px;" class="text-center bold align-text-bottom text-secondary"> {{ getPageTitle }}</div></b-col>
        <b-col class="col-3 text-light">
          <td>
          <div style="padding-top: 6px; padding-right: 10px;">
            <b-select v-model="network" class="custom-select-sm border-secondary bg-dark text-white-50">
              <b-select-option value="mainnet">Main Net</b-select-option>
              <b-select-option value="testnet">Test Net</b-select-option>
            </b-select>
          </div>
          </td>
          <td>
            <div v-if="network === 'mainnet'">
            <b-select v-model="currentNode" v-on:change="updateCurrent" :options="mainNodes" class="custom-select-sm border-secondary bg-dark text-white-50">
            </b-select>
            </div>
          </td>
          <td>
            <div v-if="network === 'testnet'">
              <b-select v-model="currentNode" v-on:change="updateCurrent" :options="testNodes" style="border-color: grey;" class="custom-select-sm border-black bg-dark text-white-50">
              </b-select>
            </div>
          </td>
        </b-col>
        <b-col class="col-2">
          <td>
            <a href="https://greymass.com/en/anchor/" target="_blank"><img src="./assets/anchor.svg" v-b-popover.hover.bottomleft="'Get Anchor Wallet'" alt="get anchor wallet" style="padding-right: 5px; padding-left: 40px; height: 32px;"> &nbsp;</a>
          </td>
          <td>
            <b-button v-on:click="anchorLogin()" :disabled="getAnchorConnected" class="btn-sm bg-transparent text-white btn-outline-dark">
              <b-icon v-if="!getAnchorConnected"  v-b-popover.hover.bottomright="'Connect to Anchor wallet'" class="text-secondary" icon="toggle2-off" font-scale="1.5"></b-icon>
              <b-icon v-if="getAnchorConnected" v-b-popover.hover.bottomright="'Anchor is connected'" class="text-success" icon="toggle-on" font-scale="1.5"></b-icon>
            </b-button>
          </td>
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
      currentNode: "https://fio.blockpane.com",
      mainNodes: [
        {value: "https://fio.blockpane.com", text: "Server: Block Pane"},
        {value: "https://fio.greymass.com", text: "Server: Greymass"},
        {value: "https://fio.eosusa.news", text: "Server: EOSUSA"},
        {value: "https://api.fio.alohaeos.com", text: "Server: Aloha EOS"},
        {value: "https://fio.zenblocks.io", text: "Server: Zen Blocks"},
        {value: "https://api.fio.currencyhub.io", text: "Server: Currency Hub"},
        {value: "https://api.fio.eosdetroit.io", text: "Server: EOS Detroit"},
        {value: "https://fio.eosphere.io", text: "Server: EOSphere"},
        {value: "https://fio.acherontrading.com", text: "Server: Acheron"},
        {value: "https://fio.eu.eosamsterdam.net", text: "Server: EOS Amsterdam"},
        {value: "https://api.fio.services", text: "Server: fio.services"},
        {value: "https://api.fio.greeneosio.com", text: "Server: Green EOSIO"},
        {value: "https://fio.cryptolions.io", text: "Server: Cryptolions"},
        {value: "https://api.fiosweden.org", text: "Server: EOS Sweden"},
        {value: "https://fio-bp.dmail.co", text: "Server: Dmail"},
        {value: "https://fio.genereos.io", text: "Server: Genereos"},
        {value: "https://fio.eosrio.io", text: "Server: EOSRIO"},
        {value: "https://fio.eoscannon.io", text: "Server: EOS Cannon"},
        {value: "https://fio.eosdublin.io", text: "Server: EOS Dublin"},
        {value: "https://fio.eosdac.io", text: "Server: EOS DAC"},
        {value: "https://fio.eos.barcelona", text: "Server: EOS Barcelona"},
        {value: "https://fio.eosargentina.io", text: "Server: EOS Argentina"},
        {value: "https://fioapi.ledgerwise.io", text: "Server: Ledgerwise"},
        {value: "https://fio-mainnet.eosblocksmith.io", text: "Server: EOS Blocksmith"},
        {value: "https://fio-za.eostribe.io", text: "Server: EOS Tribe"},
      ],
      testNodes: [
        {value: "https://fiotestnet.blockpane.com", text: "Server: Block Pane"},
        {value: "https://fiotestnet.greymass.com", text: "Server: Greymass"},
        {value: "https://testnet.fio.eosdetroit.io", text: "Server: EOS Detroit"},
        {value: "https://test.fio.eosusa.news", text: "Server: EOSUSA"},
        {value: "https://fio-testnet.genereos.io", text: "Server: GenerEOS"},
        {value: "https://api.fiotest.alohaeos.com", text: "Server: Aloha EOS"},
        {value: "https://api.fiotest.currencyhub.io", text: "Server: Currency Hub"},
        {value: "https://testnet.fioprotocol.io", text: "Server: Cryptolions"},
        {value: "https://api.testnet.fiosweden.org", text: "Server: EOS Sweden"},
        {value: "https://fio-test.eos.barcelona", text: "Server: EOS Barcelona"},
        {value: "https://fiotestnet.ledgerwise.io", text: "Server: Ledgerwise"},
        {value: "https://fio-testnet.eosblocksmith.io", text: "Server: EOS Blocksmith"},
        {value: "https://fio-bp.dmail.co:7777", text: "Server: dmail labs"},
      ],
    }
  },

  methods: {
    ...mapActions([
      'anchorLogin',
    ]),
    ...mapMutations([
        'setNetwork',
        'setPageTitle',
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

    navNames: function () {
      this.setPageTitle("Search")
      this.$router.push('names')
    },

    navFees: function () {
      this.setPageTitle("Producer Tools")
      this.$router.push('fees')
    },

    navTables: function () {
      this.setPageTitle("Developer Tools")
      this.$router.push('tables')
    },

    navProxies: function () {
      this.setPageTitle("Voting")
      this.$router.push('proxies')
    },

    navVote: function () {
      this.setPageTitle("Voting")
      this.$router.push('vote')
    },

    updateCurrent: function (){
      let self = this
      console.log(typeof self.currentNode)
      this.setNetwork({net: self.network, endpoint: self.currentNode})
      self.currentNode = self.getEndpoint
      //console.log(this.currentNode)
    },



  },

  computed: {
    ...mapGetters([
        'getAnchorConnected',
        'getEndpoint',
        'getPageTitle',
    ])
  },

  watch: {
    network: function () {
      if (this.network === "mainnet") {
        this.currentNode = this.mainNodes[0].value
      } else {
        this.currentNode = this.testNodes[0].value
      }
      this.updateCurrent()
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
