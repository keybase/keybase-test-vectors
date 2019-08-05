description: "sig3: hash of innerlink doesn't match inner field of outer"

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
          sig3_corrupt_outer : (outer) ->
            # outer[3] is InnerLinkID; corrupt the second byte of it
            outer[3][1] ^= 0x01
            outer
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
    error_substr : "sig3 error: inner link hash doesn't match inner"
  }]}
]
