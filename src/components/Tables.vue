<template>
  <div class="tables">
    <b-container class="container-xl" style="padding-top: 10px;">
      <b-row align-v="start">
        <b-col class="col-2">
          <b-select class="custom-select-sm border-dark" :options="contracts" v-on:change="selectedContract()" v-model="form.contract">
          </b-select>
        </b-col>
        <b-col class="col-2">
          <b-select class="custom-select-sm border-dark" :options="tables" v-model="form.table">
          </b-select>
        </b-col>
        <b-col class="col-2">
          <div>
            <b-button v-on:click="prevOffset()" class="btn-sm btn-dark" :disabled="form.offset === 0 || form.showAdvanced">
              <b-icon icon="arrow-left"></b-icon>
              previous
            </b-button>
          </div>
        </b-col>
        <b-col class="col-2">
          <b-form-input class="border-dark text-sm-center" id="rows" size="sm" type="number" v-model="form.numRows"></b-form-input>
          <b-form-text id="rows-help">Number of Rows</b-form-text>
        </b-col>
        <b-col class="col-2">
          <b-form-input class="border-dark text-sm-center" id="offset" size="sm" type="number" v-model="form.offset" :disabled="form.showAdvanced"></b-form-input>
          <b-form-text id="offset-help">Offset</b-form-text>
        </b-col>
        <b-col class="col-2">
          <b-button class="btn-sm btn-dark" :disabled="form.showAdvanced" v-on:click="nextOffset()" >
            next
            <b-icon icon="arrow-right"></b-icon>
          </b-button>
        </b-col>
        <b-col class="col-2">

        </b-col>

        <b-col class="col-2"></b-col>
      </b-row>

      <b-row><b-col>&nbsp;<br />&nbsp;</b-col></b-row>

      <b-row align-v="center">
       <b-col v-if="form.showAdvanced" class="col-2">
         <b-input size="sm" class="border-dark text-sm-center" v-model="form.scope"></b-input>
         <b-form-text id="scope-help">Scope</b-form-text>
       </b-col>
       <b-col class="col-1" v-if="form.showAdvanced">
         <b-input type="number" v-on:keyup.enter="executeQuery()" value="1" size="sm" class="border-dark text-sm-center" v-model="form.index"></b-input>
         <b-form-text id="index-help">Index</b-form-text>
       </b-col>
       <b-col v-if="form.showAdvanced" class="col-2">
         <b-select v-model="form.type" class="custom-select-sm border-dark">
           <b-select-option value="name">name</b-select-option>
           <b-select-option value="i64">i64</b-select-option>
           <b-select-option value="i128">i128</b-select-option>
         </b-select>
         <b-form-text id="type-help">Type</b-form-text>
       </b-col>
       <b-col v-if="form.showAdvanced" class="col-3">
         <td>
           <b-input v-on:keyup="copyToUpper()" class="border-dark text-sm-center" size="sm" v-model="form.lower"></b-input>
           <b-form-text id="lower-help">Lower bound</b-form-text>
         </td>
         <td>→</td>
         <td>
           <b-input class="border-dark text-sm-center" size="sm" v-model="form.upper"></b-input>
           <b-form-text id="upper-help">Upper Bound</b-form-text>
         </td>
       </b-col>
       <b-col clas="col-5" v-if="form.showAdvanced">
         <b-select class="custom-select-sm border-dark" v-on:change="setTypeFromTransform()"  v-model="form.transform">
           <b-select-option value="as-is">as-is</b-select-option>
           <b-select-option value="hash">string -> hash (128)</b-select-option>
         </b-select>
         <b-form-text id="transform-help">Transform</b-form-text>
       </b-col>
      </b-row>
      <b-row>
        <b-col class="col-12">
          <div>&nbsp;<br />&nbsp;</div>
        </b-col>
      </b-row>

      <b-row align-h="center">

        <b-col class="col-3">
          <b-select id="showResults" class="custom-select-sm border-dark" v-model="form.output">
            <b-select-option value="result">Query Result</b-select-option>
            <b-select-option disabled value="note">Info</b-select-option>
            <b-select-option value="json">Query request (JSON)</b-select-option>
            <b-select-option value="abi">Contract ABI</b-select-option>
            <b-select-option disabled value="note">Scaffolding</b-select-option>
            <b-select-option value="shell">shell: bash script</b-select-option>
            <b-select-option value="fio-go">code: Go/fio-go</b-select-option>
            <b-select-option value="go">code: Go/stdlib</b-select-option>
            <b-select-option value="python">code: Python3</b-select-option>
            <!-- todo: more code generators ...
            <b-select-option value="axios">code: nodejs/axios</b-select-option>
            <b-select-option value="fetch">code: browser ES6/fetch</b-select-option>
            -->
          </b-select>
          <b-form-text id="output-help">Output</b-form-text>
        </b-col>
        <b-col class="col-2">
          <b-check v-on:change="advancedChanged()" v-model="form.showAdvanced" class="border-dark"></b-check>
          <b-form-text id="rows-help">Advanced Options</b-form-text>
        </b-col>
        <b-col class="col-2">
          <b-check v-model="form.reverse" class="border-dark"></b-check>
          <b-form-text id="reverse-help">Reverse Sort Order</b-form-text>
        </b-col>
        <b-col class="col-2">
          <b-button class="btn-dark btn-sm" :disabled="disableQueryButton" v-on:click="executeQuery()">{{ queryButton }}</b-button>
        </b-col>
      </b-row>

      <b-row v-if="queryOutput !== ''"><b-col class="col-12 text-sm-center">&nbsp;<br />&nbsp;</b-col></b-row>
      <b-row v-if="queryOutput !== ''" align-v="stretch" class="bg-transparent" style="padding: 15px;">
        <div class="text-monospace font-weight-lighter text-left" id="resultOut"  style="padding: 15px;">
          <pre class="text-black" v-html="queryOutput"></pre>
       </div>
      </b-row>
    </b-container>
    <div>
    </div>
  </div>
