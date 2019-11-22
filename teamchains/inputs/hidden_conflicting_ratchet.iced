description: "ratchets from hidden to visible contradicts the secret links downloaded"

users: {
  "herb": {}
}

teams: {
  cabal : {
    links : [
      { type: "root", members: { owner: ["herb"] } },
      { type : "rotate_key" },
      { type : "rotate_key_hidden" },
      {
        type : "rotate_key"
        corruptors :
          hidden_prev : (p, state) ->
            state.p = p
            p
      },
      { type : "rotate_key_hidden" },
      {
        type : "rotate_key"
        corruptors :
          hidden_prev : (p, state) ->
            {
              link_id : state.p.link_id
              for_client : hidden_response : uncommitted_seqno : p.for_client.hidden_response.uncommitted_seqno
            }
      },
      { type : "rotate_key_hidden" },
      { type : "rotate_key" },
      { type : "rotate_key_hidden" },
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
        error_substr : "hidden team loader error: link ID at 2 fails to check against ratchet: ff5a460bcb58dac9b4918aa886ec532a9f1e43c879c14137bee1f48e2c21cf37 != ed52ca4fbe17ceebc8988d93190f73a4ddb5d431328dfccd72f49b3e0e0da4c5"
      }
    ]
  }
]
