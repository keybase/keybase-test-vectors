description: "corrupted team key prev"

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
      corrupt_prev: true
    ]
  }
}

sessions: [
  loads: [
    error: true
    error_substr: "opening prev"
  ]
,
  loads: [
    upto: 1
  ,
    error: true
    error_substr: "opening prev"
  ]
]
