Redirector
=================
This will simply redirect any request to this container to the website of your choosing. It will retain the relative path so something like http://www.myoriginalsite/test1 will redirect to http://www.myredirect/test1.

From scratch container available here:
https://hub.docker.com/r/impulse9489/redirector/

USAGE:
docker run -p PORT:PORT redirect CONTAINERPORT WEBSITE_REDIRECT

docker run -p 8082:8082 redirect 8082 https://www.tufts.edu