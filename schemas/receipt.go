package schemas

// Receipt is a JSON schema.
const Receipt = `{
  "$schema": "http://json-schema.org/schema#",
  "$id": "https://schemas.licensezero.com/1.0.0-pre/receipt.json",
  "title": "license receipt",
  "comment": "A receipt represents confirmation of the sale of a software license.",
  "type": "object",
  "required": [
    "key",
    "license",
    "signature"
  ],
  "additionalProperties": false,
  "properties": {
    "key": {
      "title": "public signing key of the broker server",
      "$ref": "key.json"
    },
    "license": {
      "title": "license manifest",
      "type": "object",
      "required": [
        "form",
        "values"
      ],
      "properties": {
        "form": {
          "title": "license form",
          "type": "string",
          "minLength": 1
        },
        "values": {
          "type": "object",
          "required": [
            "api",
            "effective",
            "buyer",
            "seller",
            "sellerID",
            "offerID",
            "orderID"
          ],
          "additionalProperties": false,
          "properties": {
            "api": {
              "title": "license API",
              "$ref": "url.json"
            },
            "effective": {
              "title": "effective date",
              "$ref": "time.json"
            },
            "expires": {
              "title": "expiration date of the license",
              "$ref": "time.json"
            },
            "offerID": {
              "title": "offer identifier",
              "type": "string",
              "format": "uuid"
            },
            "orderID": {
              "title": "order identifier",
              "type": "string",
              "format": "uuid"
            },
            "price": {
              "title": "purchase price",
              "$ref": "price.json"
            },
            "buyer": {
              "title": "buyer",
              "comment": "The buyer is the one receiving the license.",
              "type": "object",
              "required": [
                "email",
                "jurisdiction",
                "name"
              ],
              "properties": {
                "email": {
                  "type": "string",
                  "format": "email"
                },
                "jurisdiction": {
                  "$ref": "jurisdiction.json"
                },
                "name": {
                  "$ref": "name.json",
                  "examples": [
                    "Joe Buyer"
                  ]
                }
              }
            },
            "seller": {
              "title": "seller",
              "comment": "The seller is the one giving the license.",
              "type": "object",
              "required": [
                "email",
                "jurisdiction",
                "name"
              ],
              "properties": {
                "email": {
                  "type": "string",
                  "format": "email"
                },
                "jurisdiction": {
                  "$ref": "jurisdiction.json"
                },
                "name": {
                  "$ref": "name.json",
                  "examples": [
                    "Joe Seller"
                  ]
                }
              }
            },
            "sellerID": {
              "title": "seller identifier",
              "type": "string",
              "format": "uuid"
            },
            "broker": {
              "title": "license broker",
              "comment": "information on the party that sold the license, such as an agent or reseller, if the seller did not sell the license themself",
              "type": "object",
              "required": [
                "email",
                "jurisdiction",
                "name",
                "website"
              ],
              "additionalProperties": false,
              "properties": {
                "email": {
                  "type": "string",
                  "format": "email"
                },
                "jurisdiction": {
                  "$ref": "jurisdiction.json"
                },
                "name": {
                  "$ref": "name.json",
                  "example": [
                    "Artless Devices LLC"
                  ]
                },
                "website": {
                  "$ref": "url.json"
                }
              }
            }
          }
        }
      }
    },
    "signature": {
      "title": "signature of the license broker server",
      "$ref": "signature.json"
    }
  }
}`
