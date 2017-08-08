description: "Admin link signed by a user not in the team"

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
    ,
      # invalid link
      # basil tries to sign themself in
      type: "change_membership"
      signer: "basil"
      members:
        none: ["basil"]
    ]
  }
}

sessions: [
  loads: [
    error: true,
    error_type: "AdminPermissionError"
  ]
,
  loads: [
    upto: 1
  ,
    error: true,
    error_type: "AdminPermissionError"
  ]
]
