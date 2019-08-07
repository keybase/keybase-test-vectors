description: "the ratchet blinding/unblinding sent down from the server is bad"

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
          generated_ratchet : (r) ->
            r.ratchet[1] ^= 0x1
      }
    ]
  }
}

load_failure :
  error : true
  error_substr : "hidden team ratchet error: blinding check failed cbaf3f1776b85d162906f74515de1f79f7a17a3341162decff3607e721350db2 v cbae3f1776b85d162906f74515de1f79f7a17a3341162decff3607e721350db2"
  error_type_full : "hidden.RatchetError"
