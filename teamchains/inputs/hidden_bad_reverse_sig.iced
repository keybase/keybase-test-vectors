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
          reverse_sig_inputs : ({inner, outer}) ->
            inner.b.k[0].t++
      },
      { type : "rotate_key" },
    ]
  }
}

sessions: [
  { loads : [{
    error : true
    error_type : "Sig3Error"
    error_substr : "sig3 error: bad reverse signature: sig3 error: signature verification failed"
  }]}
]
