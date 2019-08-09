description: "links from the future load without an issue if they are marked 'ignore if unsupported'"

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
          dont_save_key : true
          sig3_patch_outer : (outer) ->
            outer.link_type = 1000 # link Type
            outer.ignore_if_unsupported = true # ignore if unsupported
            outer
      },
    ]
  }
}

sessions: [
  { loads : [{
    error : false
  }]}
]
