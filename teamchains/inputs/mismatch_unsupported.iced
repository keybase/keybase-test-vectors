description: "The inner and outer values of ignore_if_unsupported are mismatched"

users: {
  "herb": {}
}

teams: {
  "cabal": {
    links: [
      type: "root"
      members:
        owner: ["herb"]
    ,
      type: "unsupported"
      ignore_if_unsupported: false
      corruptors:
        payload: (payload) ->
          payload.ignore_if_unsupported_for_testing = true
    ]
  }
}

sessions: [
  loads: [
    error: true
    error_substr: "ignore_if_unsupported"
  ]
]
