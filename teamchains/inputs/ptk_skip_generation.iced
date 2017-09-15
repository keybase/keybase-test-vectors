description: "skip a generation in the per-team-key"

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
    ,
      type: "rotate_key"
      but_dont_make_a_link: true
    ,
      type: "rotate_key"
    ]
  }
}

sessions: [
  loads: [
    error: true
    error_substr: "per-team-key generation expected 3 but got 4"
  ]
]
