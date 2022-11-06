# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [2.5.0] - 2022-11-06
### Removed
- Order field "Rejected"

## [2.4.0] - 2022-11-04
### Added
- Order field "Rejected"
### Removed
- Newrelic

## [2.3.0] - 2022-10-19
### Added
- Newrelic

## [2.2.1] - 2022-09-16
### Added
- AWS textract
- Convert unicode text to ASCII

## [2.2.0] - 2022-09-09
### Changed
- Store mongodb host in stage variable

## [2.1.8] - 2022-09-05
### Added
- Priority attribute to order

## [2.1.7] - 2022-09-04
### Added
- Product attribute vs2_flight

## [2.1.6] - 2022-09-04
### Added
- Added product attribute vs2_car_pickup_address

## [2.1.5] - 2022-09-04
### Added
- More product attribute

## [2.1.4] - 2022-09-03
### Added
- Product attribute email

## [2.1.3] - 2022-09-03
### Added
- Product attribute keys

## [2.1.2] - 2022-09-03
### Added
- Mailjet template ids for priority mail

## [2.1.1] - 2022-09-03
### Changed
- Avoid using default http client

## [2.1.0] - 2022-08-21
### Added
- Parsed arrival date

## [2.0.12] - 2022-08-21
### Added
- Parse woo date string

## [2.0.11] - 2022-08-21
### Fixed
- Wrong json and bson names

## [2.0.10] - 2022-08-20
### Added
- More config values

## [2.0.9] - 2022-08-20
### Added
- Config admin domain and customer domain

## [2.0.8] - 2022-08-20
### Fixed
- Wrong bson

## [2.0.7] - 2022-08-19
### Added
- Order InvoiceDocID

## [2.0.6] - 2022-08-19
### Changed
- Mailjet timeout to 300 seconds

## [2.0.5] - 2022-08-18
### Added
- Applicant "CancelReason"

## [2.0.4] - 2022-08-17
### Added
- SanitizeCVFileName

## [2.0.3] - 2022-08-13
### Added
- Correlation ID to SNS message

## [2.0.2] - 2022-08-13
### Fixed
- Unused parameters

## [2.0.1] - 2022-08-13
### Fixed
- Missing fields in request context

## [2.0.0] - 2022-08-13
### Changed
- New way to work with context & logger

## [1.4.0] - 2022-08-12
### Added
- Global config
- Pusher send notification 

## [1.3.0] - 2022-08-11
### Added
- Text package
- Method to remove non-ascii characters

## [1.2.2] - 2022-08-10
### Added
- ResultFile PassportNumber

## [1.2.1] - 2022-08-08
### Added
- S3 NewClientForRegion

## [1.2.0] - 2022-08-08
### Added
- Shared function to upload CV

## [1.1.6] - 2022-08-05
### Added
- Invoice title

## [1.1.5] - 2022-08-03
### Added
- mongodb collection "invoices"

## [1.1.4] - 2022-08-01
### Added
- DownloadFileWithTimeout

## [1.1.3] - 2022-08-01
### Changed
- applicant_passport_number

## [1.1.2] - 2022-07-31
### Changed
- InstrumentOrderData

## [1.1.1] - 2022-07-29
### Changed
- Uptake dependencies

## [1.1.0] - 2022-07-29
### Added
- Logger

## [1.0.16] - 2022-07-28
### Added
- Generate mongodb update object from JSON patch

## [1.0.15] - 2022-07-20
### Added
- Applicant VisaSent

## [1.0.13] - 2022-07-20
### Added
- Order AllVisaSent

## [1.0.12] - 2022-07-20
### Added
- S3 pre sign client

## [1.0.11] - 2022-07-20
### Added
- Disconnect mongodb client

## [1.0.10] - 2022-07-18
### Fixed
- Wrong package name

## [1.0.9] - 2022-07-18
### Added
- SNS client

## [1.0.8] - 2022-07-18
### Added
- S3 client

## [1.0.7] - 2022-07-18
### Added
- ErrorMessage for CollectionResult

## [1.0.6] - 2022-07-18
### Added
- MongoDB CollectionResult

## [1.0.5] - 2022-07-16
### Added
- Download file

## [1.0.4] - 2022-07-13
### Added
- Package "rest"

## [1.0.3] - 2022-07-13
### Added
- Add "VisaS3Key" to Applicant

## [1.0.2] - 2022-07-13
### Added
- Add "Metas"

## [1.0.1] - 2022-07-13
### Added
- Get order from woo
- Mongo collection "orders"

## [1.0.0] - 2022-07-12
### Added
- First version
