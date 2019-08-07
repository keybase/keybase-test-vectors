description: "two rachets from visible to hidden conflict"

users: {
  "herb": {}
}

teams: {
  cabal : {
    links : [
      { type: "root", members: { owner: ["herb"] } },
      { type : "rotate_key" },
      { type : "rotate_key_hidden" },
      { type : "rotate_key" },
      {
        type : "rotate_key"
        corruptors : {
          hidden_prev : (p, state) ->
            buf = Buffer.from(p.link_id, 'hex')
            buf[0] ^= 0x01
            p.link_id = buf.toString('hex')
            p
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
        error_substr : "hidden team loader error: ratchet for seqno 1 contradicts another ratchet"
      }
    ]
  }
]
