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

sessions: [
  loads: [
    error_after_get_key: true
    error_substr: "wrong team key (enc) found at generation 1"
    then_get_key: 1 # Team loader does not check keys until they are fetched
  ]
]
