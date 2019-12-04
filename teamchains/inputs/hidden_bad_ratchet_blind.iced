description: "the ratchet blinding/unblinding sent down from the server is bad"

users: {
  "herb": {}
}

teams: {
  cabal : {
    links : [
      { type: "root", members: { owner: ["herb"] } },
      { type : "rotate_key" },
      { type : "rotate_key_hidden" }
      {
        type : "rotate_key"
        corruptors :
          generated_ratchet : (r) ->
            r.ratchet[1] ^= 0x1
      }
    ]
  }
}

load_failure :
  error : true
  error_substr : "hidden team ratchet error: blinding check failed 01ea732f329d2fd83bbe0c6bd86f3f8562dde30c0e6c66be45393c510fd5104c v 01eb732f329d2fd83bbe0c6bd86f3f8562dde30c0e6c66be45393c510fd5104c"
  error_type_full : "hidden.RatchetError"
