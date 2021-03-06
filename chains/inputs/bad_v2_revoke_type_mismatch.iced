
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
    },
    {
      label : "btc1"
      type : "btc"
      ctime : "+100"
      signer : "sib1"
      version : 2
    },
    {
      label : "btc2"
      type : "btc"
      ctime : "+100"
      signer : "sib1"
      version : 2
      revoke : sigs : [ "btc1" ]
      corrupt_v2_proof_hooks :
        pre_generate_outer : ({proof}) ->
          proof._type_v2 = () -> 6
    }
  ]
