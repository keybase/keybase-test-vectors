description: "Link signed by a user not in the team, not claiming to be an admin"

users: {
  "herb": {}
  "basil": {}
}

teams: {
  "cabal": {
    links: [
      type: "root"
      signer: "herb"
      members:
        owner: ["herb"]
    ,
      # invalid link
      # basil tries to leave, but isn't in the team
      type: "leave"
      signer: "basil"
    ]
  }
}

expect: {
  error: true,
  error_type: "AdminPermissionError"
}
