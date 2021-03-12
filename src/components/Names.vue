<template>
  <div class="names" style="text-align: center; padding-top: 10px;">
    <div class="container-lg" >
      <img alt="FIO" src="../assets/fio-dark.svg" height="64" ><br /><br /><span class="text-xl-center">Domain / Address Search</span>
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
    <p id="sans">Where to get addresses/domains:</p>
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
import axios from 'axios';
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
      const encoded = new TextEncoder().encode(this.input)
      crypto.subtle.digest("SHA-1", encoded).then( hash => {
        const hashArray = Array.from(new Uint8Array(hash)).slice(0,16).reverse()
        const hashHex = '0x' + hashArray.map(b => b.toString(16).padStart(2, '0')).join('')
        axios.post(this.getEndpoint + "/v1/chain/get_table_rows", {
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
        }).then(res=> {
          if (res.data.rows != null && res.data.rows[0] != null) {
            this.addressResult.name = res.data.rows[0].name
            this.addressResult.owner_account = res.data.rows[0].owner_account
            this.addressResult.bloks = 'https://fio.bloks.io/account/' + res.data.rows[0].owner_account
            const t = new Date(res.data.rows[0].expiration * 1000);
            this.addressResult.expiration = t.toISOString()
            this.addressResult.mappings = ""
            res.data.rows[0].addresses.forEach(r=> {
              this.addressResult.mappings += `${r.token_code}/${r.chain_code}: ${r.public_address}` + "\n"
              console.log(this.addressResult.mappings)
            })
            this.addressResult.bundleeligiblecountdown = res.data.rows[0].bundleeligiblecountdown
            this.showAddressTable = true
            console.log(this.addressResult)
          }
        }).catch(error=> {
          console.log(error)
          this.showAddressTable = false
        })
      })
    },

    getDomain() {
      const encoded = new TextEncoder().encode(this.input)
      crypto.subtle.digest("SHA-1", encoded).then( hash => {
        const hashArray = Array.from(new Uint8Array(hash)).slice(0,16).reverse()
        const hashHex = '0x' + hashArray.map(b => b.toString(16).padStart(2, '0')).join('')
        axios.post(this.getEndpoint + "/v1/chain/get_table_rows", {
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
        }).then(res=> {
          if (res.data.rows != null && res.data.rows[0] != null) {
            this.domainResult.name = res.data.rows[0].name
            this.domainResult.account = res.data.rows[0].account
            this.domainResult.bloks = 'https://fio.bloks.io/account/' + res.data.rows[0].account
            this.domainResult.is_public = "private domain"
            if (res.data.rows[0].is_public === 1) {
              this.domainResult.is_public = "domain allows public registrations"
            }
            const t = new Date(res.data.rows[0].expiration * 1000);
            this.domainResult.expiration = t.toISOString()
            this.showDomainTable = true
            console.log(this.domainResult)
          }
        }).catch(error=> {
          console.log(error)
          this.showDomainTable = false
        })
      })
    },

    available() {
      console.log("checking")
      if (this.domainOrAddress === "domain" || this.domainOrAddress === "address") {
        console.log("searching for a " + this.domainOrAddress)
        axios.post(this.getEndpoint + "/v1/chain/avail_check",
            {
              fio_name: this.input
            }
        ).then(res => {
          this.gotResponse = true;
          this.status = res.data;
          if (res.status === 400) {
            this.querySuccess = false
            this.validResult = false
            this.resultShow = true
            this.alreadyRegistered = "error"
            if (res.data.fields[0].error != null) {
              this.alreadyRegistered = res.data.fields[0].error
            }
          } else if (res.data.is_registered === 0) {
            this.querySuccess = true
            this.validResult = true
            this.resultShow = true
            this.alreadyRegistered = this.domainOrAddress + " is not registered"
          } else if (res.data.is_registered === 1) {
            this.querySuccess = false
            this.validResult = true
            this.resultShow = true
            this.alreadyRegistered = this.domainOrAddress + " is already registered"
          } else {
            this.querySuccess = false
            this.validResult = true
            this.resultShow = false
            this.alreadyRegistered = ""
          }
          if (this.validResult && this.domainOrAddress === "domain") {this.getDomain()}
          if (this.validResult && this.domainOrAddress === "address") {this.getAddress()}
        }).catch(
            error => {
              console.log(error)
              this.querySuccess = false
              this.validResult = false
              this.resultShow = true
              this.alreadyRegistered = "error"
              if (error.response.data.fields[0].error != null) {
                this.alreadyRegistered = error.response.data.fields[0].error
              }
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
