description: "server withholds the latest link"

users: {
  "herb": {}
}

teams: {
  "cabal": {
    links: [
      type: "root"
      members: owner: ["herb"]
    ,
      type: "rotate_key"
    ]
  }
}

sessions: [
  loads: [
    error: false
  ]
,
  loads: [
    upto: 1
  ,
    omit: [2] # omit these seqnos
    error: true
    error_substr: "wrong sigchain link ID"
  ]
]
