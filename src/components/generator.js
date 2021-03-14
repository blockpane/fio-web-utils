
async function hashI128(toHash) {
    const encoded = new TextEncoder().encode(toHash)
    return await crypto.subtle.digest("SHA-1", encoded).then( hash => {
        const hashArray = Array.from(new Uint8Array(hash)).slice(0, 16).reverse()
        return '0x' + hashArray.map(b => b.toString(16).padStart(2, '0')).join('')
    })
}

async function getBounds(form) {
    let bounds = ''
    if (form.transform === "hash") {
        if (form.lower !== "") {
            bounds = `"lower_bound": "${await hashI128(form.lower)}",` + "\n"
        }
        if (form.upper !== "") {
            bounds += `  "upper_bound": "${await hashI128(form.upper)}",`
        }
    } else if (form.showAdvanced === true) {
        bounds = `"lower_bound": "${form.lower}",
  "upper_bound": "${form.upper}",`
    } else {
        bounds = `"lower_bound": "${form.offset}",`
        form.type = "i64"
        form.index = 1
    }
    return bounds
}

export async function rawQuery(form) {
    const bounds = await getBounds(form)

    return `
{
  "code": "${ form.contract }",
  "scope": "${ form.scope }",
  "table": "${ form.table }",
  ${bounds}
  "limit": ${ form.numRows },
  "key_type": "${ form.type }",
  "index_position": "${ form.index }",
  "json": true,
  "reverse": ${ form.reverse }
}`
}

export function fioGoHashedQuery(form, url) {
    return `
package main

import (
	"encoding/json"
	"fmt"
	"github.com/fioprotocol/fio-go"
)

const (
	url = "${ url }"
	query = "${ form.lower }"
)

func main() {
	api, _, err := fio.NewConnection(nil, url)
	if err != nil {
		panic(err)
	}

	gtr, err := api.GetTableRowsOrder(fio.GetTableRowsOrderRequest{
		Code:       "${ form.contract }",
		Scope:      "${ form.scope }",
		Table:      "${ form.table }",
		LowerBound: fio.I128Hash(query), // converts to i128 hash
		UpperBound: fio.I128Hash(query),
		Limit:      ${ form.numRows },
		KeyType:    "i128",
		Index:      "${ form.index }",
		JSON:       true,
		Reverse:    ${ form.reverse },
	})
	if err != nil {
		panic(err)
	}

	// eos.GetTablesRowsResp.Rows is a []byte, and would normally be unmarshalled into 
	// a slice of whatever type was expected. for demonstration purposes, 
	// this pretty-prints it as indented json:
	j, err := json.MarshalIndent(json.RawMessage(gtr.Rows), "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(j))
}
`
}

export function fioGoNormalQuery(form, url) {
    let upLow = `lower = "${ form.lower }"
	upper = "${ form.upper }"`
    if (!form.showAdvanced) {
        upLow = `lower = "${ form.offset }"
        upper = ""`
        form.index = 1
        form.type = "i64"
    }
    return `
package main

import (
	"encoding/json"
	"fmt"
	"github.com/fioprotocol/fio-go"
)

const (
	url = "${ url }"
	${ upLow }
)

func main() {
	api, _, err := fio.NewConnection(nil, url)
	if err != nil {
		panic(err)
	}

	gtr, err := api.GetTableRowsOrder(fio.GetTableRowsOrderRequest{
		Code:       "${ form.contract }",
		Scope:      "${ form.scope }",
		Table:      "${ form.table }",
		LowerBound: lower,
		UpperBound: upper,
		Limit:      ${ form.numRows },
		KeyType:    "${ form.type }",
		Index:      "${ form.index }",
		JSON:       true,
		Reverse:    ${ form.reverse },
	})
	if err != nil {
		panic(err)
	}

	// eos.GetTablesRowsResp.Rows is a []byte, and would normally be unmarshalled into 
	// a slice of whatever type was expected. for demonstration purposes, 
	// this pretty-prints it as indented json:
	j, err := json.MarshalIndent(json.RawMessage(gtr.Rows), "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(j))
}
`
}

