description: "Link missing the whole team section"

users: {
  "herb": {}
  "basil": {}
}

teams: {
  "cabal": {
    links: [
      type: "root"
      signer: "herb"
      members:
        owner: ["herb"]
        admin: ["basil"]
    ,
      # invalid link: missing team section
      type: "change_membership"
      signer: "basil"
      members:
        writer: ["basil"]
      corruptors:
        sig_arg: (sig_arg) ->
          delete sig_arg.team
          sig_arg
    ]
  }
}

expect: {
  error: true
  error_substr: "no team section"
}
