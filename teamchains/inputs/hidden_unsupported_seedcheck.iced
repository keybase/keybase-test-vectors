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
      { type : "rotate_key" }
      {
        type : "rotate_key_hidden"
        corruptors :
          sig_arg : (arg) ->
            arg.per_team_keys[0].seed_check.v = 2
            arg
      }
    ]
  }
}

sessions: [
  {
    loads : [
      {
        error : true
        error_type : "LoaderError"
        error_type_full : "hidden.LoaderError"
        error_substr : "hidden team loader error: can only handle seed check version 1; got 2"
      }
    ]
  }
]
