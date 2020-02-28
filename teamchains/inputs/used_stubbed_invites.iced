description: """Case when there are multiple use invites that are "used" by
change_membership links, but non-admins load them without knowing about the
invites because of stubbing."""

users: {
  "herb": {}
  "basil": {}
  "rose": {}
  "lily": {}
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
          type: "seitan_invite_token"
          max_uses: 10
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
    }]
  }
}

sessions: [{
  loads: [
    error: false
  ]
}, {
  loads: [{
    error: false
    need_admin: false
    stub: [2] # invite is stubbed
    n_stubbed: 1
  }, {
    error: false
    need_admin: true
    n_stubbed: 0
  }]
}]
