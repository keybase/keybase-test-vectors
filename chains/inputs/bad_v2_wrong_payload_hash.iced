
mangle_hex = (x) ->
  b = new Buffer x, 'hex'
  b[0] ^= 1
  b.toString('hex')

chain :
  user : "max32"
  ctime : "now"
  expire : 10000000
  links : [
    {
      type : "eldest"
      label : "e"
      key : gen : "eddsa"
      version : 2
    },
    {
      ctime : "+100"
      label : "sib1"
      type : "sibkey"
      key : gen : "eddsa"
      signer : "e"
      version : 2
      corrupt_v2_proof_hooks :
        corrupt_ids : (d) ->
          d.outer = Buffer.concat [ d.outer, Buffer.from([1]) ]
    }
  ]
