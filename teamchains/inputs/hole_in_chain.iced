description: "Server omits a link"

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
        owner: ["basil", "herb"]
    ,
      type: "change_membership"
      signer: "basil"
      members:
        owner: ["herb"]
    ,
      type: "change_membership"
      members:
        admin: ["herb"]
    ]
  }
}

sessions: [
  loads: [
    omit: [2] # omit these seqnos
    error: true
    error_type: "PrevError"
  ]
,
  loads: [
    upto: 1
  ,
    omit: [2] # omit these seqnos
    error: true
    error_type: "PrevError"
  ]
]
