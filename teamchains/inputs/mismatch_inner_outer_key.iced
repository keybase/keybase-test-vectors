description: "different keys on the inner link (declared) and outer link (signed)"

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
      corruptors:
        force_inner_key: {user: "basil"}
    ]
  }
}

sessions: [
  loads: [
    error: true
    error_type: "WrongKidError"
  ]
,
  loads: [
    upto: 1
  ,
    error: true,
    error_type: "WrongKidError"
  ]
]
