description: "Get a link that is not known to this client but is marked as critical to loading the team"

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
      type: "unsupported_critical"
    ]
  }
}

sessions: [
  loads: [
    error: true
    error_type: "UnsupportedLinkTypeError"
  ]
,
  loads: [
    need_admin: true,
    error: true
    error_type: "UnsupportedLinkTypeError"
  ]
,
  loads: [
    upto: 1
  ,
    error: true
    error_type: "UnsupportedLinkTypeError"
  ]
,
  loads: [
    need_admin: true,
    stub: [2] # Stub these seqnos
    error: true
    error_type: "StubbedError"
  ]
]
