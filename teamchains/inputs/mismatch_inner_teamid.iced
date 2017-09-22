description: "wrong team id in inner link"

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
        payload: (payload) ->
          payload.body.team.id = "d9000000000000000000000000000024"
    ]
  }
}

sessions: [
  loads: [
    error: true
    error_substr: "wrong team ID"
  ]
,
  loads: [
    upto: 1
  ,
    error: true,
    error_substr: "wrong team ID"
  ]
]
