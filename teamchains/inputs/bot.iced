description: "Bot user test"

users: {
  "herb": {}
}

teams: {
  "cabal": {
    links: [
      type: "root"
      members:
        bot: ["herb"]
    ]
  }
}

sessions: [
  loads: [
    error: false
  ]
]
