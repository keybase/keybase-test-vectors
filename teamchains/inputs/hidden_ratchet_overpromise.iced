description: "withheld ratchet"

users: {
  "herb": {}
}

teams: {
  cabal : {
    links : [
      { type: "root", members: { owner: ["herb"] } },
      { type : "rotate_key" },
      { type : "rotate_key_hidden" }
      {
        type : "rotate_key"
        corruptors :
          ratchet_seqno : 2
      }
    ]
  }
}

sessions: [
  {
    loads : [
      {
        error : true
        error_type : "HiddenMerkleError"
        error_type_full : "libkb.HiddenMerkleError"
        error_substr : "hidden merkle client error (type 9): Server promised a hidden chain up to 2, but never received; is it withholding?"
      }
    ]
  }
]
