description: "Link from hidden to parent is bad; bad seqtype"

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
            arg.parent_chain_tail.chain_type = 1000
            arg
      },
      { type : "rotate_key" },
    ]
  }
}

sessions: [
  { loads : [{
    error : true
    error_type_full : "hidden.ParentPointerError"
    error_substr : "hidden team parent pointer error (to visible 2): wrong chain type"
  }]}
]
