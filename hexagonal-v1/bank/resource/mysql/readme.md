## mysql

### Connect database with DBeaver
- You should add client option to your mysql-connector allowPublicKeyRetrieval=true to allow the client to automatically request the public key from the server. Note that allowPublicKeyRetrieval=True could allow a malicious proxy to perform a MITM attack to get the plaintext password, so it is False by default and must be explicitly enabled.

```sh
jdbc:mysql://localhost:3306/db?allowPublicKeyRetrieval=true&useSSL=false
```