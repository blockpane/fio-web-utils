<template>
  <div class="names" style="text-align: center; padding-top: 10px;">
    <div class="container-lg" >
      <img alt="FIO" src="../assets/fio-dark.svg" height="64" ><br /><br />
      <!-- <span class="text-xl-center">Domain / Address Search</span> -->
    </div>

    <div class="container-lg">
      <br />
    </div>

    <div class="container-lg">
      <b-input-group size="md" class="mb-2">
        <b-input-group-prepend is-text>
          <b-icon icon="search"></b-icon>
        </b-input-group-prepend>
        <b-form-input v-model="input" v-on:keyup.enter="available" type="search" placeholder="domain / name@domain"></b-form-input>
        <b-button v-bind:class="{ 'alert-success': querySuccess }" v-show="resultShow">
          &nbsp; {{ alreadyRegistered }}&nbsp;
        </b-button>
      </b-input-group>
      <b-button @click="available" size="sm" class="mb-2">
        <b-icon icon="search" aria-hidden="true"></b-icon> Lookup {{ domainOrAddress }}</b-button>
    </div>
    <div class="container-lg">
      <br />
      <br />
    </div>

    <div class="container-lg" v-show="showDomainTable"><div class="d-flex justify-content-center">
      <br />
      <table class="table table-bordered table-hover table-borderless">
        <thead class="font-weight-bolder lead">
        <tr>
          <th>Name</th>
          <th>Account</th>
          <th>Public</th>
          <th>Expires</th>
        </tr>
        </thead>
        <tbody class="font-weight-bold">
        <tr>
          <td class="p-3">{{ domainResult.name }}</td>
          <td class="p-3"><b-link :href="domainResult.bloks">{{ domainResult.account }}</b-link></td>
          <td class="p-3">{{ domainResult.is_public }}</td>
          <td class="p-3">{{ domainResult.expiration }}</td>
        </tr>
        </tbody>
      </table>
    </div></div>

    <div class="container-lg" v-show="showAddressTable"><div class="d-flex justify-content-center">
      <table class="table table-bordered table-hover table-borderless">
        <thead class="font-weight-bolder lead">
        <tr>
          <th>Name</th>
          <th>Account</th>
          <th>Bundled</th>
          <th>Expires</th>
          <th>Address Mappings</th>
        </tr>
        </thead>
        <tbody class="font-weight-bold">
        <tr>
          <td class="p-3">{{ addressResult.name }}</td>
          <td class="p-3"><b-link :href="addressResult.bloks">{{ addressResult.owner_account }}</b-link></td>
          <td class="p-3">{{ addressResult.bundleeligiblecountdown }}</td>
          <td class="p-3">{{ addressResult.expiration }}</td>
          <td class="p-3 text-xl-right"><pre>{{ addressResult.mappings }}</pre></td>
        </tr>
        </tbody>
      </table>
    </div></div>

    <div class="container-lg">
      <br />
      <br />
    </div>
    <p id="sans">Where else to get addresses/domains:</p>
    <div class="container-lg">
      <div class="row">
        <div class="col-lg">
          <b-link href="https://fioprotocol.io/free-fio-addresses/" target="_blank" rel="noopener"><span class="font-weight-bold">Get a FREE address</span><br />Offered by multiple wallets</b-link>
        </div>
        <div class="col-lg">
          <b-link href="https://greymass.github.io/fio-register/" target="_blank" rel="noopener"><span class="font-weight-bold">Greymass FIO Registration Helper</span><br />Purchase using Anchor/Scatter wallets</b-link>
        </div>
        <div class="col-lg">
          <b-link href="https://reg.fioprotocol.io/" target="_blank" rel="noopener"><span class="font-weight-bold">FIO foundation registration site</span><br />Purchase using BTC/ETH/USDC</b-link>
        </div>
      </div>
    </div>
    <div class="container-lg">
      <br />
      <br />
    </div>

  </div>
</template>

