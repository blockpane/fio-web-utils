<template>

  <div class="feeVote" style="text-align: right;">
    <div style="padding-bottom: 15px; padding-top: 20px;">
      <div class="h3 text-center">
        FIO Fees
        <b-icon icon="question-circle" v-b-modal.help-modal class="text-info align-middle" v-b-popover.hover.top="'How do block-producers calculate fees?'" style="height: 18px;"></b-icon>
        <b-modal id="help-modal" size="xl"  ok-only ok-title="Close" title="About Fees">
          <div class="h4">Fee Ratios and Multipliers</div>
          <div>
            <p id="p">Fees can only be voted on by block producers</p>
            <em>What it does:</em><br />
              These calls adjust fees on the FIO chain.
              <br /><br />
              If the FIO token price fluctuates it can result in prices that don’t reflect the utility of purchasing an address. If the token price increases, and the number of tokens required to purchase a domain does not decrease it will create a disincentive for users to actively use the protocol. These mechanisms allow for adjusting the real cost of using the protocol. Most producers are targeting around a two-dollar equivalent for registering a FIO name.
              <br /><br />
              <em>There are three phases in the process of updating the fees:</em>
              <ol>
            <li>
              The producers vote for the “base” fees (aka fee ratio) using `fio.fee::setfeevote`. The amount charged for each action is up to the producer, but my personal suggestion is that the base fees from genesis (or in most cases the suggested fee when a new action is added) should be used.
            </li>
            <li>
              The second part is setting the multiplier. This is a little more complex, and in a sense requires the producer to act as a price oracle of sorts. This multiplier is applied to the fee, and is actually how a producers fee vote is determined. This is the most important step, and potentially perilous. I don’t want to make any specific recommendations on how a producer should come to this number, but I would suggest a few recommendations: don’t update too often (hourly is plenty, I only update every 4 hours,) ensure whatever price feed being consumed doesn’t give weird numbers (anything more than a few percent adjustment and my oracle ignores the price feed,) and finally don’t update for very small changes, it’s a waste of fees and requires more calls to computefees (see below,) and finally a little variability in the exact moment the fees are set can help prevent price manipulation via flash attacks and other attempts at manipulation.
            </li>
            <li>
              After the first two steps, the producer has only submitted their vote, but no updates have actually occurred until computefees is executed. This call, much like burnexpired, works in small chunks, only calculating a few of the fees at a time. So it’s likely that it needs to be called several times.
            </li>
              </ol>
            <b-link href="https://developers.fioprotocol.io/docs/bp/fee-setting">The developer documentation</b-link> does a good job of explaining how specifically to handle setting fees. Maintaining fees are one of the most important and complex jobs on the FIO Protocol and this job falls solely to the block producers. This tool exists to make that task a little easier.
          </div>
        </b-modal>
      </div>
    </div>

    <div class="container-xl" style="padding-top: 10px;">
      <b-row>
        <b-col>
          Coingecko price:
          <b-icon style="height: 18px;" icon="arrow-clockwise" v-on:click="price()"></b-icon>
          <div class="text-info text-xl-right font-weight-bolder" style="padding-top: 10px;"> ${{ gecko.market_data.current_price.usd }}</div>
        </b-col>
        <b-col class="text-left">
          Suggested Fee Ratio
          <b-input id="multi" v-model="multiplier" type="number" v-on:change="updateFeeTable()">
          </b-input>
        </b-col>
        <b-col class="text-sm-left">
          <div style="padding-top: 18px;">
            <b-check :disabled="!getAnchorConnected" v-model="onlyFees">Multiplier Only</b-check>
          </div>
        </b-col>
      </b-row>

      <b-row>
        <b-col style="text-align: center; padding-top: 14px;">
          <b-button :disabled="!getAnchorConnected" v-on:click="sendTx()">Apply {{ setFeesOnly }}</b-button>
        </b-col>
      </b-row>
      <b-row>
        <b-col>
          <div class="text-center text-info"><hr>{{ getResultMessage }}<hr></div>
        </b-col>
      </b-row>
    </div>
    <div class="container-lg">
      <b-table
          id="fee-table-transition"
          hover small striped borderless responsive="true" head-variant="dark"
          :items="feeTable" sort-by="suggested_ratio" v-bind:sort-desc="sortDescending"
          :tbody-transition-props="transProps"
          :fields="fields"
          primary-key="fee_name"
      ></b-table>
    </div>
  </div>
</template>

<script>
import { mapGetters, mapMutations, mapActions } from 'vuex'

