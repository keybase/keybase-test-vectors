##### Signed by https://keybase.io/oconnor663
```
-----BEGIN PGP SIGNATURE-----
Version: GnuPG v2

iQEcBAABCAAGBQJVfx5RAAoJEHGHa2itSC0yAqwH/iVlhumiL44NaXFnA0obvf5R
wwDuviC+yXdhL13WVU/klIjUE8VrbegP212DujsTi1XGeWt199Z56BqMY2lGfOki
Akb8a2QrsdteG/w8PrsNFwrYHhfRoXcujhfpLiMbDERAkU8Zy0f3pdQR724sXkVb
uD9lcPmbTQbCVLO6MOqeE7kB78TcautyOX5QCIYdltiC91N/YlYTGFui11EX5zb6
zOpLtTq30ODJWYFkcEQ5WeRE6lxJbTIHKwxpuUANR6lEhGJJy/y8bobFRUhIR7HM
tYnifCFIWqlFS9+atnA0v01Rbi1JDC1aWOUiZ1Z9/aYkRra9H4N8WKxxk4Zfe6w=
=qBkG
-----END PGP SIGNATURE-----

```

<!-- END SIGNATURES -->

### Begin signed statement 

#### Expect

```
size    exec  file                                          contents                                                        
              ./                                                                                                            
4979            chain_tests.json                            af684d7573c6996bfd853b4feb170d107db81d376714baffdb3ad64043b8e961
                chains/                                                                                                     
6051              bad_prev_chain.json                       53f930afd490f739b57d1de35bfb18c0e3dc9e4163a1eb898fce8bdcc98f6842
6051              bad_reverse_signature_chain.json          06940369198a81f6f89bc86dcfd21261b54d397d0ce1cb53b8ec36c522b16ed1
6051              bad_seqno_chain.json                      a411e254fdf9af78d5de5dd8036fb649ce7756fa930edd42a36110c0ab863826
8252              bad_signature_chain.json                  2d0975006ea357e9e0f64557044eed5659ffb037584e2fdc41a1b73c4c237fb6
1607              bad_uid_chain.json                        b9092ab9f79b6359a0df90eb827e18e818377f45cd30f201f76b5f0ada1bbd29
1626              bad_username_chain.json                   58e0e004357f39415b00e1efaf482cc09e8e15cebf3b8b10070f63c85fc9f7d0
1940              empty_chain.json                          3cc56d941437a5120dae27cbba205a02587830dd08fd11dcc6cc05ea18ceb616
47466             example_revokes_chain.json                d693741bfb17fbdd0bd5da570828f9c2d2c0e7676548ad3295c07e1c8b421c64
6001              expired_key_chain.json                    c2733bb8a55f5abb2c712fa2d84740cc36974237a14a3ca388c582570afbb7dd
                  inputs/                                                                                                   
400                 bad_prev_chain.cson                     a315fa2626056325f2e2961ac5fa755c1fd77f1d8e1096d2f146e4972e897be6
309                 bad_seqno_chain.cson                    f0f039ff7f04fc8b6aaa0c05c8e9e777782d136258a04e6fefd8544b3ba1fd0f
206                 bad_uid_chain.cson                      b17a1a982c5a84e3fe1a0a4d1e089887a4d7fb3450571177cdbf8544ff46c1be
191                 bad_username_chain.cson                 0c265be0f1673b86ab7d47a73dac5741b534c4f2412e40fb07c63d24d849896a
2571                example_revokes_chain.cson              1fa7bd01611ffd4e84b6ab6ec551711b136c14dd0a08a55922821fe1044fdf56
302                 expired_key_chain.cson                  1cb759331e75f16e1057882fa15f114d86184b61e1bb84eb340a47308f6ce057
541                 signed_with_revoked_key_chain.cson      62d9b4bcef4c86ebd64419e844d4c848be9f469b55ee78abb3a4beb8cab96ac9
160                 simple_chain.cson                       fc5d18080e832797349566cdf512fa46d28ede8d291d85f855fd9c9eafbdcf8a
285                 two_link_chain_for_manual_hacking.cson  8659c7b477321c83f3102f03afd1b480099016ff7ce4216f13636b563c41eda7
6051              mismatched_ctime_chain.json               47e9d2aa2dc77dd8827fb3109581131ec5b853ff5024b282ac2215c73c0af680
8244              mismatched_fingerprint_chain.json         d49b5b9b19b84b287d7d3baadca941acf102e553acb13dfbc046ac53841f8c21
8280              mismatched_kid_chain.json                 8b238dd8ddf942341a3b1f213f2b0052e1e519a73e37ff3590e06fdbe11bce11
7375              missing_kid_chain.json                    a5959e4f56b209c9d8071b924115747a31c5e5091b5b4c06d5d694bbf1af591a
7391              missing_reverse_kid_chain.json            cf5eb200d78afd221e11674581a574d82225286aff15f9d5f8739554c564bc04
42114             ralph_chain.json                          41b15ea69282b53c4b7cbd40ba2e8c18f49a62a40cf2c90cf154dfcc91107308
12536             signed_with_revoked_key_chain.json        55ac1ea124f7e9bb2473035be106b1718e288527a3583fd711c726185a047051
1607              simple_chain.json                         138d7ac6ba3dff9d210e90101156cc399de6d070eb742eecf10ee2cc9e32b890
1259    x       generate.py                                 be80746bd359bca56f1c846b029254c4e2e814c5e8a8b01d4aa32b51daf27349
                go/                                                                                                         
184267            testvectors.go                            56a0c7b5aff404cf9197667459e4893be4ee86d8ced95555f6eba22f1a26f444
                js/                                                                                                         
1829              main.js                                   ee848bb8c4120283ab0577446a4c8c0155f8297ec95fc38f48104ebf574e2f6d
322             package.json                                73c4ba55f6eb30f17e15bde23ac8945b3148c3cef19abf14a0218ebf6dc546d6
```

#### Ignore

```
/SIGNED.md
```

#### Presets

```
git      # ignore .git and anything as described by .gitignore files
dropbox  # ignore .dropbox-cache and other Dropbox-related files    
kb       # ignore anything as described by .kbignore files          
```

<!-- summarize version = 0.0.9 -->

### End signed statement

<hr>

#### Notes

With keybase you can sign any directory's contents, whether it's a git repo,
source code distribution, or a personal documents folder. It aims to replace the drudgery of:

  1. comparing a zipped file to a detached statement
  2. downloading a public key
  3. confirming it is in fact the author's by reviewing public statements they've made, using it

All in one simple command:

```bash
keybase dir verify
```

There are lots of options, including assertions for automating your checks.

For more info, check out https://keybase.io/docs/command_line/code_signing