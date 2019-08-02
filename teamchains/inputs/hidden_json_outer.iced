description: "Simplest test"

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
            bundle.o = JSON.stringify { yo : 10 }
      },
      { type : "rotate_key" },
    ]
  }
}

sessions: [
  { loads : [{
    error : true
    error_type : "CorruptInputError"
    error_substr : "illegal base64 data at input byte 0"
  }]}
]
