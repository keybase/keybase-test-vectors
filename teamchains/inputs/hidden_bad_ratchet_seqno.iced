description: "ratchet from visible to hidden has a bad sequence number (that of a previous link)"

users: {
  "herb": {}
}

teams: {
  cabal : {
    links : [
      { type: "root", members: { owner: ["herb"] } },
      { type : "rotate_key" },
      { type : "rotate_key_hidden" }
      { type : "rotate_key_hidden" }
      {
        type : "rotate_key"
        corruptors : {
          hidden_prev : (p) ->
            {
              link_id : p.link_id
              for_client : seqno : p.for_client.seqno - 1
            }
        }
      }
    ]
  }
}

sessions: [
  {
    loads : [
      {
        error : true
        error_type : "LoaderError"
        error_type_full : "hidden.LoaderError"
        error_substr : "hidden team loader error: link ID at 1 fails to check against ratchet: 2e30006c49bb1ffc8304020d132065451155e004375b731ef3514a395d04bb08 != dca50c46959d90d2000061fb4c2f00e1490897386ef65739164db73a33a58c00"
      }
    ]
  }
]
