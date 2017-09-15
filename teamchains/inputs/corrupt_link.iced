description: "corrupted link"

users: {
  "herb": {}
}

teams: {
  "cabal": {
    links: [
      type: "root"
      signer: "herb"
      members:
        owner: ["herb"]
    ,
      type: "rotate_key"
      mangle_payload: true
    ]
  }
}

sessions: [
  loads: [
    error: true
    error_substr: "unmarshaling link payload"
  ]
,
  loads: [
    upto: 1
  ,
    error: true
    error_substr: "unmarshaling link payload"
  ]
]
