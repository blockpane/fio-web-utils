<template>
  <b-container ref="pageContainer" fluid="true" style="padding-top: 10px; padding-left: 10px; padding-right: 10px">
    <div v-if="!loading" class="text-center">
      <b-check v-model="advanced">Advanced View</b-check>
    </div>
    <div class="text-center text-info" v-if="loading">
      <b-icon animation="throb" font-scale="4" icon="three-dots"></b-icon>
      <br />Loading ...
    </div>
    <div v-if="advanced && !loading" class="container-fluid">
      <BootstrapTable
          id="table" class="table-hover table-sm table-borderless"
          data-toggle="table" data-flat="true" :data="tableData"
          :columns="columns" data-card-view="false"
      ></BootstrapTable>
    </div>
    <div v-if="!advanced && !loading" class="container-xl">
      <b-container class="container-sm">
        <b-row>
          <b-col class="col-2>"><div class="text-right"> &nbsp; </div></b-col>
          <b-col class="col-2>"><div class="text-right"> Select Top:</div></b-col>
          <b-col class="col-1 align-middle>">
            <b-select v-on:change="selectVotes()" v-model="voteCount" class="custom-select-sm" value="21">
              <b-select-option value=15>15</b-select-option>
              <b-select-option value=21>21</b-select-option>
              <b-select-option value=30>30</b-select-option>
            </b-select>
          </b-col>
          <b-col class="col-2 align-middle">
            <b-select v-on:change="selectVotes()" v-model="voteBy" class="custom-select-sm" value="score">
              <b-select-option value="score">By Score</b-select-option>
              <b-select-option value="votes">By Votes</b-select-option>
            </b-select>
          </b-col>
          <b-col class="col-2>"><div class="text-right"> &nbsp; </div></b-col>
        </b-row>
      </b-container>
      <BootstrapTable
          id="simpleTable" ref="simpleTable" class="table-hover table-sm table-light"
          data-toggle="table" data-flat="true" :data="tableData"
          :columns="summaryColumns" data-card-view="false"
          height="200" data-virtual-scroll="true"
      ></BootstrapTable>
    </div>
  </b-container>
</template>

<script>
import {mapGetters} from "vuex";

