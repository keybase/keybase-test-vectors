description: "Link from hidden to parent is bad; won't fail for now until we fix checking code in the team loader"

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
  }]}
]

skip : true
