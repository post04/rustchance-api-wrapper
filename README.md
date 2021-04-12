# Explanation

This package is an API wrapper for the popular rust skin gambling website [rustchance](https://rustchance.com). I made this package because I wanted to play around with automatically doing things on rustchance but didn't see **ANY** public API documentation or **ANY** talk about one.

# TODO

- [ ] Search for more authorization required socket data
- [ ] Search for more authorization required http/s endpoints
- [ ] Add all /api/account/history/ API urls to the wrapper

# Disclaimer

In the [rustchance TOS](https://rustchance.com/page/tos) I couldn't find anything about account automization being disallowed, nor anything about using the api in this way being disallowed. With that being said if they see this project they may add that to their TOS so please do use it at your own risk.

# Note

The rustchance API is not at all documented or even meant for this type of use (non web), with that being said, rustchance has some obscure variable names that I had to go through and figure out myself. With that this wrapper is subject to human error and if you find any problems in the wrapper or in my understanding of the API please open a github issue or make a pull request!
