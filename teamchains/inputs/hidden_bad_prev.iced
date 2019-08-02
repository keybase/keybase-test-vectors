description: "Simplest test"

users: {
  "herb": {}
}

teams: {
  cabal : {
    links : [
      { type: "root", members: { owner: ["herb"] } },
      { type : "rotate_key" },
      { type : "rotate_key_hidden" }
      { type : "rotate_key" },
      {
        type : "rotate_key_hidden"
        corruptors :
          sig_arg : (arg) ->
            arg.prev[1] ^= 0x01
            arg
      }
    ]
  }
}

sessions: [
  {
    loads : [
      {
        error : true
        error_type : "SequenceError"
        error_substr : "sig3 sequencing error: prev hash mismatch at 2"
      }
    ]
  },{
    loads : [
      {
        error : false
        hidden_upto : 1
        upto : 2
      },{
        error : true
        error_type : "LoaderError"
        error_type_full : "hidden.LoaderError"
        error_substr : "hidden team loader error: prev mismatch at 2"
      }
    ]
  }
]