export default {
  name: 'FeeVote',

  data () {
    return {
      network: "mainnet",
      mainnet: "21dcae42c0182200e93f954a074011f9048a7624c6fe81d3c9541a614a88bd1c",
      testnet: "b20901380af44ef59c5918439a1f9a41d83669020319a80574b804a5f95cbd7e",
      gecko: {market_data: {current_price: {usd: 0.0 }}},
      onlyFees: false,
      feeTable: [],
      myVotes: {},
      myMult: {},
      currentPrices: {},
      multiplier: 0,
      sortDescending: true,
      fees: [
        {end_point: "add_bundled_transactions", value: 200000000},
        {end_point: "add_pub_address", value: 30000000},
        {end_point: "add_to_whitelist", value: 30000000},
        {end_point: "auth_delete", value: 20000000},
        {end_point: "auth_link", value: 20000000},
        {end_point: "auth_update", value: 50000000},
        {end_point: "burn_fio_address", value: 60000000},
        {end_point: "cancel_funds_request", value: 60000000},
        {end_point: "msig_approve", value: 20000000},
        {end_point: "msig_cancel", value: 20000000},
        {end_point: "msig_exec", value: 20000000},
        {end_point: "msig_invalidate", value: 20000000},
        {end_point: "msig_propose", value: 50000000},
        {end_point: "msig_unapprove", value: 20000000},
        {end_point: "new_funds_request", value: 60000000},
        {end_point: "proxy_vote", value: 30000000},
        {end_point: "record_obt_data", value: 60000000},
        {end_point: "register_fio_address", value: 2000000000},
        {end_point: "register_fio_domain", value: 40000000000},
        {end_point: "register_producer", value: 10000000000},
        {end_point: "register_proxy", value: 1000000000},
        {end_point: "reject_funds_request", value: 30000000},
        {end_point: "remove_all_pub_addresses", value: 60000000},
        {end_point: "remove_from_whitelist", value: 30000000},
        {end_point: "remove_pub_address", value: 60000000},
        {end_point: "renew_fio_address", value: 2000000000},
        {end_point: "renew_fio_domain", value: 40000000000},
        {end_point: "set_fio_domain_public", value: 30000000},
        {end_point: "submit_bundled_transaction", value: 30000000},
        {end_point: "submit_fee_multiplier", value: 0},
        {end_point: "submit_fee_ratios", value: 70000000},
        {end_point: "transfer_fio_address", value: 60000000},
        {end_point: "transfer_fio_domain", value: 100000000},
        {end_point: "transfer_locked_tokens", value: 100000000},
        {end_point: "transfer_tokens_pub_key", value: 100000000},
        {end_point: "unregister_producer", value: 20000000},
        {end_point: "unregister_proxy", value: 20000000},
        {end_point: "vote_producer", value: 30000000}
      ],
      fields: [
        { key: 'fee_name', sortable: true },
        { key: 'actual_fee', sortable: false },
        { key: 'my_current_vote', sortable: false },
        { key: 'suggested_ratio', sortable: true },
        { key: 'multiplier_×_price', sortable: false },
        { key: 'resulting_cost', sortable: true }
      ],
      transProps: {
        name: 'flip-list'
      },
    }
  },

  mounted () {
    this.price()
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

    setFeesOnly: function () {
      if (this.onlyFees) {
        return "Multiplier"
      }
      return "Fee Ratios and Multiplier"
    },
  },

  watch: {
    getAnchorConnected: function () {
      this.updateFeeTable()
      if (this.getAnchorConnected === false) {
        this.initNetwork()
      }
    },
    getNetwork: function () {
      this.updateFeeTable()
    },
    getEndpoint: function () {
      this.updateFeeTable()
    }
  },

  methods: {
    ...mapMutations([
        'setConnected',
        'setResult',
        'setNetwork',
        'setTransport',
        'setAnchorLink',
        'setAnchorId',
    ]),

    ...mapActions([
        'anchorLogin',
    ]),

    updateFeeTable: async function () {
      await this.getMyVotes()
      await this.getMyMult()
      await this.getActualFees()
      let result = []
      for (const fee of this.fees) {
        let myVote = "-"
        if (this.getAnchorConnected) {
          myVote = "(" + (this.myVotes[fee.end_point] / 1000000000).toFixed(4)
          myVote += " @ " + this.myMult.toFixed(4) + ") = $"
          myVote += ((this.gecko.market_data.current_price.usd * (this.myVotes[fee.end_point] / 1000000000)) * this.myMult).toFixed(2)
        }
        result.push({
          fee_name: fee.end_point,
          actual_fee: "ᵮ " + (this.currentPrices[fee.end_point] / 1000000000).toFixed(4) + " × $"
              + this.gecko.market_data.current_price.usd.toFixed(2)+ " ~ $"
              + ((this.gecko.market_data.current_price.usd * (this.currentPrices[fee.end_point] / 1000000000))).toFixed(2),
          my_current_vote: myVote,
          suggested_ratio: (fee.value / 1000000000).toString(),
          "multiplier_×_price": "× " + this.multiplier + " × $" + this.gecko.market_data.current_price.usd.toPrecision(6) + " = ",
          resulting_cost: "$ " + ((this.gecko.market_data.current_price.usd * (fee.value / 1000000000)) * this.multiplier).toFixed(2),
        })
      }
      result.sort((a, b) => (a.suggested_ratio > b.suggested_ratio) ? -1 : 1)
      this.feeTable = result
    },

    price: async function () {
      const response = await fetch('https://api.coingecko.com/api/v3/coins/fio-protocol?tickers=false&market_data=true&community_data=false&developer_data=false&sparkline=false')
      const body = await response.json()
      this.gecko = body
      this.multiplier = (2 / body.market_data.current_price.usd / 2).toPrecision(4)
      await this.updateFeeTable()
    },

    initNetwork: function () {
      this.setNetwork({net: this.network, endpoint: this.getEndpoint})
      this.setTransport(null)
      this.resetAnchorLink()
      this.setAnchorId(null)
    },

    getActualFees: async function () {
      let self = this
      const response = await fetch(self.getEndpoint + "/v1/chain/get_table_rows", {
        method: 'POST',
        mode: 'cors',
        cache: 'no-cache',
        credentials: 'same-origin',
        redirect: 'error',
        referrerPolicy: 'no-referrer',
        body: JSON.stringify(
            {
              "code": "fio.fee",
              "scope": "fio.fee",
              "table": "fiofees",
              "json": true
            }
        )
      });
      return await response.json().then((body) => {
        self.currentPrices = {}
        for (const row of body.rows) {
          self.currentPrices[row.end_point] = row.suf_amount
        }
        return true
      })
    },

    getMyVotes: async function () {
      if (!this.getAnchorConnected) {
        return
      }
      let self = this
      const response = await fetch(self.getEndpoint + "/v1/chain/get_table_rows", {
        method: 'POST',
        mode: 'cors',
        cache: 'no-cache',
        credentials: 'same-origin',
        redirect: 'error',
        referrerPolicy: 'no-referrer',
        body: JSON.stringify(
            {
              "code": "fio.fee",
              "scope": "fio.fee",
              "table": "feevotes2",
              "lower_bound": self.getAnchorIdentity.auth.actor,
              "upper_bound": self.getAnchorIdentity.auth.actor,
              "limit": 1,
              "key_type": "name",
              "index_position": "2",
              "json": true,
              "reverse": false
            }
        )
      });
      return await response.json().then((body) => {
        self.myVotes = {}
        for (const row of body.rows[0].feevotes) {
          self.myVotes[row.end_point] = row.value
        }
      })
    },

    getMyMult: async function () {
      if (!this.getAnchorConnected) {
        return
      }
      let self = this
      const response = await fetch(self.getEndpoint + "/v1/chain/get_table_rows", {
        method: 'POST',
        mode: 'cors',
        cache: 'no-cache',
        credentials: 'same-origin',
        redirect: 'error',
        referrerPolicy: 'no-referrer',
        body: JSON.stringify(
            {
              "code": "fio.fee",
              "scope": "fio.fee",
              "table": "feevoters",
              "lower_bound": self.getAnchorIdentity.auth.actor,
              "upper_bound": self.getAnchorIdentity.auth.actor,
              "limit": 1,
              "key_type": "name",
              "index_position": "1",
              "json": true,
              "reverse": false
            }
        )
      });
      const body = await response.json()
      console.log(body.rows[0])
      self.myMult = parseFloat(body.rows[0]["fee_multiplier"].toString())
    },

    sendTx: async function () {
      let self = this

      const fee = async function (endpoint) {
        const response = await fetch(self.getEndpoint + "/v1/chain/get_fee", {
          method: 'POST',
          mode: 'cors',
          cache: 'no-cache',
          credentials: 'same-origin',
          redirect: 'error',
          referrerPolicy: 'no-referrer',
          body: JSON.stringify({end_point: endpoint, fio_address: "todd@blockpane"})
        });
        return await response.json().then((body) => {
          return parseInt(body.fee)
        })
      }
      const feeMultiplier = {
        account: 'fio.fee',
        name: 'setfeemult',
        authorization: [self.getAnchorIdentity.auth],
        data: {
          multiplier: self.multiplier,
          max_fee: await fee("submit_fee_multiplier"),
          actor: self.getAnchorIdentity.auth.actor,
        },
      }
      const feeVote = {
        account: 'fio.fee',
        name: 'setfeevote',
        authorization: [self.getAnchorIdentity.auth],
        data: {
          fee_ratios: self.fees,
          max_fee: await fee("submit_fee_ratios"),
          actor: self.getAnchorIdentity.auth.actor,
        },
      }
      let actions = [feeMultiplier]
      if (!self.onlyFees) {
        actions = [feeVote, feeMultiplier]
      }
      self.getAnchorIdentity.transact({actions: actions})
          .then(({signer, transaction}) => {
            self.setResult(`Success! Transaction signed by ${signer} and broadcast with transaction id: ${transaction.id}`)
          })
          .catch((e) => {
            self.setResult(e.toString())
          })
    },

  },
}

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style>
table#fee-table-transition .flip-list-move {
  transition: transform .5s;
}
</style>
