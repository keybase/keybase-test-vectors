description: "Bot user test"

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
        bot: ["basil"]
    ]
  }
}

sessions: [
  loads: [
    error: false
  ]
]
