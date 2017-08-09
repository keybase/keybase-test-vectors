description: "Returning stubbed links should work"

users: {
  "herb": {}
  "basil": {}
}

teams: {
  "cabal": {
    links: [
      type: "root"
      signer: "basil"
      members:
        owner: ["basil"]
        writer: ["herb"]
    ,
      type: "invite"
      signer: "basil"
    ,
      type: "invite"
      signer: "basil"
    ]
  }
}

sessions: [
  loads: [
    stub: [2, 3] # Stub these seqnos
    n_stubbed: 2
  ]
,
  loads: [
    upto: 1
    n_stubbed: 0
  ,
    upto: 2
    stub: [2, 3]
    n_stubbed: 1
  ,
    upto: 3
    stub: [2, 3]
    n_stubbed: 2
  ]
,
  loads: [
    upto: 2
    stub: [2, 3]
    n_stubbed: 1
  ,
    stub: [2, 3]
    n_stubbed: 2
  ]
]
