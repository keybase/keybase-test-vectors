description: "Admin link signed by a user who was kicked out of the team"

users: {
  "herb": {}
  "basil": {}
  "pepper": {}
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
      type: "change_membership"
      signer: "herb"
      members:
        none: ["basil"]
    ,
      # invalid link
      # basil tries to use their admin privelege
      type: "change_membership"
      signer: "basil"
      members:
        reader: ["pepper"]
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
    upto: 2
  ,
    error: true,
    error_type: "AdminPermissionError"
  ]
]
