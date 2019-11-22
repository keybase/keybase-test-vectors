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
              for_client : hidden_response : uncommitted_seqno : p.for_client.hidden_response.uncommitted_seqno - 1
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
        error_substr : "hidden team loader error: link ID at 1 fails to check against ratchet: 29f4207c6e29cef873baeb1518422defead5ac0b5b3b43002f0bcec75ef43e9e != ff5a460bcb58dac9b4918aa886ec532a9f1e43c879c14137bee1f48e2c21cf37"
      }
    ]
  }
]
