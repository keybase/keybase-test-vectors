description: "Simplest test"

users: {
  "herb": {}
}

teams: {
  cabal : {
    links : [
      { type: "root", members: { owner: ["herb"] } },
      { type : "rotate_key" },
      { type : "rotate_key_hidden" }
      {
        type : "rotate_key"
        corruptors :
          hidden_prev : (p) ->
            buf = Buffer.from(p.link_id, 'hex')
            buf[0] ^= 0x01
            p.link_id = buf.toString('hex')
            p
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
        error_substr : "hidden team loader error: link ID at 1 fails to check against ratchet: dda50c46959d90d2000061fb4c2f00e1490897386ef65739164db73a33a58c00 != dca50c46959d90d2000061fb4c2f00e1490897386ef65739164db73a33a58c00"
      }
    ]
  }
]
