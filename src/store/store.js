import Vue from 'vue'
import Vuex from 'vuex'
Vue.use(Vuex)

import AnchorLinkBrowserTransport from "anchor-link-browser-transport";
import AnchorLink from "anchor-link";

export const store = new Vuex.Store({
    state: {
        pageTitle: "",
        anchorConnected: false,
        anchorIdentity: {},
        anchorTransport: {},
        anchorLink: {},
        resultMessage: "Connect to Anchor Wallet to see current vote and to submit fees",
        network: "mainnet",
        mainnet: "21dcae42c0182200e93f954a074011f9048a7624c6fe81d3c9541a614a88bd1c",
        testnet: "b20901380af44ef59c5918439a1f9a41d83669020319a80574b804a5f95cbd7e",
        chainId: "21dcae42c0182200e93f954a074011f9048a7624c6fe81d3c9541a614a88bd1c",
        endpoint: "https://fio.blockpane.com",
        mainEndpoint: "https://fio.blockpane.com",
        testEndpoint: "https://fiotestnet.blockpane.com",
    },

    getters: {
        getPageTitle: state => {return state.pageTitle},
        getAnchorConnected: state => { return state.anchorConnected },
        getAnchorIdentity: state => { return state.anchorIdentity },
        getAnchorTransport: state => { return state.anchorTransport },
        getAnchorLink: state => { return state.anchorLink},
        getResultMessage: state => { return state.resultMessage},
        getNetwork: state => { return state.network},
        getChainId: state => { return state.chainId},
        getEndpoint: state => { return state.endpoint},
    },

    mutations: {
        setPageTitle (state, newTitle) {
            state.pageTitle = newTitle
        },
        setConnected (state, isCon) {
            state.anchorConnected = isCon
        },
        setResult (state, msg) {
            state.resultMessage = msg
        },

        setNetwork (state, newEndpoint) {
            console.log(newEndpoint.net + " new endpoint: " + newEndpoint.endpoint)
            state.network = newEndpoint.net
            state.chainId = state.mainnet
            if ( newEndpoint.endpoint === undefined || newEndpoint.endpoint === "" ) {
                if (state.network === "testnet") {
                    state.endpoint = state.testEndpoint
                } else {
                    state.network = "mainnet"
                    state.endpoint = state.mainEndpoint
                    state.chainId = state.mainnet
                }
            } else {
                state.endpoint = newEndpoint.endpoint
            }
            if (newEndpoint.net === "testnet") {
                state.chainId = state.testnet
            }
            state.resultMessage = `changing network to: ${ state.network } via ${ state.endpoint }`
            state.anchorConnected = false
        },

        setTransport (state, trans) {
            state.anchorTransport = trans
        },
        resetAnchorLink (state) {
            state.anchorLink = null
        },
        setAnchorLink (state) {
            state.anchorLink = new AnchorLink({
                transport: state.anchorTransport,
                chains: [
                    {
                        chainId: state.chainId,
                        nodeUrl: state.endpoint,
                    }
                ],
            })
        },

        setAnchorId (state, id) {
            state.anchorIdentity = id
        },
    },

    actions: {
        async anchorLogin (context) {
            context.commit('setResult', "attempting wallet authentication")
            context.commit('setTransport', new AnchorLinkBrowserTransport())
            context.commit('setAnchorLink')
            let link = context.getters.getAnchorLink
            let result = {}
            // TODO: figure out how to restore a session, this doesn't work.
            try {
                result = await link.restoreSession({identifier: 'fio.utils', chainId: context.getters.getChainId})
            } catch {
                result = await link.login('fio.utils')
            }
            context.commit('setAnchorId', result.session)
            context.commit('setResult', `Authenticated as ${result.session.auth.toString()}`)
            context.commit('setConnected', true)
            if (context.getters.getAnchorIdentity === undefined || context.anchorIdentity.auth.actor === undefined
                || context.anchorIdentity.auth.actor === null) {
                context.commit('setConnected', false)
                context.commit('setResult', "Anchor login failed")
            }
        },
    },
})

