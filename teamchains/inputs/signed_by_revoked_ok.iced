description: "chain link signed by a revoked device _before_ it was revoked"

users: {
  "herb":
    keys:
      default:
        revoke:
          seqno: 1
          merkle_hashmeta: 2500 # after link 2
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
      signer: "herb"
      merkle_hashmetas: [2500] # at these hasmetas, the team chain pointed here
    ]
  }
}

sessions: [
  loads: [
    error: false
  ]
  # note: can't load just team link 1 because they revocation of the key references
  # 2500 at which point team link 2 must exist.
]
