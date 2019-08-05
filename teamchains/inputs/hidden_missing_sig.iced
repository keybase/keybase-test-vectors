description: "sig3 missing sig"

users: {
  "herb": {}
}

teams: {
  cabal : {
    links : [
      { type: "root", members: { owner: ["herb"] } },
      { type : "rotate_key" },
      {
        type : "rotate_key_hidden"
        corruptors :
          sig3_bundle : ({bundle}) ->
            delete bundle.s
      },
      { type : "rotate_key" },
    ]
  }
}

sessions: [
  { loads : [{
    error : true
    error_type : "ParseError"
    error_type_full : "sig3.ParseError"
    error_substr : "need a sig and an inner, or neither, but not one without the other (sig: false, inner: true)"
  }]}
]
