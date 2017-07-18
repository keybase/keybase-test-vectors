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
        corrupt_for_reverse_signature: ({obj}) ->
          obj.hello = 'world'
        corrupt_key_section : 'per_user_key'
    }
  ]
