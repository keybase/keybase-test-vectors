description: "request with need_admin when not an admin (client's fault)"

users: {
  "herb": {}
  "basil": {}
}

teams: {
  "cabal": {
    links: [
      type: "root"
      signer: "basil"
      members:
        owner: ["basil"]
        writer: ["herb"]
    ]
  }
}

sessions: [
  loads: [
    error: false
  ]
,
  loads: [
    need_admin: true
    error: true
    error_substr: "is not an admin"
  ]
]
