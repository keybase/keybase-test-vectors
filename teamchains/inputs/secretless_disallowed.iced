description: "server returns 'subteam-reader' response for a root team load"

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
    subteam_reader: true
    error: true
    error_substr: "unexpected subteam reader result"
  ]
]
