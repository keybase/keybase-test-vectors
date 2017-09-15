description: "request a keygen that doesn't exist yet (client's fault)"

users: {
  "herb": {}
}

teams: {
  "cabal": {
    links: [
      type: "root"
      members: owner: ["herb"]
    ]
  }
}

sessions: [
  loads: [
    error: false
  ]
,
  loads: [
    need_keygen: 2
    error: true
    error_substr: "team key generation too low"
  ]
]
