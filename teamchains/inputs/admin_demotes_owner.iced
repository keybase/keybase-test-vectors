description: "Admin tries to change owner list (demote other)"

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
      # invalid link - basil tries to demote an owner
      type: "change_membership"
      signer: "basil"
      members:
        admin: ["herb"]
    ]
  }
}

sessions: [
  loads: [
    error: true
    error_substr: "non-owner cannot demote owners"
  ]
,
  loads: [
    upto: 1
  ,
    error: true
    error_substr: "non-owner cannot demote owners"
  ]
]
