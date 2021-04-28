<template>
  <div>
    <b-table
        hover small striped borderless responsive="true" head-variant="dark"
        :items="activeProxies"
    ></b-table>
  </div>
</template>

<script>
import { mapGetters, mapMutations, mapActions } from 'vuex'

export default {
  name: 'Proxy',

  data () {
    return {
      activeProxies: [],
      activeProducers: [],
      accountMap: {},
      // tracks mirrored votes, ie non-unique
      dupVoteMap: {
        byHash: {},
        byVoter: {},
      },
    }
  },

  mounted() {
    this.refresh()
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
  },

  watch: {
    getEndpoint: function () {
      //console.log("watch triggered")
      this.dupVoteMap =  {
        byHash: {},
        byVoter: {},
      }
      this.refresh()
    },
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

    refresh: async function () {
      await this.getAllProxies()
      await this.getAllProducers()
    },

    getAllProxies: async function () {
      let self = this
      self.activeProxies = []
      const response = await fetch(self.getEndpoint + "/v1/chain/get_table_rows", {
            method: 'POST',
            mode: 'cors',
            cache: 'no-cache',
            credentials: 'same-origin',
            redirect: 'error',
            referrerPolicy: 'no-referrer',
            body: JSON.stringify({
                code: "eosio",
                scope: "eosio",
                table: "voters",
                lower_bound: "",
                upper_bound: "",
                limit: -1,
                key_type: "i64",
                index_position: "1",
                json: true,
                reverse: false,
            })
      });
      const body = await response.json()
      for (const voter of body.rows) {
        if (voter.is_proxy === 1 && voter.producers.length > 0) {
          self.activeProxies.push(voter)
          await self.hashVote(voter)
          if (voter.fioaddress !== "" && self.accountMap[voter.owner] !== null) {
            self.accountMap[voter.owner] = voter.fioaddress
          }
        }
      }
    },

    getAllProducers: async function () {
      let self = this
      self.activeProducers = []
      const response = await fetch(self.getEndpoint + "/v1/chain/get_table_rows", {
        method: 'POST',
        mode: 'cors',
        cache: 'no-cache',
        credentials: 'same-origin',
        redirect: 'error',
        referrerPolicy: 'no-referrer',
        body: JSON.stringify({
          code: "eosio",
          scope: "eosio",
          table: "producers",
          lower_bound: "",
          upper_bound: "",
          limit: -1,
          key_type: "i64",
          index_position: "1",
          json: true,
          reverse: false,
        })
      });
      const body = await response.json()
      for (const producer of body.rows) {
        if (producer.is_active === 1 ) {
          self.activeProducers.push(producer)
          if (producer.fio_address !== "" && self.accountMap[producer.owner] !== null) {
            self.accountMap[producer.owner] = producer.fio_address
          }
        }
      }
    },

    hashVote: async function (voter) {
      voter.producers.sort()
      const hashed = await crypto.subtle.digest("SHA-256", new Buffer(voter.producers.join()))
      const hex = Buffer.from(hashed).toString("hex")
      if (this.dupVoteMap.byHash[hex] === undefined) {
        this.dupVoteMap.byHash[hex] = [voter.owner]
      } else {
        if (!this.dupVoteMap.byHash[hex].contains(voter.owner)) {
          this.dupVoteMap.byHash[hex].push(voter.owner)
        }
      }
      this.dupVoteMap.byVoter[voter.owner] = hex
      //console.log(this.dupVoteMap)
    },

  },

}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
