<template>
  <div class="tables">
    <b-container class="container-xl" style="padding-top: 10px;">
      <b-row align-v="start">
        <b-col class="col-2">
          <b-select class="custom-select-sm border-dark" v-model="form.contract">
            <b-select-option value="none">Contract:</b-select-option>
          </b-select>
        </b-col>
        <b-col class="col-2">
          <b-select class="custom-select-sm border-dark" v-model="form.table">
            <b-select-option value="none">Table:</b-select-option>
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
          <b-button class="btn-sm btn-dark" v-on:click="nextOffset()" >
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
       <b-col v-if="form.showAdvanced">
         <b-input type="number" value="1" size="sm" class="border-dark text-sm-center" v-model="form.index"></b-input>
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
       <b-col v-if="form.showAdvanced" class="col-4">
         <b-container class="container-fluid">
           <b-row>
             <b-col>
               <b-input class="border-dark text-sm-center" v-on:keyup="copyToUpper()" size="sm" v-model="form.lower"></b-input>
               <b-form-text id="lower-help">Lower bound</b-form-text>
             </b-col>
             <b-col>
               <b-input class="border-dark text-sm-center" size="sm" v-model="form.upper"></b-input>
               <b-form-text id="upper-help">Upper Bound</b-form-text>
             </b-col>
           </b-row>
         </b-container>
       </b-col>
       <b-col clas="col-3" v-on="setTypeFromTransform()" v-if="form.showAdvanced">
         <b-select class="custom-select-sm border-dark" v-model="form.transform">
           <b-select-option value="as-is">as-is</b-select-option>
           <b-select-option value="i64">name -> i64</b-select-option>
           <b-select-option value="hash">string -> hash (128)</b-select-option>
         </b-select>
         <b-form-text id="transform-help">Transform</b-form-text>
       </b-col>
       <b-col>
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
            <b-select-option value="json">Query request (JSON)</b-select-option>
            <b-select-option value="abi">Contract ABI</b-select-option>
            <!-- todo: code generators
            <b-select-option value="shell">shell: bash script</b-select-option>
            <b-select-option value="axios">code: nodejs/axios</b-select-option>
            <b-select-option value="fetch">code: browser ES6/fetch</b-select-option>
            <b-select-option value="eos-go">code: Go/eos-go</b-select-option>
            <b-select-option value="fio-go">code: Go/fio-go</b-select-option>
            <b-select-option value="python">code: Python3</b-select-option>
            -->
          </b-select>
          <b-form-text id="output-help">Output</b-form-text>
        </b-col>
        <b-col class="col-2">
          <b-check v-model="form.showAdvanced" class="border-dark"></b-check>
          <b-form-text id="rows-help">Advanced Options</b-form-text>
        </b-col>
        <b-col class="col-2">
          <b-check v-model="form.reverse" class="border-dark"></b-check>
          <b-form-text id="reverse-help">Reverse Sort Order</b-form-text>
        </b-col>
        <b-col class="col-2">
          <b-button class="btn-dark btn-sm" v-on:click="executeQuery()">Execute Query</b-button>
        </b-col>
      </b-row>

      <b-row v-if="queryOutput !== ''"><b-col class="col-12 text-sm-center">Result:</b-col></b-row>
      <b-row v-if="queryOutput !== ''" align-v="stretch" class="rounded" style="background:#111; padding: 15px;">
        <div class="text-monospace font-weight-lighter text-left" id="resultOut"  style="padding: 15px;">
          <pre class="text-white" v-html="queryOutput"></pre>
       </div>
      </b-row>
    </b-container>
    <div>
    </div>
  </div>
</template>

<script>
import Prism from 'prismjs'
import { rawQuery } from './generator.js'

export default {
  name: "Tables",
  data() {
    return {
      form: {
        contract: "none",
        table: "none",
        numRows: 10,
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
      result: '',
      queryOutput: "",
    }
  },

  methods: {
    executeQuery: function () {
      this.result = `
// this is a test!
{
  "testing": "formatting"
}
`
      let languages, name
      switch (this.form.output) {
        case "eos-go":
        case "fio-go":
          languages = Prism.languages.go
          name = "go"
          break
        case "shell":
          languages = Prism.languages.bash
          name = "bash"
          break
        case "python":
          languages = Prism.languages.python
          name = "python"
          break
        case "json":
          this.result = rawQuery(this.form)
          languages = Prism.languages.javascript
          name = "javascript"
          break
        default:
          languages = Prism.languages.javascript
          name = "javascript"
      }
      this.queryOutput = Prism.highlight(this.result, languages, name)
    },

    copyToUpper: function () {
      this.form.upper = this.form.lower
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

  },

  mounted: function () {
  },

  computed: {},

}
</script>
