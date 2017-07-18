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
    }
  ]
