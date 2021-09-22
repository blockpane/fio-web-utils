package stat

import (
	"github.com/blockpane/fio-web-utils/stats-api/models"
	"sync"
)

var (
	feeMap = map[string]string{
		"add_nft":                    "addnft",
		"remove_all_nfts":            "remove_all_nfts",
		"remove_nft":                 "remallnfts",
		"add_bundled_transactions":   "addbundles",
		"add_pub_address":            "addaddress",
		"cancel_funds_request":       "cancelfndreq",
		"remove_pub_address":         "remaddress",
		"msig_unapprove":             "unapprove",
		"vote_producer":              "voteproducer",
		"register_fio_domain":        "regdomain",
		"register_producer":          "regproducer",
		"msig_approve":               "approve",
		"submit_bundled_transaction": "bundlevote",
		"msig_cancel":                "cancel",
		"auth_link":                  "linkauth",
		"new_funds_request":          "newfundsreq",
		"record_obt_data":            "recordobt",
		"auth_update":                "updateauth",
		"set_fio_domain_public":      "setdomainpub",
		"submit_fee_multiplier":      "setfeemult",
		"unregister_proxy":           "unregproxy",
		"submit_fee_ratios":          "setfeevote",
		"burn_fio_address":           "burnaddress",
		"msig_invalidate":            "invalidate",
		"register_proxy":             "regproxy",
		"remove_all_pub_addresses":   "remalladdr",
		"renew_fio_address":          "renewaddress",
		"renew_fio_domain":           "renewdomain",
		"msig_exec":                  "exec",
		"msig_propose":               "propose",
		"unregister_producer":        "unregprod",
		"auth_delete":                "deleteauth",
		"register_fio_address":       "regaddress",
		"transfer_tokens_pub_key":    "trnsfiopubky",
		"transfer_locked_tokens":     "trnsloctoks",
		"transfer_fio_domain":        "xferdomain",
		"reject_funds_request":       "rejectfndreq",
		"proxy_vote":                 "voteproxy",
		"transfer_fio_address":       "xferaddress",
	}
	feeMapMux = sync.RWMutex{}
)

func s(str string) *string {
	return &str
}

func u(u64 uint64) *uint64 {
	return &u64
}

func suggestFee() []*models.Feevote {
	return []*models.Feevote{
		{EndPoint: s("add_bundled_transactions"), Value: u(2000000000)},
		{EndPoint: s("add_nft"), Value: u(30000000)},
		{EndPoint: s("add_pub_address"), Value: u(30000000)},
		{EndPoint: s("add_to_whitelist"), Value: u(30000000)},
		{EndPoint: s("auth_delete"), Value: u(20000000)},
		{EndPoint: s("auth_link"), Value: u(20000000)},
		{EndPoint: s("auth_update"), Value: u(50000000)},
		{EndPoint: s("burn_fio_address"), Value: u(60000000)},
		{EndPoint: s("cancel_funds_request"), Value: u(60000000)},
		{EndPoint: s("msig_approve"), Value: u(20000000)},
		{EndPoint: s("msig_cancel"), Value: u(20000000)},
		{EndPoint: s("msig_exec"), Value: u(20000000)},
		{EndPoint: s("msig_invalidate"), Value: u(20000000)},
		{EndPoint: s("msig_propose"), Value: u(50000000)},
		{EndPoint: s("msig_unapprove"), Value: u(20000000)},
		{EndPoint: s("new_funds_request"), Value: u(60000000)},
		{EndPoint: s("proxy_vote"), Value: u(30000000)},
		{EndPoint: s("record_obt_data"), Value: u(60000000)},
		{EndPoint: s("register_fio_address"), Value: u(2000000000)},
		{EndPoint: s("register_fio_domain"), Value: u(40000000000)},
		{EndPoint: s("register_producer"), Value: u(10000000000)},
		{EndPoint: s("register_proxy"), Value: u(1000000000)},
		{EndPoint: s("reject_funds_request"), Value: u(30000000)},
		{EndPoint: s("remove_all_nfts"), Value: u(60000000)},
		{EndPoint: s("remove_all_pub_addresses"), Value: u(60000000)},
		{EndPoint: s("remove_from_whitelist"), Value: u(30000000)},
		{EndPoint: s("remove_nft"), Value: u(60000000)},
		{EndPoint: s("remove_pub_address"), Value: u(60000000)},
		{EndPoint: s("renew_fio_address"), Value: u(2000000000)},
		{EndPoint: s("renew_fio_domain"), Value: u(40000000000)},
		{EndPoint: s("set_fio_domain_public"), Value: u(30000000)},
		{EndPoint: s("submit_bundled_transaction"), Value: u(30000000)},
		{EndPoint: s("submit_fee_multiplier"), Value: u(10000000)}, // I'm overriding the default here, this should be cheap.
		{EndPoint: s("submit_fee_ratios"), Value: u(70000000)},
		{EndPoint: s("transfer_fio_address"), Value: u(60000000)},
		{EndPoint: s("transfer_fio_domain"), Value: u(100000000)},
		{EndPoint: s("transfer_locked_tokens"), Value: u(100000000)},
		{EndPoint: s("transfer_tokens_pub_key"), Value: u(100000000)},
		{EndPoint: s("unregister_producer"), Value: u(20000000)},
		{EndPoint: s("unregister_proxy"), Value: u(20000000)},
		{EndPoint: s("vote_producer"), Value: u(30000000)},
	}
}
