<template>
  <div class="tables">
    <b-container class="container-xl" style="padding-top: 10px;">
      <b-row align-v="start">
        <b-col class="col-1">
          <b-icon icon="question-circle" v-b-modal.tbl-help-modal class="text-info align-middle" v-b-popover.hover.top="'Help'" style="height: 18px;"></b-icon>
          <b-modal id="tbl-help-modal" size="xl"  ok-only ok-title="Close" title="Help">
            <div id="p">
              <h3>Table Lookups</h3>
              <p id="p">This tool allows performing lookups against the EOSIO state for each contract. It will pre-load the list of all FIO tables, and the ABI definitions in advance.</p>
              <h4>Tips on using the Table Search tool</h4>
              <p id="p">By default, the searches assume an integer index that is incrementing. This allows incrementing through the table page by page. If you keep getting the same page when trying to go to the next page, it's likely that the table is using a 'name' type for the index. Under the covers a name is actually an integer, but it usually a very large integer. Reverse sort does work, but you cannot page through a table when it's enabled.</p>
              <h4>Advanced options</h4>
              <p id="p">In FIO there are only 3 index types in use (though EOSIO allows many others.) In the advanced tab, you can select the type of index, which index is being used, and can perform a hash of the value in the case of i128 indexes. This is how lookups for text are performed, and these are always an index number higher than one.</p>
              <p id="p">The "By Scope" option is for tables with secondary indexes, also called a scope. If you query a table and it comes back empty, it's likely because it's scoped. A good example is the 'eosio.msig' contracts 'proposals' table. Click on the "By Scope" checkbox, and it will allow querying the scopes in use on that table. Then the "scope" field from the results can be used by un-clicking "By Scope" and entering that account name into the scope field. </p>
              <h4>Output</h4>
              <p id="p">The output dropdown allows viewing various data related to the query. By default it shows the results, but it can also show the input sent to the API server, the contract's ABI, and can generate functioning source code that will recreate the query in several programming languages.</p>
            </div>
          </b-modal>
        </b-col>
        <b-col class="col-2">
          <b-select class="custom-select-sm border-dark" :options="contracts" v-on:change="selectedContract()" v-model="form.contract">
          </b-select>
        </b-col>
        <b-col class="col-2">
          <b-select class="custom-select-sm border-dark" :options="tables" v-model="form.table">
          </b-select>
        </b-col>
        <b-col class="col-1">
          <div>
            <b-button v-on:click="prevOffset()" class="btn-sm btn-dark" :disabled="form.offset === 0 || form.showAdvanced">
              <b-icon icon="arrow-left"></b-icon>
              prev
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
        <b-col class="col-1">
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
        <b-col v-if="form.showAdvanced" class="col-1">
          <b-check v-model="form.scopeQuery" v-on:change="clickedScoped()" class="border-dark"></b-check>
          <b-form-text id="scope-help">By Scope</b-form-text>
        </b-col>
       <b-col v-if="form.showAdvanced" class="col-2">
         <b-input :disabled="form.scopeQuery" size="sm" class="border-dark text-sm-center" v-model="form.scope"></b-input>
         <b-form-text id="scope-help">Scope</b-form-text>
       </b-col>
       <b-col class="col-1" v-if="form.showAdvanced">
         <b-input type="number" :disabled="form.scopeQuery" v-on:keyup.enter="executeQuery()" value="1" size="sm" class="border-dark text-sm-center" v-model="form.index"></b-input>
         <b-form-text id="index-help">Index</b-form-text>
       </b-col>
       <b-col v-if="form.showAdvanced" class="col-2">
         <b-select :disabled="form.scopeQuery" v-model="form.type" class="custom-select-sm border-dark">
           <b-select-option value="name">name</b-select-option>
           <b-select-option value="i64">i64</b-select-option>
           <b-select-option value="i128">i128</b-select-option>
         </b-select>
         <b-form-text id="type-help">Type</b-form-text>
       </b-col>
       <b-col v-if="form.showAdvanced" class="col-4">
         <td>
           <b-input v-on:keyup="copyToUpper()" class="border-dark text-sm-center" size="sm" v-model="form.lower"></b-input>
           <b-form-text id="lower-help">Lower bound</b-form-text>
         </td>
         <td>→</td>
         <td>
           <b-input :disabled="form.transform === 'hash'" class="border-dark text-sm-center" size="sm" v-model="form.upper"></b-input>
           <b-form-text id="upper-help">Upper Bound</b-form-text>
         </td>
       </b-col>
       <b-col clas="col-4" v-if="form.showAdvanced">
         <b-select :disabled="form.scopeQuery" class="custom-select-sm border-dark" v-on:change="setTypeFromTransform()"  v-model="form.transform">
           <b-select-option value="as-is">as-is</b-select-option>
           <b-select-option value="hash">string -> hash (i128)</b-select-option>
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
            <b-select-option disabled value="">Info</b-select-option>
            <b-select-option value="json">Query request (JSON)</b-select-option>
            <b-select-option value="abi">Contract ABI</b-select-option>
            <b-select-option disabled value="">Scaffolding</b-select-option>
            <b-select-option :disabled="form.scopeQuery" value="shell">Bash script</b-select-option>
            <b-select-option :disabled="form.scopeQuery" value="browser">Browser (native)</b-select-option>
            <b-select-option :disabled="form.scopeQuery" value="fio-go">Go (fio-go)</b-select-option>
            <b-select-option :disabled="form.scopeQuery" value="go">Go (stdlib)</b-select-option>
            <b-select-option :disabled="form.scopeQuery" value="node">NodeJS (axios/crypto)</b-select-option>
            <b-select-option :disabled="form.scopeQuery" value="python">Python3 (requests/hashlib)</b-select-option>
          </b-select>
          <b-form-text id="output-help">Output</b-form-text>
        </b-col>
        <b-col class="col-2">
          <b-check v-on:change="advancedChanged()" v-model="form.showAdvanced" class="border-dark"></b-check>
          <b-form-text id="rows-help">Advanced Options</b-form-text>
        </b-col>
        <b-col class="col-2">
          <b-check :disabled="form.scopeQuery" v-model="form.reverse" class="border-dark"></b-check>
          <b-form-text id="reverse-help">Reverse Sort Order</b-form-text>
        </b-col>
        <b-col class="col-2">
          <b-button class="btn-dark btn-sm" :disabled="disableQueryButton" v-on:click="executeQuery()">{{ queryButton }}</b-button>
        </b-col>
      </b-row>

      <b-row v-if="queryOutput !== ''">
        <b-col class="col-sm-12 text-sm-left" style="padding: 10px;">
          <b-button v-clipboard:copy="result" v-clipboard:success="copyResult" class="btn-light btn-outline-secondary btn-sm">
            <b-icon v-if="copiedNotify === ''" icon="clipboard" scale="1.3"></b-icon>
            <b-icon v-if="copiedNotify !== ''" class="text-success font-weight-bolder" icon="clipboard-check" scale="1.3"></b-icon>
          </b-button>
          <span class="text-sm-left text-secondary">&nbsp; {{ copiedNotify }}</span>
        </b-col>
      </b-row>
      <b-row v-if="queryOutput !== ''" align-v="stretch" class="bg-transparent" style="padding: 15px;">
        <div class="text-monospace font-weight-lighter text-left" id="resultOut" ref="resultOut"  style="padding: 15px;">
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
import * as Scaffold from './generator.js'
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
        scopeQuery: false,
      },
      contracts: [{value: "none", text: "Contract:", disabled: true}],
      tables: [{value: "none", text: "Table:", disabled: true}],
      abis: {},
      result: '',
      queryOutput: "",
      copiedNotify: "",
    }
  },

  methods: {
    executeQuery: async function () {
      let self = this

      this.queryOutput = " "

      if (!window.isSecureContext && this.form.transform === "hash") {
        this.queryOutput = "Sorry: cannot use SubtleCrypto features in an insecure context. Don't blame me, look at the w3c!"
      }

      let languages = Prism.languages.javascript
      let name = "javascript"
      let path = "/v1/chain/get_table_rows"
      let response
      switch (this.form.output) {
        case "go":
          languages = Prism.languages.go
          name = "go"
          if (this.form.transform === "hash") {
            this.result = Scaffold.stdGoHashedQuery(this.form, this.endpoint)
            break
          }
          this.result = Scaffold.stdGoNormalQuery(this.form, this.endpoint)
          break
        case "fio-go":
          languages = Prism.languages.go
          name = "go"
          if (this.form.transform === "hash") {
            this.result = Scaffold.fioGoHashedQuery(this.form, this.endpoint)
            break
          }
          this.result = Scaffold.fioGoNormalQuery(this.form, this.endpoint)
          break
        case "shell":
          languages = Prism.languages.bash
          name = "bash"
          if (this.form.transform === "hash") {
            this.result = Scaffold.bashHashQuery(this.form, this.endpoint)
            break
          }
          this.result = await Scaffold.bashNormalQuery(this.form, this.endpoint)
          break
        case "python":
          languages = Prism.languages.python
          name = "python"
            if (this.form.transform === "hash") {
              this.result = Scaffold.pythonHashQuery(this.form, this.endpoint)
              break
            }
            this.result = Scaffold.pythonNormalQuery(this.form, this.endpoint)
          break
        case "node":
          languages = Prism.languages.javascript
          name = "javascript"
          if (this.form.transform === "hash") {
            this.result = Scaffold.nodeHashQuery(this.form, this.endpoint)
            break
          }
          this.result = Scaffold.nodeNormalQuery(this.form, this.endpoint)
          break
        case "browser":
          if (this.form.transform === "hash") {
            this.result = Scaffold.browserHashQuery(this.form, this.endpoint)
            break
          }
          this.result = Scaffold.browserNormalQuery(this.form, this.endpoint)
          break
        case "json":
          this.result = await Scaffold.rawQuery(this.form)
          break
        case "abi":
          this.result = JSON.stringify(this.abis[this.form.contract], null, 4)
          break
        case "result":
          if (this.form.scopeQuery === true) {
            path = "/v1/chain/get_table_by_scope"
          }
          response = await fetch(self.endpoint + path, {
            method: 'POST',
            mode: 'cors',
            cache: 'no-cache',
            credentials: 'same-origin',
            redirect: 'error',
            referrerPolicy: 'no-referrer',
            body: await Scaffold.rawQuery(self.form)
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

    copyResult: async function () {
      this.copiedNotify = "copied!"
      let self = this
      setTimeout(function (){
        self.copiedNotify = ""
      }, 1000)
    },

    copyToUpper: function () {
      if (this.form.scopeQuery === true) {
        return
      }
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
      if (this.form.scopeQuery === true) {
        this.form.scopeQuery = false
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

    clickedScoped: function () {
      this.form.lower = ""
      this.form.upper = ""
      if (this.form.scopeQuery === true) {
        this.form.output = "result"
        this.form.lower = "............1"
        this.form.upper = "zzzzzzzzzzzz"
      }
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
      })
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
        })
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

  watch: {
    getNetwork: function () {
      this.getContracts()
    },
    endpoint: function () {
      this.getContracts()
    }
  },

  computed: {
    ...mapGetters({
      endpoint: 'getEndpoint',
      getNetwork: 'getNetwork',
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
