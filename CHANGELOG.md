# Changelog

## 0.1.0-alpha.4 (2026-05-18)

Full Changelog: [v0.1.0-alpha.3...v0.1.0-alpha.4](https://github.com/HubSpot/hubspot-sdk-go/compare/v0.1.0-alpha.3...v0.1.0-alpha.4)

### Features

* **api:** manual updates ([b09ae91](https://github.com/HubSpot/hubspot-sdk-go/commit/b09ae91f474f81f1a664169a37994a7cf91408c8))

## 0.1.0-alpha.3 (2026-05-14)

Full Changelog: [v0.1.0-alpha.2...v0.1.0-alpha.3](https://github.com/HubSpot/hubspot-sdk-go/compare/v0.1.0-alpha.2...v0.1.0-alpha.3)

### Features

* **client:** optimize json encoder for internal types ([3f4f4f1](https://github.com/HubSpot/hubspot-sdk-go/commit/3f4f4f113cd0950beec2df3156bdd1ac0772ad2f))


### Bug Fixes

* **go:** avoid panic when http.DefaultTransport is wrapped ([56d4ce1](https://github.com/HubSpot/hubspot-sdk-go/commit/56d4ce18b0212b1ff14ab5330ab3398ee4d50a0a))


### Chores

* redact api-key headers in debug logs ([14847b4](https://github.com/HubSpot/hubspot-sdk-go/commit/14847b43b20668ac21eaf1d00a804cb20c64ca14))

## 0.1.0-alpha.2 (2026-05-07)

Full Changelog: [v0.1.0-alpha.1...v0.1.0-alpha.2](https://github.com/HubSpot/hubspot-sdk-go/compare/v0.1.0-alpha.1...v0.1.0-alpha.2)

### Features

* **api:** manual updates ([8c16968](https://github.com/HubSpot/hubspot-sdk-go/commit/8c16968fc9c2a6093c6d3b094f1e5ddc4db1e768))

## 0.1.0-alpha.1 (2026-05-06)

### Features

* **api:** manual updates ([f4c9de8](https://github.com/HubSpot/hubspot-sdk-go/commit/f4c9de8aba1ffb8e305eb0921ebf1a3d38d7db28))
* **api:** manual updates ([e9a9db6](https://github.com/HubSpot/hubspot-sdk-go/commit/e9a9db67b784f801651a864a8d808e9465db8be8))
* **api:** manual updates ([a03b901](https://github.com/HubSpot/hubspot-sdk-go/commit/a03b901bd859b6ba09e2180d1560b752fdea6904))
* **api:** manual updates ([cb3a16b](https://github.com/HubSpot/hubspot-sdk-go/commit/cb3a16b8f1f84f89da2834eae7bc5f1be4ad4841))
* **api:** manual updates ([7ac3dce](https://github.com/HubSpot/hubspot-sdk-go/commit/7ac3dce2e99ba25ba3ef22780c601f54b663e054))
* **api:** manual updates ([e1f31a3](https://github.com/HubSpot/hubspot-sdk-go/commit/e1f31a3e3a7f78806225e85704ada6a90fe896d0))
* **api:** manual updates ([d7e03d0](https://github.com/HubSpot/hubspot-sdk-go/commit/d7e03d0683a106b91d08660fff9e6219e8f71563))
* **api:** manual updates ([2ad7120](https://github.com/HubSpot/hubspot-sdk-go/commit/2ad71205c848ca882d5006f5575449997ea30b70))
* **api:** manual updates ([cc57da8](https://github.com/HubSpot/hubspot-sdk-go/commit/cc57da873c44c0b8faee9c6c578b3406f2f63b25))
* **api:** manual updates ([7d42772](https://github.com/HubSpot/hubspot-sdk-go/commit/7d427724e425dd0487f31be9ee887f7106606438))
* **api:** manual updates ([1387927](https://github.com/HubSpot/hubspot-sdk-go/commit/1387927fe717595bbbc782927a208578addaf437))
* **api:** manual updates ([d1bec2f](https://github.com/HubSpot/hubspot-sdk-go/commit/d1bec2f861d586f24d0b69383ea5735ea8bf95cd))
* **api:** manual updates ([9049db4](https://github.com/HubSpot/hubspot-sdk-go/commit/9049db4c56f8ef3e01e4d62cc0b2bdb2a19f3b31))
* **api:** manual updates ([eba438a](https://github.com/HubSpot/hubspot-sdk-go/commit/eba438a6c47fad601b3c6cbe3c5275a550ddb7c7))
* **api:** switch to 2026-03 ([442c099](https://github.com/HubSpot/hubspot-sdk-go/commit/442c099ff8d92ab32f334a9242477aa3ebf903af))
* **client:** add a convenient param.SetJSON helper ([0646173](https://github.com/HubSpot/hubspot-sdk-go/commit/0646173335fb312c98ce94657db58149fccd4771))
* **encoder:** support bracket encoding form-data object members ([ccd937c](https://github.com/HubSpot/hubspot-sdk-go/commit/ccd937c89bc796a0ab86b890d4f96e05aefa0aa9))
* **go:** add default http client with timeout ([2459870](https://github.com/HubSpot/hubspot-sdk-go/commit/2459870da6d86b82521cd2aff698e902fe679832))
* **internal:** support comma format in multipart form encoding ([af9585f](https://github.com/HubSpot/hubspot-sdk-go/commit/af9585f72ca750836c1b959ac385027b248504cb))
* support setting headers via env ([5b136a7](https://github.com/HubSpot/hubspot-sdk-go/commit/5b136a79f4e2a8d9cee8e4ffc9651a096b9ef1b3))


### Bug Fixes

* **client:** correctly specify Accept header with */* instead of empty ([ac9d395](https://github.com/HubSpot/hubspot-sdk-go/commit/ac9d39519f0ed43309f798b805560eca3c0f7f94))
* **docs:** add missing pointer prefix to api.md return types ([1a61305](https://github.com/HubSpot/hubspot-sdk-go/commit/1a61305e2849d4ec37e85c81cf98c07a4d191adb))
* **encoder:** correctly serialize NullStruct ([2e93bd6](https://github.com/HubSpot/hubspot-sdk-go/commit/2e93bd69b55ef25aa795abdc056e9bda8cc7992c))
* fix issue with unmarshaling in some cases ([e476d30](https://github.com/HubSpot/hubspot-sdk-go/commit/e476d30fea77e461013acea979379082e5aa2c33))
* **mcp:** correct code tool API endpoint ([c539ea5](https://github.com/HubSpot/hubspot-sdk-go/commit/c539ea57071fbbc6e0393f2db698dcccc9f491f9))
* prevent duplicate ? in query params ([519cc39](https://github.com/HubSpot/hubspot-sdk-go/commit/519cc39f58102fbe67d8779d2b7fa5511f3f7b26))
* rename param to avoid collision ([7ad3b32](https://github.com/HubSpot/hubspot-sdk-go/commit/7ad3b32f3227b43143a40bcb2617075ea55240df))
* skip usage tests that don't work with Prism ([f378d4f](https://github.com/HubSpot/hubspot-sdk-go/commit/f378d4fddcf71963814b664dfe801bd7b120a56a))


### Chores

* add float64 to valid types for RegisterFieldValidator ([e28bcad](https://github.com/HubSpot/hubspot-sdk-go/commit/e28bcad84ab6360385428e1c29a3fa33b228bdb1))
* avoid embedding reflect.Type for dead code elimination ([096c9fe](https://github.com/HubSpot/hubspot-sdk-go/commit/096c9fe27f326991ad15a8454cb8771c9af625ad))
* bump @stdy/cli to 0.15.3 ([9734d4d](https://github.com/HubSpot/hubspot-sdk-go/commit/9734d4db76382836f06eec65825c30318807a842))
* bump gjson version ([a38b871](https://github.com/HubSpot/hubspot-sdk-go/commit/a38b871a6eefb8ebc986bf68db2296ac9e890379))
* **ci:** skip lint on metadata-only changes ([6679af0](https://github.com/HubSpot/hubspot-sdk-go/commit/6679af008131ba93547a77fa12fe2608120758e8))
* **ci:** support opting out of skipping builds on metadata-only commits ([19a5aeb](https://github.com/HubSpot/hubspot-sdk-go/commit/19a5aebf24713d9c476438fd6fe0dcc4bd38d75e))
* **client:** fix multipart serialisation of Default() fields ([75cc1ab](https://github.com/HubSpot/hubspot-sdk-go/commit/75cc1abd56c523e10ce77f490ff393feee92634a))
* elide duplicate aliases ([dd8323c](https://github.com/HubSpot/hubspot-sdk-go/commit/dd8323cb7446b232e7a3783c327a2a94ed2322d8))
* fix empty interfaces ([90f8387](https://github.com/HubSpot/hubspot-sdk-go/commit/90f83879830bc4ced90389db0f43d0ab6a8eab4f))
* **internal:** bump steady CLI version ([6ff1c60](https://github.com/HubSpot/hubspot-sdk-go/commit/6ff1c6046bdf44cb1b65ef5b14ee3090dea575a7))
* **internal:** codegen related update ([e667362](https://github.com/HubSpot/hubspot-sdk-go/commit/e667362375625d87310d48703e55c086f4934c2b))
* **internal:** codegen related update ([1748f67](https://github.com/HubSpot/hubspot-sdk-go/commit/1748f67b260a2a6c894b22e85badb21bb3a75985))
* **internal:** grammar fix (it's -&gt; its) ([f997d9e](https://github.com/HubSpot/hubspot-sdk-go/commit/f997d9ede3d8b186ed1fa1c7321436414d4f508f))
* **internal:** more robust bootstrap script ([fbcfd12](https://github.com/HubSpot/hubspot-sdk-go/commit/fbcfd12666a5d0b9ce73eb2182af9b3163e31e66))
* **internal:** support default value struct tag ([aada93a](https://github.com/HubSpot/hubspot-sdk-go/commit/aada93a865f9a2840ab2abeecde67bf57ef9ecc8))
* **internal:** update `actions/checkout` version ([d41ae2a](https://github.com/HubSpot/hubspot-sdk-go/commit/d41ae2afbafcd08b176b5f75eae1548fac47e5c0))
* **internal:** update gitignore ([b0a5134](https://github.com/HubSpot/hubspot-sdk-go/commit/b0a51349f6dd6b9b17f7b9aa755ce0f1f3fa7d86))
* remove unnecessary error check for url parsing ([47f9bf0](https://github.com/HubSpot/hubspot-sdk-go/commit/47f9bf00d99f1770c132b9d933a23e2daf61b9dd))
* **tests:** configure mock server correctly ([1c66b5d](https://github.com/HubSpot/hubspot-sdk-go/commit/1c66b5de65bd318e379a59094445afd0885a2889))
* **tests:** switch from prism to steady ([a1be7b4](https://github.com/HubSpot/hubspot-sdk-go/commit/a1be7b473d023355bd6d1661d39dd525a4eab87a))
* update docs for api:"required" ([0608736](https://github.com/HubSpot/hubspot-sdk-go/commit/0608736862160b4f180c024a47a9963718830e34))


### Documentation

* split `api.md` by standalone resources ([937dbeb](https://github.com/HubSpot/hubspot-sdk-go/commit/937dbeba67e6951f1e868419953ce3e465c9109b))
