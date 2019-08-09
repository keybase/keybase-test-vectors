description: "repeating a PTK gen once on the hidden and once on the visible chain"

users: {
  "herb": {}
}

teams: {
  cabal : {
    links : [
      { type: "root", members: { owner: ["herb"] } }
      { type : "rotate_key" }
      { type : "rotate_key_hidden" }
      { type : "rotate_key", corruptors : ptk_gen : 3 }
    ]
  }
}

sessions: [
  {
    loads : [
      {
        error : true
        error_type_full : "hidden.RepeatPTKGenerationError"
        error_substr : "Repeated PTK Generation found at 3 (clashes a previously-loaded visible rotation)"
      }
    ]
  },{
    loads : [
      {
        error : false
        upto : 2
        hidden_upto : 1
      },{
        error : true
        error_type_full : "hidden.RepeatPTKGenerationError"
        error_substr : "Repeated PTK Generation found at 3 (clashes a previously-loaded hidden rotation)"
      }
    ]
  }
]
