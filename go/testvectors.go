package testvectors

const ChainTests = `
{
  "error_types": [
    "BAD_LINK_FORMAT",
    "NONEXISTENT_KID",
    "VERIFY_FAILED",
    "REVERSE_SIG_VERIFY_FAILED",
    "KID_MISMATCH",
    "FINGERPRINT_MISMATCH",
    "CTIME_MISMATCH",
    "INVALID_SIBKEY",
    "EXPIRED_SIBKEY",
    "WRONG_UID",
    "WRONG_USERNAME",
    "WRONG_SEQNO",
    "WRONG_PREV"
  ],
  "tests": {
    "test_ralph_sig_chain": {
      "_comment": "Ralph is a test user I created by hand on my local server. I fetched his sigs and keys from the API, and then massaged them into our input format. This test is mainly to make sure that the generated chains we're using in other tests bear some relationship to reality.  - Jack",
      "_comment2": "The eldest key for this test is not the first in the list, it's the 2nd (index 1).",
      "_comment3": "TODO: Use labels instead of indices.",
      "input": "ralph_chain.json",
      "len": 5,
      "sibkeys": 3,
      "subkeys": 2,
      "eldest": "second_eldest"
    },

    "test_simple_chain": {
      "_comment": "Test a simple chain, just one link.",
      "input": "simple_chain.json",
      "len": 1,
      "sibkeys": 1,
      "subkeys": 0
    },

    "test_error_unknown_key": {
      "_comment": "Check the case where a signing kid is simply missing from the list of available keys (as opposed to invalid for some other reason, like having been revoked).",
      "input": "missing_kid_chain.json",
      "eldest": "e",
      "err_type": "NONEXISTENT_KID"
    },

    "test_error_unknown_reverse_sig_key": {
      "_comment": "As above, but for a reverse sig.",
      "input": "missing_reverse_kid_chain.json",
      "eldest": "e",
      "err_type": "NONEXISTENT_KID"
    },

    "test_error_bad_signature": {
      "_comment": "Change some bytes from the valid signature, and confirm it gets rejected.",
      "input": "bad_signature_chain.json",
      "err_type": "VERIFY_FAILED"
    },

    "test_error_bad_reverse_signature": {
      "_comment": "Change some bytes from the valid reverse signature, and confirm it gets rejected.",
      "input": "bad_reverse_signature_chain.json",
      "err_type": "REVERSE_SIG_VERIFY_FAILED"
    },

    "test_error_mismatched_ctime": {
      "_comment": "We need to use the server-provided ctime to unbox a signature (PGP key expiry is checked at the signature level, although NaCl expiry is checked as we replay the chain). We always need to check back after unboxing to make sure the internal ctime matches what the server said. This test exercises that check.",
      "input": "mismatched_ctime_chain.json",
      "err_type": "CTIME_MISMATCH"
    },

    "test_error_mismatched_kid": {
      "_comment": "We need to use the server-provided KID to unbox a signature. We always need to check back after unboxing to make sure the internal KID matches the one we actually used. This test exercises that check. NOTE: I generated this chain by hacking some code into kbpgp to modify the payload right before it was signed.",
      "input": "mismatched_kid_chain.json",
      "err_type": "KID_MISMATCH"
    },

    "test_error_mismatched_fingerprint": {
      "_comment": "We don't use fingerprints in unboxing, but nonetheless we want to make sure that if a chain link claims to have been signed by a given fingerprint, that does in fact correspond to the KID of the PGP key that signed it. NOTE: I generated this chain by hacking some code into kbpgp to modify the payload right before it was signed.",
      "input": "mismatched_fingerprint_chain.json",
      "err_type": "FINGERPRINT_MISMATCH"
    },

    "test_revokes": {
      "_comment": "The chain is length 10, but after 2 sig revokes it should be length 8. Likewise, 6 keys are delegated, but after 2 sig revokes and 2 key revokes it should be down to 2 keys.",
      "input": "example_revokes_chain.json",
      "len": 13,
      "sibkeys": 2,
      "subkeys": 1
    },

    "test_error_revoked_key": {
      "_comment": "Try signing a link with a key that was previously revoked.",
      "input": "signed_with_revoked_key_chain.json",
      "err_type": "INVALID_SIBKEY"
    },

    "test_error_expired_key": {
      "_comment": "Try signing a link with a key that has expired.",
      "input": "expired_key_chain.json",
      "err_type": "EXPIRED_SIBKEY"
    },

    "test_error_bad_uid": {
      "input": "bad_uid_chain.json",
      "err_type": "WRONG_UID"
    },

    "test_error_bad_username": {
      "input": "bad_username_chain.json",
      "err_type": "WRONG_USERNAME"
    },

    "test_error_bad_prev": {
      "input": "bad_prev_chain.json",
      "err_type": "WRONG_PREV"
    },

    "test_error_bad_seqno": {
      "input": "bad_seqno_chain.json",
      "err_type": "WRONG_SEQNO"
    }
  }
}
`

