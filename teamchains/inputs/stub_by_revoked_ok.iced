description: "stub a chain link signed by a revoked device _before_ it was revoked"
# Note: this test doesn't actually test inflating.
# Because the loader throws out the whole cache when need_admin turns on.
# This test could be adapted to inflate when subteams are working in these tests.

users: {
  "herb": {}
  "basil":
    keys:
      default:
        revoke:
          seqno: 1
          merkle_hashmeta: 4500 # after link 4
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
      type: "new_subteam"
      signer: "basil"
      subteam:
        name: "cabal.hodlers"
        id: "8a794530c69d9eaac6f7a14433ca7225"
    ,
      type: "change_membership"
      signer: "basil"
      members:
        admin: ["herb"]
    ,
      type: "rotate_key"
      signer: "herb"
      merkle_hashmetas: [4500] # at these hashmetas, the team chain pointed here
    ]
  }
}

sessions: [
  loads: [
    stub: [2] # Stub these seqnos
    n_stubbed: 1
  ,
    need_admin: true
    n_stubbed: 0
  ]
,
  loads: [
    n_stubbed: 0
  ]

  # note: can't partial load the team
  # because the revoked key demands that
  # the team is at >= seqno 3
]
