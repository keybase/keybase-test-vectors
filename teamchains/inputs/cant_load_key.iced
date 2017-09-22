description: "can't load the (signer) key section key"

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
          payload.body.key.kid = "012000000000000000000000000000000000000000000000000000000000000000000a"
    ]
  }
}

sessions: [
  loads: [
    error: true
    error_substr: "in LoadKeyV2: key not found"
  ]
,
  loads: [
    upto: 1
  ,
    error: true,
    error_substr: "in LoadKeyV2: key not found"
  ]
]
