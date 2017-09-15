description: "Admin link signed by a user who was kicked out of the team (later so it's ok)"

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
      signer: "basil"
      members:
        reader: ["pepper"]
    ,
      type: "change_membership"
      signer: "herb"
      members:
        none: ["basil"]
    ]
  }
}

sessions: [
  loads: [
    error: false,
  ]
,
  loads: [
    upto: 2
  ,
    error: false,
  ]
]
