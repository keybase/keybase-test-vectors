description: "chain link signed by a revoked device, thus invalid"

users: {
  "herb":
    keys:
      default:
        revoke:
          seqno: 1
          merkle_hashmeta: 1500 # between team chain 1 and 2
}

teams: {
  "cabal": {
    links: [
      type: "root"
      signer: "herb"
      members:
        owner: ["herb"]
      merkle_hashmetas: [1500] # at these hasmetas, the team chain pointed here
    ,
      # invalid link - signed by a revoked device
      type: "rotate_key"
      signer: "herb"
    ]
  }
}

sessions: [
  loads: [
    error: true
    error_type: "ProofError"
    error_substr: "team link before user key revocation"
  ]
,
  loads: [
    upto: 1
  ,
    error: true
    error_type: "ProofError"
    error_substr: "team link before user key revocation"
  ]
]
