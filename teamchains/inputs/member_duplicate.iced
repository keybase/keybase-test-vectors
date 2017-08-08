description: "Duplicate members in a link"

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
        none: ["basil"]
    ,
      type: "change_membership"
      members: # this link is invalid because it contains the same user twice
        admin: ["basil"]
        none: ["basil"]
    ]
  }
}

sessions: [
  loads: [
    error: true,
    error_substr: "duplicate UID in members"
  ]
,
  loads: [
    upto: 1
  ,
    upto: 2
  ,
    error: true,
    error_substr: "duplicate UID in members"
  ]
,
  loads: [
    upto: 1
  ,
    error: true,
    error_substr: "duplicate UID in members"
  ]
]
