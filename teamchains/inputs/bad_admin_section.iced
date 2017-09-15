description: "admin section is wrong"

users: {
  "herb": {}
  "basil": {}
}

teams: {
  "cabal": {
    links: [
      type: "root"
      members: owner: ["herb"]
    ,
      type: "change_membership"
      members: writer: ["basil"]
      admin:
        team_id: "21000000000000000600000000000024"
        seqno: 1,
        seq_type: 3,
    ]
  }
}

sessions: [
  loads: [
    error: true
    error_type: "AdminNotFoundError"
  ]
,
  loads: [
    upto: 1
  ,
    error: true,
    error_type: "AdminNotFoundError"
  ]
]
