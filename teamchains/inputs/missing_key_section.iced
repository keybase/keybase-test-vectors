description: "missing the (signer) key section on the link"

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
          delete payload.body.key
    ]
  }
}

sessions: [
  loads: [
    error: true
    error_type: "NoUIDError"
  ]
,
  loads: [
    upto: 1
  ,
    error: true,
    error_type: "NoUIDError"
  ]
]
