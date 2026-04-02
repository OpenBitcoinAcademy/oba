## Random Key Generation

The earliest Bitcoin wallets generated each private key independently using a random number generator. Each key was unrelated to every other key. The wallet stored them all in a database file.

This approach had a serious backup problem. Every new key required a new backup. If a user generated 100 keys and backed up after key 50, keys 51 through 100 would be lost if the wallet file was corrupted.

Bitcoin Core originally pre-generated a pool of 100 keys to reduce the frequency of required backups. But the problem remained: independent random keys are hard to manage.
