description: """ Test if it's possible to create a multiple use invite
(infinite uses) that's valid for 1000 years. This is one of the defaults that
Keybase GUI does."""

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
          id: "54eafff3400b5bcd8b40bff3d225ab27"
          name: "YmFzZTY0IGV4YW1wbGUgc3RyCg=="
          type: "seitan_invite_token"
          max_uses: -1 # const for infinity
          etime: 33139407600 # 3020-02-23T23:00:00.000Z
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
    }]
  }
}

sessions: [
  loads: [
    error: false
  ]
]
