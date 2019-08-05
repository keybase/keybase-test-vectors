description: "sig3: missing reverse sig"

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
        corruptors : no_reverse_sig : true
      },
      { type : "rotate_key" },
    ]
  }
}

sessions: [
  { loads : [{
    error : true
    error_type : "Sig3Error"
    error_type_full : "sig3.Sig3Error"
    error_substr : "sig3 error: rotate key link is missing a reverse sig"
  }]}
]
