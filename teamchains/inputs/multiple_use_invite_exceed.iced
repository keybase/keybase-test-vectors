description: """Multiple use invite, used for invite links.

Use the invite too many times, should fail in sigchain player.
"""

users: {
  "herb": {}
  "basil": {}
  "rose": {}
  "lily": {}

  "azalea": {}
  "rosemary": {}
  "lotus": {}
}

teams: {
  "cabal": {
    links: [{
      type: "root"
      members:
        owner: ["herb"]
    }, {
      type: "invite"
      signer: "herb"
      invites:
        writer: [{
          id: "54eafff3400b5bcd8b40bff3d225ab27",
          name: "YmFzZTY0IGV4YW1wbGUgc3RyCg==",
          type: "invitelink"
          max_uses: 5
        }]
    }, {
      type: "change_membership"
      signer: "herb"
      members:
        writer: ["basil"]
      used_invites: [
        {
          id: "54eafff3400b5bcd8b40bff3d225ab27"
          uv: "basil"
        }
      ]
    }, {
      type: "change_membership"
      signer: "herb"
      members:
        writer: ["rose", "lily"]
      used_invites: [
        {
          id: "54eafff3400b5bcd8b40bff3d225ab27"
          uv: "rose"
        },
        {
          id: "54eafff3400b5bcd8b40bff3d225ab27"
          uv: "lily"
        }
      ]
    }, {
      type: "change_membership"
      members:
        writer: [ "azalea", "rosemary", "lotus" ]
      used_invites: [
        { id : "54eafff3400b5bcd8b40bff3d225ab27", uv : "azalea" }
        { id : "54eafff3400b5bcd8b40bff3d225ab27", uv : "rosemary" }
        { id : "54eafff3400b5bcd8b40bff3d225ab27", uv : "lotus" }
      ]
    }]
  }
}

sessions: [{
  loads: [{
    error: true
    error_substr: "illegal used_invites: invite 54eafff3400b5bcd8b40bff3d225ab27 is expired after 5 uses"
  }, {
    error: false
    upto: 4
  }]
}]
