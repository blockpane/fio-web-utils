
export function rawQuery(form) {
    let bounds = ''
    if (form.showAdvanced === true) {
        bounds = `"lower_bound": "${ form.lower }",
  "upper_bound": "${ form.upper }",`
    } else {
        bounds = `"lower_bound": "${ form.offset }",`
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
