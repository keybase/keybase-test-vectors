description: "wrong seqno type on inner link"

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
          payload.seq_type_for_testing = payload.seq_type
          payload.seq_type = 999
    ]
  }
}

sessions: [
  loads: [
    error: true
    error_type: "SigchainV2MismatchedFieldError"
  ]
,
  loads: [
    upto: 1
  ,
    error: true,
    error_type: "SigchainV2MismatchedFieldError"
  ]
]