export function stdGoHashedQuery(form, url) {
    return `
package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	url   = "${ url }"
	query = "${ form.lower }"
)

type payload struct {
	Code          string \`json:"code"\`
	Scope         string \`json:"scope"\`
	Table         string \`json:"table"\`
	LowerBound    string \`json:"lower_bound"\`
	UpperBound    string \`json:"upper_bound"\`
	Limit         uint64 \`json:"limit"\`
	KeyType       string \`json:"key_type"\`
	Index         string \`json:"index"\`
	IndexPosition string \`json:"index_position"\`
	Json          bool   \`json:"json"\`
	Reverse       bool   \`json:"reverse"\`
}

func main() {

	// create our hash:
	sha := sha1.New() // #nosec
	_, err := sha.Write([]byte(query))
	if err != nil {
		panic(err)
	}
	b := sha.Sum(nil)
	// reverse endianness
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	// last 16 bytes of sha1-sum, as big-endian
	hash := "0x" + hex.EncodeToString(b)[8:]

	postBody, err := json.Marshal(payload{
		Code:          "${ form.contract }",
		Scope:         "${ form.scope }",
		Table:         "${ form.table }",
		LowerBound:    hash,
		UpperBound:    hash,
		Limit:         ${ form.numRows },
		KeyType:       "i128",
		IndexPosition: "${ form.index }",
		Json:          true,
		Reverse:       ${ form.reverse },
	})
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(url + "/v1/chain/get_table_rows", "application/json", bytes.NewReader(postBody))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	j, err := json.MarshalIndent(json.RawMessage(body), "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(j))
}

`
}

export function stdGoNormalQuery(form, url) {
    let upLow = `lower = "${ form.lower }"
	upper = "${ form.upper }"`
    if (!form.showAdvanced) {
        upLow = `lower = "${ form.offset }"
        upper = ""`
        form.index = 1
        form.type = "i64"
    }
    return `
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	url   = "${ url }"
	${ upLow }
)

type payload struct {
	Code          string \`json:"code"\`
	Scope         string \`json:"scope"\`
	Table         string \`json:"table"\`
	LowerBound    string \`json:"lower_bound"\`
	UpperBound    string \`json:"upper_bound"\`
	Limit         uint64 \`json:"limit"\`
	KeyType       string \`json:"key_type"\`
	IndexPosition string \`json:"index_position"\`
	Json          bool   \`json:"json"\`
	Reverse       bool   \`json:"reverse"\`
}

func main() {

	postBody, err := json.Marshal(payload{
		Code:          "${ form.contract }",
		Scope:         "${ form.scope }",
		Table:         "${ form.table }",
		LowerBound:    lower,
		UpperBound:    upper,
		Limit:         ${ form.numRows },
		KeyType:       "${ form.type }",
		IndexPosition: "${ form.index }",
		Json:          true,
		Reverse:       ${ form.reverse },
	})
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(url + "/v1/chain/get_table_rows", "application/json", bytes.NewReader(postBody))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	j, err := json.MarshalIndent(json.RawMessage(body), "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(j))
}

`
}

export function bashHashQuery(form, url) {
    return `
#!/bin/bash

# ensure pre-requisites are in path:
hash openssl && hash xxd && hash curl && hash jq || { echo "script requires openssl, vim, curl, and jq packages"; kill 0; }
# ubuntu's rev doesn't like binary, but tac is fine, MacOS doesn't have tac, but rev is ok with binary...
REV=$(which tac || which rev || { echo "could not find 'tac' or 'rev' utility"; kill 0; }) 2>/dev/null

URL=${ url }
QUERY=${ form.lower }

# last 16 bytes of sha1-sum, as hex-encoded big-endian
HASH="0x$(echo -n \${QUERY}| openssl sha1 -binary | LC_ALL=C $REV |xxd -p | cut -c 9-40)"

curl -s "\${URL}/v1/chain/get_table_rows" -d '{
  "code": "${ form.contract }",
  "scope": "${ form.scope }",
  "table": "${ form.table }",
  "lower_bound": "'\${HASH}'",
  "upper_bound": "'\${HASH}'",
  "limit": ${ form.numRows },
  "key_type": "${ form.type }",
  "index_position": "${ form.index }",
  "json": true,
  "reverse": ${ form.reverse }
}' | jq .

`
}