</template>

<script>
import Prism from 'prismjs'
import { rawQuery, fioGoHashedQuery, fioGoNormalQuery, stdGoHashedQuery, stdGoNormalQuery, bashHashQuery, bashNormalQuery, pythonHashQuery, pythonNormalQuery} from './generator.js'
import {mapGetters} from "vuex";

export default {
  name: "Tables",
  data() {
    return {
      form: {
        contract: "none",
        table: "none",
        numRows: 25,
        offset: 0,
        scope: "",
        index: 1,
        type: "name",
        lower: "",
        upper: "",
        transform: "as-is",
        output: "result",
        showAdvanced: false,
        reverse: false,
        more: false,
      },
      contracts: [{value: "none", text: "Contract:", disabled: true}],
      tables: [{value: "none", text: "Table:", disabled: true}],
      abis: {},
      result: '',
      queryOutput: "",
    }
  },

  methods: {
    executeQuery: async function () {
      let self = this

      this.queryOutput = " "

      if (!window.isSecureContext) {
        this.queryOutput = "Sorry: cannot use 'crypto.subtle' library in an insecure context. Don't blame me, look at the w3c!"
      }

      let languages = Prism.languages.javascript
      let name = "javascript"
      let response
      switch (this.form.output) {
        case "go":
          languages = Prism.languages.go
          name = "go"
          if (this.form.transform === "hash") {
            this.result = stdGoHashedQuery(this.form, this.endpoint)
            break
          }
          this.result = stdGoNormalQuery(this.form, this.endpoint)
          break
        case "fio-go":
          languages = Prism.languages.go
          name = "go"
          if (this.form.transform === "hash") {
            this.result = fioGoHashedQuery(this.form, this.endpoint)
            break
          }
          this.result = fioGoNormalQuery(this.form, this.endpoint)
          break
        case "shell":
          languages = Prism.languages.bash
          name = "bash"
          if (this.form.transform === "hash") {
            this.result = bashHashQuery(this.form, this.endpoint)
            break
          }
          this.result = await bashNormalQuery(this.form, this.endpoint)
          break
        case "python":
          languages = Prism.languages.python
          name = "python"
            if (this.form.transform === "hash") {
              this.result = pythonHashQuery(this.form, this.endpoint)
              break
            }
            this.result = pythonNormalQuery(this.form, this.endpoint)
          break
        case "json":
          this.result = await rawQuery(this.form)
          break
        case "abi":
          this.result = JSON.stringify(this.abis[this.form.contract], null, 4)
          break
        case "result":
          response = await fetch(self.endpoint + "/v1/chain/get_table_rows", {
            method: 'POST',
            mode: 'cors',
            cache: 'no-cache',
            credentials: 'same-origin',
            redirect: 'error',
            referrerPolicy: 'no-referrer',
            body: await rawQuery(self.form)
          });
          try {
            await response.json().then((body) => {
                self.result = JSON.stringify(body, null, 4)
            })
          } catch (e) {
            self.queryOutput = e
          }

      }
      this.queryOutput = Prism.highlight(this.result, languages, name)
    },

    copyToUpper: function () {
      this.form.upper = this.form.lower
      console.log(this.form.lower)
    },

    advancedChanged: function () {
      if (this.form.showAdvanced === true) {
        this.form.index = 1
        this.form.type = "name"
        this.form.upper = ""
        this.form.lower = ""
        this.form.transform = "as-is"
        this.form.scope = this.form.contract
        return
      }
      this.form.index = 1
      this.form.type = "i64"
      this.form.transform = "as-is"
      this.form.scope = this.form.contract
    },

    setTypeFromTransform: function () {
      if (this.form.transform === 'i64') {
        this.form.type = 'i64'
      } else if (this.form.transform === 'hash') {
        this.form.type = 'i128'
      }
    },

    prevOffset: function () {
      this.form.offset = parseInt(this.form.offset) - parseInt(this.form.numRows)
      if (this.form.offset < 0) {
        this.form.offset = 0
      }
      this.executeQuery()
    },

    nextOffset: function () {
      this.form.offset = parseInt(this.form.numRows) + parseInt(this.form.offset)
      this.executeQuery()
    },

    selectedContract: function () {
      this.tables = [{value: "none", text: "Table:", disabled: true}]
      this.form.table = "none"
      this.form.scope = this.form.contract
      this.queryOutput = " "
      let self = this
      for (const table of self.abis[self.form.contract]["abi"]["tables"]) {
        self.tables.push({
          value: table.name,
          text: table.name,
          disabled: false,
        })
      }
    },

    getContracts: async function () {
      let self = this
      const response = await fetch(self.endpoint + "/v1/chain/get_table_rows", {
        method: 'POST',
        mode: 'cors',
        cache: 'no-cache',
        credentials: 'same-origin',
        redirect: 'error',
        referrerPolicy: 'no-referrer',
        body: JSON.stringify(
            {
              "code": "eosio",
              "scope": "eosio",
              "table": "abihash",
              "lower_bound": "0",
              "limit": 100,
              "json": true
            }
        )
      });
      self.queryOutput = "getting contracts ..."
      try {
        await response.json().then((body) => {
          self.contracts = [
            {value: "none", text: "Contract:", disabled: true}
          ]
          for (const row of body.rows) {
            self.contracts.push({
              value: row.owner,
              text: row.owner,
              disabled: false
            })
            self.queryOutput = "done."
          }
        })
      } catch (e) {
        self.queryOutput = e
      }
      await this.getAbis()
    },

    getAbis: async function () {
      let self = this
      self.queryOutput = "getting ABIs ..."
      for (const contract of self.contracts) {
        if (contract.value === "none") continue
        const response = await fetch(self.endpoint + "/v1/chain/get_abi", {
          method: 'POST',
          mode: 'cors',
          cache: 'no-cache',
          credentials: 'same-origin',
          redirect: 'error',
          referrerPolicy: 'no-referrer',
          body: JSON.stringify({"account_name": contract.value})
        });
        try {
          await response.json().then((body) => {
              self.abis[contract.value] = body
              self.queryOutput = "loaded abi for " + contract.value
          })
        } catch (e) {
          self.queryOutput = e
          throw e
        }
      }
      this.queryOutput = "ready for queries"
    },

  },

  mounted: function () {
    this.getContracts()
  },

  computed: {
    ...mapGetters({
      endpoint: 'getEndpoint',
    }),

    disableQueryButton: function () {
      // todo: this will need to change when adding code generators
      return !(this.form.contract !== "none" && (this.form.output === "abi" || this.form.table !== "none"))
    },

    queryButton: function () {
      if (this.form.output === "result") {
        return "Execute Query"
      } else if (this.form.output === "json") {
        return "Show Query Input"
      } else if (this.form.output === "abi") {
        return "Fetch ABI"
      }
      return "Generate Code"
    },

  },

}
</script>
