description: "Returning stubbed links should work"

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
    ,
      type: "invite"
      signer: "basil"
    ,
      type: "invite"
      signer: "basil"
    ]
  }
}

load: {
  # Stub these chain links
  stub: [2, 3]
}

expect: {
  error: false
  n_stubbed: 2
}