export async function bashNormalQuery(form, url) {
    const bounds = await getBounds(form)
    return `
#!/bin/bash

hash curl && hash jq || { echo "script requires curl and jq packages"; kill 0; }

URL=${ url }

curl -s "\${URL}/v1/chain/get_table_rows" -d '{
  "code": "${ form.contract }",
  "scope": "${ form.scope }",
  "table": "${ form.table }",
  ${ bounds }
  "limit": ${ form.numRows },
  "key_type": "${ form.type }",
  "index_position": "${ form.index }",
  "json": true,
  "reverse": ${ form.reverse }
}' | jq .

`
}

export function pythonHashQuery(form, url) {
    let rev = "False"
    if (form.reverse === true) {
        rev = "True"
    }
    return `
import hashlib
import json
import requests

query = b"${ form.lower }"
url = "${ url }"

h = hashlib.new('sha1')
h.update(query)

# last 16 bytes of sha1 hash in big endian
hash_hex = '0x' + h.digest()[15::-1].hex()

response = requests.post(url+"/v1/chain/get_table_rows", json={
    "code": "${ form.contract }",
    "scope": "${ form.scope }",
    "table": "${ form.table }",
    "lower_bound": hash_hex,
    "upper_bound": hash_hex,
    "limit": ${ form.numRows },
    "key_type": "i128",
    "index_position": "${ form.index }",
    "json": True,
    "reverse": ${ rev },
})

print(json.dumps(response.json(), indent=2))

`
}

export function pythonNormalQuery(form, url) {
    let rev = "False"
    if (form.reverse === true) {
        rev = "True"
    }
    let lower = form.lower
    let upper = form.upper
    if (!form.showAdvanced) {
        lower = form.offset
        upper = ""
        form.index = 1
        form.type = "i64"
    }

    return `
import json
import requests

lower = "${ lower }"
upper = "${ upper }"
url = "${ url }"

response = requests.post(url+"/v1/chain/get_table_rows", json={
    "code": "${ form.contract }",
    "scope": "${ form.scope }",
    "table": "${ form.table }",
    "lower_bound": lower,
    "upper_bound": upper,
    "limit": ${ form.numRows },
    "key_type": "${ form.type }",
    "index_position": "${ form.index }",
    "json": True,
    "reverse": ${ rev },
})

print(json.dumps(response.json(), indent=2))

`
}

export function nodeHashQuery(form, url) {
    return `
import axios from 'axios'
import crypto from 'crypto'

function nameHash(name) {
    const hash = crypto.createHash('sha1')
    return '0x' + hash.update(name).digest().slice(0,16).reverse().toString("hex")
}

async function getRows(hashed, url) {
    const resp = await axios.post(url + "/v1/chain/get_table_rows", {
        code: "${ form.contract }",
        scope: "${ form.scope }",
        table: "${ form.table }",
        lower_bound: hashed,
        upper_bound: hashed,
        limit: ${ form.numRows },
        key_type: "i128",
        index_position: "${ form.index }",
        json: true,
        reverse: ${ form.reverse }
    })
    return resp.data
}

const query = "${ form.lower }"
const url = "${ url }"

console.log(await getRows(nameHash(query), url))

`
}

