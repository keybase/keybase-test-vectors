description: "The seqno skips a number"

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
      # invalid link - the seqno skips ahead
      type: "change_membership"
      seqno: 4
      members:
        admin: ["herb"]
    ]
  }
}

expect: {
  error: true,
  error_substr: "expected seqno"
}
