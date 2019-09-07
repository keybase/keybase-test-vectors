description: "repeating a PTK gen twice on the visible chain"

users: {
  "herb": {}
}

teams: {
  cabal : {
    links : [
      { type: "root", members: { owner: ["herb"] } }
      { type : "rotate_key" }
      { type : "rotate_key", corruptors : ptk_gen : 2 }
    ]
  }
}

sessions: [
  {
    loads : [
      {
        error : true
        error_type_full : "teams.AppendLinkError"
        error_substr : "appending 2->3: PTK clash at generation 2"
      }
    ]
  }
]
