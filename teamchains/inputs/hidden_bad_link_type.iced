description: "bad link type advertised in outer chainlink"

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
            outer[4] = 1000
            outer
      },
      { type : "rotate_key" },
    ]
  }
}

sessions: [
  { loads : [{
    error : true
    error_type : "ParseError"
    error_type_full : "sig3.ParseError"
    error_substr : "unknown link type 1000"
  }]}
]
