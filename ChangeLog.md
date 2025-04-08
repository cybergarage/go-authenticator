# Changelog  

## v1.0.5 (2025-04-XX)
- Changed tls.Config::SetTLSConfig() to disable TLS if the specified configuration is nil

## v1.0.4 (2025-01-18)
- Updated default credential authenticator to compare credentials based on variable type
- Add a default credential authenticator interface
 
## v1.0.3 (2025-01-16) 
- Updated Query interface
  - Added setter methods for the query parameters
  - Added the EncryptFunc() method
- Updated `go-sasl` package to v1.2.4

## v1.0.2 (2024-12-31)  
- Added aliases for `Conn` and `Mechanism` interfaces  
- Updated `go-sasl` package to v1.2.3  

## v1.0.1 (2024-12-28)  
- Updated `go-sasl` package to v1.2.2  

## v1.0.0 (2024-12-26)  
- Initial release  
