description: "rotate_key missing its per_team_key section"

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
        none: ["basil"]
    ,
      type: "rotate_key"
      corruptors:
        sig_arg: (sig_arg) ->
          delete sig_arg.kms
          sig_arg
    ]
  }
}

sessions: [
  loads: [
    error: true
    error_substr: "missing per-team-key"
  ]
,
  loads: [
    upto: 1
  ,
    upto: 2
  ,
    error: true,
    error_substr: "missing per-team-key"
  ]
,
  loads: [
    upto: 1
  ,
    error: true,
    error_substr: "missing per-team-key"
  ]
]
