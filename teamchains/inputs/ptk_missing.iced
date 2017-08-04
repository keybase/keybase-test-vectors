description: "Chain is missing a ptk"

users: {
  "herb": {}
  "basil": {}
  "rose": {}
}

teams: {
  "cabal": {
    links: [
      type: "root"
      signer: "herb"
      members:
        owner: ["herb"]
        writer: ["basil"]
      corruptors:
        sig_arg: (sig_arg) ->
          delete sig_arg.kms
          sig_arg
    ]
  }
}

expect: {
  error: true
  error_substr: "missing per-team-key"
}