export function nodeNormalQuery(form, url) {
    let lower = form.lower
    let upper = form.upper
    if (!form.showAdvanced) {
        lower = form.offset
        upper = ""
        form.index = 1
        form.type = "i64"
    }

    return `
import axios from 'axios'

async function getRows(lower, upper, url) {
    const resp = await axios.post(url + "/v1/chain/get_table_rows", {
        code: "${ form.contract }",
        scope: "${ form.scope }",
        table: "${ form.table }",
        lower_bound: lower,
        upper_bound: upper,
        limit: ${ form.numRows },
        key_type: "${ form.type }",
        index_position: "${ form.index }",
        json: true,
        reverse: ${ form.reverse }
    })
    return resp.data
}

const lower = "${ lower }"
const upper = "${ upper }"
const url = "${ url }"

console.log(await getRows(lower, upper, url))

`
}

export function browserHashQuery(form, url) {
    return `
<!DOCTYPE html>
<html lang="">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width,initial-scale=1.0">
    <title>Hash Example</title>
</head>
<body>
<div><pre id="output"></pre></div>
<script>
    const url = "${ url }"
    const query = "${ form.lower }"

    async function hashI128(toHash) {
        const encoded = new TextEncoder().encode(toHash)
        // note: crypto.subtle is ONLY available in a secure context, so this only works on https sites!
        return await crypto.subtle.digest("SHA-1", encoded).then( hash => {
            const hashArray = Array.from(new Uint8Array(hash)).slice(0, 16).reverse()
            return '0x' + hashArray.map(b => b.toString(16).padStart(2, '0')).join('')
        })
    }

    async function hashedQuery(query, url) {
        try {
            const response = await fetch(url + "/v1/chain/get_table_rows", {
                method: 'POST',
                mode: 'cors',
                cache: 'no-cache',
                credentials: 'same-origin',
                redirect: 'error',
                referrerPolicy: 'no-referrer',
                body: JSON.stringify({
                    code: "${ form.contract }",
                    scope: "${ form.scope }",
                    table: "${ form.table }",
                    lower_bound: await hashI128(query),
                    upper_bound: await hashI128(query),
                    limit: ${ form.numRows },
                    key_type: "i128",
                    index_position: "${ form.index }",
                    json: true,
                    reverse: ${ form.reverse },
                })
            });
            const body = await response.json()
            document.getElementById("output").innerHTML = JSON.stringify(body, null, 4)
        } catch (e) {
            document.getElementById("output").innerHTML = e.toString()
        }
    }

    hashedQuery(query, url)
</script>
</body>
</html>

`
}

export function browserNormalQuery(form, url) {
    let lower = form.lower
    let upper = form.upper
    if (!form.showAdvanced) {
        lower = form.offset
        upper = ""
        form.index = 1
        form.type = "i64"
    }

    return `
<!DOCTYPE html>
<html lang="">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width,initial-scale=1.0">
    <title>Example</title>
</head>
<body>
    <div><pre id="output"></pre></div>
    <script>
        const url = "${ url }"
        const upper = "${ upper }"
        const lower = "${ lower }"
        
        async function query(lower, upper, url) {
            try {
                const response = await fetch(url + "/v1/chain/get_table_rows", {
                    method: 'POST',
                    mode: 'cors',
                    cache: 'no-cache',
                    credentials: 'same-origin',
                    redirect: 'error',
                    referrerPolicy: 'no-referrer',
                    body: JSON.stringify({
                        code: "${ form.contract }",
                        scope: "${ form.scope }",
                        table: "${ form.table }",
                        lower_bound: lower,
                        upper_bound: upper,
                        limit: ${ form.numRows },
                        key_type: "${ form.type }",
                        index_position: "${ form.index }",
                        json: true,
                        reverse: ${ form.reverse },
                    })
                });
                const body = await response.json()
                document.getElementById("output").innerHTML = JSON.stringify(body, null, 4)
            } catch (e) {
                document.getElementById("output").innerHTML = e.toString()
            }
        }
        
        query(lower, upper, url)
    </script>
</body>
</html>

`
}

export default {
    rawQuery,
    fioGoHashedQuery,
    fioGoNormalQuery,
    stdGoHashedQuery,
    bashHashQuery,
    bashNormalQuery,
    pythonHashQuery,
    pythonNormalQuery,
    nodeHashQuery,
    nodeNormalQuery,
    browserHashQuery,
    browserNormalQuery,
}
