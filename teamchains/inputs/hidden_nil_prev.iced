description: "Simplest test"

users: {
  "herb": {}
}

teams: {
  cabal : {
    links : [
      { type: "root", members: { owner: ["herb"] } },
      { type : "rotate_key" },
      { type : "rotate_key_hidden" }
      { type : "rotate_key" },
      {
        type : "rotate_key_hidden"
        corruptors :
          sig_arg : (arg) ->
            arg.prev = null
            arg
      }
    ]
  }
}

sessions: [
  { loads : [{
    error : true
    error_type : "SequenceError"
    error_substr : "sig3 sequencing error: bad nil prev at 2"
  }]}
]
