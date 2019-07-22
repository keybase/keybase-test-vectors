description: "Restricted bot user test"

users: {
  "herb": {}
  "basil": {}
}

teams: {
  "cabal": {
    links: [
      type: "root"
      members:
        owner: ["herb"]
        restricted_bot: ["basil"]
    ]
  }
}

sessions: [
  loads: [
    error: false
  ]
]
