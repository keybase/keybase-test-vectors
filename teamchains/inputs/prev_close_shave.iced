description: "server responds with minimum required prevs"

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
    ,
      type: "rotate_key"
    ,
      type: "rotate_key"
    ]
  }
}

sessions: [
  loads: [
    error: false
    then_get_key: 4
  ]
,
  loads: [
    upto: 1
  ,
    # omit prevs that contain the secret
    # for generations <= to this number
    omit_prevs: 1
    then_get_key: 1
  ]
,
  loads: [
    upto: 2
  ,
    omit_prevs: 2
    then_get_key: 2
  ]
,
  loads: [
    upto: 1
  ,
    upto: 2
    omit_prevs: 1
  ,
    upto: 3
    omit_prevs: 2
    then_get_key: 2
  ,
    omit_prevs: 2
    then_get_key: 4
  ]
]
