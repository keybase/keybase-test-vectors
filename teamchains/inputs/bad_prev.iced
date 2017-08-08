description: "Link has wrong prev"

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
      # invalid link: wrong prev
      type: "change_membership"
      signer: "basil"
      members:
        writer: ["basil"]
      corruptors:
        prev: -> "deadbeef00000000000000000000000000000000000000000000000000000000"
    ]
  }
}

sessions: [
  loads: [
    error: true
    error_type: "PrevError"
  ]
,
  loads: [
    upto: 1
  ,
    error: true
    error_type: "PrevError"
  ]
]
