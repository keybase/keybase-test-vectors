description: """Multiple use invite, used for invite links."""

users: {
  "herb": {}
  "basil": {}

  "rose": {}
  "lily": {}
  "jasmine": {}
  "rosemary": {}
}

teams: {
  "cabal": {
    links: [{
      type: "root"
      members:
        owner: ["herb", "basil"]
    }, {
      type: "invite"
      signer: "herb"
      invites:
        reader: [{
          id: "54eafff3400b5bcd8b40bff3d225ab27",
          name: "YmFzZTY0IGV4YW1wbGUgc3RyCg==",
          type: "seitan_invite_token"
          max_uses: 5
        }]
    }, {
      type: "invite"
      signer: "basil"
      invites:
        reader: [{
          id: "cc3aa06f02d3409e06f0cb6494ed3027",
          name: "YmFzZTY0IGV4YW1wbGUgc3RyCg==",
          type: "seitan_invite_token"
          max_uses: 999
          etime: 1897806323
        }]
    }, {
      type: "change_membership"
      signer: "herb"
      members:
        writer: ["rose"]
      used_invites: [
        {
          id: "54eafff3400b5bcd8b40bff3d225ab27"
          uv: "rose"
        }
      ]
    }, {
      type: "change_membership"
      signer: "herb"
      members:
        writer: ["lily", "jasmine"]
      used_invites: [
        {
          id: "cc3aa06f02d3409e06f0cb6494ed3027"
          uv: "lily"
        },
        {
          id: "cc3aa06f02d3409e06f0cb6494ed3027"
          uv: "jasmine"
        }
      ]
    }, {
      type: "change_membership"
      signer: "basil"
      members:
        writer: ["rosemary"]
      used_invites: [
        {
          id: "54eafff3400b5bcd8b40bff3d225ab27"
          uv: "rosemary"
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
