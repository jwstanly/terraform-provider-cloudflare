---
page_title: "cloudflare_access_ca_certificate Resource - Cloudflare"
subcategory: ""
description: |-
  Cloudflare Access can replace traditional SSH key models with short-lived certificates issued to your users based on the token generated by their Access login.
---

# cloudflare_access_ca_certificate (Resource)

Cloudflare Access can replace traditional SSH key models with short-lived certificates issued to your users based on the token generated by their Access login.

~> It's required that an `account_id` or `zone_id` is provided and in
most cases using either is fine. However, if you're using a scoped
access token, you must provide the argument that matches the token's
scope. For example, an access token that is scoped to the "example.com"
zone needs to use the `zone_id` argument.


<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `application_id` (String) The Access Application ID to associate with the CA certificate.

### Optional

- `account_id` (String) The account identifier to target for the resource. Conflicts with `zone_id`.
- `zone_id` (String) The zone identifier to target for the resource. Conflicts with `account_id`.

### Read-Only

- `aud` (String) Application Audience (AUD) Tag of the CA certificate.
- `id` (String) The ID of this resource.
- `public_key` (String) Cryptographic public key of the generated CA certificate.

