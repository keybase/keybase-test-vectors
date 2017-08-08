description: "Chain has a blank per team key encryption kid"

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
        per_team_key: (section) ->
          delete section.encryption_kid
          section
    ]
  }
}

expect: {
  error: true
  error_substr: "invalid per-team-key encryption KID"
}
