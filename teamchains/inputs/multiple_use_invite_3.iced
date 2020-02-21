description: """Join through multiple use invite, leave, and then join through that invite again."""

users: {
  "herb": {}
  "basil": {}
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
          max_uses: 999
        }]
    }, {
      type: "change_membership"
      signer: "herb"
      members:
        writer: ["basil", "lily"]
      used_invites: [
        {
          id: "54eafff3400b5bcd8b40bff3d225ab27"
          uv: "basil"
        },
        {
          id: "54eafff3400b5bcd8b40bff3d225ab27"
          uv: "lily"
        }
      ]
    }, {
      type: "leave",
      signer: "basil"
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
    }]
  }
}

sessions: [
  loads: [
    error: false
  ]
]
