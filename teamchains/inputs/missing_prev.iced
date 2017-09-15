description: "server responds with a gap in prevs"

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
    ,
      type: "rotate_key"
    ]
  }
}

sessions: [
  loads: [
    error: false
  ]
,
  loads: [
    upto: 2
  ,
    # omit prevs that contain the secret
    # for generations <= to this number
    omit_prevs: 3
    error: true
    error_substr: "gap in per-team-keys"
  ]
,
  loads: [
    omit_prevs: 3
    error: true
    error_substr: "gap in per-team-keys"
  ]
]
