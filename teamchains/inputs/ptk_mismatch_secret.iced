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
    error: true
    error_substr: "wrong sigKID"
  ]
]
