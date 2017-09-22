description: "per-team-key box has a too-high generation"

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
    ]
  }
}

sessions: [
  loads: [
    force_last_box: true
    error: true
    error_substr: "wrong latest key generation"
  ]
]
