description: "simple team with 2 generations of keys"

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
    ]
  }
}

sessions: [
  loads: [
    error: false,
    then_get_key: 1
  ]
,
  loads: [
    upto: 1
    then_get_key: 1
  ,
    upto: 3
    then_get_key: 2
  ,
    error: false,
    then_get_key: 1
  ]
,
  loads: [
    upto: 1
    then_get_key: 1
  ,
    error: false,
    then_get_key: 2
  ]
]
