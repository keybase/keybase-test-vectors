description: "Simplest test"

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
            outer[0]++
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
    error_substr : "sig3 error: can only handle sig version 3 (got 4)"
  }]}
]
