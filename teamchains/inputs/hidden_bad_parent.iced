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
          sig_arg : (arg) ->
            arg.parent_chain_tail.hash[0] ^= 0x01
            arg
      },
      { type : "rotate_key" },
    ]
  }
}

sessions: [
  { loads : [{
    error : true
    error_type : "ParseError"
    error_substr : "unknown link type 1000"
  }]}
]
