description: "Stub a link for an admin that will balk at the stub"

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
        admin: ["herb"]
    ,
      type: "invite"
      signer: "basil"
      invites:
        admin: [ {
          id: "54eafff3400b5bcd8b40bff3d225ab27",
          name: "max+be6ef086a4a5@keyba.se",
          type: "email"
        } ]

    ]
  }
}

load: {
  need_admin: true
  # Stub these chain links (by seqno)
  stub: [2]
}

expect: {
  error: true
  error_type: "StubbedError"
  error_substr: "seqno 2"
}
