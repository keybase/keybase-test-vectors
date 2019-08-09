description: "missing hidden ratchet blinding key on loader output"

users: {
  "herb": {}
}

teams: {
  cabal : {
    links : [
      { type: "root", members: { owner: ["herb"] } }
      { type : "rotate_key" }
      { type : "rotate_key_hidden" }
      {
        type : "rotate_key"
        corruptors : drop_ratchet_blinding_key : true
      }
      { type : "rotate_key_hidden"}
    ]
  }
}

sessions: [
  {
    loads : [
      {
        error : true
        error_type_full : "hidden.LoaderError"
        error_substr : "hidden team loader error: missing unblind for ratchet"
      }
    ]
  }
]
