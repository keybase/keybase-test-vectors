description: "per-team-key box is different from the chain"

users: {
  "herb": {}
}

teams: {
  "cabal": {
    links: [
      type: "root"
      signer: "herb"
      members:
        owner: ["herb"]
    ,
      type: "rotate_key"
      use_other_key: true
    ,
      type: "rotate_key"
    ]
  }
}

sessions: [
  loads: [
    error_after_get_key: true
    error_substr: "wrong team key found at generation 2"
    then_get_key: 2 # Team loader does not check keys until they are fetched
  ]
]
