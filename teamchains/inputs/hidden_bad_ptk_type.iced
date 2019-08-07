description: "unknown PTK type still should work"

users: {
  "herb": {}
}

teams: {
  cabal : {
    links : [
      { type: "root", members: { owner: ["herb"] } }
      { type : "rotate_key" }
      {
        type : "rotate_key_hidden",
        corruptors : {
          sig_arg : (arg) ->
            arg.per_team_keys.ptk_type++
            arg
        }
      }
      { type : "rotate_key" },
    ]
  }
}

sessions: [
  { loads : [{
    error : false
  }]}
]
