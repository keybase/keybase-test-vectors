description: """Admin got tricked into using old invite. Note that
this would not happen right now, because of chains are replayed.
Invite that was obsoleted would not end up in ActiveInvites, so if SBS
message arrived, admin would ignore such message becuase the obsoleted
inviteId would be invalid for them.

But: We want this chain to still be valid, because there may be teams
in the wild that had similar sigchains, and they shouldn't break. """

users: {
  "herb": {}
  "basil": {}
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
      type: "change_membership"
      signer: "herb"
      members:
        writer: ["basil"]
    ,
      # Basil leaves
      type: "leave"
      signer: "basil"
    ,
      # Assume rogue server tricked Herb with SBS message with invite
      # 54eafff3400b5bcd8b40bff3d225ab27. The invite was not explicitly
      # cleared nor completed.
      type: "change_membership"
      signer: "herb"
      completed_invites:
        # Even on new clients, where this invite would have been implicitely
        # completed, team player still should fail on this link.
        "54eafff3400b5bcd8b40bff3d225ab27": "basil"
      members:
        writer: ["basil"]
    ]
  }
}

sessions: [
  loads: [
    error: false
  ]
]
