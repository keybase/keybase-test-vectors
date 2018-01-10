description: "Obsoleting invites"

users: {
  "herb": {}
  "basil": {}
  "rose": {}
  "lily": {}
}

teams: {
  "cabal": {
    links: [
      type: "root"
      members:
        owner: ["herb"]
    ,
      type: "invite"
      signer: "herb"
      invites:
        writer: [ {
          id: "54eafff3400b5bcd8b40bff3d225ab27",
          name: "579651b0d574971040b531b66efbc519%1",
          type: "keybase"
        } ]
    ,
      type: "invite"
      signer: "herb"
      invites:
        reader: [ {
          id: "55eafff3400b5bcd8b40bff3d225ab27",
          # rose%1
          name: "618d663af0f1ec88a5a19defa65a2f19%1",
          type: "keybase"
        }, {
          id: "56eafff3400b5bcd8b40bff3d225ab27",
          # lily%1
          name: "40903c59d19feef1d67c455499304c19%1",
          type: "keybase"
        } ]
    ,
      type: "change_membership"
      signer: "herb"
      members:
        writer: ["basil"]
    ,
      type: "invite"
      signer: "herb"
      invites:
        cancel: ["55eafff3400b5bcd8b40bff3d225ab27"]
    ]
  }
}

sessions: [
  loads: [
    error: false
  ]
]
