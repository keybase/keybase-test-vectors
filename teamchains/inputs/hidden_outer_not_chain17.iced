description: "bad chain type"

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
            outer[5]++
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
    error_substr : "sig3 error: can only handle type 17 (team private hidden)"
  }]}
]
