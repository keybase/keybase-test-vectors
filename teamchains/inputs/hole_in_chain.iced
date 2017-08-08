description: "Server omits a link"

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
        owner: ["basil", "herb"]
    ,
      type: "change_membership"
      signer: "basil"
      members:
        owner: ["herb"]
    ,
      type: "change_membership"
      members:
        admin: ["herb"]
    ]
  }
}

load: {
  # Stub these chain links
  omit: [2]
}

expect: {
  error: true,
  error_type: "PrevError"
}
