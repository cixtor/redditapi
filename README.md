### Reddit API

[Reddit](https://www.reddit.com/) _(stylized as reddit, /ˈrɛdɪt/)_ is a social news aggregation, web content rating, and discussion website. Reddit's registered community members can submit content, such as text posts or direct links. Registered users can then vote submissions up or down to organize the posts and determine their position on the site's pages. The submissions with the most positive votes appear on the front page or the top of a category. Content entries are organized by areas of interest called "subreddits". The subreddit topics include news, science, gaming, movies, music, books, fitness, food, and image-sharing, among many others.

This project provides access to the [Reddit API](https://www.reddit.com/dev/api/) via OAuth2 for programs written using the Go programming language. The API implementation follows the [rules](https://github.com/reddit/reddit/wiki/API) that require the monitoring of the HTTP headers _"X-Ratelimit-Used"_ and _"X-Ratelimit-Remaining"_ to determine how many requests are allowed in the current minute which at the time of writing this document is limited to 60 requests. Once the limit is reached the API interface will stop contacting the API service until the next minute.

Notice that the HTTP authentication using OAuth2 is cached for 3,600 seconds setting the content of the Bearer key in a temporary file. When the lifetime of the key expires the API interface will request a new one and close the loop.

Be sure to check other [API Wrappers](https://github.com/reddit/reddit/wiki/API-Wrappers).
