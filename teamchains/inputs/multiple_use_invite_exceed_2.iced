description: """Multiple use invite, used for invite links.

Use the invite too many times, should fail in sigchain player.

Loading the links up to 5th link (3rd change_membership) should work - adding
"lotus" brings the invite use count to 5, which is the max_uses. Then loading
next link - adding "rosemary" - triggers an error, because 6th add for this
invite id is illegal.

"""

users: {
  "herb": {}
  "basil": {}
  "rose": {}
  "lily": {}
  "azalea": {}
  "lotus": {}
  "rosemary": {}
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
        { id: "54eafff3400b5bcd8b40bff3d225ab27", uv: "basil" }
      ]
    }, {
      type: "change_membership"
      signer: "herb"
      members:
        writer: [ "rose", "lily", "azalea" ]
      used_invites: [
        { id: "54eafff3400b5bcd8b40bff3d225ab27", uv: "rose" }
        { id: "54eafff3400b5bcd8b40bff3d225ab27", uv: "lily" }
        { id : "54eafff3400b5bcd8b40bff3d225ab27", uv : "azalea" }
      ]
    }, {
      type: "change_membership"
      members:
        writer: [ "lotus" ]
      used_invites: [
        { id : "54eafff3400b5bcd8b40bff3d225ab27", uv : "lotus" }
      ]
    }, {
      type: "change_membership"
      members:
        writer: [ "rosemary" ]
      used_invites: [
        { id : "54eafff3400b5bcd8b40bff3d225ab27", uv : "rosemary" }
      ]
    }]
  }
}

sessions: [{
  loads: [{
    error: false
    upto: 4
  }, {
    error: false
    upto: 5
  }, {
    error: true
    error_substr: "illegal used_invites: invite 54eafff3400b5bcd8b40bff3d225ab27 is expired after 5 uses"
  }]
}]
