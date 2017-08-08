description: "Chain has the wrong per team key encryption kid"

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
          section.encryption_kid = "01218ba2aa312e74a292ce6e8136fa8343bd5146acbb5b60e30ad3d29e2ae67bd53c0a"
          section
    ]
  }
}

expect: {
  error: true
  error_substr: "wrong encKID"
}