<script>
import {mapGetters} from "vuex";
export default {
  name: 'Names',
  data () {
    return {
      input: "",
      alreadyRegistered: "",
      querySuccess: false,
      validResult: true,
      resultShow: true,
      showAddressTable: false,
      showDomainTable: false,
      addressResult: {
        name: "",
        expiration: "",
        owner_account: "",
        blocks: "",
        mappings: "",
        bundleeligiblecountdown: 0,
      },
      domainResult: {
        name: "",
        account: "",
        bloks: "",
        is_public: "",
        expiration: "",
      },
      showRenewDomain: false,
      showRenewAddress: false,
      showRegDomain: false,
      showRegAddress: false,
    }
  },

  methods: {
    clear () {
      this.alreadyRegistered = ""
      this.resultShow = false
      this.showDomainTable = false
      this.showAddressTable = false
    },

    getAddress() {
      let self = this
      const encoded = new TextEncoder().encode(self.input)
      crypto.subtle.digest("SHA-1", encoded).then( hash => {
        const hashArray = Array.from(new Uint8Array(hash)).slice(0,16).reverse()
        const hashHex = '0x' + hashArray.map(b => b.toString(16).padStart(2, '0')).join('')
        fetch(self.getEndpoint + "/v1/chain/get_table_rows", {
          method: 'POST',
          mode: 'cors',
          cache: 'no-cache',
          credentials: 'same-origin',
          redirect: 'error',
          referrerPolicy: 'no-referrer',
          body: JSON.stringify(
              {
                code: "fio.address",
                scope: "fio.address",
                table: "fionames",
                lower_bound: hashHex,
                upper_bound: hashHex,
                limit: 1,
                key_type: "i128",
                index_position: "5",
                json: true,
                reverse: true
              })
        }).then(body => {
          body.json().then(res => {
            if (res.rows != null && res.rows[0] != null) {
              self.addressResult.name = res.rows[0].name
              self.addressResult.owner_account = res.rows[0].owner_account
              self.addressResult.bloks = 'https://fio.bloks.io/account/' + res.rows[0].owner_account
              const t = new Date(res.rows[0].expiration * 1000);
              self.addressResult.expiration = t.toISOString()
              self.addressResult.mappings = ""
              res.rows[0].addresses.forEach(r => {
                self.addressResult.mappings += `${r.token_code}/${r.chain_code}: ${r.public_address}` + "\n"
              })
              self.addressResult.bundleeligiblecountdown = res.rows[0].bundleeligiblecountdown
              self.showAddressTable = true
            }
          })
        }).catch(error=> {
          console.log(error)
          self.showAddressTable = false
        })
      })
    },

    getDomain() {
      let self = this
      self.showRenewDomain = false
      self.showRenewAddress = false
      self.showRegDomain = false
      self.showRegAddress = false
      const encoded = new TextEncoder().encode(self.input)
      crypto.subtle.digest("SHA-1", encoded).then( hash => {
        const hashArray = Array.from(new Uint8Array(hash)).slice(0,16).reverse()
        const hashHex = '0x' + hashArray.map(b => b.toString(16).padStart(2, '0')).join('')
        fetch(self.getEndpoint + "/v1/chain/get_table_rows", {
          method: 'POST',
          mode: 'cors',
          cache: 'no-cache',
          credentials: 'same-origin',
          redirect: 'error',
          referrerPolicy: 'no-referrer',
          body: JSON.stringify({
              code: "fio.address",
              scope: "fio.address",
              table: "domains",
              lower_bound: hashHex,
              upper_bound: hashHex,
              limit: 1,
              key_type: "i128",
              index_position: "4",
              json: true,
              reverse: true
          })
        }).then(body => {
          body.json().then(res => {
            if (res.rows != null && res.rows[0] != null) {
              self.domainResult.name = res.rows[0].name
              self.domainResult.account = res.rows[0].account
              self.domainResult.bloks = 'https://fio.bloks.io/account/' + res.rows[0].account
              self.domainResult.is_public = "private domain"
              if (res.rows[0].is_public === 1) {
                self.domainResult.is_public = "domain allows public registrations"
              }
              const t = new Date(res.rows[0].expiration * 1000);
              self.domainResult.expiration = t.toISOString()
              console.log("now: " + (parseInt(Date.now().toFixed(0)) / 1000))
              console.log("     " + parseInt(res.rows.[0].expiration))
              if ((parseInt(Date.now().toFixed(0)) / 1000) > parseInt(res.rows.[0].expiration) + 90 * 24 * 60 * 60) {
                self.domainResult.expiration += "is expired, and available to register!"
              } else if ((parseInt(Date.now().toFixed(0)) / 1000) > parseInt(res.rows.[0].expiration)) {
                self.domainResult.expiration += " is pending expiration – within 90 day grace period."
              }
              self.showDomainTable = true
            }
          })
        }).catch(error=> {
          console.log(error)
          self.showDomainTable = false
        })
      })
    },

    available() {
      let self = this
      if (self.domainOrAddress === "domain" || self.domainOrAddress === "address") {
        fetch(self.getEndpoint + "/v1/chain/avail_check",  {
              method: 'POST',
              mode: 'cors',
              cache: 'no-cache',
              credentials: 'same-origin',
              redirect: 'error',
              referrerPolicy: 'no-referrer',
              body: JSON.stringify({
                  fio_name: self.input
                })
            }
        ).then(body => {
          body.json().then(res => {
            self.gotResponse = true;
            self.status = body.status;
            if (body.status === 400) {
              self.querySuccess = false
              self.validResult = false
              self.resultShow = true
              self.alreadyRegistered = "error " + res.fields[0].error
            } else if (res.is_registered === 0) {
              self.querySuccess = true
              self.validResult = true
              self.resultShow = true
              self.alreadyRegistered = self.domainOrAddress + " is not registered"
            } else if (res.is_registered === 1) {
              self.querySuccess = false
              self.validResult = true
              self.resultShow = true
              self.alreadyRegistered = self.domainOrAddress + " is already registered"
            } else {
              self.querySuccess = false
              self.validResult = true
              self.resultShow = false
              self.alreadyRegistered = ""
            }
            if (self.validResult && self.domainOrAddress === "domain") {
              self.getDomain()
            }
            if (self.validResult && self.domainOrAddress === "address") {
              self.getAddress()
            }
          })
        }).catch(
            error => {
              console.log(error)
              self.querySuccess = false
              self.validResult = false
              self.resultShow = true
              self.alreadyRegistered = error.toString()
            }
        )}
    }
  },

  computed: {
    ...mapGetters([
      'getEndpoint',
    ]),

    domainOrAddress() {
      this.clear()
      if (this.input.match("@")) return "address";
      if (this.input.length > 0) return "domain";
      return ""
    }
  }
}
</script>

<style>
#names {
  text-align: center;
  margin-top: 60px;
}
#sans {
  font-family: Avenir, Helvetica, Arial, sans-serif;
}
</style>
