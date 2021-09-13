# wled-tests

This was a simple repo to test a bug I was encountering where WLED requests would begin timing out after the fourth successive call. It was fixed by explicitly defining a client.timeout.
