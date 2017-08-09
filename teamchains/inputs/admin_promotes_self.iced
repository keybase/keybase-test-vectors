description: "Admin tries to change owner list (boost self)"

users: {
  "herb": {}
  "basil": {}
}

teams: {
  "cabal": {
    links: [
      type: "root"
      signer: "herb"
      members:
        owner: ["herb"]
        admin: ["basil"]
    ,
      # invalid link
      # basil tries to promote itself to owner
      type: "change_membership"
      signer: "basil"
      members:
        owner: ["basil"]
    ]
  }
}

sessions: [
  loads: [
    error: true
    error_substr: "non-owner cannot add owners"
  ]
,
  loads: [
    upto: 1
  ,
    error: true
    error_substr: "non-owner cannot add owners"
  ]
]
