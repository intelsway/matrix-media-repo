# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to 
[Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

* Add webp image support. Thanks @Sorunome!
* Add apng image support. Thanks @Sorunome!

### Changed

* Remove deprecated support for restricting uploads to certain mime types.

### Fixed

* GIFs now thumbnail correctly. Thanks @Sorunome!
* Fixed empty Content-Type header on retrieved remote media. Thanks @silkeh!
* Fixed various issues with IPv6 handling. Thanks @silkeh!
* Fixed high database usage for uploads when only one datastore is present.

## [1.1.3] - July 15th, 2020

### Added

* Added options to cache access tokens for users. This prevents excessive calls to `/account/whoami` on your homeserver, particularly for appservices.
* [Documentation](https://github.com/turt2live/matrix-media-repo/blob/master/docs/contrib/delegation.md) on how to set up delegation with the media repo and Traefik. Thanks @derEisele!

### Changed

* Deprecated support for restricting uploads to certain mime types, due to inability to make it work correctly with encrypted media.
* Removed deprecated `storagePaths` config option. Please use datastores.

### Fixed

* Fixed federation with some homeserver setups (delegation with ports). Thanks @MatMaul!
* Fixed the Synapse import script to not skip duplicated media. Thanks @jaywink!
* Fixed requests to IPv6 hosts. Thanks @MatMaul!
* Removed excessive calls to the database during upload.

## [1.1.2] - April 21st, 2020

### Fixed

* Fixed templates being corrupt in the Docker image.
* Fixed `REPO_CONFIG` environment variable not being respected for auxiliary binaries in the Docker image.

### Changed

* The Docker image now uses the migrations packed into the binary instead of the in-image ones.
* Reduced log spam when someone views an export.

## [1.1.1] - March 26th, 2020

### Added

* Added pprof endpoints for debugging performance. Only enabled with a `MEDIA_PPROF_SECRET_KEY` environment variable.

### Fixed

* Fixed a few very slow memory leaks when using S3 datastores.

## [1.1.0] - March 19th, 2020

### Added

* Added support for [MSC2448](https://github.com/matrix-org/matrix-doc/pull/2448).
* Added support for specifying a `region` to the S3 provider.
* Pass-through the `Accept-Language` header for URL previews, with options to set a default.
* Experimental support for IPFS.
* Consistent inclusion of a charset for certain text `Content-Type`s.
* New metrics for the cache composition reality (`media_cache_num_live_bytes_used` and `media_cache_num_live_items`).

### Fixed

* Fixed thumbnails producing the wrong result.
* Fixed `expireAfterDays` for thumbnails potentially deleting media under some conditions.
* Fixed a bug where items could be double-counted (but not double-stored) in the cache.
* Fixed the cache metrics reporting inaccurate values.
* Fixed a general memory leak in the cache due to inaccurate counting of items in the cache.

### Changed

* Updated to Go 1.14
* Updated the Grafana dashboard and moved it in-tree.

## [1.0.2] - March 3, 2020

### Added

* Added support for a `forKinds: ["all"]` option on datastores.

### Fixed

* Fixed a bug with the cache where it would never expire old entries unless it was pressed for space.
* Fixed a bug with the cache where the minimum cache time trigger would not work.

## [1.0.1] - February 27, 2020

### Fixed

* Fix a memory leak within the cache layers.

## [1.0.0] - January 4, 2020

### Added

* Compile assets (templates and migrations) into the binary for ease of deployment.
* Added binaries to make exports and imports easier.

### Fixed

* Fix error message when an invalid access token is provided.
* Fixed imports not starting in 1.0.0-rc.2.

## [1.0.0-rc.2] - January 3, 2020

### Fixed

* Fixed exports not starting in 1.0.0-rc.1.

## [1.0.0-rc.1] - December 29, 2019

### Added

* First ever release of matrix-media-repo.
* Deduplicate media from all sources.
* Support downloads, thumbnails, URL previews, identicons.
* Support for GDPR-style media exports.
* Support for importing from a previous export (for transferring data between repos).
* Admin utilities for clearing up space and undesirable content.
* Built-in S3 (and S3-like) support.
* Animated thumbnail generation.
* Importing media from an existing Synapse homeserver.
* Support for multiple datastores/locations to store different kinds of media.
* Federation for acquiring remote media.
* Media identification ([MSC2380](https://github.com/matrix-org/matrix-doc/pull/2380)).
* Support for cloning media to the local homeserver.
* Various other features that would be expected like maximum/minimum size controls, rate limiting, etc. Check out the
  sample config for a better idea of what else is possible.

[unreleased]: https://github.com/turt2live/matrix-media-repo/compare/v1.1.3...HEAD
[1.1.3]: https://github.com/turt2live/matrix-media-repo/compare/v1.1.2...v1.1.3
[1.1.2]: https://github.com/turt2live/matrix-media-repo/compare/v1.1.1...v1.1.2
[1.1.1]: https://github.com/turt2live/matrix-media-repo/compare/v1.1.0...v1.1.1
[1.1.0]: https://github.com/turt2live/matrix-media-repo/compare/v1.0.2...v1.1.0
[1.0.2]: https://github.com/turt2live/matrix-media-repo/compare/v1.0.1...v1.0.2
[1.0.1]: https://github.com/turt2live/matrix-media-repo/compare/v1.0.0...v1.0.1
[1.0.0]: https://github.com/turt2live/matrix-media-repo/compare/v1.0.0-rc.2...v1.0.0
[1.0.0-rc.2]: https://github.com/turt2live/matrix-media-repo/compare/v1.0.0-rc.1...v1.0.0-rc.2
[1.0.0-rc.1]: https://github.com/turt2live/matrix-media-repo/releases/tag/v1.0.0-rc.1
