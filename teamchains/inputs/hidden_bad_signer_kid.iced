description: "sig3: bad signing KID"

users: {
  "herb": {}
}

teams: {
  cabal : {
    links : [
      { type: "root", members: { owner: ["herb"] } },
      { type : "rotate_key" },
      {
        type : "rotate_key_hidden"
        corruptors :
          sig3_patch_inner : (json) ->
            json.s.k[10] ^= 0x1 # corrupt a bit in the 10th byte of the signer KI
            json
      },
      { type : "rotate_key" },
    ]
  }
}

sessions: [
  { loads : [{
    error : true
    error_type : "Sig3Error"
    error_type_full : "sig3.Sig3Error"
    error_substr : "sig3 error: signature verification failed"
  }]}
]
