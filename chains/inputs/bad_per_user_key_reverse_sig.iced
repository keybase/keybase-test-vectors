chain :
  user : "max32"
  ctime : "now"
  expire : 10000000
  links : [
    {
      type : "eldest"
      label : "e"
      key : gen : "eddsa"
    },
    {
      ctime : "+100"
      label : "puk1"
      type : "per_user_key"
      signer : "e"
      corrupt_v1_proof_hooks :
        corrupt : ({obj}) ->
          obj.body.per_user_key.reverse_sig = 'a' + obj.body.per_user_key.reverse_sig[1..]
          #return obj
    }
  ]
