description: "missing PTK gen; the hidden link skipped one"

users: {
  "herb": {}
}

teams: {
  cabal : {
    links : [
      { type: "root", members: { owner: ["herb"] } }
      { type : "rotate_key" }
      { type : "rotate_key_hidden" }
      { type : "rotate_key" }
      { type : "rotate_key_hidden", corruptors : ptk_gen : 6 }
      { type : "rotate_key_hidden", corruptors : ptk_gen : 7 }
    ]
  }
}

sessions: [
  {
    loads : [
      {
        error : true
        error_substr : "loading team secrets: per-team-key not found for generation 5"
      }
    ]
  }
]
