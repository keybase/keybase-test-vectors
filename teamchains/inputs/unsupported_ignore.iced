description: "Get a link that is not known to this client but is marked as ingoreable"

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
      type: "unsupported_ignore"
    ]
  }
}

sessions: [
  loads: [
    error: false
  ]
,
  loads: [
    need_admin: true,
    error: false
  ]
,
  loads: [
    upto: 1
  ,
    error: false
  ]
,
  loads: [
    stub: [2] # Stub these seqnos
    error: false
  ]
,
  loads: [
    need_admin: true,
    stub: [2] # Stub these seqnos
    error: true,
    error_type: "StubbedError"
  ]
]
