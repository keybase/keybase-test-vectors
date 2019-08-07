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
              for_client : seqno : p.for_client.seqno
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
        error_substr : "hidden team loader error: link ID at 2 fails to check against ratchet: dca50c46959d90d2000061fb4c2f00e1490897386ef65739164db73a33a58c00 != e362136fee3c035151f6a4e62e521f853adb4308d472b959e5e5262891538b6e"
      }
    ]
  }
]