var ChainTestInputs = map[string]string{
	"bad_prev_chain.json": `{
    "chain": [
        {
            "seqno": 1,
            "prev": null,
            "sig": "g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEgtAfp8c2vIsq6o/a4fzGays16KagtWSwxbObC7mU2UngKp3BheWxvYWTFASp7ImJvZHkiOnsia2V5Ijp7Imhvc3QiOiJrZXliYXNlLmlvIiwia2lkIjoiMDEyMGI0MDdlOWYxY2RhZjIyY2FiYWEzZjZiODdmMzE5YWNhY2Q3YTI5YTgyZDU5MmMzMTZjZTZjMmVlNjUzNjUyNzgwYSIsInVpZCI6Ijc0YzM4Y2Y3Y2ViOTQ3ZjU2MzIwNDVkOGNhNWQ0ODE5IiwidXNlcm5hbWUiOiJtYXgzMiJ9LCJ0eXBlIjoiZWxkZXN0IiwidmVyc2lvbiI6MX0sImN0aW1lIjoxNDMyMjI1OTA0LCJleHBpcmVfaW4iOjEwMDAwMDAwLCJwcmV2IjpudWxsLCJzZXFfdHlwZSI6MSwic2Vxbm8iOjEsInRhZyI6InNpZ25hdHVyZSJ9o3NpZ8RALK1iab8sGnnb4waKvKzJAO2kpRvTyaCy7u6W8ZI2+OALN0SmeZZ4ZZk9A+qDWNWFaw7LRDg7/wafrn4sVjwlDKhzaWdfdHlwZSCjdGFnzQICp3ZlcnNpb24B",
            "payload_hash": "df5d9f31e681f921d15e1dc98dd875f3f299dcebdf5f4267827416a2117cb22d",
            "sig_id": "3dc0ed1000573ccde355fd4004274c2aa45aed986df5fcbc187ce1ceffdfcb560f",
            "payload_json": "{\"body\":{\"key\":{\"host\":\"keybase.io\",\"kid\":\"0120b407e9f1cdaf22cabaa3f6b87f319acacd7a29a82d592c316ce6c2ee653652780a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"type\":\"eldest\",\"version\":1},\"ctime\":1432225904,\"expire_in\":10000000,\"prev\":null,\"seq_type\":1,\"seqno\":1,\"tag\":\"signature\"}",
            "kid": "0120b407e9f1cdaf22cabaa3f6b87f319acacd7a29a82d592c316ce6c2ee653652780a",
            "ctime": 1432225904
        },
        {
            "seqno": 2,
            "prev": "deadbeef00000000000000000000000000000000000000000000000000000000",
            "sig": "g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEgtAfp8c2vIsq6o/a4fzGays16KagtWSwxbObC7mU2UngKp3BheWxvYWTFBfl7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTIwYjQwN2U5ZjFjZGFmMjJjYWJhYTNmNmI4N2YzMTlhY2FjZDdhMjlhODJkNTkyYzMxNmNlNmMyZWU2NTM2NTI3ODBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwYjQwN2U5ZjFjZGFmMjJjYWJhYTNmNmI4N2YzMTlhY2FjZDdhMjlhODJkNTkyYzMxNmNlNmMyZWU2NTM2NTI3ODBhIiwidWlkIjoiNzRjMzhjZjdjZWI5NDdmNTYzMjA0NWQ4Y2E1ZDQ4MTkiLCJ1c2VybmFtZSI6Im1heDMyIn0sInNpYmtleSI6eyJraWQiOiIwMTIwMTI2ZTRjYjQwMTI3MTAyMjQ3YjQ2MGFjNmU0ODcyYzM1YjUzMDc1YWYwNTk4Y2NiMDQwMzRiOTYxY2Y3NDBjZDBhIiwicmV2ZXJzZV9zaWciOiJnNlJpYjJSNWhxaGtaWFJoWTJobFpNT3BhR0Z6YUY5MGVYQmxDcU5yWlhuRUl3RWdFbTVNdEFFbkVDSkh0R0NzYmtoeXcxdFRCMXJ3V1l6TEJBTkxsaHozUU0wS3AzQmhlV3h2WVdURkFpdDdJbUp2WkhraU9uc2lhMlY1SWpwN0ltVnNaR1Z6ZEY5cmFXUWlPaUl3TVRJd1lqUXdOMlU1WmpGalpHRm1NakpqWVdKaFlUTm1ObUk0TjJZek1UbGhZMkZqWkRkaE1qbGhPREprTlRreVl6TXhObU5sTm1NeVpXVTJOVE0yTlRJM09EQmhJaXdpYUc5emRDSTZJbXRsZVdKaGMyVXVhVzhpTENKcmFXUWlPaUl3TVRJd1lqUXdOMlU1WmpGalpHRm1NakpqWVdKaFlUTm1ObUk0TjJZek1UbGhZMkZqWkRkaE1qbGhPREprTlRreVl6TXhObU5sTm1NeVpXVTJOVE0yTlRJM09EQmhJaXdpZFdsa0lqb2lOelJqTXpoalpqZGpaV0k1TkRkbU5UWXpNakEwTldRNFkyRTFaRFE0TVRraUxDSjFjMlZ5Ym1GdFpTSTZJbTFoZURNeUluMHNJbk5wWW10bGVTSTZleUpyYVdRaU9pSXdNVEl3TVRJMlpUUmpZalF3TVRJM01UQXlNalEzWWpRMk1HRmpObVUwT0RjeVl6TTFZalV6TURjMVlXWXdOVGs0WTJOaU1EUXdNelJpT1RZeFkyWTNOREJqWkRCaElpd2ljbVYyWlhKelpWOXphV2NpT201MWJHeDlMQ0owZVhCbElqb2ljMmxpYTJWNUlpd2lkbVZ5YzJsdmJpSTZNWDBzSW1OMGFXMWxJam94TkRNeU1qSTJNREEwTENKbGVIQnBjbVZmYVc0aU9qRXdNREF3TURBd0xDSndjbVYySWpvaVpHVmhaR0psWldZd01EQXdNREF3TURBd01EQXdNREF3TURBd01EQXdNREF3TURBd01EQXdNREF3TURBd01EQXdNREF3TURBd01EQXdNREF3TURBd01DSXNJbk5sY1Y5MGVYQmxJam94TENKelpYRnVieUk2TWl3aWRHRm5Jam9pYzJsbmJtRjBkWEpsSW4yamMybG54RUI3TWtDVlJmejFyUXc5UXlmZ3VxYm5jZHdCdmUxWUdUY29wc3hmaWJ3alJqOThJbzFuNGtBbzhpNXBCUDRXazdIaVRTc3I3REpORjZHUnRVaUpMTWtDcUhOcFoxOTBlWEJsSUtOMFlXZk5BZ0tuZG1WeWMybHZiZ0U9In0sInR5cGUiOiJzaWJrZXkiLCJ2ZXJzaW9uIjoxfSwiY3RpbWUiOjE0MzIyMjYwMDQsImV4cGlyZV9pbiI6MTAwMDAwMDAsInByZXYiOiJkZWFkYmVlZjAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwIiwic2VxX3R5cGUiOjEsInNlcW5vIjoyLCJ0YWciOiJzaWduYXR1cmUifaNzaWfEQG/DK6SGvEXIp24c8aSxWcQfx0MNtDRdo6lbK6YnZ9fZ+ic60bW9QoUzCUHD7eGtltTZnsC0lb8rqb2q+nBpiwWoc2lnX3R5cGUgo3RhZ80CAqd2ZXJzaW9uAQ==",
            "payload_hash": "bca91e71714a849c255c531c60bc6e4a69c33782edeba8afcab9fc22b286171e",
            "sig_id": "10a9ec422404d3b0b062e56c809db16e3301bd01deb4d54ee77e5895c29a79660f",
            "payload_json": "{\"body\":{\"key\":{\"eldest_kid\":\"0120b407e9f1cdaf22cabaa3f6b87f319acacd7a29a82d592c316ce6c2ee653652780a\",\"host\":\"keybase.io\",\"kid\":\"0120b407e9f1cdaf22cabaa3f6b87f319acacd7a29a82d592c316ce6c2ee653652780a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"sibkey\":{\"kid\":\"0120126e4cb40127102247b460ac6e4872c35b53075af0598ccb04034b961cf740cd0a\",\"reverse_sig\":\"g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEgEm5MtAEnECJHtGCsbkhyw1tTB1rwWYzLBANLlhz3QM0Kp3BheWxvYWTFAit7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTIwYjQwN2U5ZjFjZGFmMjJjYWJhYTNmNmI4N2YzMTlhY2FjZDdhMjlhODJkNTkyYzMxNmNlNmMyZWU2NTM2NTI3ODBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwYjQwN2U5ZjFjZGFmMjJjYWJhYTNmNmI4N2YzMTlhY2FjZDdhMjlhODJkNTkyYzMxNmNlNmMyZWU2NTM2NTI3ODBhIiwidWlkIjoiNzRjMzhjZjdjZWI5NDdmNTYzMjA0NWQ4Y2E1ZDQ4MTkiLCJ1c2VybmFtZSI6Im1heDMyIn0sInNpYmtleSI6eyJraWQiOiIwMTIwMTI2ZTRjYjQwMTI3MTAyMjQ3YjQ2MGFjNmU0ODcyYzM1YjUzMDc1YWYwNTk4Y2NiMDQwMzRiOTYxY2Y3NDBjZDBhIiwicmV2ZXJzZV9zaWciOm51bGx9LCJ0eXBlIjoic2lia2V5IiwidmVyc2lvbiI6MX0sImN0aW1lIjoxNDMyMjI2MDA0LCJleHBpcmVfaW4iOjEwMDAwMDAwLCJwcmV2IjoiZGVhZGJlZWYwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMCIsInNlcV90eXBlIjoxLCJzZXFubyI6MiwidGFnIjoic2lnbmF0dXJlIn2jc2lnxEB7MkCVRfz1rQw9QyfguqbncdwBve1YGTcopsxfibwjRj98Io1n4kAo8i5pBP4Wk7HiTSsr7DJNF6GRtUiJLMkCqHNpZ190eXBlIKN0YWfNAgKndmVyc2lvbgE=\"},\"type\":\"sibkey\",\"version\":1},\"ctime\":1432226004,\"expire_in\":10000000,\"prev\":\"deadbeef00000000000000000000000000000000000000000000000000000000\",\"seq_type\":1,\"seqno\":2,\"tag\":\"signature\"}",
            "kid": "0120b407e9f1cdaf22cabaa3f6b87f319acacd7a29a82d592c316ce6c2ee653652780a",
            "ctime": 1432226004
        }
    ],
    "keys": [
        "0120b407e9f1cdaf22cabaa3f6b87f319acacd7a29a82d592c316ce6c2ee653652780a",
        "0120126e4cb40127102247b460ac6e4872c35b53075af0598ccb04034b961cf740cd0a"
    ],
    "uid": "74c38cf7ceb947f5632045d8ca5d4819",
    "username": "max32"
}
`,
	"bad_reverse_signature_chain.json": `{
    "chain": [
        {
            "seqno": 1,
            "prev": null,
            "sig": "g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEgd1JXIa21+kSY41/Nx2jLADJcFxGLZ0Rgxug6xaFH67EKp3BheWxvYWTFASp7ImJvZHkiOnsia2V5Ijp7Imhvc3QiOiJrZXliYXNlLmlvIiwia2lkIjoiMDEyMDc3NTI1NzIxYWRiNWZhNDQ5OGUzNWZjZGM3NjhjYjAwMzI1YzE3MTE4YjY3NDQ2MGM2ZTgzYWM1YTE0N2ViYjEwYSIsInVpZCI6Ijc0YzM4Y2Y3Y2ViOTQ3ZjU2MzIwNDVkOGNhNWQ0ODE5IiwidXNlcm5hbWUiOiJtYXgzMiJ9LCJ0eXBlIjoiZWxkZXN0IiwidmVyc2lvbiI6MX0sImN0aW1lIjoxNDMyMTQ0MjQ1LCJleHBpcmVfaW4iOjEwMDAwMDAwLCJwcmV2IjpudWxsLCJzZXFfdHlwZSI6MSwic2Vxbm8iOjEsInRhZyI6InNpZ25hdHVyZSJ9o3NpZ8RAz1C4EVTGxElMcIHOJqRFTmqA0XeFS8oAweX4Bih+6yfzi6sqiOefKBg5AvG3gXeWnHdsSCzJIun8Anwuhf+zAqhzaWdfdHlwZSCjdGFnzQICp3ZlcnNpb24B",
            "payload_hash": "2eddd3fa83f7f34808526b28bb887a95c09dc81f833c88b234a4be2311dd89f4",
            "sig_id": "7c7ef1f35cf56e63b62f42d1afee7682022bed0c58981f9ca4f8601929cb923c0f",
            "payload_json": "{\"body\":{\"key\":{\"host\":\"keybase.io\",\"kid\":\"012077525721adb5fa4498e35fcdc768cb00325c17118b674460c6e83ac5a147ebb10a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"type\":\"eldest\",\"version\":1},\"ctime\":1432144245,\"expire_in\":10000000,\"prev\":null,\"seq_type\":1,\"seqno\":1,\"tag\":\"signature\"}",
            "kid": "012077525721adb5fa4498e35fcdc768cb00325c17118b674460c6e83ac5a147ebb10a",
            "ctime": 1432144245
        },
        {
            "seqno": 2,
            "prev": "2eddd3fa83f7f34808526b28bb887a95c09dc81f833c88b234a4be2311dd89f4",
            "sig": "g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEgd1JXIa21+kSY41/Nx2jLADJcFxGLZ0Rgxug6xaFH67EKp3BheWxvYWTFBfl7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTIwNzc1MjU3MjFhZGI1ZmE0NDk4ZTM1ZmNkYzc2OGNiMDAzMjVjMTcxMThiNjc0NDYwYzZlODNhYzVhMTQ3ZWJiMTBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwNzc1MjU3MjFhZGI1ZmE0NDk4ZTM1ZmNkYzc2OGNiMDAzMjVjMTcxMThiNjc0NDYwYzZlODNhYzVhMTQ3ZWJiMTBhIiwidWlkIjoiNzRjMzhjZjdjZWI5NDdmNTYzMjA0NWQ4Y2E1ZDQ4MTkiLCJ1c2VybmFtZSI6Im1heDMyIn0sInNpYmtleSI6eyJraWQiOiIwMTIwNTcyNTdlZWI4ZmVlZDVkMzMxZDYzYmEzODJiMGQ1Y2ZmYmExN2M5NWQzMmE4M2M3YTU4YjU1MjdiYWQ4NzMyNTBhIiwicmV2ZXJzZV9zaWciOiJnNlJpYjJSNWhxaGtaWFJoWTJobFpNT3BhR0Z6YUY5MGVYQmxDcU5yWlhuRUl3RWdWeVYrNjQvdTFkTXgxanVqZ3JEVnovdWhmSlhUS29QSHBZdFZKN3JZY3lVS3AzQmhlV3h2WVdURkFpdDdJbUp2WkhraU9uc2lhMlY1SWpwN0ltVnNaR1Z6ZEY5cmFXUWlPaUl3TVRJd056YzFNalUzTWpGaFpHSTFabUUwTkRrNFpUTTFabU5rWXpjMk9HTmlNREF6TWpWak1UY3hNVGhpTmpjME5EWXdZelpsT0ROaFl6VmhNVFEzWldKaU1UQmhJaXdpYUc5emRDSTZJbXRsZVdKaGMyVXVhVzhpTENKcmFXUWlPaUl3TVRJd056YzFNalUzTWpGaFpHSTFabUUwTkRrNFpUTTFabU5rWXpjMk9HTmlNREF6TWpWak1UY3hNVGhpTmpjME5EWXdZelpsT0ROaFl6VmhNVFEzWldKaU1UQmhJaXdpZFdsa0lqb2lOelJqTXpoalpqZGpaV0k1TkRkbU5UWXpNakEwTldRNFkyRTFaRFE0TVRraUxDSjFjMlZ5Ym1GdFpTSTZJbTFoZURNeUluMHNJbk5wWW10bGVTSTZleUpyYVdRaU9pSXdNVEl3TlRjeU5UZGxaV0k0Wm1WbFpEVmtNek14WkRZelltRXpPREppTUdRMVkyWm1ZbUV4TjJNNU5XUXpNbUU0TTJNM1lUVTRZalUxTWpkaVlXUTROek15TlRCaElpd2ljbVYyWlhKelpWOXphV2NpT201MWJHeDlMQ0owZVhCbElqb2ljMmxpYTJWNUlpd2lkbVZ5YzJsdmJpSTZNWDBzSW1OMGFXMWxJam94TkRNeU1UUTBNelExTENKbGVIQnBjbVZmYVc0aU9qRXdNREF3TURBd0xDSndjbVYySWpvaU1tVmtaR1F6Wm1FNE0yWTNaak0wT0RBNE5USTJZakk0WW1JNE9EZGhPVFZqTURsa1l6Z3haamd6TTJNNE9HSXlNelJoTkdKbE1qTXhNV1JrT0RsbU5DSXNJbk5sY1Y5MGVYQmxJam94TENKelpYRnVieUk2TWl3aWRHRm5Jam9pYzJsbmJtRjBkWEpsSW4yamMybG54RUFBQUFBQVFjZ09pSjlOVlB3OVIvUFZhc2l1b3cxQUtkQ1RuY0pjbXFKVFljRyswVTJDU1ZVc0tiVDZpb3A3cUV1Sm43a0Q1Tk5vejNrUkUzMVNKcXNKcUhOcFoxOTBlWEJsSUtOMFlXZk5BZ0tuZG1WeWMybHZiZ0U9In0sInR5cGUiOiJzaWJrZXkiLCJ2ZXJzaW9uIjoxfSwiY3RpbWUiOjE0MzIxNDQzNDUsImV4cGlyZV9pbiI6MTAwMDAwMDAsInByZXYiOiIyZWRkZDNmYTgzZjdmMzQ4MDg1MjZiMjhiYjg4N2E5NWMwOWRjODFmODMzYzg4YjIzNGE0YmUyMzExZGQ4OWY0Iiwic2VxX3R5cGUiOjEsInNlcW5vIjoyLCJ0YWciOiJzaWduYXR1cmUifaNzaWfEQJNGnPiaFJfzlUQRl5WxoTeDGRC7ZyJu25tx6+OqHEDx+d2kRfy7RvN05c6ZA46RIJZVb4Qp96q8QNOReDRecAyoc2lnX3R5cGUgo3RhZ80CAqd2ZXJzaW9uAQ==",
            "payload_hash": "4ccc4185a501dd16f2e2226182d3fadf8558aa59b593b5e4a239aea35581102d",
            "sig_id": "2d8dcd8fe4e7ceb9c67a76c18bbee71524d0ba0aa7c2e10298ce5688cf3cd2c80f",
            "payload_json": "{\"body\":{\"key\":{\"eldest_kid\":\"012077525721adb5fa4498e35fcdc768cb00325c17118b674460c6e83ac5a147ebb10a\",\"host\":\"keybase.io\",\"kid\":\"012077525721adb5fa4498e35fcdc768cb00325c17118b674460c6e83ac5a147ebb10a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"sibkey\":{\"kid\":\"012057257eeb8feed5d331d63ba382b0d5cffba17c95d32a83c7a58b5527bad873250a\",\"reverse_sig\":\"g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEgVyV+64/u1dMx1jujgrDVz/uhfJXTKoPHpYtVJ7rYcyUKp3BheWxvYWTFAit7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTIwNzc1MjU3MjFhZGI1ZmE0NDk4ZTM1ZmNkYzc2OGNiMDAzMjVjMTcxMThiNjc0NDYwYzZlODNhYzVhMTQ3ZWJiMTBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwNzc1MjU3MjFhZGI1ZmE0NDk4ZTM1ZmNkYzc2OGNiMDAzMjVjMTcxMThiNjc0NDYwYzZlODNhYzVhMTQ3ZWJiMTBhIiwidWlkIjoiNzRjMzhjZjdjZWI5NDdmNTYzMjA0NWQ4Y2E1ZDQ4MTkiLCJ1c2VybmFtZSI6Im1heDMyIn0sInNpYmtleSI6eyJraWQiOiIwMTIwNTcyNTdlZWI4ZmVlZDVkMzMxZDYzYmEzODJiMGQ1Y2ZmYmExN2M5NWQzMmE4M2M3YTU4YjU1MjdiYWQ4NzMyNTBhIiwicmV2ZXJzZV9zaWciOm51bGx9LCJ0eXBlIjoic2lia2V5IiwidmVyc2lvbiI6MX0sImN0aW1lIjoxNDMyMTQ0MzQ1LCJleHBpcmVfaW4iOjEwMDAwMDAwLCJwcmV2IjoiMmVkZGQzZmE4M2Y3ZjM0ODA4NTI2YjI4YmI4ODdhOTVjMDlkYzgxZjgzM2M4OGIyMzRhNGJlMjMxMWRkODlmNCIsInNlcV90eXBlIjoxLCJzZXFubyI6MiwidGFnIjoic2lnbmF0dXJlIn2jc2lnxEAAAAAAQcgOiJ9NVPw9R/PVasiuow1AKdCTncJcmqJTYcG+0U2CSVUsKbT6iop7qEuJn7kD5NNoz3kRE31SJqsJqHNpZ190eXBlIKN0YWfNAgKndmVyc2lvbgE=\"},\"type\":\"sibkey\",\"version\":1},\"ctime\":1432144345,\"expire_in\":10000000,\"prev\":\"2eddd3fa83f7f34808526b28bb887a95c09dc81f833c88b234a4be2311dd89f4\",\"seq_type\":1,\"seqno\":2,\"tag\":\"signature\"}",
            "kid": "012077525721adb5fa4498e35fcdc768cb00325c17118b674460c6e83ac5a147ebb10a",
            "ctime": 1432144345
        }
    ],
    "keys": [
        "012077525721adb5fa4498e35fcdc768cb00325c17118b674460c6e83ac5a147ebb10a",
        "012057257eeb8feed5d331d63ba382b0d5cffba17c95d32a83c7a58b5527bad873250a"
    ],
    "uid": "74c38cf7ceb947f5632045d8ca5d4819",
    "username": "max32"
}
`,
	"bad_seqno_chain.json": `{
    "chain": [
        {
            "seqno": 1,
            "prev": null,
            "sig": "g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEgsDiGCNcoqXQu5wYLlcIpCJv1Vlj7OJ3mmHN1pC8If7cKp3BheWxvYWTFASp7ImJvZHkiOnsia2V5Ijp7Imhvc3QiOiJrZXliYXNlLmlvIiwia2lkIjoiMDEyMGIwMzg4NjA4ZDcyOGE5NzQyZWU3MDYwYjk1YzIyOTA4OWJmNTU2NThmYjM4OWRlNjk4NzM3NWE0MmYwODdmYjcwYSIsInVpZCI6Ijc0YzM4Y2Y3Y2ViOTQ3ZjU2MzIwNDVkOGNhNWQ0ODE5IiwidXNlcm5hbWUiOiJtYXgzMiJ9LCJ0eXBlIjoiZWxkZXN0IiwidmVyc2lvbiI6MX0sImN0aW1lIjoxNDMyMTQ0MzQzLCJleHBpcmVfaW4iOjEwMDAwMDAwLCJwcmV2IjpudWxsLCJzZXFfdHlwZSI6MSwic2Vxbm8iOjEsInRhZyI6InNpZ25hdHVyZSJ9o3NpZ8RAzscTjEcAa9aF8Ze5MOUpDaR2BRSFObI7CtWlFLSAySxG0kKv99hqh53WjTyUL9G4CZcg14o1WSJcS9P05XI2DahzaWdfdHlwZSCjdGFnzQICp3ZlcnNpb24B",
            "payload_hash": "fafd9bbd0b7e905b96cb4b6c8022f7ee712098ef78ae048acb2fa296220d760f",
            "sig_id": "8dcc6e4d4de09e9bf0a1bf543e9bafbef589a2dce7797bd45296414a7b7cbff50f",
            "payload_json": "{\"body\":{\"key\":{\"host\":\"keybase.io\",\"kid\":\"0120b0388608d728a9742ee7060b95c229089bf55658fb389de6987375a42f087fb70a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"type\":\"eldest\",\"version\":1},\"ctime\":1432144343,\"expire_in\":10000000,\"prev\":null,\"seq_type\":1,\"seqno\":1,\"tag\":\"signature\"}",
            "kid": "0120b0388608d728a9742ee7060b95c229089bf55658fb389de6987375a42f087fb70a",
            "ctime": 1432144343
        },
        {
            "seqno": 3,
            "prev": "fafd9bbd0b7e905b96cb4b6c8022f7ee712098ef78ae048acb2fa296220d760f",
            "sig": "g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEgsDiGCNcoqXQu5wYLlcIpCJv1Vlj7OJ3mmHN1pC8If7cKp3BheWxvYWTFBfl7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTIwYjAzODg2MDhkNzI4YTk3NDJlZTcwNjBiOTVjMjI5MDg5YmY1NTY1OGZiMzg5ZGU2OTg3Mzc1YTQyZjA4N2ZiNzBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwYjAzODg2MDhkNzI4YTk3NDJlZTcwNjBiOTVjMjI5MDg5YmY1NTY1OGZiMzg5ZGU2OTg3Mzc1YTQyZjA4N2ZiNzBhIiwidWlkIjoiNzRjMzhjZjdjZWI5NDdmNTYzMjA0NWQ4Y2E1ZDQ4MTkiLCJ1c2VybmFtZSI6Im1heDMyIn0sInNpYmtleSI6eyJraWQiOiIwMTIwMTVkMDllOWYzZDEzMjQyYzBiNTMxMjI4ODI1YjI3NjIxNjNmYzRhNzJlY2EzNDM5ZjkxODY3NzM0ZGEzMDVlMjBhIiwicmV2ZXJzZV9zaWciOiJnNlJpYjJSNWhxaGtaWFJoWTJobFpNT3BhR0Z6YUY5MGVYQmxDcU5yWlhuRUl3RWdGZENlbnowVEpDd0xVeElvZ2xzblloWS94S2N1eWpRNStSaG5jMDJqQmVJS3AzQmhlV3h2WVdURkFpdDdJbUp2WkhraU9uc2lhMlY1SWpwN0ltVnNaR1Z6ZEY5cmFXUWlPaUl3TVRJd1lqQXpPRGcyTURoa056STRZVGszTkRKbFpUY3dOakJpT1RWak1qSTVNRGc1WW1ZMU5UWTFPR1ppTXpnNVpHVTJPVGczTXpjMVlUUXlaakE0TjJaaU56QmhJaXdpYUc5emRDSTZJbXRsZVdKaGMyVXVhVzhpTENKcmFXUWlPaUl3TVRJd1lqQXpPRGcyTURoa056STRZVGszTkRKbFpUY3dOakJpT1RWak1qSTVNRGc1WW1ZMU5UWTFPR1ppTXpnNVpHVTJPVGczTXpjMVlUUXlaakE0TjJaaU56QmhJaXdpZFdsa0lqb2lOelJqTXpoalpqZGpaV0k1TkRkbU5UWXpNakEwTldRNFkyRTFaRFE0TVRraUxDSjFjMlZ5Ym1GdFpTSTZJbTFoZURNeUluMHNJbk5wWW10bGVTSTZleUpyYVdRaU9pSXdNVEl3TVRWa01EbGxPV1l6WkRFek1qUXlZekJpTlRNeE1qSTRPREkxWWpJM05qSXhOak5tWXpSaE56SmxZMkV6TkRNNVpqa3hPRFkzTnpNMFpHRXpNRFZsTWpCaElpd2ljbVYyWlhKelpWOXphV2NpT201MWJHeDlMQ0owZVhCbElqb2ljMmxpYTJWNUlpd2lkbVZ5YzJsdmJpSTZNWDBzSW1OMGFXMWxJam94TkRNeU1UUTBORFF6TENKbGVIQnBjbVZmYVc0aU9qRXdNREF3TURBd0xDSndjbVYySWpvaVptRm1aRGxpWW1Rd1lqZGxPVEExWWprMlkySTBZalpqT0RBeU1tWTNaV1UzTVRJd09UaGxaamM0WVdVd05EaGhZMkl5Wm1FeU9UWXlNakJrTnpZd1ppSXNJbk5sY1Y5MGVYQmxJam94TENKelpYRnVieUk2TXl3aWRHRm5Jam9pYzJsbmJtRjBkWEpsSW4yamMybG54RUJaUG5JU0RESDNyUGFzUnVFaWJsdDJISFA1ODdkMDVVTXhxTlp3bG9HdVZtSVIyK1NaTUdmUkdXQ3VOejVaWG9hNDcrR1NFcnY2UUQ3c1Z3QjRoOXdQcUhOcFoxOTBlWEJsSUtOMFlXZk5BZ0tuZG1WeWMybHZiZ0U9In0sInR5cGUiOiJzaWJrZXkiLCJ2ZXJzaW9uIjoxfSwiY3RpbWUiOjE0MzIxNDQ0NDMsImV4cGlyZV9pbiI6MTAwMDAwMDAsInByZXYiOiJmYWZkOWJiZDBiN2U5MDViOTZjYjRiNmM4MDIyZjdlZTcxMjA5OGVmNzhhZTA0OGFjYjJmYTI5NjIyMGQ3NjBmIiwic2VxX3R5cGUiOjEsInNlcW5vIjozLCJ0YWciOiJzaWduYXR1cmUifaNzaWfEQAtTwjpVwFGBi3D2zGgfLzaYis4pQx0xBeuSd6XUA/a/MK4+gWVRGG39liy+OEVrh8F7plp97mNn4j1OAZ3t+wmoc2lnX3R5cGUgo3RhZ80CAqd2ZXJzaW9uAQ==",
            "payload_hash": "db4a190291265dceb1d047089ca40c21cea4258e2eb7c353fd7aa19729e7b0c1",
            "sig_id": "9a0eb84b985fa57a38a850189206cd0cd257db3d9c6c1f6169373ba1b5b1e4be0f",
            "payload_json": "{\"body\":{\"key\":{\"eldest_kid\":\"0120b0388608d728a9742ee7060b95c229089bf55658fb389de6987375a42f087fb70a\",\"host\":\"keybase.io\",\"kid\":\"0120b0388608d728a9742ee7060b95c229089bf55658fb389de6987375a42f087fb70a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"sibkey\":{\"kid\":\"012015d09e9f3d13242c0b531228825b2762163fc4a72eca3439f91867734da305e20a\",\"reverse_sig\":\"g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEgFdCenz0TJCwLUxIoglsnYhY/xKcuyjQ5+Rhnc02jBeIKp3BheWxvYWTFAit7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTIwYjAzODg2MDhkNzI4YTk3NDJlZTcwNjBiOTVjMjI5MDg5YmY1NTY1OGZiMzg5ZGU2OTg3Mzc1YTQyZjA4N2ZiNzBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwYjAzODg2MDhkNzI4YTk3NDJlZTcwNjBiOTVjMjI5MDg5YmY1NTY1OGZiMzg5ZGU2OTg3Mzc1YTQyZjA4N2ZiNzBhIiwidWlkIjoiNzRjMzhjZjdjZWI5NDdmNTYzMjA0NWQ4Y2E1ZDQ4MTkiLCJ1c2VybmFtZSI6Im1heDMyIn0sInNpYmtleSI6eyJraWQiOiIwMTIwMTVkMDllOWYzZDEzMjQyYzBiNTMxMjI4ODI1YjI3NjIxNjNmYzRhNzJlY2EzNDM5ZjkxODY3NzM0ZGEzMDVlMjBhIiwicmV2ZXJzZV9zaWciOm51bGx9LCJ0eXBlIjoic2lia2V5IiwidmVyc2lvbiI6MX0sImN0aW1lIjoxNDMyMTQ0NDQzLCJleHBpcmVfaW4iOjEwMDAwMDAwLCJwcmV2IjoiZmFmZDliYmQwYjdlOTA1Yjk2Y2I0YjZjODAyMmY3ZWU3MTIwOThlZjc4YWUwNDhhY2IyZmEyOTYyMjBkNzYwZiIsInNlcV90eXBlIjoxLCJzZXFubyI6MywidGFnIjoic2lnbmF0dXJlIn2jc2lnxEBZPnISDDH3rPasRuEiblt2HHP587d05UMxqNZwloGuVmIR2+SZMGfRGWCuNz5ZXoa47+GSErv6QD7sVwB4h9wPqHNpZ190eXBlIKN0YWfNAgKndmVyc2lvbgE=\"},\"type\":\"sibkey\",\"version\":1},\"ctime\":1432144443,\"expire_in\":10000000,\"prev\":\"fafd9bbd0b7e905b96cb4b6c8022f7ee712098ef78ae048acb2fa296220d760f\",\"seq_type\":1,\"seqno\":3,\"tag\":\"signature\"}",
            "kid": "0120b0388608d728a9742ee7060b95c229089bf55658fb389de6987375a42f087fb70a",
            "ctime": 1432144443
        }
    ],
    "keys": [
        "0120b0388608d728a9742ee7060b95c229089bf55658fb389de6987375a42f087fb70a",
        "012015d09e9f3d13242c0b531228825b2762163fc4a72eca3439f91867734da305e20a"
    ],
    "uid": "74c38cf7ceb947f5632045d8ca5d4819",
    "username": "max32"
}
`,
	"bad_signature_chain.json": `{
    "chain": [
        {
            "seqno": 1,
            "prev": null,
            "sig": "-----BEGIN PGP MESSAGE-----\nVersion: Keybase OpenPGP v2.0.10\nComment: https://keybase.io/crypto\n\nyMDyAnicbZG9axVBFMU3fkEeCMGIYBWcLriE+d6ZBwYFRaysLCQxj9mZO+uSl93N\n7r4kL/EVgtj4B2ir2FinEMRCRCvFYGchigSxsLKxC7q7wc5pZs65v3suw31z8mjQ\nm03O725/+bTJpt69ejsKbgwuzeygOHdj1N9Bq9BdPs0SKIsyzWrURwriSHBLY08l\njihgLZTwAFQLACV47LGVlhMUott51XY0MbGpYCHNG68Rg9Q17n/41a6ACWHaSSo4\ncAXESkUki7yWxmPNsQTBBaWWRwBaSUUxBmdjCk4bxb3jkYwdNk3cqIuLuGXK+shC\nrHnkhWQUc+GUNcJxRXQLVlBmZg0aes1sMYomIarHRath6KD5RIg2oKzSPEN90hRt\nnbY04YwyLDQTIYKtIi1hkLYEPjwhKkrYQP1sNByGqIL1wWEo6USWd6/aJM2YKk0y\nU49KQJPXK8eC2V5w4viRdhlBb3rm34o2f08Fd1/+ef7h5t6vp/OTa+P3V65eeHbv\n89yD/cUnZ/ZXitPLt4KDU0s/zj76ee5r/uLO/MWH673h3lJSP853P377foCu37/8\nF4rxoso=\n=Zv9D\n-----END PGP MESSAGE-----\n",
            "payload_hash": "eea2476670616f48f64531bcad5ed592e053671dae9aa8ce35550af03151e77a",
            "sig_id": "dfc1b98e0a9094f3f68b74f698b7a084006bb88cf26ca592c84eac49663e7fd40f",
            "payload_json": "{\"body\":{\"key\":{\"fingerprint\":\"8eb754c2bf26072e09585fee295ee854bf0c6c41\",\"host\":\"keybase.io\",\"key_id\":\"295ee854bf0c6c41\",\"kid\":\"01139d6254e48e1c681637f96af09406e54522c47ee9868200edcb2ed9a84fd476bd0a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"type\":\"eldest\",\"version\":1},\"ctime\":1432305935,\"expire_in\":10000000,\"prev\":null,\"seq_type\":1,\"seqno\":1,\"tag\":\"signature\"}",
            "kid": "01139d6254e48e1c681637f96af09406e54522c47ee9868200edcb2ed9a84fd476bd0a",
            "ctime": 1432305935
        },
        {
            "seqno": 2,
            "prev": "eea2476670616f48f64531bcad5ed592e053671dae9aa8ce35550af03151e77a",
            "sig": "-----BEGIN PGP MESSAGE-----\nVersion: Keybase OpenPGP v2.0.10\nComment: https://keybase.io/crypto\n\nyMQhAnicrVU7jONEAM0dHBIrIR06CShXaRCXZeOxPf5EOgknsZPdxInzcbIJ0S1j\ne5xMsh5/87GPEy2cREWBkKBBSNAjIdEhPhUFHaK6ElEfBS1O9tA1VzLNaGbePL15\n1nv++ZUXCkd35qVvs8d/bLkbv/7CrQvmpXL7QdHynbRYeVBc4cOErxwcJ5cr4hQr\nRQYATnYEFvKYlzCwBQkInOjKAnIZmWcEDHnIsjYvYixLgsQyDHZsi8WOjCTedXhR\nsBwGFU+KLqFzHAURoUlOK2FLhLzNWi4rMCKLGRlK0MWYlSHGEuQtl7EFmwf5xYUf\n72/k4iwU41Pi53v54vIg7zn4/1n3+kAn8jYn2a5oY0vmRRcKHMvw0JFsBB1eAvIe\nGOOIIg/naA/tOLb48KQYE+upqc9UuVBE2OWB6LqAs5BrQwe7ruxYEmsBKOXM2AEy\ny7D5ieQyyIKOCCwRW5bA5Tj2oCrCGxzF+DIm85z27f2oqo2zzrHRMI51dTBQGuph\nd0ZHOZD4tHLcunbwuBtguodt2FPmFDAzWvM9D9OkcrxIkiCulMvPvC7bURok/ozO\naKprE4USO5qeiQj0q5rubFxVG899zZ+Mm4tJbdzOjPUuqDbKczVQY0MRlpkRdaZr\n0eM2m5zBR71d3G7UYb0xWqnxTrXUga/76nxORkG7kW76rLaO7KE9bDCe19BLTrm8\ndpue22FFWh/P6NghUuxiUxRFRpeTDeezGSpTrXRxsVp4aid3sa+vtpuGMvA9LoW1\nFmnqtaZeVVdaVRCqM4pb3vm05iM00NmmGvcutHq47p1tx17Y0nvDTUvpLzuTUrVV\nWtWSkLQnV03W7wejvrNajbegm7uFhYCAPun3VMs3TdVXJwrSNWenQWF41V6I46Yz\nIj1+yq55U8mG422Uthy+LBulKc5SMKOmvM1EZKS8srswzUVD7ztVnMY1IA2Y7SIM\nmRAqZlie1FDH9/lNuF3kjx3II00PalAdajM6OCcZM4kVfRrNO82RKnR7+rxPWsiO\nfNJihr1+g27r9Hzlw244CPkaMDt4xDEmkuz6FgW5k5bOoiB25W5KRleqPpB7yxb2\nLH9lLqnWaJyPESGskrZIdEYAqsqMdLGcnmf6vLHt6ECZ0RZJz5farpn2zGWWxPMz\nhQyN1XYbmmQ7aWWjZjDvKaHb18nZWatjEhDVkJh2+2ikdM2J5OWvSNZubzNts/w2\n6lUTq8PLqxrkYraj1cXh1KMOhUa2EUStzYYydUpiHfBJVo+wZCQrWZrR9jBKXSqk\n67KRCAHVciHZRbldHRqTJeY4bLgl2DE2YakbBoFtRrQEE7oEcsk1/FLIcTO6TLFB\nkwzW+5NEWDP3ZvSeCna5QYf0qJ36cxK1T3aSBvukPw34SXFznbJiBeSHdkL2PQB4\njuUYgeHgSRHvAhLhS7JHMNfjpBjkMc5JMEZsXjiCyAhAcHnJFXjIActGDsQOlPN2\nhJwgAgdhGeUfEHMQQibvMQ5AgEVx3wcxDi+vJYHDgvrFCpurRPODyDlFyTrCxYc/\n3n+xcOeo8NKtm/vqLxy9fPu/H8JPb94ofCxnDz599/PLQj6+Pvrqw2/eWP3wW124\n+/j9t+7yf43jG4Xvgn/okz/fe/L9l+88un/zM+x9cvX6R68+qvx967XTD34nX/wL\nbL0JeQ==\n=965F\n-----END PGP MESSAGE-----\n",
            "payload_hash": "3eafef1d80fc9e3c2d559bbc4a7b57522caaefc48d3a399cf9c2d7213284b0a7",
            "sig_id": "7b0b5f4fb66663c2290d78cddb1898236d6e065eb783e9aa85c8c67ac941e1a40f",
            "payload_json": "{\"body\":{\"key\":{\"eldest_kid\":\"01139d6254e48e1c681637f96af09406e54522c47ee9868200edcb2ed9a84fd476bd0a\",\"fingerprint\":\"8eb754c2bf26072e09585fee295ee854bf0c6c41\",\"host\":\"keybase.io\",\"key_id\":\"295ee854bf0c6c41\",\"kid\":\"01139d6254e48e1c681637f96af09406e54522c47ee9868200edcb2ed9a84fd476bd0a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"sibkey\":{\"kid\":\"0113f57aef417ff13bafc5deff9db82b158632ed19202fc58f0ab5d71b7ebb63ff920a\",\"reverse_sig\":\"-----BEGIN PGP MESSAGE-----\\nVersion: Keybase OpenPGP v2.0.10\\nComment: https://keybase.io/crypto\\n\\nyMFYAnicrZI7a1RBFMdvfEFWgoFoYWHhYCWLzPuxpBG/gEpEsPA6jzPrNZu7m3vv\\nyoaQxsLGD5DGVkEsxEbESoMoEggiVpLGyvR2FurcTcTG0mmGM+d//ufHmfN27nDW\\nWdi8sfeU7770M9tv3o2za/nF+XXkhmEN9dbRMkwvGASom3y5CKiHMCHMBEkFB66B\\neKmJZCoaaSM2HEsQXFDquQIwWmqKMQTvKARjNY+BK+kCtqiLYlH2oRpVRdkkWw1O\\nCe6pi1RiRQEboUUEoEYAaMFdxF56TlLh7WHdViQ4Z2u4UAzTWwryKd4/9P+Zezy1\\nU9wz7aPy4AxXUUhGMRdBeysC18S0whqq0q5AUq/YCaNoo4vqwh0M9S9VFMpC5ETF\\nSJiz0YsAMZrgNHVE6OQMgRiKacroiK0TQRGnwDnJko5OqSq4C1UNeV30Ua8cDwap\\nWbM2apsf9OyiVlEMS9QjKembokUjnFGGJWaii2AyKirIi1aB908XjZJzMgGwNM1A\\nKiyJjFxHyQUjztsgIAiTPkwwqUiwYKzVHpgQAqfRMiIIKNUi1rCa7yORaVAOUY8m\\nStufQvZL24wrQBtbN49kC53s2NFD7TZmndn5Pzv67FL2q9nd+7D14tzDre8Ptk98\\nLTryfn6yu/Pt6pnFjztzX/LBTPYje33ePf+5NPvq+OqppcUrn+5tnj19+fPo+q33\\njyePntz5DRYt6u0=\\n=E1xp\\n-----END PGP MESSAGE-----\\n\"},\"type\":\"sibkey\",\"version\":1},\"ctime\":1432306035,\"expire_in\":10000000,\"prev\":\"eea2476670616f48f64531bcad5ed592e053671dae9aa8ce35550af03151e77a\",\"seq_type\":1,\"seqno\":2,\"tag\":\"signature\"}",
            "kid": "01139d6254e48e1c681637f96af09406e54522c47ee9868200edcb2ed9a84fd476bd0a",
            "ctime": 1432306035
        }
    ],
    "keys": [
        "-----BEGIN PGP PUBLIC KEY BLOCK-----\nVersion: Keybase OpenPGP v2.0.10\nComment: https://keybase.io/crypto\n\nxm8EVV9BDxMFK4EEACIDAwRSzcsRfEMpwQ7RvwWHPxXf97lwjf8mqqCeTZXkntaK\nLAYBBi4ZH2dPpVL2Nk6Bh0K7Zc9II0ksc0BL+z0fJ/3hIDhq9NgfusjqjiX8NYZ7\ndbeT+gyDPp5gzXaxPqrF0vXNBW1heDMywpIEExMKABoFAlVfQQ8CGwEDCwkHAxUK\nCAIeAQIXgAIZAQAKCRApXuhUvwxsQUGxAX9Y1nv+tbCxZGE4P22kbVQbi4BTyIkL\nYgNKqeoOlNxuWeDdh0xaDsimsTpEKWtSagMBgK179Xs4gsE9jPBtHGuzYpbHbROS\nhVCq9ssZjvId45D2UHFSjyAm8spEeFJzLARqEs5WBFVfQQ8SCCqGSM49AwEHAgME\nntY21MrNuVBiA28QJRFx/+nw8O6URDuF7P+a1Ou+c/mzeH8bH9NB+fozm+wt0+kU\nsAQ1rrmAdK9oXfbxFHw+VgMBCgnChwQYEwoADwUCVV9BDwUJDwmcAAIbDAAKCRAp\nXuhUvwxsQYFyAYCuu59mWoB68DUaSFY1YoIjXt10oKkJyqeaF/MDKCs4RZvnLkcM\nVfm4ANHRx8P4jTwBgKF4bBRJYKbGCh9Sdve3ivaihIjKueXbIkIwzlHKnDhH0ryF\nZy7AuYErrAMqZiEyG85SBFVfQQ8TCCqGSM49AwEHAgME2hMvKEfCqKxsX7+B70Lq\nfOOQg3mAP5vNEE9fP/O4CHS3nG7DRv4Di1vZlHE7u2OHAfHSAFq3ir935z4C0d7B\ncsLAJwQYEwoADwUCVV9BDwUJDwmcAAIbIgBqCRApXuhUvwxsQV8gBBkTCgAGBQJV\nX0EPAAoJEGcrs3re2HczDekBAJOT5B0avXmVYnPGkTwVZ3oeoXNwQEpYOm/kdPEx\nBFVhAQCKGjWboi7vVftw1cU/IFp1uh7lYWEzDlB9yvUCyd7/INoFAYC9wV1R++yE\nroIfq9hkzttJnvmuobSYaJHLMIQXnIADOBLHbHxrKInDtYqIToUoYsoBf2CMDj0r\nVwx1bGn7KprTMnu89hv/rIjnAcDYorsmqfECoXXZNBrVHS/DBpDROPPpag==\n=wdBi\n-----END PGP PUBLIC KEY BLOCK-----\n",
        "-----BEGIN PGP PUBLIC KEY BLOCK-----\nVersion: Keybase OpenPGP v2.0.10\nComment: https://keybase.io/crypto\n\nxm8EVV9BEBMFK4EEACIDAwS9o76Gzhjx/aRI+8EdNO7F011t/m6clC6VHpKASpc/\nmnygwK6qvxou750c6WgrBlFtnskjjpJcY2zqkgSKi+MKtpt6PJwWfQ1CR6aJBqNe\nSJyJVWS0yLuvf94v8u+e2c/NBW1heDMywpIEExMKABoFAlVfQRACGwEDCwkHAxUK\nCAIeAQIXgAIZAQAKCRCRTHSn/50zJ11MAYDI+TPMhrTb73FWiUxTkRvehyfWrLzk\nSFQS0RZpujJpzaaI2tE8Q/g5KU2oLEZpvdkBgPubVOmKL0LSWaaQzGaPrv7pR5+D\naGjOMf1VlUYSprgDiHrPnft1EiAlUsWsfX8yKs5WBFVfQRASCCqGSM49AwEHAgME\n77H3BHjofaqD5KYjvEkq9SAxc9i8VBuSrWplGelgeG/Ctr9UIvindI7uoRhHqAVa\nbjTFylO+CZy88cKdsRcj3wMBCgnChwQYEwoADwUCVV9BEAUJDwmcAAIbDAAKCRCR\nTHSn/50zJ+KDAX9nCQ75y87n0onZqDR3N+CC8QNvj6z+oaxvEh9w8AoF5auU9AQn\ny+/L5h5TNuQ/UkwBgLhredcA4cJ1+iBUOIQKJituuVYwvZ5wRFzRw6O7bs9XrbQZ\nZhki4WY7K64PuLvSQs5SBFVfQRATCCqGSM49AwEHAgMErT7N/XUzhm0YrnBF5uCk\nl9UIbLoiN4Zv8Vfj+KRJ8cWxQZVGKUd1lbTSQpE7js9adGEGqloqCEaxPaAYWtor\nQcLAJwQYEwoADwUCVV9BEAUJDwmcAAIbIgBqCRCRTHSn/50zJ18gBBkTCgAGBQJV\nX0EQAAoJEJda6Ks03LhjKjcBAOc8qE3v+yh8buW/vgeHbxTZRQLKaJoGSHC+TvgZ\nCe9/APwPIo89pm4CLE5vukuv0rSV3kQUF/LLkAtG2kwcvIUGO/2kAX9+aHMa/KlU\nLR4NJzhprqbXG9hWqR/R0DJLCm3HrP/J7v2glQ/yRWrAqdaYvxhASZoBfiTH1x2U\nL1C2MzFN8F/M8MQiobo5Kp4O/jLvj7leBQk/mXgaDAZrnAulV0cYn9xgfA==\n=PeLw\n-----END PGP PUBLIC KEY BLOCK-----\n"
    ],
    "uid": "74c38cf7ceb947f5632045d8ca5d4819",
    "username": "max32"
}
`,
	"bad_uid_chain.json": `{
    "chain": [
        {
            "seqno": 1,
            "prev": null,
            "sig": "g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEgj8Et28Wy41nCNQ2UVwC2dpyeKou5x1Z51AtO17SyjGsKp3BheWxvYWTFASp7ImJvZHkiOnsia2V5Ijp7Imhvc3QiOiJrZXliYXNlLmlvIiwia2lkIjoiMDEyMDhmYzEyZGRiYzViMmUzNTljMjM1MGQ5NDU3MDBiNjc2OWM5ZTJhOGJiOWM3NTY3OWQ0MGI0ZWQ3YjRiMjhjNmIwYSIsInVpZCI6ImRlYWRiZWVmMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwIiwidXNlcm5hbWUiOiJtYXgzMiJ9LCJ0eXBlIjoiZWxkZXN0IiwidmVyc2lvbiI6MX0sImN0aW1lIjoxNDMyMTQ0NTk1LCJleHBpcmVfaW4iOjEwMDAwMDAwLCJwcmV2IjpudWxsLCJzZXFfdHlwZSI6MSwic2Vxbm8iOjEsInRhZyI6InNpZ25hdHVyZSJ9o3NpZ8RAc4WxvRU1roBEB1HStIiE2d48q4cOLgHlYSu9sGawOYOc4nvDvrzPKq72jrTgjSJFOuwa9t0bQDpXMTD1/s/XA6hzaWdfdHlwZSCjdGFnzQICp3ZlcnNpb24B",
            "payload_hash": "4ffdc8466a0ffa380da8c4ef6477cc80e3a80734eda97d148f659bc0410c7344",
            "sig_id": "84083329f6ebe6fc7157cb7be43a6ea7526db482416949d1f54eb5159eb8c4890f",
            "payload_json": "{\"body\":{\"key\":{\"host\":\"keybase.io\",\"kid\":\"01208fc12ddbc5b2e359c2350d945700b6769c9e2a8bb9c75679d40b4ed7b4b28c6b0a\",\"uid\":\"deadbeef000000000000000000000000\",\"username\":\"max32\"},\"type\":\"eldest\",\"version\":1},\"ctime\":1432144595,\"expire_in\":10000000,\"prev\":null,\"seq_type\":1,\"seqno\":1,\"tag\":\"signature\"}",
            "kid": "01208fc12ddbc5b2e359c2350d945700b6769c9e2a8bb9c75679d40b4ed7b4b28c6b0a",
            "ctime": 1432144595
        }
    ],
    "keys": [
        "01208fc12ddbc5b2e359c2350d945700b6769c9e2a8bb9c75679d40b4ed7b4b28c6b0a"
    ],
    "uid": "74c38cf7ceb947f5632045d8ca5d4819",
    "username": "max32"
}
`,
	"bad_username_chain.json": `{
    "chain": [
        {
            "seqno": 1,
            "prev": null,
            "sig": "g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEguI0gTYv5jJi0OvFbTgAsh7xXonZn2UXGOIZlClWJN/oKp3BheWxvYWTFATF7ImJvZHkiOnsia2V5Ijp7Imhvc3QiOiJrZXliYXNlLmlvIiwia2lkIjoiMDEyMGI4OGQyMDRkOGJmOThjOThiNDNhZjE1YjRlMDAyYzg3YmM1N2EyNzY2N2Q5NDVjNjM4ODY2NTBhNTU4OTM3ZmEwYSIsInVpZCI6Ijc0YzM4Y2Y3Y2ViOTQ3ZjU2MzIwNDVkOGNhNWQ0ODE5IiwidXNlcm5hbWUiOiJnYXJiYWdlX3VzZXIifSwidHlwZSI6ImVsZGVzdCIsInZlcnNpb24iOjF9LCJjdGltZSI6MTQzMjE0NDY0MSwiZXhwaXJlX2luIjoxMDAwMDAwMCwicHJldiI6bnVsbCwic2VxX3R5cGUiOjEsInNlcW5vIjoxLCJ0YWciOiJzaWduYXR1cmUifaNzaWfEQG0IAj1rTzCBW2H64Hh66oer8vzOAmiL70khO1/49o8hL/QjXRUm8/8aabGrHmESu7FC+TU6oyvO1sH++bK2LgSoc2lnX3R5cGUgo3RhZ80CAqd2ZXJzaW9uAQ==",
            "payload_hash": "a867b0e83da675b6076e3ba7e5f34e2020ea96836fba50e6e1a4092c9b653aac",
            "sig_id": "a9d87f33ed10c1295bedb5b7369b9af11bbb69fc38872b8a7a85cb660049dc480f",
            "payload_json": "{\"body\":{\"key\":{\"host\":\"keybase.io\",\"kid\":\"0120b88d204d8bf98c98b43af15b4e002c87bc57a27667d945c63886650a558937fa0a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"garbage_user\"},\"type\":\"eldest\",\"version\":1},\"ctime\":1432144641,\"expire_in\":10000000,\"prev\":null,\"seq_type\":1,\"seqno\":1,\"tag\":\"signature\"}",
            "kid": "0120b88d204d8bf98c98b43af15b4e002c87bc57a27667d945c63886650a558937fa0a",
            "ctime": 1432144641
        }
    ],
    "keys": [
        "0120b88d204d8bf98c98b43af15b4e002c87bc57a27667d945c63886650a558937fa0a"
    ],
    "uid": "74c38cf7ceb947f5632045d8ca5d4819",
    "username": "max32"
}
`,
	"example_revokes_chain.json": `{
    "chain": [
        {
            "seqno": 1,
            "prev": null,
            "sig": "g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEgKvldOpi5fFXv0twJsMnGBv0LXqp+NbWaV14zbexQTUkKp3BheWxvYWTFASp7ImJvZHkiOnsia2V5Ijp7Imhvc3QiOiJrZXliYXNlLmlvIiwia2lkIjoiMDEyMDJhZjk1ZDNhOThiOTdjNTVlZmQyZGMwOWIwYzljNjA2ZmQwYjVlYWE3ZTM1YjU5YTU3NWUzMzZkZWM1MDRkNDkwYSIsInVpZCI6Ijc0YzM4Y2Y3Y2ViOTQ3ZjU2MzIwNDVkOGNhNWQ0ODE5IiwidXNlcm5hbWUiOiJtYXgzMiJ9LCJ0eXBlIjoiZWxkZXN0IiwidmVyc2lvbiI6MX0sImN0aW1lIjoxNDMyMTQ0NzAwLCJleHBpcmVfaW4iOjEwMDAwMDAwLCJwcmV2IjpudWxsLCJzZXFfdHlwZSI6MSwic2Vxbm8iOjEsInRhZyI6InNpZ25hdHVyZSJ9o3NpZ8RAelPb6xpVpK5Kdk57aOoiA438uOVw8Ha6eHBGydFnWYLw12EWT02pKXQYnIvCVSQNOJIMeXXn99hPYLSS3cv2DKhzaWdfdHlwZSCjdGFnzQICp3ZlcnNpb24B",
            "payload_hash": "69a9b17aca289091b869bddb00385857147515fb4c335d157fd0be86f0bb5e9c",
            "sig_id": "71e44e4c690d0d7ac6bcb244d9f72cf173bfae15e3fa7d4cb0ee2f24c3d6de9f0f",
            "payload_json": "{\"body\":{\"key\":{\"host\":\"keybase.io\",\"kid\":\"01202af95d3a98b97c55efd2dc09b0c9c606fd0b5eaa7e35b59a575e336dec504d490a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"type\":\"eldest\",\"version\":1},\"ctime\":1432144700,\"expire_in\":10000000,\"prev\":null,\"seq_type\":1,\"seqno\":1,\"tag\":\"signature\"}",
            "kid": "01202af95d3a98b97c55efd2dc09b0c9c606fd0b5eaa7e35b59a575e336dec504d490a",
            "ctime": 1432144700
        },
        {
            "seqno": 2,
            "prev": "69a9b17aca289091b869bddb00385857147515fb4c335d157fd0be86f0bb5e9c",
            "sig": "g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEgKvldOpi5fFXv0twJsMnGBv0LXqp+NbWaV14zbexQTUkKp3BheWxvYWTFBfl7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTIwMmFmOTVkM2E5OGI5N2M1NWVmZDJkYzA5YjBjOWM2MDZmZDBiNWVhYTdlMzViNTlhNTc1ZTMzNmRlYzUwNGQ0OTBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwMmFmOTVkM2E5OGI5N2M1NWVmZDJkYzA5YjBjOWM2MDZmZDBiNWVhYTdlMzViNTlhNTc1ZTMzNmRlYzUwNGQ0OTBhIiwidWlkIjoiNzRjMzhjZjdjZWI5NDdmNTYzMjA0NWQ4Y2E1ZDQ4MTkiLCJ1c2VybmFtZSI6Im1heDMyIn0sInNpYmtleSI6eyJraWQiOiIwMTIwZDIwZmU5ODJlNGU3YmExYTA4ZGQ2MGRkNTI3YTAyN2JiZTAyMjVkNGZlODllZWY5NzkwMjBjNWI5ZWZlZGJmYjBhIiwicmV2ZXJzZV9zaWciOiJnNlJpYjJSNWhxaGtaWFJoWTJobFpNT3BhR0Z6YUY5MGVYQmxDcU5yWlhuRUl3RWcwZy9wZ3VUbnVob0kzV0RkVW5vQ2U3NENKZFQraWU3NWVRSU1XNTcrMi9zS3AzQmhlV3h2WVdURkFpdDdJbUp2WkhraU9uc2lhMlY1SWpwN0ltVnNaR1Z6ZEY5cmFXUWlPaUl3TVRJd01tRm1PVFZrTTJFNU9HSTVOMk0xTldWbVpESmtZekE1WWpCak9XTTJNRFptWkRCaU5XVmhZVGRsTXpWaU5UbGhOVGMxWlRNek5tUmxZelV3TkdRME9UQmhJaXdpYUc5emRDSTZJbXRsZVdKaGMyVXVhVzhpTENKcmFXUWlPaUl3TVRJd01tRm1PVFZrTTJFNU9HSTVOMk0xTldWbVpESmtZekE1WWpCak9XTTJNRFptWkRCaU5XVmhZVGRsTXpWaU5UbGhOVGMxWlRNek5tUmxZelV3TkdRME9UQmhJaXdpZFdsa0lqb2lOelJqTXpoalpqZGpaV0k1TkRkbU5UWXpNakEwTldRNFkyRTFaRFE0TVRraUxDSjFjMlZ5Ym1GdFpTSTZJbTFoZURNeUluMHNJbk5wWW10bGVTSTZleUpyYVdRaU9pSXdNVEl3WkRJd1ptVTVPREpsTkdVM1ltRXhZVEE0WkdRMk1HUmtOVEkzWVRBeU4ySmlaVEF5TWpWa05HWmxPRGxsWldZNU56a3dNakJqTldJNVpXWmxaR0ptWWpCaElpd2ljbVYyWlhKelpWOXphV2NpT201MWJHeDlMQ0owZVhCbElqb2ljMmxpYTJWNUlpd2lkbVZ5YzJsdmJpSTZNWDBzSW1OMGFXMWxJam94TkRNeU1UUTBPREF3TENKbGVIQnBjbVZmYVc0aU9qRXdNREF3TURBd0xDSndjbVYySWpvaU5qbGhPV0l4TjJGallUSTRPVEE1TVdJNE5qbGlaR1JpTURBek9EVTROVGN4TkRjMU1UVm1ZalJqTXpNMVpERTFOMlprTUdKbE9EWm1NR0ppTldVNVl5SXNJbk5sY1Y5MGVYQmxJam94TENKelpYRnVieUk2TWl3aWRHRm5Jam9pYzJsbmJtRjBkWEpsSW4yamMybG54RUJqcWhUQWpnaXNEVFpJbnhjaVNReVhGZE1CVlJhbDk4R3pZMVp1VkxwVU9VNWd5cjdIazh4cnBSSDdmUHRYcFdHUnArdVpzaWdrVlZtVS8xbXJMY0FPcUhOcFoxOTBlWEJsSUtOMFlXZk5BZ0tuZG1WeWMybHZiZ0U9In0sInR5cGUiOiJzaWJrZXkiLCJ2ZXJzaW9uIjoxfSwiY3RpbWUiOjE0MzIxNDQ4MDAsImV4cGlyZV9pbiI6MTAwMDAwMDAsInByZXYiOiI2OWE5YjE3YWNhMjg5MDkxYjg2OWJkZGIwMDM4NTg1NzE0NzUxNWZiNGMzMzVkMTU3ZmQwYmU4NmYwYmI1ZTljIiwic2VxX3R5cGUiOjEsInNlcW5vIjoyLCJ0YWciOiJzaWduYXR1cmUifaNzaWfEQHBvlAFnu1sE4GKJKz7iSTkCMiGM/tlf4gDtrbKOYrlWTiceY/LdWffEGgUqX/Rox95eaf9j+InG0nMlHKtPcwqoc2lnX3R5cGUgo3RhZ80CAqd2ZXJzaW9uAQ==",
            "payload_hash": "29ef238a6283a937300a13b3e5542a92e3b9014380c0798979a9b27133d6c8ee",
            "sig_id": "77a2a70d5c8af0f1201027076504065e80ddd6d5ff2a622eb57e149dc8fa00470f",
            "payload_json": "{\"body\":{\"key\":{\"eldest_kid\":\"01202af95d3a98b97c55efd2dc09b0c9c606fd0b5eaa7e35b59a575e336dec504d490a\",\"host\":\"keybase.io\",\"kid\":\"01202af95d3a98b97c55efd2dc09b0c9c606fd0b5eaa7e35b59a575e336dec504d490a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"sibkey\":{\"kid\":\"0120d20fe982e4e7ba1a08dd60dd527a027bbe0225d4fe89eef979020c5b9efedbfb0a\",\"reverse_sig\":\"g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEg0g/pguTnuhoI3WDdUnoCe74CJdT+ie75eQIMW57+2/sKp3BheWxvYWTFAit7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTIwMmFmOTVkM2E5OGI5N2M1NWVmZDJkYzA5YjBjOWM2MDZmZDBiNWVhYTdlMzViNTlhNTc1ZTMzNmRlYzUwNGQ0OTBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwMmFmOTVkM2E5OGI5N2M1NWVmZDJkYzA5YjBjOWM2MDZmZDBiNWVhYTdlMzViNTlhNTc1ZTMzNmRlYzUwNGQ0OTBhIiwidWlkIjoiNzRjMzhjZjdjZWI5NDdmNTYzMjA0NWQ4Y2E1ZDQ4MTkiLCJ1c2VybmFtZSI6Im1heDMyIn0sInNpYmtleSI6eyJraWQiOiIwMTIwZDIwZmU5ODJlNGU3YmExYTA4ZGQ2MGRkNTI3YTAyN2JiZTAyMjVkNGZlODllZWY5NzkwMjBjNWI5ZWZlZGJmYjBhIiwicmV2ZXJzZV9zaWciOm51bGx9LCJ0eXBlIjoic2lia2V5IiwidmVyc2lvbiI6MX0sImN0aW1lIjoxNDMyMTQ0ODAwLCJleHBpcmVfaW4iOjEwMDAwMDAwLCJwcmV2IjoiNjlhOWIxN2FjYTI4OTA5MWI4NjliZGRiMDAzODU4NTcxNDc1MTVmYjRjMzM1ZDE1N2ZkMGJlODZmMGJiNWU5YyIsInNlcV90eXBlIjoxLCJzZXFubyI6MiwidGFnIjoic2lnbmF0dXJlIn2jc2lnxEBjqhTAjgisDTZInxciSQyXFdMBVRal98GzY1ZuVLpUOU5gyr7Hk8xrpRH7fPtXpWGRp+uZsigkVVmU/1mrLcAOqHNpZ190eXBlIKN0YWfNAgKndmVyc2lvbgE=\"},\"type\":\"sibkey\",\"version\":1},\"ctime\":1432144800,\"expire_in\":10000000,\"prev\":\"69a9b17aca289091b869bddb00385857147515fb4c335d157fd0be86f0bb5e9c\",\"seq_type\":1,\"seqno\":2,\"tag\":\"signature\"}",
            "kid": "01202af95d3a98b97c55efd2dc09b0c9c606fd0b5eaa7e35b59a575e336dec504d490a",
            "ctime": 1432144800
        },
        {
            "seqno": 3,
            "prev": "29ef238a6283a937300a13b3e5542a92e3b9014380c0798979a9b27133d6c8ee",
            "sig": "g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEg0g/pguTnuhoI3WDdUnoCe74CJdT+ie75eQIMW57+2/sKp3BheWxvYWTFBfl7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTIwMmFmOTVkM2E5OGI5N2M1NWVmZDJkYzA5YjBjOWM2MDZmZDBiNWVhYTdlMzViNTlhNTc1ZTMzNmRlYzUwNGQ0OTBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwZDIwZmU5ODJlNGU3YmExYTA4ZGQ2MGRkNTI3YTAyN2JiZTAyMjVkNGZlODllZWY5NzkwMjBjNWI5ZWZlZGJmYjBhIiwidWlkIjoiNzRjMzhjZjdjZWI5NDdmNTYzMjA0NWQ4Y2E1ZDQ4MTkiLCJ1c2VybmFtZSI6Im1heDMyIn0sInNpYmtleSI6eyJraWQiOiIwMTIwMzFmZTc1Y2FlMTZlMzY0OTQ5Zjk5YjZmYzRmZTc5ZmFiMGIwN2QxMjM0ZGQ1MGI5ZWFmYjBkNzMxNTcxYmY2YzBhIiwicmV2ZXJzZV9zaWciOiJnNlJpYjJSNWhxaGtaWFJoWTJobFpNT3BhR0Z6YUY5MGVYQmxDcU5yWlhuRUl3RWdNZjUxeXVGdU5rbEorWnR2eFA1NStyQ3dmUkkwM1ZDNTZ2c05jeFZ4djJ3S3AzQmhlV3h2WVdURkFpdDdJbUp2WkhraU9uc2lhMlY1SWpwN0ltVnNaR1Z6ZEY5cmFXUWlPaUl3TVRJd01tRm1PVFZrTTJFNU9HSTVOMk0xTldWbVpESmtZekE1WWpCak9XTTJNRFptWkRCaU5XVmhZVGRsTXpWaU5UbGhOVGMxWlRNek5tUmxZelV3TkdRME9UQmhJaXdpYUc5emRDSTZJbXRsZVdKaGMyVXVhVzhpTENKcmFXUWlPaUl3TVRJd1pESXdabVU1T0RKbE5HVTNZbUV4WVRBNFpHUTJNR1JrTlRJM1lUQXlOMkppWlRBeU1qVmtOR1psT0RsbFpXWTVOemt3TWpCak5XSTVaV1psWkdKbVlqQmhJaXdpZFdsa0lqb2lOelJqTXpoalpqZGpaV0k1TkRkbU5UWXpNakEwTldRNFkyRTFaRFE0TVRraUxDSjFjMlZ5Ym1GdFpTSTZJbTFoZURNeUluMHNJbk5wWW10bGVTSTZleUpyYVdRaU9pSXdNVEl3TXpGbVpUYzFZMkZsTVRabE16WTBPVFE1WmprNVlqWm1ZelJtWlRjNVptRmlNR0l3TjJReE1qTTBaR1ExTUdJNVpXRm1ZakJrTnpNeE5UY3hZbVkyWXpCaElpd2ljbVYyWlhKelpWOXphV2NpT201MWJHeDlMQ0owZVhCbElqb2ljMmxpYTJWNUlpd2lkbVZ5YzJsdmJpSTZNWDBzSW1OMGFXMWxJam94TkRNeU1UUTBPVEF3TENKbGVIQnBjbVZmYVc0aU9qRXdNREF3TURBd0xDSndjbVYySWpvaU1qbGxaakl6T0dFMk1qZ3pZVGt6TnpNd01HRXhNMkl6WlRVMU5ESmhPVEpsTTJJNU1ERTBNemd3WXpBM09UZzVOemxoT1dJeU56RXpNMlEyWXpobFpTSXNJbk5sY1Y5MGVYQmxJam94TENKelpYRnVieUk2TXl3aWRHRm5Jam9pYzJsbmJtRjBkWEpsSW4yamMybG54RUFjcEZVaUVWUkQyeXZxSWEwdDkzVnluamRmOGxTcnBWWGV2aHNIdDA2TGVLc1IrU0RYYTk4RUpvNlltL2ZqcFd4Smo4ZWpzbktCTGFncUV0UWZFakFNcUhOcFoxOTBlWEJsSUtOMFlXZk5BZ0tuZG1WeWMybHZiZ0U9In0sInR5cGUiOiJzaWJrZXkiLCJ2ZXJzaW9uIjoxfSwiY3RpbWUiOjE0MzIxNDQ5MDAsImV4cGlyZV9pbiI6MTAwMDAwMDAsInByZXYiOiIyOWVmMjM4YTYyODNhOTM3MzAwYTEzYjNlNTU0MmE5MmUzYjkwMTQzODBjMDc5ODk3OWE5YjI3MTMzZDZjOGVlIiwic2VxX3R5cGUiOjEsInNlcW5vIjozLCJ0YWciOiJzaWduYXR1cmUifaNzaWfEQCsPuE8lW37tjH5kUlKa7ysGMyqGN6QdhRcdxwX3J6rtk1IYeXhNX7ll3kVeLLBj5t84wtp6zDm1W5LdT/Aciwioc2lnX3R5cGUgo3RhZ80CAqd2ZXJzaW9uAQ==",
            "payload_hash": "c3deb587280bf27e7b7554c0304c5b58c4183d888306c65386d7a57287ba0a8d",
            "sig_id": "e65e67f0e1748ccfb4c1b05c84c87eb8311771a5d6bb908093edfb68a4e691180f",
            "payload_json": "{\"body\":{\"key\":{\"eldest_kid\":\"01202af95d3a98b97c55efd2dc09b0c9c606fd0b5eaa7e35b59a575e336dec504d490a\",\"host\":\"keybase.io\",\"kid\":\"0120d20fe982e4e7ba1a08dd60dd527a027bbe0225d4fe89eef979020c5b9efedbfb0a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"sibkey\":{\"kid\":\"012031fe75cae16e364949f99b6fc4fe79fab0b07d1234dd50b9eafb0d731571bf6c0a\",\"reverse_sig\":\"g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEgMf51yuFuNklJ+ZtvxP55+rCwfRI03VC56vsNcxVxv2wKp3BheWxvYWTFAit7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTIwMmFmOTVkM2E5OGI5N2M1NWVmZDJkYzA5YjBjOWM2MDZmZDBiNWVhYTdlMzViNTlhNTc1ZTMzNmRlYzUwNGQ0OTBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwZDIwZmU5ODJlNGU3YmExYTA4ZGQ2MGRkNTI3YTAyN2JiZTAyMjVkNGZlODllZWY5NzkwMjBjNWI5ZWZlZGJmYjBhIiwidWlkIjoiNzRjMzhjZjdjZWI5NDdmNTYzMjA0NWQ4Y2E1ZDQ4MTkiLCJ1c2VybmFtZSI6Im1heDMyIn0sInNpYmtleSI6eyJraWQiOiIwMTIwMzFmZTc1Y2FlMTZlMzY0OTQ5Zjk5YjZmYzRmZTc5ZmFiMGIwN2QxMjM0ZGQ1MGI5ZWFmYjBkNzMxNTcxYmY2YzBhIiwicmV2ZXJzZV9zaWciOm51bGx9LCJ0eXBlIjoic2lia2V5IiwidmVyc2lvbiI6MX0sImN0aW1lIjoxNDMyMTQ0OTAwLCJleHBpcmVfaW4iOjEwMDAwMDAwLCJwcmV2IjoiMjllZjIzOGE2MjgzYTkzNzMwMGExM2IzZTU1NDJhOTJlM2I5MDE0MzgwYzA3OTg5NzlhOWIyNzEzM2Q2YzhlZSIsInNlcV90eXBlIjoxLCJzZXFubyI6MywidGFnIjoic2lnbmF0dXJlIn2jc2lnxEAcpFUiEVRD2yvqIa0t93Vynjdf8lSrpVXevhsHt06LeKsR+SDXa98EJo6Ym/fjpWxJj8ejsnKBLagqEtQfEjAMqHNpZ190eXBlIKN0YWfNAgKndmVyc2lvbgE=\"},\"type\":\"sibkey\",\"version\":1},\"ctime\":1432144900,\"expire_in\":10000000,\"prev\":\"29ef238a6283a937300a13b3e5542a92e3b9014380c0798979a9b27133d6c8ee\",\"seq_type\":1,\"seqno\":3,\"tag\":\"signature\"}",
            "kid": "0120d20fe982e4e7ba1a08dd60dd527a027bbe0225d4fe89eef979020c5b9efedbfb0a",
            "ctime": 1432144900
        },
        {
            "seqno": 4,
            "prev": "c3deb587280bf27e7b7554c0304c5b58c4183d888306c65386d7a57287ba0a8d",
            "sig": "g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEgMf51yuFuNklJ+ZtvxP55+rCwfRI03VC56vsNcxVxv2wKp3BheWxvYWTFBfl7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTIwMmFmOTVkM2E5OGI5N2M1NWVmZDJkYzA5YjBjOWM2MDZmZDBiNWVhYTdlMzViNTlhNTc1ZTMzNmRlYzUwNGQ0OTBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwMzFmZTc1Y2FlMTZlMzY0OTQ5Zjk5YjZmYzRmZTc5ZmFiMGIwN2QxMjM0ZGQ1MGI5ZWFmYjBkNzMxNTcxYmY2YzBhIiwidWlkIjoiNzRjMzhjZjdjZWI5NDdmNTYzMjA0NWQ4Y2E1ZDQ4MTkiLCJ1c2VybmFtZSI6Im1heDMyIn0sInNpYmtleSI6eyJraWQiOiIwMTIwZmY1YWZmM2I1NjAyNmQ1MGY2MjI1ZmY1YzUzZmVkOWJiOWMyN2U1MjMwYWJlNjM2YzVlMDNiZTFjODQxZGQ3NjBhIiwicmV2ZXJzZV9zaWciOiJnNlJpYjJSNWhxaGtaWFJoWTJobFpNT3BhR0Z6YUY5MGVYQmxDcU5yWlhuRUl3RWcvMXIvTzFZQ2JWRDJJbC8xeFQvdG03bkNmbEl3cStZMnhlQTc0Y2hCM1hZS3AzQmhlV3h2WVdURkFpdDdJbUp2WkhraU9uc2lhMlY1SWpwN0ltVnNaR1Z6ZEY5cmFXUWlPaUl3TVRJd01tRm1PVFZrTTJFNU9HSTVOMk0xTldWbVpESmtZekE1WWpCak9XTTJNRFptWkRCaU5XVmhZVGRsTXpWaU5UbGhOVGMxWlRNek5tUmxZelV3TkdRME9UQmhJaXdpYUc5emRDSTZJbXRsZVdKaGMyVXVhVzhpTENKcmFXUWlPaUl3TVRJd016Rm1aVGMxWTJGbE1UWmxNelkwT1RRNVpqazVZalptWXpSbVpUYzVabUZpTUdJd04yUXhNak0wWkdRMU1HSTVaV0ZtWWpCa056TXhOVGN4WW1ZMll6QmhJaXdpZFdsa0lqb2lOelJqTXpoalpqZGpaV0k1TkRkbU5UWXpNakEwTldRNFkyRTFaRFE0TVRraUxDSjFjMlZ5Ym1GdFpTSTZJbTFoZURNeUluMHNJbk5wWW10bGVTSTZleUpyYVdRaU9pSXdNVEl3Wm1ZMVlXWm1NMkkxTmpBeU5tUTFNR1kyTWpJMVptWTFZelV6Wm1Wa09XSmlPV015TjJVMU1qTXdZV0psTmpNMll6VmxNRE5pWlRGak9EUXhaR1EzTmpCaElpd2ljbVYyWlhKelpWOXphV2NpT201MWJHeDlMQ0owZVhCbElqb2ljMmxpYTJWNUlpd2lkbVZ5YzJsdmJpSTZNWDBzSW1OMGFXMWxJam94TkRNeU1UUTFNREF3TENKbGVIQnBjbVZmYVc0aU9qRXdNREF3TURBd0xDSndjbVYySWpvaVl6TmtaV0kxT0RjeU9EQmlaakkzWlRkaU56VTFOR013TXpBMFl6VmlOVGhqTkRFNE0yUTRPRGd6TURaak5qVXpPRFprTjJFMU56STROMkpoTUdFNFpDSXNJbk5sY1Y5MGVYQmxJam94TENKelpYRnVieUk2TkN3aWRHRm5Jam9pYzJsbmJtRjBkWEpsSW4yamMybG54RUNzSzBoU3V6UVRKVS9STjNNNzYyRzJ0N2NrSk1DRGtVSzdqQTZyZzNwRnJveTBYRVp3TGhQQWlNODdtelVWVEJRSGU5VFVIbXJFb2Y5YWVoVTg2NElEcUhOcFoxOTBlWEJsSUtOMFlXZk5BZ0tuZG1WeWMybHZiZ0U9In0sInR5cGUiOiJzaWJrZXkiLCJ2ZXJzaW9uIjoxfSwiY3RpbWUiOjE0MzIxNDUwMDAsImV4cGlyZV9pbiI6MTAwMDAwMDAsInByZXYiOiJjM2RlYjU4NzI4MGJmMjdlN2I3NTU0YzAzMDRjNWI1OGM0MTgzZDg4ODMwNmM2NTM4NmQ3YTU3Mjg3YmEwYThkIiwic2VxX3R5cGUiOjEsInNlcW5vIjo0LCJ0YWciOiJzaWduYXR1cmUifaNzaWfEQHKgrodQ97nLF0vdV+Tvde8zto29t6XWwGbOeNvAri3JGVCCH5U0Pab/4RkZt5aSvv+rP2/AWlJqiA+ea/pikA2oc2lnX3R5cGUgo3RhZ80CAqd2ZXJzaW9uAQ==",
            "payload_hash": "e89c2dc5a08333fe90d9af7dcb5ec251feb72f87daf44b281ffc9645bb49f025",
            "sig_id": "506bf48c368a84c9e9ab6fc22438359241ee12784c72f2e68c2afa1d3dd71aae0f",
            "payload_json": "{\"body\":{\"key\":{\"eldest_kid\":\"01202af95d3a98b97c55efd2dc09b0c9c606fd0b5eaa7e35b59a575e336dec504d490a\",\"host\":\"keybase.io\",\"kid\":\"012031fe75cae16e364949f99b6fc4fe79fab0b07d1234dd50b9eafb0d731571bf6c0a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"sibkey\":{\"kid\":\"0120ff5aff3b56026d50f6225ff5c53fed9bb9c27e5230abe636c5e03be1c841dd760a\",\"reverse_sig\":\"g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEg/1r/O1YCbVD2Il/1xT/tm7nCflIwq+Y2xeA74chB3XYKp3BheWxvYWTFAit7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTIwMmFmOTVkM2E5OGI5N2M1NWVmZDJkYzA5YjBjOWM2MDZmZDBiNWVhYTdlMzViNTlhNTc1ZTMzNmRlYzUwNGQ0OTBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwMzFmZTc1Y2FlMTZlMzY0OTQ5Zjk5YjZmYzRmZTc5ZmFiMGIwN2QxMjM0ZGQ1MGI5ZWFmYjBkNzMxNTcxYmY2YzBhIiwidWlkIjoiNzRjMzhjZjdjZWI5NDdmNTYzMjA0NWQ4Y2E1ZDQ4MTkiLCJ1c2VybmFtZSI6Im1heDMyIn0sInNpYmtleSI6eyJraWQiOiIwMTIwZmY1YWZmM2I1NjAyNmQ1MGY2MjI1ZmY1YzUzZmVkOWJiOWMyN2U1MjMwYWJlNjM2YzVlMDNiZTFjODQxZGQ3NjBhIiwicmV2ZXJzZV9zaWciOm51bGx9LCJ0eXBlIjoic2lia2V5IiwidmVyc2lvbiI6MX0sImN0aW1lIjoxNDMyMTQ1MDAwLCJleHBpcmVfaW4iOjEwMDAwMDAwLCJwcmV2IjoiYzNkZWI1ODcyODBiZjI3ZTdiNzU1NGMwMzA0YzViNThjNDE4M2Q4ODgzMDZjNjUzODZkN2E1NzI4N2JhMGE4ZCIsInNlcV90eXBlIjoxLCJzZXFubyI6NCwidGFnIjoic2lnbmF0dXJlIn2jc2lnxECsK0hSuzQTJU/RN3M762G2t7ckJMCDkUK7jA6rg3pFroy0XEZwLhPAiM87mzUVTBQHe9TUHmrEof9aehU864IDqHNpZ190eXBlIKN0YWfNAgKndmVyc2lvbgE=\"},\"type\":\"sibkey\",\"version\":1},\"ctime\":1432145000,\"expire_in\":10000000,\"prev\":\"c3deb587280bf27e7b7554c0304c5b58c4183d888306c65386d7a57287ba0a8d\",\"seq_type\":1,\"seqno\":4,\"tag\":\"signature\"}",
            "kid": "012031fe75cae16e364949f99b6fc4fe79fab0b07d1234dd50b9eafb0d731571bf6c0a",
            "ctime": 1432145000
        },
        {
            "seqno": 5,
            "prev": "e89c2dc5a08333fe90d9af7dcb5ec251feb72f87daf44b281ffc9645bb49f025",
            "sig": "g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEg/1r/O1YCbVD2Il/1xT/tm7nCflIwq+Y2xeA74chB3XYKp3BheWxvYWTFBfl7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTIwMmFmOTVkM2E5OGI5N2M1NWVmZDJkYzA5YjBjOWM2MDZmZDBiNWVhYTdlMzViNTlhNTc1ZTMzNmRlYzUwNGQ0OTBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwZmY1YWZmM2I1NjAyNmQ1MGY2MjI1ZmY1YzUzZmVkOWJiOWMyN2U1MjMwYWJlNjM2YzVlMDNiZTFjODQxZGQ3NjBhIiwidWlkIjoiNzRjMzhjZjdjZWI5NDdmNTYzMjA0NWQ4Y2E1ZDQ4MTkiLCJ1c2VybmFtZSI6Im1heDMyIn0sInNpYmtleSI6eyJraWQiOiIwMTIwODE4YjYzODA3NGQwYjE1ZTc2NmFmN2E2MWMwOGZiNWZkMjMxNGNmMGUyZDBkNWExMDRkODM0N2QyMWU2ZjVlZDBhIiwicmV2ZXJzZV9zaWciOiJnNlJpYjJSNWhxaGtaWFJoWTJobFpNT3BhR0Z6YUY5MGVYQmxDcU5yWlhuRUl3RWdnWXRqZ0hUUXNWNTJhdmVtSEFqN1g5SXhUUERpME5XaEJOZzBmU0htOWUwS3AzQmhlV3h2WVdURkFpdDdJbUp2WkhraU9uc2lhMlY1SWpwN0ltVnNaR1Z6ZEY5cmFXUWlPaUl3TVRJd01tRm1PVFZrTTJFNU9HSTVOMk0xTldWbVpESmtZekE1WWpCak9XTTJNRFptWkRCaU5XVmhZVGRsTXpWaU5UbGhOVGMxWlRNek5tUmxZelV3TkdRME9UQmhJaXdpYUc5emRDSTZJbXRsZVdKaGMyVXVhVzhpTENKcmFXUWlPaUl3TVRJd1ptWTFZV1ptTTJJMU5qQXlObVExTUdZMk1qSTFabVkxWXpVelptVmtPV0ppT1dNeU4yVTFNak13WVdKbE5qTTJZelZsTUROaVpURmpPRFF4WkdRM05qQmhJaXdpZFdsa0lqb2lOelJqTXpoalpqZGpaV0k1TkRkbU5UWXpNakEwTldRNFkyRTFaRFE0TVRraUxDSjFjMlZ5Ym1GdFpTSTZJbTFoZURNeUluMHNJbk5wWW10bGVTSTZleUpyYVdRaU9pSXdNVEl3T0RFNFlqWXpPREEzTkdRd1lqRTFaVGMyTm1GbU4yRTJNV013T0daaU5XWmtNak14TkdObU1HVXlaREJrTldFeE1EUmtPRE0wTjJReU1XVTJaalZsWkRCaElpd2ljbVYyWlhKelpWOXphV2NpT201MWJHeDlMQ0owZVhCbElqb2ljMmxpYTJWNUlpd2lkbVZ5YzJsdmJpSTZNWDBzSW1OMGFXMWxJam94TkRNeU1UUTFNVEF3TENKbGVIQnBjbVZmYVc0aU9qRXdNREF3TURBd0xDSndjbVYySWpvaVpUZzVZekprWXpWaE1EZ3pNek5tWlRrd1pEbGhaamRrWTJJMVpXTXlOVEZtWldJM01tWTROMlJoWmpRMFlqSTRNV1ptWXprMk5EVmlZalE1WmpBeU5TSXNJbk5sY1Y5MGVYQmxJam94TENKelpYRnVieUk2TlN3aWRHRm5Jam9pYzJsbmJtRjBkWEpsSW4yamMybG54RUJONGREU2tscENRVlRnU1NiZkZNUEVHRk8rOVpGMHNtbE4wbkhSYkVNUDFIc09nbnVUZ0hhcXRCdEFJMGZPM09oWkl6N21LVkFlQUpESTI0dERyUjBNcUhOcFoxOTBlWEJsSUtOMFlXZk5BZ0tuZG1WeWMybHZiZ0U9In0sInR5cGUiOiJzaWJrZXkiLCJ2ZXJzaW9uIjoxfSwiY3RpbWUiOjE0MzIxNDUxMDAsImV4cGlyZV9pbiI6MTAwMDAwMDAsInByZXYiOiJlODljMmRjNWEwODMzM2ZlOTBkOWFmN2RjYjVlYzI1MWZlYjcyZjg3ZGFmNDRiMjgxZmZjOTY0NWJiNDlmMDI1Iiwic2VxX3R5cGUiOjEsInNlcW5vIjo1LCJ0YWciOiJzaWduYXR1cmUifaNzaWfEQG/vX+rlHcSTSNAPh3S/pMcGrxGBYNZyQQrTvfN8etjjd0bT8VqL0NoOQgedykvgcQoOcdWnC3qiJUY/p4nucQyoc2lnX3R5cGUgo3RhZ80CAqd2ZXJzaW9uAQ==",
            "payload_hash": "33b022b996f1faaff00ad6eafdb9f5745ec0c0baa08588cab74017edeaf0f233",
            "sig_id": "b0d3a36fc061b317a260514eff8f5ac1b8ab1cfc60c54731e79e979f9e638b320f",
            "payload_json": "{\"body\":{\"key\":{\"eldest_kid\":\"01202af95d3a98b97c55efd2dc09b0c9c606fd0b5eaa7e35b59a575e336dec504d490a\",\"host\":\"keybase.io\",\"kid\":\"0120ff5aff3b56026d50f6225ff5c53fed9bb9c27e5230abe636c5e03be1c841dd760a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"sibkey\":{\"kid\":\"0120818b638074d0b15e766af7a61c08fb5fd2314cf0e2d0d5a104d8347d21e6f5ed0a\",\"reverse_sig\":\"g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEggYtjgHTQsV52avemHAj7X9IxTPDi0NWhBNg0fSHm9e0Kp3BheWxvYWTFAit7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTIwMmFmOTVkM2E5OGI5N2M1NWVmZDJkYzA5YjBjOWM2MDZmZDBiNWVhYTdlMzViNTlhNTc1ZTMzNmRlYzUwNGQ0OTBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwZmY1YWZmM2I1NjAyNmQ1MGY2MjI1ZmY1YzUzZmVkOWJiOWMyN2U1MjMwYWJlNjM2YzVlMDNiZTFjODQxZGQ3NjBhIiwidWlkIjoiNzRjMzhjZjdjZWI5NDdmNTYzMjA0NWQ4Y2E1ZDQ4MTkiLCJ1c2VybmFtZSI6Im1heDMyIn0sInNpYmtleSI6eyJraWQiOiIwMTIwODE4YjYzODA3NGQwYjE1ZTc2NmFmN2E2MWMwOGZiNWZkMjMxNGNmMGUyZDBkNWExMDRkODM0N2QyMWU2ZjVlZDBhIiwicmV2ZXJzZV9zaWciOm51bGx9LCJ0eXBlIjoic2lia2V5IiwidmVyc2lvbiI6MX0sImN0aW1lIjoxNDMyMTQ1MTAwLCJleHBpcmVfaW4iOjEwMDAwMDAwLCJwcmV2IjoiZTg5YzJkYzVhMDgzMzNmZTkwZDlhZjdkY2I1ZWMyNTFmZWI3MmY4N2RhZjQ0YjI4MWZmYzk2NDViYjQ5ZjAyNSIsInNlcV90eXBlIjoxLCJzZXFubyI6NSwidGFnIjoic2lnbmF0dXJlIn2jc2lnxEBN4dDSklpCQVTgSSbfFMPEGFO+9ZF0smlN0nHRbEMP1HsOgnuTgHaqtBtAI0fO3OhZIz7mKVAeAJDI24tDrR0MqHNpZ190eXBlIKN0YWfNAgKndmVyc2lvbgE=\"},\"type\":\"sibkey\",\"version\":1},\"ctime\":1432145100,\"expire_in\":10000000,\"prev\":\"e89c2dc5a08333fe90d9af7dcb5ec251feb72f87daf44b281ffc9645bb49f025\",\"seq_type\":1,\"seqno\":5,\"tag\":\"signature\"}",
            "kid": "0120ff5aff3b56026d50f6225ff5c53fed9bb9c27e5230abe636c5e03be1c841dd760a",
            "ctime": 1432145100
        },
        {
            "seqno": 6,
            "prev": "33b022b996f1faaff00ad6eafdb9f5745ec0c0baa08588cab74017edeaf0f233",
            "sig": "g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEggYtjgHTQsV52avemHAj7X9IxTPDi0NWhBNg0fSHm9e0Kp3BheWxvYWTFBfl7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTIwMmFmOTVkM2E5OGI5N2M1NWVmZDJkYzA5YjBjOWM2MDZmZDBiNWVhYTdlMzViNTlhNTc1ZTMzNmRlYzUwNGQ0OTBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwODE4YjYzODA3NGQwYjE1ZTc2NmFmN2E2MWMwOGZiNWZkMjMxNGNmMGUyZDBkNWExMDRkODM0N2QyMWU2ZjVlZDBhIiwidWlkIjoiNzRjMzhjZjdjZWI5NDdmNTYzMjA0NWQ4Y2E1ZDQ4MTkiLCJ1c2VybmFtZSI6Im1heDMyIn0sInNpYmtleSI6eyJraWQiOiIwMTIwZDU2ODZjZDAzNmJiMzI1M2NjMDI2Nzc1OTAzNjQ4ODBhYmI5OTQ1NzY1Nzk4Y2FjNTBhMjg1YmFjZGEzNjI2MDBhIiwicmV2ZXJzZV9zaWciOiJnNlJpYjJSNWhxaGtaWFJoWTJobFpNT3BhR0Z6YUY5MGVYQmxDcU5yWlhuRUl3RWcxV2hzMERhN01sUE1BbWQxa0RaSWdLdTVsRmRsZVl5c1VLS0Z1czJqWW1BS3AzQmhlV3h2WVdURkFpdDdJbUp2WkhraU9uc2lhMlY1SWpwN0ltVnNaR1Z6ZEY5cmFXUWlPaUl3TVRJd01tRm1PVFZrTTJFNU9HSTVOMk0xTldWbVpESmtZekE1WWpCak9XTTJNRFptWkRCaU5XVmhZVGRsTXpWaU5UbGhOVGMxWlRNek5tUmxZelV3TkdRME9UQmhJaXdpYUc5emRDSTZJbXRsZVdKaGMyVXVhVzhpTENKcmFXUWlPaUl3TVRJd09ERTRZall6T0RBM05HUXdZakUxWlRjMk5tRm1OMkUyTVdNd09HWmlOV1prTWpNeE5HTm1NR1V5WkRCa05XRXhNRFJrT0RNME4yUXlNV1UyWmpWbFpEQmhJaXdpZFdsa0lqb2lOelJqTXpoalpqZGpaV0k1TkRkbU5UWXpNakEwTldRNFkyRTFaRFE0TVRraUxDSjFjMlZ5Ym1GdFpTSTZJbTFoZURNeUluMHNJbk5wWW10bGVTSTZleUpyYVdRaU9pSXdNVEl3WkRVMk9EWmpaREF6Tm1KaU16STFNMk5qTURJMk56YzFPVEF6TmpRNE9EQmhZbUk1T1RRMU56WTFOems0WTJGak5UQmhNamcxWW1GalpHRXpOakkyTURCaElpd2ljbVYyWlhKelpWOXphV2NpT201MWJHeDlMQ0owZVhCbElqb2ljMmxpYTJWNUlpd2lkbVZ5YzJsdmJpSTZNWDBzSW1OMGFXMWxJam94TkRNeU1UUTFNakF3TENKbGVIQnBjbVZmYVc0aU9qRXdNREF3TURBd0xDSndjbVYySWpvaU16TmlNREl5WWprNU5tWXhabUZoWm1Zd01HRmtObVZoWm1SaU9XWTFOelExWldNd1l6QmlZV0V3T0RVNE9HTmhZamMwTURFM1pXUmxZV1l3WmpJek15SXNJbk5sY1Y5MGVYQmxJam94TENKelpYRnVieUk2Tml3aWRHRm5Jam9pYzJsbmJtRjBkWEpsSW4yamMybG54RUM3OVpzTDJvd2pjamFXK2hLeWplVFgwMXJFYkFlQ0pWMU5oalBKSTZQUUJPZ0tMcFFGYldzcnhsZTgyQit6ZkN5NkJGaFRsZm1yRUxpZG5WV0QrRzRBcUhOcFoxOTBlWEJsSUtOMFlXZk5BZ0tuZG1WeWMybHZiZ0U9In0sInR5cGUiOiJzaWJrZXkiLCJ2ZXJzaW9uIjoxfSwiY3RpbWUiOjE0MzIxNDUyMDAsImV4cGlyZV9pbiI6MTAwMDAwMDAsInByZXYiOiIzM2IwMjJiOTk2ZjFmYWFmZjAwYWQ2ZWFmZGI5ZjU3NDVlYzBjMGJhYTA4NTg4Y2FiNzQwMTdlZGVhZjBmMjMzIiwic2VxX3R5cGUiOjEsInNlcW5vIjo2LCJ0YWciOiJzaWduYXR1cmUifaNzaWfEQBy2uxeE6Sr8KcG9R1EEnKklC8sR8tFF3x48QqRPFPSlb8CZs1YekMVcFmjJDaANUpdokzc1dDP+gi5S61Fjkwqoc2lnX3R5cGUgo3RhZ80CAqd2ZXJzaW9uAQ==",
            "payload_hash": "81356413d54fb759fdb6f3bdbcf020c7208fba6f3635a81469f9581be2be2340",
            "sig_id": "f52a55c59023f5e188e85cf0933e0301b1143b462c7f361abcdc7989492b1bf00f",
            "payload_json": "{\"body\":{\"key\":{\"eldest_kid\":\"01202af95d3a98b97c55efd2dc09b0c9c606fd0b5eaa7e35b59a575e336dec504d490a\",\"host\":\"keybase.io\",\"kid\":\"0120818b638074d0b15e766af7a61c08fb5fd2314cf0e2d0d5a104d8347d21e6f5ed0a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"sibkey\":{\"kid\":\"0120d5686cd036bb3253cc02677590364880abb9945765798cac50a285bacda362600a\",\"reverse_sig\":\"g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEg1Whs0Da7MlPMAmd1kDZIgKu5lFdleYysUKKFus2jYmAKp3BheWxvYWTFAit7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTIwMmFmOTVkM2E5OGI5N2M1NWVmZDJkYzA5YjBjOWM2MDZmZDBiNWVhYTdlMzViNTlhNTc1ZTMzNmRlYzUwNGQ0OTBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwODE4YjYzODA3NGQwYjE1ZTc2NmFmN2E2MWMwOGZiNWZkMjMxNGNmMGUyZDBkNWExMDRkODM0N2QyMWU2ZjVlZDBhIiwidWlkIjoiNzRjMzhjZjdjZWI5NDdmNTYzMjA0NWQ4Y2E1ZDQ4MTkiLCJ1c2VybmFtZSI6Im1heDMyIn0sInNpYmtleSI6eyJraWQiOiIwMTIwZDU2ODZjZDAzNmJiMzI1M2NjMDI2Nzc1OTAzNjQ4ODBhYmI5OTQ1NzY1Nzk4Y2FjNTBhMjg1YmFjZGEzNjI2MDBhIiwicmV2ZXJzZV9zaWciOm51bGx9LCJ0eXBlIjoic2lia2V5IiwidmVyc2lvbiI6MX0sImN0aW1lIjoxNDMyMTQ1MjAwLCJleHBpcmVfaW4iOjEwMDAwMDAwLCJwcmV2IjoiMzNiMDIyYjk5NmYxZmFhZmYwMGFkNmVhZmRiOWY1NzQ1ZWMwYzBiYWEwODU4OGNhYjc0MDE3ZWRlYWYwZjIzMyIsInNlcV90eXBlIjoxLCJzZXFubyI6NiwidGFnIjoic2lnbmF0dXJlIn2jc2lnxEC79ZsL2owjcjaW+hKyjeTX01rEbAeCJV1NhjPJI6PQBOgKLpQFbWsrxle82B+zfCy6BFhTlfmrELidnVWD+G4AqHNpZ190eXBlIKN0YWfNAgKndmVyc2lvbgE=\"},\"type\":\"sibkey\",\"version\":1},\"ctime\":1432145200,\"expire_in\":10000000,\"prev\":\"33b022b996f1faaff00ad6eafdb9f5745ec0c0baa08588cab74017edeaf0f233\",\"seq_type\":1,\"seqno\":6,\"tag\":\"signature\"}",
            "kid": "0120818b638074d0b15e766af7a61c08fb5fd2314cf0e2d0d5a104d8347d21e6f5ed0a",
            "ctime": 1432145200
        },
        {
            "seqno": 7,
            "prev": "81356413d54fb759fdb6f3bdbcf020c7208fba6f3635a81469f9581be2be2340",
            "sig": "g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEg1Whs0Da7MlPMAmd1kDZIgKu5lFdleYysUKKFus2jYmAKp3BheWxvYWTFAoF7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTIwMmFmOTVkM2E5OGI5N2M1NWVmZDJkYzA5YjBjOWM2MDZmZDBiNWVhYTdlMzViNTlhNTc1ZTMzNmRlYzUwNGQ0OTBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwZDU2ODZjZDAzNmJiMzI1M2NjMDI2Nzc1OTAzNjQ4ODBhYmI5OTQ1NzY1Nzk4Y2FjNTBhMjg1YmFjZGEzNjI2MDBhIiwidWlkIjoiNzRjMzhjZjdjZWI5NDdmNTYzMjA0NWQ4Y2E1ZDQ4MTkiLCJ1c2VybmFtZSI6Im1heDMyIn0sInN1YmtleSI6eyJraWQiOiIwMTIxMzBiZjU3Y2MwMTY4ZjJhZGQzYTYxZGI1ZmNiOTNkMTlmN2MwZTQyM2EyNjZmZDVkNWE5OWQ1NmQ4MWMzNWM1NTBhIiwicGFyZW50X2tpZCI6IjAxMjBkNTY4NmNkMDM2YmIzMjUzY2MwMjY3NzU5MDM2NDg4MGFiYjk5NDU3NjU3OThjYWM1MGEyODViYWNkYTM2MjYwMGEiLCJyZXZlcnNlX3NpZyI6bnVsbH0sInR5cGUiOiJzdWJrZXkiLCJ2ZXJzaW9uIjoxfSwiY3RpbWUiOjE0MzIxNDUzMDAsImV4cGlyZV9pbiI6MTAwMDAwMDAsInByZXYiOiI4MTM1NjQxM2Q1NGZiNzU5ZmRiNmYzYmRiY2YwMjBjNzIwOGZiYTZmMzYzNWE4MTQ2OWY5NTgxYmUyYmUyMzQwIiwic2VxX3R5cGUiOjEsInNlcW5vIjo3LCJ0YWciOiJzaWduYXR1cmUifaNzaWfEQOhVHMGds7KBnZCG0IkbF7E8nEpHYhoFI2s54a69Cl8LikH6JdcCkAkFrHM7Oeu2gOFIHk5/w3XU7z51n+DKBw6oc2lnX3R5cGUgo3RhZ80CAqd2ZXJzaW9uAQ==",
            "payload_hash": "ba371fdc2fc138718b7183204c92a2ae0eb1f7b3717f457a68ac7add0ff34bcb",
            "sig_id": "81824a207a2a3ecbc958074167346990f0f5e61fd4b0ebd22b7c7b1b0987b7920f",
            "payload_json": "{\"body\":{\"key\":{\"eldest_kid\":\"01202af95d3a98b97c55efd2dc09b0c9c606fd0b5eaa7e35b59a575e336dec504d490a\",\"host\":\"keybase.io\",\"kid\":\"0120d5686cd036bb3253cc02677590364880abb9945765798cac50a285bacda362600a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"subkey\":{\"kid\":\"012130bf57cc0168f2add3a61db5fcb93d19f7c0e423a266fd5d5a99d56d81c35c550a\",\"parent_kid\":\"0120d5686cd036bb3253cc02677590364880abb9945765798cac50a285bacda362600a\",\"reverse_sig\":null},\"type\":\"subkey\",\"version\":1},\"ctime\":1432145300,\"expire_in\":10000000,\"prev\":\"81356413d54fb759fdb6f3bdbcf020c7208fba6f3635a81469f9581be2be2340\",\"seq_type\":1,\"seqno\":7,\"tag\":\"signature\"}",
            "kid": "0120d5686cd036bb3253cc02677590364880abb9945765798cac50a285bacda362600a",
            "ctime": 1432145300
        },
        {
            "seqno": 8,
            "prev": "ba371fdc2fc138718b7183204c92a2ae0eb1f7b3717f457a68ac7add0ff34bcb",
            "sig": "g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEg1Whs0Da7MlPMAmd1kDZIgKu5lFdleYysUKKFus2jYmAKp3BheWxvYWTFAoF7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTIwMmFmOTVkM2E5OGI5N2M1NWVmZDJkYzA5YjBjOWM2MDZmZDBiNWVhYTdlMzViNTlhNTc1ZTMzNmRlYzUwNGQ0OTBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwZDU2ODZjZDAzNmJiMzI1M2NjMDI2Nzc1OTAzNjQ4ODBhYmI5OTQ1NzY1Nzk4Y2FjNTBhMjg1YmFjZGEzNjI2MDBhIiwidWlkIjoiNzRjMzhjZjdjZWI5NDdmNTYzMjA0NWQ4Y2E1ZDQ4MTkiLCJ1c2VybmFtZSI6Im1heDMyIn0sInN1YmtleSI6eyJraWQiOiIwMTIxNTYwMWM0YzRiNDZlMTlmMzU2NTRhMWRlYmM4ZGMwZTNmNjgwMTNjMzUxN2ViODA5ZTVmMzY1NDE3MjZmZDMzNjBhIiwicGFyZW50X2tpZCI6IjAxMjBkNTY4NmNkMDM2YmIzMjUzY2MwMjY3NzU5MDM2NDg4MGFiYjk5NDU3NjU3OThjYWM1MGEyODViYWNkYTM2MjYwMGEiLCJyZXZlcnNlX3NpZyI6bnVsbH0sInR5cGUiOiJzdWJrZXkiLCJ2ZXJzaW9uIjoxfSwiY3RpbWUiOjE0MzIxNDU0MDAsImV4cGlyZV9pbiI6MTAwMDAwMDAsInByZXYiOiJiYTM3MWZkYzJmYzEzODcxOGI3MTgzMjA0YzkyYTJhZTBlYjFmN2IzNzE3ZjQ1N2E2OGFjN2FkZDBmZjM0YmNiIiwic2VxX3R5cGUiOjEsInNlcW5vIjo4LCJ0YWciOiJzaWduYXR1cmUifaNzaWfEQMUM7CpuUEq39bXVQE8Idxp7Ezolki56TrLsua40gmrzdj7bTFmj0YeEqRegn50dxqgAzoi/eqECouYHAdxb7Amoc2lnX3R5cGUgo3RhZ80CAqd2ZXJzaW9uAQ==",
            "payload_hash": "f28bad7570323729cbfbb3cd945e03b4ab9957450006fc339e297f3788add3c0",
            "sig_id": "a535a3972951c2077647d567949c73209672afd5c586fe7a0f4407d5894099000f",
            "payload_json": "{\"body\":{\"key\":{\"eldest_kid\":\"01202af95d3a98b97c55efd2dc09b0c9c606fd0b5eaa7e35b59a575e336dec504d490a\",\"host\":\"keybase.io\",\"kid\":\"0120d5686cd036bb3253cc02677590364880abb9945765798cac50a285bacda362600a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"subkey\":{\"kid\":\"01215601c4c4b46e19f35654a1debc8dc0e3f68013c3517eb809e5f36541726fd3360a\",\"parent_kid\":\"0120d5686cd036bb3253cc02677590364880abb9945765798cac50a285bacda362600a\",\"reverse_sig\":null},\"type\":\"subkey\",\"version\":1},\"ctime\":1432145400,\"expire_in\":10000000,\"prev\":\"ba371fdc2fc138718b7183204c92a2ae0eb1f7b3717f457a68ac7add0ff34bcb\",\"seq_type\":1,\"seqno\":8,\"tag\":\"signature\"}",
            "kid": "0120d5686cd036bb3253cc02677590364880abb9945765798cac50a285bacda362600a",
            "ctime": 1432145400
        },
        {
            "seqno": 9,
            "prev": "f28bad7570323729cbfbb3cd945e03b4ab9957450006fc339e297f3788add3c0",
            "sig": "g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEg1Whs0Da7MlPMAmd1kDZIgKu5lFdleYysUKKFus2jYmAKp3BheWxvYWTFAoF7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTIwMmFmOTVkM2E5OGI5N2M1NWVmZDJkYzA5YjBjOWM2MDZmZDBiNWVhYTdlMzViNTlhNTc1ZTMzNmRlYzUwNGQ0OTBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwZDU2ODZjZDAzNmJiMzI1M2NjMDI2Nzc1OTAzNjQ4ODBhYmI5OTQ1NzY1Nzk4Y2FjNTBhMjg1YmFjZGEzNjI2MDBhIiwidWlkIjoiNzRjMzhjZjdjZWI5NDdmNTYzMjA0NWQ4Y2E1ZDQ4MTkiLCJ1c2VybmFtZSI6Im1heDMyIn0sInN1YmtleSI6eyJraWQiOiIwMTIxNWNiOTAyZTRiNmEwNzgzZjQ1NThlNGI0ODNmMDkxODUyZWI1ODVhZWNkZWMwMTkzNTZlYjgxM2FjNTA0MGQwMjBhIiwicGFyZW50X2tpZCI6IjAxMjBkNTY4NmNkMDM2YmIzMjUzY2MwMjY3NzU5MDM2NDg4MGFiYjk5NDU3NjU3OThjYWM1MGEyODViYWNkYTM2MjYwMGEiLCJyZXZlcnNlX3NpZyI6bnVsbH0sInR5cGUiOiJzdWJrZXkiLCJ2ZXJzaW9uIjoxfSwiY3RpbWUiOjE0MzIxNDU1MDAsImV4cGlyZV9pbiI6MTAwMDAwMDAsInByZXYiOiJmMjhiYWQ3NTcwMzIzNzI5Y2JmYmIzY2Q5NDVlMDNiNGFiOTk1NzQ1MDAwNmZjMzM5ZTI5N2YzNzg4YWRkM2MwIiwic2VxX3R5cGUiOjEsInNlcW5vIjo5LCJ0YWciOiJzaWduYXR1cmUifaNzaWfEQJEPySgBxHJeapOmXGKSEDP7H2EL4pB78+oM6j22Qa5CaEIzC1NPE/nGlpb3a1/iVxORoMWLohakYoHEgJVYOAaoc2lnX3R5cGUgo3RhZ80CAqd2ZXJzaW9uAQ==",
            "payload_hash": "d56391bc51ce01b9c6851edd75d7b5e8d8b3c7622c5d7600b71437075adaa232",
            "sig_id": "d5f5e75c809a6c1207d35791dd8361c68b5a7a0de9853d9ee647ba3f1620b9e00f",
            "payload_json": "{\"body\":{\"key\":{\"eldest_kid\":\"01202af95d3a98b97c55efd2dc09b0c9c606fd0b5eaa7e35b59a575e336dec504d490a\",\"host\":\"keybase.io\",\"kid\":\"0120d5686cd036bb3253cc02677590364880abb9945765798cac50a285bacda362600a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"subkey\":{\"kid\":\"01215cb902e4b6a0783f4558e4b483f091852eb585aecdec019356eb813ac5040d020a\",\"parent_kid\":\"0120d5686cd036bb3253cc02677590364880abb9945765798cac50a285bacda362600a\",\"reverse_sig\":null},\"type\":\"subkey\",\"version\":1},\"ctime\":1432145500,\"expire_in\":10000000,\"prev\":\"f28bad7570323729cbfbb3cd945e03b4ab9957450006fc339e297f3788add3c0\",\"seq_type\":1,\"seqno\":9,\"tag\":\"signature\"}",
            "kid": "0120d5686cd036bb3253cc02677590364880abb9945765798cac50a285bacda362600a",
            "ctime": 1432145500
        },
        {
            "seqno": 10,
            "prev": "d56391bc51ce01b9c6851edd75d7b5e8d8b3c7622c5d7600b71437075adaa232",
            "sig": "g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEg1Whs0Da7MlPMAmd1kDZIgKu5lFdleYysUKKFus2jYmAKp3BheWxvYWTFAhl7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTIwMmFmOTVkM2E5OGI5N2M1NWVmZDJkYzA5YjBjOWM2MDZmZDBiNWVhYTdlMzViNTlhNTc1ZTMzNmRlYzUwNGQ0OTBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwZDU2ODZjZDAzNmJiMzI1M2NjMDI2Nzc1OTAzNjQ4ODBhYmI5OTQ1NzY1Nzk4Y2FjNTBhMjg1YmFjZGEzNjI2MDBhIiwidWlkIjoiNzRjMzhjZjdjZWI5NDdmNTYzMjA0NWQ4Y2E1ZDQ4MTkiLCJ1c2VybmFtZSI6Im1heDMyIn0sInJldm9rZSI6eyJraWQiOiIwMTIwZDIwZmU5ODJlNGU3YmExYTA4ZGQ2MGRkNTI3YTAyN2JiZTAyMjVkNGZlODllZWY5NzkwMjBjNWI5ZWZlZGJmYjBhIn0sInR5cGUiOiJyZXZva2UiLCJ2ZXJzaW9uIjoxfSwiY3RpbWUiOjE0MzIxNDU2MDAsImV4cGlyZV9pbiI6MTAwMDAwMDAsInByZXYiOiJkNTYzOTFiYzUxY2UwMWI5YzY4NTFlZGQ3NWQ3YjVlOGQ4YjNjNzYyMmM1ZDc2MDBiNzE0MzcwNzVhZGFhMjMyIiwic2VxX3R5cGUiOjEsInNlcW5vIjoxMCwidGFnIjoic2lnbmF0dXJlIn2jc2lnxECvZgjAO0Re95cU0cwiw2IMaomncIU6ir6lbQup6QM789tiDoed42Xf9+vKXB+lyHmNQwPOt6Qd/cYHafREYR8HqHNpZ190eXBlIKN0YWfNAgKndmVyc2lvbgE=",
            "payload_hash": "601b3462742a14955c12017a7bbf5615cb3cad8bea04734d00dd61d492ad1952",
            "sig_id": "4d4a63b6914fc2750ff9e03f7f5f54ea62cfaf91c17d4311b0991fbedefbff760f",
            "payload_json": "{\"body\":{\"key\":{\"eldest_kid\":\"01202af95d3a98b97c55efd2dc09b0c9c606fd0b5eaa7e35b59a575e336dec504d490a\",\"host\":\"keybase.io\",\"kid\":\"0120d5686cd036bb3253cc02677590364880abb9945765798cac50a285bacda362600a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"revoke\":{\"kid\":\"0120d20fe982e4e7ba1a08dd60dd527a027bbe0225d4fe89eef979020c5b9efedbfb0a\"},\"type\":\"revoke\",\"version\":1},\"ctime\":1432145600,\"expire_in\":10000000,\"prev\":\"d56391bc51ce01b9c6851edd75d7b5e8d8b3c7622c5d7600b71437075adaa232\",\"seq_type\":1,\"seqno\":10,\"tag\":\"signature\"}",
            "kid": "0120d5686cd036bb3253cc02677590364880abb9945765798cac50a285bacda362600a",
            "ctime": 1432145600
        },
        {
            "seqno": 11,
            "prev": "601b3462742a14955c12017a7bbf5615cb3cad8bea04734d00dd61d492ad1952",
            "sig": "g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEg1Whs0Da7MlPMAmd1kDZIgKu5lFdleYysUKKFus2jYmAKp3BheWxvYWTFAhl7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTIwMmFmOTVkM2E5OGI5N2M1NWVmZDJkYzA5YjBjOWM2MDZmZDBiNWVhYTdlMzViNTlhNTc1ZTMzNmRlYzUwNGQ0OTBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwZDU2ODZjZDAzNmJiMzI1M2NjMDI2Nzc1OTAzNjQ4ODBhYmI5OTQ1NzY1Nzk4Y2FjNTBhMjg1YmFjZGEzNjI2MDBhIiwidWlkIjoiNzRjMzhjZjdjZWI5NDdmNTYzMjA0NWQ4Y2E1ZDQ4MTkiLCJ1c2VybmFtZSI6Im1heDMyIn0sInJldm9rZSI6eyJraWQiOiIwMTIwZDIwZmU5ODJlNGU3YmExYTA4ZGQ2MGRkNTI3YTAyN2JiZTAyMjVkNGZlODllZWY5NzkwMjBjNWI5ZWZlZGJmYjBhIn0sInR5cGUiOiJyZXZva2UiLCJ2ZXJzaW9uIjoxfSwiY3RpbWUiOjE0MzIxNDU3MDAsImV4cGlyZV9pbiI6MTAwMDAwMDAsInByZXYiOiI2MDFiMzQ2Mjc0MmExNDk1NWMxMjAxN2E3YmJmNTYxNWNiM2NhZDhiZWEwNDczNGQwMGRkNjFkNDkyYWQxOTUyIiwic2VxX3R5cGUiOjEsInNlcW5vIjoxMSwidGFnIjoic2lnbmF0dXJlIn2jc2lnxEALFHOFpo85gpD2rMddo8E1U6ZxP3wRFOyj9eoxDGFxdtwwNbWCNb8bO0U1MsaiMFJCd/WOyg2NEHEcSR+5stsKqHNpZ190eXBlIKN0YWfNAgKndmVyc2lvbgE=",
            "payload_hash": "dce3f0a140b4549038a21b43aec620d5362c1ee99b462cc2c8d7c8efbf4cf233",
            "sig_id": "be6da4ea06918c7f5f6255834c9e41da9b94c75745edbe64659fe885c3c432a10f",
            "payload_json": "{\"body\":{\"key\":{\"eldest_kid\":\"01202af95d3a98b97c55efd2dc09b0c9c606fd0b5eaa7e35b59a575e336dec504d490a\",\"host\":\"keybase.io\",\"kid\":\"0120d5686cd036bb3253cc02677590364880abb9945765798cac50a285bacda362600a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"revoke\":{\"kid\":\"0120d20fe982e4e7ba1a08dd60dd527a027bbe0225d4fe89eef979020c5b9efedbfb0a\"},\"type\":\"revoke\",\"version\":1},\"ctime\":1432145700,\"expire_in\":10000000,\"prev\":\"601b3462742a14955c12017a7bbf5615cb3cad8bea04734d00dd61d492ad1952\",\"seq_type\":1,\"seqno\":11,\"tag\":\"signature\"}",
            "kid": "0120d5686cd036bb3253cc02677590364880abb9945765798cac50a285bacda362600a",
            "ctime": 1432145700
        },
        {
            "seqno": 12,
            "prev": "dce3f0a140b4549038a21b43aec620d5362c1ee99b462cc2c8d7c8efbf4cf233",
            "sig": "g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEg1Whs0Da7MlPMAmd1kDZIgKu5lFdleYysUKKFus2jYmAKp3BheWxvYWTFAhx7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTIwMmFmOTVkM2E5OGI5N2M1NWVmZDJkYzA5YjBjOWM2MDZmZDBiNWVhYTdlMzViNTlhNTc1ZTMzNmRlYzUwNGQ0OTBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwZDU2ODZjZDAzNmJiMzI1M2NjMDI2Nzc1OTAzNjQ4ODBhYmI5OTQ1NzY1Nzk4Y2FjNTBhMjg1YmFjZGEzNjI2MDBhIiwidWlkIjoiNzRjMzhjZjdjZWI5NDdmNTYzMjA0NWQ4Y2E1ZDQ4MTkiLCJ1c2VybmFtZSI6Im1heDMyIn0sInJldm9rZSI6eyJraWRzIjpbIjAxMjAzMWZlNzVjYWUxNmUzNjQ5NDlmOTliNmZjNGZlNzlmYWIwYjA3ZDEyMzRkZDUwYjllYWZiMGQ3MzE1NzFiZjZjMGEiXX0sInR5cGUiOiJyZXZva2UiLCJ2ZXJzaW9uIjoxfSwiY3RpbWUiOjE0MzIxNDU4MDAsImV4cGlyZV9pbiI6MTAwMDAwMDAsInByZXYiOiJkY2UzZjBhMTQwYjQ1NDkwMzhhMjFiNDNhZWM2MjBkNTM2MmMxZWU5OWI0NjJjYzJjOGQ3YzhlZmJmNGNmMjMzIiwic2VxX3R5cGUiOjEsInNlcW5vIjoxMiwidGFnIjoic2lnbmF0dXJlIn2jc2lnxEBw0vZNoLLGZ+77z68uKIPVtTY9L+LB/Le7Qy+8ViETnpz1SMN8TFzNVtoTZvtCp8lUVzYqhCCGCczNkT6BfJkJqHNpZ190eXBlIKN0YWfNAgKndmVyc2lvbgE=",
            "payload_hash": "8689d6aa08e1a0818f955de82520fa736f61ced7eb640070c45660e127d63bda",
            "sig_id": "cb6682dcddf0e407635260320b21957c2df306f81813a5a7c8507936b296c5150f",
            "payload_json": "{\"body\":{\"key\":{\"eldest_kid\":\"01202af95d3a98b97c55efd2dc09b0c9c606fd0b5eaa7e35b59a575e336dec504d490a\",\"host\":\"keybase.io\",\"kid\":\"0120d5686cd036bb3253cc02677590364880abb9945765798cac50a285bacda362600a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"revoke\":{\"kids\":[\"012031fe75cae16e364949f99b6fc4fe79fab0b07d1234dd50b9eafb0d731571bf6c0a\"]},\"type\":\"revoke\",\"version\":1},\"ctime\":1432145800,\"expire_in\":10000000,\"prev\":\"dce3f0a140b4549038a21b43aec620d5362c1ee99b462cc2c8d7c8efbf4cf233\",\"seq_type\":1,\"seqno\":12,\"tag\":\"signature\"}",
            "kid": "0120d5686cd036bb3253cc02677590364880abb9945765798cac50a285bacda362600a",
            "ctime": 1432145800
        },
        {
            "seqno": 13,
            "prev": "8689d6aa08e1a0818f955de82520fa736f61ced7eb640070c45660e127d63bda",
            "sig": "g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEg1Whs0Da7MlPMAmd1kDZIgKu5lFdleYysUKKFus2jYmAKp3BheWxvYWTFAhh7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTIwMmFmOTVkM2E5OGI5N2M1NWVmZDJkYzA5YjBjOWM2MDZmZDBiNWVhYTdlMzViNTlhNTc1ZTMzNmRlYzUwNGQ0OTBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwZDU2ODZjZDAzNmJiMzI1M2NjMDI2Nzc1OTAzNjQ4ODBhYmI5OTQ1NzY1Nzk4Y2FjNTBhMjg1YmFjZGEzNjI2MDBhIiwidWlkIjoiNzRjMzhjZjdjZWI5NDdmNTYzMjA0NWQ4Y2E1ZDQ4MTkiLCJ1c2VybmFtZSI6Im1heDMyIn0sInJldm9rZSI6eyJzaWdfaWQiOiI1MDZiZjQ4YzM2OGE4NGM5ZTlhYjZmYzIyNDM4MzU5MjQxZWUxMjc4NGM3MmYyZTY4YzJhZmExZDNkZDcxYWFlMGYifSwidHlwZSI6InJldm9rZSIsInZlcnNpb24iOjF9LCJjdGltZSI6MTQzMjE0NTkwMCwiZXhwaXJlX2luIjoxMDAwMDAwMCwicHJldiI6Ijg2ODlkNmFhMDhlMWEwODE4Zjk1NWRlODI1MjBmYTczNmY2MWNlZDdlYjY0MDA3MGM0NTY2MGUxMjdkNjNiZGEiLCJzZXFfdHlwZSI6MSwic2Vxbm8iOjEzLCJ0YWciOiJzaWduYXR1cmUifaNzaWfEQAVTx2OScgQI1tbC07Cg1VfRW4Zow+p7PFX3g4HCck5ZBjf/2wsj1JYihQZgmDnm8IFVdmUEXNui5F2MVxvYSweoc2lnX3R5cGUgo3RhZ80CAqd2ZXJzaW9uAQ==",
            "payload_hash": "2a8200e9f3f59f2db54f404a51a727784571586282b7786755a17196ec2f722e",
            "sig_id": "33214539afe47c134f16b520728da821a8f0f0259e954805d11a0ea285309f700f",
            "payload_json": "{\"body\":{\"key\":{\"eldest_kid\":\"01202af95d3a98b97c55efd2dc09b0c9c606fd0b5eaa7e35b59a575e336dec504d490a\",\"host\":\"keybase.io\",\"kid\":\"0120d5686cd036bb3253cc02677590364880abb9945765798cac50a285bacda362600a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"revoke\":{\"sig_id\":\"506bf48c368a84c9e9ab6fc22438359241ee12784c72f2e68c2afa1d3dd71aae0f\"},\"type\":\"revoke\",\"version\":1},\"ctime\":1432145900,\"expire_in\":10000000,\"prev\":\"8689d6aa08e1a0818f955de82520fa736f61ced7eb640070c45660e127d63bda\",\"seq_type\":1,\"seqno\":13,\"tag\":\"signature\"}",
            "kid": "0120d5686cd036bb3253cc02677590364880abb9945765798cac50a285bacda362600a",
            "ctime": 1432145900
        },
        {
            "seqno": 14,
            "prev": "2a8200e9f3f59f2db54f404a51a727784571586282b7786755a17196ec2f722e",
            "sig": "g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEg1Whs0Da7MlPMAmd1kDZIgKu5lFdleYysUKKFus2jYmAKp3BheWxvYWTFAhh7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTIwMmFmOTVkM2E5OGI5N2M1NWVmZDJkYzA5YjBjOWM2MDZmZDBiNWVhYTdlMzViNTlhNTc1ZTMzNmRlYzUwNGQ0OTBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwZDU2ODZjZDAzNmJiMzI1M2NjMDI2Nzc1OTAzNjQ4ODBhYmI5OTQ1NzY1Nzk4Y2FjNTBhMjg1YmFjZGEzNjI2MDBhIiwidWlkIjoiNzRjMzhjZjdjZWI5NDdmNTYzMjA0NWQ4Y2E1ZDQ4MTkiLCJ1c2VybmFtZSI6Im1heDMyIn0sInJldm9rZSI6eyJzaWdfaWQiOiI1MDZiZjQ4YzM2OGE4NGM5ZTlhYjZmYzIyNDM4MzU5MjQxZWUxMjc4NGM3MmYyZTY4YzJhZmExZDNkZDcxYWFlMGYifSwidHlwZSI6InJldm9rZSIsInZlcnNpb24iOjF9LCJjdGltZSI6MTQzMjE0NjAwMCwiZXhwaXJlX2luIjoxMDAwMDAwMCwicHJldiI6IjJhODIwMGU5ZjNmNTlmMmRiNTRmNDA0YTUxYTcyNzc4NDU3MTU4NjI4MmI3Nzg2NzU1YTE3MTk2ZWMyZjcyMmUiLCJzZXFfdHlwZSI6MSwic2Vxbm8iOjE0LCJ0YWciOiJzaWduYXR1cmUifaNzaWfEQK2ipVbmehRpRKz8yoBDxNlxipjLFEfSFZpbz3HBB7pOCSQKcEQP9jDn4Das48z08CS6Bb59ZUD+hlYjI6joWQ+oc2lnX3R5cGUgo3RhZ80CAqd2ZXJzaW9uAQ==",
            "payload_hash": "d0ff71f5026ff1a92b8baec3b06336df8d46f961077f39442fb019553b2b0a18",
            "sig_id": "67d4abca611b70d909cbfa3fcf3311320791a57309f54799a6e2ae49db858a9e0f",
            "payload_json": "{\"body\":{\"key\":{\"eldest_kid\":\"01202af95d3a98b97c55efd2dc09b0c9c606fd0b5eaa7e35b59a575e336dec504d490a\",\"host\":\"keybase.io\",\"kid\":\"0120d5686cd036bb3253cc02677590364880abb9945765798cac50a285bacda362600a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"revoke\":{\"sig_id\":\"506bf48c368a84c9e9ab6fc22438359241ee12784c72f2e68c2afa1d3dd71aae0f\"},\"type\":\"revoke\",\"version\":1},\"ctime\":1432146000,\"expire_in\":10000000,\"prev\":\"2a8200e9f3f59f2db54f404a51a727784571586282b7786755a17196ec2f722e\",\"seq_type\":1,\"seqno\":14,\"tag\":\"signature\"}",
            "kid": "0120d5686cd036bb3253cc02677590364880abb9945765798cac50a285bacda362600a",
            "ctime": 1432146000
        },
        {
            "seqno": 15,
            "prev": "d0ff71f5026ff1a92b8baec3b06336df8d46f961077f39442fb019553b2b0a18",
            "sig": "g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEg1Whs0Da7MlPMAmd1kDZIgKu5lFdleYysUKKFus2jYmAKp3BheWxvYWTFAmB7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTIwMmFmOTVkM2E5OGI5N2M1NWVmZDJkYzA5YjBjOWM2MDZmZDBiNWVhYTdlMzViNTlhNTc1ZTMzNmRlYzUwNGQ0OTBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwZDU2ODZjZDAzNmJiMzI1M2NjMDI2Nzc1OTAzNjQ4ODBhYmI5OTQ1NzY1Nzk4Y2FjNTBhMjg1YmFjZGEzNjI2MDBhIiwidWlkIjoiNzRjMzhjZjdjZWI5NDdmNTYzMjA0NWQ4Y2E1ZDQ4MTkiLCJ1c2VybmFtZSI6Im1heDMyIn0sInJldm9rZSI6eyJzaWdfaWRzIjpbIjc3YTJhNzBkNWM4YWYwZjEyMDEwMjcwNzY1MDQwNjVlODBkZGQ2ZDVmZjJhNjIyZWI1N2UxNDlkYzhmYTAwNDcwZiIsImIwZDNhMzZmYzA2MWIzMTdhMjYwNTE0ZWZmOGY1YWMxYjhhYjFjZmM2MGM1NDczMWU3OWU5NzlmOWU2MzhiMzIwZiJdfSwidHlwZSI6InJldm9rZSIsInZlcnNpb24iOjF9LCJjdGltZSI6MTQzMjE0NjEwMCwiZXhwaXJlX2luIjoxMDAwMDAwMCwicHJldiI6ImQwZmY3MWY1MDI2ZmYxYTkyYjhiYWVjM2IwNjMzNmRmOGQ0NmY5NjEwNzdmMzk0NDJmYjAxOTU1M2IyYjBhMTgiLCJzZXFfdHlwZSI6MSwic2Vxbm8iOjE1LCJ0YWciOiJzaWduYXR1cmUifaNzaWfEQH9HYHx8/DWAwcC5rRKcNQ+Bwk//XSc196Nr89E6dS1XrKYGurIL0+wqWvfRn1NxLT/zSvcwAXT/OoEu5PtHDguoc2lnX3R5cGUgo3RhZ80CAqd2ZXJzaW9uAQ==",
            "payload_hash": "159fd5164720001518c6a6d194c755ed6bca0eeedce5ac3f409f67fe629e655d",
            "sig_id": "577d4cc5654ced3ed4740b35eac9312701bb46177a68249bd44fec4a274669680f",
            "payload_json": "{\"body\":{\"key\":{\"eldest_kid\":\"01202af95d3a98b97c55efd2dc09b0c9c606fd0b5eaa7e35b59a575e336dec504d490a\",\"host\":\"keybase.io\",\"kid\":\"0120d5686cd036bb3253cc02677590364880abb9945765798cac50a285bacda362600a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"revoke\":{\"sig_ids\":[\"77a2a70d5c8af0f1201027076504065e80ddd6d5ff2a622eb57e149dc8fa00470f\",\"b0d3a36fc061b317a260514eff8f5ac1b8ab1cfc60c54731e79e979f9e638b320f\"]},\"type\":\"revoke\",\"version\":1},\"ctime\":1432146100,\"expire_in\":10000000,\"prev\":\"d0ff71f5026ff1a92b8baec3b06336df8d46f961077f39442fb019553b2b0a18\",\"seq_type\":1,\"seqno\":15,\"tag\":\"signature\"}",
            "kid": "0120d5686cd036bb3253cc02677590364880abb9945765798cac50a285bacda362600a",
            "ctime": 1432146100
        },
        {
            "seqno": 16,
            "prev": "159fd5164720001518c6a6d194c755ed6bca0eeedce5ac3f409f67fe629e655d",
            "sig": "g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEgKvldOpi5fFXv0twJsMnGBv0LXqp+NbWaV14zbexQTUkKp3BheWxvYWTFAhl7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTIwMmFmOTVkM2E5OGI5N2M1NWVmZDJkYzA5YjBjOWM2MDZmZDBiNWVhYTdlMzViNTlhNTc1ZTMzNmRlYzUwNGQ0OTBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwMmFmOTVkM2E5OGI5N2M1NWVmZDJkYzA5YjBjOWM2MDZmZDBiNWVhYTdlMzViNTlhNTc1ZTMzNmRlYzUwNGQ0OTBhIiwidWlkIjoiNzRjMzhjZjdjZWI5NDdmNTYzMjA0NWQ4Y2E1ZDQ4MTkiLCJ1c2VybmFtZSI6Im1heDMyIn0sInJldm9rZSI6eyJraWQiOiIwMTIxMzBiZjU3Y2MwMTY4ZjJhZGQzYTYxZGI1ZmNiOTNkMTlmN2MwZTQyM2EyNjZmZDVkNWE5OWQ1NmQ4MWMzNWM1NTBhIn0sInR5cGUiOiJyZXZva2UiLCJ2ZXJzaW9uIjoxfSwiY3RpbWUiOjE0MzIxNDYyMDAsImV4cGlyZV9pbiI6MTAwMDAwMDAsInByZXYiOiIxNTlmZDUxNjQ3MjAwMDE1MThjNmE2ZDE5NGM3NTVlZDZiY2EwZWVlZGNlNWFjM2Y0MDlmNjdmZTYyOWU2NTVkIiwic2VxX3R5cGUiOjEsInNlcW5vIjoxNiwidGFnIjoic2lnbmF0dXJlIn2jc2lnxEBRTqKeefnwZwlcVwjun1IsWKFhkP3xNdSliTIq/YVuuBOS4Iw66xtvShzdfwrqUR569G4mcU0AkLV320hWVVQGqHNpZ190eXBlIKN0YWfNAgKndmVyc2lvbgE=",
            "payload_hash": "ba47575973c4e7f3faf6c6c3ee346a8f5d7a1c37f5c03bc0df7e15c5603774fd",
            "sig_id": "00ce703cbc1ec4576392064e378aaf5fd78983a934e4cfae97af3d6d8231743e0f",
            "payload_json": "{\"body\":{\"key\":{\"eldest_kid\":\"01202af95d3a98b97c55efd2dc09b0c9c606fd0b5eaa7e35b59a575e336dec504d490a\",\"host\":\"keybase.io\",\"kid\":\"01202af95d3a98b97c55efd2dc09b0c9c606fd0b5eaa7e35b59a575e336dec504d490a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"revoke\":{\"kid\":\"012130bf57cc0168f2add3a61db5fcb93d19f7c0e423a266fd5d5a99d56d81c35c550a\"},\"type\":\"revoke\",\"version\":1},\"ctime\":1432146200,\"expire_in\":10000000,\"prev\":\"159fd5164720001518c6a6d194c755ed6bca0eeedce5ac3f409f67fe629e655d\",\"seq_type\":1,\"seqno\":16,\"tag\":\"signature\"}",
            "kid": "01202af95d3a98b97c55efd2dc09b0c9c606fd0b5eaa7e35b59a575e336dec504d490a",
            "ctime": 1432146200
        },
        {
            "seqno": 17,
            "prev": "ba47575973c4e7f3faf6c6c3ee346a8f5d7a1c37f5c03bc0df7e15c5603774fd",
            "sig": "g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEgKvldOpi5fFXv0twJsMnGBv0LXqp+NbWaV14zbexQTUkKp3BheWxvYWTFAhh7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTIwMmFmOTVkM2E5OGI5N2M1NWVmZDJkYzA5YjBjOWM2MDZmZDBiNWVhYTdlMzViNTlhNTc1ZTMzNmRlYzUwNGQ0OTBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwMmFmOTVkM2E5OGI5N2M1NWVmZDJkYzA5YjBjOWM2MDZmZDBiNWVhYTdlMzViNTlhNTc1ZTMzNmRlYzUwNGQ0OTBhIiwidWlkIjoiNzRjMzhjZjdjZWI5NDdmNTYzMjA0NWQ4Y2E1ZDQ4MTkiLCJ1c2VybmFtZSI6Im1heDMyIn0sInJldm9rZSI6eyJzaWdfaWQiOiJhNTM1YTM5NzI5NTFjMjA3NzY0N2Q1Njc5NDljNzMyMDk2NzJhZmQ1YzU4NmZlN2EwZjQ0MDdkNTg5NDA5OTAwMGYifSwidHlwZSI6InJldm9rZSIsInZlcnNpb24iOjF9LCJjdGltZSI6MTQzMjE0NjMwMCwiZXhwaXJlX2luIjoxMDAwMDAwMCwicHJldiI6ImJhNDc1NzU5NzNjNGU3ZjNmYWY2YzZjM2VlMzQ2YThmNWQ3YTFjMzdmNWMwM2JjMGRmN2UxNWM1NjAzNzc0ZmQiLCJzZXFfdHlwZSI6MSwic2Vxbm8iOjE3LCJ0YWciOiJzaWduYXR1cmUifaNzaWfEQKXqUkdDSab896qNZKjMQ9sF3n3SbDS66JvGc9vFZ7/sBXNjTlt2SWRy40ukK3VIGE10D6QT8DBiEeSz5ffAmwaoc2lnX3R5cGUgo3RhZ80CAqd2ZXJzaW9uAQ==",
            "payload_hash": "ae4ed6b838a8395e9f44f7e1e5fbad442134de72691dcfc5dc04eb9b47da00f2",
            "sig_id": "ee77d80410f106b29f3b6294ae5de84308166188ba01e64d7ce3e1628044ce380f",
            "payload_json": "{\"body\":{\"key\":{\"eldest_kid\":\"01202af95d3a98b97c55efd2dc09b0c9c606fd0b5eaa7e35b59a575e336dec504d490a\",\"host\":\"keybase.io\",\"kid\":\"01202af95d3a98b97c55efd2dc09b0c9c606fd0b5eaa7e35b59a575e336dec504d490a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"revoke\":{\"sig_id\":\"a535a3972951c2077647d567949c73209672afd5c586fe7a0f4407d5894099000f\"},\"type\":\"revoke\",\"version\":1},\"ctime\":1432146300,\"expire_in\":10000000,\"prev\":\"ba47575973c4e7f3faf6c6c3ee346a8f5d7a1c37f5c03bc0df7e15c5603774fd\",\"seq_type\":1,\"seqno\":17,\"tag\":\"signature\"}",
            "kid": "01202af95d3a98b97c55efd2dc09b0c9c606fd0b5eaa7e35b59a575e336dec504d490a",
            "ctime": 1432146300
        }
    ],
    "keys": [
        "01202af95d3a98b97c55efd2dc09b0c9c606fd0b5eaa7e35b59a575e336dec504d490a",
        "0120d20fe982e4e7ba1a08dd60dd527a027bbe0225d4fe89eef979020c5b9efedbfb0a",
        "012031fe75cae16e364949f99b6fc4fe79fab0b07d1234dd50b9eafb0d731571bf6c0a",
        "0120ff5aff3b56026d50f6225ff5c53fed9bb9c27e5230abe636c5e03be1c841dd760a",
        "0120818b638074d0b15e766af7a61c08fb5fd2314cf0e2d0d5a104d8347d21e6f5ed0a",
        "0120d5686cd036bb3253cc02677590364880abb9945765798cac50a285bacda362600a",
        "012130bf57cc0168f2add3a61db5fcb93d19f7c0e423a266fd5d5a99d56d81c35c550a",
        "01215601c4c4b46e19f35654a1debc8dc0e3f68013c3517eb809e5f36541726fd3360a",
        "01215cb902e4b6a0783f4558e4b483f091852eb585aecdec019356eb813ac5040d020a"
    ],
    "uid": "74c38cf7ceb947f5632045d8ca5d4819",
    "username": "max32"
}
`,
	"expired_key_chain.json": `{
    "chain": [
        {
            "seqno": 1,
            "prev": null,
            "sig": "g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEgDTTfqXtH50CdFdtB6s5oGZiTHBiSB8+PU6MnAafmgr4Kp3BheWxvYWTFASN7ImJvZHkiOnsia2V5Ijp7Imhvc3QiOiJrZXliYXNlLmlvIiwia2lkIjoiMDEyMDBkMzRkZmE5N2I0N2U3NDA5ZDE1ZGI0MWVhY2U2ODE5OTg5MzFjMTg5MjA3Y2Y4ZjUzYTMyNzAxYTdlNjgyYmUwYSIsInVpZCI6Ijc0YzM4Y2Y3Y2ViOTQ3ZjU2MzIwNDVkOGNhNWQ0ODE5IiwidXNlcm5hbWUiOiJtYXgzMiJ9LCJ0eXBlIjoiZWxkZXN0IiwidmVyc2lvbiI6MX0sImN0aW1lIjoxNDMyMTQ0NzUwLCJleHBpcmVfaW4iOjEsInByZXYiOm51bGwsInNlcV90eXBlIjoxLCJzZXFubyI6MSwidGFnIjoic2lnbmF0dXJlIn2jc2lnxEAi+mx5hSdvAHor5WhocvcOzK0M/JjQcB6/6BNzIm7DHXopkvVGK6qXrg6QvydIoZ0OW8lD23eWrKdIHaEITJUOqHNpZ190eXBlIKN0YWfNAgKndmVyc2lvbgE=",
            "payload_hash": "cc2f840119fc2734c852c41c3274b69e53fd6771f15a31b2c06bb57c6a9f1752",
            "sig_id": "bdc3c02698b224ba0e27ed3c9782106ecc83ffa5c5a436e438b3de7dbc42554e0f",
            "payload_json": "{\"body\":{\"key\":{\"host\":\"keybase.io\",\"kid\":\"01200d34dfa97b47e7409d15db41eace681998931c189207cf8f53a32701a7e682be0a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"type\":\"eldest\",\"version\":1},\"ctime\":1432144750,\"expire_in\":1,\"prev\":null,\"seq_type\":1,\"seqno\":1,\"tag\":\"signature\"}",
            "kid": "01200d34dfa97b47e7409d15db41eace681998931c189207cf8f53a32701a7e682be0a",
            "ctime": 1432144750
        },
        {
            "seqno": 2,
            "prev": "cc2f840119fc2734c852c41c3274b69e53fd6771f15a31b2c06bb57c6a9f1752",
            "sig": "g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEgDTTfqXtH50CdFdtB6s5oGZiTHBiSB8+PU6MnAafmgr4Kp3BheWxvYWTFBep7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTIwMGQzNGRmYTk3YjQ3ZTc0MDlkMTVkYjQxZWFjZTY4MTk5ODkzMWMxODkyMDdjZjhmNTNhMzI3MDFhN2U2ODJiZTBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwMGQzNGRmYTk3YjQ3ZTc0MDlkMTVkYjQxZWFjZTY4MTk5ODkzMWMxODkyMDdjZjhmNTNhMzI3MDFhN2U2ODJiZTBhIiwidWlkIjoiNzRjMzhjZjdjZWI5NDdmNTYzMjA0NWQ4Y2E1ZDQ4MTkiLCJ1c2VybmFtZSI6Im1heDMyIn0sInNpYmtleSI6eyJraWQiOiIwMTIwNjM1NTc5NDlhY2NlMzI0NWNlNzQwNjM2MzJlOWRkMmE2MWNjYzAzNTBlOWJlMTAwMDEwN2E4Y2VlMGQyMTMxODBhIiwicmV2ZXJzZV9zaWciOiJnNlJpYjJSNWhxaGtaWFJoWTJobFpNT3BhR0Z6YUY5MGVYQmxDcU5yWlhuRUl3RWdZMVY1U2F6T01rWE9kQVkyTXVuZEttSE13RFVPbStFQUFRZW96dURTRXhnS3AzQmhlV3h2WVdURkFpUjdJbUp2WkhraU9uc2lhMlY1SWpwN0ltVnNaR1Z6ZEY5cmFXUWlPaUl3TVRJd01HUXpOR1JtWVRrM1lqUTNaVGMwTURsa01UVmtZalF4WldGalpUWTRNVGs1T0Rrek1XTXhPRGt5TURkalpqaG1OVE5oTXpJM01ERmhOMlUyT0RKaVpUQmhJaXdpYUc5emRDSTZJbXRsZVdKaGMyVXVhVzhpTENKcmFXUWlPaUl3TVRJd01HUXpOR1JtWVRrM1lqUTNaVGMwTURsa01UVmtZalF4WldGalpUWTRNVGs1T0Rrek1XTXhPRGt5TURkalpqaG1OVE5oTXpJM01ERmhOMlUyT0RKaVpUQmhJaXdpZFdsa0lqb2lOelJqTXpoalpqZGpaV0k1TkRkbU5UWXpNakEwTldRNFkyRTFaRFE0TVRraUxDSjFjMlZ5Ym1GdFpTSTZJbTFoZURNeUluMHNJbk5wWW10bGVTSTZleUpyYVdRaU9pSXdNVEl3TmpNMU5UYzVORGxoWTJObE16STBOV05sTnpRd05qTTJNekpsT1dSa01tRTJNV05qWXpBek5UQmxPV0psTVRBd01ERXdOMkU0WTJWbE1HUXlNVE14T0RCaElpd2ljbVYyWlhKelpWOXphV2NpT201MWJHeDlMQ0owZVhCbElqb2ljMmxpYTJWNUlpd2lkbVZ5YzJsdmJpSTZNWDBzSW1OMGFXMWxJam94TkRNeU1UUTBPRFV3TENKbGVIQnBjbVZmYVc0aU9qRXNJbkJ5WlhZaU9pSmpZekptT0RRd01URTVabU15TnpNMFl6ZzFNbU0wTVdNek1qYzBZalk1WlRVelptUTJOemN4WmpFMVlUTXhZakpqTURaaVlqVTNZelpoT1dZeE56VXlJaXdpYzJWeFgzUjVjR1VpT2pFc0luTmxjVzV2SWpveUxDSjBZV2NpT2lKemFXZHVZWFIxY21VaWZhTnphV2ZFUUg3aU00dnRmeDlQUWdteFgzZVhmQmVQNEJodEErQ0d1ZEtvR01ZMGZFZGo0TEZQbjk4RE9Od1YvTWUrUU56WFlma01xWnQ4UnAvRVMySHp4cjdxZXcyb2MybG5YM1I1Y0dVZ28zUmhaODBDQXFkMlpYSnphVzl1QVE9PSJ9LCJ0eXBlIjoic2lia2V5IiwidmVyc2lvbiI6MX0sImN0aW1lIjoxNDMyMTQ0ODUwLCJleHBpcmVfaW4iOjEsInByZXYiOiJjYzJmODQwMTE5ZmMyNzM0Yzg1MmM0MWMzMjc0YjY5ZTUzZmQ2NzcxZjE1YTMxYjJjMDZiYjU3YzZhOWYxNzUyIiwic2VxX3R5cGUiOjEsInNlcW5vIjoyLCJ0YWciOiJzaWduYXR1cmUifaNzaWfEQJywhknmAyuC+7rqV/WdOLVOXDrcZJaVVI/W+fTmJVbLJKA/HnK2uqrTscFIGYVnnFZpGaYT0cg5LQ8t2HUDIweoc2lnX3R5cGUgo3RhZ80CAqd2ZXJzaW9uAQ==",
            "payload_hash": "d945735df7741a001ceeee778c79fc4d3dfd55d2fa5d6b672b7bd198395ab746",
            "sig_id": "79579a28083e801d565a330e3c8a7d0cd76dda85f07603976ccb016f259cfadf0f",
            "payload_json": "{\"body\":{\"key\":{\"eldest_kid\":\"01200d34dfa97b47e7409d15db41eace681998931c189207cf8f53a32701a7e682be0a\",\"host\":\"keybase.io\",\"kid\":\"01200d34dfa97b47e7409d15db41eace681998931c189207cf8f53a32701a7e682be0a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"sibkey\":{\"kid\":\"012063557949acce3245ce74063632e9dd2a61ccc0350e9be1000107a8cee0d213180a\",\"reverse_sig\":\"g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEgY1V5SazOMkXOdAY2MundKmHMwDUOm+EAAQeozuDSExgKp3BheWxvYWTFAiR7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTIwMGQzNGRmYTk3YjQ3ZTc0MDlkMTVkYjQxZWFjZTY4MTk5ODkzMWMxODkyMDdjZjhmNTNhMzI3MDFhN2U2ODJiZTBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwMGQzNGRmYTk3YjQ3ZTc0MDlkMTVkYjQxZWFjZTY4MTk5ODkzMWMxODkyMDdjZjhmNTNhMzI3MDFhN2U2ODJiZTBhIiwidWlkIjoiNzRjMzhjZjdjZWI5NDdmNTYzMjA0NWQ4Y2E1ZDQ4MTkiLCJ1c2VybmFtZSI6Im1heDMyIn0sInNpYmtleSI6eyJraWQiOiIwMTIwNjM1NTc5NDlhY2NlMzI0NWNlNzQwNjM2MzJlOWRkMmE2MWNjYzAzNTBlOWJlMTAwMDEwN2E4Y2VlMGQyMTMxODBhIiwicmV2ZXJzZV9zaWciOm51bGx9LCJ0eXBlIjoic2lia2V5IiwidmVyc2lvbiI6MX0sImN0aW1lIjoxNDMyMTQ0ODUwLCJleHBpcmVfaW4iOjEsInByZXYiOiJjYzJmODQwMTE5ZmMyNzM0Yzg1MmM0MWMzMjc0YjY5ZTUzZmQ2NzcxZjE1YTMxYjJjMDZiYjU3YzZhOWYxNzUyIiwic2VxX3R5cGUiOjEsInNlcW5vIjoyLCJ0YWciOiJzaWduYXR1cmUifaNzaWfEQH7iM4vtfx9PQgmxX3eXfBeP4BhtA+CGudKoGMY0fEdj4LFPn98DONwV/Me+QNzXYfkMqZt8Rp/ES2Hzxr7qew2oc2lnX3R5cGUgo3RhZ80CAqd2ZXJzaW9uAQ==\"},\"type\":\"sibkey\",\"version\":1},\"ctime\":1432144850,\"expire_in\":1,\"prev\":\"cc2f840119fc2734c852c41c3274b69e53fd6771f15a31b2c06bb57c6a9f1752\",\"seq_type\":1,\"seqno\":2,\"tag\":\"signature\"}",
            "kid": "01200d34dfa97b47e7409d15db41eace681998931c189207cf8f53a32701a7e682be0a",
            "ctime": 1432144850
        }
    ],
    "keys": [
        "01200d34dfa97b47e7409d15db41eace681998931c189207cf8f53a32701a7e682be0a",
        "012063557949acce3245ce74063632e9dd2a61ccc0350e9be1000107a8cee0d213180a"
    ],
    "uid": "74c38cf7ceb947f5632045d8ca5d4819",
    "username": "max32"
}
`,
	"mismatched_ctime_chain.json": `{
    "chain": [
        {
            "seqno": 1,
            "prev": null,
            "sig": "g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEgYnFfW7DX7muUv2yy4q0mOPrRSbfaann5TLJr0k5nYK0Kp3BheWxvYWTFASp7ImJvZHkiOnsia2V5Ijp7Imhvc3QiOiJrZXliYXNlLmlvIiwia2lkIjoiMDEyMDYyNzE1ZjViYjBkN2VlNmI5NGJmNmNiMmUyYWQyNjM4ZmFkMTQ5YjdkYTZhNzlmOTRjYjI2YmQyNGU2NzYwYWQwYSIsInVpZCI6Ijc0YzM4Y2Y3Y2ViOTQ3ZjU2MzIwNDVkOGNhNWQ0ODE5IiwidXNlcm5hbWUiOiJtYXgzMiJ9LCJ0eXBlIjoiZWxkZXN0IiwidmVyc2lvbiI6MX0sImN0aW1lIjoxNDMyMTQ0ODkyLCJleHBpcmVfaW4iOjEwMDAwMDAwLCJwcmV2IjpudWxsLCJzZXFfdHlwZSI6MSwic2Vxbm8iOjEsInRhZyI6InNpZ25hdHVyZSJ9o3NpZ8RA8+M6gsewyY6d9b1C6LL4wtWqU8mGvPdRLltLCIYAsNkswHR7A5s7PtNqu43Inw6daGP/2FE0GqkSPoyiAaaoDKhzaWdfdHlwZSCjdGFnzQICp3ZlcnNpb24B",
            "payload_hash": "38e015983046ec66fe1e05f79afa4fb794c007d884238d67fa0d61404a7dc3c0",
            "sig_id": "428cf6551603a81e245dbd3831d962639600411fdf56a0a11e52192e567debc50f",
            "payload_json": "{\"body\":{\"key\":{\"host\":\"keybase.io\",\"kid\":\"012062715f5bb0d7ee6b94bf6cb2e2ad2638fad149b7da6a79f94cb26bd24e6760ad0a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"type\":\"eldest\",\"version\":1},\"ctime\":1432144892,\"expire_in\":10000000,\"prev\":null,\"seq_type\":1,\"seqno\":1,\"tag\":\"signature\"}",
            "kid": "012062715f5bb0d7ee6b94bf6cb2e2ad2638fad149b7da6a79f94cb26bd24e6760ad0a",
            "ctime": 1432144892
        },
        {
            "seqno": 2,
            "prev": "38e015983046ec66fe1e05f79afa4fb794c007d884238d67fa0d61404a7dc3c0",
            "sig": "g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEgYnFfW7DX7muUv2yy4q0mOPrRSbfaann5TLJr0k5nYK0Kp3BheWxvYWTFBfl7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTIwNjI3MTVmNWJiMGQ3ZWU2Yjk0YmY2Y2IyZTJhZDI2MzhmYWQxNDliN2RhNmE3OWY5NGNiMjZiZDI0ZTY3NjBhZDBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwNjI3MTVmNWJiMGQ3ZWU2Yjk0YmY2Y2IyZTJhZDI2MzhmYWQxNDliN2RhNmE3OWY5NGNiMjZiZDI0ZTY3NjBhZDBhIiwidWlkIjoiNzRjMzhjZjdjZWI5NDdmNTYzMjA0NWQ4Y2E1ZDQ4MTkiLCJ1c2VybmFtZSI6Im1heDMyIn0sInNpYmtleSI6eyJraWQiOiIwMTIwMDNmYTc2MTYwMTEwNjkzZmM2MTNjZWQ1NGVlNjdkNmE1MTlkYjZjYmFlOWY0ZDg2NjRjMzYwOWMwNzRmM2E3MjBhIiwicmV2ZXJzZV9zaWciOiJnNlJpYjJSNWhxaGtaWFJoWTJobFpNT3BhR0Z6YUY5MGVYQmxDcU5yWlhuRUl3RWdBL3AyRmdFUWFUL0dFODdWVHVaOWFsR2R0c3V1bjAyR1pNTmduQWRQT25JS3AzQmhlV3h2WVdURkFpdDdJbUp2WkhraU9uc2lhMlY1SWpwN0ltVnNaR1Z6ZEY5cmFXUWlPaUl3TVRJd05qSTNNVFZtTldKaU1HUTNaV1UyWWprMFltWTJZMkl5WlRKaFpESTJNemhtWVdReE5EbGlOMlJoTm1FM09XWTVOR05pTWpaaVpESTBaVFkzTmpCaFpEQmhJaXdpYUc5emRDSTZJbXRsZVdKaGMyVXVhVzhpTENKcmFXUWlPaUl3TVRJd05qSTNNVFZtTldKaU1HUTNaV1UyWWprMFltWTJZMkl5WlRKaFpESTJNemhtWVdReE5EbGlOMlJoTm1FM09XWTVOR05pTWpaaVpESTBaVFkzTmpCaFpEQmhJaXdpZFdsa0lqb2lOelJqTXpoalpqZGpaV0k1TkRkbU5UWXpNakEwTldRNFkyRTFaRFE0TVRraUxDSjFjMlZ5Ym1GdFpTSTZJbTFoZURNeUluMHNJbk5wWW10bGVTSTZleUpyYVdRaU9pSXdNVEl3TURObVlUYzJNVFl3TVRFd05qa3pabU0yTVROalpXUTFOR1ZsTmpka05tRTFNVGxrWWpaalltRmxPV1kwWkRnMk5qUmpNell3T1dNd056Um1NMkUzTWpCaElpd2ljbVYyWlhKelpWOXphV2NpT201MWJHeDlMQ0owZVhCbElqb2ljMmxpYTJWNUlpd2lkbVZ5YzJsdmJpSTZNWDBzSW1OMGFXMWxJam94TkRNeU1UUTBPVGt5TENKbGVIQnBjbVZmYVc0aU9qRXdNREF3TURBd0xDSndjbVYySWpvaU16aGxNREUxT1Rnek1EUTJaV00yTm1abE1XVXdOV1kzT1dGbVlUUm1ZamM1TkdNd01EZGtPRGcwTWpNNFpEWTNabUV3WkRZeE5EQTBZVGRrWXpOak1DSXNJbk5sY1Y5MGVYQmxJam94TENKelpYRnVieUk2TWl3aWRHRm5Jam9pYzJsbmJtRjBkWEpsSW4yamMybG54RUIxTHVaQi9KRGZZQVZVbjdrZGM1aEVkZU1RSHFqMzh3T0Z5aEF0VUVKNWV3NG5vM2VWSXI3Z2o5cW5CL2dQUHdNb2paeUxuaUgzRmFKYjVsdkNEVWNJcUhOcFoxOTBlWEJsSUtOMFlXZk5BZ0tuZG1WeWMybHZiZ0U9In0sInR5cGUiOiJzaWJrZXkiLCJ2ZXJzaW9uIjoxfSwiY3RpbWUiOjE0MzIxNDQ5OTIsImV4cGlyZV9pbiI6MTAwMDAwMDAsInByZXYiOiIzOGUwMTU5ODMwNDZlYzY2ZmUxZTA1Zjc5YWZhNGZiNzk0YzAwN2Q4ODQyMzhkNjdmYTBkNjE0MDRhN2RjM2MwIiwic2VxX3R5cGUiOjEsInNlcW5vIjoyLCJ0YWciOiJzaWduYXR1cmUifaNzaWfEQLs0S2APzYjYogGfzk71zXPFZCQIaWRutp7hbGhRFb8E+O1lgf3rq0rPY026KfwDmC+XVuWP4fx6tRiT3Za1nQyoc2lnX3R5cGUgo3RhZ80CAqd2ZXJzaW9uAQ==",
            "payload_hash": "b25f5bc2a42b746d18a63500434639c793b187b58ddb1b5f917f2f19e56e984d",
            "sig_id": "b884f8210ddbf8ea7b649ef262fb61b8f4ac746918f3e0de0815d01a33b73f800f",
            "payload_json": "{\"body\":{\"key\":{\"eldest_kid\":\"012062715f5bb0d7ee6b94bf6cb2e2ad2638fad149b7da6a79f94cb26bd24e6760ad0a\",\"host\":\"keybase.io\",\"kid\":\"012062715f5bb0d7ee6b94bf6cb2e2ad2638fad149b7da6a79f94cb26bd24e6760ad0a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"sibkey\":{\"kid\":\"012003fa76160110693fc613ced54ee67d6a519db6cbae9f4d8664c3609c074f3a720a\",\"reverse_sig\":\"g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEgA/p2FgEQaT/GE87VTuZ9alGdtsuun02GZMNgnAdPOnIKp3BheWxvYWTFAit7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTIwNjI3MTVmNWJiMGQ3ZWU2Yjk0YmY2Y2IyZTJhZDI2MzhmYWQxNDliN2RhNmE3OWY5NGNiMjZiZDI0ZTY3NjBhZDBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwNjI3MTVmNWJiMGQ3ZWU2Yjk0YmY2Y2IyZTJhZDI2MzhmYWQxNDliN2RhNmE3OWY5NGNiMjZiZDI0ZTY3NjBhZDBhIiwidWlkIjoiNzRjMzhjZjdjZWI5NDdmNTYzMjA0NWQ4Y2E1ZDQ4MTkiLCJ1c2VybmFtZSI6Im1heDMyIn0sInNpYmtleSI6eyJraWQiOiIwMTIwMDNmYTc2MTYwMTEwNjkzZmM2MTNjZWQ1NGVlNjdkNmE1MTlkYjZjYmFlOWY0ZDg2NjRjMzYwOWMwNzRmM2E3MjBhIiwicmV2ZXJzZV9zaWciOm51bGx9LCJ0eXBlIjoic2lia2V5IiwidmVyc2lvbiI6MX0sImN0aW1lIjoxNDMyMTQ0OTkyLCJleHBpcmVfaW4iOjEwMDAwMDAwLCJwcmV2IjoiMzhlMDE1OTgzMDQ2ZWM2NmZlMWUwNWY3OWFmYTRmYjc5NGMwMDdkODg0MjM4ZDY3ZmEwZDYxNDA0YTdkYzNjMCIsInNlcV90eXBlIjoxLCJzZXFubyI6MiwidGFnIjoic2lnbmF0dXJlIn2jc2lnxEB1LuZB/JDfYAVUn7kdc5hEdeMQHqj38wOFyhAtUEJ5ew4no3eVIr7gj9qnB/gPPwMojZyLniH3FaJb5lvCDUcIqHNpZ190eXBlIKN0YWfNAgKndmVyc2lvbgE=\"},\"type\":\"sibkey\",\"version\":1},\"ctime\":1432144992,\"expire_in\":10000000,\"prev\":\"38e015983046ec66fe1e05f79afa4fb794c007d884238d67fa0d61404a7dc3c0\",\"seq_type\":1,\"seqno\":2,\"tag\":\"signature\"}",
            "kid": "012062715f5bb0d7ee6b94bf6cb2e2ad2638fad149b7da6a79f94cb26bd24e6760ad0a",
            "ctime": 1432100000
        }
    ],
    "keys": [
        "012062715f5bb0d7ee6b94bf6cb2e2ad2638fad149b7da6a79f94cb26bd24e6760ad0a",
        "012003fa76160110693fc613ced54ee67d6a519db6cbae9f4d8664c3609c074f3a720a"
    ],
    "uid": "74c38cf7ceb947f5632045d8ca5d4819",
    "username": "max32"
}
`,
	"mismatched_fingerprint_chain.json": `{
    "chain": [
        {
            "seqno": 1,
            "prev": null,
            "sig": "-----BEGIN PGP MESSAGE-----\nVersion: Keybase OpenPGP v2.0.10\nComment: https://keybase.io/crypto\n\nyMDvAnicbZG7axVBFMY3RoVcEQIRCSIkTL3I7Dzu7lzbFEljJ4KELPM4cx1ys7vZ\n3bskXPInWJgqpBC1SGFlbyPGB6KIiL2CGFDsRAULdXaDnVPMnO+c33wzh/Pk7HTQ\nm7u/vHDm/blP+1MvHz0dB1fTmzMTpHKzjQYTtA7dYV02hLIoXVajAbJYqEhThmWM\nrd+UtIrQPnBBCafCJoJwIjVBIbqRV+0Nb6NkBZdc7nNepM747H/49a6Ao4gaE2uq\niSSaMaIg0sJwAoqbGCjTShCB+8YyoRRlJAL/DR9LIfrWYgwmxtLbjTu7mGmaaBtr\nUILFlvcpwYybREtuWBKJFqygzOQGeHpDblGCdkJUbxethpEB30SIGigrl2doEPmi\nrl1LR4yStomEhQi2CldC6loCH68QFSU0aJCNR6MQVbCZHptGncjyLqrl0D9TuWEm\n63EJaOfx2slgrhecPnWiHUbQm5n9N6Lf94I/zR3Y273VvLq9+uvD5Ojaef79+dc3\n+dri4rfNg8vvmo9TweefR4cLP97OPpwvXj9gq8/SlbsrF64sHV4vvlyky/DiL4nZ\npE0=\n=S2XN\n-----END PGP MESSAGE-----\n",
            "payload_hash": "2fe005da2e66271bc5904ce234a0960c018e6a408fd4ee12aa1c03b31dc8cd81",
            "sig_id": "a486871a4705c9d1690125cd6ba2f6bd4ec74b129a466ad0dc0b78b0dd0bf3630f",
            "payload_json": "{\"body\":{\"key\":{\"fingerprint\":\"f09b1c340a70f0a7bafb236e5932539f89252ac2\",\"host\":\"keybase.io\",\"key_id\":\"5932539f89252ac2\",\"kid\":\"0113dd7c3c2a2c442be1c9d52eb5d7e34cb92906df49bb3421e0a749ba996ff00ed70a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"type\":\"eldest\",\"version\":1},\"ctime\":1432325384,\"expire_in\":10000000,\"prev\":null,\"seq_type\":1,\"seqno\":1,\"tag\":\"signature\"}",
            "kid": "0113dd7c3c2a2c442be1c9d52eb5d7e34cb92906df49bb3421e0a749ba996ff00ed70a",
            "ctime": 1432325384
        },
        {
            "seqno": 2,
            "prev": "2fe005da2e66271bc5904ce234a0960c018e6a408fd4ee12aa1c03b31dc8cd81",
            "sig": "-----BEGIN PGP MESSAGE-----\nVersion: Keybase OpenPGP v2.0.10\nComment: https://keybase.io/crypto\n\nyMQeAnicrZU7jNxEGMcvvEQORUIJRSo43O7ldjx+n3TFetdr7+7t6/Zpa8XJ9oz3\nZh9jr1+73iiK0lGQghZBD2noKRGpUlGlTMGjQUJQIKVB4N0EpUFUTOOZ+b75+6fP\n/n/z3Y3XDw5vPTI+eOfZez99du3JYy45GFw+PLzLOD7KmNO7zBzvH3iBcBRfzgli\nThnAshxCksu50IYuz0MHs66CBIgdAUmY411HgQoQkccrjsPxkMXAlvK5rSii5wGA\nkQRs5pjxCJ3iMAgJjXNZhG3kYJzH/3vkB6/8aHcih3PsCJ8QP9/LF5d7PEHhoMAp\nnqxAAdou3MX+X+5kLyfxLie7nuRiR+ElTxA5CHgBya4tIF5mlV1ihENqL3GevbQ3\nHGTuHTMRcV4W9RWVk2NxtsxhmUcOBwRWcnjPyWE9qHDAFt2cESIJCjwQPcS6HOR5\nWeJYx8OIhza7pwpxisMIX0Zkmsve2Q1V02uto47eOWpqvV5J1/a7EzrME4lPT48a\nLyp41A4w3aWl8AScsGBCy/5yiWl8enQVx0F0Wiy+qnXRDbMg9id0QrNmdVyixA2t\nRtHeXKjVpsvVUwN0a2TeI16/bDS0ZR9vk21rbAzNUaOXEVPXRwXPWEbJRTKhkpIp\n5fXCiEqrWqGUbKJhKYrPS8ONJtZ66kqqqiYczoOu5skV2gcztLDazWJnmXqe4lEi\nT2gsoqVfcMjaLAw3CepYBamTtlLHAHgcxWmr3NcqMeBsTWqoYB1LlVV8EQeaBIg9\naGLBnVA7nFcrQ7PWbJLInelTMoSNcZXMui1Dw7be5Ls1cxBAPWbNlWwJPVSmw0Ab\nymPa2xh1S5jQZneVVcLBlpSNorRp2uwaDbkLYzprCOv2RUFLN+p8IG1M1ySCvKir\nlVZQpl5S5NrtYFz2owldKDat+qC3MuV+ZdHX+gsftTPdWk2j8LyUog1bQSwbuZkm\nng/P54KeBF7jqr1O6qabS+gTCgSyuKDTvlkBfp8OSSlr81W9rmqFgin0ylq/Mt8a\nAShtbWtQHwlkNKqOhCHXO49sZTwrORPqdJVQzjqd6moeyvOqCsl6LvVICqzZVCub\ny0CfXnBquKyPLGE81oB0ZVpqMpiWrvhgZoGcoa/Wa0tH3/hRjbgNPeityKKvN4hq\njXF50NVHs9Z5go3maLREVlMvrJep5WRdWDF91R60pxPKtkBgnXtdn9PrVttrwUUG\nYC1rbzeFMN5WV3Sx5SOxSWm00YxMKYwsEo0NThQTRyW8l/+T6CrtKIIrYVGW0Jrv\nCMWkE86K5mws5pM0c7zWgi/2LWXrjQzOM4VuYatIYlJwlpLA4tmEBt3Uc/KPprqp\n312fnU3omR5H8YTu3aO1Kv/iqJ2z4yzYOf2lwY+Z9IXLmFM2D7ox2fUBludg3p54\nmT9m8CYgIb4ku4yXve2YCXIb5yLQwwAIyIZYFKHEOq6gAN7FkONtoIjABayMRZsH\nsod4jFlo26wLOIdjkSu7SGbz10d4dfkCid0vqM+cwpzSnu4hp9SOkxAz97796I2D\nW4cHb7352q71Hxxef/efC+H79w/+un/2FH86+OLt7f0r/5OvnuBHH5YePBVu/Hnz\n9h3m698W22sHzx88rv0SPr+5aT7Mvrx+6/ZJ138G//jm589Lv/7w9PcfP/4bjE0M\nZg==\n=4X5j\n-----END PGP MESSAGE-----\n",
            "payload_hash": "dffbf787f044211a12ac4f601ff477c351675634e8a6d6fc8aba468b7ced26f0",
            "sig_id": "deed996f790f6b4441a4dcd06f610e281a7c782111368c5da66806ade383347d0f",
            "payload_json": "{\"body\":{\"key\":{\"eldest_kid\":\"0113dd7c3c2a2c442be1c9d52eb5d7e34cb92906df49bb3421e0a749ba996ff00ed70a\",\"fingerprint\":\"deadbeef00000000000000000000000000000000\",\"host\":\"keybase.io\",\"key_id\":\"5932539f89252ac2\",\"kid\":\"0113dd7c3c2a2c442be1c9d52eb5d7e34cb92906df49bb3421e0a749ba996ff00ed70a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"sibkey\":{\"kid\":\"0113bdd73a83e84db30517b4fb3ddf2930a6c52e2d725406fd1c32448731bfed42a10a\",\"reverse_sig\":\"-----BEGIN PGP MESSAGE-----\\nVersion: Keybase OpenPGP v2.0.10\\nComment: https://keybase.io/crypto\\n\\nyMFXAnicrZK/axRBFMc3JvH0QIikSifTCHKEmTezuzNXHVYWKSyiYGGW+fHmsuRu\\n79y9CwlHsAqI+AuxsVAstLAVxE6ISBq7FBY2VkpQEf8DnT0jdlZOM/Pmvff9fni8\\nt6dmo+biwY+VxudPZ+7PvNvbH0eXstvNCTEDt03aE7KB0wt7DqtRtpE70iaUMe5c\\narkFDVYIMMiscjGgiV2KXFijQNHEeaGM4QIYUp2Gt1Yq8Z5SdCnVpEV8XnSxHJZ5\\nMQqyDrUziCH/7xMa1wdV3RHgjK5wOR+EvxBkU7xYcYi58lJBDNpCnfu/3OOpXCos\\nl9anFo0SqY8TDlTETlodOyGZqgsrLAvdx1Dd11scyE6LVLk5GupfKhOwuJYcpXCG\\n05ilRngTYD0oTnViAyO4FGJBE++Y5SCETDkzHp0AzaZUJW5iWWFW5V3SLsa9XjAb\\nbQ9r8yPPFqkr8kFB2iwk7Siv0ZjgECYmpGgR3BrmJWZ5XXE07hYZBuUgAh4pjZ0G\\nTBJImbGxosIicKGpSqilTGKiBZXeCUQGWjNLueHMWWmdZMG+wmvZbyQ2DYoBaUOg\\n1N0pZLfQo3GJZOfN2ly02IyOzx+rtzFqnlz4s6MnnsxEHy9+WZisXH366ubBi4fn\\ndhvP95c7e687dw4P5/uPrj/YjX6uPrvybfNl4/TZ9zfWH3fY5Q+z976u+bm751ej\\npQvfb+EvBcvoQw==\\n=Gtst\\n-----END PGP MESSAGE-----\\n\"},\"type\":\"sibkey\",\"version\":1},\"ctime\":1432325484,\"expire_in\":10000000,\"prev\":\"2fe005da2e66271bc5904ce234a0960c018e6a408fd4ee12aa1c03b31dc8cd81\",\"seq_type\":1,\"seqno\":2,\"tag\":\"signature\"}",
            "kid": "0113dd7c3c2a2c442be1c9d52eb5d7e34cb92906df49bb3421e0a749ba996ff00ed70a",
            "ctime": 1432325484
        }
    ],
    "keys": [
        "-----BEGIN PGP PUBLIC KEY BLOCK-----\nVersion: Keybase OpenPGP v2.0.10\nComment: https://keybase.io/crypto\n\nxm8EVV+NCBMFK4EEACIDAwSxdoYt7Gw2ZhzqSD7gxAANbFbNJFDslJ0hqd6blexK\nu9a2iSGdnWI7zz8o0IJj+mv+Vbf9IYJ4QWnDMIK8JHPP60BfqIkrjcSZn45LY8q9\nkkQlXP88yFjIrAqxca7GWJrNBW1heDMywpIEExMKABoFAlVfjQgCGwEDCwkHAxUK\nCAIeAQIXgAIZAQAKCRBZMlOfiSUqwmTwAYCzAH6p3Y1AhHLupXtmfoLn2LpbIDCf\nbgTRUwKDYTrM9/Si2w2mDwDUde893GQ8TXgBfROubbwMm7mko9kVG08/+Dgsq297\n+681BmBRwCWt7pJLTAfNG0cCrvSs12DePfHXYs5WBFVfjQgSCCqGSM49AwEHAgME\n+3JxQRgAllBCy62bYqxbi3ABKdvFtay8GhNDasLD9qZLdAebWJmHbuM1FebN4V8+\nXuBqNhDePGCE0vT1F8Sr5gMBCgnChwQYEwoADwUCVV+NCAUJDwmcAAIbDAAKCRBZ\nMlOfiSUqwjyeAYD6kut9MSuZozqQOfOS80WCqJe+mrf8ifCp1FvndALhnfo9F9qZ\nVja22hB3wWvxkOoBgLXVrYq6cmbRrmbqRSZ+4XpyoCNFQer6wsa6fvyTdMhcyBUi\nlTYjwzKKOxFopfemv85SBFVfjQgTCCqGSM49AwEHAgMECIKXe20+/fgR1Sv3iCGo\nGGH1tCUkdBsTy9ooXhtw52HA1ua0D94Uij2fOinADT5nRfSA601HUVLtIbl9kie2\nbMLAJwQYEwoADwUCVV+NCAUJDwmcAAIbIgBqCRBZMlOfiSUqwl8gBBkTCgAGBQJV\nX40IAAoJEKpIHwvfFOWaqoUBANYAO2WFA+1ykf52R2zDZWjktVd+bs8VlVVGU6ek\n8EV/AQDmKu+utOxSyXMRIe6sHCnMS7TcKHF1tpCuBv0dahwWOYqcAYD2Q2ntJ/W0\nI+rjw2T8uHbDIpMVERfg48E/YV2c2nJMAl6EPSbCZoCWs/PuXZfOQp0Bf0HqVZ3L\nBm/NhmTS3jvcpiFzJ/oM/RJf0mxVIV3ECTTvCF6nKaXCCTnKe3seNADTVA==\n=gaPN\n-----END PGP PUBLIC KEY BLOCK-----\n",
        "-----BEGIN PGP PUBLIC KEY BLOCK-----\nVersion: Keybase OpenPGP v2.0.10\nComment: https://keybase.io/crypto\n\nxm8EVV+NCRMFK4EEACIDAwRX2ASuVhbQWsemulbD6U1PRTZV/tAUb7P4IeRzklSs\ntWk9coRRpepboFjTWiAZjOz/rpVLd+DgUfEcUt8yH7jb3K8meEbiVGZLgeBL6hPY\nXpm73qI1saVTCA31AXSdJZnNBW1heDMywpIEExMKABoFAlVfjQkCGwEDCwkHAxUK\nCAIeAQIXgAIZAQAKCRCd/xzIsqmWA2Q6AYC5ayGbUfWSOhd9OiS83Dtf+91Ege7z\nECVUym9+RHIUS7mQaOV8lQu0zhugvsqlmTwBf2eQMX+xvMtM5YdDI33lZLqibTJS\nx36yabMk+l1C3jdNQOKYL2iJU3dJMjlaTiWE8M5WBFVfjQkSCCqGSM49AwEHAgME\nyz7k0POd3bNOpuz4Jc2djhmtgIlyCHaeUgfAnrP7QqlC8Q8WzzmBFz+od1EM5tEQ\nRB07+YGPFL1Kbq/DO8OJAgMBCgnChwQYEwoADwUCVV+NCQUJDwmcAAIbDAAKCRCd\n/xzIsqmWAzVQAYC33nryFGGQnMSRQjqm3K9xe/l8/Zz+UjWBBviEh96ccsngJGLw\nKadsWyYuOuos2y4BgO1ChHwRNiSGXrpe1bBD1HJCvPxOhqQvNaqyjK7BYdPW6HhR\nnzGp0HlAnNLb7lj2/c5SBFVfjQkTCCqGSM49AwEHAgMEMVlURn6Qe9pHS2I5pFII\nLru49zxqsoOws8+oz6ojqWxJUZeb1lEHrcb1kvNpTL0gs/TAQMJY+LFwU7E/PeEc\nO8LAJwQYEwoADwUCVV+NCQUJDwmcAAIbIgBqCRCd/xzIsqmWA18gBBkTCgAGBQJV\nX40JAAoJENLwTAfm5SCSIh8BAM7UnJtmN0gMpEfx2Z7BS0CjKm+dGkuMz4mCW9U2\nsrDHAP4ukoSjPtSs//isCS+474UOfRu6MGubTWB4qFTSzOl9Aj8SAX9q07KDj3Vn\neSOQEaqzB8WFDK/Ye2+Vz7WUgzXpGt5ZE+D8buy92f2pN1bgfnL9s2UBf3IrS3pC\nrqgRPXzqJSbJk8Pg3PRtEF19LLII2QMsWaIypTOtGtNyR/mNmQTgqnzikw==\n=euv7\n-----END PGP PUBLIC KEY BLOCK-----\n"
    ],
    "uid": "74c38cf7ceb947f5632045d8ca5d4819",
    "username": "max32"
}
`,
	"mismatched_kid_chain.json": `{
    "chain": [
        {
            "seqno": 1,
            "prev": null,
            "sig": "-----BEGIN PGP MESSAGE-----\nVersion: Keybase OpenPGP v2.0.10\nComment: https://keybase.io/crypto\n\nyMDzAnicbZG/axRBFMdXTYQcCJIUYulAmrDovJ3ZmdkDBcEqlSBWCR7z8xxy2T13\n946LRxrBJhaCKHZ2goKFnWghIRoQhKCdWPoHSBAr8cTZDXZOM/P9vs/7Pob3/tSJ\nqLN0dmdl58PB5d1jH3f3R9F1M340RaowW6g7RRu2vZzP+7Yclj6vUReRTEgADkwb\nJYAZcFoYh9PEcDBOOCZERiEFFKObRdV0hBglK3veF8ELoudNcP/Db7QFDEAowVkm\nWaozxbFzWhPlpEsxYGlkwhkmMhFGc0gsKCO1YKAoU6CNTZmkDMsQN2rjONVEaMe1\nVRnlLmUkwTQ1QsvUUAFZA1a2zOWmDfSmnJAEbceo3ho22g6MDZ+I0diWlS9y1IVQ\n1LVvaKAkYZQT4DGyk6Evbc83BD46MRqWdoy6+WgwiFFlb/WOQqEVedG+atkPYyrf\nz2U9Ki3a3rsxFy11opPzx5tlRJ2F0/9WdFdHs8X+2qs3c7+ffn229unOjxfn5s88\nXP7y9sq9Wf5y9fG7B4fRnwuX9hbu82+/ru5f/Hl7Oll/ffj9M3++fO3J4mx15YCs\n/wUBuaVy\n=jTqg\n-----END PGP MESSAGE-----\n",
            "payload_hash": "fb7a58e58f6acde118349495d44438bbd8577e88795ed81c9149fda3b4628d76",
            "sig_id": "d92e49fb62ac5009f53989807cc82113fc5a806a8b717097f0ca8be8822553840f",
            "payload_json": "{\"body\":{\"key\":{\"fingerprint\":\"398a11716cdb816d1fc8df052d71df8f68894151\",\"host\":\"keybase.io\",\"key_id\":\"2d71df8f68894151\",\"kid\":\"011343099a65c9b70ffcc3bfaf5010ada27603a28dc712e1bdac861b46b1cde56a460a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"type\":\"eldest\",\"version\":1},\"ctime\":1432647317,\"expire_in\":10000000,\"prev\":null,\"seq_type\":1,\"seqno\":1,\"tag\":\"signature\"}",
            "kid": "011343099a65c9b70ffcc3bfaf5010ada27603a28dc712e1bdac861b46b1cde56a460a",
            "ctime": 1432647317
        },
        {
            "seqno": 2,
            "prev": "fb7a58e58f6acde118349495d44438bbd8577e88795ed81c9149fda3b4628d76",
            "sig": "-----BEGIN PGP MESSAGE-----\nVersion: Keybase OpenPGP v2.0.10\nComment: https://keybase.io/crypto\n\nyMQ2AnicpVVLiOREAJ31iyuCuHpYT2Mf3XGTSiqpZGTB/mZ60p/pTv9tHJJUJZ3u\ndOXb+fSy6E1YFBXBFU+CiwcvK3jZi7CIiih49ORlj4JHr6LpmZW97M26FFX16uXx\nkvfy/XOP7128dPnmqzd/+rV878IvP/CbvSFOPrleMjyclw6vl1bkbCIuJlF8unJw\n6bDEAsBDnpVlXRRM2UCsZZkmb1i6JbCA1bHOIZHldU7CJgIcAQbWTUkEBhQNYGIi\niDoUWb10ULIcapPQDx0aF7S8LOkAICCa2JCAiIFlSthiBQ4jgC3JEiVJhkAAxcWF\nF+1uFOIMPSJXHa/YKxanZ/IegT/XjYmODUIs9n+Ngm5zRoegyUumhUxiyBBZgshz\nLBSwZOoChhKQd8CIhFRfkwK91jOeK904KEWO8cDUh25ignSe42S2MKBwjLU4CbAm\nhJgQicMcx/OmgCAwBGLwQBQ41oDAQsjgrMIlxJ+7GZKEhBE5jRy7oH1tNyp1pdnZ\nP1FO9tt1TSsr9bPdOR0VQMejh/vquYP7XZ/QHSzhrrJXATunVW+9JjQ+3F/EsR8d\nMsxDrxkzzP3Ym9M5zduNaZk6ZjhTkQ76FaU9SuquPGtMB7DB5nJ+tN3SaOg0PBW7\nxGXXBIs5MTZCjPBKGXXmtBDoq2O2rqbrJtXKHtey+71RUx2vbco1nOZivBqaRNhk\nbUPTp/qsOznpGinaipusCtylPqekGjGtyRWIyLQ7CSGZxZaUhWwWLVmmF7DZ2OYH\n6qQvB2UjBjhfNqJm2W0oUdw/MnUtzuZ02Rb6CsuvlFYtW6ycSqAdqd10tbJPJora\nU8bVXtKt5LBdlx372PMFtue6iZS3eobqjhercuGWeSwEbAU6dSQ4NZrVaTYcao6r\nTCrNBsdwQeIrdthdVxdWPon9lsNMhyhjREZKZdrxj+dUCBb50PDtzqzbr7JeCgEU\n0i0YeEq8aJhhNWyOV5Nl13DVanwc+DkBg2CSe5avKnX9OFXnVGuvYN2eTNtkVhu2\ne4NeQJcr78hNR80eVWqtZt33tWEaa4NgqUYdRVWTXpg0u/piUjOWpDGntaEhd5PU\ni6Ab9DEMh+VE8/rJqtP3Us3ut+RQPU7XU6bFtlGYjWDHDzwq5APNGGhmZQTnVI/1\nVdcuD2ez0JcbSyVotcZqLrjDuh8chWnVh0e+WAYdRh+p/nCZjUaLZCtU7I2zEIP6\nUJtTn7jt5mDTUHrblt2arjtQcDXIR9nJkXBliddLc7S1GORtvSsCOlmpXcb1CIMk\nO2eMFqsP5nSRYprEJ6G7RWS8TE2eAcIVk5/I8jpxmKOUE7qN1FJxhjYDeWIhylhD\nJh7LEkGyi7vFV03pjKMuL9SmvIZw89qcXgvrNWVOz9JT79QekahdsuPc3yX9QcAP\nSsl5ykqHoDg0Y2fXAwDynAiLLKODEsl8JySnzg7xoFwOSn4R44LEMpAuSEQoakwv\nOhMAiYcylItmgZCXDANLAkJEkpAsECwBUwZQtrDOFyVb9C4Si8dHJDg9lwTOFtQr\nHXKFSt0+E2kXb2sTktKN7956Yu/Sxb2nnnxsV/17F595/r8fwv1be//El9MPnz3w\nPrj957feK9Jjb37duf2p+nH88vv3P7rzXv70nQt7P/99943fX/greOnLFy8tfvzi\nt9z9fPL6Z2/fu7X/zTt3v3r3j38BlbAMxw==\n=U60b\n-----END PGP MESSAGE-----\n",
            "payload_hash": "b0c1a6a4f8af46483e1c249ef49a727c991d601537cee9e799d414357f1edddf",
            "sig_id": "3edc748f9de94269229d930ca1090c78caf3ade6148bfd1571290eb6c05284860f",
            "payload_json": "{\"body\":{\"key\":{\"eldest_kid\":\"011343099a65c9b70ffcc3bfaf5010ada27603a28dc712e1bdac861b46b1cde56a460a\",\"fingerprint\":\"398a11716cdb816d1fc8df052d71df8f68894151\",\"host\":\"keybase.io\",\"key_id\":\"2d71df8f68894151\",\"kid\":\"deadbeef00000000000000000000000000000000000000000000000000000000000000\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"sibkey\":{\"kid\":\"0113de7a32290a1112e0f2810c44dee82d2233c5741b5eb316520b41f77b2fd1f7360a\",\"reverse_sig\":\"-----BEGIN PGP MESSAGE-----\\nVersion: Keybase OpenPGP v2.0.10\\nComment: https://keybase.io/crypto\\n\\nyMFYAnicrZK7a1RBGMVvEl9ZFYT4F0y9yHzznsUiFoKdlel0med6yebu5t7dkGVN\\nEGIpKW0EKwmInSAo2LgRQVIKWmgn2FiIhWkUce5uxMbSaYaZOXPObw7z6uxC1lja\\neCs/LX+47eYOXr4eZtf8xr0xsj0/Qq0xWg3TKXR9qAbt1dyjFsIAlFGstRHcaStx\\njM5RG03kGLDxhkiBqSHKOwkkgPXGKQGWCQvOBy4ME9igJop50Qllv8yLQbKlWhkA\\nCcJ5q0B4iE75iDnxEnxUUSilGXBIF2/2qvpGgrOmChfyXtpLi/YU7x/6/8w9nNpJ\\n5qhyUbpgNZORC0ow4145wz1ToGthFcrCrIWkXjOblKCtJqpye1TqXyofpKGEaJwK\\nSMk4EgXYMeZDUMQTQqnjkoHlwVIQnGDLIEppSUwtSTqjKsNGKKvQrvIOahXDbjeF\\nDUb9Ovwos4lqRd4rUAvSoRvkNRowSgRL9rKJwmY/L0M7rxV4Npqon5yTSbTScBV4\\natakOgAUZZrp9FjGqLLWKy5lUEpqHrwCp4Hp6A1N/aVKpUjxVVhvz5Bguih6qEUS\\npelMITuFGQzLgLYmN45lS43sxPH5+jdmjcVzf/7ozo+57PkKO/loe/78gy/bL0aT\\nhwdnvtPrlz7eWjwc3/15+c3X99mvi/Hw25OFwfKdx7uT9Xf7n/fU/tW98e79ldOn\\nnnZ2nl35DY3S7dI=\\n=rEDG\\n-----END PGP MESSAGE-----\\n\"},\"type\":\"sibkey\",\"version\":1},\"ctime\":1432647417,\"expire_in\":10000000,\"prev\":\"fb7a58e58f6acde118349495d44438bbd8577e88795ed81c9149fda3b4628d76\",\"seq_type\":1,\"seqno\":2,\"tag\":\"signature\"}",
            "kid": "011343099a65c9b70ffcc3bfaf5010ada27603a28dc712e1bdac861b46b1cde56a460a",
            "ctime": 1432647417
        }
    ],
    "keys": [
        "-----BEGIN PGP PUBLIC KEY BLOCK-----\nVersion: Keybase OpenPGP v2.0.10\nComment: https://keybase.io/crypto\n\nxm8EVWR2lRMFK4EEACIDAwTHLrzQ9a9ejgFgBkMeCo3GZxk+st5RQT7ogWdd/ReZ\nx+a0RSzVQoJ1YWxMOH1gXe2g3Ez89shdUwdRhAIqiHu+3Dpju9gLULaggupPHTcH\nDubfmoXLs9XZlWdfwnebobDNBW1heDMywpIEExMKABoFAlVkdpUCGwEDCwkHAxUK\nCAIeAQIXgAIZAQAKCRAtcd+PaIlBUYbrAYDbujynJ196wGcCxBatpJkEvWfl33HT\n7h7LflVFH7gRpVBWy2qJxhwpMJXySlhKFfcBf06AZbOg3B5UaNh7e+xp6J/8Ey6j\nN1zYxdnrDFqjIyPQO1nLLkQx9vDVT6iMwuki+s5WBFVkdpUSCCqGSM49AwEHAgME\nuJXDKpNHbnnKFnjHZYBo7/j945n6YC/w0uEmEspqY1pAGhDBr83onBBiTJrDMM88\nrzA2voGwkdsFTlFDtps2BQMBCgnChwQYEwoADwUCVWR2lQUJDwmcAAIbDAAKCRAt\ncd+PaIlBUYC5AXsFr9Q8NeRhUqhu5QQ1q2ZRHc6noVTCYx5rljmHACiqFfM6MjFy\nGjfXfnm1HSe3cfUBgOSj+1sDKahuZg/F0NC4fPSNFdlDKbP2awmCX8fxBOePN+zm\nZgfA+9Fn3PINLiUkqM5SBFVkdpUTCCqGSM49AwEHAgMEkVmgczonctdljbaeqhVK\n24Y0PWoqn5inr6/kIG2jez1UW1Io8F4Nyc7NZl+UPHE1+hKht2Y0438Jg89cE9aY\noMLAJwQYEwoADwUCVWR2lQUJDwmcAAIbIgBqCRAtcd+PaIlBUV8gBBkTCgAGBQJV\nZHaVAAoJEBmKKorIzkHAcm8A/3rb4Kbt4BHk93U4AaNukEZWS+kPqKGJGEJBHsas\ni1IeAP4qFQIvPcUFgCBwM0aOvQoIKY3u0XqlfPAmY6YJo/mf/w9BAYD0afRCV6U3\nCPUZshdS5jO5/3Y23QoSfI2ibq6UWgAmKXo7rgLwsSJpFIriCIN9vrwBfjneWxQ9\nOb/HHvXW0jBP3qMZhfLQdQPM+FYBYIpoES8UNblYyNkPXVXyojuNn0pKHg==\n=4ZY+\n-----END PGP PUBLIC KEY BLOCK-----\n",
        "-----BEGIN PGP PUBLIC KEY BLOCK-----\nVersion: Keybase OpenPGP v2.0.10\nComment: https://keybase.io/crypto\n\nxm8EVWR2lRMFK4EEACIDAwQNQ19PK1/SLtX48z0hz3dPFw7m8wgGQHKZtr9+cSSh\n93/8rC/WzmAwwA42iswhS/n0QCIe00JsKvku5MoU1wtJa0klQfKOpf2JMgx56Ei8\nWoEfWkVQBYjbZo2TwHzwIUjNBW1heDMywpIEExMKABoFAlVkdpUCGwEDCwkHAxUK\nCAIeAQIXgAIZAQAKCRD6HZfQ/y2R/RY+AX0VG+yecb/FPERfFpID8PgDwisbDYYV\nu9n7rxWXweO4r/XLo8KQR9FmTwlNgua5oX4Bf3do5Lro6gTWsqDvxcYo/6BDCC6s\nU0GlFSYX9lo+mJNY5Fhj5BM/YEFznEMVrmEFAc5WBFVkdpUSCCqGSM49AwEHAgME\nICRjm0O9+5c5ylqcWEhpK42kJOvN2WzMsuAXPve8YXwXZWoI4mBKWlhpnxMD2i2U\nNDj9FgbsOWxAYTwBNl8/wQMBCgnChwQYEwoADwUCVWR2lQUJDwmcAAIbDAAKCRD6\nHZfQ/y2R/QaIAYCEzFPDc3EWSjl3cudokf4K4/dd9nODpNqB/PaGex2tQv4jaetr\nRf/HcTCqICyuBUQBfAvDoRiShMCUqIg+/Uj279VbdDw5ncGpMF8qeM8CxyTartVL\nwTPEsN3DC9puTb59/s5SBFVkdpUTCCqGSM49AwEHAgMEzpgc1Zzf4tsT3K8Fp6ld\n6drFm6y/vnUDbRW1OiHgn3xHPYjOHtb5SKPTg7Hcaa9JinnJcGVwPpkRYm7ixRVE\nFMLAJwQYEwoADwUCVWR2lQUJDwmcAAIbIgBqCRD6HZfQ/y2R/V8gBBkTCgAGBQJV\nZHaVAAoJEHbTN+NA24BjqYkA+gNxwnl8x2yCNjHtKT7vhLBZYvL+qXyUYpOVWjaH\nbe6yAP0QZNQTWZxt5ArLWBpjE7PrwjG7Nu2ljvGA7Z0K7CBpj1w3AYDP9Cuq7xnQ\nG4PL8NfiRvgyvaHiUFYuRQA2z97u3zdBsZzN53dj85Q8s7i5niUfSdMBgL5Fu0H8\nJGBMySeMF03CLsEOS9saWXVqTqY/kcc09A91bhQ0shxPBbpXZgHisQEIZA==\n=h1vu\n-----END PGP PUBLIC KEY BLOCK-----\n"
    ],
    "uid": "74c38cf7ceb947f5632045d8ca5d4819",
    "username": "max32"
}
`,
	"missing_kid_chain.json": `{
    "chain": [
        {
            "seqno": 1,
            "prev": null,
            "sig": "-----BEGIN PGP MESSAGE-----\nVersion: Keybase OpenPGP v2.0.10\nComment: https://keybase.io/crypto\n\nyMD3AnicbVG/SxxBGF01CjmIBI4UtgMhCJsws7O7M7uNiIUQCSGQNCF4zM58ew6e\nu+fu3qmIlWCVJuA/IEKaQIrYpMkPJAkRhSvENNpYaGGjpSBiZlfS5WPgm/e9N+8x\nfD/uDVi1er33aPfB2f7vvt3vPzvWK7U2sYyiVC2hcBnNQtVinTQha2c6KVCIVOzG\nipHYA0yxkswNvJgTRxDsE8kw9SXnhsXIRjNpXr4wNpHI4YlOzcyAhlZm+h/9bEVg\nQmiEAYQUWAQKRIyF9CNg4GHGmOIBMB74Jp3ziDhY+R6JCI+oIlx5rlTmYGHsOpUd\ncyXlMmYSosBlsedTB7ue4lJ4yuUkKIU5ZImYA6OeE4vUQSs2KpbaJYaWAvMJG3Uh\ny3WaoJAYUha6VBOXOiacUmwjWGzrDBq6VODbslE7gy4Kk06rZaMc5hu3pqQCSVrd\nCtE0MbluJqLoZIBWtqfvWPWaNTTYXy7Dqt29/29FWwt9Vu/l2/drf76NfhnfOnm9\n0Qpf1I9XNz+PPR7onY5fjUx+sm4mrw/ere5tXv768PThxZF4NrEz1X84vHO+3fv6\ncf3g+Zu/KrimzA==\n=9oxM\n-----END PGP MESSAGE-----\n",
            "payload_hash": "a9e948e5353741bd48065c251c7b8224ff52d8de16f4f824b7548e074dd99d67",
            "sig_id": "3d80007ad5e64c9367fe566b03c1918bff04d522537787d7385fdc4bed274a740f",
            "payload_json": "{\"body\":{\"key\":{\"fingerprint\":\"df4fd71f5e030dc7495f812a1061c7036c881f50\",\"host\":\"keybase.io\",\"key_id\":\"1061c7036c881f50\",\"kid\":\"0113b0eeaca0a9deaf0ac6be7e50777d89e789603088b120d651b18b3d18d54cd4cd0a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"type\":\"eldest\",\"version\":1},\"ctime\":1432651330,\"expire_in\":10000000,\"prev\":null,\"seq_type\":1,\"seqno\":1,\"tag\":\"signature\"}",
            "kid": "0113b0eeaca0a9deaf0ac6be7e50777d89e789603088b120d651b18b3d18d54cd4cd0a",
            "ctime": 1432651330
        },
        {
            "seqno": 2,
            "prev": "a9e948e5353741bd48065c251c7b8224ff52d8de16f4f824b7548e074dd99d67",
            "sig": "-----BEGIN PGP MESSAGE-----\nVersion: Keybase OpenPGP v2.0.10\nComment: https://keybase.io/crypto\n\nyMQmAnicrVXPq+NEHH+76oIrorAsiAd5xIOHPreZ/JxURZo27Ut/JOmP9LUl7CPJ\nzLRpm0mapH3teyze9iYK6lFBlD0J4tWjrigsugcP3vwDPHjXBTHtW9nLHh1Cwsx8\nvp/55Aufz9x/8ZmD6zduPHzjwc0/fv3pyoMfiqsDG92tXDBehLZM6YKZ4/0HLxBO\ns9N5gJgSwwLAeyzGru+yroKwS1jXlzwsY5GVZRlBBctQkViehdADHIskEXgAejwC\nEImCj/KHdZkjhgR0gpM4CWiW0yIiECQDIuK8EvmyoIgEAs4FrAR8meUlH8J8l80L\np1G6q8jFeW6KbwVRvpZPTvfynoL/n3Wv9nSy4PPQJ7KPPUWQiSjxHCuICPquiAQI\nlB0wxQl1Q5yjQ3fDc8ydIyYNvMdNfaKKk0SMZdHDPIKQ5xUZChwBvIgwkmRPFjgB\nQUJ4EctE4DzPZbEn5ScCUYKeh1x3ryrBa5yk+DQNJjntm7uhanXdOLTq1mFb6/XK\ndW2/6tBBDgwiWjpsXnbw0Iwx3cHW3C32FmAdWonCENOsdDjNsjgtFYtPel30k22c\nRQ516LZdOynTwE/GDcPddDb1NqJgMEa23gmkqs2PvJrY58aqRSZqECXDFl2I7YmM\n6ADQ8TxrOVRC9tquHWvetozVUSEYTOsw6BFVXwWxvpiBQQa3rBANZ4LbmZv9omUp\nZB6Sc65dn5s8mTv0DLPxat3nfJMshALpi81l0R9qBufVSJewICv3M1VrW4M+GG1n\n/FJnz/R0MyubumFIqaY51Gts2UhTJ/G4P2z2AtRLt2cgHnfkpj1ZhLOWZdSwb6LZ\n5swDdjtoVdtxI+wi0DfTahVIXYfqFbWOl0hoRnoI1HIaDMxRz0XAatsbaCwFN9AD\nWrUHtM0f+z24UgMeW8V+MRyu3Xhh+Q5tY67ih2o/iZqdWV1frY/HddwUG9xkkjTL\n6flkOcFhd1Bnh5E18CgiSq84qHWa5YqdtpsQOHQk0R4te5Vp2V8vBie+v1Sqo3LD\nFoyuO4HHTUkLNGCMpRFYLrXYnSZ1tSN26bKmWNN2Ixk4dCladWIUzqebuGYFYofC\nbnVQ8I36qm6QzWZTzVkaFWU6r/rNtYEkV9XPN9tOOB13E007O3PoZhQaCHZ7ZiUJ\nRiAzmh1FtZtQ0dR50V4Yq8Vouk1aM97tJmsNs8mqKo8621YWTYOZFXG2QxthOk1h\ny2Y3rU4htmoduPANMUQmUu2Roit41bBmvtiex1DY9mRiFq1I0SSl0IgLRUue5hrG\nWOZdrm8WfHlYgGuvJom+tOyThqz4J+uthIayIaf8mmOVoneiwH5ZNs6VIrG6qxEh\nFYcqq+35DBVWyGSrDn3HjwLRoXvnaEb1KW7auTrbxjuXPzb3EbO+dBhTAvmmnwW7\nDADCzuj5mz1i8CYOEnwa7BDs5Thi4tzCOYmrYEWAWORFXhaAl4cJm/8CJ+ap5kGO\nEwgROQQRBlKemJATPFnM4awsIKQoeWDkx6d4eXopCewnNGJKXK7SnexFTqibrRLM\n3Pnu9rMHN64fXHvu6i72D64///J/l8E988rBn5//mIHvP7390kc3H/7y4Tfv/nzR\ne+WD1y/mtX/+voY+1l+7cvD+l5y5uKff/eToEf7ts6+4L+4/0vDbb7337QtXfz/6\n6+tX/wU94wjD\n=mdjZ\n-----END PGP MESSAGE-----\n",
            "payload_hash": "a6eb21956740c5d06a82248710b0243ed7f4a9d83a0ab0ad345971cb0532d3f6",
            "sig_id": "03a631ece380553b7a422c0d53ff9492a0a036594580304fbbea14af928761c10f",
            "payload_json": "{\"body\":{\"key\":{\"eldest_kid\":\"0113b0eeaca0a9deaf0ac6be7e50777d89e789603088b120d651b18b3d18d54cd4cd0a\",\"fingerprint\":\"df4fd71f5e030dc7495f812a1061c7036c881f50\",\"host\":\"keybase.io\",\"key_id\":\"1061c7036c881f50\",\"kid\":\"0113b0eeaca0a9deaf0ac6be7e50777d89e789603088b120d651b18b3d18d54cd4cd0a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"sibkey\":{\"kid\":\"0113265ee75be3d883397842f135ded67b7424d8ff35e7f42bba0eb67f51568bbdaa0a\",\"reverse_sig\":\"-----BEGIN PGP MESSAGE-----\\nVersion: Keybase OpenPGP v2.0.10\\nComment: https://keybase.io/crypto\\n\\nyMFWAnicrZJNaxQxGMdn1VZdUIQi6DU3YbF5T2ZBPfgBiorXLnl5Mg7dnV1nZktL\\n6dUvUFHEbyAeBY+iVhG8iSfBIuipIlj1Vt8y04oXj4aQkOT/PP9fkmfz2MGkO3fk\\nwe0puvT2cOfl4+fT5Kq/cXEN2bFfRf01tATtBEMPVT1Yyj3qI0wIsxjAOINN6sEE\\nbJy0oEBgpZTXKSidSsyw1pZQ7KUglmjLPNFecOdjxwb1UMiLDMpJmRd1TOsDD16R\\nICBGeqd4KoIm1BAsiVOYSad1PMUx8Nq4aiIinDUVnM3HcS8uBi3eP/T/mXvaplPc\\nMe2CcmBTroKQjGIuvHZGeK5J2ggrKAszgqgemRVG0XoPVbndf9S/VFQKACUsMK81\\nY6nSnAbChAcvlVWccq9DYAJU4NRag8HK6EiE1NZ6Y1qqEpahrGBQ5RnqF9PhMJrV\\nq5PGfN+zhxpFPi5Qn8RDV+cNGuGNfxxxD8HKJC9hkDcKvNd6aBIzxyQmhZRrEEww\\nxYmNd8RSOCriY1tNKQ9BUK89EBk/UlNulYhyrLj3aRrvEe0ruD7YQyLtohijPo2U\\nJmshs8LU0xLQ+pPFQ8lcN5mdOdBUY9I9euJPjc5Mkp84yS7fO/Po9E69+Jp+/P7h\\nxZe73a2TO+c7X+8vbF65c6qTfJ79cWvy6dX7N7s3v209/bW98TA7Nz9/fPRuYffC\\n9uyzjd+udO0D\\n=coi5\\n-----END PGP MESSAGE-----\\n\"},\"type\":\"sibkey\",\"version\":1},\"ctime\":1432651430,\"expire_in\":10000000,\"prev\":\"a9e948e5353741bd48065c251c7b8224ff52d8de16f4f824b7548e074dd99d67\",\"seq_type\":1,\"seqno\":2,\"tag\":\"signature\"}",
            "kid": "0113b0eeaca0a9deaf0ac6be7e50777d89e789603088b120d651b18b3d18d54cd4cd0a",
            "ctime": 1432651430
        }
    ],
    "keys": [
        "-----BEGIN PGP PUBLIC KEY BLOCK-----\nVersion: Keybase OpenPGP v2.0.10\nComment: https://keybase.io/crypto\n\nxm8EVWSGQxMFK4EEACIDAwTUgD4K0SSmhOHwgdWN2M8LT0NVQWSwgIYd0jeH1fM5\nJsE/Dsuem9GQwHLGwUILIINMzOPHbPscjOC7Rf3+8uAHiTDj00ejRgDoJpCKIXjH\nEM9MMxofqG0eDvEldTWKZibNBW1heDMywpIEExMKABoFAlVkhkMCGwEDCwkHAxUK\nCAIeAQIXgAIZAQAKCRB0waG+wsm+NVgWAYC5Q86W0VFs6IGqOgD2HIBAhShCo0Bp\nOivlkcK4iO7xFfYfbXusj74byp13ct8Qzk8Bf2VY8vJF6eha8dfcngiCZsEOAGVE\n3abIgE/YOERD/xm1sfzYs50QHNSj+LBU7pBqLs5WBFVkhkMSCCqGSM49AwEHAgME\naomMzHfI91CQZmMq3mYYh1szK9lBxoK8jYMHa0iVxCKd3mcTtvRxMBEak6XSo2U2\ni8i6iv25vyTO1yfJPcmBOwMBCgnChwQYEwoADwUCVWSGQwUJDwmcAAIbDAAKCRB0\nwaG+wsm+NXxRAYDlZo+z7gpcXFyILtxBbhBKz0qqueRVBfkCP8y4ANJJSjH6MnC2\nEcLwj6BNX8yq3TEBewf4GinIZHMXWj8hOH8KKtyidPLhR5WPCgPe3SYP0lZzFYgO\nb1YYrIVicghxSB6UTM5SBFVkhkMTCCqGSM49AwEHAgMEZ5qBf/Nw3XSK0OSz23Kz\nslf6ixA7xs+TNE72MwuzoVDGypTHTF5tz9Oz1/oDH8eCzZHLDHtF3gtgyyTv+864\nvMLAJwQYEwoADwUCVWSGQwUJDwmcAAIbIgBqCRB0waG+wsm+NV8gBBkTCgAGBQJV\nZIZDAAoJEAi0l3UiUd0H8R4BALO/bP6Ir0dBdzHMjsRitZ5cf1gNlTOsALhfTaRm\nCigeAQD2/vmEl6C7TA1hlSuvGgp7Z7UPgUgJizoYJGu2NpAwgqEQAYCtFBHE8FmQ\nKpbovcWv783nqAiwsVRQPgefcE7tZsneUgBa5pT3tFmsPBH0229y8IoBgPcPRToY\nLELCkzhNSfsxY6UI7Qlp57lFIjnxrDIrz5kYrVRmNeLIrbZK+56xv6x3rQ==\n=vv7t\n-----END PGP PUBLIC KEY BLOCK-----\n"
    ],
    "uid": "74c38cf7ceb947f5632045d8ca5d4819",
    "username": "max32",
    "label_kids": {
        "e": "0113b0eeaca0a9deaf0ac6be7e50777d89e789603088b120d651b18b3d18d54cd4cd0a",
        "sib1": "0113265ee75be3d883397842f135ded67b7424d8ff35e7f42bba0eb67f51568bbdaa0a"
    },
    "label_sigs": {
        "e": "3d80007ad5e64c9367fe566b03c1918bff04d522537787d7385fdc4bed274a740f",
        "sib1": "03a631ece380553b7a422c0d53ff9492a0a036594580304fbbea14af928761c10f"
    }
}
`,
	"missing_reverse_kid_chain.json": `{
    "chain": [
        {
            "seqno": 1,
            "prev": null,
            "sig": "-----BEGIN PGP MESSAGE-----\nVersion: Keybase OpenPGP v2.0.10\nComment: https://keybase.io/crypto\n\nyMDuAnicZZE9b9RAEIYdEpA4gUBKl3LLw0T7Ya/XJ9EGpUIKBAoUjvXO2GxyZzu2\nL+RDJ2jCD6BAAgoafgM1IEJzQkhItEgUdHRIIAoU1o6omGbn3XnmnV3Nu/PzXm/x\n0+cLo9/9w/m52eujibcOj1YOSFLAHhkckC3sjtTmGVZlZfOGDEgSc4PSgBJcSaV4\nClzHFGRCNY3RQKQlTQE48cm9om47nE2ia1y2hbtzYmih9fmf3+oKlDHBqJIQIYRR\niCmgjFDEjINOGRcyUCHFgIuAGa4ppG44UJkI7R4VBzyEQFHt7CadXRQYoUwaGUzi\nIEpDKTgNQlBGtyCLW7DGKtdjdPRY7wpOpj5p9spW4wjQfcInO1jVtsjJgLmiaWxL\ns0BwGTIhhE9wt7QVDm1L0JPwSVnhDhnkk9HIJzVuD09MWSfyossanbkxtc1y3Uwq\nJNO3dxa8xZ535vSpdhle7+zFfys63JjzPq78Wt3cKM9dW7t16eXS7M+Hy9P9pQfk\nuP/02c2r+vmad3x7/cqXH1k1vv7maHN2/9X7hezFk/705/ev/Yc3Ht9d/vYXEwyi\nWA==\n=H6Xl\n-----END PGP MESSAGE-----\n",
            "payload_hash": "00169139cc7d3d241bcd0537fa7dcf1f1d6c13f2ef581bf402424fbfd169bed1",
            "sig_id": "ea8a3358e12cdaa444d211fa71ffc4b421040244fd5b72ea2b387d722c11d8170f",
            "payload_json": "{\"body\":{\"key\":{\"fingerprint\":\"b92ce6cd83286882fd2a90d6b0a09ecd7a60fdd2\",\"host\":\"keybase.io\",\"key_id\":\"b0a09ecd7a60fdd2\",\"kid\":\"01131086d7ed575efde67e3912daf12364850e42341c2a0df0d6d06b3acd89425d480a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"type\":\"eldest\",\"version\":1},\"ctime\":1432651333,\"expire_in\":10000000,\"prev\":null,\"seq_type\":1,\"seqno\":1,\"tag\":\"signature\"}",
            "kid": "01131086d7ed575efde67e3912daf12364850e42341c2a0df0d6d06b3acd89425d480a",
            "ctime": 1432651333
        },
        {
            "seqno": 2,
            "prev": "00169139cc7d3d241bcd0537fa7dcf1f1d6c13f2ef581bf402424fbfd169bed1",
            "sig": "-----BEGIN PGP MESSAGE-----\nVersion: Keybase OpenPGP v2.0.10\nComment: https://keybase.io/crypto\n\nyMQyAnicrZU7jONEGMf3OEDiJB7SFVBQrIKugOxtPDN+Bh1S4jhONnFeTjaJL2Jl\nz8NxnDiJ7ThOTidEARIlFRUFFSUdBSUCIUGBhEQHNRJPUVBwDTi7h665kilszfj/\nff7Np/l/88Wz149u3Pzu++fnf7/2zvVr33z5+uZoQN6t3ss5S7LLFe/lfHr5onNC\no/jC90iumOMAQICTRSJRIkgCZYSKEkUKgMRmACKRlwWO8hDxAEObI4wjIuFEB9mY\nyAoPBcLLnJ07yTEvcGm4Cr0gztI6CsRUzCQIyqIsQ0agrWShDmdzCsVEskWOEQKz\nwOkyOkRkcI4d0VNvma1lk4tLvMfo/2fuzWU6icdIxkzC1FF4iQkighwvEBnbByFQ\nDsKIhoG9oJl6YacI5u6f5CLPeVjUR1RUoUBCPHUY4QUAOYlIDCkytHmKAWYCoNQW\nOIHjFYlS3gGiQ3kO2BBKLCuRJFxShTShYUQvIs/N0t4+jLKm11vHHb1zbGimWdK1\ny9VJcJ4JvWVQPG5cVfC4vaLBQZbAU+4UcJNAXS4WNIiLx9M4XkXFQuFRrQs43K3i\n5SSYBDujSkqBh0OrLtmgV64aJGGGtskbqjvWtf1svHaMdl4gaTTtDdRKXXNtVaBW\ne8iGptCQJoEoQcmflvv5tTsoq4grAXfj+jUtneFyh9fc5Xo4H6penY/TVmqMdwM5\n3VqJVMBI2LcbBTQJJH9n+xAbTHKSPpvLfAybsdzya2upE06bIzcdmSNv0GxXhK3m\nlGxStnZ1Hah6r22spktjErizeqMklmCj0tS3XUefWvZq6XmCb6R7c1ftcpoKNMWq\nLPzFpttwBv50fEZJrRrNC4vUVBeTgPQCtG60NLlqBo1uoeT07Hg2X9vLLpm1ulyp\nU14mih3aGtVxOCqJkjyU9Fpfp51Cz9l3sl3gs33bVBWzxqe1Oef352Ubm+5wFUvD\nruUMt9uwsTXT2bmJLAs3m8P6fLSJV6E42y73lmbVM4aS7o9XW9vQqn7DwWaHtcrG\nwAzXjfrO4GemQUF3sBgLrZ0Odw03PR/PR1Fjtx3kzwnnifok8PhNCnV1PkoZW8e+\nF5vNpGr3W/zaYEJlYPv8WVgdsV49nOE9l0RCc6g6a34oYma0zOl8ErQXhicrYqmk\nipbMqedeydWZrpZdZhnlulBe4W3VdALfiLpJz4or42ZVnlpYWnDTqFbuKpOA29g9\nLnBqfgupleV47ofDURXhuS8Df5/nidLvLvbYLid7Q6Ro7FGn3xtWZFnOj0Yz1upM\nggQBmO026ESgVWNKKiciwfsWSgpWrSI02UoU8x0Ew1Z/pMhkCPJWgpypjGddQdnU\nKlF2Js/arJIWEiRMnUKBr81WzbXVvXNnEtw5H/HZwb90kdaqPMZZB4fHu9XB8Q+N\nfpJLrtyWK4LsI469Qz8APIKikD3RSY6mKy+kF95BwV2Nk9wqs/OhO3BAVABSMJYI\nIpAHDiacgDLfSwQzwAARMUAMUibIwGE8B3nIs6yPZGEOJSD7fUTXF1dI4HISLHNF\nmFHa7iWkG9jxJqS5+5+/+eTRzRtHTz/1xOEKOLrxzAv/XQz3xkf/0LsPHrxU/eRH\n9+PbP9984/db5T86qvnq6U+/fPbrtPJX8aNrRx/231t89Vz8weju2688/dYPL37a\nvf0+2nx98tufL4vfPtje+heUNA3B\n=okGo\n-----END PGP MESSAGE-----\n",
            "payload_hash": "525de3a305ed703e4419add54649cedd083464f3b76f296c0bd9c7ca2bff3bb9",
            "sig_id": "eca9bd4e3be4d3be48501d5be8bc01efdbad86b51c9887a11279c44a9eb7de4b0f",
            "payload_json": "{\"body\":{\"key\":{\"eldest_kid\":\"01131086d7ed575efde67e3912daf12364850e42341c2a0df0d6d06b3acd89425d480a\",\"fingerprint\":\"b92ce6cd83286882fd2a90d6b0a09ecd7a60fdd2\",\"host\":\"keybase.io\",\"key_id\":\"b0a09ecd7a60fdd2\",\"kid\":\"01131086d7ed575efde67e3912daf12364850e42341c2a0df0d6d06b3acd89425d480a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"sibkey\":{\"kid\":\"0113e9e1734ebfd451207d7f3982a4ec1cf51eea5050497ee4b16be401a227fa90750a\",\"reverse_sig\":\"-----BEGIN PGP MESSAGE-----\\nVersion: Keybase OpenPGP v2.0.10\\nComment: https://keybase.io/crypto\\n\\nyMFdAnicrZI7a1RBFMdvfMEu+MCgYGEzjYqbMO+5dxshRUCDIEgaC5eZOWfWS5K7\\n6727khBT+qgUBC30A1gugkHExjcBP4EgoqWlWCiI4txNxMYyU8xwZv7/c35zOK/3\\n7kyak2cMf7bvTfl84t2Lt8NkHq7PrhLXgxXSXiULOD5wEbAadBZyIG1CGROMphoM\\ngjIKA6A2KDLGwQbGhZapoii5kMxzSyFQ0EC1E9ZDmkmuQKbUkhYJedHFsl/mxSCm\\ndRn3qKNE8FSnKQ/AbRatjlqaoQdjNQ0APBov9araEeGcrXA678W7GHTGeP/RbzP3\\ncJzOSC9SH4xHl0kTlBacSgWpt7WQZbWwwrKwSxjVS3ZZcLLWIlXutpr6jwozZEZI\\ndAGkYpwaMEFkKbcSPfNBMUSrqKIyM4jSMe1QUmY5NyG2yKgxVYlXsKywU+Vd0i6G\\ni4ux2GClXxffqtkitSLvFaTN4qMf5DUak4JrFXfRIrjcz0vs5LWCbq4W6cfMNShl\\nOmMi896AAC6Z80CViAgGfGCBgfZMBI5BpcwFSbnkMsQvRZtDYLF8hZc7m0hsHBQ9\\n0uaR0nbHkN3CDoYlkrWXF3clk81kz+4d9TQmzcaBvzM6e3YiebTRWD888+XXjfNP\\nv312jSMnPs1NHf9x8v6dczN3v/ZHD5Lfp66+P32rNTX98dW1+Zv3bh8cjQ59uHDs\\n6JOfDx/v35hb//4HjpLqZQ==\\n=VX4o\\n-----END PGP MESSAGE-----\\n\"},\"type\":\"sibkey\",\"version\":1},\"ctime\":1432651433,\"expire_in\":10000000,\"prev\":\"00169139cc7d3d241bcd0537fa7dcf1f1d6c13f2ef581bf402424fbfd169bed1\",\"seq_type\":1,\"seqno\":2,\"tag\":\"signature\"}",
            "kid": "01131086d7ed575efde67e3912daf12364850e42341c2a0df0d6d06b3acd89425d480a",
            "ctime": 1432651433
        }
    ],
    "keys": [
        "-----BEGIN PGP PUBLIC KEY BLOCK-----\nVersion: Keybase OpenPGP v2.0.10\nComment: https://keybase.io/crypto\n\nxm8EVWSGRRMFK4EEACIDAwTGPp5WCEi30jKYsqP+DUV1NWMxi53YT9axzEbrEyU/\n7UdKsWZPZJnOsYpjCr8+VXqIHY56kO1oiR0HqpnoJhsRZ1imNqHq4L4d/QMhyNGs\nICqaFQJsSShlwQwvUffq0urNBW1heDMywpIEExMKABoFAlVkhkUCGwEDCwkHAxUK\nCAIeAQIXgAIZAQAKCRCwoJ7NemD90gI7AX9zOItdXE2tVZuNus0+HvqcVI4cvlXD\nqt/0HTIly/mDP8tgNcPSpgOuzPVCQy+dnYMBfiGFbDQnwge9Cea0tHHIbY8qAOax\niX6yRDuMPaqhbleIKaKU1hJBrtDMLz6pnHLtn85WBFVkhkUSCCqGSM49AwEHAgME\nHZOEO11yQS3sZtB43vEHD9YbSHpvhOtgArJl4/dQB5Jr4DFrmBAA5WPlUbF8EBgc\nJ7oD6qVGETtJDX1R78qo0QMBCgnChwQYEwoADwUCVWSGRQUJDwmcAAIbDAAKCRCw\noJ7NemD90tS0AX9gy15Vv10RalrRSHhjcT/BW6r6MGHy/Z+w33tUAND7EUDBwXvc\n4u4i/+bTOJY6TcYBf2fN+q2Zcajt91IJiqd/4su/SwcpJb9mcHJpc4M1wXen6rFn\nmHVVB47WTJpk+ORFcc5SBFVkhkUTCCqGSM49AwEHAgMEBHRJTnd1BZp6x9KrE1p1\nPJ4McaYM8lu1JysuL+BJdV/yw98XnKBIv8nhnBikJgntcj1jAcWyixuQJvbreZn+\nD8LAJwQYEwoADwUCVWSGRQUJDwmcAAIbIgBqCRCwoJ7NemD90l8gBBkTCgAGBQJV\nZIZFAAoJENLVD2z5KoUDfcsBAKpn+wqezR1uPhU2aG1YRgwTlSr2Qkf/n1Ob8gss\nkCilAP4+zmdLJgv017cEYrMAxozL/zTFhCRns+BBMLnJCaRNDov4AX93/JS4CYca\naw7nB4SNhxHaMgyV6jhsXmO1psuC8uAVqtkQPuo3hf6ttcftfxbnv9gBgNJwoYyk\nkEStMA8gnT4841eHX2JMzuZ9kfqYcKuf+rqkSgX+wyMnuU4rKtKWHc+mZg==\n=pMMp\n-----END PGP PUBLIC KEY BLOCK-----\n"
    ],
    "uid": "74c38cf7ceb947f5632045d8ca5d4819",
    "username": "max32",
    "label_kids": {
        "e": "01131086d7ed575efde67e3912daf12364850e42341c2a0df0d6d06b3acd89425d480a",
        "sib1": "0113e9e1734ebfd451207d7f3982a4ec1cf51eea5050497ee4b16be401a227fa90750a"
    },
    "label_sigs": {
        "e": "ea8a3358e12cdaa444d211fa71ffc4b421040244fd5b72ea2b387d722c11d8170f",
        "sib1": "eca9bd4e3be4d3be48501d5be8bc01efdbad86b51c9887a11279c44a9eb7de4b0f"
    }
}
`,
	"ralph_chain.json": `{
  "username": "ralph",
  "uid": "bf65266d0d8df3ad5d1b367f578e6819",
  "chain": [
    {
      "seqno": 1,
      "payload_hash": "f4f8779457bc3053ee4f04637934ace1705344aeee8a4b38dcffbbf9c9268d4e",
      "sig_id": "55b70fe0e7d5d8605e6607de1daf77c3e862050ad6ae491b5b7829e05130bf580f",
      "sig_id_short": "VbcP4OfV2GBeZgfeHa93w-hiBQrWrkkbW3gp",
      "kid": "0101dbd7f3aaf965df1f60839b970f92b5f368046d693d5ea4e20aa564cab672820c0a",
      "sig": "-----BEGIN PGP MESSAGE-----\nVersion: Keybase OpenPGP v2.0.8\nComment: https://keybase.io/crypto\n\nyMHhAnicbVFbSFRBGF4rIzdbLOuhoqAxuu22ncuey6xmWFFBIVb0YCnLnJ0568H1\n7Hb2oqIrWwkSUkGZ+GAG3ZAeukiCpLkiIYEJ2kXzRSg0SGQDRci0yxzJIGgeZpjv\n/77vn2/+l2uWW6wpaVOt+Rva68dS+nvmI5Yzx97UVwElgCuBuwqUksVD1XQfMYKG\npoeBG6iizLugC0kcIwsQ0p1DhPci6PIiHmJJIRBipBDgACWBkKmgNgoKEacWoBi9\neDRM0f/wSxcLDMuwWMGSyiOkQlHAKquKjMxDBUqMCjlFUHlRZlwiFiGPBYJchGMQ\nEkRqqIgSJ3OMl0HULrJop6iiwIkiZrCMqSUWMKvwoqQKkkxEmYUmMUQMHZURyjaQ\nP1gCYg5AoajmJWb+pRLBWAv/Sw+TUJgYf0XhyqCJlhPF80fvUTQd0/+jsigxQlpA\nB26WMr1hzTRgXTxNJsgC6wCkIqgZxKOZDEGiCelygKBBosCtR/x+803n9QCt0kbI\nR/uENJ+OwhGDgFhvT9EKS4rVsjJ1mTlDizUtY2myte2rLDeSZ5ubh/Zdi7UTpu/j\ns11FTi7JJ9TMo7aiK0O+TeLzvtGqd6fzMivu5naNOpn01WV3Be5O/ER8z60PLV6Y\n+YJEh34Wz/ZeLs0erBl7W9y48wBM+BYOFcRc3U79/ffD+V9q5wau385Jn9/deaHh\n0qyjI9u1/oGRtOcWfDp5b+NU9XjQOpLX5YjO5D+9HxtOTrRlPII2jRnP+Tzds3m7\nezCrcWDHj7zH9tYyUBiPObNuavpkUwTMCfHpjrriznOJq/vnXk96LiaD/pkt5b+E\n/t7uU3sbvz15tXX0SLa9sMmaVtNQXWdLtKxdaMtY17Hta+7BhyMTzQt2/bhtOPU3\nnTkY3w==\n=DD0n\n-----END PGP MESSAGE-----",
      "payload_json": "{\"body\":{\"key\":{\"fingerprint\":\"f683494a72085992082ae3ca94ca39d7be99dabe\",\"host\":\"keybase.io\",\"key_id\":\"94ca39d7be99dabe\",\"kid\":\"0101dbd7f3aaf965df1f60839b970f92b5f368046d693d5ea4e20aa564cab672820c0a\",\"uid\":\"bf65266d0d8df3ad5d1b367f578e6819\",\"username\":\"ralph\"},\"service\":{\"name\":\"reddit\",\"username\":\"testerralph\"},\"type\":\"web_service_binding\",\"version\":1},\"ctime\":1430835851,\"expire_in\":157680000,\"prev\":null,\"seqno\":1,\"tag\":\"signature\"}",
      "sig_type": 2,
      "ctime": 1430835851,
      "etime": 1588515851,
      "rtime": null,
      "sig_status": 1,
      "prev": null,
      "proof_id": "af85baace455686a5e809a10",
      "proof_type": 4,
      "proof_text_check": "\n\nyMHhAnicbVFbSFRBGF4rIzdbLOuhoqAxuu22ncuey6xmWFFBIVb0YCnLnJ0568H1\n7Hb2oqIrWwkSUkGZ+GAG3ZAeukiCpLkiIYEJ2kXzRSg0SGQDRci0yxzJIGgeZpjv\n/77vn2/+l2uWW6wpaVOt+Rva68dS+nvmI5Yzx97UVwElgCuBuwqUksVD1XQfMYKG\npoeBG6iizLugC0kcIwsQ0p1DhPci6PIiHmJJIRBipBDgACWBkKmgNgoKEacWoBi9\neDRM0f/wSxcLDMuwWMGSyiOkQlHAKquKjMxDBUqMCjlFUHlRZlwiFiGPBYJchGMQ\nEkRqqIgSJ3OMl0HULrJop6iiwIkiZrCMqSUWMKvwoqQKkkxEmYUmMUQMHZURyjaQ\nP1gCYg5AoajmJWb+pRLBWAv/Sw+TUJgYf0XhyqCJlhPF80fvUTQd0/+jsigxQlpA\nB26WMr1hzTRgXTxNJsgC6wCkIqgZxKOZDEGiCelygKBBosCtR/x+803n9QCt0kbI\nR/uENJ+OwhGDgFhvT9EKS4rVsjJ1mTlDizUtY2myte2rLDeSZ5ubh/Zdi7UTpu/j\ns11FTi7JJ9TMo7aiK0O+TeLzvtGqd6fzMivu5naNOpn01WV3Be5O/ER8z60PLV6Y\n+YJEh34Wz/ZeLs0erBl7W9y48wBM+BYOFcRc3U79/ffD+V9q5wau385Jn9/deaHh\n0qyjI9u1/oGRtOcWfDp5b+NU9XjQOpLX5YjO5D+9HxtOTrRlPII2jRnP+Tzds3m7\nezCrcWDHj7zH9tYyUBiPObNuavpkUwTMCfHpjrriznOJq/vnXk96LiaD/pkt5b+E\n/t7uU3sbvz15tXX0SLa9sMmaVtNQXWdLtKxdaMtY17Hta+7BhyMTzQt2/bhtOPU3\nnTkY3w==\n",
      "check_data_json": "{\"name\":\"reddit\",\"username\":\"testerralph\"}",
      "remote_id": "",
      "api_url": null,
      "human_url": null,
      "proof_state": 7,
      "proof_status": 0,
      "retry_count": 0,
      "hard_fail_count": 0,
      "last_check": 1430835852,
      "last_success": null,
      "version": 0,
      "fingerprint": "f683494a72085992082ae3ca94ca39d7be99dabe"
    },
    {
      "seqno": 2,
      "payload_hash": "bdc9430e56f7d59eb9e6709cb719b4a8479ad3272c046e3c915a113e0892f8a7",
      "sig_id": "32c195e0fca5c310237f6180b29315afd2cfc4ddce9cb307b50cc4a9172bdbd30f",
      "sig_id_short": "MsGV4PylwxAjf2GAspMVr9LPxN3OnLMHtQzE",
      "kid": "0101dbd7f3aaf965df1f60839b970f92b5f368046d693d5ea4e20aa564cab672820c0a",
      "sig": "-----BEGIN PGP MESSAGE-----\nVersion: Keybase OpenPGP v2.0.8\nComment: https://keybase.io/crypto\n\nyMIQAnicbZJbSFRBGMd33a21JbsQFV0wPBVFSJ37nNnSpKDbQ0UlQhe2mTMzuwfb\nS2eP642NSIrKh3wQMgojMs0IMcyH7KZZLy09dH3oYg+RdqMLEZFGNEcsCJqHGeb/\n/b8f8/2Z23k+T9A74cP5jdO76we82b5Qhad03YOWWgEnSLUQqhXK6ejBrHiE2knb\nijtCSGC6oahQRUAWDQ1CvsuIKiaCqokUSACmEBKEqVAoRBMpt4NjMErRpVaCa/wS\ntghX/+MvHy2IkigRTABTEGJQ1wiTmC4aCsQQiAzKWGOKboiqTnSoEI0ilcoiQprO\ngVgHsiGLpog4rmIUh5muybpORGIQjiQakbCiA6YBg+qGBF1jitpxFKPcbaM9yaiQ\nKRS4lLZM6s4/VnIqLceh9r9+h6a49rfLqU66aiXF4TFAGFtxwgPkbWlqp6xEXAhJ\n3Gk6lguQVIWPphlALxRoVdKyadhyHRrgI/JVKCRtmnZjV5kBAFQ1gE1F1BRKVcYz\nUABUVGRSCXBNVRGl1EAqVgxiMoYxgyaUdYOobrwpujeeEEIyfyaKcGTKisSRU2FT\nIXOrd6ff4w16xo/Lcb+AJzhhyp+PMSTmejrfFQy1X27+OLO/Z/LZWHBFpvMKe2K/\nbiw7PjUdrjy6c97Kc9saJW/s1Podfq1uX+3TlsH6o5ULh0/kPPu1o6Z4c6ZkV11u\nnOYfPpM9MuIPRtnzNqXg0TIjFvha9rmv5cdsWvzK77m4YNqnu48/VfW31g7EFiXq\nNkxquD/Um9d+GkVnPPB9+/kQvOwpuT67FXdFam4u7TqQCfjWZPOzDYOXqjsXG59b\nr395I83qPpg350ejtGTgwsSO0q2HcvZfqw88n78q2QS2tNH84ZIb7xt9y1d9752S\n2zH57srm7vKiL0xc33Rv99zlq9du6D148e1gUeDF9o6T8Nkd/9Vj6ZFIzaaPvwGx\nDCYq\n=4hWA\n-----END PGP MESSAGE-----",
      "payload_json": "{\"body\":{\"key\":{\"fingerprint\":\"f683494a72085992082ae3ca94ca39d7be99dabe\",\"host\":\"keybase.io\",\"key_id\":\"94ca39d7be99dabe\",\"kid\":\"0101dbd7f3aaf965df1f60839b970f92b5f368046d693d5ea4e20aa564cab672820c0a\",\"uid\":\"bf65266d0d8df3ad5d1b367f578e6819\",\"username\":\"ralph\"},\"service\":{\"name\":\"twitter\",\"username\":\"testerralph\"},\"type\":\"web_service_binding\",\"version\":1},\"ctime\":1430835876,\"expire_in\":157680000,\"prev\":\"f4f8779457bc3053ee4f04637934ace1705344aeee8a4b38dcffbbf9c9268d4e\",\"seqno\":2,\"tag\":\"signature\"}",
      "sig_type": 2,
      "ctime": 1430835876,
      "etime": 1588515876,
      "rtime": null,
      "sig_status": 1,
      "prev": "f4f8779457bc3053ee4f04637934ace1705344aeee8a4b38dcffbbf9c9268d4e",
      "proof_id": "9d8b0ec56bbda063e5241b10",
      "proof_type": 2,
      "proof_text_check": "Verifying myself: I am ralph on Keybase.io. MsGV4PylwxAjf2GAspMVr9LPxN3OnLMHtQzE /",
      "check_data_json": "{\"name\":\"twitter\",\"username\":\"testerralph\"}",
      "remote_id": "595594965541064704",
      "api_url": "https://twitter.com/testerralph/status/595594965541064704",
      "human_url": "https://twitter.com/testerralph/status/595594965541064704",
      "proof_state": 5,
      "proof_status": 1,
      "retry_count": 1,
      "hard_fail_count": 0,
      "last_check": 1430835881,
      "last_success": 1430835881,
      "version": 1,
      "fingerprint": "f683494a72085992082ae3ca94ca39d7be99dabe"
    },
    {
      "seqno": 3,
      "payload_hash": "4cc9ac722ffcd53b5de91a37b3a7ede9e80cc3893eb64bdc5988b18eb26a97d6",
      "sig_id": "4200f01dc416139995cd534d577ef5c14b39fde69a6e57ccd3f1b6d9212b53fc0f",
      "sig_id_short": "QgDwHcQWE5mVzVNNV371wUs5_eaablfM0_G2",
      "kid": "0101c304e8c86c8f4b6773478eed4d05e9ffdddc81c7068c50db1b5bad9a904f5f890a",
      "sig": "-----BEGIN PGP MESSAGE-----\nVersion: Keybase OpenPGP v2.0.8\nComment: https://keybase.io/crypto\n\nyMIRAnicbZJ/aI1RGMfvHZdcxpLtD7RysoQ73ve+73ve817dpikzkpKVP65u57zn\nnO1t3Hu973vvzELajJjLsCxDITVKSZuhxQwh/7BsiSn/XG1+FVqK/DjvQinn/HF6\nvs/3+fQ8p+du/gRf0H/ie0/94iWnyvyPbkfSvqpVT4saAEnSehBpALVs/OFWoprZ\nKdtKuCACiBmGHCschxkRV9UUGVFEDQ0TLCPOsQplrjAKQqAm6XgVAkOww5ZYSaGJ\nIG5Rof7HXzuekGRJNhVJZchE0ERcJVDXFVVHjFGVShozOKeUmkg2dQkiU5MokYlG\nMDWwIalc48iQsMClx3GEQy0MIZVEk1zBVKMyUaDONQGESDY8o8PsBN7ChNvGm1M1\nYEcICCljmcyb/3fKrbNcl9n/+l3mCO1vlVuf8tQ6RuK/AXFiJaj4QFGWYbZjJRMg\nIgun6VoeQFYVCSkwrMIQYNtSls3ilufQdIgkcUIgZbOMNwY1DeFlGuQ61QxGDAZ1\nyTCJLhtExUjVDUyVsB42JRUyxTRkDcuywiRkhDnCOvBG2ppIgogi2sTVAulY1Qns\npm0GdvT3xSb6/EHfpECetwK+4JSCP4vRu2fyj7LCn5vO95VP2+jMel46ve/b3OyG\njoHDu0pocVOnYa0YGrx29VhbO/h8q7JluHGdFRut7CjomXO2u2NLNIQr5795dZpc\n+XBr2tjxHBs7dGaA3wvsHGwcKh/JP5DX+6T8kO7ezsaO7PeN5Hbff7fww9Zs07Kl\nu5yW3M2VRWvCB8HjpRezjdGmVrrv3Eh38mT/TlrjwPbQpzs/jGAq+WxW4dvKsRkL\nLn166RbldS1fFr8Drs/uHdo+83Xz9tU3Wi42ZFpjP9FoSdtDe2rL0ZNzvrad3hu4\nolVcfvF476J5pSNv1w32fEk3+7r82su1Dypao8NnAlWnhqO59Rfq8z/a74s7X/wC\nT2QsJA==\n=eyGB\n-----END PGP MESSAGE-----",
      "payload_json": "{\"body\":{\"key\":{\"fingerprint\":\"bc26fa3fa2ebebe45318d8d95aba18ffa461f3ed\",\"host\":\"keybase.io\",\"key_id\":\"5aba18ffa461f3ed\",\"kid\":\"0101c304e8c86c8f4b6773478eed4d05e9ffdddc81c7068c50db1b5bad9a904f5f890a\",\"uid\":\"bf65266d0d8df3ad5d1b367f578e6819\",\"username\":\"ralph\"},\"service\":{\"name\":\"twitter\",\"username\":\"testerralph\"},\"type\":\"web_service_binding\",\"version\":1},\"ctime\":1430836246,\"expire_in\":157680000,\"prev\":\"bdc9430e56f7d59eb9e6709cb719b4a8479ad3272c046e3c915a113e0892f8a7\",\"seqno\":3,\"tag\":\"signature\"}",
      "sig_type": 2,
      "ctime": 1430836246,
      "etime": 1588516246,
      "rtime": null,
      "sig_status": 0,
      "prev": "bdc9430e56f7d59eb9e6709cb719b4a8479ad3272c046e3c915a113e0892f8a7",
      "proof_id": "35dc42fe28fdab9085277e10",
      "proof_type": 2,
      "proof_text_check": "Verifying myself: I am ralph on Keybase.io. QgDwHcQWE5mVzVNNV371wUs5_eaablfM0_G2 /",
      "check_data_json": "{\"name\":\"twitter\",\"username\":\"testerralph\"}",
      "remote_id": "595596523460796418",
      "api_url": "https://twitter.com/testerralph/status/595596523460796418",
      "human_url": "https://twitter.com/testerralph/status/595596523460796418",
      "proof_state": 1,
      "proof_status": 1,
      "retry_count": 5,
      "hard_fail_count": 0,
      "last_check": 1431006765,
      "last_success": 1431006765,
      "version": 1,
      "fingerprint": "bc26fa3fa2ebebe45318d8d95aba18ffa461f3ed"
    },
    {
      "seqno": 4,
      "payload_hash": "6069efa710426907e15a734f5f15e32730aa02b133feb8cf718b22cf44faa640",
      "sig_id": "5c4da82af7e168cc4c6708be87d28c8558c7e8801b826d7af8b7049d246c32600f",
      "sig_id_short": "XE2oKvfhaMxMZwi-h9KMhVjH6IAbgm16-LcE",
      "kid": "0101c304e8c86c8f4b6773478eed4d05e9ffdddc81c7068c50db1b5bad9a904f5f890a",
      "sig": "-----BEGIN PGP MESSAGE-----\nComment: https://keybase.io/download\nVersion: Keybase Go 0.1.12 (linux)\n\nxA0DAAoBmv23eSsunT8By+F0AOIAAAAA63siYm9keSI6eyJkZXZpY2UiOnsiaWQi\nOiJmOWY1MmEyMzUwYmYwODA0ODA0MzQ3ZGJhODNjYTkxOCIsInR5cGUiOiJkZXNr\ndG9wIiwia2lkIjoiMDEyMGFmYWM5MzgzMTMyYjMxNzMzNjg2NmI0Y2U0ODgzNjkw\nNmNhZDg4ZGUxOGE5MmJkZGYyOWZmYjk0OWRiOGJjNDIwYSIsImRlc2NyaXB0aW9u\nIjoic3R1ZmYiLCJzdGF0dXMiOjF9LCJrZXkiOnsiZWxkZXN0X2tpZCI6IjAxMDFj\nMzA0ZThjODZjOGY0YjY3NzM0NzhlZWQ0ZDA1ZTlmZmRkZGM4MWM3MDY4YzUwZGIx\nYjViYWQ5YTkwNGY1Zjg5MGEiLCJmaW5nZXJwcmludCI6ImJjMjZmYTNmYTJlYmVi\nZTQ1MzE4ZDhkOTVhYmExOGZmYTQ2MWYzZWQiLCJob3N0Ijoia2V5YmFzZS5pbyIs\nImtleV9pZCI6IjVBQkExOEZGQTQ2MUYzRUQiLCJraWQiOiIwMTAxYzMwNGU4Yzg2\nYzhmNGI2NzczNDc4ZWVkNGQwNWU5ZmZkZGRjODFjNzA2OGM1MGRiMWI1YmFkOWE5\nMDRmNWY4OTBhIiwidWlkIjoiYmY2NTI2NmQwZDhkZjNhZDVkMWIzNjdmNTc4ZTY4\nMTkiLCJ1c2VybmFtZSI6InJhbHBoIn0sInNpYmtleSI6eyJraWQiOiIwMTIwYWZh\nYzkzODMxMzJiMzE3MzM2ODY2YjRjZTQ4ODM2OTA2Y2FkODhkZTE4YTkyYmRkZjI5\nZmZiOTQ5ZGI4YmM0MjBhIiwicmV2ZXJzZV9zaWciOiJnNlJpYjJSNWhxaGtaWFJo\nWTJobFpNT3BhR0Z6YUY5MGVYQmxDcU5yWlhuRUl3RWdyNnlUZ3hNck1YTTJobXRN\nNUlnMmtHeXRpTjRZcVN2ZDhwLzdsSjI0dkVJS3AzQmhlV3h2WVdURkJCeDdJbUp2\nWkhraU9uc2laR1YyYVdObElqcDdJbWxrSWpvaVpqbG1OVEpoTWpNMU1HSm1NRGd3\nTkRnd05ETTBOMlJpWVRnelkyRTVNVGdpTENKMGVYQmxJam9pWkdWemEzUnZjQ0lz\nSW10cFpDSTZJakF4TWpCaFptRmpPVE00TXpFek1tSXpNVGN6TXpZNE5qWmlOR05s\nTkRnNE16WTVNRFpqWVdRNE9HUmxNVGhoT1RKaVpHUm1NamxtWm1JNU5EbGtZamhp\nWXpReU1HRWlMQ0prWlhOamNtbHdkR2x2YmlJNkluTjBkV1ptSWl3aWMzUmhkSFZ6\nSWpveGZTd2lhMlY1SWpwN0ltVnNaR1Z6ZEY5cmFXUWlPaUl3TVRBeFl6TXdOR1U0\nWXpnMll6aG1OR0kyTnpjek5EYzRaV1ZrTkdRd05XVTVabVprWkdSak9ERmpOekEy\nT0dNMU1HUmlNV0kxWW1Ga09XRTVNRFJtTldZNE9UQmhJaXdpWm1sdVoyVnljSEpw\nYm5RaU9pSmlZekkyWm1FelptRXlaV0psWW1VME5UTXhPR1E0WkRrMVlXSmhNVGht\nWm1FME5qRm1NMlZrSWl3aWFHOXpkQ0k2SW10bGVXSmhjMlV1YVc4aUxDSnJaWGxm\nYVdRaU9pSTFRVUpCTVRoR1JrRTBOakZHTTBWRUlpd2lhMmxrSWpvaU1ERXdNV016\nTURSbE9HTTRObU00WmpSaU5qYzNNelEzT0dWbFpEUmtNRFZsT1dabVpHUmtZemd4\nWXpjd05qaGpOVEJrWWpGaU5XSmhaRGxoT1RBMFpqVm1PRGt3WVNJc0luVnBaQ0k2\nSW1KbU5qVXlOalprTUdRNFpHWXpZV1ExWkRGaU16WTNaalUzT0dVMk9ERTVJaXdp\nZFhObGNtNWhiV1VpT2lKeVlXeHdhQ0o5TENKemFXSnJaWGtpT25zaWEybGtJam9p\nTURFeU1HRm1ZV001TXpnek1UTXlZak14TnpNek5qZzJObUkwWTJVME9EZ3pOamt3\nTm1OaFpEZzRaR1V4T0dFNU1tSmtaR1l5T1dabVlqazBPV1JpT0dKak5ESXdZU0lz\nSW5KbGRtVnljMlZmYzJsbklqcHVkV3hzZlN3aWRIbHdaU0k2SW5OcFltdGxlU0lz\nSW5abGNuTnBiMjRpT2pGOUxDSmpiR2xsYm5RaU9uc2libUZ0WlNJNkltdGxlV0po\nYzJVdWFXOGdaMjhnWTJ4cFpXNTBJaXdpZG1WeWMybHZiaUk2SWpBdU1TNHhNaUo5\nTENKamRHbHRaU0k2TVRRek1EZ3pOak0zT1N3aVpYaHdhWEpsWDJsdUlqbzVORFl3\nT0RBd01Dd2liV1Z5YTJ4bFgzSnZiM1FpT25zaVkzUnBiV1VpT2pFME16QTRNell5\nTkRjc0ltaGhjMmdpT2lJNFkyRXpPV1UyTlRjME4yTTJNMlZtTkRjd05UazVZVGMx\nT0dVNE5qYzBNelV6T0ROaE5EaGlZemswWldVd09EY3pZamhpTnpVMFpUSXhPREk1\nTXpaaVptSmtZVE5tWkRVNFpUQmlaVEJoT0dNeU5ETXhZalV5WWpRMk5XRmxabVEz\nWkdZM1lXVm1ZakJtWW1NMU9UWTFOV1F6WVdWaU9UVTROamMwWVdVNFlTSXNJbk5s\nY1c1dklqbzNNRFI5TENKd2NtVjJJam9pTkdOak9XRmpOekl5Wm1aalpEVXpZalZr\nWlRreFlUTTNZ6WpOaE4yVmtaVGxsT0RCall6TTRPVE5sWWpZMFltUmpOVGs0T0dJ\neE9HVmlNalpoT1Rka05pSXNJbk5sY1c1dklqbzBMQ0owWVdjaU9pSnphV2R1WVhS\nMWNtVWlmYU56YVdmRVFQQ1pXa0dBMkM3WFIxMS9PQk84SXU1dHdxdVdzWHczL3ZY\nbnh5Sk50NWwyRmh0dGVueWxlOVo3Tis5S09WZXkvMVNvVk5Sb0sxU2RVT3U3S3lM\nYnJRQ29jMmxuWDNSNWNHVWdvM1JoWjgwQ0FxZDJaWEp6YVc5dUFRPT0ifSwidHlw\nZSI6InNpYmtleSIsInZlcnNpb24iOjF9LCJjbGllbnQiOnsibmFtZSI6ImtleWJh\nc2UuaW8gZ28gY2xpZW50IiwidmVyc2lvbiI6IjAuMS4xMiJ9LCJjdGltZSI6MTQz\nMDgzNjM3OSwiZXhwaXJlX2luIjo5NDYwODAwMCwibWVya2xlX3Jvb3QiOnsiY3Rp\nbWUiOjE0MzA4MzYyNDcsImhhc2giOiI4Y2EzOWU2NTc0N2M2M2VmNDcwNTk5YTc1\nOGU4Njc0MzUzODNhNDhiYzk0ZWUwODczYjhiNzU0ZTIxODI5MzZiZmJkYTNmZDU4\nZTBiZTBhOGMyNDMxYjUyYjQ2NWFlZmQ3ZGY3YWVmYjBmYmM1OTY1NWQz52FlYjk1\nODY3NGFlOGEiLCJzZXFubyI6NzA0fSwicHJldiI6IjRjYzlhYzcyMmZmY2Q1M2I1\nZGU5MWEzN2IzYTdlZGU5ZTgwY2MzODkzZWI2NGJkYzU5ODhiMThlYjI2YTk3ZDYi\nLCJzZXFubyI6NCwidGFnIjoic2lnbmF0dXJl4SJ9AMLAXAQAAQoAEAUCVUjUmwkQ\nmv23eSsunT8AAIv8CACQz2HZror6UIxV1ypJu1qluRnNq8oWz1rDdsZDeFk+JeUi\nVeACVwTbyvgofmTn0Yt/qXha4fbwU6+dCJIA31dSOvY4/JTVnZ2T39OrmCxYU8gr\nCyJlXqQGkErKlGKx300P+Ep6yZg8ZJgK3/1yLB5PhRmKSQRb9bEStF9xVCLVUefc\n84lUe9BnSD9XZpkfKmMwJkgIF2O2HQ919jPCqUIW8GkjWuILTcteyVAGv8fn3tn7\nESVwt5Cg6mwsjfC8zfDGLz81yw9xOHxhb+fgc1L01coqCxSVo4MYlV6KwOztsnbp\nbGdcbgZ/DD5krt8lsCABbPZeuiuGQlHHvhd6pQeG\n=E0qo\n-----END PGP MESSAGE-----",
      "payload_json": "{\"body\":{\"device\":{\"id\":\"f9f52a2350bf0804804347dba83ca918\",\"type\":\"desktop\",\"kid\":\"0120afac9383132b317336866b4ce48836906cad88de18a92bddf29ffb949db8bc420a\",\"description\":\"stuff\",\"status\":1},\"key\":{\"eldest_kid\":\"0101c304e8c86c8f4b6773478eed4d05e9ffdddc81c7068c50db1b5bad9a904f5f890a\",\"fingerprint\":\"bc26fa3fa2ebebe45318d8d95aba18ffa461f3ed\",\"host\":\"keybase.io\",\"key_id\":\"5ABA18FFA461F3ED\",\"kid\":\"0101c304e8c86c8f4b6773478eed4d05e9ffdddc81c7068c50db1b5bad9a904f5f890a\",\"uid\":\"bf65266d0d8df3ad5d1b367f578e6819\",\"username\":\"ralph\"},\"sibkey\":{\"kid\":\"0120afac9383132b317336866b4ce48836906cad88de18a92bddf29ffb949db8bc420a\",\"reverse_sig\":\"g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEgr6yTgxMrMXM2hmtM5Ig2kGytiN4YqSvd8p/7lJ24vEIKp3BheWxvYWTFBBx7ImJvZHkiOnsiZGV2aWNlIjp7ImlkIjoiZjlmNTJhMjM1MGJmMDgwNDgwNDM0N2RiYTgzY2E5MTgiLCJ0eXBlIjoiZGVza3RvcCIsImtpZCI6IjAxMjBhZmFjOTM4MzEzMmIzMTczMzY4NjZiNGNlNDg4MzY5MDZjYWQ4OGRlMThhOTJiZGRmMjlmZmI5NDlkYjhiYzQyMGEiLCJkZXNjcmlwdGlvbiI6InN0dWZmIiwic3RhdHVzIjoxfSwia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTAxYzMwNGU4Yzg2YzhmNGI2NzczNDc4ZWVkNGQwNWU5ZmZkZGRjODFjNzA2OGM1MGRiMWI1YmFkOWE5MDRmNWY4OTBhIiwiZmluZ2VycHJpbnQiOiJiYzI2ZmEzZmEyZWJlYmU0NTMxOGQ4ZDk1YWJhMThmZmE0NjFmM2VkIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJrZXlfaWQiOiI1QUJBMThGRkE0NjFGM0VEIiwia2lkIjoiMDEwMWMzMDRlOGM4NmM4ZjRiNjc3MzQ3OGVlZDRkMDVlOWZmZGRkYzgxYzcwNjhjNTBkYjFiNWJhZDlhOTA0ZjVmODkwYSIsInVpZCI6ImJmNjUyNjZkMGQ4ZGYzYWQ1ZDFiMzY3ZjU3OGU2ODE5IiwidXNlcm5hbWUiOiJyYWxwaCJ9LCJzaWJrZXkiOnsia2lkIjoiMDEyMGFmYWM5MzgzMTMyYjMxNzMzNjg2NmI0Y2U0ODgzNjkwNmNhZDg4ZGUxOGE5MmJkZGYyOWZmYjk0OWRiOGJjNDIwYSIsInJldmVyc2Vfc2lnIjpudWxsfSwidHlwZSI6InNpYmtleSIsInZlcnNpb24iOjF9LCJjbGllbnQiOnsibmFtZSI6ImtleWJhc2UuaW8gZ28gY2xpZW50IiwidmVyc2lvbiI6IjAuMS4xMiJ9LCJjdGltZSI6MTQzMDgzNjM3OSwiZXhwaXJlX2luIjo5NDYwODAwMCwibWVya2xlX3Jvb3QiOnsiY3RpbWUiOjE0MzA4MzYyNDcsImhhc2giOiI4Y2EzOWU2NTc0N2M2M2VmNDcwNTk5YTc1OGU4Njc0MzUzODNhNDhiYzk0ZWUwODczYjhiNzU0ZTIxODI5MzZiZmJkYTNmZDU4ZTBiZTBhOGMyNDMxYjUyYjQ2NWFlZmQ3ZGY3YWVmYjBmYmM1OTY1NWQzYWViOTU4Njc0YWU4YSIsInNlcW5vIjo3MDR9LCJwcmV2IjoiNGNjOWFjNzIyZmZjZDUzYjVkZTkxYTM3YjNhN2VkZTllODBjYzM4OTNlYjY0YmRjNTk4OGIxOGViMjZhOTdkNiIsInNlcW5vIjo0LCJ0YWciOiJzaWduYXR1cmUifaNzaWfEQPCZWkGA2C7XR11/OBO8Iu5twquWsXw3/vXnxyJNt5l2Fhttenyle9Z7N+9KOVey/1SoVNRoK1SdUOu7KyLbrQCoc2lnX3R5cGUgo3RhZ80CAqd2ZXJzaW9uAQ==\"},\"type\":\"sibkey\",\"version\":1},\"client\":{\"name\":\"keybase.io go client\",\"version\":\"0.1.12\"},\"ctime\":1430836379,\"expire_in\":94608000,\"merkle_root\":{\"ctime\":1430836247,\"hash\":\"8ca39e65747c63ef470599a758e867435383a48bc94ee0873b8b754e2182936bfbda3fd58e0be0a8c2431b52b465aefd7df7aefb0fbc59655d3aeb958674ae8a\",\"seqno\":704},\"prev\":\"4cc9ac722ffcd53b5de91a37b3a7ede9e80cc3893eb64bdc5988b18eb26a97d6\",\"seqno\":4,\"tag\":\"signature\"}",
      "sig_type": 1,
      "ctime": 1430836379,
      "etime": 1525444379,
      "rtime": null,
      "sig_status": 0,
      "prev": "4cc9ac722ffcd53b5de91a37b3a7ede9e80cc3893eb64bdc5988b18eb26a97d6",
      "proof_id": null,
      "proof_type": null,
      "proof_text_check": null,
      "check_data_json": null,
      "remote_id": null,
      "api_url": null,
      "human_url": null,
      "proof_state": null,
      "proof_status": null,
      "retry_count": null,
      "hard_fail_count": null,
      "last_check": null,
      "last_success": null,
      "version": null,
      "fingerprint": "bc26fa3fa2ebebe45318d8d95aba18ffa461f3ed"
    },
    {
      "seqno": 5,
      "payload_hash": "ba29de60fd420e579cfb241c3419bd781763a28f0ee9be18c5f58a1ae1e1a017",
      "sig_id": "b29a6b21766b48215e3ce7b75d6a9bd61b36e6212bf5b050b8b318cd972f742f0f",
      "sig_id_short": "spprIXZrSCFePOe3XWqb1hs25iEr9bBQuLMY",
      "kid": "0120afac9383132b317336866b4ce48836906cad88de18a92bddf29ffb949db8bc420a",
      "sig": "g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEgr6yTgxMrMXM2hmtM5Ig2kGytiN4YqSvd8p/7lJ24vEIKp3BheWxvYWTFBAp7ImJvZHkiOnsiZGV2aWNlIjp7ImlkIjoiZjlmNTJhMjM1MGJmMDgwNDgwNDM0N2RiYTgzY2E5MTgiLCJ0eXBlIjoiZGVza3RvcCIsImtpZCI6IjAxMjE2N2YxYjAyOTFhMDdmZGIxNGI3Mjk0NzUxNDdhYmYzN2QzMWM5NDc3ZDhkNjlkY2JmYTg3NDhjYjc2MDEwZDYxMGEiLCJkZXNjcmlwdGlvbiI6InN0dWZmIiwic3RhdHVzIjoxfSwia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTAxYzMwNGU4Yzg2YzhmNGI2NzczNDc4ZWVkNGQwNWU5ZmZkZGRjODFjNzA2OGM1MGRiMWI1YmFkOWE5MDRmNWY4OTBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwYWZhYzkzODMxMzJiMzE3MzM2ODY2YjRjZTQ4ODM2OTA2Y2FkODhkZTE4YTkyYmRkZjI5ZmZiOTQ5ZGI4YmM0MjBhIiwidWlkIjoiYmY2NTI2NmQwZDhkZjNhZDVkMWIzNjdmNTc4ZTY4MTkiLCJ1c2VybmFtZSI6InJhbHBoIn0sInN1YmtleSI6eyJraWQiOiIwMTIxNjdmMWIwMjkxYTA3ZmRiMTRiNzI5NDc1MTQ3YWJmMzdkMzFjOTQ3N2Q4ZDY5ZGNiZmE4NzQ4Y2I3NjAxMGQ2MTBhIiwicGFyZW50X2tpZCI6IjAxMjBhZmFjOTM4MzEzMmIzMTczMzY4NjZiNGNlNDg4MzY5MDZjYWQ4OGRlMThhOTJiZGRmMjlmZmI5NDlkYjhiYzQyMGEifSwidHlwZSI6InN1YmtleSIsInZlcnNpb24iOjF9LCJjbGllbnQiOnsibmFtZSI6ImtleWJhc2UuaW8gZ28gY2xpZW50IiwidmVyc2lvbiI6IjAuMS4xMiJ9LCJjdGltZSI6MTQzMDgzNjM3OSwiZXhwaXJlX2luIjo5NDYwODAwMCwibWVya2xlX3Jvb3QiOnsiY3RpbWUiOjE0MzA4MzYyNDcsImhhc2giOiI4Y2EzOWU2NTc0N2M2M2VmNDcwNTk5YTc1OGU4Njc0MzUzODNhNDhiYzk0ZWUwODczYjhiNzU0ZTIxODI5MzZiZmJkYTNmZDU4ZTBiZTBhOGMyNDMxYjUyYjQ2NWFlZmQ3ZGY3YWVmYjBmYmM1OTY1NWQzYWViOTU4Njc0YWU4YSIsInNlcW5vIjo3MDR9LCJwcmV2IjoiNjA2OWVmYTcxMDQyNjkwN2UxNWE3MzRmNWYxNWUzMjczMGFhMDJiMTMzZmViOGNmNzE4YjIyY2Y0NGZhYTY0MCIsInNlcW5vIjo1LCJ0YWciOiJzaWduYXR1cmUifaNzaWfEQBeo141bOEMnr0m/IHl72SGtJ/8wf9sVkw8uZRU0O3nZFpF0jrydXbQOMpZOv8c7eAEa7QdecLbdsmlZUtlaUAWoc2lnX3R5cGUgo3RhZ80CAqd2ZXJzaW9uAQ==",
      "payload_json": "{\"body\":{\"device\":{\"id\":\"f9f52a2350bf0804804347dba83ca918\",\"type\":\"desktop\",\"kid\":\"012167f1b0291a07fdb14b729475147abf37d31c9477d8d69dcbfa8748cb76010d610a\",\"description\":\"stuff\",\"status\":1},\"key\":{\"eldest_kid\":\"0101c304e8c86c8f4b6773478eed4d05e9ffdddc81c7068c50db1b5bad9a904f5f890a\",\"host\":\"keybase.io\",\"kid\":\"0120afac9383132b317336866b4ce48836906cad88de18a92bddf29ffb949db8bc420a\",\"uid\":\"bf65266d0d8df3ad5d1b367f578e6819\",\"username\":\"ralph\"},\"subkey\":{\"kid\":\"012167f1b0291a07fdb14b729475147abf37d31c9477d8d69dcbfa8748cb76010d610a\",\"parent_kid\":\"0120afac9383132b317336866b4ce48836906cad88de18a92bddf29ffb949db8bc420a\"},\"type\":\"subkey\",\"version\":1},\"client\":{\"name\":\"keybase.io go client\",\"version\":\"0.1.12\"},\"ctime\":1430836379,\"expire_in\":94608000,\"merkle_root\":{\"ctime\":1430836247,\"hash\":\"8ca39e65747c63ef470599a758e867435383a48bc94ee0873b8b754e2182936bfbda3fd58e0be0a8c2431b52b465aefd7df7aefb0fbc59655d3aeb958674ae8a\",\"seqno\":704},\"prev\":\"6069efa710426907e15a734f5f15e32730aa02b133feb8cf718b22cf44faa640\",\"seqno\":5,\"tag\":\"signature\"}",
      "sig_type": 1,
      "ctime": 1430836379,
      "etime": 1525444379,
      "rtime": null,
      "sig_status": 0,
      "prev": "6069efa710426907e15a734f5f15e32730aa02b133feb8cf718b22cf44faa640",
      "proof_id": null,
      "proof_type": null,
      "proof_text_check": null,
      "check_data_json": null,
      "remote_id": null,
      "api_url": null,
      "human_url": null,
      "proof_state": null,
      "proof_status": null,
      "retry_count": null,
      "hard_fail_count": null,
      "last_check": null,
      "last_success": null,
      "version": null,
      "fingerprint": ""
    },
    {
      "seqno": 6,
      "payload_hash": "cdd5ae013a6a879c605ba9611c3efff8dbfc473b9b14b2b99e8a6c00f62bbbba",
      "sig_id": "bce8a4dd488f24d42ce89ee667d549ab723d80e8b0258915af3fa8f3495dde620f",
      "sig_id_short": "vOik3UiPJNQs6J7mZ9VJq3I9gOiwJYkVrz-o",
      "kid": "0120afac9383132b317336866b4ce48836906cad88de18a92bddf29ffb949db8bc420a",
      "sig": "g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEgr6yTgxMrMXM2hmtM5Ig2kGytiN4YqSvd8p/7lJ24vEIKp3BheWxvYWTFCX97ImJvZHkiOnsiZGV2aWNlIjp7ImlkIjoiZmIxNWQ5ZDg1ODMxOThlNjI0ODZhNjIzYWYxYjRkMTgiLCJ0eXBlIjoid2ViIiwia2lkIjoiMDEyMGZkYWY4MTkyNzhjYjY5MDE3OGQzOGY4OWZmMWU0NDY0YzZhMzkzMjBmNzU2MzJiMWI1NzM0ODIwYmQxMDM3YzcwYSIsInN0YXR1cyI6MX0sImtleSI6eyJlbGRlc3Rfa2lkIjoiMDEwMWMzMDRlOGM4NmM4ZjRiNjc3MzQ3OGVlZDRkMDVlOWZmZGRkYzgxYzcwNjhjNTBkYjFiNWJhZDlhOTA0ZjVmODkwYSIsImhvc3QiOiJrZXliYXNlLmlvIiwia2lkIjoiMDEyMGFmYWM5MzgzMTMyYjMxNzMzNjg2NmI0Y2U0ODgzNjkwNmNhZDg4ZGUxOGE5MmJkZGYyOWZmYjk0OWRiOGJjNDIwYSIsInVpZCI6ImJmNjUyNjZkMGQ4ZGYzYWQ1ZDFiMzY3ZjU3OGU2ODE5IiwidXNlcm5hbWUiOiJyYWxwaCJ9LCJzaWJrZXkiOnsia2lkIjoiMDEyMGZkYWY4MTkyNzhjYjY5MDE3OGQzOGY4OWZmMWU0NDY0YzZhMzkzMjBmNzU2MzJiMWI1NzM0ODIwYmQxMDM3YzcwYSIsInJldmVyc2Vfc2lnIjoiZzZSaWIyUjVocWhrWlhSaFkyaGxaTU9wYUdGemFGOTBlWEJsQ3FOclpYbkVJd0VnL2ErQmtuakxhUUY0MDQrSi94NUVaTWFqa3lEM1ZqS3h0WE5JSUwwUU44Y0twM0JoZVd4dllXVEZBNjE3SW1KdlpIa2lPbnNpWkdWMmFXTmxJanA3SW1sa0lqb2labUl4TldRNVpEZzFPRE14T1RobE5qSTBPRFpoTmpJellXWXhZalJrTVRnaUxDSjBlWEJsSWpvaWQyVmlJaXdpYTJsa0lqb2lNREV5TUdaa1lXWTRNVGt5TnpoallqWTVNREUzT0dRek9HWTRPV1ptTVdVME5EWTBZelpoTXprek1qQm1OelUyTXpKaU1XSTFOek0wT0RJd1ltUXhNRE0zWXpjd1lTSXNJbk4wWVhSMWN5STZNWDBzSW10bGVTSTZleUpsYkdSbGMzUmZhMmxrSWpvaU1ERXdNV016TURSbE9HTTRObU00WmpSaU5qYzNNelEzT0dWbFpEUmtNRFZsT1dabVpHUmtZemd4WXpjd05qaGpOVEJrWWpGaU5XSmhaRGxoT1RBMFpqVm1PRGt3WVNJc0ltaHZjM1FpT2lKclpYbGlZWE5sTG1sdklpd2lhMmxrSWpvaU1ERXlNR0ZtWVdNNU16Z3pNVE15WWpNeE56TXpOamcyTm1JMFkyVTBPRGd6Tmprd05tTmhaRGc0WkdVeE9HRTVNbUprWkdZeU9XWm1ZamswT1dSaU9HSmpOREl3WVNJc0luVnBaQ0k2SW1KbU5qVXlOalprTUdRNFpHWXpZV1ExWkRGaU16WTNaalUzT0dVMk9ERTVJaXdpZFhObGNtNWhiV1VpT2lKeVlXeHdhQ0o5TENKemFXSnJaWGtpT25zaWEybGtJam9pTURFeU1HWmtZV1k0TVRreU56aGpZalk1TURFM09HUXpPR1k0T1dabU1XVTBORFkwWXpaaE16a3pNakJtTnpVMk16SmlNV0kxTnpNME9ESXdZbVF4TURNM1l6Y3dZU0lzSW5KbGRtVnljMlZmYzJsbklqcHVkV3hzZlN3aWRIbHdaU0k2SW5OcFltdGxlU0lzSW5abGNuTnBiMjRpT2pGOUxDSmpiR2xsYm5RaU9uc2libUZ0WlNJNkltdGxlV0poYzJVdWFXOGdaMjhnWTJ4cFpXNTBJaXdpZG1WeWMybHZiaUk2SWpBdU1TNHhNaUo5TENKamRHbHRaU0k2TVRRek1EZ3pOak0zT1N3aVpYaHdhWEpsWDJsdUlqbzVORFl3T0RBd01Dd2liV1Z5YTJ4bFgzSnZiM1FpT25zaVkzUnBiV1VpT2pFME16QTRNell5TkRjc0ltaGhjMmdpT2lJNFkyRXpPV1UyTlRjME4yTTJNMlZtTkRjd05UazVZVGMxT0dVNE5qYzBNelV6T0ROaE5EaGlZemswWldVd09EY3pZamhpTnpVMFpUSXhPREk1TXpaaVptSmtZVE5tWkRVNFpUQmlaVEJoT0dNeU5ETXhZalV5WWpRMk5XRmxabVEzWkdZM1lXVm1ZakJtWW1NMU9UWTFOV1F6WVdWaU9UVTROamMwWVdVNFlTSXNJbk5sY1c1dklqbzNNRFI5TENKd2NtVjJJam9pWW1FeU9XUmxOakJtWkRReU1HVTFOemxqWm1JeU5ERmpNelF4T1dKa056Z3hOell6WVRJNFpqQmxaVGxpWlRFNFl6Vm1OVGhoTVdGbE1XVXhZVEF4TnlJc0luTmxjVzV2SWpvMkxDSjBZV2NpT2lKemFXZHVZWFIxY21VaWZhTnphV2ZFUUN6UGxXT216WE0yeXFsZjl4dnpJL3N0OU5vWmwwTjFDNWR3YUl6UGdFK1MzN3hQbGI1L1ExNFUvU0dnbjRxSXZ1VHNERjBucS8xeE90QUJpYUo0eVFxb2MybG5YM1I1Y0dVZ28zUmhaODBDQXFkMlpYSnphVzl1QVE9PSJ9LCJ0eXBlIjoic2lia2V5IiwidmVyc2lvbiI6MX0sImNsaWVudCI6eyJuYW1lIjoia2V5YmFzZS5pbyBnbyBjbGllbnQiLCJ2ZXJzaW9uIjoiMC4xLjEyIn0sImN0aW1lIjoxNDMwODM2Mzc5LCJleHBpcmVfaW4iOjk0NjA4MDAwLCJtZXJrbGVfcm9vdCI6eyJjdGltZSI6MTQzMDgzNjI0NywiaGFzaCI6IjhjYTM5ZTY1NzQ3YzYzZWY0NzA1OTlhNzU4ZTg2NzQzNTM4M2E0OGJjOTRlZTA4NzNiOGI3NTRlMjE4MjkzNmJmYmRhM2ZkNThlMGJlMGE4YzI0MzFiNTJiNDY1YWVmZDdkZjdhZWZiMGZiYzU5NjU1ZDNhZWI5NTg2NzRhZThhIiwic2Vxbm8iOjcwNH0sInByZXYiOiJiYTI5ZGU2MGZkNDIwZTU3OWNmYjI0MWMzNDE5YmQ3ODE3NjNhMjhmMGVlOWJlMThjNWY1OGExYWUxZTFhMDE3Iiwic2Vxbm8iOjYsInRhZyI6InNpZ25hdHVyZSJ9o3NpZ8RAVPFae3R7YfGC+zzBXi6Kffak6g6aHRWy+zyPRMEqce12ElhFTwr5ahmG3hu+5ZmT9ZZ0nj76hnKzUpPKEXr4DKhzaWdfdHlwZSCjdGFnzQICp3ZlcnNpb24B",
      "payload_json": "{\"body\":{\"device\":{\"id\":\"fb15d9d8583198e62486a623af1b4d18\",\"type\":\"web\",\"kid\":\"0120fdaf819278cb690178d38f89ff1e4464c6a39320f75632b1b5734820bd1037c70a\",\"status\":1},\"key\":{\"eldest_kid\":\"0101c304e8c86c8f4b6773478eed4d05e9ffdddc81c7068c50db1b5bad9a904f5f890a\",\"host\":\"keybase.io\",\"kid\":\"0120afac9383132b317336866b4ce48836906cad88de18a92bddf29ffb949db8bc420a\",\"uid\":\"bf65266d0d8df3ad5d1b367f578e6819\",\"username\":\"ralph\"},\"sibkey\":{\"kid\":\"0120fdaf819278cb690178d38f89ff1e4464c6a39320f75632b1b5734820bd1037c70a\",\"reverse_sig\":\"g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEg/a+BknjLaQF404+J/x5EZMajkyD3VjKxtXNIIL0QN8cKp3BheWxvYWTFA617ImJvZHkiOnsiZGV2aWNlIjp7ImlkIjoiZmIxNWQ5ZDg1ODMxOThlNjI0ODZhNjIzYWYxYjRkMTgiLCJ0eXBlIjoid2ViIiwia2lkIjoiMDEyMGZkYWY4MTkyNzhjYjY5MDE3OGQzOGY4OWZmMWU0NDY0YzZhMzkzMjBmNzU2MzJiMWI1NzM0ODIwYmQxMDM3YzcwYSIsInN0YXR1cyI6MX0sImtleSI6eyJlbGRlc3Rfa2lkIjoiMDEwMWMzMDRlOGM4NmM4ZjRiNjc3MzQ3OGVlZDRkMDVlOWZmZGRkYzgxYzcwNjhjNTBkYjFiNWJhZDlhOTA0ZjVmODkwYSIsImhvc3QiOiJrZXliYXNlLmlvIiwia2lkIjoiMDEyMGFmYWM5MzgzMTMyYjMxNzMzNjg2NmI0Y2U0ODgzNjkwNmNhZDg4ZGUxOGE5MmJkZGYyOWZmYjk0OWRiOGJjNDIwYSIsInVpZCI6ImJmNjUyNjZkMGQ4ZGYzYWQ1ZDFiMzY3ZjU3OGU2ODE5IiwidXNlcm5hbWUiOiJyYWxwaCJ9LCJzaWJrZXkiOnsia2lkIjoiMDEyMGZkYWY4MTkyNzhjYjY5MDE3OGQzOGY4OWZmMWU0NDY0YzZhMzkzMjBmNzU2MzJiMWI1NzM0ODIwYmQxMDM3YzcwYSIsInJldmVyc2Vfc2lnIjpudWxsfSwidHlwZSI6InNpYmtleSIsInZlcnNpb24iOjF9LCJjbGllbnQiOnsibmFtZSI6ImtleWJhc2UuaW8gZ28gY2xpZW50IiwidmVyc2lvbiI6IjAuMS4xMiJ9LCJjdGltZSI6MTQzMDgzNjM3OSwiZXhwaXJlX2luIjo5NDYwODAwMCwibWVya2xlX3Jvb3QiOnsiY3RpbWUiOjE0MzA4MzYyNDcsImhhc2giOiI4Y2EzOWU2NTc0N2M2M2VmNDcwNTk5YTc1OGU4Njc0MzUzODNhNDhiYzk0ZWUwODczYjhiNzU0ZTIxODI5MzZiZmJkYTNmZDU4ZTBiZTBhOGMyNDMxYjUyYjQ2NWFlZmQ3ZGY3YWVmYjBmYmM1OTY1NWQzYWViOTU4Njc0YWU4YSIsInNlcW5vIjo3MDR9LCJwcmV2IjoiYmEyOWRlNjBmZDQyMGU1NzljZmIyNDFjMzQxOWJkNzgxNzYzYTI4ZjBlZTliZTE4YzVmNThhMWFlMWUxYTAxNyIsInNlcW5vIjo2LCJ0YWciOiJzaWduYXR1cmUifaNzaWfEQCzPlWOmzXM2yqlf9xvzI/st9NoZl0N1C5dwaIzPgE+S37xPlb5/Q14U/SGgn4qIvuTsDF0nq/1xOtABiaJ4yQqoc2lnX3R5cGUgo3RhZ80CAqd2ZXJzaW9uAQ==\"},\"type\":\"sibkey\",\"version\":1},\"client\":{\"name\":\"keybase.io go client\",\"version\":\"0.1.12\"},\"ctime\":1430836379,\"expire_in\":94608000,\"merkle_root\":{\"ctime\":1430836247,\"hash\":\"8ca39e65747c63ef470599a758e867435383a48bc94ee0873b8b754e2182936bfbda3fd58e0be0a8c2431b52b465aefd7df7aefb0fbc59655d3aeb958674ae8a\",\"seqno\":704},\"prev\":\"ba29de60fd420e579cfb241c3419bd781763a28f0ee9be18c5f58a1ae1e1a017\",\"seqno\":6,\"tag\":\"signature\"}",
      "sig_type": 1,
      "ctime": 1430836379,
      "etime": 1525444379,
      "rtime": null,
      "sig_status": 0,
      "prev": "ba29de60fd420e579cfb241c3419bd781763a28f0ee9be18c5f58a1ae1e1a017",
      "proof_id": null,
      "proof_type": null,
      "proof_text_check": null,
      "check_data_json": null,
      "remote_id": null,
      "api_url": null,
      "human_url": null,
      "proof_state": null,
      "proof_status": null,
      "retry_count": null,
      "hard_fail_count": null,
      "last_check": null,
      "last_success": null,
      "version": null,
      "fingerprint": ""
    },
    {
      "seqno": 7,
      "payload_hash": "ab5399ff8a1bc85c80e621d19462ced75095544c84a854f9b9f1dbde90fd1a0e",
      "sig_id": "c55da1d425a0f9ee917b5cd3b35b31135ccb0241e51ff93cb926dbb3bba1f98b0f",
      "sig_id_short": "xV2h1CWg-e6Re1zTs1sxE1zLAkHlH_k8uSbb",
      "kid": "0120fdaf819278cb690178d38f89ff1e4464c6a39320f75632b1b5734820bd1037c70a",
      "sig": "g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEg/a+BknjLaQF404+J/x5EZMajkyD3VjKxtXNIIL0QN8cKp3BheWxvYWTFA/B7ImJvZHkiOnsiZGV2aWNlIjp7ImlkIjoiZmIxNWQ5ZDg1ODMxOThlNjI0ODZhNjIzYWYxYjRkMTgiLCJ0eXBlIjoid2ViIiwia2lkIjoiMDEyMWI4NWIyNWFkZmRjNDgzYWUwMDQzNTU3NGJkODI3MmM1Y2M5MzZjNDYzZDE3OWQ4ODkwMGM0Yjk2YTM3NGNkNmMwYSIsInN0YXR1cyI6MX0sImtleSI6eyJlbGRlc3Rfa2lkIjoiMDEwMWMzMDRlOGM4NmM4ZjRiNjc3MzQ3OGVlZDRkMDVlOWZmZGRkYzgxYzcwNjhjNTBkYjFiNWJhZDlhOTA0ZjVmODkwYSIsImhvc3QiOiJrZXliYXNlLmlvIiwia2lkIjoiMDEyMGZkYWY4MTkyNzhjYjY5MDE3OGQzOGY4OWZmMWU0NDY0YzZhMzkzMjBmNzU2MzJiMWI1NzM0ODIwYmQxMDM3YzcwYSIsInVpZCI6ImJmNjUyNjZkMGQ4ZGYzYWQ1ZDFiMzY3ZjU3OGU2ODE5IiwidXNlcm5hbWUiOiJyYWxwaCJ9LCJzdWJrZXkiOnsia2lkIjoiMDEyMWI4NWIyNWFkZmRjNDgzYWUwMDQzNTU3NGJkODI3MmM1Y2M5MzZjNDYzZDE3OWQ4ODkwMGM0Yjk2YTM3NGNkNmMwYSIsInBhcmVudF9raWQiOiIwMTIwZmRhZjgxOTI3OGNiNjkwMTc4ZDM4Zjg5ZmYxZTQ0NjRjNmEzOTMyMGY3NTYzMmIxYjU3MzQ4MjBiZDEwMzdjNzBhIn0sInR5cGUiOiJzdWJrZXkiLCJ2ZXJzaW9uIjoxfSwiY2xpZW50Ijp7Im5hbWUiOiJrZXliYXNlLmlvIGdvIGNsaWVudCIsInZlcnNpb24iOiIwLjEuMTIifSwiY3RpbWUiOjE0MzA4MzYzNzksImV4cGlyZV9pbiI6OTQ2MDgwMDAsIm1lcmtsZV9yb290Ijp7ImN0aW1lIjoxNDMwODM2MjQ3LCJoYXNoIjoiOGNhMzllNjU3NDdjNjNlZjQ3MDU5OWE3NThlODY3NDM1MzgzYTQ4YmM5NGVlMDg3M2I4Yjc1NGUyMTgyOTM2YmZiZGEzZmQ1OGUwYmUwYThjMjQzMWI1MmI0NjVhZWZkN2RmN2FlZmIwZmJjNTk2NTVkM2FlYjk1ODY3NGFlOGEiLCJzZXFubyI6NzA0fSwicHJldiI6ImNkZDVhZTAxM2E2YTg3OWM2MDViYTk2MTFjM2VmZmY4ZGJmYzQ3M2I5YjE0YjJiOTllOGE2YzAwZjYyYmJiYmEiLCJzZXFubyI6NywidGFnIjoic2lnbmF0dXJlIn2jc2lnxEBHqe4TZEOiNGSd8Vanr53/W1aMX1wD7PSk8ag8lnvN2jmzrfzIz6PqRyiFFVRFnDuk3buCEDvQZ2EsA/aZfLUEqHNpZ190eXBlIKN0YWfNAgKndmVyc2lvbgE=",
      "payload_json": "{\"body\":{\"device\":{\"id\":\"fb15d9d8583198e62486a623af1b4d18\",\"type\":\"web\",\"kid\":\"0121b85b25adfdc483ae00435574bd8272c5cc936c463d179d88900c4b96a374cd6c0a\",\"status\":1},\"key\":{\"eldest_kid\":\"0101c304e8c86c8f4b6773478eed4d05e9ffdddc81c7068c50db1b5bad9a904f5f890a\",\"host\":\"keybase.io\",\"kid\":\"0120fdaf819278cb690178d38f89ff1e4464c6a39320f75632b1b5734820bd1037c70a\",\"uid\":\"bf65266d0d8df3ad5d1b367f578e6819\",\"username\":\"ralph\"},\"subkey\":{\"kid\":\"0121b85b25adfdc483ae00435574bd8272c5cc936c463d179d88900c4b96a374cd6c0a\",\"parent_kid\":\"0120fdaf819278cb690178d38f89ff1e4464c6a39320f75632b1b5734820bd1037c70a\"},\"type\":\"subkey\",\"version\":1},\"client\":{\"name\":\"keybase.io go client\",\"version\":\"0.1.12\"},\"ctime\":1430836379,\"expire_in\":94608000,\"merkle_root\":{\"ctime\":1430836247,\"hash\":\"8ca39e65747c63ef470599a758e867435383a48bc94ee0873b8b754e2182936bfbda3fd58e0be0a8c2431b52b465aefd7df7aefb0fbc59655d3aeb958674ae8a\",\"seqno\":704},\"prev\":\"cdd5ae013a6a879c605ba9611c3efff8dbfc473b9b14b2b99e8a6c00f62bbbba\",\"seqno\":7,\"tag\":\"signature\"}",
      "sig_type": 1,
      "ctime": 1430836379,
      "etime": 1525444379,
      "rtime": null,
      "sig_status": 0,
      "prev": "cdd5ae013a6a879c605ba9611c3efff8dbfc473b9b14b2b99e8a6c00f62bbbba",
      "proof_id": null,
      "proof_type": null,
      "proof_text_check": null,
      "check_data_json": null,
      "remote_id": null,
      "api_url": null,
      "human_url": null,
      "proof_state": null,
      "proof_status": null,
      "retry_count": null,
      "hard_fail_count": null,
      "last_check": null,
      "last_success": null,
      "version": null,
      "fingerprint": ""
    }
  ],
  "keys": [
    "-----BEGIN PGP PUBLIC KEY BLOCK-----\nVersion: Keybase OpenPGP v2.0.8\nComment: https://keybase.io/crypto\n\n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\n=BVOL\n-----END PGP PUBLIC KEY BLOCK-----",
    "-----BEGIN PGP PUBLIC KEY BLOCK-----\nVersion: Keybase OpenPGP v2.0.8\nComment: https://keybase.io/crypto\n\n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\n=sF0S\n-----END PGP PUBLIC KEY BLOCK-----",
    "0120afac9383132b317336866b4ce48836906cad88de18a92bddf29ffb949db8bc420a",
    "012167f1b0291a07fdb14b729475147abf37d31c9477d8d69dcbfa8748cb76010d610a",
    "0121b85b25adfdc483ae00435574bd8272c5cc936c463d179d88900c4b96a374cd6c0a",
    "0120fdaf819278cb690178d38f89ff1e4464c6a39320f75632b1b5734820bd1037c70a"
  ],
  "label_kids": {
    "second_eldest": "0101c304e8c86c8f4b6773478eed4d05e9ffdddc81c7068c50db1b5bad9a904f5f890a"
  }
}
`,
	"signed_with_revoked_key_chain.json": `{
    "chain": [
        {
            "seqno": 1,
            "prev": null,
            "sig": "g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEgNkE13AJPYlXXSNyO4DTlRVMLDNln4Gji/wPNwVgvxwYKp3BheWxvYWTFASp7ImJvZHkiOnsia2V5Ijp7Imhvc3QiOiJrZXliYXNlLmlvIiwia2lkIjoiMDEyMDM2NDEzNWRjMDI0ZjYyNTVkNzQ4ZGM4ZWUwMzRlNTQ1NTMwYjBjZDk2N2UwNjhlMmZmMDNjZGMxNTgyZmM3MDYwYSIsInVpZCI6Ijc0YzM4Y2Y3Y2ViOTQ3ZjU2MzIwNDVkOGNhNWQ0ODE5IiwidXNlcm5hbWUiOiJtYXgzMiJ9LCJ0eXBlIjoiZWxkZXN0IiwidmVyc2lvbiI6MX0sImN0aW1lIjoxNDMyMTQ1NzYzLCJleHBpcmVfaW4iOjEwMDAwMDAwLCJwcmV2IjpudWxsLCJzZXFfdHlwZSI6MSwic2Vxbm8iOjEsInRhZyI6InNpZ25hdHVyZSJ9o3NpZ8RAaH1Gcdn5h66O/2llNDAaZKGzZAEQBuXgTN9JuxAa8hV2up2yn144/aRucaRx3lE/fYVDJsgy6lg/fJI4c+i6CahzaWdfdHlwZSCjdGFnzQICp3ZlcnNpb24B",
            "payload_hash": "9aa395f90c7c4cc42a9e3934d7b88d85318c7a8f4be72def31b3ed9e1f1e8088",
            "sig_id": "53e8e98d74d1bb3ab2710a2ce955e6dfc111c400478d7f4dd18a359cdd79350f0f",
            "payload_json": "{\"body\":{\"key\":{\"host\":\"keybase.io\",\"kid\":\"0120364135dc024f6255d748dc8ee034e545530b0cd967e068e2ff03cdc1582fc7060a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"type\":\"eldest\",\"version\":1},\"ctime\":1432145763,\"expire_in\":10000000,\"prev\":null,\"seq_type\":1,\"seqno\":1,\"tag\":\"signature\"}",
            "kid": "0120364135dc024f6255d748dc8ee034e545530b0cd967e068e2ff03cdc1582fc7060a",
            "ctime": 1432145763
        },
        {
            "seqno": 2,
            "prev": "9aa395f90c7c4cc42a9e3934d7b88d85318c7a8f4be72def31b3ed9e1f1e8088",
            "sig": "g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEgNkE13AJPYlXXSNyO4DTlRVMLDNln4Gji/wPNwVgvxwYKp3BheWxvYWTFBfl7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTIwMzY0MTM1ZGMwMjRmNjI1NWQ3NDhkYzhlZTAzNGU1NDU1MzBiMGNkOTY3ZTA2OGUyZmYwM2NkYzE1ODJmYzcwNjBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwMzY0MTM1ZGMwMjRmNjI1NWQ3NDhkYzhlZTAzNGU1NDU1MzBiMGNkOTY3ZTA2OGUyZmYwM2NkYzE1ODJmYzcwNjBhIiwidWlkIjoiNzRjMzhjZjdjZWI5NDdmNTYzMjA0NWQ4Y2E1ZDQ4MTkiLCJ1c2VybmFtZSI6Im1heDMyIn0sInNpYmtleSI6eyJraWQiOiIwMTIwNzg3OTA4ODUxZDU3OTEzYjQwZTRiNzllMjFjZmNlM2NhNzMxZTMxMDA1MjUwOWVmZWU3NWI4OTY5ZWE3NWY0ZDBhIiwicmV2ZXJzZV9zaWciOiJnNlJpYjJSNWhxaGtaWFJoWTJobFpNT3BhR0Z6YUY5MGVYQmxDcU5yWlhuRUl3RWdlSGtJaFIxWGtUdEE1TGVlSWMvT1BLY3g0eEFGSlFudjduVzRscDZuWDAwS3AzQmhlV3h2WVdURkFpdDdJbUp2WkhraU9uc2lhMlY1SWpwN0ltVnNaR1Z6ZEY5cmFXUWlPaUl3TVRJd016WTBNVE0xWkdNd01qUm1OakkxTldRM05EaGtZemhsWlRBek5HVTFORFUxTXpCaU1HTmtPVFkzWlRBMk9HVXlabVl3TTJOa1l6RTFPREptWXpjd05qQmhJaXdpYUc5emRDSTZJbXRsZVdKaGMyVXVhVzhpTENKcmFXUWlPaUl3TVRJd016WTBNVE0xWkdNd01qUm1OakkxTldRM05EaGtZemhsWlRBek5HVTFORFUxTXpCaU1HTmtPVFkzWlRBMk9HVXlabVl3TTJOa1l6RTFPREptWXpjd05qQmhJaXdpZFdsa0lqb2lOelJqTXpoalpqZGpaV0k1TkRkbU5UWXpNakEwTldRNFkyRTFaRFE0TVRraUxDSjFjMlZ5Ym1GdFpTSTZJbTFoZURNeUluMHNJbk5wWW10bGVTSTZleUpyYVdRaU9pSXdNVEl3TnpnM09UQTRPRFV4WkRVM09URXpZalF3WlRSaU56bGxNakZqWm1ObE0yTmhOek14WlRNeE1EQTFNalV3T1dWbVpXVTNOV0k0T1RZNVpXRTNOV1kwWkRCaElpd2ljbVYyWlhKelpWOXphV2NpT201MWJHeDlMQ0owZVhCbElqb2ljMmxpYTJWNUlpd2lkbVZ5YzJsdmJpSTZNWDBzSW1OMGFXMWxJam94TkRNeU1UUTFPRFl6TENKbGVIQnBjbVZmYVc0aU9qRXdNREF3TURBd0xDSndjbVYySWpvaU9XRmhNemsxWmprd1l6ZGpOR05qTkRKaE9XVXpPVE0wWkRkaU9EaGtPRFV6TVRoak4yRTRaalJpWlRjeVpHVm1NekZpTTJWa09XVXhaakZsT0RBNE9DSXNJbk5sY1Y5MGVYQmxJam94TENKelpYRnVieUk2TWl3aWRHRm5Jam9pYzJsbmJtRjBkWEpsSW4yamMybG54RURpc2tKVmNVc042bkM4YzdCbmdLR05aYU1XL0duMWlhZTBhMFdMSFJ2bnlQYkxKcUoxYUEzSnF6ZDVLY3NqVGFFSzMvSktqQkZ3NDUzOHRDK3QwT0VQcUhOcFoxOTBlWEJsSUtOMFlXZk5BZ0tuZG1WeWMybHZiZ0U9In0sInR5cGUiOiJzaWJrZXkiLCJ2ZXJzaW9uIjoxfSwiY3RpbWUiOjE0MzIxNDU4NjMsImV4cGlyZV9pbiI6MTAwMDAwMDAsInByZXYiOiI5YWEzOTVmOTBjN2M0Y2M0MmE5ZTM5MzRkN2I4OGQ4NTMxOGM3YThmNGJlNzJkZWYzMWIzZWQ5ZTFmMWU4MDg4Iiwic2VxX3R5cGUiOjEsInNlcW5vIjoyLCJ0YWciOiJzaWduYXR1cmUifaNzaWfEQCGlvk5VtnWDNGcSe2rq1u00pf9m8w/yLdiq/0A/d6MYhtgWrtny6hiGamcrPT1bS6e9P7XDreP7m3Vtd82Kuguoc2lnX3R5cGUgo3RhZ80CAqd2ZXJzaW9uAQ==",
            "payload_hash": "d845211b2ba21e87141d51553c43ecc3701fbf37203d63e5b82ede4c0fa023bc",
            "sig_id": "ecdbdc174a75a79fb175b9c4729765eaf0d03d1b9d910f244f90293b6f96e8b40f",
            "payload_json": "{\"body\":{\"key\":{\"eldest_kid\":\"0120364135dc024f6255d748dc8ee034e545530b0cd967e068e2ff03cdc1582fc7060a\",\"host\":\"keybase.io\",\"kid\":\"0120364135dc024f6255d748dc8ee034e545530b0cd967e068e2ff03cdc1582fc7060a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"sibkey\":{\"kid\":\"0120787908851d57913b40e4b79e21cfce3ca731e310052509efee75b8969ea75f4d0a\",\"reverse_sig\":\"g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEgeHkIhR1XkTtA5LeeIc/OPKcx4xAFJQnv7nW4lp6nX00Kp3BheWxvYWTFAit7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTIwMzY0MTM1ZGMwMjRmNjI1NWQ3NDhkYzhlZTAzNGU1NDU1MzBiMGNkOTY3ZTA2OGUyZmYwM2NkYzE1ODJmYzcwNjBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwMzY0MTM1ZGMwMjRmNjI1NWQ3NDhkYzhlZTAzNGU1NDU1MzBiMGNkOTY3ZTA2OGUyZmYwM2NkYzE1ODJmYzcwNjBhIiwidWlkIjoiNzRjMzhjZjdjZWI5NDdmNTYzMjA0NWQ4Y2E1ZDQ4MTkiLCJ1c2VybmFtZSI6Im1heDMyIn0sInNpYmtleSI6eyJraWQiOiIwMTIwNzg3OTA4ODUxZDU3OTEzYjQwZTRiNzllMjFjZmNlM2NhNzMxZTMxMDA1MjUwOWVmZWU3NWI4OTY5ZWE3NWY0ZDBhIiwicmV2ZXJzZV9zaWciOm51bGx9LCJ0eXBlIjoic2lia2V5IiwidmVyc2lvbiI6MX0sImN0aW1lIjoxNDMyMTQ1ODYzLCJleHBpcmVfaW4iOjEwMDAwMDAwLCJwcmV2IjoiOWFhMzk1ZjkwYzdjNGNjNDJhOWUzOTM0ZDdiODhkODUzMThjN2E4ZjRiZTcyZGVmMzFiM2VkOWUxZjFlODA4OCIsInNlcV90eXBlIjoxLCJzZXFubyI6MiwidGFnIjoic2lnbmF0dXJlIn2jc2lnxEDiskJVcUsN6nC8c7BngKGNZaMW/Gn1iae0a0WLHRvnyPbLJqJ1aA3Jqzd5KcsjTaEK3/JKjBFw4538tC+t0OEPqHNpZ190eXBlIKN0YWfNAgKndmVyc2lvbgE=\"},\"type\":\"sibkey\",\"version\":1},\"ctime\":1432145863,\"expire_in\":10000000,\"prev\":\"9aa395f90c7c4cc42a9e3934d7b88d85318c7a8f4be72def31b3ed9e1f1e8088\",\"seq_type\":1,\"seqno\":2,\"tag\":\"signature\"}",
            "kid": "0120364135dc024f6255d748dc8ee034e545530b0cd967e068e2ff03cdc1582fc7060a",
            "ctime": 1432145863
        },
        {
            "seqno": 3,
            "prev": "d845211b2ba21e87141d51553c43ecc3701fbf37203d63e5b82ede4c0fa023bc",
            "sig": "g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEgNkE13AJPYlXXSNyO4DTlRVMLDNln4Gji/wPNwVgvxwYKp3BheWxvYWTFAhh7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTIwMzY0MTM1ZGMwMjRmNjI1NWQ3NDhkYzhlZTAzNGU1NDU1MzBiMGNkOTY3ZTA2OGUyZmYwM2NkYzE1ODJmYzcwNjBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwMzY0MTM1ZGMwMjRmNjI1NWQ3NDhkYzhlZTAzNGU1NDU1MzBiMGNkOTY3ZTA2OGUyZmYwM2NkYzE1ODJmYzcwNjBhIiwidWlkIjoiNzRjMzhjZjdjZWI5NDdmNTYzMjA0NWQ4Y2E1ZDQ4MTkiLCJ1c2VybmFtZSI6Im1heDMyIn0sInJldm9rZSI6eyJraWQiOiIwMTIwNzg3OTA4ODUxZDU3OTEzYjQwZTRiNzllMjFjZmNlM2NhNzMxZTMxMDA1MjUwOWVmZWU3NWI4OTY5ZWE3NWY0ZDBhIn0sInR5cGUiOiJyZXZva2UiLCJ2ZXJzaW9uIjoxfSwiY3RpbWUiOjE0MzIxNDU5NjMsImV4cGlyZV9pbiI6MTAwMDAwMDAsInByZXYiOiJkODQ1MjExYjJiYTIxZTg3MTQxZDUxNTUzYzQzZWNjMzcwMWZiZjM3MjAzZDYzZTViODJlZGU0YzBmYTAyM2JjIiwic2VxX3R5cGUiOjEsInNlcW5vIjozLCJ0YWciOiJzaWduYXR1cmUifaNzaWfEQIN1Yd5VY6JwixVs2xEJhQ0cz7PqHm00OCbuNFamehp6RfMzDkFK9L8YfUwkKoRQ89czbCeDw4my5vT/n4eutgyoc2lnX3R5cGUgo3RhZ80CAqd2ZXJzaW9uAQ==",
            "payload_hash": "3fd60b4b2317261b678b92b23dfc46e3d369b8d94a79b5b3dbfc2581de6a6741",
            "sig_id": "384d5b0d6ce5497364589b4fc3c907c0dc26b264aeb523b2212367cb96b170920f",
            "payload_json": "{\"body\":{\"key\":{\"eldest_kid\":\"0120364135dc024f6255d748dc8ee034e545530b0cd967e068e2ff03cdc1582fc7060a\",\"host\":\"keybase.io\",\"kid\":\"0120364135dc024f6255d748dc8ee034e545530b0cd967e068e2ff03cdc1582fc7060a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"revoke\":{\"kid\":\"0120787908851d57913b40e4b79e21cfce3ca731e310052509efee75b8969ea75f4d0a\"},\"type\":\"revoke\",\"version\":1},\"ctime\":1432145963,\"expire_in\":10000000,\"prev\":\"d845211b2ba21e87141d51553c43ecc3701fbf37203d63e5b82ede4c0fa023bc\",\"seq_type\":1,\"seqno\":3,\"tag\":\"signature\"}",
            "kid": "0120364135dc024f6255d748dc8ee034e545530b0cd967e068e2ff03cdc1582fc7060a",
            "ctime": 1432145963
        },
        {
            "seqno": 4,
            "prev": "3fd60b4b2317261b678b92b23dfc46e3d369b8d94a79b5b3dbfc2581de6a6741",
            "sig": "g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEgeHkIhR1XkTtA5LeeIc/OPKcx4xAFJQnv7nW4lp6nX00Kp3BheWxvYWTFBfl7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTIwMzY0MTM1ZGMwMjRmNjI1NWQ3NDhkYzhlZTAzNGU1NDU1MzBiMGNkOTY3ZTA2OGUyZmYwM2NkYzE1ODJmYzcwNjBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwNzg3OTA4ODUxZDU3OTEzYjQwZTRiNzllMjFjZmNlM2NhNzMxZTMxMDA1MjUwOWVmZWU3NWI4OTY5ZWE3NWY0ZDBhIiwidWlkIjoiNzRjMzhjZjdjZWI5NDdmNTYzMjA0NWQ4Y2E1ZDQ4MTkiLCJ1c2VybmFtZSI6Im1heDMyIn0sInNpYmtleSI6eyJraWQiOiIwMTIwN2E5M2FiMzUwMDNmNDgyNjJkZDU1OTNiZDAxYmU1M2QzYWNkYTBjMDlkYjVkYWM2MTg3YzZjNmUxZDYwYjllMDBhIiwicmV2ZXJzZV9zaWciOiJnNlJpYjJSNWhxaGtaWFJoWTJobFpNT3BhR0Z6YUY5MGVYQmxDcU5yWlhuRUl3RWdlcE9yTlFBL1NDWXQxVms3MEJ2bFBUck5vTUNkdGRyR0dIeHNiaDFndWVBS3AzQmhlV3h2WVdURkFpdDdJbUp2WkhraU9uc2lhMlY1SWpwN0ltVnNaR1Z6ZEY5cmFXUWlPaUl3TVRJd016WTBNVE0xWkdNd01qUm1OakkxTldRM05EaGtZemhsWlRBek5HVTFORFUxTXpCaU1HTmtPVFkzWlRBMk9HVXlabVl3TTJOa1l6RTFPREptWXpjd05qQmhJaXdpYUc5emRDSTZJbXRsZVdKaGMyVXVhVzhpTENKcmFXUWlPaUl3TVRJd056ZzNPVEE0T0RVeFpEVTNPVEV6WWpRd1pUUmlOemxsTWpGalptTmxNMk5oTnpNeFpUTXhNREExTWpVd09XVm1aV1UzTldJNE9UWTVaV0UzTldZMFpEQmhJaXdpZFdsa0lqb2lOelJqTXpoalpqZGpaV0k1TkRkbU5UWXpNakEwTldRNFkyRTFaRFE0TVRraUxDSjFjMlZ5Ym1GdFpTSTZJbTFoZURNeUluMHNJbk5wWW10bGVTSTZleUpyYVdRaU9pSXdNVEl3TjJFNU0yRmlNelV3TURObU5EZ3lOakprWkRVMU9UTmlaREF4WW1VMU0yUXpZV05rWVRCak1EbGtZalZrWVdNMk1UZzNZelpqTm1VeFpEWXdZamxsTURCaElpd2ljbVYyWlhKelpWOXphV2NpT201MWJHeDlMQ0owZVhCbElqb2ljMmxpYTJWNUlpd2lkbVZ5YzJsdmJpSTZNWDBzSW1OMGFXMWxJam94TkRNeU1UUTJNRFl6TENKbGVIQnBjbVZmYVc0aU9qRXdNREF3TURBd0xDSndjbVYySWpvaU0yWmtOakJpTkdJeU16RTNNall4WWpZM09HSTVNbUl5TTJSbVl6UTJaVE5rTXpZNVlqaGtPVFJoTnpsaU5XSXpaR0ptWXpJMU9ERmtaVFpoTmpjME1TSXNJbk5sY1Y5MGVYQmxJam94TENKelpYRnVieUk2TkN3aWRHRm5Jam9pYzJsbmJtRjBkWEpsSW4yamMybG54RURKTU42a1lCQTlvd1ZJS1dSZSs5czJPdGkwTkt0VHpqcnNkSWVmdmZiM3Q5MlBHRnptMEE1RDR0OTczTXNHVHFjZ3VJem9hak52eXNPVXc1b0FPVFFQcUhOcFoxOTBlWEJsSUtOMFlXZk5BZ0tuZG1WeWMybHZiZ0U9In0sInR5cGUiOiJzaWJrZXkiLCJ2ZXJzaW9uIjoxfSwiY3RpbWUiOjE0MzIxNDYwNjMsImV4cGlyZV9pbiI6MTAwMDAwMDAsInByZXYiOiIzZmQ2MGI0YjIzMTcyNjFiNjc4YjkyYjIzZGZjNDZlM2QzNjliOGQ5NGE3OWI1YjNkYmZjMjU4MWRlNmE2NzQxIiwic2VxX3R5cGUiOjEsInNlcW5vIjo0LCJ0YWciOiJzaWduYXR1cmUifaNzaWfEQDKHVTNDoxJs1UIrMI4CNhFzFdXtmbWe9F09gk2Rq5TbpMzSWGCCY1VKbHyzuSxEtX9rhBOcGkD8+O5Ktf809Ayoc2lnX3R5cGUgo3RhZ80CAqd2ZXJzaW9uAQ==",
            "payload_hash": "ff34e6ffa732ada20e3e2413f82fddb835f4c981c9acf43033da7ff8275333f5",
            "sig_id": "7eb99ec28f5cea7163e1e3dc05fae2690c7d040fd6b99a70ebe2266786e637500f",
            "payload_json": "{\"body\":{\"key\":{\"eldest_kid\":\"0120364135dc024f6255d748dc8ee034e545530b0cd967e068e2ff03cdc1582fc7060a\",\"host\":\"keybase.io\",\"kid\":\"0120787908851d57913b40e4b79e21cfce3ca731e310052509efee75b8969ea75f4d0a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"sibkey\":{\"kid\":\"01207a93ab35003f48262dd5593bd01be53d3acda0c09db5dac6187c6c6e1d60b9e00a\",\"reverse_sig\":\"g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEgepOrNQA/SCYt1Vk70BvlPTrNoMCdtdrGGHxsbh1gueAKp3BheWxvYWTFAit7ImJvZHkiOnsia2V5Ijp7ImVsZGVzdF9raWQiOiIwMTIwMzY0MTM1ZGMwMjRmNjI1NWQ3NDhkYzhlZTAzNGU1NDU1MzBiMGNkOTY3ZTA2OGUyZmYwM2NkYzE1ODJmYzcwNjBhIiwiaG9zdCI6ImtleWJhc2UuaW8iLCJraWQiOiIwMTIwNzg3OTA4ODUxZDU3OTEzYjQwZTRiNzllMjFjZmNlM2NhNzMxZTMxMDA1MjUwOWVmZWU3NWI4OTY5ZWE3NWY0ZDBhIiwidWlkIjoiNzRjMzhjZjdjZWI5NDdmNTYzMjA0NWQ4Y2E1ZDQ4MTkiLCJ1c2VybmFtZSI6Im1heDMyIn0sInNpYmtleSI6eyJraWQiOiIwMTIwN2E5M2FiMzUwMDNmNDgyNjJkZDU1OTNiZDAxYmU1M2QzYWNkYTBjMDlkYjVkYWM2MTg3YzZjNmUxZDYwYjllMDBhIiwicmV2ZXJzZV9zaWciOm51bGx9LCJ0eXBlIjoic2lia2V5IiwidmVyc2lvbiI6MX0sImN0aW1lIjoxNDMyMTQ2MDYzLCJleHBpcmVfaW4iOjEwMDAwMDAwLCJwcmV2IjoiM2ZkNjBiNGIyMzE3MjYxYjY3OGI5MmIyM2RmYzQ2ZTNkMzY5YjhkOTRhNzliNWIzZGJmYzI1ODFkZTZhNjc0MSIsInNlcV90eXBlIjoxLCJzZXFubyI6NCwidGFnIjoic2lnbmF0dXJlIn2jc2lnxEDJMN6kYBA9owVIKWRe+9s2Oti0NKtTzjrsdIefvfb3t92PGFzm0A5D4t973MsGTqcguIzoajNvysOUw5oAOTQPqHNpZ190eXBlIKN0YWfNAgKndmVyc2lvbgE=\"},\"type\":\"sibkey\",\"version\":1},\"ctime\":1432146063,\"expire_in\":10000000,\"prev\":\"3fd60b4b2317261b678b92b23dfc46e3d369b8d94a79b5b3dbfc2581de6a6741\",\"seq_type\":1,\"seqno\":4,\"tag\":\"signature\"}",
            "kid": "0120787908851d57913b40e4b79e21cfce3ca731e310052509efee75b8969ea75f4d0a",
            "ctime": 1432146063
        }
    ],
    "keys": [
        "0120364135dc024f6255d748dc8ee034e545530b0cd967e068e2ff03cdc1582fc7060a",
        "0120787908851d57913b40e4b79e21cfce3ca731e310052509efee75b8969ea75f4d0a",
        "01207a93ab35003f48262dd5593bd01be53d3acda0c09db5dac6187c6c6e1d60b9e00a"
    ],
    "uid": "74c38cf7ceb947f5632045d8ca5d4819",
    "username": "max32"
}
`,
	"simple_chain.json": `{
    "chain": [
        {
            "seqno": 1,
            "prev": null,
            "sig": "g6Rib2R5hqhkZXRhY2hlZMOpaGFzaF90eXBlCqNrZXnEIwEggg05gbQjjIO9nT6AnnoCfDEbP8W+/Qc9K/tIlLm+hrgKp3BheWxvYWTFASp7ImJvZHkiOnsia2V5Ijp7Imhvc3QiOiJrZXliYXNlLmlvIiwia2lkIjoiMDEyMDgyMGQzOTgxYjQyMzhjODNiZDlkM2U4MDllN2EwMjdjMzExYjNmYzViZWZkMDczZDJiZmI0ODk0YjliZTg2YjgwYSIsInVpZCI6Ijc0YzM4Y2Y3Y2ViOTQ3ZjU2MzIwNDVkOGNhNWQ0ODE5IiwidXNlcm5hbWUiOiJtYXgzMiJ9LCJ0eXBlIjoiZWxkZXN0IiwidmVyc2lvbiI6MX0sImN0aW1lIjoxNDMxOTc5MjE0LCJleHBpcmVfaW4iOjEwMDAwMDAwLCJwcmV2IjpudWxsLCJzZXFfdHlwZSI6MSwic2Vxbm8iOjEsInRhZyI6InNpZ25hdHVyZSJ9o3NpZ8RABjZFgXiUCNQR16g+RRL4xUtp+O2DI72+1gbMwoUJVlMxHGcXItlBFSFjfWGtOxU7//fYS0aVEH02q17IUCOXBqhzaWdfdHlwZSCjdGFnzQICp3ZlcnNpb24B",
            "payload_hash": "223e25fd456ea18cc39b608bac8b5f35302d240d5f95e6e42e5a85b1136ac5e7",
            "sig_id": "f0a633afc45db7d611a0e484e2db45800e87033f7c4bc1f62577bc1eba156f060f",
            "payload_json": "{\"body\":{\"key\":{\"host\":\"keybase.io\",\"kid\":\"0120820d3981b4238c83bd9d3e809e7a027c311b3fc5befd073d2bfb4894b9be86b80a\",\"uid\":\"74c38cf7ceb947f5632045d8ca5d4819\",\"username\":\"max32\"},\"type\":\"eldest\",\"version\":1},\"ctime\":1431979214,\"expire_in\":10000000,\"prev\":null,\"seq_type\":1,\"seqno\":1,\"tag\":\"signature\"}",
            "kid": "0120820d3981b4238c83bd9d3e809e7a027c311b3fc5befd073d2bfb4894b9be86b80a",
            "ctime": 1431979214
        }
    ],
    "keys": [
        "0120820d3981b4238c83bd9d3e809e7a027c311b3fc5befd073d2bfb4894b9be86b80a"
    ],
    "uid": "74c38cf7ceb947f5632045d8ca5d4819",
    "username": "max32"
}
`,
}
