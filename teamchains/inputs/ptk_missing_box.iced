description: "no per-team-key box in response"

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
    ]
  }
}

sessions: [
  loads: [
    omit_box: true
    error: true
    error_substr: "no key box"
  ]
,
  loads: [
    upto: 1
  ,
    omit_box: true
    error: true
    error_substr: "no key box"
  ]
]
