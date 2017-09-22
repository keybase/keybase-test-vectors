description: "corrupted team key box"

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
      corrupt_box: true
    ]
  }
}

sessions: [
  loads: [
    error: true
    error_substr: "opening key box: box.Open failure"
  ]
,
  loads: [
    upto: 1
  ,
    error: true
    error_substr: "opening key box: box.Open failure"
  ]
]
