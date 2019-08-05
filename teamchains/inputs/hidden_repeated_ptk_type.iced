description: "repeating a PTK type twice in the same link; they are supposed to be destinct (and should only be READER_PTK)"

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
            arg.per_team_keys.push arg.per_team_keys[0]
            arg
        }
      }
      { type : "rotate_key" },
    ]
  }
}

sessions: [
  { loads : [{
    error : true
    error_type : "ParseError"
    error_type_full : "sig3.ParseError"
    error_substr : "duplicated PTK type: READER"
  }]}
]
