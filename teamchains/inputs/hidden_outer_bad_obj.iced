description: "sig3 bad outer formatting"

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
            bundle.o = "lAECAwQ=" # (pack [1,2,3,4]).toString('base64')
      },
      { type : "rotate_key" },
    ]
  }
}

sessions: [
  { loads : [{
    error : true
    error_substr : "msgpack decode error [pos 4]: only encoded map or array can be decoded into a slice (0)"
  }]}
]
