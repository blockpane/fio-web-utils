
async function hashI128(toHash) {
    const encoded = new TextEncoder().encode(toHash)
    return await crypto.subtle.digest("SHA-1", encoded).then( hash => {
        const hashArray = Array.from(new Uint8Array(hash)).slice(0, 16).reverse()
        return '0x' + hashArray.map(b => b.toString(16).padStart(2, '0')).join('')
    })
}

export async function rawQuery(form) {
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

export default {
    rawQuery,
}
