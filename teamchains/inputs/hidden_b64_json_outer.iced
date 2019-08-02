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
            bundle.o = Buffer.from(JSON.stringify({ yo : 10 }), "utf8").toString('base64')
      },
      { type : "rotate_key" },
    ]
  }
}

sessions: [
  { loads : [{
    error : true
    error_type : "ParseError"
    error_substr : "need an encoded msgpack array"
  }]}
]
