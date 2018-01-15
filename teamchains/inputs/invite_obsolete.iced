description: """Obsoleting invites - This chain has 3 invites added in
two links, first one with one invite ("basil%1") and second with two
invites ("rose%1", "lily%1"). Then, basil user is added as a member
using change_membership link, but it does not explicitly complete
invite. Team player should mark "basil%1" invite as obsolete. Last
link cancels "rose%1" invite. After replaying the chain, team should
have two members (herb and basil), and one active invite
("lily%1")."""

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
          # basil%1
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
