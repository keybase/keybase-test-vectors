description: "stub a chain link signed by a revoked device, thus invalid"
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
          merkle_hashmeta: 1500 # after link 1
}

teams: {
  "cabal": {
    links: [
      type: "root"
      signer: "basil"
      members:
        owner: ["basil"]
        admin: ["herb"]
      merkle_hashmetas: [1500] # at these hashmetas, the team chain pointed here
    ,
      # invalid link - signed by a revoked device
      type: "invite"
      signer: "basil"
      invites:
        writer: [ {
          id: "54eafff3400b5bcd8b40bff3d225ab27",
          name: "max+be6ef086a4a5@keyba.se",
          type: "email"
        } ]
    ,
      type: "rotate_key"
      signer: "herb"
    ]
  }
}

sessions: [
  loads: [
    need_admin: true
    error: true
    error_type: "ProofError"
    error_substr: "team link before user key revocation"
  ]
,
  loads: [
    stub: [2] # Stub these seqnos
    n_stubbed: 1
    # signer isn't checked yet for the stubbed link
  ,
    need_admin: true
    error: true
    error_type: "ProofError"
    error_substr: "team link before user key revocation"
    # now the link is inflated and the signer is checked
  ]
,
  loads: [
    upto: 2
    stub: [2] # Stub these seqnos
  ,
    need_admin: true
    error: true
    error_type: "ProofError"
    error_substr: "team link before user key revocation"
  ]
,
  loads: [
    upto: 2
    error: true
    error_type: "ProofError"
    error_substr: "team link before user key revocation"
  ]
,
  loads: [
    upto: 1
  ,
    upto: 2
    stub: [2] # Stub these seqnos
  ,
    need_admin: true
    n_stubbed: 0
    error: true
    error_type: "ProofError"
    error_substr: "team link before user key revocation"
  ]
]