export default {
  name: 'Vote',

  data () {
    return {
      data: [],
      counter: 0,
      advanced: false,
      loading: true,
      voteBy: "score",
      voteCount: 21,
      columns: [
        {title: 'FIO Name', align: 'right', field: 'address', formatter: this.linkFormatter},
        {title: '', align: 'center', field: 'svg', formatter: this.svgFormatter},
        {title: '', align: 'center', field: 'missing_excluded', formatter: this.boolInvertedFormatter},
        {title: 'Actual', align: 'center', field: 'extra.rank', sortable: true, formatter: this.rankFormatter},
        {title: 'Delta', align: 'center', field: 'rank_delta', sortable: true, formatter: this.deltaFormatter},
        {title: 'Rank', align: 'center', field: 'rank', sortable: true, formatter: this.rankFormatter},
        {title: 'Score', align: 'center', field: 'score', sortable: true, formatter: this.scoreFormatter},
        {title: 'Fees', align: 'center', style: "border-left-style: solid;", field: 'fee_vote_30d', sortable: true, formatter: this.zeroFormatter},
        {title: 'MSIG', align: 'center', field: 'msig_30d', sortable: true, formatter: this.zeroFormatter},
        {title: 'Claims', align: 'center', field: 'bpclaim_1d', sortable: true, formatter: this.zeroFormatter},
        {title: 'TPID Claims', align: 'center', field: 'tpidclaim_1d', sortable: true, formatter: this.zeroFormatter},
        {title: 'Compute', align: 'center', field: 'compute', sortable: true, formatter: this.zeroFormatter},
        {title: 'Burn', align: 'center', field: 'burnexpired_1d', sortable: true, formatter: this.zeroFormatter},
        {title: 'CPU', align: 'center', field: 'cpu_score', sortable: true, formatter: this.zeroFormatter},
        {title: 'Keys', align: 'center', field: 'diff_sign_key', formatter: this.boolFormatter},
        {title: 'BP Json', align: 'center', field: 'bp_json', formatter: this.boolFormatter},
        {title: 'Json CORS', align: 'center', field: 'bp_json_cors', formatter: this.boolFormatter},
        {title: 'URL', align: 'center', field: 'valid_url', formatter: this.boolFormatter},
        {title: 'Perms', align: 'center', field: 'using_linked_auth_or_msig', formatter: this.boolFormatter},
        {title: '30d Claims', align: 'center', field: 'has_claimed_30d', formatter: this.boolFormatter},
      ],
      summaryColumns: [
        {title: 'FIO Name', align: 'right', sortable: true, field: 'address', formatter: this.linkFormatter},
        {title: "Vote", checkbox: "true", field: "voteFor"},
        {title: '', align: 'center', field: 'svg', formatter: this.svgFormatter},
        {title: '', align: 'center', field: 'rank', formatter: this.thumbsUp},
        {title: '', align: 'center', field: 'score', visible: false},
        {title: '', align: 'center', field: 'missing_excluded', formatter: this.boolInvertedFormatter},
        {title: 'mm ᵮ Voted', align: 'center', field: 'extra.total_votes', sortable: true, formatter: this.mmFormatter},
        {title: 'By Votes', align: 'center', field: 'extra.rank', sortable: true, formatter: this.rankFormatter},
        {title: 'Proxy Score', align: 'center', field: 'rank', sortable: true, formatter: this.rankFormatter},
        {title: 'Last Seen (UTC)', style: "text-monospace", align: 'right', field: 'last_seen', sortable: true},
      ],
    }
  },

  computed: {
    ...mapGetters([
      'getAnchorConnected',
      'getAnchorIdentity',
      'getAnchorTransport',
      'getAnchorLink',
      'getResultMessage',
      'getNetwork',
      'getChainId',
      'getEndpoint',
    ]),
    tableData: function () {
      return this.data
    },
  },

  watch: {
    getNetwork: function () {
      this.fetchData()
    },
  },

  beforeMount() {
    this.fetchData()
  },


  methods: {

    selectVotes: function () {
      for (let i = 0; i < this.tableData.length; i++) {
        this.tableData[i].voteFor = false
        if (this.voteBy === "score" && this.tableData[i].rank <= this.voteCount) {
          this.tableData[i].voteFor = true
        } else if (this.voteBy === "votes" && this.tableData[i].extra.rank <= this.voteCount) {
          this.tableData[i].voteFor = true
        } else {
          this.tableData[i].voteFor = false
        }
        console.log(this.tableData[i].voteFor)
      }
      this.$refs.simpleTable.$forceUpdate()
    },

    fetchData: async function() {
      this.loading = true
      const prods = await this.getProducers()
      let src = 'https://snap.blockpane.com/ranks.json'
      if (this.getNetwork === "testnet") {
        src = 'https://snap.blockpane.com/testnet-ranks.json'
      }
      const response = await fetch(src, {
        method: 'GET',
        mode: 'cors',
        cache: 'no-cache',
        //credentials: 'same-origin',
        redirect: 'error',
        referrerPolicy: 'no-referrer',
      })
      this.data = await response.json()
      for (let i = 0; i < this.tableData.length; i++) {
        this.data[i].rank = i + 1
        this.data[i].extra = prods[this.data[i].account]
        this.data[i].rank_delta = (i + 1) - this.data[i].extra.rank
        this.data[i].voteFor = (i + 1) <= 21
        this.data[i].last_seen = await this.lastActive(this.data[i].account)
      }
      this.data.sort((a, b) => (a.extra.rank > b.extra.rank) ? 1 : -1)
      this.loading = false
    },

    lastActive: async function (account) {
      let self = this
      const response = await fetch(self.getEndpoint + "/v1/history/get_actions", {
        method: 'POST',
        mode: 'cors',
        cache: 'no-cache',
        //credentials: 'same-origin',
        redirect: 'error',
        referrerPolicy: 'no-referrer',
        body: JSON.stringify(
            {
              "account_name": account,
              "pos": -1,
              "offset": -1
            }
        )
      });
      const body = await response.json()
      if (body.actions === null || body.actions === undefined) {
        return "unknown"
      }
      return body.actions[0].block_time
    },

    scoreFormatter: function (value, row) {
      if (row.missing_excluded === true) {
        return `<div style="color: #000000;">${value}</div>`
      }
      if (value > 100) {
        return `<div style="color: #09641f;">${value}</div>`
      } else if (value > 50) {
        return `<div style="color: #0c8551;">${value}</div>`
      } else if (value > 25) {
        return `<div style="color: #2656ca;">${value}</div>`
      } else if (value > 0) {
        return `<div style="color: #555555;">${ value }</div>`
      }
      return `<div style="color: darkred;">${ value }<div>`
    },

    rankFormatter: function (value, row) {
      if (row.missing_excluded === true) {
        return `<div style="color: #000000;">${value}</div>`
      }
      if (value <= 15) {
        return `<div style="color: #0a8100;">${value}</div>`
      } else if (value <= 21) {
        return `<div style="color: #003f91;">${value}</div>`
      }
      return `<div style="color: #580000;">${value}<div>`
    },

    linkFormatter: function (value, row) {
      return `<div class="text-lg-right">` +
          `<a target="bloks" href="https://fio.bloks.io/account/${ row.account }">${ value }</a></div>`
    },

    mmFormatter: function (value) {
      return (parseFloat(value)/ 1000000).toFixed(2)
    },

    deltaFormatter: function (value, row) {
      let color = "black"
      let extra = " &nbsp; &nbsp;"
      if (parseInt(value) > 0) {
        color = "#580000"
        if (row.extra.rank <= 21 && row.extra.rank + row.rank_delta > 21) {
          extra = "↓"
        }
      } else if (parseInt(value) < 0) {
        color = "#09641f"
        if (row.extra.rank > 21 && row.extra.rank + row.rank_delta <= 21) {
          extra = "↑"
        }
      }
      return `<div style="color: ${ color };">${ value } ${ extra }</div>`
    },

    thumbsUp: function (value, row) {
      if (row.score < 0) {
        return '<span style="color: #AA3333;">⚠️</span>'
      } else if (parseInt(value) <= 21) {
        return '<span style="color: #33AA33;">✓</span>'
      }
        return ""
    },

    zeroFormatter: function (value) {
      if (value === 0) {
        return ''
      }
      return value
    },

    boolInvertedFormatter: function (value) {
      if (value) {
        return '<div style="color: red">🚫</div>'
      }
      return ""
    },

    svgFormatter: function (value) {
      if (value.match('^https://[a-zA-Z0-9_/.-]+.svg$')) {
        return `<img height="28" width="28" src="${ value }"/>`
      }
      return ''
    },

    boolFormatter: function (value, row) {
      if (row.missing_excluded === true) {
        return '<div style="color: orange">-</div>'
      }
      let color = 'red'
      let symbol = '𐄂';
      if (value) {
        color = 'green'
        symbol = '✓'
      }
      return '<div style="color: ' + color + '">' +
          symbol +
          '</div>'
    },

    getProducers: async function () {
      const location = {
        10: "East Asia",
        20: "Australia",
        30: "West Asia",
        40: "Africa",
        50: "Europe",
        60: "East North America",
        70: "South America",
        80: "West North America",
      }
      const response = await fetch(this.getEndpoint + '/v1/chain/get_producers', {
        method: 'POST',
        mode: 'cors',
        cache: 'no-cache',
        credentials: 'same-origin',
        redirect: 'error',
        referrerPolicy: 'no-referrer',
      })
      let producers = {}
      const body = await response.json()
      if (body.producers !== null && body.producers !==  undefined) {
        for (let i = 0; i < body.producers.length; i++) {
          producers[ body.producers[i].owner] = {
            rank: i + 1,
            total_votes: parseFloat( body.producers[i].total_votes.toString().split(".")[0]) / 1000000000.0,
            url:  body.producers[i].url,
            last_claim_time:  body.producers[i].last_claim_time,
            location: location[ body.producers[i].location]
          }
        }
      }
      return producers
    },

  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
